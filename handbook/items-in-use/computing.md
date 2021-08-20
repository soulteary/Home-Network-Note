# 💻 主机资源

> 提供运算能力的本地设备，[历史设备见文档](../devices/README.md)。

| 序号 | 资源类型 | 明细 | 网络 | 储存 | 开始服务 |
| --- | --- | --- | --- | --- | --- |
| 1 | 编码机器 | MacBook Pro 16 | 千兆LAN & 5G Wi-Fi | 32GBRAM / 2T  | 2019 |
| 2 | 资源机器 | ThinkBook 15 | 千兆LAN & 5G Wi-Fi | 40GBRAM / 2T  | 2021 |
| 3 | 便携机器 | ThinkPad L14 | 千兆LAN & 5G Wi-Fi | 64GBRAM / 1T  | 2021 |
| 4 | 资源机器 | Nuc8i5BEH | 千兆LAN / 5G Wi-Fi | 64GBRAM / 2T  | 2021 |
| 5 | 网络设备 | Nuc7CJYH | 千兆LAN | 8GBRAM / 256GB | 2021 |
| x | 码字机器 | MacBook Pro m1 | 千兆LAN & 5G Wi-Fi | 8GBRAM / 512GB | 2021 |

## 简要说明

- [1] 在 19 年首发入手，这款键盘相比较之前有巨大改善，缓解了养宠物的我需要不定时访问“苹果售后”清理键盘中的猫毛的问题，性能也让我非常满意，比公司配备的16寸i7 2.6GHz设备性能足足高了 20～30%。
- [2] 无独显版本，搭载 7nm Zen3 5800u，性能彪悍，作为补充资源机器购置。
- [3] 无独显版本，搭载 7nm Zen2 4750u 仅 45w 峰值功耗，性能非常强，核心数也非常多，作为一台便携的“服务器”使用，用于拓展本地开发资源，提供一个冗余一些的资源跑测试服务，替换之前使用的 [HP EliteDesk G4 800](./handbook/devices/HP-EliteDesk-G4-800.md) 小型工作站。
- [4] 入手原因见[这篇文章](https://soulteary.com/2021/01/31/nuc-notes-linux-system.html)。在随后不断添置和更新设备后，这台设备职能更新为软件测试资源，提供搭建各种开源软件分布式环境场景所使用的虚拟机环境，极大的降低了笔记本发热的程度。
- [5] 将群晖上的容器服务迁移至此，解决群晖跑容器，硬盘不会休眠的问题。同时提供稳定的“回家网络”、部分个人公开 Wiki 资源。
- [-] 13-inch, m1 吃螃蟹，性能非常强，逼近 i9 的设备，妹纸打字机，测试应用构建时使用。

## 为什么主力是笔记本而非台式机

下面有两篇文章，系统的讲解了为什么，以及如何做，感兴趣可以进行阅读。

- [廉价的家用工作站方案：前篇](https://soulteary.com/2021/07/02/cheap-home-workstation-solution-part-one.html)
- [AMD 4750u 及 5800u 笔记本安装 Ubuntu 20.04](https://soulteary.com/2021/07/04/install-ubuntu-20-04-on-amd-4750u-and-5800u-laptops.html)

使用上面方案的计算资源，再配合一些轻量的云主机，基本上什么实验和环境就都能轻松模拟啦。

## NUC 不跑黑苹果，可以用来干嘛

当然是做虚拟化，以及为群晖上的容器服务换个地方，让群晖可以睡个好觉，让硬盘可以活的更久一些。

- [NUC 折腾笔记 - 安装 ESXi 7](https://soulteary.com/2021/06/22/nuc-notes-install-esxi7.html)
- [NUC 折腾笔记 - Linux 系统篇](https://soulteary.com/2021/01/31/nuc-notes-linux-system.html)
- [NUC 折腾笔记 - 储存能力测试](https://soulteary.com/2021/02/02/nuc-notes-storage-ability-test.html)
