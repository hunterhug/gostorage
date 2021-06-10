# 文件存储库

```
type Interface interface {
	Init() error                                  // 初始化
	CreateDir(name string) error                  // 创建文件夹
	RemoveDir(name string, all bool) error        // 删除文件夹
	CreateFile(name string, data []byte) error    // 创建文件
	ReadFile(name string) ([]byte, error)         // 读取文件
	RemoveFile(name string) error                 // 移除文件
	IsExist(name string) bool                     // 文件是否存在
	IsDir(name string) bool                       // 是否是目录
	IsFile(name string) bool                      // 是否是文件
	GetFileInfo(name string) (os.FileInfo, error) // 获取文件信息
}
```

## 例子

```
package main

import (
	"fmt"
	"github.com/hunterhug/gostorage"
	"github.com/hunterhug/marmot/util"
)

func main() {
	myStorage := gostorage.NewLocalStorage(util.CurDir())
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

```

