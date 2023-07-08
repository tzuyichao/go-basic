package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/samuel/go-zookeeper/zk"
)

func main() {
	// 檢查命令行參數
	if len(os.Args) < 2 {
		fmt.Println("請提供至少一個ZooKeeper伺服器的位址")
		return
	}

	// 取得命令行參數中的ZooKeeper伺服器位址
	serverList := os.Args[1]
	servers := strings.Split(serverList, ",")

	// 連接到ZooKeeper伺服器
	conn, _, err := zk.Connect(servers, time.Second)
	if err != nil {
		fmt.Printf("無法連接到ZooKeeper伺服器：%v\n", err)
		return
	}
	defer conn.Close()

	// 實現 ls 命令
	listCommand := func(path string) {
		children, _, err := conn.Children(path)
		if err != nil {
			fmt.Printf("無法取得目錄 %s 的內容：%v\n", path, err)
			return
		}

		fmt.Printf("目錄 %s 的內容：\n", path)
		for _, child := range children {
			fmt.Println(child)
		}
	}

	// 實現 get 命令
	getCommand := func(path string) {
		data, _, err := conn.Get(path)
		if err != nil {
			fmt.Printf("無法讀取檔案 %s 的內容：%v\n", path, err)
			return
		}

		fmt.Printf("檔案 %s 的內容：\n%s\n", path, string(data))
	}

	// 解析命令
	if len(os.Args) < 4 {
		fmt.Println("請提供要執行的命令：ls <目錄> 或 get <檔案路徑>")
		return
	}

	command := os.Args[2]
	path := os.Args[3]

	switch command {
	case "ls":
		listCommand(path)
	case "get":
		getCommand(path)
	default:
		fmt.Println("無效的命令")
	}
}
