package main

import (
	"line-sticker-downloader-go/stickerhandler"
)

func main() {
	webSourceUrl := "https://store.line.me/stickershop/product/21802595/en"
	h := stickerhandler.New(webSourceUrl)
	h.Run()
}
