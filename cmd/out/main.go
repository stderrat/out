package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	// "github.com/sirupsen/logrus"
)

func main() {
	run()
}

func run() {
	rootCmd := &cobra.Command{
		Use:   filepath.Base(os.Args[0]),
		Short: "Tests prerequisites for installing OpenShift with the UPI method",
		Long:  "",
		Run: runRootCmd,
	}

	rootCmd.Execute()
}

func runRootCmd(cmd *cobra.Command, args []string) {
	fmt.Println("hello world!")
}
