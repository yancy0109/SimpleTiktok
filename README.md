# SimpleTiktok
越想越不队  |  抖音项目  |  第三届字节跳动青训营

## 数据库信息
mysql脚本文件/SQL/douyin.sql用于本地数据库测试

## 运行时环境配置
执行 go mod tidy  
再对.env文件进行修改
MYSQL_USER                  数据库用户名  
MYSQL_PASSWORD              数据库莫玛  
MYSQL_DBNAME                sql文件生成得数据库对应名字(默认douyin)  
MYSQL_HOST                  数据库连接地址  
MYSQL_PORT                  数据库连接地址  
RESOURCE_DIRECTORY          当前项目运行所在的ip地址（填写公网ip/局域网下ip）  

## 视频上传功能需要安装
使用FFMPEG处理本地视频提取首帧作为封面，通过exec.Command()调用命令，需要以下环境
### Windows
#### FFmpeg安装
FFmpeg链接(https://ffmpeg.org/)  
选择windows安装 解压安装包至特定文件夹  
将 安装目录下xxx\bin 加入至环境变量  
ffmpeg -version 测试输出版本信息

#视频流接口实现上传
有逻辑漏洞和代码不规范请多多指教

#### mysql省略
### Linux
#### FFmpeg安装即可
sudo apt install ffmpeg
FFmpeg链接(https://ffmpeg.org/)
#### Mysql
##### 安装
sudo apt install mysql-server mysql-client  
##### 查看初始账号密码
sudo cat /etc/mysql/debian.conf  
##### 以初始账号密码登录
mysql -u___ -p____
##### 修改root账户密码为root
ALTER USER 'root'@'localhost' IDENTIFIED WITH mysql_native_password BY 'root';
flush privileges;

