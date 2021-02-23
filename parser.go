package main

import (
	"regexp"
	"strings"
)

// PackageParser 用于实现一个取出 package 的解析器
type PackageParser struct {
	// 用于匹配的正则表达式
	re string
	// package 出现在正则表达式 match 中的索引
	index int
	// 对拿出来的 match 进行何种处理
	postParse func(string) []string
}

// ParseFile reads a string and returns packages parsed from it
func ParseFile(s string) map[string]bool {
	s = strings.ToLower(s)

	// add more parsers!
	parsers := []PackageParser{
		{
			// reUsePackage 是从 `\usepackage[options]{pkg}` 命令中匹配出 pkg
			// `(?m)^` 用于匹配行首
			re:    `(?m)^ *\\usepackage(\[.*?\])?{(.*?)}`,
			index: 2,
			postParse: func(s string) []string {
				// if \usepackage{pkg1,pkg2,...}
				if strings.Contains(s, ",") {
					p := strings.Split(s, ",")
					for i, v := range p {
						// remove spaces
						p[i] = strings.ReplaceAll(v, " ", "")
					}
					return p
				}
				return []string{s}
			},
		},
		{
			// reRequirePackage 从 `\RequirePackage[option list]{package name}[release date]` 中匹配出 pkg
			re:    `(?m)^ *\\requirepackage(\[.*?\])?{(.*?)}`,
			index: 2,
			postParse: func(s string) []string {
				return []string{s}
			},
		},
		{
			// reRequirePackageWithOptions 从 `\RequirePackageWithOptions{package name}[release date]` 中匹配出 pkg
			re:    `(?m)^ *\\requirepackagewithoptions{(.*?)}`,
			index: 1,
			postParse: func(s string) []string {
				return []string{s}
			},
		},
		{
			re:    `(?m)^ *\\documentclass{(.*?)}`,
			index: 1,
			postParse: func(s string) []string {
				return []string{s}
			},
		},
	}

	packages := map[string]bool{}

	for _, parser := range parsers {
		r := regexp.MustCompile(parser.re)
		matches := r.FindAllStringSubmatch(s, -1)

		for _, match := range matches {
			m := match[parser.index]
			p := parser.postParse(m)

			for _, s := range p {
				packages[s] = true
			}
		}
	}

	return packages
}

/*
// ParseFile reads a string and returns packages parsed from it
func ParseFileOld(s string) map[string]bool {
	// 首先把接收到的 string 全部转为小写
	s = strings.ToLower(s)

	const (
		// reUsePackage 是从 `\usepackage[options]{pkg}` 命令中匹配出 pkg
		// `(?m)^` 用于匹配行首
		reUsePackage = `(?m)^ *\\usepackage(\[.*?\])?{(.*?)}`
		// reRequirePackage 从 `\RequirePackage[option list]{package name}[release date]` 中匹配出 pkg
		reRequirePackage = `(?m)^ *\\requirepackage(\[.*?\])?{(.*?)}`
		// reRequirePackageWithOptions 从 `\RequirePackageWithOptions{package name}[release date]` 中匹配出 pkg
		// 虽然此 command 并没有 `[options]`，但为了 parseFromRegex 的兼容性，我们扔保留正则中的 `(\[.*?\])?`部分
		reRequirePackageWithOptions = `(?m)^ *\\requirepackagewithoptions(\[.*?\])?{(.*?)}`
	)

	var (
		res = []string{reUsePackage, reRequirePackage, reRequirePackageWithOptions}

		packages = map[string]bool{}
	)

	for _, re := range res {
		r := regexp.MustCompile(re)
		matches := r.FindAllStringSubmatch(s, -1)

		for _, match := range matches {
			packages[match[2]] = true
		}
	}

	return packages
}
*/

func getPackagesNeedInstall(s string) map[string]bool {
	p := ParseFile(s)
	packages := GetPackages(p)

	return packages
}
