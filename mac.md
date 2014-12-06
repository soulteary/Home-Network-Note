## Mac 环境配置

### HTTP

	HTTP Server是最简单的server之一了。

系统自带，但是考虑到要和Win一致，不妨使用`XMAPP`，目前最新的`XMAPP`由`Bitnami`打包。

### MySQL

XMAPP自带，使用v`5.6`。

### 简单配置

- `http.conf`引入`vhosts`，简单配置`vhosts`即可。

### 已知BUG

- Mac下重启进程或者休眠后，MySQL会假死，`sudo killall mysqld`即可（如果你有其他的地方也起了mysql，那么ps自己kill吧）
- 升级系统的时候，会丢失文件（xmappfiles被替换），使用软链或者指定其他的`documentRoot`吧。



