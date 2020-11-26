package console

import (
	"fmt"
	"github.com/go-ex/console/contract"
	"io/ioutil"
	"plugin"
	"sort"
	"strconv"
)

// http://patorjk.com/software/taag/#p=testall&f=Big&t=Gocmd
var banner = `  _____                  __
 / ___/__  ______ _  ___/ /
/ (_ / _ \/ __/  ' \/ _  / 
\___/\___/\__/_/_/_/\_,_/  

`

func (c *Console) RunHelp() {
	fmt.Print(banner)

	fmt.Println("Usage:")
	fmt.Println("  command [options] [arguments] [has]")
	fmt.Println("Base Has Param:")
	fmt.Println("  -d  守护进程启动")
	fmt.Println("  -h  显示帮助信息参数")
	fmt.Println("Available commands:")

	// 命令排序
	arrCmd := map[string]contract.Command{}
	for _, path := range GetPluginPath() {
		fileInfoList, err := ioutil.ReadDir(path)
		if err != nil {
			continue
		}

		for _, file := range fileInfoList {
			fileName := file.Name()
			plug, err := plugin.Open(path + fileName)
			if err != nil {
				fmt.Println(err)
				continue
			}

			inf, err := plug.Lookup("Command")
			if err != nil {
				fmt.Println(err)
				continue
			}

			cmd := inf.(contract.Command)
			// 移除 .so 后缀
			arrCmd[fileName[:len(fileName)-3]] = cmd
		}
	}

	var keys []string
	var macLen int
	for cmdName := range arrCmd {
		keys = append(keys, cmdName)
		tempLen := len(cmdName)
		if tempLen > macLen {
			macLen = tempLen
		}
	}
	sort.Strings(keys)
	macLen += 4
	for _, cmdName := range keys {
		kv := arrCmd[cmdName]
		fmt.Println(echoSpace("  "+cmdName, macLen), kv.Configure().Description)
	}
}

// 字符串不足，空格补充输出
func echoSpace(str string, mac int) string {
	strCon := strconv.Itoa(mac)
	return fmt.Sprintf("%-"+strCon+"s", str)
}

// 某个命令需要帮助时
func (c *Console) RunCommandHelp(name string, con contract.Configure) error {
	fmt.Println("Usage:")
	fmt.Print("  ", name)
	for _, ArgParam := range con.Input.Argument {
		fmt.Print(" <", ArgParam.Name, ">")
	}
	fmt.Println()
	for _, ArgParam := range con.Input.Option {
		if len(ArgParam.Default) >= 1 {
			fmt.Println(echoSpace("    -"+ArgParam.Name, 25), "= "+ArgParam.Default)
		}
	}

	if len(con.Input.Argument) > 0 {
		fmt.Println("Arguments:")
		for _, ArgParam := range con.Input.Argument {
			fmt.Println(echoSpace("  "+ArgParam.Name, 25), ArgParam.Description)
		}
	}

	if len(con.Input.Option) > 0 {
		fmt.Println("Option:")
		for _, ArgParam := range con.Input.Option {
			fmt.Println(echoSpace("  -"+ArgParam.Name, 25), ArgParam.Description)
		}
	}

	if len(con.Input.Has) > 0 {
		fmt.Println("Has:")
		for _, ArgParam := range con.Input.Has {
			fmt.Println(echoSpace("  "+ArgParam.Name, 25), ArgParam.Description)
		}
	}

	fmt.Println("Description:")
	fmt.Println("  ", con.Description)

	return nil
}
