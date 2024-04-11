package handlers

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	consulapi "github.com/hashicorp/consul/api"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"

	"github.com/e2b-dev/infra/packages/orchestrator/internal/api"
	"github.com/e2b-dev/infra/packages/orchestrator/internal/consul"
	"github.com/e2b-dev/infra/packages/orchestrator/internal/instance"
	"github.com/e2b-dev/infra/packages/shared/pkg/smap"
	"github.com/e2b-dev/infra/packages/shared/pkg/telemetry"
)

const shortNodeIDLength = 8
const fcVersionsDir = "/fc-versions"
const kernelDir = "/fc-kernels"
const kernelMountDir = "/fc-vm"
const kernelName = "vmlinux.bin"
const uffdBinaryName = "uffd"
const fcBinaryName = "firecracker"

var (
	nodeID   = os.Getenv("NODE_ID")
	clientID = nodeID[:shortNodeIDLength]
)

type APIStore struct {
	Ctx       context.Context
	instances *smap.Map[*instance.Instance]
	dns       *instance.DNS
	tracer    trace.Tracer
	consul    *consulapi.Client
}

func NewAPIStore() *APIStore {
	fmt.Println("Initializing API store")

	ctx := context.Background()

	dns, err := instance.NewDNS()
	if err != nil {
		panic(err)
	}

	consulClient, err := consul.New(ctx)
	if err != nil {
		panic(err)
	}

	return &APIStore{
		Ctx:       ctx,
		tracer:    otel.Tracer("orchestrator"),
		consul:    consulClient,
		dns:       dns,
		instances: smap.New[*instance.Instance](),
	}
}

func (a *APIStore) Close() {}

// This function wraps sending of an error in the Error format, and
// handling the failure to marshal that.
func (a *APIStore) sendAPIStoreError(c *gin.Context, code int, message string) {
	apiErr := api.Error{
		Code:    int32(code),
		Message: message,
	}

	c.Error(fmt.Errorf(message))
	c.JSON(code, apiErr)
}

func (a *APIStore) GetHealth(c *gin.Context) {
	c.String(http.StatusOK, "Health check successful")
}

func (a *APIStore) PostSandboxes(c *gin.Context) {
	ctx := c.Request.Context()
	tracer := otel.Tracer("create")

	body, err := parseBody[api.Sandbox](ctx, c)
	if err != nil {
		fmt.Println("Error parsing request body")

		a.sendAPIStoreError(c, http.StatusBadRequest, "failed to parse request body")
		return
	}

	instance, err := instance.NewInstance(
		ctx,
		tracer,
		a.consul,
		&instance.InstanceConfig{
			EnvID:                 body.EnvID,
			NodeID:                nodeID,
			InstanceID:            body.InstanceID,
			TraceID:               body.TraceID,
			TeamID:                body.TeamID,
			KernelVersion:         body.KernelVersion,
			EnvsDisk:              body.EnvsDisk,
			KernelsDir:            kernelDir,
			KernelMountDir:        kernelMountDir,
			KernelName:            kernelName,
			HugePages:             body.HugePages,
			UFFDBinaryPath:        filepath.Join(fcVersionsDir, body.FirecrackerVersion, uffdBinaryName),
			FirecrackerBinaryPath: filepath.Join(fcVersionsDir, body.FirecrackerVersion, fcBinaryName),
			AllocID:               "alloc-id",
		},
		a.dns,
		&body,
	)
	if err != nil {
		errMsg := fmt.Errorf("failed to create instance: %w", err)
		telemetry.ReportCriticalError(ctx, errMsg)

		a.sendAPIStoreError(c, http.StatusInternalServerError, "failed to create instance")

		return
	}

	body.ClientID = &clientID

	a.instances.Insert(body.InstanceID, instance)

	go func() {
		tracer := otel.Tracer("close")
		defer instance.CleanupAfterFCStop(context.Background(), tracer, a.consul, a.dns)
		defer a.instances.Remove(body.InstanceID)

		err := instance.FC.Wait()
		if err != nil {
			errMsg := fmt.Errorf("failed to wait for FC: %w", err)
			telemetry.ReportCriticalError(ctx, errMsg)
		}
	}()

	c.JSON(http.StatusCreated, api.NewSandbox{
		SandboxID: body.InstanceID,
		ClientID:  clientID,
	})
}

func (a *APIStore) GetSandboxes(c *gin.Context) {
	var sandboxes []*api.Sandbox

	for _, instance := range a.instances.Items() {
		sandboxes = append(sandboxes, instance.Request)
	}

	c.JSON(http.StatusOK, sandboxes)
}

func (a *APIStore) DeleteSandboxesSandboxID(c *gin.Context, sandboxID string) {
	ctx := c.Request.Context()

	tracer := otel.Tracer("delete")

	instance, ok := a.instances.Get(sandboxID)
	if !ok {
		a.sendAPIStoreError(c, http.StatusNotFound, "sandbox not found")
		return
	}

	err := instance.FC.Stop(ctx, tracer)
	defer instance.CleanupAfterFCStop(ctx, tracer, a.consul, a.dns)
	if err != nil {
		errMsg := fmt.Errorf("failed to stop FC: %w", err)

		telemetry.ReportCriticalError(ctx, errMsg)
		a.sendAPIStoreError(c, http.StatusInternalServerError, "failed to stop FC")

		return
	}

	c.Status(http.StatusNoContent)
}
