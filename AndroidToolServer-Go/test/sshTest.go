package main

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"log"
	"os"
)

type SshConfig struct {
	User     string
	Password string
	Host     string
	Port     string
}

func main() {
	zhang := &SshConfig{
		User:     "root",
		Password: "123456",
		Host:     "192.168.200.130",
		Port:     "22",
	}

	cfg := &ssh.ClientConfig{
		User:            zhang.User,
		Auth:            []ssh.AuthMethod{ssh.Password(zhang.Password)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// 连接SSH服务器
	addr := fmt.Sprintf("%s:%s", zhang.Host, zhang.Port)
	client, err := ssh.Dial("tcp", addr, cfg)
	if err != nil {
		fmt.Println("连接虚拟机失败")
	}
	defer client.Close()
	// 创建一个会话
	session, err := client.NewSession()
	if err != nil {
		log.Fatalf("Failed to create session: %s", err)
	}
	defer session.Close()

	// 运行命令
	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	if err := session.Run("cd /home && ls"); err != nil {
		log.Fatalf("Failed to run: %s", err)
	}

}
