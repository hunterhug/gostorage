package gostorage

import "os"

// 对象存储鸭子接口
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
