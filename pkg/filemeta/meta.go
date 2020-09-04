package filemeta

import (
	"fmt"
	"os"
	"syscall"
)

func GetMeta(path string) error {
	fileInfo, _ := os.Stat(path)
	fileSys := fileInfo.Sys().(*syscall.Win32FileAttributeData)
	fileAttributes:= fileSys.FileAttributes
	fmt.Println(fileAttributes)

	//stat, err := os.Stat(path)
	//if err != nil {
	//	return err
	//}
	//sys := stat.Sys()
	//fmt.Println(sys)
	return nil
}
