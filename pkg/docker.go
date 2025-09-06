package docker

import (
	"bytes"
	"errors"
	"os/exec"
	"strings"
)

// Helper provides methods to interact with Docker.
type Helper struct{}

// IsRunning checks if Docker daemon is running.
func (h *Helper) IsRunning() bool {
	cmd := exec.Command("docker", "info")
	err := cmd.Run()
	return err == nil
}

// Version returns the Docker version.
func (h *Helper) Version() (string, error) {
	out, err := exec.Command("docker", "version", "--format", "{{.Server.Version}}").Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}

// ListImages returns a list of Docker images.
func (h *Helper) ListImages() ([]string, error) {
	out, err := exec.Command("docker", "images", "--format", "{{.Repository}}:{{.Tag}}").Output()
	if err != nil {
		return nil, err
	}
	lines := strings.Split(strings.TrimSpace(string(out)), "\n")
	if len(lines) == 1 && lines[0] == "" {
		return []string{}, nil
	}
	return lines, nil
}

// ListContainers returns a list of running Docker containers.
func (h *Helper) ListContainers() ([]string, error) {
	out, err := exec.Command("docker", "ps", "--format", "{{.Names}}").Output()
	if err != nil {
		return nil, err
	}
	lines := strings.Split(strings.TrimSpace(string(out)), "\n")
	if len(lines) == 1 && lines[0] == "" {
		return []string{}, nil
	}
	return lines, nil
}

// RunCommand runs a custom docker command and returns its output.
func (h *Helper) RunCommand(args ...string) (string, error) {
	if len(args) == 0 {
		return "", errors.New("no command specified")
	}
	cmd := exec.Command("docker", args...)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out
	err := cmd.Run()
	return out.String(), err
}