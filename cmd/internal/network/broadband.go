package network

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

type BroadBand struct {
	TLDR string `yaml:"TLDR"`
	Full string `yaml:"Full"`
	List []struct {
		Name string `yaml:"Name"`
		Type string `yaml:"Type"`
		Note string `yaml:"Note"`
	} `yaml:"List"`
	Description string
}

func GetBoardBand(baseDir string) (broadband BroadBand, err error) {
	buf, err := os.ReadFile(filepath.Join(baseDir, "broadband", "info.yml"))
	if err != nil {
		fmt.Println("read broadband yaml error: ", err)
		return broadband, err
	}

	err = yaml.Unmarshal(buf, &broadband)
	if err != nil {
		fmt.Println("unmarshal broadband yaml error: ", err)
		return broadband, err
	}

	buf, err = os.ReadFile(filepath.Join(baseDir, "broadband", "info.md"))
	if err != nil {
		fmt.Println("read broadband md error: ", err)
		return broadband, err
	}

	broadband.Description = string(buf)

	return broadband, nil
}

func MakeBroadBandTemplate(broadband BroadBand) string {
	devicesIntro := []string{}

	if broadband.TLDR != "" {
		devicesIntro = append(devicesIntro, fmt.Sprintf("> %s", broadband.TLDR), "")
	}

	if len(broadband.List) > 0 {
		devicesIntro = append(devicesIntro, "| 资源类型 | 明细 | 备注 |")
		devicesIntro = append(devicesIntro, "| --- | --- | --- |")
		for _, item := range broadband.List {
			devicesIntro = append(devicesIntro, fmt.Sprintf("| %s | %s | %s |", item.Name, item.Type, item.Note))
		}
		devicesIntro = append(devicesIntro, "")
	}

	if broadband.Full != "" {
		devicesIntro = append(devicesIntro, broadband.Full, "")
	}

	devicesIntro = append(devicesIntro, broadband.Description)
	return strings.Join(devicesIntro, "\n")
}

func GetBroadbandNetworkMarkdown(baseDir string) string {
	boardband, err := GetBoardBand(baseDir)
	if err != nil {
		log.Fatalln("get boardband error: ", err)
	}
	return MakeBroadBandTemplate(boardband)
}

func UpdateBroadbandNetworkReadme(baseDir string) {
	buf, err := os.ReadFile(filepath.Join(baseDir, "broadband", "full.md"))
	if err != nil {
		log.Fatalln("read broadband markdown error: ", err)
	}

	boardband, err := GetBoardBand(baseDir)
	if err != nil {
		log.Fatalln("get boardband error: ", err)
	}
	boardband.Full = ""

	result := strings.ReplaceAll(string(buf), "${BROADBAND_NETWORK}", MakeBroadBandTemplate(boardband))
	os.WriteFile(filepath.Join(baseDir, "broadband.md"), []byte(result), 0644)
}
