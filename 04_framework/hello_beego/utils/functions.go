package utils

import (
	"crypto/md5"
	"fmt"
)

/**
 *  Md5
 */
func Md5(buf []byte) string {
	hash := md5.New()
	hash.Write(buf)
	return fmt.Sprintf("%x", hash.Sum(nil))
}
