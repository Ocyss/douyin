package cmd

import (
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/Ocyss/douyin/cmd/flags"
	"github.com/Ocyss/douyin/internal/bootstrap"
	"github.com/Ocyss/douyin/utils"
	"github.com/sirupsen/logrus"
)

var (
	pid     = -1
	pidFile string
)

// initServer 初始化服务
func initServer() {
	// 配置日志格式
	formatter := logrus.TextFormatter{
		ForceColors:               true,
		EnvironmentOverrideColors: true,
		TimestampFormat:           "2006-01-02 15:04:05",
		FullTimestamp:             true,
		DisableQuote:              true,
	}
	logrus.SetFormatter(&formatter)
	// 服务初始化
	bootstrap.InitConf()
	bootstrap.InitLog()
	bootstrap.InitDb()
	bootstrap.InitRdb()
	rand.Seed(time.Now().Unix())
}

// initDaemon 守护进程初始化
func initDaemon() {
	pidFile = filepath.Join(flags.DataDir, "pid")
	if utils.Exists(pidFile) {
		bytes, err := os.ReadFile(pidFile)
		if err != nil {
			logrus.Fatal("无法读取pid文件，", err)
		}
		id, err := strconv.Atoi(string(bytes))
		if err != nil {
			logrus.Fatal("无法转换pid，", err)
		}
		pid = id
	}
}
