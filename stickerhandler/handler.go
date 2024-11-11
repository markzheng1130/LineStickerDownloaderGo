package stickerhandler

import (
	"bufio"
	"fmt"
	"net/http"
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
}

func (h *Handler) ParseStickerUrlList() error {
	resp, err := http.Get(h.WebSourceUrl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		line := scanner.Text()
		err = h.ParseStickerInfoFromLine(line)
		if err != nil {
			return err
		}
	}

	return nil
}
