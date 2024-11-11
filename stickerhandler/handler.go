package stickerhandler

import (
	"bufio"
	"fmt"
	stickertype "line-sticker-downloader-go/misc"
	"net/http"
)

type Handler struct {
	WebSourceUrl   string
	StickerType    stickertype.Type
	StickerUrlList []string
}

func New(webSourceUrl string) *Handler {
	h := &Handler{}
	h.WebSourceUrl = webSourceUrl
	h.StickerType = stickertype.UnDefined
	h.StickerUrlList = []string{}

	return h
}

func (h *Handler) Run() {
	fmt.Printf("Handler: %#v\n", h)

	_, err := h.ParseStickerUrlList()
	if err != nil {
		fmt.Printf("Error when parsing sticker URL list: %v", err)
		return
	}
}

func (h *Handler) ParseStickerUrlList() (string, error) {
	resp, err := http.Get(h.WebSourceUrl)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)

	for scanner.Scan() {
		line := scanner.Text()
		h.ParseStickerInfoFromLine(line)
	}

	return "", nil
}
