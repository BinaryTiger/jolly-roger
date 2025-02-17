package cmd

import (
	"jolly_roger/test"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

// FIXME
func TestPrintConfig(t *testing.T) {
	t.Skip("not implemented correctly, skipping")

	viper.AddConfigPath("..")
	viper.SetConfigType("toml")
	viper.SetConfigName("jolly_roger.toml")

	if err := viper.ReadInConfig(); err != nil {
		t.Errorf("couldn't read config file: %v", err)
	} else {
		expected := `
config file: jolly_roger.toml
=======================
storage:
  engine: sqlite3
  connection_string: file:local.db
unested_config: tycho
		`

		actual := test.CaptureOutput(func() {
			printConfig()
		})

		assert.Equal(t, expected, actual, "what does this print")
	}
}
