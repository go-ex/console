package console

import (
	"errors"
	"fmt"
	"github.com/go-ex/console/contract"
	"os"
	"os/exec"
	"path/filepath"
	"plugin"
	"strings"
)

func New() contract.Console {
	return &Console{}
}

type Console struct {
}

func (c *Console) HasCommand() bool {
	argsLen := len(os.Args)

	if argsLen < 2 {
		return false
	} else {
		cmdName := os.Args[1] + ".so"
		for _, path := range GetPluginPath() {
			if hasPlugin(path + cmdName) {
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

	for _, path := range GetPluginPath() {
		pluginPath := path + cmdName
		if hasPlugin(pluginPath) {
			plug, err := plugin.Open(pluginPath)
			if err != nil {
				return err
			}

			inf, err := plug.Lookup("Command")
			if err != nil {
				return err
			}

			cmd := inf.(contract.Command)
			configure := cmd.Configure()
			input := contract.Input{
				Has:      map[string]bool{},
				Argument: map[string]string{},
				Option:   map[string]string{},
			}

			err = input.Parsed(configure.Input)
			if err != nil {
				return err
			}

			if input.GetHas("-h") {
				return c.RunCommandHelp(name, configure)
			}
			// 如果有守护进程方式启动参数，拦截，并且转换后台启动
			if input.GetHas("-d") {
				for key, str := range os.Args {
					if str == "-d" {
						os.Args[key] = "-d=true"
						break
					}
				}
				command := exec.Command(os.Args[0], os.Args[1:]...)
				out, err := os.OpenFile("/dev/null", os.O_RDWR, 0)
				if err == nil {
					command.Stdout = out
				}
				return command.Start()
			} else if input.GetOption("d") == "true" {
				// 命令转换为后台的传入
				input.Has["-d"] = true
			}

			cmd.Execute(input)

			return nil
		}
	}

	return errors.New("无法加载插件")
}

func GetWd() string {
	dir, _ := os.Getwd()

	return dir
}

func GetRoot() string {
	paths, fileName := filepath.Split(os.Args[0])
	ext := filepath.Ext(fileName)
	abs := strings.TrimSuffix(fileName, ext)

	return paths + abs
}

// 命令插件可能在的目录，当前目录、执行目录
func GetPluginPath() []string {
	return []string{GetWd() + "/script/plugin/", GetRoot() + "/"}
}

func hasPlugin(file string) bool {
	_, err := os.Stat(file)
	if err != nil {
		// no such file or dir
		return false
	}
	return true
}
