package main

import (
	"regexp"
	"strings"
)

// ParseFile reads a string and returns packages parsed from it
func ParseFile(s string) map[string]bool {
	// 首先把接收到的 string 全部转为小写，以防正则匹配失败
	s = strings.ToLower(s)

	// reUsePackage 是从 `\usepackage[options]{pkg}` 命令中匹配出 pkg
	// `(?m)^` 用于匹配行首
	reUsePackage := `(?m)^ *\\usepackage(\[.*?\])?{(.*?)}`
	// reRequirePackage 从 `\RequirePackage[option list]{package name}[release date]` 中匹配出 pkg
	reRequirePackage := `(?m)^ *\\requirepackage(\[.*?\])?{(.*?)}`
	// reRequirePackageWithOptions 从 `\RequirePackageWithOptions{package name}[release date]` 中匹配出 pkg
	// 虽然此 command 并没有 `[options]`，但为了 parseFromRegex 的兼容性，我们扔保留正则中的 `(\[.*?\])?`部分
	reRequirePackageWithOptions := `(?m)^ *\\requirepackagewithoptions(\[.*?\])?{(.*?)}`

	res := []string{reUsePackage, reRequirePackage, reRequirePackageWithOptions}

	packages := map[string]bool{}

	for _, re := range res {
		r := regexp.MustCompile(re)
		matches := r.FindAllStringSubmatch(s, -1)

		for _, match := range matches {
			packages[match[2]] = true
		}
	}

	return packages
}

func getPackagesNeedInstall(s string) map[string]bool {
	p := ParseFile(s)
	packages := GetPackages(p)

	return packages
}
