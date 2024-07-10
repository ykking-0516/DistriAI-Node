package docker

import (
	"DistriAI-Node/config"
	docker_utils "DistriAI-Node/docker/utils"
	"DistriAI-Node/pattern"
	"DistriAI-Node/utils"
	logs "DistriAI-Node/utils/log_utils"
	"bufio"
	"context"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)

// Check if a Docker image exists locally
func ImageExistOrPull(imageName string) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return err
	}
	cli.NegotiateAPIVersion(ctx)

	isCreated, _ := docker_utils.ImageExist(ctx, cli, imageName)
	if !isCreated {
		// If image does not exist locally, pull it from Docker registry
		if err := docker_utils.PullImage(imageName); err != nil {
			return err
		}
	}
	return nil
}

// Run a Docker container and retrieves a score from its logs
func RunScoreContainer(isGPU bool) (float64, error) {
	oldScore := 0.0

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return oldScore, err
	}
	cli.NegotiateAPIVersion(ctx)

	hostConfig := &container.HostConfig{
		AutoRemove: true,
	}
	if isGPU {
		hostConfig.Runtime = "nvidia"
		hostConfig.Resources = container.Resources{
			DeviceRequests: []container.DeviceRequest{
				{
					Count:        -1,
					Capabilities: [][]string{{"gpu"}},
				},
			},
		}
	}

	containerID, err := docker_utils.RunContainer(ctx, cli, pattern.SCORE_CONTAINER,
		&container.Config{
			Image: pattern.SCORE_NAME,
		},
		hostConfig)
	if err != nil {
		return oldScore, err
	}

	reader, err := cli.ContainerLogs(ctx, containerID, types.ContainerLogsOptions{
		ShowStdout: true,
		ShowStderr: true,
		Follow:     true,
		Timestamps: true})
	if err != nil {
		return oldScore, err
	}
	defer reader.Close()

	scanner1 := bufio.NewScanner(reader)
	for scanner1.Scan() {
		out := scanner1.Text()
		index := strings.Index(out, "Score:")
		if index > 0 {
			scoreStr := strings.TrimSpace(out[index+len("Score:"):])
			newScore, err := strconv.ParseFloat(scoreStr, 64)
			oldScore = (oldScore + newScore) / 2
			if err != nil {
				return oldScore, err
			}
		}
	}
	return oldScore, nil
}

func RunWorkspaceContainer(isGPU bool, mlToken string) (string, error) {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return "", err
	}
	cli.NegotiateAPIVersion(ctx)

	containerName := pattern.ML_WORKSPACE_CONTAINER
	containerConfig := &container.Config{
		Image: pattern.ML_WORKSPACE_NAME,
		Env: []string{
			fmt.Sprintf("AUTHENTICATE_VIA_JUPYTER=%s", mlToken),
		},
		Tty: true,
	}

	portBind := nat.PortMap{
		nat.Port("8080/tcp"): []nat.PortBinding{
			{
				HostIP:   "0.0.0.0",
				HostPort: config.GlobalConfig.Console.WorkPort,
			},
		}}

	if config.GlobalConfig.Console.ExpandPort1 != "" {
		hostPort1 := config.GlobalConfig.Console.ExpandPort1
		port1 := fmt.Sprintf("%s/tcp", hostPort1)
		logs.Normal(fmt.Sprintf("port1: %s", port1))
		portBind[nat.Port(port1)] = []nat.PortBinding{
			{
				HostIP:   "0.0.0.0",
				HostPort: hostPort1,
			}}
	}
	if config.GlobalConfig.Console.ExpandPort2 != "" {
		hostPort2 := config.GlobalConfig.Console.ExpandPort2
		port2 := fmt.Sprintf("%s/tcp", hostPort2)
		logs.Normal(fmt.Sprintf("port2: %s", port2))
		portBind[nat.Port(port2)] = []nat.PortBinding{
			{
				HostIP:   "0.0.0.0",
				HostPort: hostPort2,
			}}
	}
	if config.GlobalConfig.Console.ExpandPort3 != "" {
		hostPort3 := config.GlobalConfig.Console.ExpandPort3
		port3 := fmt.Sprintf("%s/tcp", hostPort3)
		logs.Normal(fmt.Sprintf("port3: %s", port3))
		portBind[nat.Port(port3)] = []nat.PortBinding{
			{
				HostIP:   "0.0.0.0",
				HostPort: hostPort3,
			}}
	}
	hostConfig := &container.HostConfig{
		PortBindings: portBind,
		Binds: []string{
			fmt.Sprintf("%s/ml-workspace:/workspace", config.GlobalConfig.Console.WorkDirectory),
			"myvolume:/data",
		},
		RestartPolicy: container.RestartPolicy{
			Name: "always",
		},
		ShmSize: 512 * 1024 * 1024, // 512MB
	}
	if isGPU {
		containerName = pattern.ML_WORKSPACE_GPU_CONTAINER
		containerConfig.Image = pattern.ML_WORKSPACE_GPU_NAME
		hostConfig.Runtime = "nvidia"
		hostConfig.Resources = container.Resources{
			DeviceRequests: []container.DeviceRequest{
				{
					Count:        -1,
					Capabilities: [][]string{{"gpu"}},
				},
			},
		}
	}

	isExists, containerID := docker_utils.ContainerExists(ctx, cli, containerName)
	if isExists {
		if err := StopWorkspaceContainer(containerID); err != nil {
			return containerID, fmt.Errorf("> StopWorkspaceContainer: %v", err)
		}
	}

	containerID, err = docker_utils.RunContainer(ctx, cli, containerName,
		containerConfig,
		hostConfig)
	if err != nil {
		return "", err
	}
	return containerID, nil
}

func TestRunWorkspaceContainer(isGPU bool, mlToken string) (string, error) {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return "", err
	}
	cli.NegotiateAPIVersion(ctx)

	containerName := pattern.ML_WORKSPACE_CONTAINER
	if isGPU {
		containerName = pattern.ML_WORKSPACE_GPU_CONTAINER
	}

	isExists, containerID := docker_utils.ContainerExists(ctx, cli, containerName)
	if isExists {
		if err := StopWorkspaceContainer(containerID); err != nil {
			return containerID, fmt.Errorf("> StopWorkspaceContainer: %v", err)
		}
	}

	cmd := exec.Command("sudo", "docker", "run", "-d")
	cmd.Args = append(cmd.Args, "-p", fmt.Sprintf("%s:8080", config.GlobalConfig.Console.WorkPort))

	if config.GlobalConfig.Console.ExpandPort1 != "" {
		cmd.Args = append(cmd.Args, "-p", fmt.Sprintf("%s:%s", config.GlobalConfig.Console.ExpandPort1, config.GlobalConfig.Console.ExpandPort1))
	}
	if config.GlobalConfig.Console.ExpandPort2 != "" {
		cmd.Args = append(cmd.Args, "-p", fmt.Sprintf("%s:%s", config.GlobalConfig.Console.ExpandPort2, config.GlobalConfig.Console.ExpandPort2))
	}
	if config.GlobalConfig.Console.ExpandPort3 != "" {
		cmd.Args = append(cmd.Args, "-p", fmt.Sprintf("%s:%s", config.GlobalConfig.Console.ExpandPort3, config.GlobalConfig.Console.ExpandPort3))
	}

	if isGPU {
		cmd.Args = append(cmd.Args, "--runtime=nvidia")
		cmd.Args = append(cmd.Args, "--gpus", "all")
	}

	cmd.Args = append(cmd.Args, "--env", fmt.Sprintf("AUTHENTICATE_VIA_JUPYTER=%s", mlToken))

	cmd.Args = append(cmd.Args, "--name", pattern.ML_WORKSPACE_CONTAINER)
	cmd.Args = append(cmd.Args, "-v", fmt.Sprintf("%s/ml-workspace:/workspace", config.GlobalConfig.Console.WorkDirectory))
	cmd.Args = append(cmd.Args, "--shm-size", "512m")
	cmd.Args = append(cmd.Args, "--restart", "always")

	cmd.Args = append(cmd.Args, pattern.ML_WORKSPACE_GPU_NAME)

	logs.Normal(fmt.Sprintf("Command: %v", strings.Join(cmd.Args, " ")))

	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("> exec.Command: %s, output: %s", err.Error(), string(output))
	}
	outputStr := string(output)
	if len(outputStr) > 64 {
		outputStr = outputStr[:64]
	}
	return outputStr, nil
}

func RunDeployContainer(isGPU bool, downloadURL []string) (string, error) {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return "", err
	}
	cli.NegotiateAPIVersion(ctx)
	isExists, containerID := docker_utils.ContainerExists(ctx, cli, pattern.MODELS_DEPLOY_CONTAINER)
	if isExists {
		if err := StopWorkspaceContainer(containerID); err != nil {
			return containerID, fmt.Errorf("> StopWorkspaceContainer: %v", err)
		}
	}

	if len(downloadURL) == 0 {
		return "", fmt.Errorf("> downloadURL is empty")
	}

	cmd := exec.Command("sudo", "docker", "run", "-d")

	cmd.Args = append(cmd.Args, "-p", fmt.Sprintf("%s:7860", config.GlobalConfig.Console.WorkPort))

	if isGPU {
		cmd.Args = append(cmd.Args, "--runtime=nvidia")
		cmd.Args = append(cmd.Args, "--gpus", "all")
	}

	cmd.Args = append(cmd.Args, "--name", pattern.MODELS_DEPLOY_CONTAINER)

	host, path, err := utils.SplitURL(downloadURL[0])
	if err != nil {
		return "", fmt.Errorf("> SplitURL downloadURL[0]: %v", err)
	}

	cmd.Args = append(cmd.Args, "--env", fmt.Sprintf("DOWNLOAD_URL=%s", host))
	cmd.Args = append(cmd.Args, "--env", fmt.Sprintf("DEPLOY_FILE=%s", strings.TrimPrefix(path, "/")))
	if len(downloadURL) == 2 {
		_, path, err = utils.SplitURL(downloadURL[1])
		if err != nil {
			return "", fmt.Errorf("> SplitURL downloadURL[1]: %v", err)
		}
		cmd.Args = append(cmd.Args, "--env", fmt.Sprintf("REQUIREMENTS=%s", strings.TrimPrefix(path, "/")))
	}

	cmd.Args = append(cmd.Args, "--restart", "always")

	cmd.Args = append(cmd.Args, pattern.MODELS_DEPLOY_NAME)

	logs.Normal(fmt.Sprintf("Command: %v", strings.Join(cmd.Args, " ")))

	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("> exec.Command: %s, output: %s", err.Error(), string(output))
	}
	outputStr := string(output)
	if len(outputStr) > 64 {
		outputStr = outputStr[:64]
	}
	return outputStr, nil
}

func StopWorkspaceContainer(containerID string) error {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return err
	}
	cli.NegotiateAPIVersion(ctx)

	if err := docker_utils.StopAndRemoveContainer(ctx, cli, containerID); err != nil {
		return err
	}

	dir := config.GlobalConfig.Console.WorkDirectory + "/ml-workspace"
	err = os.RemoveAll(dir)
	if err != nil {
		return err
	}

	return nil
}
