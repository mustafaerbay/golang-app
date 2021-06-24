package cmd

import (
	"context"
	"fmt"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"

	// "github.com/mustafaerbay/golang-app/functions"
	"github.com/spf13/cobra"
)

var getContainersCmd = &cobra.Command{
	Use:   "Cont",
	Short: "Get Containers",
	Long:  `This command Containers`,
	Run: func(cmd *cobra.Command, args []string) {
		getContainers()
	},
}

func init()  {
	rootCmd.AddCommand(getContainersCmd)
}


func getContainers() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Println("container.ID		container.Names		container.SizeRw	container.SizeRootFs")
	// elements := [][]string{}
	for _, container := range containers {
		cont := container.ID
		cont10 := cont[0:9]
		fmt.Println(
			cont10,"	",
			strings.Trim(strings.Join(container.Names,""), "[/]"),"	",
			container.SizeRw,
			container.SizeRootFs,
			// functions.ByteCountSI(container.SizeRw),
			// functions.ByteCountSI(container.SizeRootFs),
		)
		
		// elements = append(elements, container.Names)
		// fmt.Print(elements)
	}
	fmt.Println("*******************")
	fmt.Println("container.SizeRw",containers[0].SizeRw)
	fmt.Println("container.SizeRootFs",containers[0].SizeRootFs)

	
}