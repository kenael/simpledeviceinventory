package main

import (
	"fmt"
	debian "github.com/kenael/simpledeviceinventory/debian"
	"github.com/kenael/simpledeviceinventory/system"
	"github.com/kenael/simpledeviceinventory/user"
	"github.com/kenael/simpledeviceinventory/client"
)

func main() {

	// Get Installed Packages
	// Ubuntu Default: /var/lib/dpkg/status
	pack, err := debian.GetPackagesJSON("/var/lib/dpkg/status")
	if err != nil {
		fmt.Println("ERROR", err)
	}
	fmt.Println(string(pack))
	
	machineID, system, err := system.GetSystemInfo() 
	if err != nil {
		fmt.Println("Error ", err)
	}
	fmt.Println(machineID)
	fmt.Println(string(system))

	user, err := user.GetKnownUser() 
	if err != nil {
		fmt.Println(err)

	}
	fmt.Println(string(user))

	// test
	client.SendSystem( machineID, system)
	client.SendPackages(machineID, pack)

}