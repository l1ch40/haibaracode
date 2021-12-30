package docker

import (
	"context"
	"errors"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"haibaracode/dto"
)

type DockerService struct {
}

func (ds DockerService) Create(dto dto.ContainerDto) (interface{}, error) {
	ctx := context.Background()
	cli, err := client.NewClient("tcp://localhost:1234", "v1.12", nil, nil)
	if err != nil {
		return nil, errors.New("创建容器失败。")
	}

	_, err = cli.ImagePull(ctx, dto.Image, types.ImagePullOptions{})
	if err != nil {
		return nil, errors.New("创建容器失败。")
	}

	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: dto.Image,
	}, nil, nil, nil, "")
	if err != nil {
		return nil, errors.New("创建容器失败。")
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		return nil, errors.New("创建容器失败。")
	}

	return resp.ID, nil
}

func (ds DockerService) Delete(dto dto.ContainerDto) error {
	ctx := context.Background()
	cli, err := client.NewClient("tcp://localhost:1234", "v1.12", nil, nil)
	if err != nil {
		return errors.New("删除容器失败。")
	}
	if err := cli.ContainerRemove(ctx, dto.CID, types.ContainerRemoveOptions{}); err != nil {
		fmt.Println(err)
		return errors.New("删除容器失败。")
	}
	return nil
}
