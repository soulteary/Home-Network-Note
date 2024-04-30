package devices

import (
	"fmt"
	"strings"
)

func MakeLabel(info string, infoType string) string {
	switch strings.ToLower(infoType) {
	case "system":
		osInfo := strings.Split(info, ":")
		osName := osInfo[0]
		brand := osName
		if strings.ToLower(brand) == "macos" {
			brand = "apple"
		}
		return fmt.Sprintf("![%s](https://img.shields.io/badge/%s-brightgreen?style=flat-square&logo=%s)", info, strings.ReplaceAll(strings.ReplaceAll(info, ":", "-"), " ", "%20"), brand)
	case "cpu":
		cpuInfo := strings.Split(info, ":")
		cpuName := cpuInfo[0]
		cpuSpeed := cpuInfo[1]
		cpuBrand := strings.Split(cpuName, " ")[0]
		return fmt.Sprintf("![CPU:%s](https://img.shields.io/badge/CPU-%s(%s)-brightgreen?style=flat-square&logo=%s)", cpuName, strings.ReplaceAll(strings.ReplaceAll(cpuName, " ", "%20"), "-", "--"), strings.ReplaceAll(cpuSpeed, " ", "%20"), cpuBrand)
	case "gpu":
		gpuInfo := strings.Split(info, ":")
		gpuName := gpuInfo[0]
		gpuExtra := gpuInfo[1]
		gpuBrand := strings.Split(gpuName, " ")[0]
		return fmt.Sprintf("![GPU:%s](https://img.shields.io/badge/GPU-%s-brightgreen?style=flat-square&logo=%s)", gpuName, strings.ReplaceAll(strings.ReplaceAll(gpuName, " ", "%20"), "-", "--")+"%20"+strings.ReplaceAll(gpuExtra, " ", "%20"), gpuBrand)
	case "ram":
		return fmt.Sprintf("![RAM:%s](https://img.shields.io/badge/RAM-%s-brightgreen?style=flat-square)", info, info)
	case "disk":
		return fmt.Sprintf("![Disk:%s](https://img.shields.io/badge/Disk-%s-brightgreen?style=flat-square)", info, info)
	case "lan":
		return fmt.Sprintf("![LAN:%s](https://img.shields.io/badge/LAN-%s-brightgreen?style=flat-square)", info, info)
	case "wifi":
		return fmt.Sprintf("![WiFi:%s](https://img.shields.io/badge/WiFi-%s-brightgreen?style=flat-square)", info, info)
	case "thunderbolt":
		return fmt.Sprintf("![ThunderBolt:%s](https://img.shields.io/badge/ThunderBolt-%s-brightgreen?style=flat-square)", info, info)
	}
	return ""
}

func MakeTemplate(device Device) (result string) {
	var draft []string
	title := fmt.Sprintf("### %s", device.Name)
	if device.TurnOff != "" {
		title = fmt.Sprintf("%s ![Turn On:%s](https://img.shields.io/badge/Turn%%20On-%s-brightgreen?style=flat-square)", title, device.TurnOn, device.TurnOn)
		title = fmt.Sprintf("%s ![Turn Off:%s](https://img.shields.io/badge/Turn%%20Off-%s-red?style=flat-square)", title, device.TurnOff, device.TurnOff)
	} else {
		title = fmt.Sprintf("%s ![Turn On:%s](https://img.shields.io/badge/Turn%%20On-%s-brightgreen?style=flat-square)", title, device.TurnOn, device.TurnOn)
	}
	draft = append(draft, title, "")

	tldr := fmt.Sprintf("> %s", device.TLDR)
	draft = append(draft, tldr, "")

	badges := []string{}
	if device.Info.Syetem != "" {
		badges = append(badges, MakeLabel(device.Info.Syetem, "system"))
	}
	if device.Info.CPU != "" {
		badges = append(badges, MakeLabel(device.Info.CPU, "cpu"))
	}
	if device.Info.GPU != "" {
		badges = append(badges, MakeLabel(device.Info.GPU, "gpu"))
	}
	if device.Info.RAM != "" {
		badges = append(badges, MakeLabel(device.Info.RAM, "ram"))
	}
	if device.Info.Disk != "" {
		badges = append(badges, MakeLabel(device.Info.Disk, "disk"))
	}
	if device.Info.LAN != "" {
		badges = append(badges, MakeLabel(device.Info.LAN, "lan"))
	}
	if device.Info.WiFi != "" {
		badges = append(badges, MakeLabel(device.Info.WiFi, "wifi"))
	}

	draft = append(draft, strings.Join(badges, " "), "")

	draft = append(draft, device.Description)

	return strings.Join(draft, "\n")
}

func GetDeviceMarkdown(baseDir string) string {
	devicesIntro := []string{}
	for _, device := range GetPowerOnList(baseDir) {
		devicesIntro = append(devicesIntro, MakeTemplate(device))
	}
	return strings.Join(devicesIntro, "\n")
}
