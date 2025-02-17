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

func CreateSSHCommand(host string, hostname string, user string, port int, identifyFile string) string {
	// ssh -i ~/.ssh/bastion_key -p 22 bastionuser@bastion.example.com
	return fmt.Sprintf("ssh -i %s -p %s %s@%s", identifyFile, port, user, hostname)
}
