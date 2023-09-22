package env

import (
	"context"

	"github.com/docker/docker/client"
	"go.opentelemetry.io/otel"

	_ "embed"
)

func MockBuild(envID, buildID, provisionEnvScript string) {
	ctx := context.Background()

	tracer := otel.Tracer("test")

	client, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	contextsPath := "/mnt/disks/docker-contexts/v1"
	registry := "us-central1-docker.pkg.dev/e2b-prod/custom-environments"
	envsPath := "/mnt/disks/fc-envs/v1"
	kernelImagePath := "/fc-vm/vmlinux.bin"
	firecrackerBinaryPath := "/usr/bin/firecracker"
	envsPipelinePath := "/mnt/disks/envs-pipeline"
	envdName := "envd"
	contextFileName := "context.tar.gz"
	vCPUCount := int64(1)
	memoryMB := int64(512)
	diskSizeMB := int64(512)

	e := Env{
		BuildID:               buildID,
		EnvID:                 envID,
		EnvsPath:              envsPath,
		VCpuCount:             vCPUCount,
		MemoryMB:              memoryMB,
		DockerContextsPath:    contextsPath,
		DockerRegistry:        registry,
		KernelImagePath:       kernelImagePath,
		DiskSizeMB:            diskSizeMB,
		FirecrackerBinaryPath: firecrackerBinaryPath,
		ProvisionScript:       provisionEnvScript,
		EnvsPipelinePath:      envsPipelinePath,
		EnvdName:              envdName,
		ContextFileName:       contextFileName,
	}

	err = e.Build(ctx, tracer, client)
	if err != nil {
		panic(err)
	}
}