package gostorage

import (
	"errors"
	"github.com/hunterhug/gostorage/seaweedfs/core"
	"github.com/hunterhug/marmot/miner"
	"os"
	"path"
	"strings"
)

type SeaWeedFs struct {
	NetworkUrl string
	TimeOut    int
	Debug      bool
}

func NewSeaWeedFs(url string, timeOut int, debug bool) (Interface, error) {
	s := new(SeaWeedFs)
	s.Debug = debug
	s.NetworkUrl = url
	s.TimeOut = timeOut
	err := s.Init()
	return s, err
}

func (s *SeaWeedFs) Init() error {
	core.SetTimeOut(s.TimeOut)
	if s.Debug {
		core.SetDebug()
	}
	return nil

}
func (*SeaWeedFs) CreateDir(name string) error {
	return nil
}

func (s *SeaWeedFs) GetFileInfo(name string) (os.FileInfo, error) {
	return nil, nil

}

func (s *SeaWeedFs) ReadFile(name string) ([]byte, error) {
	_, filePath, err := fixFilePath(name)
	if err != nil {
		return nil, err
	}

	Fs := core.NewClient(s.NetworkUrl)

	return Fs.Download(filePath)
}

func (*SeaWeedFs) RemoveDir(name string, all bool) error {
	return nil
}

func fixFilePath(filePath string) (string, string, error) {
	if filePath == "" {
		return "", filePath, errors.New("filepath invalid")
	}
	fileName := path.Base(filePath)
	if fileName == "." || fileName == "/" {
		return fileName, filePath, errors.New("filepath invalid")
	}

	if strings.HasPrefix(filePath, ".") {
		filePath = strings.TrimLeft(filePath, ".")
	}

	if !strings.HasPrefix(filePath, "/") {
		filePath = "/" + filePath
	}

	return fileName, filePath, nil
}
func (s *SeaWeedFs) CreateFile(name string, data []byte) error {
	fileName, filePath, err := fixFilePath(name)
	if err != nil {
		return err
	}
	Fs := core.NewClient(s.NetworkUrl)

	if len(data) != 0 {
		info, err := Fs.Upload(data, fileName, filePath)
		if err != nil {
			return err
		}

		miner.Log().Debugf("upload net:%#v,%s", info, filePath)
	}
	return nil
}

func (s *SeaWeedFs) RemoveFile(name string) error {
	_, filePath, err := fixFilePath(name)
	if err != nil {
		return err
	}
	Fs := core.NewClient(s.NetworkUrl)
	return Fs.DELETE(filePath)

}

func (*SeaWeedFs) IsExist(name string) bool {
	return true
}

func (*SeaWeedFs) IsDir(name string) bool {
	return true
}

func (*SeaWeedFs) IsFile(name string) bool {
	return true
}
