package bootstrap

import (
	"encoding/json"
	"github.com/Ocyss/douyin/cmd/flags"
	"github.com/Ocyss/douyin/internal/conf"
	"github.com/Ocyss/douyin/utils"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

var configPath string

func InitConf() {
	configPath = filepath.Join(flags.DataDir, "config.json")
	if !utils.Exists(configPath) {
		// 配置文件不存在，创建默认配置
		log.Info("没检测到配置文件，将进行初始化 config.json.")
		basePath := filepath.Dir(configPath)
		err := os.MkdirAll(basePath, 0766)
		if err != nil {
			log.Fatalf("无法创建文件夹, %s", err)
		}
		conf.Conf = conf.DefaultConfig()
		defaultData, _ := json.MarshalIndent(conf.Conf, "", "  ")
		err = os.WriteFile(configPath, defaultData, 0666)
		if err != nil {
			log.Fatalf("配置文件写入错误，请检查,{%s}", err)
		}
		return
	}
	file, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatalf("配置读取错误，请检查,{%s}", err)
	}
	err = json.Unmarshal(file, &conf.Conf)
	if err != nil {
		log.Fatalf("配置文件解析错误，请检查,{%s}", err)
	}
}
