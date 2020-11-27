package cmd

import (
	"fmt"
	"github.com/go-ex/console/contract"
)

type Hello string

func (h *Hello) Configure() contract.Configure {
	return contract.Configure{
		Description: "命令说明",
		Input: contract.Argument{
			// 必须参数 := command name
			Argument: []contract.ArgArgument{
				{
					Name:        "name",
					Description: "问候名字",
				},
			},
			// 可选参数 command name -age=18
			Option: []contract.ArgParam{
				{
					Name:        "age",
					Description: "年龄",
					Default:     "30",
				},
			},
			// bool 可选参数 command name -h | input.GetHas("-h")
			Has: nil,
		},
	}
}

func (h *Hello) Execute(input contract.Input) {
	fmt.Println("hello " + input.GetArgument("name"))
	fmt.Println("age " + input.GetOption("age"))
	fmt.Println(input)
}
