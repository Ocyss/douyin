package cmd

import (
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"

	"github.com/Ocyss/douyin/cmd/flags"
	"github.com/Ocyss/douyin/internal/bootstrap"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "初始化配置",
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("开始初始化配置文件")
		if bootstrap.InitConf() == 0 {
			log.Info("已有配置文件,开始进行备份")
			err := os.Rename(filepath.Join(flags.DataDir, "config.json"), filepath.Join(flags.DataDir, "config.old.json"))
			if err != nil {
				log.Fatal("改名失败...")
			}
			bootstrap.InitConf()
		}
		log.Info("Ok!")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
