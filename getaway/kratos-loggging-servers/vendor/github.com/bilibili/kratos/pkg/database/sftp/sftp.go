package sftp

import (
	xtime "github.com/bilibili/kratos/pkg/time"
	"github.com/pkg/sftp"
	"path"
	"time"
	"net"
	"os"
	"golang.org/x/crypto/ssh"
	"github.com/bilibili/kratos/pkg/log"
)

const (
	REMOTE_SFTP_IMAGE ="/SFTP/images"
	REMOTE_SFTP_VODIE ="/SFTP/vedio"
	REMOTE_SFTP_PDF ="/SFTP/pdf"
	LOCAL_IMAGES=""
	LOCAL_VEDIO=""
)

// FTP FTP.
type FtpConfigs struct {
	Addr       string
	User       string
	Password   string
	RemotePath map[string]string
	Timeout    xtime.Duration
	LocalPath  map[string]string
}


/**

 */
func NewUploadListSftpFile(conf *FtpConfigs,remoteFilePath string,localFilePath []string)(error) {

	// 这里换成实际的 SSH 连接的 用户名，密码，主机名或IP，SSH端口
	sftpClient, err := sftpConn(conf)
	defer sftpClient.Close()
	if err != nil {
		log.Info("sftpClient conn fail err=%s",err.Error())
		return err
	}

	for _, value := range localFilePath {
		//上传
		srcFile, err := os.Open(value)
		defer srcFile.Close()

		if err != nil {
			log.Info("sftpClient Open fail err=%s",err.Error())
			return err
		}

		var remoteFileName = path.Base(value)
		dstFile, err := sftpClient.Create(path.Join(remoteFilePath, remoteFileName))
		defer dstFile.Close()
		if err != nil {
			log.Info("sftpClient Create fail err=%s",err.Error())
			return err
		}

		buf := make([]byte, 1024)
		for {
			n, _ := srcFile.Read(buf)
			if n == 0 {
				break
			}
			dstFile.Write(buf)
		}

		os.Remove(value)

	}

	return err
}


/**
conf 连接sftp参数
remoteFilePath 远程路径
localFilePath 本地路径

var remoteDir = "/IMAGES"
var localFilePath = "/Users/a747/go/src/github.com/KXX747/gobase/sftp/570cab3a-443e-4d58-9e9a-31668ca2fa8e .jpg"

 */
func NewUploadSftpFile(conf *FtpConfigs,remoteFilePath,localFilePath string)(error){

	// 这里换成实际的 SSH 连接的 用户名，密码，主机名或IP，SSH端口
	sftpClient, err := sftpConn(conf)
	defer sftpClient.Close()
	if err != nil {
		log.Info("sftpClient conn fail err=%s",err.Error())
		return err
	}

	//上传
	srcFile, err := os.Open(localFilePath)
	defer srcFile.Close()

	if err != nil {
		log.Info("sftpClient Open fail err=%s",err.Error())
		return err
	}

	var remoteFileName = path.Base(localFilePath)
	dstFile, err := sftpClient.Create(path.Join(remoteFilePath, remoteFileName))
	defer dstFile.Close()
	if err != nil {
		log.Info("sftpClient Create fail err=%s",err.Error())
		return err
	}

	buf := make([]byte, 1024)
	for {
		n, _ := srcFile.Read(buf)
		if n == 0 {
			break
		}
		dstFile.Write(buf)
	}

	return err
}

/**
conf 连接sftp参数
remoteFilePath 远程路径
localFilePath 本地路径
var remoteFilePath = "/IMAGES/570cab3a-443e-4d58-9e9a-31668ca2fa8e .jpg"
var localDir = "/Users/a747/go/src/github.com/KXX747/gobase/sftp/download/"

 */
func NewDownLoadSftpFile(conf *FtpConfigs,remoteFilePath,localFilePath string) (error) {
	// 这里换成实际的 SSH 连接的 用户名，密码，主机名或IP，SSH端口
	sftpClient, err := sftpConn(conf)
	defer sftpClient.Close()
	if err != nil {
		log.Info("sftpClient conn fail err=%s",err.Error())
		return err
	}

	//
	srcFile, err := sftpClient.Open(remoteFilePath)
	defer srcFile.Close()
	if err != nil {
		log.Info("sftpClient Open fail err=%s",err.Error())
		return err
	}

	var localFileName = path.Base(remoteFilePath)
	dstFile, err := os.Create(path.Join(localFilePath, localFileName))
	if err != nil {
		log.Info("sftpClient Create fail err=%s",err.Error())
		return err
	}
	defer dstFile.Close()

	if _, err = srcFile.WriteTo(dstFile); err != nil {
		log.Info("sftpClient WriteTo fail err=%s",err.Error())
		return err
	}



	return err
}

/**
conn sftp
 */
func sftpConn(conf *FtpConfigs)(*sftp.Client, error) {
	var (
		auth         []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
		sshClient    *ssh.Client
		sftpClient   *sftp.Client
		err          error
	)
	// get auth method
	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(conf.Password))

	clientConfig = &ssh.ClientConfig{
		User:    conf.User,
		Auth:    auth,
		Timeout: 30 * time.Second,
		// ssh报错 ssh: must specify HostKeyCallback  解决方式 //需要验证服务端，不做验证返回nil就可以，点击HostKeyCallback看源码就知道了
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}
	// connet to ssh
	//addr = fmt.Sprintf("%s:%d", host, port)
	addr=conf.Addr
	sshClient, err = ssh.Dial("tcp",addr, clientConfig);
	if  err != nil {
		return nil, err
	}

	// create sftp client
	if sftpClient, err = sftp.NewClient(sshClient); err != nil {
		return nil, err
	}

	return sftpClient, nil
}
