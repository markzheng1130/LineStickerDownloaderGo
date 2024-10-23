package main

import (
	"fmt"
	"line-sticker-downloader-go/stickerHandler"
)

func main() {
	stickerUrl := "https://store.line.me/stickershop/product/21802595/en"
	fmt.Printf("stickerUrl %s", stickerUrl)
	h := stickerHandler.StickerHandler.New{}
	fmt.Printf("[%#v]", h)
}
