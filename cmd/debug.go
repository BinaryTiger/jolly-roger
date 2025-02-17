package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// DebugCmd represents the debug command
var DebugCmd = &cobra.Command{
	Use:   "debug",
	Short: "print useful runtime information",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		printConfig()
	},
}

func init() {
	rootCmd.AddCommand(DebugCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// debugCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// debugCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func printConfig() {
	fmt.Println("config file:", viper.ConfigFileUsed())
	fmt.Println("=======================")
	unfoldSettings(viper.AllSettings(), 0)
}

func unfoldSettings(m map[string]interface{}, indent int) {
	indentStr := strings.Repeat("  ", indent)

	for k, v := range m {
		switch val := v.(type) {
		case map[string]interface{}:
			fmt.Printf("%s%s:\n", indentStr, k)
			unfoldSettings(val, indent+1)
		case []interface{}:
			fmt.Printf("%s%s:\n", indentStr, k)
			for _, item := range val {
				switch i := item.(type) {
				case map[string]interface{}:
					unfoldSettings(i, indent+1)
				default:
					fmt.Printf("%s  - %v\n", indentStr, i)
				}
			}
		default:
			fmt.Printf("%s%s: %v\n", indentStr, k, v)
		}
	}
}
