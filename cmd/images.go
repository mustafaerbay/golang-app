package cmd

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	docker "github.com/fsouza/go-dockerclient"
	"github.com/mustafaerbay/golang-app/functions"
	"github.com/spf13/cobra"
)

var getImagesCmd = &cobra.Command{
	Use:   "images",
	Short: "Get images",
	Long:  `This command images`,
	Run: func(cmd *cobra.Command, args []string) {
		getImages()
	},
}

func init()  {
	rootCmd.AddCommand(getImagesCmd)
}

func getImages() {
	client, err := docker.NewClientFromEnv()
	if err != nil {
		panic(err)
	}
	imgs, err := client.ListImages(docker.ListImagesOptions{All: false})
	if err != nil {
		panic(err)
	}
	for _, img := range imgs {
		fmt.Println("ID: ", img.ID)
		fmt.Println("RepoTags: ", img.RepoTags)
		fmt.Println("Created: ", img.Created)
		fmt.Println("Size: ", img.Size)
		fmt.Println("VirtualSize: ", img.VirtualSize)
		fmt.Println("VirtualSize: ", functions.ByteCountSI(img.VirtualSize))
		fmt.Println("ParentId: ", img.ParentID)
		fmt.Println("")
	}
}

func getImages2() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		fmt.Println(container.ID)
	}
}