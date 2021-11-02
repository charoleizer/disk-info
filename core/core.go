package core

import (
	"syscall"

	"github.com/charoleizer/disk-info/models"
)

func DiskUsage(path string) (disk models.DiskStatus) {
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

const (
	MAXPERCENT = 100
)
