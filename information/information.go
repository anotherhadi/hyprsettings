package information

import (
	"fmt"
	"os/exec"
	"strings"

	hyprsettings_utils "github.com/anotherhadi/hyprsettings/utils"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

func getProcessorName() (string, error) {
	info, err := cpu.Info()
	if err != nil {
		return "", err
	}
	return info[0].ModelName, nil
}

func getGPUName() (string, error) {
	cmd := exec.Command("lspci", "-v")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if strings.Contains(line, "VGA compatible controller") {
			parts := strings.Split(line, ":")
			return strings.TrimSpace(parts[len(parts)-1]), nil
		}
	}
	return "", nil
}

func getDiskSize() (float64, error) {
	partitions, err := disk.Partitions(true)
	if err != nil {
		return 0, err
	}
	var totalSize uint64
	for _, partition := range partitions {
		usage, err := disk.Usage(partition.Mountpoint)
		if err == nil {
			totalSize += usage.Total
		}
	}
	return float64(totalSize) / (1024 * 1024 * 1024), nil
}

func getRAMSize() (float64, error) {
	info, err := mem.VirtualMemory()
	if err != nil {
		return 0, err
	}
	return float64(info.Total) / (1024 * 1024 * 1024), nil
}

func getUsername() (string, error) {
	info, err := host.Info()
	if err != nil {
		return "", err
	}
	return info.Hostname, nil
}
func Information() hyprsettings_utils.Page {
	processorName, err := getProcessorName()
	if err != nil {
		panic(err)
	}
	gpuName, err := getGPUName()
	if err != nil {
		panic(err)
	}
	diskSize, err := getDiskSize()
	if err != nil {
		panic(err)
	}
	ramSize, err := getRAMSize()
	if err != nil {
		panic(err)
	}
	username, err := getUsername()
	if err != nil {
		panic(err)
	}
	return hyprsettings_utils.Page{
		Title:       "Informations",
		Description: "System & Hyprsettings informations",
		PageType:    "info",
		Setting: hyprsettings_utils.Setting{
			Info: hyprsettings_utils.Info{
				Paragraph: []string{
					"System informations:",
					"Username: " + username,
					"CPU: " + processorName,
					"GPU: " + gpuName,
					fmt.Sprintf("Ram Size: %.2fGb", ramSize),
					fmt.Sprintf("Disk Size: %.2fGb", diskSize),
					"",
					"Hyprsettings by @anotherhadi",
				},
			},
		},
	}
}
