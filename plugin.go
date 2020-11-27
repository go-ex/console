package console

import (
	"github.com/go-ex/console/contract"
	"gopkg.in/ini.v1"
	"os"
	"os/exec"
)

// 命令插件可能在的目录，当前目录、执行目录
func (c *Console) GetPluginPath() []string {
	return []string{
		c.programPath + c.config.Section("plugin").Key("root_path").String() + "/",
		GetWd() + c.config.Section("plugin").Key("worker_path").String() + "/",
	}
}

// 执行这个插件
func (c *Console) RunPlugin(path string, name string, cmd contract.Command) error {
	configure := cmd.Configure()
	input := contract.Input{
		Has:      map[string]bool{},
		Argument: map[string]string{},
		Option:   map[string]string{},
	}
	// 如果有配置文件，覆盖option默认值（扩展名.ini）
	if hasFile(path + name + ".ini") {
		defConfig, err := ini.Load(path + name + ".ini")
		if err == nil {
			for i, kv := range configure.Input.Option {
				has := defConfig.Section("option").HasKey(kv.Name)
				if has {
					configure.Input.Option[i].Default = defConfig.Section("option").Key(kv.Name).String()
				}
			}
		}
	}
	// 从命令行解析参数
	err := input.Parsed(configure.Input)
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

func hasFile(file string) bool {
	_, err := os.Stat(file)
	if err != nil {
		// no such file or dir
		return false
	}
	return true
}
