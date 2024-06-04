```bash
$ pwd
/Users/huzhi/work/code/go_code/cobra/cobra-demo

$ git init .
提示：使用 'master' 作为初始分支的名称。这个默认分支名称可能会更改。要在新仓库中
提示：配置使用初始分支名，并消除这条警告，请执行：
提示：
提示：	git config --global init.defaultBranch <名称>
提示：
提示：除了 'master' 之外，通常选定的名字有 'main'、'trunk' 和 'development'。
提示：可以通过以下命令重命名刚创建的分支：
提示：
提示：	git branch -m <name>
已初始化空的 Git 仓库于 /Users/huzhi/work/code/go_code/cobra/cobra-demo/.git/

$ tree -a .
.
├── .git
│   ├── HEAD
│   ├── config
│   ├── description
│   ├── hooks
│   │   ├── applypatch-msg.sample
│   │   ├── commit-msg.sample
│   │   ├── fsmonitor-watchman.sample
│   │   ├── post-update.sample
│   │   ├── pre-applypatch.sample
│   │   ├── pre-commit.sample
│   │   ├── pre-merge-commit.sample
│   │   ├── pre-push.sample
│   │   ├── pre-rebase.sample
│   │   ├── pre-receive.sample
│   │   ├── prepare-commit-msg.sample
│   │   ├── push-to-checkout.sample
│   │   ├── sendemail-validate.sample
│   │   └── update.sample
│   ├── info
│   │   └── exclude
│   ├── objects
│   │   ├── info
│   │   └── pack
│   └── refs
│       ├── heads
│       └── tags
└── README.md

10 directories, 19 files
$

$ git status
位于分支 master

尚无提交

未跟踪的文件:
  （使用 "git add <文件>..." 以包含要提交的内容）
	README.md

提交为空，但是存在尚未跟踪的文件（使用 "git add" 建立跟踪）

$ go mod init github.com/lanzhiwang/cobra-demo
go: creating new go.mod: module github.com/lanzhiwang/cobra-demo

$ git status
位于分支 master

尚无提交

未跟踪的文件:
  （使用 "git add <文件>..." 以包含要提交的内容）
	README.md
	go.mod

提交为空，但是存在尚未跟踪的文件（使用 "git add" 建立跟踪）
$

$ go mod init github.com/lanzhiwang/cobra-demo
go: creating new go.mod: module github.com/lanzhiwang/cobra-demo
$

#######################################################################################3

$ go install github.com/spf13/cobra-cli@v1.3.0
go: downloading github.com/spf13/cobra-cli v1.3.0
go: downloading github.com/spf13/cobra v1.3.0
go: downloading gopkg.in/ini.v1 v1.66.2
go: downloading github.com/spf13/afero v1.6.0
go: downloading golang.org/x/sys v0.0.0-20211210111614-af8b64212486
$

$ cobra-cli
Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.

Usage:
  cobra-cli [command]

Available Commands:
  add         Add a command to a Cobra Application
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  init        Initialize a Cobra Application

Flags:
  -a, --author string    author name for copyright attribution (default "YOUR NAME")
      --config string    config file (default is $HOME/.cobra.yaml)
  -h, --help             help for cobra-cli
  -l, --license string   name of license for the project
      --viper            use Viper for configuration

Use "cobra-cli [command] --help" for more information about a command.

$ cobra-cli init --help
Initialize (cobra init) will create a new application, with a license
and the appropriate structure for a Cobra-based CLI application.

Cobra init must be run inside of a go module (please run "go mod init <MODNAME>" first)

Usage:
  cobra-cli init [path] [flags]

Aliases:
  init, initialize, initialise, create

Flags:
  -h, --help   help for init

Global Flags:
  -a, --author string    author name for copyright attribution (default "YOUR NAME")
      --config string    config file (default is $HOME/.cobra.yaml)
  -l, --license string   name of license for the project
      --viper            use Viper for configuration
$

# $ cobra-cli init -a hzhilamp@163.com -l MIT
$ cobra-cli init
Your Cobra application is ready at
/Users/huzhi/work/code/go_code/cobra/cobra-demo
$

$ git status
位于分支 master

尚无提交

未跟踪的文件:
  （使用 "git add <文件>..." 以包含要提交的内容）
	LICENSE
	README.md
	cmd/
	go.mod
	go.sum
	main.go

提交为空，但是存在尚未跟踪的文件（使用 "git add" 建立跟踪）
$

$ tree -a .
.
├── LICENSE
├── README.md
├── cmd
│   └── root.go
├── go.mod
├── go.sum
└── main.go

17 directories, 31 files
$

$ go build main.go

$ ./main
A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.


$ ./main help
A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.


$ ./main -h
A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.

$

#####################################################################

$ cobra-cli add version
version created at /Users/huzhi/work/code/go_code/cobra/cobra-demo

$ git status
位于分支 master

尚无提交

要提交的变更：
  （使用 "git rm --cached <文件>..." 以取消暂存）
	新文件：   LICENSE
	新文件：   README.md
	新文件：   cmd/root.go
	新文件：   go.mod
	新文件：   go.sum
	新文件：   main.go

未跟踪的文件:
  （使用 "git add <文件>..." 以包含要提交的内容）
	cmd/version.go

$

$ go build main.go

$ ./main
A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.

Usage:
  cobra-demo [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  version     A brief description of your command

Flags:
  -h, --help     help for cobra-demo
  -t, --toggle   Help message for toggle

Use "cobra-demo [command] --help" for more information about a command.

$ ./main help
A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.

Usage:
  cobra-demo [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  version     A brief description of your command

Flags:
  -h, --help     help for cobra-demo
  -t, --toggle   Help message for toggle

Use "cobra-demo [command] --help" for more information about a command.

$ ./main version
version called
$

#####################################################################

$ cobra-cli add subconmand1
subconmand1 created at /Users/huzhi/work/code/go_code/cobra/cobra-demo

$ git status
位于分支 master

尚无提交

要提交的变更：
  （使用 "git rm --cached <文件>..." 以取消暂存）
	新文件：   LICENSE
	新文件：   README.md
	新文件：   cmd/root.go
	新文件：   cmd/version.go
	新文件：   go.mod
	新文件：   go.sum
	新文件：   main.go

未跟踪的文件:
  （使用 "git add <文件>..." 以包含要提交的内容）
	cmd/subconmand1.go

$



```
