package cmd

import (
    "fmt"

    "github.com/spf13/cobra"

    "aux3.xyz/common"
)

var toolscmd = &cobra.Command{
    Use: "tools",
    Short: "tools to run locally for interacting with AUX3",
    Long: `includes things like generating public/private keys and more`,
    Run: func(cmd *cobra.Command, args []string) {
        priv, pub := common.GenerateKey()
        fmt.Printf("private: %s\n\npublic: %s\n", priv, pub)
    },
}
