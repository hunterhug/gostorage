package storage

import (
	"fmt"
	"testing"
	"github.com/hunterhug/gostorage/seaweedfs"
)

func TestLocalOSS_CreateDir(t *testing.T) {
	//MyVOSS = &LocalOSS{Root: "/home/hunterhug"}
	MyVOSS = &seaweedfs.SeaWeedFs{NetworkUrl: "192.168.0.101:38888", Debug: true, TimeOut: 10}
	err := MyVOSS.Init()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	dir := "./dir"
	name := "./dir/a.asv"
	err = MyVOSS.CreateDir(dir)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = MyVOSS.CreateFile(name, []byte("sss"))
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	s, err := MyVOSS.ReadFile(name)
	fmt.Println(s, err)
}
