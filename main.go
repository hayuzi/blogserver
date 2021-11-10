package main

import (
	"github.com/hayuzi/blogserver/cmd"
	"log"
)

func main() {
	// 使用cobra改造入口文件
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute err: %v", err)
	}

}
