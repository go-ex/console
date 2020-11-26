package contract

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

// 参数操作
type Input struct {
	// 是否有参数 【名称string】默认值bool
	Has map[string]bool
	// 必须输入参数 【命令位置】【赋值名称】默认值
	Argument map[string]string
	// 可选输入参数 【赋值名称（开头必须是-）】默认值
	Option map[string]string
}

var BaseInputHas = map[string]ArgParam{
	"-d": {Name: "-d", Description: "守护进程启动"},
	"-h": {Name: "-h", Description: "显示帮助信息"},
}

// 参数解析
func (i *Input) Parsed(Config Argument) error {
	args := os.Args[2:]

	Config.Has = append(Config.Has, BaseInputHas["-d"])
	Config.Has = append(Config.Has, BaseInputHas["-h"])

	for _, ArgParam := range Config.Has {
		for _, strArg := range args {
			if ArgParam.Name == strArg {
				i.Has[ArgParam.Name] = true
			}
		}
		_, ok := i.Has[ArgParam.Name]
		if !ok {
			i.Has[ArgParam.Name] = false
		}
	}
	// 帮助参数 -h 不需要配置
	helpCmd := "-h"
	i.Has[helpCmd] = false
	for _, strArg := range args {
		if helpCmd == strArg {
			i.Has[helpCmd] = true
			return nil
		}
	}

	// 必须值
	lenArgument := len(args)
	for mustInt, kv := range Config.Argument {
		if lenArgument <= mustInt {
			// 不存在，报错,并且输出帮助命令
			fmt.Println("必须输入参数:" + kv.Name)
			return errors.New("必须输入参数:" + kv.Name)
		} else {
			i.Argument[kv.Name] = args[mustInt]
		}
	}
	// 选项值
	for _, kv := range Config.Option {
		i.Option[kv.Name] = kv.Default
	}
	var strArgKy, strValue string
	for _, strArg := range args {
		startIndex := strings.Index(strArg, "-")
		if startIndex == 0 {
			stopIndex := strings.Index(strArg, "=")
			if stopIndex < 0 {
				// 不存在 = 号
				strArgKy = strArg[startIndex+1:]
				defaultValue, _ := i.Option[strArgKy]
				strValue = defaultValue
			} else {
				strArgKy = strArg[startIndex+1 : stopIndex]
				strValue = strArg[stopIndex+1:]

			}
			if strArgKy != "" {
				i.Option[strArgKy] = strValue
			}
		}
	}
	return nil
}

// 参数
func (i *Input) GetHas(key string) bool {
	value, ok := i.Has[key]
	if !ok {
		return false
	}
	return value
}

// 参数
func (i *Input) GetArgument(key string) string {
	value, ok := i.Argument[key]
	if !ok {
		return ""
	}
	return value
}

// 参数
func (i *Input) GetOption(key string) string {
	value, ok := i.Option[key]
	if !ok {
		return ""
	}
	return value
}

// 是否后台启动
func (i *Input) IsDaemon() bool {
	value, ok := i.Has["-d"]
	if !ok {
		return false
	}
	return value
}
