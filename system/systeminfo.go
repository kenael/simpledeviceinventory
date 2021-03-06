package system

import (
	"encoding/json"
	
	"github.com/zcalusic/sysinfo"
)

func GetSystemInfo() (string, []byte, error) {
	var si sysinfo.SysInfo

	si.GetSysInfo()

	data, err := json.MarshalIndent(&si, "", "  ")
	if err != nil {
		return "", nil, err
	}

	return si.Node.MachineID, data, nil
}