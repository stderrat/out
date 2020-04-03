package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	installconfig "github.com/stderrat/out/pkg/installconfig"
)

func newTestCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "test",
		Short: "Test OpenShift requirements before installation",
		Run:   runTestCommand,
	}

	return cmd
}

func runTestCommand(cmd *cobra.Command, args []string) {
	logrus.Info("Running test command")

	err := installconfig.Load("../../test/testdata")
	if err != nil {
		logrus.Info("error")
	}
}
