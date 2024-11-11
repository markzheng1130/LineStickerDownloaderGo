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
	// Retrieve all sticker URLs.
	err := h.ParseStickerUrlList()
	if err != nil {
		fmt.Printf("Error when parsing sticker URL list: %v", err)
		return
	}

	fmt.Printf("[sticker URL list]%v\n", h.StickerUrlList)

	// Parallel download and collect all.
	// (TBD)

	// Save to local.
	// (TBD)

}
