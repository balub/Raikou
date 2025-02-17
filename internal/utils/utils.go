package utils

import (
	"fmt"
	"os"
)

func GetDeviceHomeDir() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error: ", err)
	}

	return homeDir
}

func GetSSHDir() string {
	homeDir := GetDeviceHomeDir()
	sshDir := fmt.Sprintf("%s/.ssh", homeDir)

	return sshDir
}

func GetSSHConfigPath() string {
	sshDir := GetSSHDir()
	sshConfigPath := fmt.Sprintf("%s/mock-config", sshDir)

	return sshConfigPath
}
