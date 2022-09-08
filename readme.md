

## gomemo

#### [查看中文文档](https://github.com/1911447416/gomemo/blob/master/readme_ch.md)

------

A small exercise based on gin+gorm development, the front-end code is bootstrap take over



### 1.Download github source code

```
git clone https://github.com/1911447416/gomemo
```

### 2.Configure MySQL, create the required database and tables

- If you do not have mysql installed,  [install mysql](https://github.com/1911447416/gomemo/blob/master/mysql.md) first
- Recommended to use **Navicat Premmium** for visual database management

#### Login database

```mysql
mysql -u root -p
```
#### Create databases and tables
##### mysql8
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

##### mysql 5.6
```mysql
CREATE DATABASE gin DEFAULT CHARSET=utf8;
USE gin;
CREATE TABLE `memos` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(255) DEFAULT NULL,
  `status` int(10) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=25 DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;
SET NAMES utf8;
SET FOREIGN_KEY_CHECKS = 0;
BEGIN;
INSERT INTO `memos` VALUES (1, 'test', 1);
COMMIT;
SET FOREIGN_KEY_CHECKS = 1;
```

### 3.Create a folder and configuration file in the main directory

- For Mac OS, **ForkLift** is recommended for file management, and **iTerm** is recommended for writing command line

`cd gomemo`

```
mkdir conf
cd conf
touch app.ini
vim app.ini
```

### 4.Edit the configuration file, press i, fill in the mysql information, and then :wq save to exit

```vim app.ini
port = 8080

[mysql]
ip       = 127.0.0.1
port     = 3306
user     = root
password = passwd
database = gin
```

### 5.Firewall release of port 8080

```
#Add the ports that need to be opened：
firewall-cmd --add-port=8080/tcp --permanent
#Reload the added port：
firewall-cmd --reload
```

### 6.Check whether the port is successfully opened

`firewall-cmd --query-port=8080/tcp`

### 7.Grant permission to run in the background

```
chmod 777 ./main
nohup ./main &
```

After starting, just open `http://Your-server's-public-network-URL:8080/` in your browser

If the firewall blocks it, run ``systemctl stop firewalld`` to turn off the firewall

[![vB4yOs.png](https://s1.ax1x.com/2022/08/17/vB4yOs.png)](https://imgtu.com/i/vB4yOs)
