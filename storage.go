package storage

import (
	"errors"
	"io/ioutil"
	"os"
	"path"
)

// 默认存储接口
var (
	MyVOSS     VOSS = new(LocalOSS) // 这个给鸭子类型用
	MyLocalOSS VOSS = new(LocalOSS) // 永远都是本地存储
)

// 对象存储鸭子接口
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

type LocalOSS struct {
	Root string
}

func (l *LocalOSS) Init() error {
	if l.Root == "" {
		return nil
	}
	if path.IsAbs(l.Root) {
		return nil
	}
	return errors.New("must abs path")

}
func (l *LocalOSS) CreateDir(name string) error {
	if l.Root != "" {
		name = path.Join(l.Root, name)
	}
	return os.MkdirAll(name, 0777)
}

func (l *LocalOSS) GetFileInfo(name string) (interface{}, error) {
	if l.Root != "" {
		name = path.Join(l.Root, name)
	}
	info, err := os.Stat(name)
	return info, err
}

func (l *LocalOSS) ReadFile(name string) ([]byte, error) {
	if l.Root != "" {
		name = path.Join(l.Root, name)
	}
	return ioutil.ReadFile(name)
}

func (l *LocalOSS) RemoveDir(name string, all bool) error {
	if l.Root != "" {
		name = path.Join(l.Root, name)
	}
	if all {
		return os.RemoveAll(name)
	}
	return os.Remove(name)
}

func (l *LocalOSS) CreateFile(name string, data []byte) error {
	if l.Root != "" {
		name = path.Join(l.Root, name)
	}
	err := ioutil.WriteFile(name, data, 0777)
	return err
}

func (l *LocalOSS) RemoveFile(name string) error {
	if l.Root != "" {
		name = path.Join(l.Root, name)
	}
	return os.Remove(name)
}

func (l *LocalOSS) IsExist(name string) bool {
	if l.Root != "" {
		name = path.Join(l.Root, name)
	}
	f, err := os.Open(name)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	f.Close()
	return true
}

func (l *LocalOSS) IsDir(name string) bool {
	if l.Root != "" {
		name = path.Join(l.Root, name)
	}
	info, err := os.Stat(name)
	if err != nil {
		return false
	} else {
		if info.IsDir() {
			return true
		} else {
			return false
		}
	}
}

func (l *LocalOSS) IsFile(name string) bool {
	if l.Root != "" {
		name = path.Join(l.Root, name)
	}
	info, err := os.Stat(name)
	if err != nil {
		return false
	} else {
		if info.IsDir() {
			return false
		} else {
			return true
		}
	}
}

func Init() error {
	return MyVOSS.Init()

}
func CreateDir(name string) error {
	return MyVOSS.CreateDir(name)
}

func GetFileInfo(name string) (interface{}, error) {
	return MyVOSS.GetFileInfo(name)
}

func ReadFile(name string) ([]byte, error) {
	return MyVOSS.ReadFile(name)
}

func RemoveDir(name string, all bool) error {
	return MyVOSS.RemoveDir(name, all)
}

func CreateFile(name string, data []byte) error {
	return MyVOSS.CreateFile(name, data)
}

func RemoveFile(name string) error {
	return MyVOSS.RemoveFile(name)
}

func IsExist(name string) bool {
	return MyVOSS.IsExist(name)
}

func IsDir(name string) bool {
	return MyVOSS.IsDir(name)
}

func IsFile(name string) bool {
	return MyVOSS.IsFile(name)
}
