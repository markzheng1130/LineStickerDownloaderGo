package stickerhandler

import (
	"fmt"
)

type Handler struct {
	WebSourceUrl   string
	StickerUrlList []string
}

func New(webSourceUrl string) *Handler {
	return &Handler{
		WebSourceUrl:   webSourceUrl,
		StickerUrlList: []string{},
	}
}

func (h *Handler) Run() {
	err := h.ParseStickerUrlList()
	if err != nil {
		fmt.Printf("Error when parsing sticker URL list: %v", err)
		return
	}

	fmt.Printf("[sticker URL list]%v\n", h.StickerUrlList)
}
