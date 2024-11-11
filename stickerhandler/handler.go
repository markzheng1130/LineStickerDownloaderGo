package stickerhandler

import (
	"fmt"
)

type Handler struct {
	Url   string
	Count int
	Body  []byte
}

func New(url string) *Handler {
	h := &Handler{}
	h.Url = url
	h.Count = 0

	return h
}

func (h *Handler) Run() {
	fmt.Printf("Handler: %#v\n", h)
	body, err := h.DownloadPage()

	if err != nil {
		fmt.Printf("Error when downloading web source: %v", err)
		return
	}

	fmt.Printf("body: %v", string(body))
}
