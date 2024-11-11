package stickerhandler

import (
	"fmt"
	stickertype "line-sticker-downloader-go/misc"
	"strings"
)

func (h *Handler) ParseStickerInfoFromLine(line string) {
	s := strings.TrimSpace(line)

	// Parse sticker type
	stickerTypeToken := "FnStickerPreviewItem"
	if (strings.Contains(s, stickerTypeToken)) && (h.StickerType == stickertype.UnDefined) {
		stickerType := strings.Split(s, " ")[3]
		h.StickerType = stickertype.Type(stickerType)
		fmt.Printf("[sticker type][%v]\n", h.StickerType)
	}

	// parse sticker URL
	stickerUrlToken1 := "data-preview"
	stickerUrlToken2 := "main.png"
	if (strings.Contains(s, stickerUrlToken1)) && (!strings.Contains(s, stickerUrlToken2)) {
		fmt.Printf("[sticker URL][%v]\n", s)
	}

}
