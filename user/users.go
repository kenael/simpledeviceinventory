package user    

import (
    "fmt"
    "path/filepath"
	"syscall"
    "os"
	"os/user"
	"strconv"
	"encoding/json"
)

var systemuser = []*user.User{}


func GetKnownUser() ([]byte, error) {
	var UID int
	files, err := filepath.Glob("/home/*")
    if err != nil {
		fmt.Printf("ERROR  /home exists? ", err)
        return nil, err
    }
    
	for _, file := range files {
		if file != "/home/lost+found" {
			info, _ := os.Stat(file)
			if stat, ok := info.Sys().(*syscall.Stat_t); ok {
				UID = int(stat.Uid)
				
			} else {
				// we are not in linux, this won't work anyway in windows, 
				// but maybe you want to log warnings
				UID = os.Getuid()
			}
			lookupuser, err := user.LookupId(strconv.Itoa(int(UID)))
			if err == nil {
				systemuser = append(systemuser, lookupuser)
			}
			// Debug
			// fmt.Printf("Entry: %v UID: %v Username: %s\n", file, UID, systemuser.Username)
		}
	}
	dataUser, err := json.Marshal(systemuser)
	if err != nil {
		fmt.Println("ERROR Encode Json! ",err)
	}
	return dataUser, nil
}