package main

import (
	"fmt"
	"github.com/fsouza/go-dockerclient"
)

var endpoint = "unix:///var/run/docker.sock"
var docker_client, _ = docker.NewClient(endpoint)

func test() {

	container, err := docker_client.CreateContainer(*dockeroptions())

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	dockerstart(container.ID, container.HostConfig)

}

func dockeroptions() *docker.CreateContainerOptions {
	var binds []string
	volumes := map[string]struct{}{}
	volumeDefinitions := make(map[string]string)
	volumeDefinitions["/home/toon/Documents/poc-provisioning-queue/tmp"] = "/tmp"

	for hostPath, containerPath := range volumeDefinitions {
		binds = append(binds, hostPath+":"+containerPath)
		volumes[containerPath] = struct{}{}
	}
	hostConfig := &docker.HostConfig{
		Binds: binds,
	}

	conf := &docker.Config{
		Tty:     true,
		Image:   "alpine",
		Cmd:     []string{"sh"},
		Volumes: volumes,
	}
	opts := &docker.CreateContainerOptions{
		Name:       "test",
		Config:     conf,
		HostConfig: hostConfig,
	}
	return opts

}

func dockerstart(id string, hostConfig *docker.HostConfig) {
	docker_client.StartContainer(id, hostConfig)

}

func dockerbuild() {

}
