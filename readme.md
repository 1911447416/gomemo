

## gomemo·备忘录

基于gin+gorm开发的小练习，前端代码是bootstrap拿过来的

## 首先下载github上的代码

```
git clone https://github.com/1911447416/gomemo
```

## 配置MySQL

### 1.创建所需要的数据库和表

#### 登录数据库

```mysql
mysql -u root -p
```

```mysql
CREATE DATABASE gin DEFAULT CHARSET=utf8mb4;
USE gin;
CREATE TABLE `memos` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(255) DEFAULT NULL,
  `status` int(10) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=25 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;
BEGIN;
INSERT INTO `memos` VALUES (1, 'test', 1);
COMMIT;
SET FOREIGN_KEY_CHECKS = 1;
```

### 2.在`/gomemo/`目录下创建文件夹及配置文件

```
mkdir conf
cd conf
touch app.ini
vim app.ini
```

### 3.编辑配置文件，按下i键，将mysql的信息填入，然后:wq保存退出

```vim app.ini
port = 8080

[mysql]
ip       = 127.0.0.1
port     = 3306
user     = root
password = passwd
database = gin
```

### 4.授予权限后台运行

```
chmod 777 ./main
nohup ./main &
```

### 5.防火墙放行8080端口

```
#添加指定需要开放的端口：
firewall-cmd --add-port=443/tcp --permanent
#重载入添加的端口：
firewall-cmd --reload
```

### 6.查询端口是否开启成功

`firewall-cmd --query-port=123/tcp`

### 7.浏览器运行`http://你的服务器公网网址:8080/`即可