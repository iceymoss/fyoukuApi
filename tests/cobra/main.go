package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func main() {
	var rootCmd = &cobra.Command{Use: "fyouku"}

	var consumeCmd = &cobra.Command{
		Use:   "consume-top",
		Short: "Starts the MQ consumer service",
		Run: func(cmd *cobra.Command, args []string) {
			// 在这里编写启动 MQ 消费端服务的代码
			fmt.Println("Starting MQ consumer service...")
		},
	}

	rootCmd.AddCommand(consumeCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
