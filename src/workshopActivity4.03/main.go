package main

import (
	"fmt"
	"os"
)

func supportedLocales() []string {
	var locales []string
	locales = append(locales, "en_US")
	locales = append(locales, "en_CN")
	locales = append(locales, "fr_CN")
	locales = append(locales, "fr_FR")
	locales = append(locales, "ru_RU")
	return locales
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("參數不足")
		os.Exit(1)
	}
	supported := supportedLocales()
	target := os.Args[1]
	for i:=0; i<len(supported); i++ {
		if supported[i] == target {
			fmt.Println("Locale passed is supported")
			os.Exit(0)
		}
	}
	fmt.Println("Locale not supported:", target)
}