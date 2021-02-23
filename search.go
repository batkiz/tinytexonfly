package main

import (
	"fmt"
	"strings"
)

func Search(s string) {
	list := strings.Split(Files, "\n")

	for _, f := range list {
		if strings.Contains(f, s) {
			fmt.Println(f)
		}
	}
}
