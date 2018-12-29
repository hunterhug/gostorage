# 网络文件服务部署


## 制作镜像

```
sudo ./docker_build.sh
```

## 启动

```
sudo ./install.sh
```

## 测试

上传文件:

```
> curl -F file=@report.js "http://127.0.0.1:38888/javascript/"
{"name":"report.js","size":866,"fid":"7,0254f1f3fd","url":"http://localhost:8081/7,0254f1f3fd"}
```

获取文件内容:

```
> curl  "http://127.0.0.1:38888/javascript/report.js"   # get the file content
```

上传文件使用不一样的名字:

```
> curl -F file=@report.js "http://127.0.0.1:38888/javascript/new_name.js"    # upload the file with a different name
{"name":"report.js","size":866,"fid":"3,034389657e","url":"http://localhost:8081/3,034389657e"}
```

获取文件夹下文件:

```
> curl -H "Accept: application/json" "http://127.0.0.1:38888/javascript/?pretty=y"   # list all files under /javascript/
{
  "Directory": "/javascript/",
  "Files": [
    {
      "name": "new_name.js",
      "fid": "3,034389657e"
    },
    {
      "name": "report.js",
      "fid": "7,0254f1f3fd"
    }
  ],
  "Subdirectories": null
}
```

翻页:

```
> curl  -H "Accept: application/json" "http://127.0.0.1:38888/javascript/?pretty=y&lastFileName=new_name.js&limit=2"
{
  "Directory": "/javascript/",
  "Files": [
    {
      "name": "report.js",
      "fid": "7,0254f1f3fd"
    }
  ]
}
```

删除文件:

```
> curl -X DELETE "http://127.0.0.1:38888/javascript/report.js"
```

网络文件服务查看: http://127.0.0.1:39333

文件Filer界面: http://127.0.0.1:38888