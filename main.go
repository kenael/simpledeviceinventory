package main

import (
	"fmt"
	debian "github.com/kenael/simpledeviceinventory/debian"
)

func main() {
	pack, err := debian.GetPackagesJSON("debian/status")
	if err != nil {
		fmt.Println("ERROR", err)
	}
	fmt.Println(string(pack))
	
	
}