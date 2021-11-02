package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/charoleizer/disk-info/core"
	"github.com/charoleizer/disk-info/utils"
)

func GetDiskUsage() string {
	disk := core.DiskUsage("/")

	disk.Total = utils.RoundTo(disk.Total/GB, 2)
	disk.Used = utils.RoundTo(disk.Used/GB, 2)
	disk.Available = utils.RoundTo(disk.Available/GB, 2)
	disk.Used_Percent = utils.RoundTo(disk.Used_Percent, 2)
	disk.Available_Percent = utils.RoundTo(disk.Available_Percent, 2)

	response, _ := json.Marshal(disk)

	return string(response)

}

func InfiniteLoop() {
	var info map[string]interface{}

	for {
		json.Unmarshal([]byte(GetDiskUsage()), &info)

		if info["available-percent"].(float64) >= 90 {
			fmt.Println("Alerta - Mais de 90% do disco está em uso")
		}

		time.Sleep(time.Second)

	}
}

func main() {
	fmt.Println("Monitoramento de disco iniciado")
	fmt.Println("Quando este servidor atingir 90% da capacidade de disco, você será notificado.")
	InfiniteLoop()
}

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)
