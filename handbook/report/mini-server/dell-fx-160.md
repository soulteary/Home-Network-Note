# 测试迷你服务器性能

## 测试时间

2017.10.14

## 测试环境

测试软件：iperf3 (Docker Ver.)

服务器环境：

```
iperf 3.0.7
Linux 3.10.102 #15101 SMP Fri May 5 12:00:40 CST 2017 x86_64 GNU/Linux
```

测试终端环境：

```
iperf 3.2 (cJSON 1.5.2)
Darwin 16.7.0 Darwin Kernel Version 16.7.0: Thu Jun 15 17:36:27 PDT 2017; root:xnu-3789.70.16~2/RELEASE_X86_64 x86_64
Optional features available: sendfile / zerocopy, authentication
```

## 测试命令


服务器：

```
docker run -it --rm -p 5201:5201 networkstatic/iperf3 -s -f M
```

终端：

```
iperf3 -c <路由IP地址> -f M
```

## 测试结果

路由器：

```
[ ID] Interval           Transfer     Bandwidth
[  5]   0.00-1.00   sec  58.8 MBytes  58.8 MBytes/sec
[  5]   1.00-2.00   sec  43.1 MBytes  43.1 MBytes/sec
[  5]   2.00-3.00   sec  62.5 MBytes  62.5 MBytes/sec
[  5]   3.00-4.00   sec  62.3 MBytes  62.3 MBytes/sec
[  5]   4.00-5.00   sec  64.0 MBytes  64.0 MBytes/sec
[  5]   5.00-6.00   sec  67.8 MBytes  67.8 MBytes/sec
[  5]   6.00-7.00   sec  67.1 MBytes  67.1 MBytes/sec
[  5]   7.00-8.00   sec  67.8 MBytes  67.8 MBytes/sec
[  5]   8.00-9.00   sec  62.7 MBytes  62.7 MBytes/sec
[  5]   9.00-10.00  sec  36.3 MBytes  36.3 MBytes/sec
[  5]  10.00-10.03  sec  2.03 MBytes  65.0 MBytes/sec
- - - - - - - - - - - - - - - - - - - - - - - - -
[ ID] Interval           Transfer     Bandwidth
[  5]   0.00-10.03  sec   595 MBytes  59.3 MBytes/sec                  sender
[  5]   0.00-10.03  sec   594 MBytes  59.2 MBytes/sec                  receiver
```

终端：

```
[ ID] Interval           Transfer     Bitrate
[  5]   0.00-1.00   sec  59.0 MBytes  58.8 MBytes/sec
[  5]   1.00-2.00   sec  43.9 MBytes  44.0 MBytes/sec
[  5]   2.00-3.00   sec  64.1 MBytes  64.3 MBytes/sec
[  5]   3.00-4.00   sec  62.3 MBytes  62.2 MBytes/sec
[  5]   4.00-5.00   sec  62.6 MBytes  62.6 MBytes/sec
[  5]   5.00-6.00   sec  69.3 MBytes  69.3 MBytes/sec
[  5]   6.00-7.00   sec  67.0 MBytes  67.0 MBytes/sec
[  5]   7.00-8.00   sec  68.1 MBytes  68.1 MBytes/sec
[  5]   8.00-9.00   sec  61.6 MBytes  61.6 MBytes/sec
[  5]   9.00-10.00  sec  36.7 MBytes  36.7 MBytes/sec
- - - - - - - - - - - - - - - - - - - - - - - - -
[ ID] Interval           Transfer     Bitrate
[  5]   0.00-10.00  sec   595 MBytes  59.5 MBytes/sec                  sender
[  5]   0.00-10.00  sec   594 MBytes  59.4 MBytes/sec                  receiver
```
