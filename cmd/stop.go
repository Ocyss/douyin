package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
)

var stopCmd = &cobra.Command{
	Use: "stop",
	Run: func(cmd *cobra.Command, args []string) {
		stop()
	},
}

func stop() {
	initDaemon()
	if pid == -1 {
		log.Info("似乎还没有启动。尝试使用 `./douyin start` 启动服务器.")
		return
	}
	process, err := os.FindProcess(pid)
	if err != nil {
		log.Errorf("无法按pid找到进程：%d，原因: %v", pid, process)
		return
	}
	err = process.Kill()
	if err != nil {
		log.Errorf("无法终止进程 %d: %v", pid, err)
	} else {
		log.Info("杀死进程: ", pid)
	}
	err = os.Remove(pidFile)
	if err != nil {
		log.Errorf("pid 文件未能正常删除")
	}
	pid = -1
}

func init() {
	rootCmd.AddCommand(stopCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// stopCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// stopCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
