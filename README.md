🚧 **Under Construction / 持续更新** 🚧

![Project Start in 2014](https://img.shields.io/badge/Project%20Start-2014-brightgreen?style=flat-square&logo=azuredataexplorer) ![GitHub last commit](https://img.shields.io/github/last-commit/soulteary/Home-Network-Note.svg?label=Last%20Update&style=flat-square&logo=github) ![GitHub](https://img.shields.io/github/license/soulteary/Home-Network-Note.svg?style=flat-square&logo=lichess)

**这几年积累/分享了不少实践内容，是时候把这个项目的坑填上了。**

## 项目说明

```text
记录搭建兼顾学习娱乐的家用网络环境的过程，折腾过的一些软硬件小经验。

目前的网络方案从 2016 年使用至今，非常稳定，整体架构几乎没有变化。（日常在线 25~40 台设备，峰值 50+ ）

文档中的方案和方案中的配置会尽可能**保持简单**，确保各种服务在运行一年之后，我依旧能够对软硬件进行**轻松简单**的“维护、升级以及替换”操作。
```

## 主要场景和关键词

列举常见家用网络场景的一些核心诉求的关键词。

|     | 设备联网 | 备份数据 | 下载上传 | 数据同步 | 开发学习 | 游戏娱乐 |
| --- | --- | --- | --- | --- | --- | --- |
| **核心指标** | 安全、稳定   | 安全、可靠 | 高速 | 无感知 |  流畅  | 流畅 |
| **重要因素** | 简单、易维护  | 高效     | 易用 |  准确  |  省心  | 舒适 |
| **可选因素** | 网速、组网模式 | 易用     | 安全 | 跨平台 | 冗余保障 | |


- [在GitHub上查看手册](./handbook/SUMMARY.md)
- [浏览部署后的网站](https://home-network-note.suyang.wiki/)

### 场景举例

- [1] 家庭环境中的手机、平台、电脑、NAS 等设备，均能够快速的完成数据自动同步，针对宝贵的照片或文档的备份能够高效完成。
- [2] 网络环境能够支持各种数据的周期性备份（每日、每周），针对代码、文档、笔记能够进行额外的数据备份保护。
- [3] 稳定的网络，能够保障日常观看在线视频、进行低延迟的游戏对战、资料的快速下载，以及开发学习相关试验诉求的数据获取和分析。
- [4] 当数十台设备接入网络的时候，网络需要依旧保持低延迟，以及流畅的内外网交互，无须针对路由器相关网络设备或联网电子设备、智能设备、虚拟机进行额外的操作，包括并不仅限于重启、执行复杂脚本监控等。
- [5] 内网设备和运行服务均处于相对安全的网络环境中，允许以安全的方式从外网进入开发环境网络，使用网络内的设备进行开发调试。


## 不完全设备列表

```text
如果你考虑入手一些设备，比如笔电、路由、网卡、显示器、硬盘、移动和娱乐设备，或许可以从这里的到一些参考信息，节约点银子，让荷包💰保持郁郁葱葱。
```

[设备清单(待补全)](./handbook/devices/device-list.md) | [显示器相关](./handbook/display/history.md) | [💣 断电不再使用的设备](./handbook/devices/obsolete.md)

## 屋内日常保持联网设备清单及简明方案

![🏠 网络结构](./handbook/network/assets/topology/network-2017.10.png)

- [家用 20～30 台在线设备参考网络方案](./handbook/network/network-topology.md)
- [无线网络方案及性能数据](./handbook/network/wifi-topology.md)
- [有线网络方案及性能数据](./handbook/network/lan-topology.md)

## 🌈 宽带资源

> 不敢想假如家里没有稳定的网络会怎样

| 资源类型 | 明细 | 备注 |
| --- | --- | --- |
| 北京联通 | 固网 1000M | 下行 900Mbps / 上行 40Mbps |
| 北京联通 | 4G 路由 | 备份网络，保障智能家居和监控使用，及远程维护主网络设备 |
| 北京电信 | 5G | 主网络，用于日常上网、热点 |
| 北京联通 | 5G | 备份网络，用于补充信号覆盖 |

👉 [完整说明](https://github.com/soulteary/Home-Network-Note/blob/main/network/broadband.md) / **简要说明** 👇

- [1] 如果没有特殊需求，建议简化不必要的多线宽带，避免策略路由带来的各种问题，以及避免使用使用软路由聚合不同类型宽带，带来后续维护上的麻烦。
- [2] 带宽使用过程中不建议依赖任何提速软件，避免当软件不可用时，带宽质量严重受损，以及带来的额外维护“提速软件”运行环境带来的成本。
- [3] 个人体验原因，已停用移动电话卡，取消原本无线三网接入。



## ⭐️ 路由网关

> 影响网络质量的核心设备，负责部分网络安全事务，[历史设备见文档](https://github.com/soulteary/Home-Network-Note/blob/main/handbook/devices/device-list.md)。

| 资源类型 | 明细 | 网络 | 开始服务 |
| --- | --- | --- | --- |
| 交换机 | NETGEAR GS116E ^1 | 千兆LAN x16 | 2017 |
| 路由器 | XiaoMi BE6500 Pro ^2 | 2G WiFi / 5G Wi-Fi（WiFi7） / 2.5G LAN | 2023 |
| 路由器 | Newifi D2 Plus ^3 | 2G WiFi / 5G WiFi / 千兆LAN | 2021 |
| 路由器 | Xiaomi Mini 青春版 ^4 | 2G WiFi / 百兆LAN | 2016 |
| 路由器 | 施耐德旅行插座 ^5 | 2G WiFi / 百兆LAN | 2018 |

👉 [完整说明](https://github.com/soulteary/Home-Network-Note/blob/main/network/gateway.md) / **简要说明** 👇

- [1] 十六口千兆交换机，用于扩展主路由网络吞吐能力，带 Web 管理界面，最大功耗仅10w，铁壳散热，替换之前的[八口交换机](https://github.com/soulteary/Home-Network-Note/blob/main/deprecate/NetGear%20GS308.md)。
- [2] 扩展主路由的 AP 能力，提供屋内设备以 WiFi6 - WiFi7 模式，进行高速无线接入，以及米家本地智能策略网关，用于替代之前的 [Mi AIoT AX 3600](https://github.com/soulteary/Home-Network-Note/blob/main/deprecate/Mi%20AIoT%20Router%20AX3600.md)。
- [3] 21 年再次更新，替换到之前已经使用三年的[普通版本](https://github.com/soulteary/Home-Network-Note/blob/main/deprecate/Newifi%20D2%20(Newifi3).md)，全千兆的四口主路由（二级路由），拥有 512M 的内存和铁壳、大面积铜片散热器，延续 Newifi 极高的性价比，一度使用两台相同规格的设备作为拨号路由器和二级路由。相关文章：[近期家用设备（NUC、猫盘、路由器）散热升级记录](https://soulteary.com/2021/10/14/recent-heat-dissipation-upgrade-record-of-homelab.html)。
- [4] 相比较[正常体积的 Mini](https://github.com/soulteary/Home-Network-Note/blob/main/deprecate/Xiaomi%20Mi%20WiFi%20Mini.md) 或者随后出品的 [AC 2100](https://github.com/soulteary/Home-Network-Note/blob/main/deprecate/Xiaomi%20Mi%20Router%20AC2100.md)，青春版就小巧玲珑多了。除了小巧方便、颜值较高之外，路由功耗极低、三方固件也支持各种各样的 DIY，除了百兆带宽在当前看来比较不够用之外，一切都很好。如果是用在调试开发场景，那么还是比较合适的。
- [5] 此插座自带一个简易的热点 WiFi 功能，如果你需要插上设备就自动组网，可以使用上面的设备，如果你没有自动组网等需求，那么这个能让你上网的插座，用起来体验还不错，唯一槽点，插座本身发热比较严重，或许对网络稳定性/质量有一定影响。



## 💻 主机资源

> 提供运算能力的本地设备，[历史设备见文档](./handbook/devices/device-list.md)。

### MacBook Air (15-inch, 2023) ![Turn On:2023](https://img.shields.io/badge/Turn%20On-2023-brightgreen?style=flat-square)

> 轻便、长续航，一切刚刚好的打字机。

![MacOS:Sonoma 14.2](https://img.shields.io/badge/MacOS-Sonoma%2014.2-brightgreen?style=flat-square&logo=apple) ![CPU:Apple M2](https://img.shields.io/badge/CPU-Apple%20M2(4P+4E)-brightgreen?style=flat-square&logo=Apple) ![RAM:24GB](https://img.shields.io/badge/RAM-24GB-brightgreen?style=flat-square) ![Disk:1TB](https://img.shields.io/badge/Disk-1TB-brightgreen?style=flat-square) ![LAN:1000M](https://img.shields.io/badge/LAN-1000M-brightgreen?style=flat-square) ![WiFi:WiFi6](https://img.shields.io/badge/WiFi-WiFi6-brightgreen?style=flat-square)

2023 年，把主力打字机替换为了拥有更好续航，更少发热，运行几乎没有噪音的 M2 版本的 Air。

随着 Apple 停止生产 Intel 芯片的设备，除了验证构建软件的场景外，再继续持有设备，并作为主力设备有一些不明智，尤其是我的 [19 款高配版本 MBP](https://github.com/soulteary/Home-Network-Note/blob/main/deprecate/MacBook%20Pro%20(16-inch%2C%202019).md)在关键时刻出现了两次过热黑屏。

性能方面，一方面我将更需要密集计算的任务移动到了其他的设备中，另外一方面，这台设备的性能比我之前的 Intel I9 处理器只高不低（单核+37%、多核+15%、FP32+83%），虽然在构建 Docker、使用一些小众软件过程中有细节的影响，但总归体验还是提升的。

日常不论是出差还是在外通勤、临时会议，支撑一天的电量毫无压力，加上笔记本的重量非常友好，或许我的下一台打字机设备，还会是 Air。

- [官方设备规格链接](https://support.apple.com/zh-cn/111346)

### ThinkBook 14 2023 (Intel) ![Turn On:2023](https://img.shields.io/badge/Turn%20On-2023-brightgreen?style=flat-square)

> 多核心、高性能、满血雷电的便携计算平台。

![Ubuntu:23.10](https://img.shields.io/badge/Ubuntu-23.10-brightgreen?style=flat-square&logo=Ubuntu) ![CPU:Intel I7 13700H](https://img.shields.io/badge/CPU-Intel%20I7%2013700H(14C(6P+8E)20T)-brightgreen?style=flat-square&logo=Intel) ![RAM:64GB](https://img.shields.io/badge/RAM-64GB-brightgreen?style=flat-square) ![Disk:4TB](https://img.shields.io/badge/Disk-4TB-brightgreen?style=flat-square) ![LAN:1000M](https://img.shields.io/badge/LAN-1000M-brightgreen?style=flat-square) ![WiFi:WiFi6](https://img.shields.io/badge/WiFi-WiFi6-brightgreen?style=flat-square)

2023 年国庆的时候，考虑到各种原因，我将主力输入设备换成了 MacBook Air，这几年的 Mac 设备只有 Apple Silicon 芯片（M1、M2、M3 这类 ARM 芯片），好多应用和 Docker 的兼容性真的还是有很大问题。

我就开始想，能不能多带一台轻便易携多设备出门，解决日常的开发体验、方案验证问题，还可以避免糟糕网络带来的糟糕的远程使用体验和可访问性的问题，并且成本还低。当然，主力输入设备重量降低到了 1.5kg，也是我可以再考虑多带一台设备的原因。

在一番简单筛选后，找到了联想 ThinkBook 14，14寸的机身，原生支持雷电4（CPU而非外置芯片），拥有双内存插槽、双硬盘槽、有 13500h 或 13700h 两种 CPU 型号选择的笔记本，设备重量也只有 1.4kg，基础功耗不到 30w。

这台设备的算力相比较之前的 Intel I9 MacBook Pro，至少性能提升 38%，有一说一，五千块的基础价格，还是非常香的。

当我需要额外算力，又不需要显卡的时候，带两台笔记本，使用雷电线将两台设备连接到一起后，我就拥有了更多的核心、标准的 Linux 操作系统环境、更大的存储空间，甚至是一块副屏。

- [官方设备规格链接](https://psref.lenovo.com/syspool/Sys/PDF/ThinkBook/ThinkBook_14_G6_IRL/ThinkBook_14_G6_IRL_Spec.pdf)
- 相关文章记录
    - 《[家用工作站方案：ThinkBook 14 2023 版](https://soulteary.com/2023/11/09/home-workstation-solution-thinkbook-14-2023.html)》

### ThinkPad L14 Gen 1 (AMD) ![Turn On:2021,2022](https://img.shields.io/badge/Turn%20On-2021,2022-brightgreen?style=flat-square)

> 低成本、多核心数、高内存规格的笔记本。

![Ubuntu:22.04](https://img.shields.io/badge/Ubuntu-22.04-brightgreen?style=flat-square&logo=Ubuntu) ![CPU:AMD Zen2 4750u](https://img.shields.io/badge/CPU-AMD%20Zen2%204750u(8C16T,%201.7GHz)-brightgreen?style=flat-square&logo=AMD) ![RAM:64GB](https://img.shields.io/badge/RAM-64GB-brightgreen?style=flat-square) ![Disk:2TB](https://img.shields.io/badge/Disk-2TB-brightgreen?style=flat-square) ![LAN:1000M](https://img.shields.io/badge/LAN-1000M-brightgreen?style=flat-square) ![WiFi:802.11ac](https://img.shields.io/badge/WiFi-802.11ac-brightgreen?style=flat-square)

之前一直在寻找核心数多、功耗低、支持 64GB 内存、相对便携的无显卡笔记本设备，以做“廉价的服务器”使用，并替代早先时候购置的 [HP EliteDesk G4 800](./handbook/devices/HP-EliteDesk-G4-800.md) 小型工作站，直到遇到了搭载 AMD Zen2 4750u 的 ThinkPad L14。

这台设备满载仅 45w 的功耗，性能极强，核心数也非常多，特别适合长时间跑容器服务，来扩展本地的计算资源。美中不足的是，设备只有一条固态硬盘插槽可用。

2022年，设备价格进一步下降，增加了一台相同 CPU 配置的新设备作为冗余资源（32G/1TB）。

- [官方设备规格链接](https://psref.lenovo.com/syspool/Sys/PDF/ThinkPad/ThinkPad_L14_Gen_1_AMD/ThinkPad_L14_Gen_1_AMD_Spec.pdf)
- 相关文章记录
    - 《[廉价的家用工作站方案：前篇](https://soulteary.com/2021/07/02/cheap-home-workstation-solution-part-one.html)》
    - 《[AMD 4750u 及 5800u 笔记本安装 Ubuntu 20.04](https://soulteary.com/2021/07/04/install-ubuntu-20-04-on-amd-4750u-and-5800u-laptops.html)》
- 同宗同源，但已出掉的 5800u (ThinkBook 15)
    - 《[AMD 5800u 笔记本折腾 Proxmox VE 7.0 虚拟化](https://soulteary.com/2021/10/23/amd-5800u-notebook-toss-proxmox-ve-7-0-virtualization.html)》、《[装在笔记本里的私有云环境：准备篇](https://soulteary.com/2021/10/28/private-cloud-environment-installed-in-a-notebook-preparation.html)》、《[装在笔记本里的私有云环境：监控篇](https://soulteary.com/2021/10/30/private-cloud-environment-installed-in-a-notebook-monitor.html)》、《[装在笔记本里的私有云环境：网络存储篇（上）](https://soulteary.com/2021/11/04/private-cloud-environment-installed-in-a-notebook-storage-part-1.html)》、 《[装在笔记本里的私有云环境：网络存储篇（中）](https://soulteary.com/2021/11/09/private-cloud-environment-installed-in-a-notebook-storage-part-2.html)》、 《[装在笔记本里的私有云环境：K8s 集群准备](https://soulteary.com/2022/11/29/private-cloud-environment-installed-in-a-notebook-k8s-cluster-preparation.html)》

### Intel NUC8i5BEH ![Turn On:2021](https://img.shields.io/badge/Turn%20On-2021-brightgreen?style=flat-square)

> 一台多面手，目前变成了一台 Apple TV。

![MacOS:Ventura 13.2](https://img.shields.io/badge/MacOS-Ventura%2013.2-brightgreen?style=flat-square&logo=apple) ![CPU:Intel I5-8259U](https://img.shields.io/badge/CPU-Intel%20I5--8259U(4C8T,%202.3GHz)-brightgreen?style=flat-square&logo=Intel) ![RAM:64GB](https://img.shields.io/badge/RAM-64GB-brightgreen?style=flat-square) ![Disk:2TB](https://img.shields.io/badge/Disk-2TB-brightgreen?style=flat-square) ![LAN:1000M](https://img.shields.io/badge/LAN-1000M-brightgreen?style=flat-square) ![WiFi:802.11ac](https://img.shields.io/badge/WiFi-802.11ac-brightgreen?style=flat-square)

最初的入手原因见[这篇文章](https://soulteary.com/2021/01/31/nuc-notes-linux-system.html)，随后在不断添置和更新新设备后，这台设备在职能转变为了一台安装了 ESXi 的编写虚拟机“母鸡”，支持了大量开源软件、尤其是具备分布式使用场景的软件的构建和部署测试相关的工作，极大的解放和降低了我的那台动不动就会风扇喧嚣起来的 Intel 芯片的 MacBook。

随着设备越来越多，虚拟化相关的任务不再需要它执行，目前它变回了一台 “Apple TV”，安静的放在真正的 Apple TV 旁边，为我能够看到视频网站弹幕继续做出它的贡献。

- [官方设备规格链接](https://ark.intel.com/content/www/cn/zh/ark/products/126148/intel-nuc-kit-nuc8i5beh.html)
- 相关文章记录
    - 《[NUC 折腾笔记 - Linux 系统篇](https://soulteary.com/2021/01/31/nuc-notes-linux-system.html)》
    - 《[NUC 折腾笔记 - 储存能力测试](https://soulteary.com/2021/02/02/nuc-notes-storage-ability-test.html)》
    - 《[NUC 折腾笔记 - 安装 ESXi 7](https://soulteary.com/2021/06/22/nuc-notes-install-esxi7.html)》
    - 《[近期家用设备（NUC、猫盘、路由器）散热升级记录](https://soulteary.com/2021/10/14/recent-heat-dissipation-upgrade-record-of-homelab.html#%E6%9B%B4%E6%8D%A2-nuc8-%E7%9A%84%E6%9C%BA%E7%AE%B1%E5%92%8C%E9%A3%8E%E6%89%87)》

> 相同设计风格的、曾经用于跑多种服务的 **Intel NUC7CJYH (2021 购入)**，已经转手群友。

2021 年，决定让群晖专注存储，将群晖上运行的容器都迁移到过这台小机器。在解决了群晖的硬盘因为跑容器，出现的勤奋运转绝不休眠的问题后，在这台机器上搭建和运行了两年个人 Wiki。

- [官方设备规格链接](https://www.intel.cn/content/www/cn/zh/products/sku/126135/intel-nuc-kit-nuc7cjyh/specifications.html)
- 相关文章记录
    - 《[NUC7 一台并不弱的 Mini PC](https://soulteary.com/2021/10/19/nuc7-a-mini-pc-that-is-not-weak.html)》

### Dell OptiPlex 3060 ![Turn On:2022](https://img.shields.io/badge/Turn%20On-2022-brightgreen?style=flat-square)

> 重新建设中，用于获取和处理各种 RSS 信息。

![VMware:ESXi 6.7u3](https://img.shields.io/badge/VMware-ESXi%206.7u3-brightgreen?style=flat-square&logo=VMware) ![CPU:I3-9100T](https://img.shields.io/badge/CPU-I3--9100T(4C4T,%203.10GHz)-brightgreen?style=flat-square&logo=I3-9100T) ![RAM:16GB](https://img.shields.io/badge/RAM-16GB-brightgreen?style=flat-square) ![Disk:512GB](https://img.shields.io/badge/Disk-512GB-brightgreen?style=flat-square) ![LAN:1000M](https://img.shields.io/badge/LAN-1000M-brightgreen?style=flat-square) ![WiFi:802.11ac](https://img.shields.io/badge/WiFi-802.11ac-brightgreen?style=flat-square)


重新建设中，用于获取和处理各种 RSS 信息。

### Lenovo 9000K ![Turn On:2023](https://img.shields.io/badge/Turn%20On-2023-brightgreen?style=flat-square)

> 目前已重置，待重新投入使用。

![Ubuntu:23.04](https://img.shields.io/badge/Ubuntu-23.04-brightgreen?style=flat-square&logo=Ubuntu) ![CPU:Intel I9-13900KF](https://img.shields.io/badge/CPU-Intel%20I9--13900KF(24C32T,%203.0GHz)-brightgreen?style=flat-square&logo=Intel) ![GPU:Nvidia 4090](https://img.shields.io/badge/GPU-Nvidia%204090%2024GB-brightgreen?style=flat-square&logo=Nvidia) ![RAM:64GB](https://img.shields.io/badge/RAM-64GB-brightgreen?style=flat-square) ![Disk:2TB](https://img.shields.io/badge/Disk-2TB-brightgreen?style=flat-square) ![LAN:1000M](https://img.shields.io/badge/LAN-1000M-brightgreen?style=flat-square) ![WiFi:802.11ac](https://img.shields.io/badge/WiFi-802.11ac-brightgreen?style=flat-square)

机器性能不错，用来做本地模型验证。

- [官方设备规格链接](https://item.lenovo.com.cn/product/1026169.html)

### Mac Pro (2013) ![Turn On:2023](https://img.shields.io/badge/Turn%20On-2023-brightgreen?style=flat-square)

> 为 128GB ECC RAM、高颜值、长期密集运算准备的设备。

![MacOS:Ventura 13.2](https://img.shields.io/badge/MacOS-Ventura%2013.2-brightgreen?style=flat-square&logo=apple) ![CPU:Intel E5-2697v2](https://img.shields.io/badge/CPU-Intel%20E5--2697v2(12C24T,%202.7GHz)-brightgreen?style=flat-square&logo=Intel) ![GPU:AMD FirePro D300](https://img.shields.io/badge/GPU-AMD%20FirePro%20D300%204GBx2-brightgreen?style=flat-square&logo=AMD) ![RAM:128GB(ECC)](https://img.shields.io/badge/RAM-128GB(ECC)-brightgreen?style=flat-square) ![Disk:2TB](https://img.shields.io/badge/Disk-2TB-brightgreen?style=flat-square) ![LAN:1000M](https://img.shields.io/badge/LAN-1000M-brightgreen?style=flat-square) ![WiFi:802.11ac](https://img.shields.io/badge/WiFi-802.11ac-brightgreen?style=flat-square)

ECC RAM 保障密集计算时，数据绝对正确，线程足够多即使对比前两年的设备也毫不逊色。唯一缺点，相对功耗较高。

- [官方设备规格链接](https://support.apple.com/kb/SP697?viewlocale=zh_CN&locale=zh_CN)
- 相关文章记录
    - 《[廉价的家庭工作站设备改造记录：苹果垃圾桶（Mac Pro 2013）](https://soulteary.com/2023/02/04/cheap-home-workstation-solution-mac-pro-2013.html.html)》。

### HP EliteDesk 800G6 ![Turn On:2023](https://img.shields.io/badge/Turn%20On-2023-brightgreen?style=flat-square)

> 使用各种淘汰硬件重置的另外一台全闪存 NAS。

![VMware:ESXi 7.0u3](https://img.shields.io/badge/VMware-ESXi%207.0u3-brightgreen?style=flat-square&logo=VMware) ![CPU:Intel I3-9100T](https://img.shields.io/badge/CPU-Intel%20I3--9100T(4C4T,%203.10GHz)-brightgreen?style=flat-square&logo=Intel) ![RAM:32GB](https://img.shields.io/badge/RAM-32GB-brightgreen?style=flat-square) ![Disk:5.5TB](https://img.shields.io/badge/Disk-5.5TB-brightgreen?style=flat-square) ![LAN:1000M](https://img.shields.io/badge/LAN-1000M-brightgreen?style=flat-square) ![WiFi:802.11ac](https://img.shields.io/badge/WiFi-802.11ac-brightgreen?style=flat-square)

已重置，待重新投入使用。

- [官方设备规格链接](https://h20195.www2.hp.com/v2/getpdf.aspx/c06648254.pdf)
- 相关文章记录
    - 《[低成本搭建一台家庭存储服务器：前篇](https://soulteary.com/2023/01/15/building-a-home-storage-server-at-low-cost-part-one.html)》
    - 《[低成本搭建一台 Unraid 家庭存储服务器：中篇](https://soulteary.com/2023/01/21/building-a-home-unraid-storage-server-at-low-cost-part-two.html)》
    - 《[低成本搭建一台家庭存储服务器：全闪存篇](https://soulteary.com/2023/09/17/building-a-home-storage-server-at-low-cost-part-pure-flash.html)》

### Intel NUC9i5QNX ![Turn On:2023](https://img.shields.io/badge/Turn%20On-2023-brightgreen?style=flat-square)

> 构建桌面全闪存 NAS / DAS 的准系统主机。

![VMware:ESXi 7.0u3](https://img.shields.io/badge/VMware-ESXi%207.0u3-brightgreen?style=flat-square&logo=VMware) ![CPU:Intel I5-9300H](https://img.shields.io/badge/CPU-Intel%20I5--9300H(4C8T,%202.4GHz)-brightgreen?style=flat-square&logo=Intel) ![RAM:16GB](https://img.shields.io/badge/RAM-16GB-brightgreen?style=flat-square) ![Disk:2TB](https://img.shields.io/badge/Disk-2TB-brightgreen?style=flat-square) ![LAN:2000M](https://img.shields.io/badge/LAN-2000M-brightgreen?style=flat-square) ![WiFi:WiFi6](https://img.shields.io/badge/WiFi-WiFi6-brightgreen?style=flat-square)

这台设备入手原因见[这篇文章](https://soulteary.com/2023/08/31/cheap-pure-flash-thunderbolt-nas-tossing-notes-the-choice-of-networking-solutions.html)，目前正在使用它进行全闪存雷电 NAS 的搭建。

- [官方设备规格链接](https://ark.intel.com/content/www/us/en/ark/products/190104/intel-nuc-9-extreme-kit-nuc9i5qnx.html)
- 相关文章记录
    - 《[廉价的全闪存雷电 NAS 折腾笔记：组网方案的选择](https://soulteary.com/2023/08/31/cheap-pure-flash-thunderbolt-nas-tossing-notes-the-choice-of-networking-solutions.html)》
    - 《[廉价的全闪存雷电 NAS 折腾笔记：NUC9 操作系统踩坑](https://soulteary.com/2023/09/12/cheap-pure-flash-thunderbolt-nas-tossing-notes-nuc9-operating-system-pitfalls.html)》

### MacBook Pro M2 ![Turn On:2022](https://img.shields.io/badge/Turn%20On-2022-brightgreen?style=flat-square)

> 因为是 ARM 架构，使用比较少的设备，当然，偶尔用它来测试“模型”。

![MacOS:Ventura 13.2](https://img.shields.io/badge/MacOS-Ventura%2013.2-brightgreen?style=flat-square&logo=apple) ![CPU:Apple M2](https://img.shields.io/badge/CPU-Apple%20M2(4P+4E)-brightgreen?style=flat-square&logo=Apple) ![RAM:16GB](https://img.shields.io/badge/RAM-16GB-brightgreen?style=flat-square) ![Disk:512GB](https://img.shields.io/badge/Disk-512GB-brightgreen?style=flat-square) ![LAN:1000M](https://img.shields.io/badge/LAN-1000M-brightgreen?style=flat-square) ![WiFi:802.11ac](https://img.shields.io/badge/WiFi-802.11ac-brightgreen?style=flat-square)


- [官方设备规格链接](https://support.apple.com/kb/SP870?locale=zh_CN)
- 相关文章记录
    - 《[在搭载 M1 及 M2 芯片 MacBook设备上玩 Stable Diffusion 模型](https://soulteary.com/2022/12/10/play-the-stable-diffusion-model-on-macbook-devices-with-m1-and-m2-chips.html)》


👉 [完整说明](./handbook/items-in-use/computing.md) / **简要说明** 👇


## 🚚 储存资源

> 用来持久化保存资料，开始服务时间从作为存储角色开始计算。[历史设备见文档](./handbook/devices/device-list.md)。

| 资源类型 | 明细 | 网络 | 储存 | 开始服务 |
| --- | --- | --- | --- | --- |
| 主力备份 | Synology DS 920+ ^1 | 千兆LAN | 17TB (8TB Raid1 / 8TB SHR / 1TB SSD) + 0.5TB SSD Cache | 2020 |
| 资源冷备 | 硬盘若干 ^2 | - | - | 2016 / 2018 / 2020 |
| 长期备份 | Canon G3800 ^3 | 2G WIFI | -      | 2019 |
| 清理备份 | Deli 9920 碎纸机 ^4 | - | -      | 2017 |
| 电力保障 | APC BR550G ^5 | - | - | 2017/2019 |
| 电力保障 | APC BK650M2 ^6 | - | - | 2023 |

👉 [完整说明](./handbook/items-in-use/storage.md) / **简要说明** 👇 

- [1] 新品发布时入手的 DS920+ 四盘位机器，取代之前服役了许久的 [Synology DS 718+](./handbook/notes/2021-ds718-plus-hard-drive-replacement-record.md) 和定制的 [HP Gen8 MicroServer](./handbook/devices/HP-Gen8-MicroServer.md)。使用3组盘来区分对待不同场景的数据，针对临时下载数据，使用 SSD 进行数据落地，对于诸如软件资源等持久性存储的一般数据使用 SHR 模式的磁盘存储，而对于宝贵的照片数据则使用 Raid 1 进行储存，并搭配 SSD Cache 对重复查询的数据进行缓存。
- [2] 因为各种原因腾出来的闲置磁盘。
- [3] 打印不失为一种相对稳定的持久化保存方案，之前因为放置太久坏掉过一台，迫于搬家后打印需求变多，又买了一台。价格便宜，非常好用。
- [4] 干掉持久化的纸质存储，最靠谱的莫过于加密级别的粉碎了，尤其是相对隐私敏感的内容。
- [5] 在所有电源都带稳流稳压作用后，添加一台UPS可以进一步防止市电闪断带来的数据损坏或者写输出脏掉的问题。在第一块使用了两年后，更换了一块电池，继续战斗，产品质量靠谱。
- [6] 为数据备份设备单独准备的后备电源。

- 《[硬件笔记：组装“固态 U 盘”的八年，从 100 块到 1000 块](https://soulteary.com/2023/09/16/eight-years-of-assembling-solid-state-usb-disk-from-100-to-1000.html)》


## 📱 移动设备 & 🎮 游戏设备

> 强依赖网络进行交互的娱乐设备。[历史设备见文档](./handbook/devices/device-list.md)。

| 编号 | 资源类型 | 明细 | 网络 | 储存 | 开始服务 |
| --- | --- | --- | --- | --- | --- |
| 1 | 游戏机 | Switch 续航版 | 5G WIFI | 500G | 2020 |
| 2 | 游戏机 | Switch Lite  | 5G WIFI | 500G | 2020 |
| 3 | 游戏机 | PS4 | 2G WIFI | 500G HDD | 2017 |
| 4 | 游戏机 | PS4 Pro | 2G WIFI | 500G SSD | 2017 |
| 5 | 游戏机 | PSVx2 ^2 | 2G WIFI | 16G / 64G | 2015 / 2016 |
| 6 | 游戏机 | 3DSx2 ^3 | 2G WIFI | 64G / 64G | 2014 / 2016 |
| 7 | 平板 | iPad Air2 | 4G / 5G WIFI | 128G (改) | 2015 |
| 8 | 平板 | iPad Pro 10' ^4 | 5G WIFI | 256G | 2018 |
| 9 | 平板 | iPad Pro 12' ^4 | 5G WIFI | 256G | 2018 |
| 10 | 爪机 | iPhone 14 Pro | 5G / 5G WIFI | 512G  | 2022 |
| 11 | 爪机 | Redmi 11T Pro | 5G / 5G WIFI | 256G  | 2022 |
| 12 | 爪机 | 海信 A7cc | 5G / 5G WIFI | 128G  | 2021 |
| 13 | 爪机 | iP3GS ^1 | 2G / 2G WIFI | 忽略 | 2017 |

👉 [完整说明](./handbook/items-in-use/entertainment.md) / **简要说明** 👇

- [1] 出掉了17年购置的初代后，将状态变更为一正一破，好处是可以联机的游戏，可以大号带小号玩（比如动森），另外可以做PS4手柄，玩不需要 L2R2 键的游戏体验还可以。
- [2] 高铁候车、飞机候机、团建出门、旅游出行必备，小巧可爱，有 GBA XL 的感觉。
- [3] 家中吃灰。
- [4] 巫师3专用机。
- [5] 灵魂献祭专用机。
- [6] 塞尔达、逆转、火纹专用机。
- [7] 电子相册专用机。
- [8] 电子笔记 + 电脑副屏 + 王者荣耀专用机。
- [9] 电子笔记 + 电脑副屏 + 王者荣耀专用机。
- [10] 12Pro 的机器感觉有些卡，恰逢新版本发布，正好升级。
- [11] 接替 K30Pro，主要作为导航机器使用。
- [12] 微信读书阅读器，体验很棒。
- [13] 年度最值手机，作为monitor使用，极低功耗，可以愉快跑脚本，已购两台。


## 🔮 智能设备 & 周边

待更新

> 相比较前些年的智能设备，这些年的设备的体验越来越好了。[历史设备见文档](./handbook/devices/device-list.md)。

| 编号 | 资源类型 | 明细 | 网络 |备注 |
| --- | --- | --- | --- | --- |
| 1 | 网络音箱  | 小米音箱 Pro | 2G WIFI | 2019年  |
| 2 | 蓝牙音箱 | 飞利浦 TAVS700 蓝牙音响 | BT | 2021年  |
| 3 | 空气净化器  | 小米净化器 Pro | 2G WIFI | 2019年  |
| 4 | 网络网关  | 小米 x3 | 2G WIFI | 2017年 / 2021年 |
| 5 | 网络插座  | 小米 xN | 2G WIFI | 2017年 / 2021年 |
| 6 | 传感装置  | 小米 xN | 2G WIFI/ZigBee | 2017年 / 2021年 |
| 7 | 网络摄像头 | 水滴 x3 | 2G WIFI | 2015年、2016年、2017年、2018年 |
| 8 | 网络摄像头 | 小方等 x2 | 2G WIFI | 2018年 |
| 9 | 网络盒子 | 小米盒子4SPro | WIFI6 | 2021年 |
| 10 | 网络盒子 | AppleTV 6 | WIFI6 | 2021年 |


## 简要说明

👉 [完整说明](./handbook/items-in-use/iot.md) / **简要说明** 👇 

- [1] 音质尚可，在尝试过两台组网进行全屋播放后，最终还是让一台推出的服务的舞台，剩余一台作为“小爱同学”使用。
- [2] 因为AppleTV和小爱协作音响召唤“小爱同学”，只好换了一个简单的设备。
- [3] 拯救雾霾天和猫主子入厕后的空气质量，“手动档”挺好使，其他自动档位比较“弱鸡”，正常运行基本静音，总体点赞。
- [4] 小米网关+空调插座，对于空间利用率很高，而且支持ZigBee/WiFi，还能网络调试。
- [5] 小米开放局域网协议之后，把控客都换成了小米，支持编程这件事太好了。
- [6] 门磁可以避免出门老想着有没有关好门的问题，烟雾等传感器避免检查厨房，很省心，光照传感器可以制作围绕床帏的夜灯等，简单实用。
- [7] 2015年从小米换到360就一直使用，目前使用还好，小问题是启动音有点惊悚，4台设备经常会有一台显示离线，我所在的小区，想稳定使用需要额外提供移动热点，目前考虑替换掉。
- [8] 小方支持自定义固件，折腾通过，计划对接群晖后替换所有360摄像头。
    - [-]控客的插座APP不是很好用，尤其是不给SDK，无法定制开发，国内版本也不支持简单DIY（需要拆+编程器），故弃用。
- [9] 默认广告有一些多，而且遇到有趣的视频想看弹幕不能看挺遗憾的，乐播 SDK 投屏质量感觉不是很好，而且有强制广告，但是视频资源还是挺多的。
- [10] 最强投屏盒子，支持AirPlay2，无线投屏完美使用4K。
