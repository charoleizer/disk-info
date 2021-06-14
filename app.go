package main

import (
	"encoding/json"
	"fmt"
	"math"
	"syscall"
)

type DiskStatus struct {
	Total             float64 `json:"total"`
	Used              float64 `json:"used"`
	Available         float64 `json:"available"`
	Used_Percent      float64 `json:"used-percent"`
	Available_Percent float64 `json:"available-percent"`
}

func DiskUsage(path string) (disk DiskStatus) {
	fs := syscall.Statfs_t{}
	err := syscall.Statfs(path, &fs)
	if err != nil {
		return
	}
	disk.Total = float64(fs.Blocks) * float64(fs.Bsize)
	disk.Available = float64(fs.Bfree) * float64(fs.Bsize)
	disk.Used = disk.Total - disk.Available

	disk.Used_Percent = float64(disk.Used) * MAXPERCENT / float64(disk.Total)
	disk.Available_Percent = float64(disk.Available) * MAXPERCENT / float64(disk.Total)

	return
}

func Round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func RoundTo(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(Round(num*output)) / output
}

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB

	MAXPERCENT = 100
)

func GetDiskUsage() string {
	disk := DiskUsage("/")

	disk.Total = RoundTo(disk.Total/GB, 2)
	disk.Used = RoundTo(disk.Used/GB, 2)
	disk.Available = RoundTo(disk.Available/GB, 2)
	disk.Used_Percent = RoundTo(disk.Used_Percent, 2)
	disk.Available_Percent = RoundTo(disk.Available_Percent, 2)

	response, _ := json.Marshal(disk)

	return string(response)

}

func main() {
	fmt.Println(GetDiskUsage())

}
