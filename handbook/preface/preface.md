# 写在前面

> 大家因为各种各样的原因入坑 `HomeLab`，希望我的一些记录，可以帮助此刻阅读的你，节约一些时间和经济成本。

---

## 项目说明

_**文档内容正在进行大幅度调整，如果你发现了问题，欢迎联系或反馈我，谢谢**_

记录搭建兼顾学习娱乐的家用网络环境的过程，折腾过的一些软硬件小经验。

目前的网络方案从 2016 年使用至今，非常稳定，整体架构几乎没有变化。（日常在线20~30台设备，峰值50+）

文档中的方案和方案中的配置会尽可能**保持简单**，确保各种服务在运行一年之后，我依旧能够对软硬件进行**轻松简单**的“维护、升级以及替换”操作。

## 主要场景和关键词

列举常见家用网络场景的一些核心诉求的关键词。

|     | 设备联网 | 备份数据 | 下载上传 | 数据同步 | 开发学习 | 游戏娱乐 |
| --- | --- | --- | --- | --- | --- | --- |
| **核心指标** | 安全、稳定   | 安全、可靠 | 高速 | 无感知 |  流畅  | 流畅 |
| **重要因素** | 简单、易维护  | 高效     | 易用 |  准确  |  省心  | 舒适 |
| **可选因素** | 网速、组网模式 | 易用     | 安全 | 跨平台 | 冗余保障 | |

## 我个人的一些习惯

## 保持工作习惯，保持熟悉感

> 尽可能利用系统优势，但是保证一定的方案可替代性

- 一般事务使用 `Mac OSX`，MBP不更换系统。
- 持续编码使用 `Mac OSX`，MBP系统底层环境尽可能干净，使用内部、外部虚拟机、Docker保证环境隔离，发布测试集中构建使用虚拟化方案，保证环境独立和结果稳定。
- 游戏娱乐使用 `iOS`/`Windows`/`游戏机`，在遇到问题的时候，可以获得最广泛的资料参考和软件支持，减少时间浪费。

## 跨设备开发的可能性，保持一定的灵活性

- 可以接受牺牲一定的性能，或者在某种场景下不是最优方案，但是要保障可一定程度的迁移灵活性。
- 项目构建发布流程期望通用，且可以方便移植到其他系统的机器上，需要使用虚拟化方式实现。
  - 虚拟化技术方案：虚拟机（整体服务） && Docker（单一应用）
  - 虚拟化系统选择：非 `Win` 和 `macOS` 之外，系统基础镜像选择 `Ubuntu` / `Alpine`

## 娱乐体验

- 折腾也是乐趣，折腾的东西要能够有明确目标。
- 目前不再使用多服务商网络聚合
    - 达不到稳定无感知，时不时花时间维护一下，没什么意义，不如选择一家服务商，使用相对高品质的服务。
- 游戏使用游戏终端来进行，游戏体验更好，也避免了工作、学习设备中途可能被其他事物打断游戏进程。
- 智能插座、监控使用不同的路由设备，限制对内网进行访问，以免出现漏洞被恶意利用。

## 其他记录

- [博客中的折腾笔记](https://soulteary.com/subject/life/)
