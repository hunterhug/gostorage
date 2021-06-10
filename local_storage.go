package gostorage

import (
	"errors"
	"io/ioutil"
	"os"
	"path"
)

type LocalStorage struct {
	Root string
}

func NewLocalStorage(root string) Interface {
	s := new(LocalStorage)
	s.Root = root
	return s
}

func (l *LocalStorage) Init() error {
	if l.Root == "" {
		return nil
	}
	if path.IsAbs(l.Root) {
		return nil
	}
	return errors.New("must abs path")

}
func (l *LocalStorage) CreateDir(name string) error {
	if l.Root != "" {
		name = path.Join(l.Root, name)
	}
	return os.MkdirAll(name, 0777)
}

func (l *LocalStorage) GetFileInfo(name string) (os.FileInfo, error) {
	if l.Root != "" {
		name = path.Join(l.Root, name)
	}
	info, err := os.Stat(name)
	return info, err
}

func (l *LocalStorage) ReadFile(name string) ([]byte, error) {
	if l.Root != "" {
		name = path.Join(l.Root, name)
	}
	return ioutil.ReadFile(name)
}

func (l *LocalStorage) RemoveDir(name string, all bool) error {
	if l.Root != "" {
		name = path.Join(l.Root, name)
	}
	if all {
		return os.RemoveAll(name)
	}
	return os.Remove(name)
}

func (l *LocalStorage) CreateFile(name string, data []byte) error {
	if l.Root != "" {
		name = path.Join(l.Root, name)
	}
	err := ioutil.WriteFile(name, data, 0777)
	return err
}

func (l *LocalStorage) RemoveFile(name string) error {
	if l.Root != "" {
		name = path.Join(l.Root, name)
	}
	return os.Remove(name)
}

func (l *LocalStorage) IsExist(name string) bool {
	if l.Root != "" {
		name = path.Join(l.Root, name)
	}
	f, err := os.Open(name)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	_ = f.Close()
	return true
}

func (l *LocalStorage) IsDir(name string) bool {
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

func (l *LocalStorage) IsFile(name string) bool {
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
