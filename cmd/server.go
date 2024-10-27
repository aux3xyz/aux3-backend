package cmd

import (
    "fmt"

    "github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
    Use: "server",
    Short: "start server to listen",
    Long: `Hit Ctrl+D to exit, will run until terminated`,
    Run: func(cmd *cobra.Command, args []string) {
        // TODO start threads and other things
        // TODO listen for SIGTERM/SIGINT
        fmt.Println("Hello, world!")
    },
}
