package server

import (
    "database/sql"
    "database/sql/driver"
    "encoding/json"
    "errors"
    "log"
	"github.com/kenael/simpledeviceinventory/env"
	// "github.com/zcalusic/sysinfo"

    _ "github.com/lib/pq"
)

type System struct {
    MachineID    string
    Attrs Attrs
}

type Attrs map[string]interface{}

func (a Attrs) Value() (driver.Value, error) {
    return json.Marshal(a)
}

func (a *Attrs) Scan(value interface{}) error {
    b, ok := value.([]byte)
    if !ok {
        return errors.New("type assertion to []byte failed")
    }

    return json.Unmarshal(b, &a)
}

func StoreSystemValues(machineID string, payload []byte) error {
    db, err := sql.Open("postgres", "postgres://root:testtest22@192.168.0.15/test01?sslmode=disable")
    if err != nil {
        log.Fatal(err)
    }
	defer db.Close()
	system := new(System)
	json.Unmarshal([]byte(payload), &system.Attrs)
	
	log.Printf("test %v", system.Attrs)
	system.MachineID = machineID
    _, err = db.Exec("INSERT INTO system (machineid,attrs) VALUES($1, $2) ON CONFLICT (machineid) DO UPDATE SET attrs = $2;",system.MachineID, system.Attrs)
    if err != nil {
        log.Fatal(err)
    }

    return err
}

func StorePackagesValues(machineID string, payload []env.Package) error {
    db, err := sql.Open("postgres", "postgres://root:testtest22@192.168.0.15/test01?sslmode=disable")
    if err != nil {
        log.Fatal(err)
    }
	defer db.Close()
	system := new(System)
	system.Attrs = Attrs{"installed": payload}
	
	system.MachineID = machineID
    _, err = db.Exec("INSERT INTO packages (machineid,attrs) VALUES($1, $2) ON CONFLICT (machineid) DO UPDATE SET attrs = $2;",system.MachineID, system.Attrs)
	
    if err != nil {
        log.Fatal(err)
    }

    return err
}


func StoreSystemUser(machineID string, payload []byte) error {
    db, err := sql.Open("postgres", "postgres://root:testtest22@192.168.0.15/test01?sslmode=disable")
    if err != nil {
        log.Fatal(err)
    }
	defer db.Close()

	var users []string

	system := new(System)
	json.Unmarshal([]byte(payload), &users)
	system.Attrs = Attrs{"users": users}

	system.MachineID = machineID
    _, err = db.Exec("INSERT INTO systemuser (machineid,attrs) VALUES($1, $2) ON CONFLICT (machineid) DO UPDATE SET attrs = $2;",system.MachineID, system.Attrs)
    if err != nil {
        log.Fatal(err)
    }

    return err
}
