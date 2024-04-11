package consul

import (
	"context"
	"fmt"
	"github.com/e2b-dev/infra/packages/shared/pkg/telemetry"
	consul "github.com/hashicorp/consul/api"
	"os"
)

var consulToken = os.Getenv("CONSUL_TOKEN")

func New(ctx context.Context) (*consul.Client, error) {
	config := consul.DefaultConfig()
	config.Token = consulToken

	consulClient, err := consul.NewClient(config)
	if err != nil {
		errMsg := fmt.Errorf("failed to initialize Consul client: %w", err)
		telemetry.ReportCriticalError(ctx, errMsg)

		return nil, errMsg
	}
	return consulClient, nil
}
