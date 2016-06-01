## 项目说明

    记录搭建家用娱乐网络环境。

### 目的

- 满足高效备份数据/系统需求。
- 满足切换设备无缝写码需求。
- 满足切换设备无缝调试需求。
- 构建自己的发布流程。
- 添加自己想要的持续集成工具。
- 满足高速家用网络传输需求。
- 满足偶尔无聊的游戏需求。

## 保持联网设备清单

目前省钱(屌丝)方案中使用的设备:

- 宽带资源:
    - 北京联通50M   多拨可到7~9MB/s (搬家后受限办理小区宽带);
    - 北京联通4M    使用迅雷快鸟，可以叠到6MB/s
- 主机资源:
    - 工作机器: MacBook Pro (2014年，i7 2.2GHz，16GBRAM，Retina，千兆LAN & 5G WIFI)
    - 编码机器: MacBook Pro (2014年，i7 2.5GHz，16GBRAM，Retina，千兆LAN & 5G WIFI)
    - 资源机器: HASEE Z7    (2015年，i7 2.6GHz，32GBRAM，千兆LAN & 5G WIFI)
    - 功能机器: N3700组装机  (2016年，N3700 1.6GHz，8GBRAM，千兆LANx4 & 5G WIFI)
- 储存备份:
    - 辅助备份: WD MY CLOUD 3T (2015年，千兆LAN)
    - 主要备份: WD MY CLOUD 4T (2014年，千兆LAN)
    - 临时储存: N3520 组装机 (2016年，N3520 2.166GHz，4GBRAM，千兆LAN)
- 移动设备:
    - iPad Air2     (2015年，5G WIFI & 4G)
    - S7E           (2016年，5G WIFI & 4G)
    - PSV           (2015年，2G WIFI)
- 无线AP:
    - NETGEAR WNDR4300  (2014年，全千兆，双频，双128MB，其实蛮靠谱的，入了两台了)
    - Xiaomi Mini 第一版 (2015年/2016年，刷机之后蛮稳定的，硬伤是百兆LAN口，入了三台了)
    - Xiaomi Mini 青春版 (2016年，功耗极低，入了一台)
- 智能设备:
    - 智能插座 x5       （2015年）
    - 智能摄像头 x1      (2015年，之前用的小米，后来换了360水滴)
- 历史设备: (已断电)
    - 工控机 N270, 945GM x2 (双网口x1, 单网口x1), D525 x1, D425 x1

### 使用过的设备历史记录（以网络设备为主/娱乐辅助为辅）:

- 主机:
    - [HASEE Z7](./devices/HASEE-Z78172R2.md)
    - 2013年 MacBook Pro (2013年，i7 2GHz，8G内存，Retina)
    - 2012年12月 [Lenovo Y485](./devices/Lenovo-Y485.md)
- 路由器:
    - [NETGEAR WNDR4300 750M](./devices/NetGear-WNDR4300.md)
    - [NETGEAR WNDR2000 300M](./devices/NetGear-WNR2000.md)
    - [XiaoMi Route Mini](./devices/XiaoMi-Route-Mini.md)
    - [TP-LINK MR12U](./devices/TP-LINK-MR12U.md)
    - 2013年01月 [TP-LINK WN703N](./devices/TP-LINK-WR703N.md)
    - 2014年05月 [NETGEAR WN1000RP](./devices/NetGear-WN1000RP.md)
- 储存:
    - 2016年04月 [WD My Cloud 4T](./devices/WD-My-Cloud-4T.md)
    - 2014年12月 [WD My Cloud 3T](./devices/WD-My-Cloud-3T.md)
- 移动:
    - 2015年10月 [iPad Air2 JP Ver](./devices/iPad-Air2-JP.md)
    - 2015年09月 [iPad Air2 CN Ver](./devices/iPad-Air2-CN.md)
    - 2015年05月 [iPad Air2 Cellular](./devices/iPad-Air2-Cell.md)
    - 2015年01月 [Kubi iwork7](./devices/KB-iwork7.md)
    - 2015年01月 [Kubi iwork8](./devices/KB-iwork8.md)
    - 2013年06月 [Kindle Paper White 1st](./devices/Kindle-paper-white-1.md)
- 网卡:
    - 2015年02月 [HUAWEI EC8201 3G Card](./devices/HUAWEI-EC8201-3G-Card.md)
    - 2013年8月 [HUAWEI E261 3G Card](./devices/HUAWEI-E261-3G-Card.md)
- 显示器:
    - 2412Hx2 (DP+MiniDP+HDMI) + 2312HM(HDMI+VGA)
    - 2013年11月 [Dell U2312HM]()

## 需求分析

- 因为工作环境需求，Mac依旧需要使用OSx。
- 构建发布流程期望通用且可以方便移植到其他系统的机器上，需要使用虚拟化方式实现。
	- 虚拟化技术方案：虚拟机 or Docker
	- 虚拟化系统选择：非Win和OSx之外，系统选择范围 Ubuntu或CentOS
- 为了玩使命召唤，使用Z7作为Windows下的娱乐机。
- 高速网络利用：
	- 和舍友共用TPLINK WR841一只，(固件原厂不折腾)放家中作主路由使用，主路由和4300之间用超六类线连接，4300和笔记本之间用超六类连接，以便再扩充设备后，内部可以进行千兆数据交换。
	- 无线利用，详见 #无线使用#
	- 有线利用，详见 #有线使用#
- 持续集成工具功能需求：
	- 版本控制
	- 代码编译
	- 代码构建
	- 代码同步（自用不一定使用打包分发的方式）
	- 代码测试 & Lint
	- 切换发布版本和环境（灰度）
- 发布流程想折腾的需求：
	- 分离环境
		- 日常（本地+虚拟机）
		- 预发（虚拟机二套环境）
		- 线上（远程VPSx3）
	- 分离代码
		- 服务脚本（偏后）
		- 前端资源
		- 数据层代码（DB）
	- CLI TRIGGER EVENT
- 无缝写码环境 && 无缝测试环境
	- 路由表策略 && DHCP半动态
- 无缝DLNA使用

## 具体实践

- [基础网络使用](network.md)
- [远程访问及数据交换](remote.md)
- [无线使用](wifi.md)
- [日志收集和查看](log.md)
- [虚拟机使用](vm.md)
- [有线使用](lan.md)