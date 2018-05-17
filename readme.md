🚧 **Under Construction** 🚧

不定期慢慢更新。LAST UPDATE: 2018.05.17

## 项目说明

    记录搭建家用兼顾学习和娱乐网络环境的一些事情，以及折腾过的一些硬件的小经验。


## 主要需求

| 功能 |备份数据|数据交换|无线接入|数据同步|workflow|开发学习|游戏娱乐|
| --- | --- | --- | --- | --- | --- | --- | --- |
| **核心** | 安全 | 高速 | 安全 | 无感知 | 易定制 | 流畅 | 流畅 |
| **重要** | 高效 | 易用 | 简单 | 准确| 省心 | 省心 | 舒适 |
| **可选** | 易用 | 安全 | 快速 | 全平台 | - | 冗余保障 | - |
| **参考** | - | [高速数据交换](./high-speed-data-exchange.md) | [安全的网络接入](./secure-network-access.md) | [免维护的数据同步](./data-exchange.md) | - | - | - | - |

### 额外需求

- 方案或者配置需要尽可能简化，半年、一年以上后如果需要调整，可以轻松升级、维护、修理设备。
- 网络中存在20台以上设备在线时，网络依旧需要通畅。


### 具体场景

- [1] 两个人的手机、平板、电脑 能够自动同步、备份照片等数据。
- [2] 电脑数据每日自动备份，自己玩的代码进行特殊备份。
- [3] 在线视频、在线游戏、实验性质的数据抓取可以在稳定的网络下进行。
- [4] 支持由于平时懒各种设备不断电，以及两个人都回家后10～20台设备在线时，网络通畅。
- [5] 设备处于相对安全的网络环境中。
- [6] 支持简单易用的workflow。


## 历史设备列表

    如果你考虑入手一些设备(主机/路由/网卡/显示器/储存/移动设备/娱乐/...)，或许可以从这里得到一些参考信息。

[待补全的设备清单](./device-list-history.md) | [显示器相关](./display-history.md)


## 当前屋内常保持联网设备清单及方案

![网络结构](./assets/img/network-2017.10.png)

- [家用10～20台在线设备可参考网络方案](./network.md)
- [无线网络的方案及性能数据](./wifi.md)
- [有线网络的方案及性能数据](./lan.md)


### 🌈 宽带资源

> 不敢想假如家里没有稳定的网络会怎样

| 资源类型 | 明细 | 备注 |
| --- | --- | --- |
| 北京联通 | 500M | 下行跑满，上行实际20M |

- [1] 简化不必要的多线宽带，避免策略路由带来的各种问题，以及需要使用软路由聚合不同类型宽带。
- [2] 不依赖任何提速软件，避免额外的维护成本，当软件不可用时，带宽质量严重受损。


### ⭐️ 路由网关

> 影响网络质量的核心设备，负责部分网络安全事务

| 资源类型 | 明细 | 网络 | 储存 | 开始服务 |
| --- | --- | --- | --- | --- |
| 交换机 | NETGEAR GS116E ^1 | 千兆LAN x16 | - | 2017 |
| 路由器 | newifi D2 ^2 | 2G WIFI / 5G WIFI / 千兆LAN | 8G | 2016 |
| 路由器 | newifi D2 ^3 | 2G WIFI / 5G WIFI / 千兆LAN | 8G | 2017 |
| 路由器 | Xiaomi Mini 第一版 ^4 | 2G WIFI / 5G WIFI / 百兆LAN | - | 2015 / 2016 |
| 路由器 | Xiaomi Mini 青春版 ^5 | 2G WIFI / 百兆LAN | - | 2016 |

- [1] 八口千兆交换机，带管理界面，自带铁壳散热，目前感觉最超值的一个设备。
- [2] 功耗相比较它的性能来说尚可，大内存业界良心，作为二级路由，可以做到半年以上不重启，三方适配的固件功能十分强大。
- [3] 趁产品退市之前购入，作为主路由使用，辅助光猫进行拨号操作。
- [4] 硬伤是全部百兆LAN口，不适合200M+宽带入户的情况。刷机之后十分稳定，家用设备多可做主路由来拨号使用，家用设备少，不存在什么问题，已入三台，偶尔开启，作为开发调试路由使用。
- [5] 功耗极低，小巧方便，适合旅游或者临时需要网络进行调试的场景，三方适配的固件功能强大，已入三台，如果公司不限制使用自建路由作为调试环境，强烈建议入一台。
- [6] ROG使用了一阵，发现性能过剩，定制化不足，还是老老实实用linux系统吧，或者后面能够进一步梳理好一些服务后再说...
- [7] 设备越来越多后，功能独立一些的软路由会更加适用。


### 💻 主机资源

> 提供运算能力的设备

| 资源类型 | 明细 | 网络 | 储存 | 开始服务 |
| --- | --- | --- | --- | --- |
| 资源机器 | DELL FX 160 x1 ^1 | 千兆LAN / 2G WIFI | 4GBRAM / mixed  | 2017 |
| 编码机器 | MacBook Pro ^2 | 千兆LAN & 5G WIFI | 16GBRAM / 1T  | 2017 |
| 公司机器 | MacBook Pro ^3 | 千兆LAN / 5G WIFI | 16GBRAM / 250GB  | 2017 |
| 投影机器 | MacBook Pro ^4 | 千兆LAN & 5G WIFI | 16GBRAM / 500GB  | 2014 |
| 资源机器 | HASEE Z7    ^5 | 千兆LAN & 5G WIFI | 32GBRAM / mixed  | 2015 |
| 功能机器 | N3700组装机  ^6 | 千兆LANx4 & 5G WIFI | 16GBRAM / mixed | 2016 |
| 代码仓库 | N3520 组装机 (N3520 2.166GHz) ^7 | 千兆LAN | 4GBRAM / mixed | 2016 |

- [1] N330，1.6GHz，跑负荷不重的应用，最基础的转码，其实足够了，[网络性能测试](./report/mini-server/dell-fx-160.md)。
- [2] 2017 15-inch, i7 2.9GHz 系统越来越差，键盘特别垃圾，不要买不要买不要买！！！
- [3] 2017 13-inch, i5 2.5GHz 系统越来越差，键盘特别垃圾，不要买不要买不要买！！！
- [4] Mid 2014, i7 2.5GHz 已经碾压了三家公司(淘宝／美团／阿里云)配的笔记本了, Orz。
- [5] i7 2.6GHz, 配合`hotkey`软件调节节能模式，性能表现极好。
- [6] N3700 1.6GHz
- [7] N3520 [网络性能测试](./report/mini-server/n3520.md)。
- MBP 从13款用到17款，感觉越来越垃圾（12寸MB除外），或许未来考虑买个好的Windows Laptop，走黑苹果+Win双系统路线了。

### 🚚 储存资源

> 用来持久化保存资料（开始服务从作为存储角色开始计算）

| 资源类型 | 明细 | 网络 | 储存 | 开始服务 |
| --- | --- | --- | --- | --- |
| 资源备份 | Synology DS 718+ ^1 | 千兆LAN | 4T HDD / 1T HDD | 2017 |
| 资源冷备 | 硬盘x4 ^2 | - | 1T HDD | 2016 |
| 主要备份 | WD MY CLOUD 4T | 千兆LAN | 4T HDD | 2015 |
| 辅助备份 | WD MY CLOUD 3T | 千兆LAN | 3T HDD | 2014 |
| 计划撤换 | Synology DS 115j ^3 | 千兆LAN | 4T HDD | 2017 |
| 长期备份 | Canon G3800 ^4 | 2G WIFI | -      | 2016 |
| 清理备份 | Deli 9920 碎纸机 ^5 | - | -      | 2017 |
| 电力保障 | APC BR550G ^6 | - | - | 2017 |
| 临时储存 | HASEE Z7 | 千兆LAN | 1T SSD | 2017 |
| 临时储存 | DSM 兼容机 | 千兆LAN | 300+ SSD / 1T HDD | 2017 |
| 临时储存 | N3520 组装机    | 千兆LAN | 300G HDD | 2016 |
| ~~代码备份~~ | ~~Ubuntu平板~~ | ~~2G WIFI~~ | ~~128 SSD~~ | ~~2016~~ |

- [1] DS718+ 服役快一年了，使用下来十分稳定，后续储存空间完全不足的时候，考虑换一个四盘位的使用，附：[DS718+ 网络性能测试](./report/mini-server/ds718.md)。
- [2] 每个季度使用资源冷备的盘进行一次全量的重新备份，避免增量备份出现问题，恢复时间过长或者不能完整恢复。
- [3] 原计划当DS718+备胎，然而发现好像没有什么使用价值，吃灰很久，计划撤掉，附：[DS115j 网络性能测试](./report/mini-server/ds115j.md)。
- [4] 打印不失为一种相对稳定的持久化保存方案。
- [5] 干掉持久化的纸质存储，最靠谱的莫过于加密级别的粉碎了。
- [5] 在所有电源都带稳流稳压作用后，添加一台UPS防止市电闪断带来的数据损坏或者写输出脏掉的问题。
- 正在缓慢的把DSM兼容机的储存慢慢移动到白裙 & 冷备中，以及慢慢把所有非高频HDD都换成SSD。

### 📱 移动设备 & 🎮 游戏设备

> 强依赖无线进行交互的设备

| 资源类型 | 明细 | 网络 | 储存 | 开始服务 |
| --- | --- | --- | --- | --- |
| 平板 | iPad Air2 | 4G / 5G WIFI | 128G (改) | 2015 |
| 爪机 | iP8p | 4G / 5G WIFI | 256G | 2017 |
| 爪机 | iP7  | 4G / 5G WIFI | 128G | 2016 |
| 爪机 | iP6p | 4G / 5G WIFI | 64G  | 2014 |
| 爪机 | iPse | 4G / 5G WIFI | 64G  | 2017 |
| 爪机 | iP5s ^1 | 4G / 5G WIFI | 忽略  | 2017 |
| 爪机 | iP5 ^2 | 3G / 5G WIFI | 16G | 2017 |
| 爪机 | iP3GS ^3 | 2G / 2G WIFI | 忽略 | 2017 |
| 爪机 | S7E ^4 | 4G / 5G WIFI | 忽略 | 2016 |
| 游戏机 | PSVx2 ^5 | 2G WIFI | 16G / 64G | 2015 / 2016 |
| 游戏机 | 3DSx2 ^6 | 2G WIFI | 64G / 64G | 2014 / 2016 |
| 游戏机 | PS4 Pro | 2G WIFI | 500G SSD | 2017 |
| 游戏机 | PS4 | 2G WIFI | 500G HDD | 2017 |

- [1] 之前diy外壳的玩具，在之前公司上班的时候比较反感要安装手机证书，于是作为工作备机使用，隔离工作和生活数据，不需要考虑备份功能。
- [2] 另外一个diy外壳的玩具，作为家里MP4使用，投个音乐不用考虑电话进来，游戏音干扰，挺好的。
- [3] 年度最值手机，作为monitor使用，极低功耗，可以愉快跑脚本，已购两台。
- [4] 没有特别需求，不会再入三星的机器了，稍后补充吐槽（关于闪付／更新），目前当NFC充值POS机使用，冲地铁卡。
- [5] 一正一破，好处是可以联机的游戏，可以大号带小号玩，另外可以做PS4手柄，玩不需要L2R2键的游戏还可以。
- [6] 游戏性比SONY强，联机赛高！
- 耐心等待18年新Pad上市中。


### 🔮 智能设备

> 所谓智能不过是可编程或者扩展了原有的使用者能力

| 资源类型 | 明细 | 网络 |备注 |
| --- | --- | --- | --- |
| 网络插座  | 小米 xN ^1 | 2G WIFI | 2017年  |
| 网络网关  | 小米 x2 ^2 | 2G WIFI | 2017年  |
| 传感装置  | 小米 xN ^3 | 2G WIFI/ZigBee | 2017年  |
| 网络摄像头 | 水滴 x4 ^2 | 2G WIFI | 2015年、2016年、2017年、2018年 |

- [1] 小米开放局域网协议之后，把控客都换成了小米，支持编程这件事太好了。
- [2] 小米网关+空调插座，对于空间利用率很高，而且支持ZigBee/WiFi，还能网络调试。
- [3] 门磁可以避免出门老想着有没有关好门的问题，烟雾等传感器避免检查厨房，还是省心。
- [4] 2015年从小米换到360就一直使用，目前使用还好，小问题是启动音有点惊悚，4台设备经常会有一台显示离线。
- 控客的插座APP不是很好用，尤其是不给SDK，无法定制开发，国内版本也不支持简单DIY（需要拆+编程器），故弃用。


### 💣 历史设备

> 断电的设备

| 资源类型 | 明细 | 备注 | 开始服务 |
| --- | --- | --- | --- |
| 交换机 | NETGEAR GS308 ^1 | 已出闲置，千兆LAN x8 ，由于有强力党的加入 | 2016 |
| ~~路由器~~ | ~~NETGEAR WNDR4300 ^2~~ | 已出闲置，全千兆，双频，双128MB，其实蛮靠谱的，入了两台了 | 2014 |
| ~~工控机~~ | ~~N270, 945GM x2 ^3~~ | 双网口x1, 单网口x1，由于来了一堆强力党，直接扔掉了 | 2015 |
| ~~工控机~~ | ~~D525~~ ^4 | 已送人 | 2016 |
| ~~工控机~~ | ~~D425~~ ^5 | 已丢弃 | 2015 |
| ~~代码仓库~~ | ~~ONDA v939 Core M~~ ^6 | 4G / 2G WIFI | 4GBRAM / 128GB | 2016 |
| ~~资源备份x3~~ | ~~DSM 兼容机~~ ^7 | 千兆LAN | - | 2016 / 2017 |

- [1] 不带E的其实也够用，但是需要有线连接的设备多了之后，口不够用了。
- [2] 用于软路由的AP稳定性特别棒，有NetGear的辐射合格的背书，使用放心。不过芯片比较老了，支持的固件和适合运行的自定义功能有限，故放弃使用。
- [3] N3700 组装机虚拟的NAS代替了工控机作为下载和储存机的作用，故废弃。
- [4] 功能与N3700 组装机重复，不过配个启动U盘，刷个群晖／跑几个简单的docker镜像还不错，20W的功率感觉略鸡肋，7x24跑到是也还好。
- [5] 同上，功率差不多，但是性能差太多，更加鸡肋...
- [6] Core M 5y10c 2.0GHz，其实性能很不错，能完美跑Ubuntu/OSx/Win/Android，然而散热不足导致会降频，不能忍耐，已转出。
- [7] 主机素质不错，但是我入了新的设备，机柜放不下了，如果屋子大些就好了。


## 具体实践

- [组件家用迷你服务器](./build-mini-home-server.md)
- [家用10～20台在线设备可参考网络方案](./network.md)
- [远程访问及数据交换](remote.md) ^1
- [日志收集和查看](log.md) ^1
- [虚拟化技术应用](./virtual.md) ^1
- [1] 待更新。

## 一些观点

- [2016年上半年 自组x86软路由/迷你服务器的一些考虑](./think-about-x86-route-2016.md)


