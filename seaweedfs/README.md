# 部署服务

https://github.com/chrislusf/seaweedfs/wiki

## 制作镜像

```
cd docker
docker build -t seaweed:latest .
docker tag seaweed:latest 192.168.0.101:8888/public/seaweed:latest
docker push  192.168.0.101:8888/public/seaweed:latest
```

启动前准备:

```
mkdir /storage/app/seaweed
mkdir /storage/app/seaweed/master
mkdir /storage/app/seaweed/v1
mkdir /storage/app/seaweed/config

vim /storage/app/seaweed/config/filer.toml

[mysql]
# ./weed.exe scaffold filer
enabled = true
hostname = "192.168.0.101"
port = 3306
username = "root"
password = "123456"
database = "seafile"              # create or use an existing database
connection_max_idle = 2
connection_max_open = 100
```

请建数据库:

```
sudo docker exec -it zendaomysql mysql -uroot -P13306 -p

create database voicefile default character set utf8mb4 collate utf8mb4_unicode_ci;
use voicefile;
CREATE TABLE IF NOT EXISTS filemeta (
dirhash     BIGINT        COMMENT 'first 64 bits of MD5 hash value of directory field',
name        VARCHAR(1000) COMMENT 'directory or file name',
directory   VARCHAR(4096) COMMENT 'full path to parent directory',
meta        BLOB,
PRIMARY KEY (dirhash, name)
) DEFAULT CHARSET=utf8;
```

启动
```
docker-compose up -d
```

测试：

```
# Basic Usage:
> curl -F file=@report.js "http://192.168.0.101:38888/javascript/"
{"name":"report.js","size":866,"fid":"7,0254f1f3fd","url":"http://localhost:8081/7,0254f1f3fd"}
> curl  "http://192.168.0.101:38888/javascript/report.js"   # get the file content
...
> curl -F file=@report.js "http://192.168.0.101:38888/javascript/new_name.js"    # upload the file with a different name
{"name":"report.js","size":866,"fid":"3,034389657e","url":"http://localhost:8081/3,034389657e"}
> curl  -H "Accept: application/json" "http://192.168.0.101:38888/javascript/?pretty=y"            # list all files under /javascript/
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


curl  -H "Accept: application/json" "http://192.168.0.101:38888/javascript/?pretty=y&lastFileName=new_name.js&limit=2"
{
  "Directory": "/javascript/",
  "Files": [
    {
      "name": "report.js",
      "fid": "7,0254f1f3fd"
    }
  ]
}


curl -X DELETE "http://192.168.0.101:38888/assets/report.js"
```

可打开:

http://192.168.0.101:39333/

http://192.168.0.101:38888