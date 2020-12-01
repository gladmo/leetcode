# leetcode
leetcode 本地化题库，可本地刷题/测试/调试

在线执行 golang 代码跑算法，可以使用
[Golang 在线代码执行](https://runcode.6cm.co/go): https://runcode.6cm.co/go

本项目只包含问题及原始网站提供的代码片断，想看题解及代码实现的可查看源站题解或者：[LeetCode-Go](https://github.com/halfrost/LeetCode-Go)

本地刷题及单步调试算法是本仓库存在的意义。

## Intro
本仓库是 [`leetcode-cn.com`](https://leetcode-cn.com) 本地化题库，包含本地化测试用例、常用 go 工具包的一个仓库。

本仓库主要包含三个功能：

- leetcode 在线题目本地化 (questions 目录)
- 本地化测试用例，本地调试 (cmd 目录)
- 常用刷题 go 语言工具包 (utils 目录)

## 主要功能
### 本地化测试
```shell script
./leetcode test 133
# or
./leetcode test clone-graph
# or 
./leetcode test https://leetcode-cn.com/problems/clone-graph/
```

可以得到测试结果:
```shell script
+------+--------------------------------+----------+----------+
| 序号  |            用例名称             | 测试状态  | 失败原因  |
+------+--------------------------------+----------+----------+
| 1/4  | test-nil                       | PASS     |          |
| 2/4  | test-empty                     | PASS     |          |
| 3/4  | test-[[2],[1]]                 | PASS     |          |
| 4/4  | test-[[2,4],[1,3],[2,4],[1,3]] | PASS     |          |
+------+--------------------------------+----------+----------+
```

### 本地化 leetcode-cn.com 题目
```shell script
./leetcode get 133
# or
./leetcode get clone-graph
# or 
./leetcode get https://leetcode-cn.com/problems/clone-graph/
```

得到输出：
```shell script
+----------+------------------------------+
| 标题      | Clone Graph                  |
+----------+------------------------------+
| 问题ID    |                          133 |
+----------+------------------------------+
| 中文标题  | 克隆图                        |
+----------+------------------------------+
| 标签      | 深度优先搜索,广度优先搜索,图     |
+----------+------------------------------+
| 难度      | 中等                         |
+----------+------------------------------+
```

## Q & A
```
Q: 为何会有这个仓库？
A: 刚开始刷题后发觉大部分时间花在了构造测试用例上，我觉得这个是可以大家共享的。而在本地可以跑测试，
对于解错题时的逐行调试代码尤为重要，对于加深对算法的理解会有帮助。

Q: 为何现在只有 golang 的测试用例？
A: 本仓库是为了我辅助刷题使用，欢迎其它语言刷题者提 PR 丰富仓库。

Q: 为何仓库的题目不全？
A: 虽然可以使用 ./leetcode get ... 将题目全本地化，虽然工具已经很完善可以生成大部分
的 golang 代码，但是所有测试用例还需要手动来创建。我在努力刷题中，每刷过的题目都会将测
试用例一并更新到仓库中。欢迎大家一起贡献你的测试用例。

Q: 我是刷题者，这个仓库对我有何作者？
A: 首先，如果你使用 golang 刷题时，按题号检索时， questions/serial 目录包含了 简单/中等/困难 
三个层次的问题，你可以选择题号做题。按标签刷题的朋友可以在 questions/tags 找到不同标签的题目，
注意在测试的时候，需要使用 (leetcode test --tag [you-choose-tag] ... )
需要带上 --tag  [you-choose-tag]

Q: 我想贡献代码，我应该怎么做？
A: 首先感谢你对项目的支持。
1. 现有问题 markdown 或者测试用例修改
已有问题文件夹下的 README.md(问题 markdown 文件) 或问题下的golang文件夹中的 main.go(测试用例)，
提交前使用 leetcode clear 将 solution 代码复原，！！！这个操作会将 serial 目录下的所有 solution 
目录下的代码恢复到与 leetcode-cn.com 提供的代码片断一致（这么做的目录是为了给刷题者创建一个原生的刷题环境，
而且大家写的代码在 solution 也会反复覆盖）

2. 与 leetcode-cn.com 交互代码部分可在 leet 目录修改
3. leetcode 命令行问题可在 cmd/command 目录修改
4. 添加刷题工具库的可以添加到 utils 目录下，添加前需要先确认功能还不存在。

再次感谢！
```

## 快速开始
```shell script
git --depth 1 clone https://github.com/gladmo/leetcode.git
cd leetcode
go build -o leetcode cmd/leetcode.go
./leetcode
```

你可以看到如下输出
```shell script
leetcode cli

Usage:
  leetcode [command]

Available Commands:
  backup      backup you complete questions to solutions
  base        clear & replace all question use you specified (backup all unbanked)
  clear       set questions to default (backup all unbanked)
  get         get leet question from leetcode-cn.com
  help        Help about any command
  info        print leet question info
  test        test you code and analyse
  version     Print the version number of leetcode cli

Flags:
  -h, --help   help for leetcode

Use "leetcode [command] --help" for more information about a command.
```

本仓库是作者刷题目过程中不断完善的一个仓库，后续投入开发的时间将会减少，一起刷题目的小伙伴可以一起来贡献你的测试用例。
### 如果你是一名刷题者

你只需要使用：

`leetcode test ...` 解答完问题，跑测试用例
```shell script
+------+--------------------------------+----------+----------+
| 序号  |            用例名称             | 测试状态  | 失败原因  |
+------+--------------------------------+----------+----------+
| 1/4  | test-nil                       | PASS     |          |
| 2/4  | test-empty                     | PASS     |          |
| 3/4  | test-[[2],[1]]                 | PASS     |          |
| 4/4  | test-[[2,4],[1,3],[2,4],[1,3]] | PASS     |          |
+------+--------------------------------+----------+----------+
```

`leetcode info ...` 查看问题基础信息
```shell script
+----------+------------------------------+
| 标题      | Clone Graph                  |
+----------+------------------------------+
| 问题ID    |                          133 |
+----------+------------------------------+
| 中文标题  | 克隆图                        |
+----------+------------------------------+
| 标签      | 深度优先搜索,广度优先搜索,图     |
+----------+-----------------------------+
| 难度      | 中等                         |
+----------+------------------------------+
```

### 如果你是一名贡献者
```shell script
首先感谢你对项目的支持。
1. 现有问题 markdown 或者测试用例修改
已有问题文件夹下的 README.md(问题 markdown 文件) 或问题下的golang文件夹中的 main.go(测试用例)，
提交前使用 leetcode clear 将 solution 代码复原，！！！这个操作会将 serial 目录下的所有 solution 
目录下的代码恢复到与 leetcode-cn.com 提供的代码片断一致（这么做的目录是为了给刷题者创建一个原生的刷题环境，
而且大家写的代码在 solution 也会反复覆盖）

2. 与 leetcode-cn.com 交互代码部分可在 leet 目录修改
3. leetcode 命令行问题可在 cmd/command 目录修改
4. 添加刷题工具库的可以添加到 utils 目录下，添加前需要先确认功能还不存在。

再次感谢！
```

## TODO
- [x] leetcode backup 命令
- [ ] leetcode login
- [ ] leetcode publish
