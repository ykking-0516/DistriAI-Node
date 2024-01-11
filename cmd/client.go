package cmd

import (
	"DistriAI-Node/chain/subscribe"
	"DistriAI-Node/core_task"
	"DistriAI-Node/docker"
	logs "DistriAI-Node/utils/log_utils"
	"fmt"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/urfave/cli"
)

var ClientCommand = cli.Command{
	Name:  "node",
	Usage: "Starting or terminating a node program.",
	Subcommands: []cli.Command{
		{
			Name:  "start",
			Usage: "Upload hardware configuration and initiate listening events.",
			Action: func(c *cli.Context) error {
				distriWrapper, hwInfo, chainInfo, err := core_task.GetDistri(true)
				if err != nil {
					logs.Error(err.Error())
					return nil
				}

				machine, err := distriWrapper.GetMachine()
				if err != nil {
					logs.Error(fmt.Sprintf("Error: %v", err))
					return nil
				}

				if machine.Metadata == "" {
					logs.Normal("Machine does not exist")
					hash, err := distriWrapper.AddMachine(*hwInfo)
					if err != nil {
						logs.Error(fmt.Sprintf("Error block : %v, msg : %v\n", hash, err))
						return nil
					}
				} else {
					logs.Normal("Machine already exists")
					if machine.Status.String() == "Renting" {
						logs.Error(fmt.Sprintf("Machine is Renting, status: %v", machine.Status.String()))
						return nil
					}
				}

				subscribeBlocks := subscribe.NewSubscribeWrapper(chainInfo)

				for {
					time.Sleep(1 * time.Second)

					logs.Normal("=============== Start subscription")
					order, err := subscribeBlocks.SubscribeEvents(hwInfo)
					logs.Normal("=============== End subscription")
					if err != nil {
						logs.Error(err.Error())
						time.Sleep(3 * time.Minute)
						logs.Normal("Restart subscription")
						continue
					}

					if order.Metadata == "" {
						logs.Error("order metadata is empty")
						continue
					}

					if order.Status.String() != "Training" {
						logs.Error(fmt.Sprintf("Order status is not training, status: %v", order.Status.String()))
						continue
					}

					logs.Result(fmt.Sprintf("Start order. OrderAccount: %v", subscribeBlocks.ProgramDistriOrder))
					spew.Dump(order)

					isGPU := false
					if hwInfo.GPUInfo.Number > 0 {
						isGPU = true
					}
					containerID, err := docker.RunWorkspaceContainer(isGPU)
					if err != nil {
						logs.Error(fmt.Sprintln("RunWorkspaceContainer error: ", err))
						return nil
					}
					if core_task.StartTimer(distriWrapper, order) {
						err = core_task.OrderComplete(distriWrapper, order.Metadata, isGPU, containerID)
					} else {
						err = core_task.OrderFailed(distriWrapper, order.Metadata, order.Buyer, containerID)
					}
					if err != nil {
						logs.Error(fmt.Sprintln("Order end error: ", err))
						return nil
					}
				}
			},
		},
		{
			Name:  "stop",
			Usage: "Stop the client.",
			Action: func(c *cli.Context) error {
				distriWrapper, hwInfo, _, err := core_task.GetDistri(false)
				if err != nil {
					logs.Error(err.Error())
					return nil
				}

				machine, err := distriWrapper.GetMachine()
				if err != nil {
					logs.Error(fmt.Sprintf("Error: %v", err))
					return nil
				}
				if machine.Status.String() != "Idle" {
					logs.Error(fmt.Sprintf("Machine is not idle, status: %v", machine.Status.String()))
					return nil
				}

				hash, err := distriWrapper.RemoveMachine(*hwInfo)
				if err != nil {
					logs.Error(fmt.Sprintf("Error block : %v, msg : %v\n", hash, err))
					return nil
				}
				return nil
			},
		},
	},
}
