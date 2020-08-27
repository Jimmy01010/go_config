package common

import (
	"fmt"
	"os"
	"path"
	"regexp"
	"syscall"
)

/**
 * 验证dir是否存在，如果不存在则创建它
 */
func CheckAndMakeDir(dir string) error {
	if fi, err := os.Stat(dir); os.IsNotExist(err) {
		if crt_err := os.MkdirAll(dir, os.ModePerm); crt_err != nil {
			return fmt.Errorf("Failed to create directory. Attempt to create '%v' by %v has failed: %v", dir, os.Geteuid(), crt_err)
		}
	} else {
		if !fi.IsDir() {
			return fmt.Errorf("'%v' should be a directory", dir)
		}
		// Linux系统调用函数，检查调用进程是否可以对指定的文件执行某种操作。这里判断读写权限。
		if err := syscall.Access(dir, syscall.O_RDWR); err != nil {
			return fmt.Errorf("Directory '%v' is not writable by %v: %v", dir, os.Geteuid(), err)
		}

	}
	return nil
}


/**
 * 删除dir中与此模式匹配的所有文件
 */
func RemoveDirContent(dir string, pattern string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()

	re, err := regexp.Compile(pattern)
	if err != nil {
		return fmt.Errorf("Failed to compile regex pattern '%v': %v", pattern, err)
	}

	// 读取该目录下的每一个文件
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		if re.MatchString(name) {
			if err := os.RemoveAll(path.Join(dir, name)); err != nil {
				return fmt.Errorf("Failed to remove '%v': %v", name, err)
			}
		}
	}
	return nil
}