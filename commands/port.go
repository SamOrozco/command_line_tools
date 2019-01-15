package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"os/exec"
	"strings"
)

type LocalApplication struct {
	Name string
	PID  string
}

func NewPortCommand() *cobra.Command {
	portCommand := &cobra.Command{
		Use:   "port",
		Short: "port command utils",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
		},
	}

	portShowCommand := &cobra.Command{
		Use:   "show",
		Short: "Show applications blocking the given port",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				println("No port argument specified")
			}
			port := args[0]
			println()
			println(fmt.Sprintf("Checking for port %s", port))
			println("--------------------------------------")
			applications, err := getApplicationIdsForPort(port)
			if err != nil {
				panic(err)
			}
			if len(applications) < 1 {
				println("No applications using specified port")
			}
 			for _, val := range applications {
				println(fmt.Sprintf("Application Name: %s", val.Name))
				println(fmt.Sprintf("Application PID: %s", val.PID))
				println("--------------------------------------")
			}
		},
	}


	portKillCommand := &cobra.Command{
		Use:   "kill",
		Short: "Kills all applications with the specified port",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			if args == nil || len(args) < 1 {
				println("No port argument specified")
			}
			port := args[0]
			println()
			println(fmt.Sprintf("Kill applications for port %s", port))
			println("--------------------------------------")
			applications, err := getApplicationIdsForPort(port)
			if err != nil {
				panic(err)
			}
			if len(applications) < 1 {
				println("No applications using specified port")
			}
			for _, val := range applications {
				println(fmt.Sprintf("Application PID: %s", val.PID))
				_ = killApplicationByPID(val.PID)
			}
		},
	}

	portCommand.AddCommand(portShowCommand, portKillCommand)
	return portCommand
}

func getApplicationIdsForPort(port string) ([]LocalApplication, error) {
	shellCommand := exec.Command("lsof", fmt.Sprintf("-i:%s", port))
	output, err := shellCommand.Output()
	if err != nil {
		return make([]LocalApplication, 0), nil
	}

	// we know the lsof command has a header so we want to skip the first row
	stringOutput := string(output)
	rows := strings.Split(stringOutput, "\n")
	if len(rows) <= 1 {
		println("No applications running on that port")
	}

	responses := rows[1 : len(rows)-1]
	applications := make([]LocalApplication, 0)
	for _, resp := range responses {
		pieces := strings.Fields(resp)
		name := pieces[0]
		pid := pieces[1]
		applications = append(applications, LocalApplication{Name:name, PID:pid})
	}
	return applications, nil
}


func killApplicationByPID(pid string) string {
	cmd := exec.Command("kill", "-9", pid)
	val, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	return string(val)
}
