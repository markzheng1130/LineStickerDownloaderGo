package main

import (
	"line-sticker-downloader-go/stickerhandler"
)

func main() {
	// webSourceUrl := "https://store.line.me/stickershop/product/21802595/en" // type=static, url=staticUrl
	// webSourceUrl := "https://store.line.me/stickershop/product/12378/zh-Hant" //type=animation, url=animationUrl
	// webSourceUrl := "https://store.line.me/stickershop/product/7121/en" //type=popup, url=popupUrl
	// webSourceUrl := "https://store.line.me/stickershop/product/18085/zh-Hant" //type=popup_sound, url=popupUrl
	// webSourceUrl := "https://store.line.me/emojishop/product/5bc84438031a6704f8cff722/en" //type=static, url=staticUrl
	webSourceUrl := "https://store.line.me/stickershop/product/17054/zh-Hant" //type=name, url=staticUrl
	h := stickerhandler.New(webSourceUrl)
	h.Run()
}
