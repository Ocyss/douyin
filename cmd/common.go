package cmd

import (
	"github.com/Ocyss/douyin/cmd/flags"
	"github.com/Ocyss/douyin/utils"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"strconv"
)

var pid = -1
var pidFile string

// initDaemon 守护进程初始化
func initDaemon() {
	pidFile = filepath.Join(flags.DataDir, "pid")
	if utils.Exists(pidFile) {
		bytes, err := os.ReadFile(pidFile)
		if err != nil {
			log.Fatal("无法读取pid文件，", err)
		}
		id, err := strconv.Atoi(string(bytes))
		if err != nil {
			log.Fatal("无法转换pid，", err)
		}
		pid = id
	}
}
