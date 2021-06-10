package gostorage

import (
	"fmt"
	"github.com/hunterhug/marmot/util"
	"testing"
)

func TestLocalOSS_CreateDir(t *testing.T) {
	myStorage := NewLocalStorage(util.CurDir())
	//myStorage, err := NewSeaWeedFs("192.168.0.101:38888", 10, true)
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}

	dir := "./data"
	name := "./data/a.asv"
	err := myStorage.CreateDir(dir)
	if err != nil {
		fmt.Println("create dir err:", err.Error())
		return
	}

	err = myStorage.CreateFile(name, []byte("a.asv text data"))
	if err != nil {
		fmt.Println("create file err:", err.Error())
		return
	}

	raw, err := myStorage.ReadFile(name)
	if err != nil {
		fmt.Println("read file err:", err.Error())
		return
	}

	fmt.Println(string(raw))

	result, err := myStorage.GetFileInfo(name)
	if err != nil {
		fmt.Println("get file info err:", err.Error())
		return
	}

	fmt.Println(result.IsDir(), myStorage.IsFile(name), result.ModTime())

	err = myStorage.RemoveDir(dir, false)
	if err != nil {
		fmt.Println("remove dir err:", err.Error())
	}

	err = myStorage.RemoveDir(dir, true)
	if err != nil {
		fmt.Println("remove dir err:", err.Error())
		return
	}
}
