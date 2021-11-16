package cmd

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use:   "",
	Short: "",
	Long:  "",
}

var ConfigPath string

func Execute() error {
	return RootCmd.Execute()
}

func init() {
	serverCmd.Flags().StringVarP(&ConfigPath, "config", "c", "config/", "指定要使用的配置文件路径")
}
