package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "lsync",
	Short: "sync log files to multi-cloud storage (like S3) and logs services (like CloudWatch)",
	Long: `lsync is a command line utility which integrates with AWS, GCP and Azure
for syncing and setting up logs and dashboards`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
