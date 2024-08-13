# Git 仓库文件合并工具

这是一个用 Go 编写的命令行工具，用于从指定的 Git 仓库中克隆或拉取文件，并合并指定后缀的文件到一个单独的文件中。

## 功能

- 从指定的 Git 仓库克隆或拉取文件。
- 合并指定后缀的文件到一个单独的文件中。
- 支持多个文件后缀，用逗号分隔。

## 使用方法

### 安装

1. 确保你已经安装了 Go 环境。
2. 克隆本仓库或下载源码。
3. 在项目目录下运行 `go build` 生成可执行文件。

### 命令行参数

- `-repo`：Git 仓库的 URL，必需参数。
- `-ext`：要合并的文件后缀，多个后缀用逗号分隔，默认为 `md`。

### 示例

```sh
./git-merge-tool -repo https://github.com/user/repo.git -ext md,txt
```

上述命令将从 `https://github.com/user/repo.git` 克隆或拉取仓库，并合并所有 `.md` 和 `.txt` 文件到一个名为 `merged.md` 的文件中。

## 代码结构

- `main.go`：主程序入口，处理命令行参数和逻辑流程。
- `README.md`：本说明文档。

## 主要函数

- `dirExists(dir string) bool`：检查目录是否存在。
- `getRepoURL(dir string) (string, error)`：获取仓库的 URL。
- `cloneRepo(repoURL, dir string) error`：克隆 Git 仓库到指定目录。
- `pullRepo(dir string) error`：执行 `git pull` 操作。
- `mergeFilesWithExtensions(dir, outputFile string, extensions []string) error`：遍历目录中的所有指定后缀的文件并合并。

## 依赖

- Go 语言标准库

## 许可证

本项目采用 MIT 许可证，详情请参见 [LICENSE](LICENSE) 文件。

## 贡献

欢迎提交 Issue 和 Pull Request。