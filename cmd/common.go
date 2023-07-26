package cmd

import (
	"github.com/Ocyss/douyin/cmd/flags"
	"github.com/Ocyss/douyin/internal/bootstrap"
	"github.com/Ocyss/douyin/utils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"strconv"
)

var pid = -1
var pidFile string

// initServer 初始化服务
func initServer() {
	var baseDir, dataDir string
	var err error
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
	cobra.OnInitialize(bootstrap.InitConf)
	cobra.OnInitialize(bootstrap.InitLog)
	cobra.OnInitialize(bootstrap.InitDb)

	rootCmd.PersistentFlags().StringVar(&dataDir, "data", "data", "修改配置文件路径")
	rootCmd.PersistentFlags().BoolVar(&flags.Debug, "debug", false, "Debug 模式（更多的日志输出）")
	rootCmd.PersistentFlags().BoolVar(&flags.Dev, "dev", false, "开发环境")
	rootCmd.PersistentFlags().BoolVar(&flags.LogStd, "log-std", false, "日志强制打印到控制台")
	rootCmd.PersistentFlags().BoolVar(&flags.Memory, "memory", false, "使用内存数据库")
	flags.Pro = !flags.Dev
	// 获取可执行文件路径
	if baseDir, err = os.Executable(); err != nil {
		logrus.Fatal(err)
	}
	flags.ExPath = filepath.Dir(baseDir)
	flags.DataDir = filepath.Join(flags.ExPath, dataDir)
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
