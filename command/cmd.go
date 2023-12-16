package command

import (
	"fmt"
	"fyoukuApi/command/mq"
	"github.com/spf13/cobra"
	"os"
)

func Execute() {
	var rootCmd = &cobra.Command{
		Use:   "fyouku",
		Short: "fyouku is 是启动项目任务的命令",
	}

	rootCmd.AddCommand(mq.Top)
	rootCmd.AddCommand(mq.Send)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println("任务启动失败: ", err)
		os.Exit(1)
	}
}
