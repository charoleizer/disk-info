package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/charoleizer/disk-info/core"
	"github.com/charoleizer/disk-info/notify/telegram"
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

		if info["available-percent"].(float64) <= 10 {
			telegram.Notify("⚠️ Alerta - Apenas 10% do espaço em disco etá disponível.")
		}

		time.Sleep(time.Minute)

	}
}

func main() {
	fmt.Println("Monitoramento de disco iniciado")
	InfiniteLoop()
}

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)
