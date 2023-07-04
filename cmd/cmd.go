package cmd

import (
	"kvm_backup/router"

	"github.com/spf13/cobra"
)

func RunInfo() {
	var rootCmd = &cobra.Command{
		Use:  "main",
		Long: "kvm备份管理",
	}
	rootCmd.AddCommand(actionInfo())
	rootCmd.Execute()
}

/*
运行web服务
方式1：go run main.go web
方式2：go build;./kvm_backup web
*/
func actionInfo() *cobra.Command {
	var a = &cobra.Command{}
	a.Use = "web" //不能有空格
	a.Long = "[启动web服务]"
	a.Execute()

	a.Run = func(cmd *cobra.Command, args []string) {
		router.RouterInfo()
		router.Run()
	}
	return a
}
