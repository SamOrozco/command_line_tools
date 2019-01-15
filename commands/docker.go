package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"os/exec"
	"strings"
)

func NewDockerCommand() *cobra.Command {
	containerCommand := &cobra.Command{
		Use:   "cntr",
		Short: "Docker Container utils",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
		},
	}

	containerKillFirstCommand := &cobra.Command{
		Use:   "kf",
		Short: "Kill first running docker container",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			println(fmt.Sprintf("checking for running container"))
			println("--------------------------------------")
			containerNames, err := getRunningContainers()
			if err != nil {
				panic(err)
			}
			if len(containerNames) < 1 {
				println("no running containers")
			}
			killContainer(containerNames[0])
		},
	}

	containerKillAll := &cobra.Command{
		Use:   "ka",
		Short: "Kill all running docker containers",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			println(fmt.Sprintf("checking for running containers"))
			println("--------------------------------------")
			containerNames, err := getRunningContainers()
			if err != nil {
				panic(err)
			}
			if len(containerNames) < 1 {
				println("no running containers")
			}
			killContainer(containerNames...)
		},
	}

	containerCommand.AddCommand(containerKillFirstCommand, containerKillAll)
	return containerCommand
}

const containerIdIndex = 0

func getRunningContainers() ([]string, error) {
	shellCommand := exec.Command("docker", "container", "ls")
	output, err := shellCommand.Output()
	if err != nil {
		println(err.Error())
		return make([]string, 0), nil
	}
	// we know the lsof command has a header so we want to skip the first row
	stringOutput := string(output)
	rows := strings.Split(stringOutput, "\n")
	// removing first row because we know it is headers
	rows = rows[1 : len(rows)-1]
	if len(rows) < 1 {
		println("no running containers")
		return make([]string, 0), nil
	}
	containerIds := make([]string, 0)
	for _, row := range rows {
		columns := strings.Fields(row)
		containerIds = append(containerIds, columns[containerIdIndex])
	}

	return containerIds, nil
}

func killContainer(containerIds ... string) {
	for _, id := range containerIds {
		shellCommand := exec.Command("docker", "kill", id)
		output, err := shellCommand.Output()
		if err != nil {
			println(err.Error())
		}
		println(string(output))
	}
}
