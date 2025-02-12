package ssh

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/raikou/internal/utils"
)

type SSHosts struct {
	Host         string
	HostName     string
	User         string
	Port         int
	IdentityFile string
}

func GetSSHConfigFile() (*string, error) {
	configFilePath := utils.GetSSHConfigPath()
	data, err := os.ReadFile(configFilePath)
	if err != nil {
		return nil, fmt.Errorf("Error reading ssh config file: %v", err)
	}

	result := string(data)
	return &result, nil
}

func ParseSSHConfigFile() (*[]SSHosts, error) {
	configFile, err := GetSSHConfigFile()
	if err != nil {
		return nil, err
	}

	configFileWithoutComments := removeComments(*configFile)
	hostsList := GetHostsList(configFileWithoutComments)

	var sshHosts []SSHosts
	for _, host := range hostsList {
		if len(host) == 0 {
			continue
		}
		sshHosts = append(sshHosts, ParseHostProps(host))
	}

	return &sshHosts, nil
}

func ParseHostProps(host string) SSHosts {
	var sshHost SSHosts
	lines := strings.Split(host, "\n")
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		trimmedLine := strings.TrimSpace(line)

		// Host
		if len(strings.Split(trimmedLine, " ")) == 1 {
			sshHost.Host = trimmedLine
		}

		// HostName
		if strings.Contains(trimmedLine, "HostName") {
			sshHost.HostName = strings.Split(trimmedLine, "HostName ")[1]
		}

		// User
		if strings.Contains(trimmedLine, "User") {
			sshHost.User = strings.Split(trimmedLine, "User ")[1]
		}

		// Port
		if strings.Contains(trimmedLine, "Port") {
			portNum, _ := strconv.Atoi(strings.Split(trimmedLine, "Port ")[1])
			sshHost.Port = portNum
		}

		// IdentityFile
		if strings.Contains(trimmedLine, "IdentityFile") {
			sshHost.IdentityFile = strings.Split(trimmedLine, "IdentityFile ")[1]
		}
	}

	return sshHost

}

func GetHostsList(fileString string) []string {
	var data []string
	hosts := strings.Split(fileString, "Host ")
	data = append(data, hosts...)
	return data
}

func removeComments(configFile string) string {
	var configFileWithoutComments string

	lines := strings.Split(configFile, "\n")
	for _, line := range lines {
		if strings.Contains(line, "#") {
			continue
		}
		configFileWithoutComments += line + "\n"
	}

	return configFileWithoutComments
}

func Print() {
	hosts, err := ParseSSHConfigFile()
	if err != nil {
		panic(err)
	}

	var tableBuilder strings.Builder

	// Build table header
	tableBuilder.WriteString("| Host | HostName | User | Port | IdentityFile |\n")
	tableBuilder.WriteString("|------|----------|------|------|--------------|\n")

	// Build each host row in the table
	for _, host := range *hosts {
		tableBuilder.WriteString(fmt.Sprintf("| %s | %s | %s | %d | %s |\n", host.Host, host.HostName, host.User, host.Port, host.IdentityFile))
	}

	fmt.Println(tableBuilder.String())
}
