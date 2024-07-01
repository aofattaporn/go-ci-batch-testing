package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	// appStart   time.Time
	configFile string
	rootCmd    = &cobra.Command{
		Use:   "dwpr-ci-batch",
		Short: "dwpr-ci-batch",
	}
)

func init() {
	rootCmd.PersistentFlags().StringVar(&configFile, "config", "env/config.yaml", "config file (default is env/config.yaml)")
}

func Excecute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
