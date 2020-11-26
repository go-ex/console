package contract

type Command interface {
	Configure() Configure
	Execute(input Input)
}

type Configure struct {
	// 说明
	Description string
	// 输入定义
	Input Argument
}

// 参数存储
type ArgParam struct {
	Name        string // 名称
	Description string // 说明
	Default     string // 默认值
}

// 参数存储
type ArgArgument struct {
	Name        string // 名称
	Description string // 说明
}

// 参数设置结构
type Argument struct {
	// 是否有参数 【名称string】
	Has []ArgParam
	// 必须输入参数 【命令位置】【赋值名称】默认值
	Argument []ArgArgument
	// 可选输入参数 【赋值名称（开头必须是-）】默认值
	Option []ArgParam
}
