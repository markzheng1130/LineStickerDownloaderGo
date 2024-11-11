package stickerhandler

import (
	"fmt"
	"io"
	"net/http"
)

func (h *Handler) DownloadPage() ([]byte, error) {
	resp, err := http.Get(h.Url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, err
}

func (h *Handler) ParseUrl() []string {
	fmt.Printf("Parse URL from: %v\n", h.Url)
	urlList := []string{}
	return urlList
}
