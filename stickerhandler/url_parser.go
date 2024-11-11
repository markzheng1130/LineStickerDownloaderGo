package stickerhandler

import (
	"encoding/json"
	"fmt"
	"html"
	"strings"
)

var TypeUrlMapping = map[string]string{
	"static":      "staticUrl",    // 一般靜態
	"animation":   "animationUrl", // 動圖png
	"popup":       "popupUrl",     // 全畫面
	"popup_sound": "popupUrl",     // 全畫面+有聲音
	"name":        "staticUrl",    // 可編輯文字
}

func (h *Handler) ParseStickerInfoFromLine(line string) error {
	stickerUrlToken1 := "data-preview=" // should contains this
	stickerUrlToken2 := "main.png"      // should not contains this

	if (strings.Contains(line, stickerUrlToken1)) && (!strings.Contains(line, stickerUrlToken2)) {
		// trim line.
		line = strings.TrimSpace(line)
		line = html.UnescapeString(line)
		line = line[len(stickerUrlToken1):] // should remove prefix
		line = strings.Trim(line, ">")      // should trim this
		line = strings.Trim(line, "'")      // should trim this
		fmt.Printf("[sticker URL][%v]\n", line)

		// retrieve the origin JSON data.
		var m map[string]interface{}
		if err := json.Unmarshal([]byte(line), &m); err != nil {
			return err
		}

		// retrieve the sticker URL.
		if stickerType, ok := m["type"].(string); ok {
			stickerUrlAttributeName := TypeUrlMapping[stickerType]
			stickerUrl := m[stickerUrlAttributeName]
			fmt.Printf("[stickerUrl][%v]\n", stickerUrl)
		}
	}

	return nil
}
