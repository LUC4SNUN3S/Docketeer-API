package repositories

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

type ContainerRepository struct {
	cli *client.Client
}

func NewContainerRepository() (*ContainerRepository, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, nil)
	if err != nil {
		return nil, err
	}

	return &ContainerRepository{cli: cli}, nil
}

func (r *ContainerRepository) ListContainers() ([]types.Container, error) {
	return r.cli.ContainerList(context.Background(), container.ListOptions{All: false})
}
