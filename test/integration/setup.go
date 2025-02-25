package integration

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

// SetupTestInfrastructure provisions the NATS server using OpenTofu
func SetupTestInfrastructure() (string, func(), error) {
	// Get the absolute path to the infrastructure directory
	infraDir, err := filepath.Abs("./test/infrastructure")
	if err != nil {
		return "", nil, fmt.Errorf("failed to get infrastructure directory path: %w", err)
	}

	// Initialize OpenTofu
	cmd := exec.Command("tofu", "init")
	cmd.Dir = infraDir
	if err := cmd.Run(); err != nil {
		return "", nil, fmt.Errorf("failed to initialize OpenTofu: %w", err)
	}

	// Apply the configuration
	cmd = exec.Command("tofu", "apply", "-auto-approve")
	cmd.Dir = infraDir
	var outBuf bytes.Buffer
	cmd.Stdout = &outBuf
	if err := cmd.Run(); err != nil {
		return "", nil, fmt.Errorf("failed to apply OpenTofu configuration: %w", err)
	}

	// Extract the NATS URI from the output
	cmd = exec.Command("tofu", "output", "nats_uri")
	cmd.Dir = infraDir
	var natsURIBuf bytes.Buffer
	cmd.Stdout = &natsURIBuf
	if err := cmd.Run(); err != nil {
		// Try to clean up before returning error
		cleanupCmd := exec.Command("tofu", "destroy", "-auto-approve")
		cleanupCmd.Dir = infraDir
		cleanupCmd.Run()
		return "", nil, fmt.Errorf("failed to get NATS URI: %w", err)
	}

	natsURI := string(bytes.TrimSpace(natsURIBuf.Bytes()))
	
	// Wait for NATS to be ready
	time.Sleep(2 * time.Second)

	// Return cleanup function
	cleanup := func() {
		cmd := exec.Command("tofu", "destroy", "-auto-approve")
		cmd.Dir = infraDir
		cmd.Run()
	}

	return natsURI, cleanup, nil
}
