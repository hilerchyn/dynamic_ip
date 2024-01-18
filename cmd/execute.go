package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string

	rootCmd = &cobra.Command{
		Use:   "dynamic-ip-cli",
		Short: "Dynamic IP binding",
		Long:  "This is a CLI tool for GTEX that binds dynamic IP to specific domain.",
	}
)

func initConfig() {
	if cfgFile != "" {
		// 加载配置文件
		viper.SetConfigFile(cfgFile)
	} else {
		// 查询home目录
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".gtex_dynamic_ip")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gtex_dynamic_ip.yaml)")
	rootCmd.PersistentFlags().StringP("author", "a", "chen.tao", "author name for copyright attribution")
	rootCmd.PersistentFlags().Bool("viper", true, "use viper for configuration")

	viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
	viper.SetDefault("author", "chen.tao <hilerchyn@gmail.com>")

}
