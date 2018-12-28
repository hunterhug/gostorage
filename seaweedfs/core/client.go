package core

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hunterhug/marmot/miner"
	"strings"
)

type FilerClient struct {
	Address string
	w       *miner.Worker
}

type FileList struct {
	Path         string     `json:"Path"`         // 文件路径
	Limit        int        `json:"Limit"`        // 限制数量
	LastFileName string     `json:"LastFileName"` // 最后的文件名
	Entries      []FileInfo `json:"Entries"`
	//Replication string `json:"Replication"`
}

type FileInfo struct {
	Id       string                   // 文件ID
	Size     int                      // 文件大小,Byte
	IsDir    bool                     // 是否是目录
	FullPath string `json:"FullPath"` // 文件路口
	Mtime    string `json:"Mtime"`    // 修改时间
	Crtime   string `json:"Crtime"`   // 创建时间

	Chunks []Chunks `json:"chunks"`
}

type Chunks struct {
	FileId string `json:"file_id"`
	Size   int    `json:"size"`
}

type UploadResult struct {
	FileName string `json:"name"`
	Size     int    `json:"size"`
	FileId   string `json:"fid"`
}

func afterAction(c context.Context, m *miner.Worker) {
	m.ClearAll()
}

func NewClient(address string) *FilerClient {
	if strings.HasSuffix(address, "/") {
		address = strings.TrimRight(address, "/")
	}

	if strings.HasPrefix(address, "https://") || strings.HasPrefix(address, "http://") {

	} else {
		address = "http://" + address
	}

	api := miner.NewAPI()
	api.SetAfterAction(afterAction)
	client := new(FilerClient)
	client.Address = address
	client.w = api
	return client
}

func SetDebug() {
	miner.SetLogLevel(miner.DEBUG)
}

func SetTimeOut(times int) {
	miner.DefaultTimeOut = times
}

// 获取目录信息， 如果是文件返回二进制
// err  : HTTP请求正常则nil
// exist: 目录存在为true，否则为false
// isDir: 如果是文件，返回false
// dirInfo  目录则返回目录信息
func (c *FilerClient) GetDirInfo(dir string, lastFileName string, Limit int) (err error, exist bool, isDir bool, dirInfo *FileList, fileData []byte) {
	exist = true
	isDir = true
	if !strings.HasPrefix(dir, "/") {
		dir = "/" + dir
	}

	//if !strings.HasSuffix(dir, "/") {
	//	//	dir = dir + "/"
	//	//}

	path := fmt.Sprintf("%s%s", c.Address, dir)

	temp := c.w.SetUrl(path)
	if lastFileName != "" {
		temp.SetFormParm("lastFileName", lastFileName)
	}

	if Limit <= 0 {
		Limit = 100
	}
	temp.SetFormParm("limit", fmt.Sprintf("%d", Limit))

	temp.SetHeaderParm("Accept", miner.HTTPJSONContentType)
	data, err := temp.Get()
	if c.w.UrlStatuscode == 404 {
		exist = false
		return
	}
	if err != nil {
		return
	}

	file := c.w.Response.Header.Get("Content-Disposition")
	if strings.Contains(file, "inline") {
		isDir = false
		fileData = data
		return
	}

	dirInfo = new(FileList)
	err = json.Unmarshal(data, dirInfo)

	if len(dirInfo.Entries) == 0 {
		exist = false
		return
	}

	for k, v := range dirInfo.Entries {
		if len(v.Chunks) == 0 {
			dirInfo.Entries[k].IsDir = true
			continue
		}
		dirInfo.Entries[k].Id = v.Chunks[0].FileId
		dirInfo.Entries[k].Size = v.Chunks[0].Size
	}
	return
}

// 上传会覆盖之前的操作
func (c *FilerClient) Upload(data []byte, fileName, uploadFilePath string) (b *UploadResult, err error) {
	if !strings.HasPrefix(uploadFilePath, "/") {
		uploadFilePath = "/" + uploadFilePath
	}

	// 必须确保不能覆盖目录
	err, exist, isDir, _, _ := c.GetDirInfo(uploadFilePath, "", 1)
	if err != nil {
		return
	}
	if exist && isDir {
		err = errors.New("this is dir can not replace with a file")
		return
	}

	path := fmt.Sprintf("%s%s", c.Address, uploadFilePath)
	result, err := c.w.SetUrl(path).SetBData(data).SetFileInfo(fileName, "file").PostFILE()
	if err != nil {
		return
	}

	b = new(UploadResult)
	err = json.Unmarshal(result, b)
	return
}

func (c *FilerClient) Download(uploadFilePath string) (b []byte, err error) {
	if !strings.HasPrefix(uploadFilePath, "/") {
		uploadFilePath = "/" + uploadFilePath
	}
	// 必须确保不能覆盖目录
	err, exist, isDir, _, b := c.GetDirInfo(uploadFilePath, "", 1)
	if err != nil {
		return
	}
	if exist && !isDir {
		return
	}

	if !exist {
		err = errors.New("not exist")
		return
	}
	err = errors.New("is a dir")
	return
}

func (c *FilerClient) DELETE(uploadFilePath string) error {
	if !strings.HasPrefix(uploadFilePath, "/") {
		uploadFilePath = "/" + uploadFilePath
	}
	// 必须确保不能覆盖目录
	err, exist, isDir, _, _ := c.GetDirInfo(uploadFilePath, "", 1)
	if err != nil {
		return err
	}
	if exist && !isDir {
		path := fmt.Sprintf("%s%s", c.Address, uploadFilePath)
		_, err = c.w.SetUrl(path).Delete()
		return err
	}

	if !exist {
		return nil
	}

	err = errors.New("is a dir, you must delete file")
	return err
}
