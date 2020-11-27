# console
控制台工具，插件式的cli执行器; 

可以使用同一个命令，在不同项目目录下执行，可以使用不同的子命令执行（包括得到不同帮助命令）。

特别适合维护多个项目时，部分项目有有相通的处理，又有部分命令不一样处理。

# 使用

查看 https://github.com/go-ex/console/tree/main/example

# 配置
默认下，扩展加载会在以下两个目录寻找插件

主程序配置 {your_cli}.ini
~~~~
[plugin]
# 启动程序存储目录下的 root + root_path
root_path = /plugin
# 执行命令当前目录下的  worker + worker_path
worker_path = /script/plugin
~~~~

插件命令配置,查看 https://github.com/go-ex/console/tree/main/plugin/hello

编译插件
~~~~
go build -buildmode=plugin -o script/plugin/hello.so plugin/hello/main.go
~~~~

创建配置 script/plugin/hello.ini
~~~~
[option]
age=18
~~~~

执行命令 `./example/hello hello 张三` 输出的命令就是`18`

执行 `./example/hello hello 张三` 等价于 `./example/hello hello 张三 18` 

如果没有配置文件，那么就是 `30`;

# 其他
make run
~~~~
  _____                  __
 / ___/__  ______ _  ___/ /
/ (_ / _ \/ __/  ' \/ _  / 
\___/\___/\__/_/_/_/\_,_/  

Usage:
  command [options] [arguments] [has]
Base Has Param:
  -d  守护进程启动
  -h  显示帮助信息参数
Available commands:
  hello   命令说明

~~~~
帮助命令
~~~~
my@ctfang console % ./example/cli hello -h
Usage:
  hello <name>
    -age                  = 30
Arguments:
  name                    问候名字
Option:
  -age                    年龄
Description:
   命令说明
~~~~