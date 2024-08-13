package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	// 定义命令行参数
	repoURL := flag.String("repo", "", "URL of the Git repository")
	fileExtensions := flag.String("ext", "md", "File extensions to scan, separated by commas. Use '*' for all text files")

	// 解析命令行参数
	flag.Parse()

	// 检查必要的参数是否提供
	if *repoURL == "" {
		fmt.Println("Error: repository URL is required")
		flag.Usage()
		os.Exit(1)
	}

	// 获取当前目录
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Failed to get current directory: %v\n", err)
		return
	}

	// 设置临时目录和输出文件路径为当前目录
	tempDir := filepath.Join(currentDir, "gitrepo")
	outputFile := filepath.Join(currentDir, "merged.md")

	// 检查临时目录是否存在，并判断仓库地址是否一致
	if dirExists(tempDir) {
		repoURLInDir, err := getRepoURL(tempDir)
		if err != nil {
			fmt.Printf("Failed to get repository URL from temp directory: %v\n", err)
			return
		}
		if repoURLInDir == *repoURL {
			// 执行git pull操作
			err = pullRepo(tempDir)
			if err != nil {
				fmt.Printf("Failed to pull repository: %v\n", err)
				return
			}
		} else {
			// 删除临时目录并重新克隆仓库
			err = os.RemoveAll(tempDir)
			if err != nil {
				fmt.Printf("Failed to remove temp directory: %v\n", err)
				return
			}
			err = cloneRepo(*repoURL, tempDir)
			if err != nil {
				fmt.Printf("Failed to clone repository: %v\n", err)
				return
			}
		}
	} else {
		// 克隆仓库到临时目录
		err = cloneRepo(*repoURL, tempDir)
		if err != nil {
			fmt.Printf("Failed to clone repository: %v\n", err)
			return
		}
	}

	// 遍历仓库中的所有指定后缀的文件并合并
	var extensions []string
	if *fileExtensions == "*" {
		extensions = nil // Use nil to indicate all text files
	} else {
		extensions = strings.Split(*fileExtensions, ",")
	}
	err = mergeFilesWithExtensions(tempDir, outputFile, extensions)
	if err != nil {
		fmt.Printf("Failed to merge files: %v\n", err)
		return
	}

	fmt.Println("Files merged successfully!")
}

// 检查目录是否存在
func dirExists(dir string) bool {
	info, err := os.Stat(dir)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

// 获取仓库URL
func getRepoURL(dir string) (string, error) {
	cmd := exec.Command("git", "-C", dir, "config", "--get", "remote.origin.url")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(output)), nil
}

// 克隆Git仓库到指定目录
func cloneRepo(repoURL, dir string) error {
	cmd := exec.Command("git", "clone", repoURL, dir)
	return cmd.Run()
}

// 执行git pull操作
func pullRepo(dir string) error {
	cmd := exec.Command("git", "-C", dir, "pull")
	return cmd.Run()
}

// 遍历目录中的所有指定后缀的文件并合并
func mergeFilesWithExtensions(dir, outputFile string, extensions []string) error {
	var mergedContent strings.Builder

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			if extensions == nil {
				// If extensions is nil, check if it's a text file
				if isTextFile(path) {
					return mergeFile(path, &mergedContent)
				}
			} else {
				for _, ext := range extensions {
					if strings.HasSuffix(info.Name(), "."+ext) {
						return mergeFile(path, &mergedContent)
					}
				}
			}
		}
		return nil
	})

	if err != nil {
		return err
	}

	return os.WriteFile(outputFile, []byte(mergedContent.String()), 0644)
}

// 合并单个文件
func mergeFile(path string, mergedContent *strings.Builder) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(mergedContent, file)
	if err != nil {
		return err
	}
	mergedContent.WriteString("\n\n") // 添加分隔符
	return nil
}

// 检查文件是否为文本文件
func isTextFile(path string) bool {
	file, err := os.Open(path)
	if err != nil {
		return false
	}
	defer file.Close()

	// 读取文件的前512字节
	buffer := make([]byte, 512)
	n, err := file.Read(buffer)
	if err != nil && err != io.EOF {
		return false
	}

	// 使用http.DetectContentType来检测文件类型
	contentType := http.DetectContentType(buffer[:n])
	return strings.HasPrefix(contentType, "text/")
}
