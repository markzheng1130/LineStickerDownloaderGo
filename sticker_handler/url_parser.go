package sticker_handler

import (
	"bufio"
	"encoding/json"
	"fmt"
	"html"
	"net/http"
	"strings"
)

var TypeToUrlMapping = map[string]string{
	"static":      "staticUrl",    // 一般靜態
	"animation":   "animationUrl", // 一般動圖
	"popup":       "popupUrl",     // 全畫面動圖
	"popup_sound": "popupUrl",     // 全畫面動圖 + 有聲音
	"name":        "staticUrl",    // 可編輯文字
}

func (h *Handler) ParseStickerUrlList() error {
	resp, err := http.Get(h.WebSourceUrl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() { // read body line by line.
		line := scanner.Text()
		err = h.ParseStickerUrl(line)
		if err != nil {
			return err
		}
	}

	return nil
}

func (h *Handler) ParseStickerUrl(line string) error {
	stickerUrlToken1 := "data-preview=" // should contain this.
	stickerUrlToken2 := "main.png"      // should not contain this.

	if (strings.Contains(line, stickerUrlToken1)) && (!strings.Contains(line, stickerUrlToken2)) {
		// trim line in orderly.
		line = strings.TrimSpace(line)      // should trim space
		line = line[len(stickerUrlToken1):] // should remove prefix.
		line = strings.Trim(line, ">")      // should trim this (when emoji).
		line = strings.Trim(line, "'")      // should trim this.
		line = html.UnescapeString(line)    // for human readable.
		fmt.Printf("[sticker data struct][%v]\n", line)

		// retrieve the origin JSON data.
		var m map[string]interface{}
		if err := json.Unmarshal([]byte(line), &m); err != nil {
			return err
		}

		// retrieve the sticker URL.
		if stickerType, ok := m["type"].(string); ok {
			stickerUrlAttributeName := TypeToUrlMapping[stickerType]
			if stickerUrl, ok := m[stickerUrlAttributeName].(string); ok {
				h.StickerUrlList = append(h.StickerUrlList, stickerUrl)
				fmt.Printf("[sticker URL][%v]\n", stickerUrl)
			}
		}
	}

	return nil
}
