package main

import (
	_ "embed"
	"log"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
)

//go:embed data.txt
var Files string

// GetPackages returns all the packages need to install
func GetPackages(p map[string]bool) map[string]bool {
	// load *database*
	list := strings.Split(Files, "\n")

	// get installed packages
	installedPackages := GetInstalledPackages()

	packages := map[string]bool{}

	for _, f := range list {
		basename := path.Base(f)
		name := strings.TrimSuffix(basename, filepath.Ext(basename))

		dir := path.Dir(f)
		pkg := path.Base(dir)

		// 如果此时的文件名 在 传进来的“包名”里
		if _, ok := p[name]; ok {
			if _, found := packages[pkg]; found {
				continue
			}

			if _, found := installedPackages[pkg]; found {
				continue
			}

			packages[pkg] = true
		}
	}

	// 最后再把一些特例删除掉
	for _, p := range notPackages() {
		delete(packages, p)
	}

	return packages
}

// GetInstalledPackages 获得本地已经安装了的包
func GetInstalledPackages() map[string]bool {
	output, err := exec.Command(
		"tlmgr",
		"info",
		"--only-installed",
		"--data",
		"name",
	).Output()

	if err != nil {
		log.Fatal(err)
	}

	// if on windows, `\r\n` should be converted to `\n`
	s := strings.ReplaceAll(string(output), "\r\n", "\n")
	packages := make(map[string]bool)

	p := strings.Split(s, "\n")
	for _, s := range p {
		packages[s] = true
	}

	return packages
}

// notPackages is a *blacklist* for packages
// need time to collect
func notPackages() []string {
	return []string{
		"base",
		"config",
		"tools",
	}
}
