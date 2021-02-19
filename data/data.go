package main

import (
	"bufio"
	"github.com/imroc/req"
	"log"
	"os"
	"regexp"
)

const (
	texliveSpecUrl = `https://raw.githubusercontent.com/clearlinux-pkgs/texlive/master/texlive.spec`
	specFile       = "texlive.spec"
	dataFile       = `data.txt`
	regexPattern   = `^/.*?texmf-dist/(.*?)$`
)

func main() {
	removeFiles()

	downloadData()

	trimData()
}

// downloadData download the TeXLive files list
func downloadData() {
	r, err := req.Get(texliveSpecUrl)
	if err != nil {
		log.Fatal(err)
	}

	if err = r.ToFile(specFile); err != nil {
		log.Fatal(err)
	}

	log.Printf("%v downloaded\n", specFile)
}

// trimData processes the files list and convert it to data.txt
func trimData() {
	file, err := os.Open(specFile)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err = file.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	f, err := os.OpenFile(dataFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("processing")
	scanner := bufio.NewScanner(file)

	// read texlive.spec line by line and deal with it
	for scanner.Scan() {
		re := regexp.MustCompile(regexPattern)
		// if not matched, next line
		if !re.Match(scanner.Bytes()) {
			continue
		}

		// if matched, extract the data needed
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

// removeFiles remove texlive.spec and data.txt
func removeFiles() {
	var err error

	// remove texlive.spec
	if err = os.Remove(specFile); err != nil {
		log.Println(err)
	}
	log.Printf("%v removed\n", specFile)

	// remove data.txt
	if err = os.Remove(dataFile); err != nil {
		log.Println(err)
	}
	log.Printf("%v removed\n", dataFile)
}
