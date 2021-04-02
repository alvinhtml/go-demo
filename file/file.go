package file

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func getFolderModTime() time.Time {
	path := "/Users/alvin/coder/go/src/alvinhtml.com/go-demo"
	oFile, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		fmt.Printf("file path: %s", path)
	}

	fileInfo, err := oFile.Stat()

	if err != nil {
		fmt.Println(err)
	}

	modify := fileInfo.ModTime()

	fmt.Println(modify)

	return modify
}

func SubModTime() {
	modify := getFolderModTime()
	// duration := time.Since(modify)
	now := time.Now()
	duration := now.Sub(modify)
	fmt.Println(duration.Hours())
}

// ForEachFolder 遍历文件夹
func ForEachFolder(path string) {
	// filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
	// 	if f.IsDir() {
	// 		fmt.Printf(path)
	// 	}
	// 	return err
	// })

	files, err := ioutil.ReadDir(path)

	if err != nil {
		fmt.Printf("error: %s", err)
	} else {
		for _, v := range files {
			if v.IsDir() && time.Since(v.ModTime()).Minutes() > 5 { // test
				fmt.Printf("%s/%s", path, v.Name())
				// fmt.Printf("File Name: %s, Type: %s, Folder: %t \n", v.Name(), v.Mode(), v.IsDir())
				fmt.Printf("上次修改时间距离现在已经过了: %f 小时。 \n", time.Since(v.ModTime()).Hours())
			}
		}
	}
}

func init() {
	// 遍历文件夹
	path := "/Users/alvin/coder/go"

	ticker := time.NewTicker(time.Second * 3)

	for _ = range ticker.C {
		ForEachFolder(path)
	}
}
