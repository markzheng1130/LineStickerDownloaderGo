package stickerhandler

import (
	"encoding/json"
	"fmt"
	"html"
	"strings"
)

var TypeUrlMapping = map[string]string{
	"static":      "staticUrl",
	"animation":   "animationUrl",
	"popup":       "popupUrl",
	"popup_sound": "popupUrl",
	"name":        "staticUrl",
}

func (h *Handler) ParseStickerInfoFromLine(line string) error {
	line = strings.TrimSpace(line)
	line = html.UnescapeString(line)

	// parse sticker URL
	stickerUrlToken1 := "data-preview=" // should contains this
	stickerUrlToken2 := "main.png"      // should not contains this
	if (strings.Contains(line, stickerUrlToken1)) && (!strings.Contains(line, stickerUrlToken2)) {

		line = line[len(stickerUrlToken1):] // should remove prefix
		line = strings.Trim(line, "'")      // should trim single-quote
		fmt.Printf("[sticker URL][%v]\n", line)

		var m map[string]interface{}
		if err := json.Unmarshal([]byte(line), &m); err != nil {
			return err
		}

		if stickerType, ok := m["type"].(string); ok {
			stickerUrlAttributeName := TypeUrlMapping[stickerType]
			stickerUrl := m[stickerUrlAttributeName]
			fmt.Printf("[stickerUrl][%v]\n", stickerUrl)
		}
	}

	return nil
}
