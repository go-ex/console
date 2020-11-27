package console

import (
	"gopkg.in/ini.v1"
	"os"
	"path/filepath"
)

func (c *Console) Init() {
	c.programPath, c.programName = filepath.Split(os.Args[0])

	c.config, _ = ini.Load(defaultConfig)

	err := c.config.Append(c.programPath + c.programName + ".ini")
	if err != nil {
		_ = c.config.SaveTo(c.programPath + c.programName + ".ini")
	}
}

var defaultConfig = []byte(`
[plugin]
# 启动程序存储目录下的 root + root_path
root_path = /plugin
# 执行命令当前目录下的  worker + worker_path
worker_path = /script/plugin
`)
