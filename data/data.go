package main

import (
	"bufio"
	"github.com/imroc/req"
	"log"
	"os"
	"regexp"
)

var (
	texliveSpecUrl = `https://raw.githubusercontent.com/clearlinux-pkgs/texlive/master/texlive.spec`
	dataPath       = `data.txt`
	regexPattern   = `^/.*?texmf-dist/(.*?)$`
)

func main() {
	removeFiles()

	downloadData()

	trimData()
}

func downloadData() {
	r, err := req.Get(texliveSpecUrl)
	if err != nil {
		log.Fatal(err)
	}

	err = r.ToFile("texlive.spec")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("texlive.spec downloaded")
}

func trimData() {
	file, err := os.Open("texlive.spec")
	if err != nil {
		log.Fatal(err)
	}
	defer func() { err = file.Close() }()

	f, err := os.OpenFile(dataPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	log.Println("processing")

	for scanner.Scan() {
		re := regexp.MustCompile(regexPattern)
		if !re.Match(scanner.Bytes()) {
			continue
		}

		matches := re.FindAllStringSubmatch(scanner.Text(), -1)
		for _, match := range matches {
			if _, err := f.WriteString(match[1] + "\n"); err != nil {
				log.Fatal(err)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}

	log.Println("DONE!")
}

func removeFiles() {
	var err error
	err = os.Remove("texlive.spec")
	if err != nil {
		log.Println(err)
	}
	log.Println("texlive.spec removed")

	err = os.Remove("data.txt")
	if err != nil {
		log.Println(err)
	}
	log.Println("data.txt removed")
}
