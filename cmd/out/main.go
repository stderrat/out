package main

import (
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var (
	rootOpts struct {
		dir      string
		logLevel string
	}
)

func main() {
	rootCmd := newRootCmd()
	rootCmd.Execute()
}

func newRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   filepath.Base(os.Args[0]),
		Short: "Tests prerequisites for installing OpenShift with the UPI method",
		Long:  "",
		SilenceUsage:     true,
		// Run: runRootCmd,
	}

	rootCmd.AddCommand(newTestCommand())
	rootCmd.PersistentFlags().StringVar(&rootOpts.dir, "dir", ".", "directory containing install-config.yaml")

	return rootCmd
}

// func runRootCmd(cmd *cobra.Command, args []string) {
// 	fmt.Println("hello world!")
// }
