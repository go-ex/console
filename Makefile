# 编译命令插件
build_plugin:
	go build -buildmode=plugin -o script/plugin/hello.so plugin/hello/main.go

# 编译命令插件
build_plugin_debug:
	go build -buildmode=plugin -gcflags="all=-N -l" -o script/plugin/hello.so plugin/hello/main.go


