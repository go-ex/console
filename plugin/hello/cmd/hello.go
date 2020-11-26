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
			Has: nil,
			Argument: []contract.ArgArgument{
				{
					Name:        "name",
					Description: "问候名字",
				},
			},
			Option: nil,
		},
	}
}

func (h *Hello) Execute(input contract.Input) {
	fmt.Println("hello " + input.GetArgument("name"))
	fmt.Println(input)
}
