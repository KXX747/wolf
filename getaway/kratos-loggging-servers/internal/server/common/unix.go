package common

import (
	"path"
	"os"
	"fmt"
)

func ListenUNIX(addr string) (error) {
	dirname := path.Dir(addr)
	info, err := os.Stat(dirname)
	if err != nil {
		if !os.IsNotExist(err) {
			return  err
		}
		if err := os.MkdirAll(dirname, 0755); err != nil {
			return  fmt.Errorf("create directory %s error: %s", dirname, err)
		}
	}
	if err == nil && !info.IsDir() {
		return  fmt.Errorf("%s is already exists and not a directory", dirname)
	}
	if _, err := os.Stat(addr); err == nil {
		// remove old socket file
		os.Remove(addr)
	}
	//conn, err := net.ListenPacket("unixgram", addr)
	//if err != nil {
	//	return nil, err
	//}
	// make file permission to 666, so php can wirte span to this socket
	return  os.Chmod(addr, 0666)
}


func RemoveFilePath(path string)  {

	os.Remove(path)
}