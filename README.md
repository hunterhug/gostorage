# 多功能鸭子存储接口

`MyVOSS` 为鸭子类型.

```
type VOSS interface {
	Init() error                                  // 初始化
	CreateDir(name string) error                  // 创建文件夹
	RemoveDir(name string, all bool) error        // 删除文件夹
	CreateFile(name string, data []byte) error    // 创建文件
	ReadFile(name string) ([]byte, error)         // 读取文件
	RemoveFile(name string) error                 // 移除文件
	IsExist(name string) bool                     // 文件是否存在
	IsDir(name string) bool                       // 是否是目录
	IsFile(name string) bool                      // 是否是文件
	GetFileInfo(name string) (interface{}, error) // 获取文件信息
}
```

默认为本地存储,目前支持 `seaweeedfs` :

```
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
```