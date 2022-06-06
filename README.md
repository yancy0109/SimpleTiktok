# SimpleTiktok
越想越不队  |  抖音项目  |  第三节字节跳动青训营

mysql脚本文件douyin.sql用于本地数据库测试


# 视频上传功能
使用FFMPEG处理本地视频提取首帧作为封面，通过exec.Command()调用命令，需要以下环境
## Windows
### FFmpeg安装
FFmpeg链接(https://ffmpeg.org/)  
选择windows安装 解压安装包至特定文件夹  
将 安装目录下xxx\bin 加入至环境变量  
ffmpeg -version 测试输出版本信息

#视频流接口实现上传
有逻辑漏洞和代码不规范请多多指教


### mysql省略
## Linux
### FFmpeg安装即可
sudo apt install ffmpeg
FFmpeg链接(https://ffmpeg.org/)
### Mysql
#### 安装
sudo apt install mysql-server mysql-client  
#### 查看初始账号密码
sudo cat /etc/mysql/debian.conf  
#### 以初始账号密码登录
mysql -u___ -p____
#### 修改root账户密码为root
ALTER USER 'root'@'localhost' IDENTIFIED WITH mysql_native_password BY 'root';
flush privileges;