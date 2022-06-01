# SimpleTiktok
越想越不队  |  抖音项目  |  第三节字节跳动青训营

mysql脚本文件douyin.sql用于本地数据库测试


# 视频上传功能
使用FFMPEG处理本地视频提取首帧作为封面，需要以下环境  
下面是linux安装,Windows也可以安装ffmpeg，但是就很麻烦，需要安装linux命令支持之类的？我使用的Ubuntu  
FFmpeg链接(https://ffmpeg.org/)
## Mysql
### 安装
sudo apt install mysql-server mysql-client  
### 查看初始账号密码
sudo cat /etc/mysql/debian.conf  
### 以初始账号密码登录
mysql -u___ -p____
### 修改root账户密码为root
ALTER USER 'root'@'localhost' IDENTIFIED WITH mysql_native_password BY 'root';
flush privileges;
## FFmpeg安装即可
sudo apt install ffmpeg