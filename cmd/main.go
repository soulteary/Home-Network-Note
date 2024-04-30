package main

import (
	"log"
	"os"
	"strings"

	"github.com/soulteary/home-network-note/internal/devices"
	"github.com/soulteary/home-network-note/internal/network"
)

func main() {
	template, err := os.ReadFile("./template.md")
	if err != nil {
		log.Fatalln("read template error: ", err)
	}

	devices.UpdatePowerOffDevice("../deprecate")
	devicesTemplate := devices.GetDeviceMarkdown("../power-on")

	newMarkdown := strings.ReplaceAll(string(template), "${COMPUTING_DEVICES}", devicesTemplate)

	broadbandTemplate := network.GetBroadbandNetworkMarkdown("../network")
	newMarkdown = strings.ReplaceAll(newMarkdown, "${BROADBAND_NETWORK}", broadbandTemplate)
	network.UpdateBroadbandNetworkReadme("../network")

	gatewayTemplate := network.GetGatewayNetworkMarkdown("../network")
	newMarkdown = strings.ReplaceAll(newMarkdown, "${GATEWAY_NETWORK}", gatewayTemplate)
	network.UpdateGatewayNetworkReadme("../network")

	os.WriteFile("../README.md", []byte(newMarkdown), 0644)
}
