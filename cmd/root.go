package cmd

import (
    "github.com/spf13/cobra"
)

var (
    rootCmd = &cobra.Command{
        Use: "aux3-server",
        Short: "Listen for EVM events and perform actions",
        Long: `Subscribe to events by publishing to a smart contract and funding.
                When an event occurs an action will be triggered according to the
                configuration. Notifications are the main target currently.`,
    }
)

func Execute() error {
    return rootCmd.Execute()
}

func init() {
    rootCmd.AddCommand(serverCmd)
}
