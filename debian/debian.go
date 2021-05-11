package debian

import (
	"github.com/kenael/go-dpkg"
	//"fmt"
	"encoding/json"
	"strings"
)

type Package struct {
	Package string `json:"package"`
	Version string `json:"version"`
	Status string `json:"status"`
}

func GetPackagesJSON(file string) ([]byte , error ) {
	getPackage := []Package{}

	pack := Package{}

	packages, err := dpkg.ReadPackagesFromFile(file)
	if err != nil {
		//fmt.Printf("Error: %v\n", err)
		return nil, err
	}

	// Need only necassary Items
	for _, pkg := range packages {
	//	fmt.Printf("%-32s %-8s %s \n", pkg.Package, pkg.Version, pkg.Status)
		pack.Package = pkg.Package
		pack.Version = pkg.Version
		inst := strings.Split(pkg.Status, " ")
		pack.Status = inst[2]
		getPackage = append(getPackage, pack)

	}

	//fmt.Printf("Read %d packages\n", len(packages))
	//fmt.Printf("Read %d getPackage\n", len(getPackage))
	packJson,_ := json.Marshal(getPackage)
	//fmt.Println(string(packJson))
	return packJson, nil
}