package network

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

type Gateway struct {
	TLDR string `yaml:"TLDR"`
	Full string `yaml:"Full"`
	List []struct {
		Name    string `yaml:"Name"`
		Type    string `yaml:"Type"`
		Note    string `yaml:"Note"`
		Date    string `yaml:"Date"`
		Storage string `yaml:"Storage,omitempty"`
	} `yaml:"List"`
	Description string
}

func GetGateway(baseDir string) (gateway Gateway, err error) {
	buf, err := os.ReadFile(filepath.Join(baseDir, "gateway", "info.yml"))
	if err != nil {
		fmt.Println("read gateway yaml error: ", err)
		return gateway, err
	}

	err = yaml.Unmarshal(buf, &gateway)
	if err != nil {
		fmt.Println("unmarshal gateway yaml error: ", err)
		return gateway, err
	}

	buf, err = os.ReadFile(filepath.Join(baseDir, "gateway", "info.md"))
	if err != nil {
		fmt.Println("read gateway md error: ", err)
		return gateway, err
	}

	gateway.Description = string(buf)

	return gateway, nil
}

func MakeGatewayTemplate(gateway Gateway) string {
	devicesIntro := []string{}

	if gateway.TLDR != "" {
		devicesIntro = append(devicesIntro, fmt.Sprintf("> %s", gateway.TLDR), "")
	}

	if len(gateway.List) > 0 {
		devicesIntro = append(devicesIntro, "| 资源类型 | 明细 | 网络 | 开始服务 |")
		devicesIntro = append(devicesIntro, "| --- | --- | --- | --- |")
		for _, item := range gateway.List {
			if item.Storage == "" {
				item.Storage = "-"
			}
			devicesIntro = append(devicesIntro, fmt.Sprintf("| %s | %s | %s | %s |", item.Type, item.Name, item.Note, item.Date))
		}
		devicesIntro = append(devicesIntro, "")
	}

	if gateway.Full != "" {
		devicesIntro = append(devicesIntro, gateway.Full, "")
	}

	devicesIntro = append(devicesIntro, gateway.Description)
	return strings.Join(devicesIntro, "\n")
}

func GetGatewayNetworkMarkdown(baseDir string) string {
	gateway, err := GetGateway(baseDir)
	if err != nil {
		log.Fatalln("get gateway error: ", err)
	}
	return MakeGatewayTemplate(gateway)
}

func UpdateGatewayNetworkReadme(baseDir string) {
	buf, err := os.ReadFile(filepath.Join(baseDir, "gateway", "full.md"))
	if err != nil {
		log.Fatalln("read gateway markdown error: ", err)
	}

	gateway, err := GetGateway(baseDir)
	if err != nil {
		log.Fatalln("get gateway error: ", err)
	}
	gateway.Full = ""

	result := strings.ReplaceAll(string(buf), "${GATEWAY_NETWORK}", MakeGatewayTemplate(gateway))
	os.WriteFile(filepath.Join(baseDir, "gateway.md"), []byte(result), 0644)
}
