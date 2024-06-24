package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"github.com/mitchellh/go-homedir"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

type Cli struct {
	user       string
	pwd        string
	keyPath    string
	ip         string
	port       string
	sshClient  *ssh.Client
	sftpClient *sftp.Client
}

func NewSSHClient(user, keyPath, pwd, ip, port string) Cli {
	SshClient = Cli{
		user:    user,
		pwd:     pwd,
		ip:      ip,
		port:    port,
		keyPath: keyPath,
	}
	return SshClient
}

// 不使用 HostKey， 使用密码
func (c *Cli) getConfig_nokey() *ssh.ClientConfig {
	// sshKeyPath := "/Users/darren/.ssh/id_rsa.zzygmail"

	config := &ssh.ClientConfig{
		User: c.user,
		Auth: []ssh.AuthMethod{publicKeyAuthFunc(c.keyPath)},
		// Auth: []ssh.AuthMethod{
		// 	ssh.Password(c.pwd),
		// },
		Timeout:         30 * time.Second,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	return config
}

func publicKeyAuthFunc(kPath string) ssh.AuthMethod {
	keyPath, err := homedir.Expand(kPath)
	if err != nil {
		log.Fatal("find key's home dir failed", err)
	}
	key, err := ioutil.ReadFile(keyPath)
	if err != nil {
		log.Fatal("ssh key file read failed", err)
	}
	// Create the Signer for this private key.
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		log.Fatal("ssh key signer failed", err)
	}
	return ssh.PublicKeys(signer)
}
func (c *Cli) Connect() error {
	config := c.getConfig_nokey()

	client, err := ssh.Dial("tcp", c.ip+":"+c.port, config)
	if err != nil {
		return fmt.Errorf("connect server error: %w", err)
	}
	sftps, err := sftp.NewClient(client)
	if err != nil {
		return fmt.Errorf("new sftp client error: %w", err)
	}

	c.sshClient = client
	c.sftpClient = sftps
	return nil
}

func (c Cli) Run(cmd string) (string, error) {

	if c.sshClient == nil {
		if err := c.Connect(); err != nil {
			return "", err
		}
	}

	session, err := c.sshClient.NewSession()
	if err != nil {
		return "", fmt.Errorf("create new session error: %w", err)
	}
	defer session.Close()

	buf, err := session.CombinedOutput(cmd)
	return string(buf), err
}

func (c Cli) DownloadFile(remoteFile, localFile string) (int, error) {
	if c.sshClient == nil {
		if err := c.Connect(); err != nil {
			return -1, err
		}
	}
	source, err := c.sftpClient.Open(remoteFile)
	if err != nil {
		return -1, fmt.Errorf("sftp client open file error: %w", err)
	}
	defer source.Close()

	target, err := os.OpenFile(localFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return -1, fmt.Errorf("open local file error: %w", err)
	}
	defer target.Close()

	n, err := io.Copy(target, source)
	if err != nil {
		return -1, fmt.Errorf("write file error: %w", err)
	}
	return int(n), nil
}

var SshClient Cli

func findFiles(folder string) []string {
	// client := NewSSHClient(username, password, ip, port)
	// SshClient = client
	// 1.运行远程命令
	cmd := `find ` + folder + `  -type f \( -name "*.iyaml" -o -name "*.yaml" \)`
	backinfo, err := SshClient.Run(cmd)
	if err != nil {
		fmt.Printf("failed to run shell,err=[%v]\n", err)
		return nil
	}
	fmt.Printf("%v back info: \n[%v]\n", cmd, backinfo)
	list := strings.Split(backinfo, "\n")
	return list
}

type FileCfg struct {
	Name  string
	Value string
}

func parseFile(file string, tag []string) []FileCfg {
	// SshClient.DownloadFile(file)

	return nil
}

func save2Db(accountName string, hostName string, folder string, data []FileCfg) {

}

func mainss() {
	username := "zhangzhy"
	keyPath := "/Users/darren/.ssh/id_rsa.zzy-m1-hj"
	ip := "earth"
	port := "22"
	client := NewSSHClient(username, keyPath, "", ip, port)
	cmd := `ls`
	backinfo, err := client.Run(cmd)
	if err != nil {
		fmt.Printf("failed to run shell,err=[%v]\n", err)
		return
	}
	fmt.Printf("%v back info: \n[%v]\n", cmd, backinfo)
	// // 2. 上传一文件
	filename := "Foo.txt"
	// WriteFile(filename, []byte("hello ssh\r\n"))
	// // 上传
	// n, err := client.UploadFile(filename, "/tmp/"+filename)
	// if err != nil {
	// 	fmt.Printf("upload failed: %v\n", err)
	// 	return
	// }
	// 3. 显示该文件
	// cmd = "cat " + "/tmp/" + filename
	// backinfo, err = client.Run(cmd)
	// if err != nil {
	// 	fmt.Printf("run cmd faild: %v\n", err)
	// 	return
	// }
	// fmt.Printf("%v back info: \n[%v]\n", cmd, backinfo)
	// 4. 下载该文件到本地
	n, err := client.DownloadFile("/tmp/"+filename, "fo2o.txt")
	if err != nil {
		fmt.Printf("download failed: %v\n", err)
		return
	}
	fmt.Printf("download file[%v] ok, size=[%d]\n", filename, n)
}
