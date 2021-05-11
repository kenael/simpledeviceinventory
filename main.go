package main

import (
	"fmt"
	debian "github.com/kenael/simpledeviceinventory/debian"
	"github.com/kenael/simpledeviceinventory/system"
)

func main() {

	// Get Installed Packages
	pack, err := debian.GetPackagesJSON("/var/lib/dpkg/status")
	if err != nil {
		fmt.Println("ERROR", err)
	}
	fmt.Println(string(pack))
	
	system, err := system.GetSystemInfo() 
	if err != nil {
		fmt.Println("Error ", err)
	}
	fmt.Println(string(system))
}