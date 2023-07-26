package cmd

import (
	"fmt"
	"github.com/Ocyss/douyin/cmd/flags"
	"github.com/Ocyss/douyin/internal/bootstrap"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

var rootCmd = &cobra.Command{
	Use:   "douyin",
	Short: "抖音让每一个人看见并连接更大的世界，鼓励表达、沟通和记录，激发创造，丰富人们的精神世界，让现实生活更美好，你所热爱的，就是你的生活。",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(`douyin v0.0.1
- start 	启动守护进程
- stop 		停止守护进程
- server  	前台启动服务
- update 	进行更新
- version 	显示版本
- help 		显示帮助
- config 	显示配置信息
- init 		初始化配置信息
- system 	系统服务注册
- migrate 	迁移数据库`)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
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
