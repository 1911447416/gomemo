

# Centos 7 Detailed Steps to Install MySQL

## 1.1 MySQL Installation

### 1.1.1 Download the **wget** command

```
yum -y install wget
```

### 1.1.2 Download mysql installer online

```
wget https://dev.mysql.com/get/mysql57-community-release-el7-8.noarch.rpm
```

### 1.1.3 Installing MySQL

```
rpm -ivh mysql57-community-release-el7-8.noarch.rpm
```

### 1.1.4 Install mysql service

- First go to the `cd /etc/yum.repos.d/` directory.

```
cd /etc/yum.repos.d/
```

- Install MySQL services (this process can be a bit slow)

```
yum -y install mysql-server
```

linux install MySQL with the following error.

[![vqpWDg.md.png](https://s1.ax1x.com/2022/09/08/vqpWDg.md.png)](https://imgse.com/i/vqpWDg)

Cause: MySQL GPG key has expired.

Solution: Execute the following command to solve the problem

```rpm --import https://repo.mysql.com/RPM-GPG-KEY-mysql-2022```

 

### 1.1.5 Starting MySQL

```
systemctl start mysqld
```

## 1.2 Change the MySQL temporary password

MySQL will have a temporary password after successful installation, we can use the `grep` command to check the temporary password to log in to MySQL first, and then change the MySQL password.

### 1.2.1 Obtaining a temporary MySQL password

```
grep 'temporary password' /var/log/mysqld.log
```

[![vqpLKU.md.png](https://s1.ax1x.com/2022/09/08/vqpLKU.md.png)](https://imgse.com/i/vqpLKU)

 

 1.2.2 Use temporary password to log in first

```
mysql -uroot -p
```

### 1.2.3 Change MySQL password checksum strength to low risk

```
set global validate_password_policy=LOW;
```

### 1.2.4 Change the password length of MySQL

```
set global validate_password_length=5;
```

### 1.2.5 Change MySQL Password

```
ALTER USER 'root'@'localhost' IDENTIFIED BY 'admin'; 
```

## 1.3 Allow remote access

### 1.3.1 First, turn off Cenots' firewall

```
sudo systemctl disable firewalld
```

### 1.3.2 Modify MySQL to allow anyone to connect

1）First login to MySQL

```
mysql -uroot -padmin
```

2）Switching to mysql data

```
use mysql;
```

3）View user table

```
select Host,User from user;
```

[![vqpOrF.md.png](https://s1.ax1x.com/2022/09/08/vqpOrF.md.png)](https://imgse.com/i/vqpOrF)

 

 Found that the `root` user only allows the `localhost` host to log in

4）Modify to allow access to any address

```
update user set Host='%' where User='root';
```

5）Refresh Permissions

```
flush privileges;
```

### 1.3.3 Testing with Navicat Connection Tool

[![vqpvVJ.md.png](https://s1.ax1x.com/2022/09/08/vqpvVJ.md.png)](https://imgse.com/i/vqpvVJ)

 

 [![vq998x.md.png](https://s1.ax1x.com/2022/09/08/vq998x.md.png)](https://imgse.com/i/vq998x)

 

