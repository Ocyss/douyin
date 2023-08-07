package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/Ocyss/douyin/internal/bootstrap"
	"github.com/Ocyss/douyin/internal/conf"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "查看配置",
	Run: func(cmd *cobra.Command, args []string) {
		bootstrap.InitConf()
		s, err := json.MarshalIndent(conf.Conf, "", "  ")
		if err != nil {
			log.Fatal("未知错误: ", err)
		}
		fmt.Println(string(s))
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
