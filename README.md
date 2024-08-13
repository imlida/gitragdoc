# Git仓库文件合并工具

## 简介

这是一个用Go语言编写的命令行工具，用于克隆或更新Git仓库，并将仓库中的特定文件类型或所有文本文件合并到一个输出文件中。

## 功能

- 克隆指定的Git仓库
- 如果仓库已存在，则更新（pull）仓库
- 合并仓库中的特定文件类型（如.md文件）
- 支持合并仓库中的所有文本文件
- 将合并后的内容保存到一个输出文件中

## 使用方法

### 编译

首先，确保你的系统已安装Go语言环境。然后，在项目目录下运行以下命令编译程序：

```bash
go build -o git-repo-merger main.go
```

### 运行

使用以下命令行参数运行程序：

```bash
./git-repo-merger -repo <仓库URL> -ext <文件扩展名>
```

参数说明：
- `-repo`：Git仓库的URL（必需）
- `-ext`：要合并的文件扩展名，用逗号分隔。使用 "*" 表示合并所有文本文件（可选，默认为 "md"）

### 示例

1. 合并仓库中所有的Markdown文件：

```bash
./git-repo-merger -repo https://github.com/username/repo.git -ext md
```

2. 合并仓库中所有的文本文件：

```bash
./git-repo-merger -repo https://github.com/username/repo.git -ext "*"
```

3. 合并仓库中的Markdown和TXT文件：

```bash
./git-repo-merger -repo https://github.com/username/repo.git -ext md,txt
```

## 输出

程序将在当前目录下创建以下内容：

- `gitrepo` 文件夹：包含克隆或更新的Git仓库
- `merged.md` 文件：包含所有合并后的文件内容

## 注意事项

- 确保你有足够的权限来克隆指定的Git仓库
- 如果指定的仓库已经存在于本地，程序将尝试更新它而不是重新克隆
- 合并大型仓库或包含大量文件的仓库可能需要一些时间
- 使用 "*" 作为文件扩展名时，程序将尝试合并所有被检测为文本文件的文件

## 贡献

欢迎提交问题报告和合并请求来帮助改进这个工具。