package stickerhandler

import "fmt"

type StickerHandler struct {
	StickerUrl   string
	StickerCount int
}

func (*StickerHandler) New(stickerUrl string) *StickerHandler {
	h := &StickerHandler{}
	h.StickerUrl = stickerUrl
	h.StickerCount = 0

	return h
}

func (*StickerHandler) Run() {
	fmt.Printf("handler is running...")
}
