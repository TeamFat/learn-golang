/**
 * datetime: 2017-06-09
 * 递归遍历目录
 */
package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {

	filepath.Walk("../", func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		return nil
	})

}
