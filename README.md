DGMR（Docker + Go + MySQL + Redis）



## 目錄結構

```
/
├── services                    服務目錄
│   ├── mysql                   MySQL 文件
│   └── redis                   Redis 文件
├── logs                        日誌
├── docker-compose.sample.yml   Docker 服務配置範例
├── env.smaple                  環境範例
└── app                        程式碼
```

## 使用說明
1. 本地安装
    - `git`
    - `Docker`(系统需为Linux，Windows 10 Build 15063+，或MacOS 10.12+，且必须要`64`位）
    - `docker-compose 1.7.0+`
2. `clone`项目：
    ```
    $ git clone https://github.com/sam80719/DGMR
    ```
3. 如果不是`root`用户，还需将当前用户加入`docker`用户组：
    ```
    $ sudo gpasswd -a ${USER} docker
    ```
4. 拷贝并命名配置文件（Windows系统请用`copy`命令），启动：
    ```
    $ cd DGMR                                           
    $ cp env.sample .env                                
    $ cp docker-compose.sample.yml docker-compose.yml
    ```

###  查看docker network
```sh
ifconfig docker0
```

### 進入docker go
```sh
docker exec -it {$CONTAINER ID} /bin/sh
```

### 網頁phpMyAdmin網址
```shell
http://localhost:8080/
```

## 使用log
Log文件生成的位置依赖于conf下各log配置的值。
### MySQL日誌
MySQL容器中的MySQL使用的是`mysql`用户启动，無法在`/var/log`下的增加日誌文件。
所以，我们把MySQL的日誌放在與data一样的目錄，即项目的`mysql`目录下，對應容器中的`/var/lib/mysql/`目錄。
```bash
slow-query-log-file     = /var/lib/mysql/mysql.slow.log
log-error               = /var/lib/mysql/mysql.error.log
```
以上是mysql.conf中的日誌文件的配置。



## 常见问题
### Docker容器時間
容器时间在.env文件中配置`TZ`變數，所有支持的時區请看[time.wiki](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones)或者[PHP timezone](https://www.php.net/manual/zh/timezones.php)。

### Docker使用cron 
[Docker使用cron](https://www.awaimai.com/2615.html)

### 如何連接MySQL和Redis
兩個狀況，

第一種，在**ˇ代碼中**。
容器與容器是`expose`端口聯絡，而且在同一个`networks`下，所以連接的`host`参数直接用容器名稱，`port`参数就是容器内部的端口。更多请参考[《docker-compose ports和expose的區別》](https://www.awaimai.com/2138.html)。
第二種，**在主機中**通过**命令行**或者**Navicat**等工具連接。主機要連接mysql和redis的话，要求容器必須經過`ports`把端口映射到主機了。以 mysql 為例，`docker-compose.yml`文件中有這樣的`ports`配置：`3306:3306`，就是主機的3306和容器的3306端口形成了映射，所以我们可以這樣連接：
```bash
$ mysql -h127.0.0.1 -uroot -proot -P3306
$ redis-cli -h127.0.0.1
```
這裡`host`參數不能用localhost是因為他默認是通過sock文件與mysql通信，而容器與主機文件系统已經隔離，所以需要通過TCP方式連接，所以需要指定IP。



