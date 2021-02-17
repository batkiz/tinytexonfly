package main

import (
	"log"
	"os"
	"path/filepath"
)

// ReadFile 接收一个字符串 filename 作为文件名，读取该文件内容并以 string 返回
func ReadFile(filename string) string {
	s, err := os.ReadFile(filename)
	if err != nil {
		log.Println(err)
	}
	return string(s)
}

// ReadPath 接收一个 string path 为文件夹路径，读取该文件夹内所有内容并以 string 返回
func ReadPath(path string) string {
	var files []string
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			// 仅读取文件类型为 tex, dtx, cls 的文件
			switch filepath.Ext(info.Name()) {
			case "tex",
				"dtx",
				"cls":
				files = append(files, path)
			}
		}
		return nil
	})
	if err != nil {
		log.Println(err)
	}

	s := ""

	for _, file := range files {
		fileContent := ReadFile(file)
		s += fileContent
	}

	return s
}
