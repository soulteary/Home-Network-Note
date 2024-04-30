之前一直在寻找核心数多、功耗低、支持 64GB 内存、相对便携的无显卡笔记本设备，以做“廉价的服务器”使用，并替代早先时候购置的 [HP EliteDesk G4 800](./handbook/devices/HP-EliteDesk-G4-800.md) 小型工作站，直到遇到了搭载 AMD Zen2 4750u 的 ThinkPad L14。

这台设备满载仅 45w 的功耗，性能极强，核心数也非常多，特别适合长时间跑容器服务，来扩展本地的计算资源。美中不足的是，设备只有一条固态硬盘插槽可用。

2022年，设备价格进一步下降，增加了一台相同 CPU 配置的新设备作为冗余资源（32G/1TB）。

- [官方设备规格链接](https://psref.lenovo.com/syspool/Sys/PDF/ThinkPad/ThinkPad_L14_Gen_1_AMD/ThinkPad_L14_Gen_1_AMD_Spec.pdf)
- 相关文章记录
    - 《[廉价的家用工作站方案：前篇](https://soulteary.com/2021/07/02/cheap-home-workstation-solution-part-one.html)》
    - 《[AMD 4750u 及 5800u 笔记本安装 Ubuntu 20.04](https://soulteary.com/2021/07/04/install-ubuntu-20-04-on-amd-4750u-and-5800u-laptops.html)》
- 同宗同源，但已出掉的 5800u (ThinkBook 15)
    - 《[AMD 5800u 笔记本折腾 Proxmox VE 7.0 虚拟化](https://soulteary.com/2021/10/23/amd-5800u-notebook-toss-proxmox-ve-7-0-virtualization.html)》、《[装在笔记本里的私有云环境：准备篇](https://soulteary.com/2021/10/28/private-cloud-environment-installed-in-a-notebook-preparation.html)》、《[装在笔记本里的私有云环境：监控篇](https://soulteary.com/2021/10/30/private-cloud-environment-installed-in-a-notebook-monitor.html)》、《[装在笔记本里的私有云环境：网络存储篇（上）](https://soulteary.com/2021/11/04/private-cloud-environment-installed-in-a-notebook-storage-part-1.html)》、 《[装在笔记本里的私有云环境：网络存储篇（中）](https://soulteary.com/2021/11/09/private-cloud-environment-installed-in-a-notebook-storage-part-2.html)》、 《[装在笔记本里的私有云环境：K8s 集群准备](https://soulteary.com/2022/11/29/private-cloud-environment-installed-in-a-notebook-k8s-cluster-preparation.html)》
