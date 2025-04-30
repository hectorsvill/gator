package main

import (
	"fmt"

	"github.com/hectorsvill/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(cfg.DBURL)
	err = cfg.SetUser("lane")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(cfg.UserName)
}
