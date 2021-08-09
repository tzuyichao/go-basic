package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
	"syscall"
)

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	defer cleanUp()
	fmt.Println("程式執行中(按下ctrl + c來中斷)")
MainLoop:
    for {
		s := <- sigs
		switch s {
		case syscall.SIGINT:
			fmt.Println("程序中斷:", s)
			break MainLoop
		case syscall.SIGTERM:
			fmt.Println("程序中止:", s)
			break MainLoop
		}
	}
	fmt.Println("程式結果")
}

func cleanUp() {
	fmt.Println("進行清理作業...")
	for i := 0; i <= 10; i++ {
		fmt.Printf("刪除檔案 %v ...(Demo)\n", i)
		time.Sleep(time.Millisecond * 100)
	}
}