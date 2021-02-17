package main

import (
	"fmt"
)

// TODO: really executing the command
func ExecTlmgrInstall(packages map[string]bool) error {
	if len(packages) == 0 {
		fmt.Println("nothing to install, congratulations!")
		return nil
	}
	list := ""
	for k := range packages {
		list += k + " "
	}

	fmt.Printf("tlmgr install %v\n", list)
	/*
		out, err := exec.Command("tlmgr", "install", stringArrayToString(packages)).Output()

		if err != nil {
			return err
		}
		fmt.Printf("%v", string(out))
	*/
	return nil
}
