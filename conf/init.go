package conf

import (
	"fmt"
	"giligili/cache"
	"giligili/model"
	"giligili/tasks"
	"giligili/util"
	"github.com/joho/godotenv"
	"os"
)

// Init 初始化配置项
func Init() {
	// 从本地读取环境变量
	godotenv.Load()

	// 设置日志级别
	util.BuildLogger(os.Getenv("LOG_LEVEL"))

	// 读取翻译文件
	if err := LoadLocales("conf/locales/zh-cn.yaml"); err != nil {
		util.Log().Panic("翻译文件加载失败", err)
	}

	// 连接数据库
	model.Database(os.Getenv("MYSQL_DSN"))
	cache.Redis()

	// 启动定时任务
	tasks.CronJob()

	// 创建资源路径
	isExist := func(path string) bool {
		_, err := os.Stat(path)
		if err != nil {
			if os.IsExist(err) {
				return true
			}
			if os.IsNotExist(err) {
				return false
			}
			fmt.Println(err)
			return false
		}
		return true
	}
	if !isExist("./resource") {
		err := os.MkdirAll("./resource", os.ModePerm)
		if err != nil {
			util.Log().Panic("创建路径 ./resource 失败", err)
		}
		util.Log().Info("路径 ./resource 不存在，已创建")
	}
	if !isExist("./resource/cover") {
		err := os.MkdirAll("./resource/cover", os.ModePerm)
		if err != nil {
			util.Log().Panic("创建路径 ./resource/cover 失败", err)
		}
		util.Log().Info("路径 ./resource/cover 不存在，已创建")
	}
	if !isExist("./resource/video") {
		err := os.MkdirAll("./resource/video", os.ModePerm)
		if err != nil {
			util.Log().Panic("创建路径 ./resource/video 失败", err)
		}
		util.Log().Info("路径 ./resource/video 不存在，已创建")
	}
	if !isExist("./resource/avatar") {
		err := os.MkdirAll("./resource/avatar", os.ModePerm)
		if err != nil {
			util.Log().Panic("创建路径 ./resource/avatar 失败", err)
		}
		util.Log().Info("路径 ./resource/avatar 不存在，已创建")
	}
}
