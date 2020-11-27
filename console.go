package console

import (
	"errors"
	"fmt"
	"github.com/go-ex/console/contract"
	"gopkg.in/ini.v1"
	"os"
	"plugin"
)

func New() contract.Console {
	return &Console{}
}

type Console struct {
	programName string // 编译后的程序名称,命令
	programPath string // 编译后的程序绝对路径

	config *ini.File
}

func (c *Console) HasCommand() bool {
	argsLen := len(os.Args)

	if argsLen < 2 {
		return false
	} else {
		cmdName := os.Args[1] + ".so"
		for _, path := range c.GetPluginPath() {
			if hasFile(path + cmdName) {
				return true
			}
		}

		fmt.Println("不存在的命令:" + os.Args[1])
	}
	return false
}

func (c *Console) RunCommand() error {
	name := os.Args[1]
	cmdName := name + ".so"

	for _, path := range c.GetPluginPath() {
		pluginPath := path + cmdName
		if hasFile(pluginPath) {
			plug, err := plugin.Open(pluginPath)
			if err != nil {
				return err
			}

			inf, err := plug.Lookup("Command")
			if err != nil {
				return err
			}

			cmd, ok := inf.(contract.Command)
			if ok {
				return c.RunPlugin(path, name, cmd)
			}
		}
	}

	return errors.New("无法加载插件")
}

func GetWd() string {
	dir, _ := os.Getwd()

	return dir
}
