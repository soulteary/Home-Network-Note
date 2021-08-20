# Mac 环境配置

## HTTP／PROXY

- 优先使用 Docker NGINX／APACHE
- 次之使用 Node Server

## RDC

- 优先使用 Docker MySQL／Mariadb
- 次之使用虚拟机中的DB Server

## Cache

- 优先使用 Docker Redis
- 次之使用 Brew Redis


## 已知BUG

- Mac下重启进程或者休眠后，如果是系统中安装了MySQL，进行可能会假死，`sudo killall mysqld`即可。
- 升级系统版本的时候，如果使用了XMAPP之类的套件，可能会丢失软链接生成的`xmappfiles`。
