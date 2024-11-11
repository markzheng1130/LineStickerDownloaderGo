package stickerhandler

import (
	"fmt"
	"io"
	"net/http"
)

type Handler struct {
	Url   string
	Count int
}

func New(url string) *Handler {
	h := &Handler{}
	h.Url = url
	h.Count = 0

	return h
}

func (h *Handler) Run() {
	fmt.Printf("Handler: %#v\n", h)

	webSource, err := h.DownloadWebSource()
	if err != nil {
		fmt.Printf("Error when downloading web source: %v", err)
		return
	}

	stickerUrlList, err := h.ParseStickerUrl(webSource)
	if err != nil {
		fmt.Printf("Error when parsing sticker url: %v", err)
		return
	}
}

func (h *Handler) DownloadWebSource() ([]byte, error) {
	resp, err := http.Get(h.Url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, err
	}

	webSource, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return webSource, nil
}
