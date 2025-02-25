package integration

import (
	"os"
	"testing"

	"github.com/spf13/viper"
	"your-module-path/app/stream"
)

func TestMain(m *testing.M) {
	// Setup test infrastructure
	natsURI, cleanup, err := SetupTestInfrastructure()
	if err != nil {
		panic("Failed to setup test infrastructure: " + err.Error())
	}
	
	// Ensure cleanup runs after tests
	defer cleanup()
	
	// Run tests
	exitCode := m.Run()
	
	// Exit with the same code
	os.Exit(exitCode)
}

func TestNatsStreamEngine(t *testing.T) {
	// Setup viper with test configuration
	viper.SetConfigType("toml")
	viper.Set("stream.engine", "nats")
	viper.Set("stream.connection_string", "nats://localhost:4222")
	
	// Create stream engine
	streamEngine, err := stream.NewFromViperSettings()
	if err != nil {
		t.Fatalf("Failed to create stream engine: %v", err)
	}
	
	// Test publishing a message
	err = streamEngine.Pub("test.subject", []byte("test message"))
	if err != nil {
		t.Fatalf("Failed to publish message: %v", err)
	}
	
	// Add more assertions as needed
}
