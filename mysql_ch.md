

# Centos7 安装MySQL详细步骤

## 1.1 MySQL安装

### 1.1.1 下载wget命令

```
yum -y install wget
```

### 1.1.2 在线下载mysql安装包

```
wget https://dev.mysql.com/get/mysql57-community-release-el7-8.noarch.rpm
```

### 1.1.3 安装MySQL

```
rpm -ivh mysql57-community-release-el7-8.noarch.rpm
```

### 1.1.4 安装mysql服务

- 首先进入`cd /etc/yum.repos.d/`目录。

```
cd /etc/yum.repos.d/
```

- 安装MySQL服务（这个过程可能有点慢）

```
yum -y install mysql-server
```

linux安装MySQL时报错：

[![vqpWDg.md.png](https://s1.ax1x.com/2022/09/08/vqpWDg.md.png)](https://imgse.com/i/vqpWDg)

原因：MySQL GPG 密钥已过期导致

解决办法：执行一下命令，解决

```rpm --import https://repo.mysql.com/RPM-GPG-KEY-mysql-2022```

 

### 1.1.5 启动MySQL

```
systemctl start mysqld
```

## 1.2 修改MySQL临时密码

MySQL安装成功后会有一个临时密码，我们可以使用`grep`命令查看临时密码先登录进去MySQL，然后修改MySQL密码。

### 1.2.1 获取MySQL临时密码

```
grep 'temporary password' /var/log/mysqld.log
```

[![vqpLKU.md.png](https://s1.ax1x.com/2022/09/08/vqpLKU.md.png)](https://imgse.com/i/vqpLKU)

 

 1.2.2 使用临时密码先登录

```
mysql -uroot -p
```

### 1.2.3 把MySQL的密码校验强度改为低风险

```
set global validate_password_policy=LOW;
```

### 1.2.4 修改MySQL的密码长度

```
set global validate_password_length=5;
```

### 1.2.5 修改MySQL密码

```
ALTER USER 'root'@'localhost' IDENTIFIED BY 'admin'; 
```

## 1.3 允许远程访问

### 1.3.1 首先要关闭Cenots的防火墙

```
sudo systemctl disable firewalld
```

### 1.3.2 修改MySQL允许任何人连接

1）首先登录MySQL

```
mysql -uroot -padmin
```

2）切换到mysql数据

```
use mysql;
```

3）查看user表

```
select Host,User from user;
```

[![vqpOrF.md.png](https://s1.ax1x.com/2022/09/08/vqpOrF.md.png)](https://imgse.com/i/vqpOrF)

 

 发现`root`用户只允许`localhost`主机登录登录

4）修改为允许任何地址访问

```
update user set Host='%' where User='root';
```

5）刷新权限

```
flush privileges;
```

### 1.3.3 使用Navicat连接工具测试

[![vqpvVJ.md.png](https://s1.ax1x.com/2022/09/08/vqpvVJ.md.png)](https://imgse.com/i/vqpvVJ)

 

 [![vq998x.md.png](https://s1.ax1x.com/2022/09/08/vq998x.md.png)](https://imgse.com/i/vq998x)

 

