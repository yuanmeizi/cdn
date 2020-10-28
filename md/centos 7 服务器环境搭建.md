># centos 7 服务器环境搭建

![img](https://bkimg.cdn.bcebos.com/pic/2cf5e0fe9925bc31137974de55df8db1cb13704b?x-bce-process=image/watermark,image_d2F0ZXIvYmFpa2U4MA==,g_7,xp_5,yp_5)

# 1.安装Docker 容器

环境查看

```shell
uname -r
3.10.0-1127.13.1.el7.x86_64
```

#系统版本

[root@iZbp1brcdl34et225vwwxkZ ~]# cat /etc/os-release 
NAME="CentOS Linux"
VERSION="7 (Core)"
ID="centos"
ID_LIKE="rhel fedora"
VERSION_ID="7"
PRETTY_NAME="CentOS Linux 7 (Core)"
ANSI_COLOR="0;31"
CPE_NAME="cpe:/o:centos:centos:7"
HOME_URL="https://www.centos.org/"
BUG_REPORT_URL="https://bugs.centos.org/"

CENTOS_MANTISBT_PROJECT="CentOS-7"
CENTOS_MANTISBT_PROJECT_VERSION="7"
REDHAT_SUPPORT_PRODUCT="centos"
REDHAT_SUPPORT_PRODUCT_VERSION="7"

卸载旧版本

```shell
yum remove docker \
                  docker-client \
                  docker-client-latest \
                  docker-common \
                  docker-latest \
                  docker-latest-logrotate \
                  docker-logrotate \
                  docker-engine
```

安装需要安装包

```shell
 yum install -y yum-utils

#国外镜像地址
yum-config-manager \
    --add-repo \
    https://download.docker.com/linux/centos/docker-ce.repo
    
    #阿里云的镜像地址
yum-config-manager \
    --add-repo \
http://mirrors.aliyun.com/docker-ce/linux/centos/docker-ce.repo
```

更亲软件包索引

```
yum makecache fast
```

安装docker 相关

```shell
yum install docker-ce docker-ce-cli containerd.io
```

镜像加速器

配置阿里云的 镜像加速器

针对Docker客户端版本大于 1.10.0 的用户

您可以通过修改daemon配置文件/etc/docker/daemon.json来使用加速器

```shell
sudo mkdir -p /etc/docker
sudo tee /etc/docker/daemon.json <<-'EOF'
{
  "registry-mirrors": ["https://gwjmb4x7.mirror.aliyuncs.com"]
}
EOF
sudo systemctl daemon-reload
sudo systemctl restart docker
```



# 2.Docker搭建MySQL8镜像

![img](https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1598591773398&di=cd59a2220356735341319a638cd70569&imgtype=0&src=http%3A%2F%2Fdingyue.nosdn.127.net%2Fp%3D5RUpVvb55sjwopuRJLfA05pPbkS6DVGwL1M3yZ1ASe71536287396404.jpg)

1.安装Docker

环境:`14.04.5 LTS`

由于docker安装可以直接看官方文档，部署简单。不赘述。Docker安装成功后，查看版本:

```
root@ubuntudoc:~# docker --version
Docker version 18.06.3-ce, build d7080c1
```

2.查看docker镜像

```shell
root@ubuntudoc:~# docker search mysql
NAME                              DESCRIPTION                                     STARS               OFFICIAL            AUTOMATED
mysql                             MySQL is a widely used, open-source relation…   9238                [OK]                
mariadb                           MariaDB is a community-developed fork of MyS…   3294                [OK]                
mysql/mysql-server                Optimized MySQL Server Docker images. Create…   681                                     [OK]
```

3.拉取mysql官方镜像

```shell
docker pull mysql  
```

4. 查看目前的镜像

```shell
root@ubuntudoc:~# docker images
REPOSITORY                TAG                 IMAGE ID            CREATED             SIZE
nginx                     latest              6678c7c2e56c        10 days ago         127MB
mysql                     latest              9b51d9275906        10 days ago         547MB
dongweiming/web_develop   dev                 43fb02d9c1a3        3 years ago         294MB
```

5.运行docker mysql镜像

```shell
docker run -p 3306:3306 --name zsdmysql -e MYSQL_ROOT_PASSWORD=123456 -d mysql
```

上述命令的参数，有如下含义:

1. `--name`指定了你要取的名字。
2. `-p`对应，需要映射出来的端口。比如:3306:3306,意识表示为zsdmysql的容器里面的3306端口对应我外面这个虚拟机的3306端口。
3. `-e`是mysql的命令，设置root的密码为123456
4. `-d`是运行的镜像，这里是`mysql` 容器镜像

6.查看目前运行的容器

```shell
root@ubuntudoc:~# docker container ls
CONTAINER ID        IMAGE               COMMAND                  CREATED             STATUS              PORTS                               NAMES
3958ab15ea05        mysql               "docker-entrypoint.s…"   8 seconds ago       Up 6 seconds        0.0.0.0:3306->3306/tcp, 33060/tcp   zsdmysql
```

7.进入MySQL

```shell
root@ubuntudoc:~# docker exec -it zsdmysql bash
### 下述代表容器里面的情况了
root@3958ab15ea05:/# df -h
Filesystem      Size  Used Avail Use% Mounted on
overlay          53G  3.7G   47G   8% /
tmpfs            64M     0   64M   0% /dev
tmpfs          1000M     0 1000M   0% /sys/fs/cgroup
/dev/dm-0        53G  3.7G   47G   8% /etc/hosts
shm              64M     0   64M   0% /dev/shm
tmpfs          1000M     0 1000M   0% /proc/acpi
tmpfs          1000M     0 1000M   0% /proc/scsi
tmpfs          1000M     0 1000M   0% /sys/firmware

root@3958ab15ea05:/# mysql -uroot -p123456
mysql: [Warning] Using a password on the command line interface can be insecure.
Welcome to the MySQL monitor.  Commands end with ; or \g.
Your MySQL connection id is 8
Server version: 8.0.19 MySQL Community Server - GPL

Copyright (c) 2000, 2020, Oracle and/or its affiliates. All rights reserved.

Oracle is a registered trademark of Oracle Corporation and/or its
affiliates. Other names may be trademarks of their respective
owners.

Type 'help;' or '\h' for help. Type '\c' to clear the current input statement.

mysql> 
```

---------------------------------------------------分割线--------------------------------------------------------

8. 建立目录映射的MySQL容器

```shell
docker run -p 3306:3306 --name mysql -v /opt/mysql/conf:/etc/mysql -v /opt/mysql/logs:/var/log/mysql -v /opt/mysql/data:/var/lib/mysql -v /opt/mysql/mysql-files:/var/lib/mysql-files -e MYSQL_ROOT_PASSWORD=123456 -d mysql
```

#  3.Docker搭建redis镜像 

![img](https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1598591839071&di=e90c62f4ed3da86f3a950f591589fc2e&imgtype=0&src=http%3A%2F%2F5b0988e595225.cdn.sohucs.com%2Fimages%2F20180327%2F34adc98d775145f0b23c5fa67217af1d.png)

```shell
 docker pull redis:latest
 
 docker images
 
 docker run -itd --name redis -p 6379:6379 redis
 #运行redis 带密码
#docker run -d --name myredis -p 6379:6379 redis --requirepass "mypassword"
 
```

# 4.Docker搭建nginx镜像

![img](https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1598591957982&di=ec6b707488ab19bae703c8362f9c21e3&imgtype=0&src=http%3A%2F%2Fimg2.imgtn.bdimg.com%2Fit%2Fu%3D292861757%2C2192443207%26fm%3D214%26gp%3D0.jpg)

```shell
 docker pull nginx:latest

 docker run -di --name nginx -p 8088:80 -v /root/nginx/dist:/usr/share/nginx/html nginx

 docker exec -it nginx bash
```

# 5.Docker搭建 gitlab



### 

![img](https://img2020.cnblogs.com/blog/1570658/202005/1570658-20200505132256422-1140052201.png)

 

 

\# 1、查找GitLab镜像

```shell
docker search gitlab
```

\# 2、拉取gitlab docker镜像 

```shell
#拉取英文版社区版镜像
docker pull gitlab``/gitlab-ce``:latest
#拉取英文版社区版镜像
docker pull twang2218/gitlab-ce-zh:latest
```

　　

\# 3、运行GitLab并运行容器 

```shell
docker run \
 -itd  \
 -p 8082:80 \
 -p 8022:22 \
 -v /usr/local/gitlab/etc:/etc/gitlab  \
 -v /usr/local/gitlab/log:/var/log/gitlab \
 -v /usr/local/gitlab/opt:/var/opt/gitlab \
 --restart always \
 --privileged=true \
 --name gitlab \
 twang2218/gitlab-ce-zh
```

　　

命令解释：
-i 以交互模式运行容器，通常与 -t 同时使用命令解释：

-t 为容器重新分配一个伪输入终端，通常与 -i 同时使用

-d 后台运行容器，并返回容器ID

-p 9980:80 将容器内80端口映射至宿主机9980端口，这是访问gitlab的端口

-p 9922:22 将容器内22端口映射至宿主机9922端口，这是访问ssh的端口

-v /usr/local/gitlab-test/etc:/etc/gitlab 将容器/etc/gitlab目录挂载到宿主机/usr/local/gitlab-test/etc目录下，若宿主机内此目录不存在将会自动创建，其他两个挂载同这个一样

--restart always 容器自启动

--privileged=true 让容器获取宿主机root权限

--name gitlab-test 设置容器名称为gitlab-test

gitlab/gitlab-ce 镜像的名称，这里也可以写镜像ID

 

\# 重点：接下来的配置请在容器内进行修改，不要在挂载到宿主机的文件上进行修改。否则可能出现配置更新不到容器内，或者是不能即时更新到容器内，导致gitlab启动成功，但是无法访问

\# 4、进入容器内

```shell
docker ``exec` `-it gitlab-``test` `/bin/bash
```


\# 5、修改gitlab.rb （先查看下一个步骤再决定是否进行本步骤，本步骤是可以跳过的） 

```shell
# 打开文件``vi` `/etc/gitlab/gitlab``.rb` `# 这个文件是全注释掉了的，所以直接在首行添加如下配置` ` ` `# gitlab访问地址，可以写域名。如果端口不写的话默认为80端口``eaxternal_url ``'http://192.168.52.128:9980'``# ssh主机ip``gitlab_rails[``'gitlab_ssh_host'``] = ``'192.168.52.128'``# ssh连接端口``gitlab_rails[``'gitlab_shell_ssh_port'``] = 9922
```

 ![img](https://img2020.cnblogs.com/blog/1570658/202005/1570658-20200505132817626-825668641.png)

 

 

\# 6、修改gitlab.yml (这一步原本不是必须的，因为gitlab.rb内配置会覆盖这个，为了防止没有成功覆盖所以我在这里进行配置，当然你也可以选择不修改gitlab.rb直接修改这里)

```shell
# 打开文件``vi` `/opt/gitlab/embedded/service/gitlab-rails/config/gitlab``.yml` ` ` `# 配置一：找到gitlab标签，将其子标签如下修改
```

![img](https://img2020.cnblogs.com/blog/1570658/202005/1570658-20200505132904586-1802669239.png)

```shell
# 配置解释：` `# host：访问的IP` `# port：访问的端口` `# 以上两个和gitlab.rb内eaxternal_url的配置保持一致` `# ssh_host：ssh主机ip，和gitlab.rb内gitlab_rails['gitlab_ssh_host']保持一致` ` `  `# 配置二：找到gitlab_shell标签下的ssh_port，将其修改为9922``#（和gitlab.rb内gitlab_rails['gitlab_shell_ssh_port'] 保持一致）
```

![img](https://img2020.cnblogs.com/blog/1570658/202005/1570658-20200505132916555-1413565128.png)

```shell
# 保存并退出``:wq
```

\# 7、让修改后的配置生效

```shell
gitlab-ctl reconfigure
```

\# 8、重启gitlab 

```shell
gitlab-ctl restart
```

　　

\# 9、退出容器 

```shell
exit
```

\# 10、在游览器输入如下地址，访问gitlab（eaxternal_url配置的就是这个） 

```shell
http:``//192``.168.52.128:8082
```

如果访问不成功的话： 

（1）   进入容器查看gitlab.rb和gitlab.yml文件是否配置成功

（2）   查看防火墙是否开放8082,8822端口

 

\# 11、第一次访问默认是root账户，会需要修改密码（密码至少8位数，出现如下界面就基本上部署成功了）

 ![img](https://img2020.cnblogs.com/blog/1570658/202005/1570658-20200505133043567-300965168.png)

 

 

 

 

\# 12、输入新密码后进行登录

```shell
root``/gitlab123456
```

![img](https://img2020.cnblogs.com/blog/1570658/202005/1570658-20200505133128818-783933212.png)

 

 

\# 13、登录后界面如下图，创建一个新项目

![img](https://img2020.cnblogs.com/blog/1570658/202005/1570658-20200505133148648-594747237.png)

 

 


\# 14、输入信息，创建项目

![img](https://img2020.cnblogs.com/blog/1570658/202005/1570658-20200505133211758-1012465388.png)

 

\# 15、检查这里两处的ip的端口是否是你配置的

![img](https://img2020.cnblogs.com/blog/1570658/202005/1570658-20200505133243558-1204562365.png)

 

 

如果这里两处的端口和ip和配置的不一样，再次进入容器检查gitlab.yml文件

如下图显示：

 ![img](https://img2020.cnblogs.com/blog/1570658/202005/1570658-20200505133300359-309082101.png)

 

 ![img](https://img2020.cnblogs.com/blog/1570658/202005/1570658-20200505133307880-1155931546.png)

 

若不一样，将其修改为你的配置，然后:wq，再进行下面的操作

```shell
gitlab-ctl reconfigure` `gitlab-ctl restart
```



# 6.docker搭建PostgreSQL镜像

![img](https://ss0.bdstatic.com/70cFvHSh_Q1YnxGkpoWK1HF6hhy/it/u=2657638547,1216478257&fm=26&gp=0.jpg)

安装docker步骤略过

1、拉取postgresql镜像

```shell
docker search postgresql

docker pull postgres
```

2、创建本地卷，数据卷可以在容器之间共享和重用， 默认会一直存在，即使容器被删除（`docker volume inspect `pgdata可查看数据卷的本地位置）

```
docker volume create pgdata
```

3、启动容器

```shell
docker run --name postgres2 -e POSTGRES_PASSWORD=password -p 5432:5432 -v pgdata:/var/lib/postgresql/data -d postgres
```

![img](https://img2018.cnblogs.com/i-beta/1660349/201911/1660349-20191114230953551-384672462.png)

 

 4、进入postgres容器执行sql

```shell
docker exec -it postgres2 bash

psql -h localhost -p 5432 -U postgres --password
```

![img](https://img2018.cnblogs.com/i-beta/1660349/201911/1660349-20191114231316599-1839621880.png)

 

至此，postgresql安装成功。

**拉取postgresql可视化工具pgadmin4**

拉取postgresql可视化工具pgadmin4：

```
docker pull dpage/pgadmin4
```

**运行pgadmin4：**

```
docker run -d -p 5433:80 --name pgadmin4 -e PGADMIN_DEFAULT_EMAIL=test@123.com -e PGADMIN_DEFAULT_PASSWORD=123456 dpage/pgadmin4
```

如图：

![img](https://img2020.cnblogs.com/blog/866435/202007/866435-20200708093934272-294499201.png)

 **3、查看服务**

![img](https://img2020.cnblogs.com/blog/866435/202007/866435-20200708094300993-2063681145.png)



**4、打开浏览器访问pgadmin4：[http://ip:5433/](http://localhost:5433/) 如图**

![img](https://img2020.cnblogs.com/blog/866435/202007/866435-20200708094344581-140450036.png)

 输入我们设置的邮箱test@123.com和密码123456，点击Login

![img](https://img2020.cnblogs.com/blog/866435/202007/866435-20200708094442100-795864585.png)

 连接server：

![img](https://img2020.cnblogs.com/blog/866435/202007/866435-20200708094522472-471431013.png)

 打开

![img](https://img2020.cnblogs.com/blog/866435/202007/866435-20200708094754519-601542865.png)

 链接配置

![img](https://img2020.cnblogs.com/blog/866435/202007/866435-20200708095620222-1731249262.png)

 点击

![img](https://img2020.cnblogs.com/blog/866435/202007/866435-20200708095903166-661734408.png)

 提示无法解析,其实宿主机的hosts文件里host.docker.internal对应的还是容器IP 

那你用postgres那个容器的IP去连 

查看容器ip

```
docker exec -it 836  bash  //进入容器 836为这个容器的id 
cat /etc/hosts    //查看容器的ip
```

如图

![img](https://img2020.cnblogs.com/blog/866435/202007/866435-20200708100259634-1439945846.png)

 更改链接

![img](https://img2020.cnblogs.com/blog/866435/202007/866435-20200708100434567-979375562.png)

 结果

![img](https://img2020.cnblogs.com/blog/866435/202007/866435-20200708100516879-636487470.png)



# 7.Docker搭建部署禅道



**简介**：本文主要介绍如何使用Docker方式部署、升级禅道系统。

一、环境准备

运行环境需成功部署Docker服务，推荐使用Docker 18版本以上，对主机环境没有要求。

可通过命令查看Docker版本。

```shell
docker -v
```

![img](https://cdn.chanzhi.org/web/data/upload/202005/f_9a7b01d95138a2c54f74293b3758a6f1.webp?v=1065)

二、下载禅道镜像

目前支持在线下载和离线导入两种部署禅道镜像的方式，可根据自己环境进行选择。

**1、在线下载**

禅道镜像已放于Docker Hub上，地址为： https://hub.docker.com/r/easysoft/zentao/tags

```shell
docker pull easysoft/zentao:latest
```

![img](https://cdn.chanzhi.org/web/data/upload/202007/f_02d34af2aeef30714af3d8d91a01c560.webp?v=1065)

禅道版本和镜像tag对应关系如下：



| 禅道版本 | 镜像标签                |
| -------- | ----------------------- |
| 开源版   | 以数字开头，如：12.3.3  |
| 专业版   | 以pro开头，如：pro8.8.3 |
| 企业版   | 以biz开头，如：biz3.7.2 |

**1、创建docker网络驱动**

```shell
sudo docker network create --subnet=[ip范围] [网络驱动名]
```

其中，ip范围：例如172.172.172.0/24的意思ip可以指定范围为172.172.172.1到172.172.172.254；
网络驱动名：创建的网络驱动名，可随意指定；

```shell
sudo docker network create --subnet=172.172.172.0/24 zentaonet
```

 **2、启动禅道容器**

命令格式如下：

```shell
sudo docker run --name [容器名] -p [主机端口]:80 --network=[网络驱动名] --ip [容器IP] --mac-address [mac地址] -v [主机禅道目录]:/www/zentaopms -v [主机mysql目录]:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=[数据库密码] -d easysoft/zentao:[镜像标签]
```

其中，容器名：启动的容器名字，可随意指定；

主机端口：主机端口为web访问端口；

网络驱动名：刚刚创建的网络驱动名；

容器IP：在网络驱动范围内选择一个作为该容器的固定ip；

mac地址：指定固定的mac地址，建议范围为02:42:ac:11:00:00 到 02:42:ac:11:ff:ff；

主机禅道目录：必须指定，方便禅道代码、附件等数据的持久化，非升级情况需指定空目录；

主机mysql目录：必须指定，方便禅道数据持久化，非升级情况需指定空目录；

数据库密码： 容器内置mysql用户名为root,默认密码123456，如果不修改可以不指定该变量，如果想更改密码可以设置 MYSQL_ROOT_PASSWORD变量来更改密码；

镜像标签：禅道版本。

例如：在主机上创建空目录/www/zentaopms和/www/mysqldata，执行如下命令

```js
sudo docker run --name zentao -p 8000:80 --network=zentaonet --ip 172.172.172.172 --mac-address 02:42:ac:11:00:00 -v /www/zentaopms:/www/zentaopms -v /www/mysqldata:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=123456 -d easysoft/zentao
```

**3、查看容器是否启动成功**

执行如下命令查看容器是否启动成功，如果没有则启动失败，去掉-d选项进行前台运行调试容器，如有任何问题请咨询禅道商务同事。

```js
docker ps
```

![img](https://cdn.chanzhi.org/web/data/upload/202007/f_a8a420dcc9ec65b35ba57210b9352053.webp?v=1065)

四、安装禅道

浏览器直接访问 http://容器ip:宿主机映射端口

![img](https://cdn.chanzhi.org/web/data/upload/202007/f_c776c982daa1ce9bc09ad4bf43bc37af.webp?v=1065)

![img](https://cdn.chanzhi.org/web/data/upload/202005/f_1bb09045e4067734761f1a52e77ba95a.webp?v=1065)

![img](https://cdn.chanzhi.org/web/data/upload/202007/f_6013de793097e4eeaf77987c21386596.webp?v=1065)

![img](https://cdn.chanzhi.org/web/data/upload/202005/f_810dc7ceedde0fc6fcf3b65065aa5988.webp?v=1065)

![img](https://cdn.chanzhi.org/web/data/upload/202005/f_0fa6a7cd54129811e5e141eef0e2ea8e.webp?v=1065)





![img](https://cdn.chanzhi.org/web/data/upload/202005/f_225035ced06b017c6a49710a16a3714e.webp?v=1065) ![img](https://cdn.chanzhi.org/web/data/upload/202007/f_07f7b5a284c7adff5c4c6946e0170cd2.webp?v=1065)

**五、升级禅道**

1.先停止禅道容器，为备份数据做准备

命令格式如下：

```bsh
docker stop [容器名]
```

例如：

```bsh
docker stop zentao 
```

2 **.备份禅道数据**

为安全起见，将上文所述的[主机禅道目录]和[主机mysql目录]进行备份，例如将/www/zentaopms和/www/mysqldata进行 拷贝至主机其他目录。

3.获取最新版源码包，开源版从禅道官网获取，其他版本可咨询禅道商务获取

4.解压缩后，覆盖 [主机禅道目录]，例如解压缩后覆盖主机/www/zentaopms目录

5.启禅道容器

```bsh
docker start [容器名]
```

例如：

```bsh
docker start zentao 
```

6.访问upgrade.php，选择对应的版本，按照提示进行即可

开源版、专业版、企业版的升级相似，具体可以参考开源版禅道的升级 http://www.zentao.net/help-read-78960.html。

**六、安装svn、git客户端  (如果不使用svn、git集成的话，不用安装)**

\1. 进入容器



```bsh
docker exec -it [容器名] /bin/bash
```

  \2. 执行命令



```bsh
apt-get install git -y
apt-get install subversion -y
```

**七、FAQ**

 A：如果修改容器里的mysql等配置文件

 Q：使用如下命令



```bsh
docker exec -it [容器名] /bin/bash
```

比如进入禅道容器，修改my.cnf后保存，重启mysql服务



```bsh
docker exec -it zentao /bin/bash
vi /etc/mysql/my.cnf 
```

同理，修改apache、php.ini的配置也类似，以下列出容器内相关组件的路径。

容器内apache配置文件目录：/etc/apache2/

容器内禅道目录：/www/zentaopms

容器内mysql配置文件目录： /etc/mysql/

容器内php配置文件目录：/etc/php/7.0/apache2



A：为什么容器挂载出来的/www/zentaopms目录是空文件夹
Q：在禅道12.3.1、12.2.stable、biz3.7、pro8.8这四个镜像中，禅道的挂载目录为/app/zentaopms，需修改挂载目录



例如：在主机上创建空目录/app/zentaopms和/app/mysqldata，执行如下命令

```shell
sudo docker run --name zentao -p 80:80 --network=zentaonet --ip 172.172.172.172 --mac-address 02:42:ac:11:00:00 -v /app/zentaopms:/app/zentaopms -v /app/mysqldata:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=123456 -d easysoft/zentao:12.3.3
```