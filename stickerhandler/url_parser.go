package stickerhandler

import (
	"encoding/json"
	"fmt"
	"html"
	"strings"
)

func (h *Handler) ParseStickerInfoFromLine(line string) error {
	line = strings.TrimSpace(line)
	line = html.UnescapeString(line)

	// parse sticker URL
	stickerUrlToken1 := "data-preview=" // should contains this
	stickerUrlToken2 := "main.png"      // should not contains this
	if (strings.Contains(line, stickerUrlToken1)) && (!strings.Contains(line, stickerUrlToken2)) {

		line = line[len(stickerUrlToken1):] // should remove this prefix
		line = strings.Trim(line, "'")      // should trim single-quote from 2 end
		fmt.Printf("[sticker URL][%v]\n", line)

		var m map[string]interface{} // 用1個map來接，這是固定寫法，才有辨法剖析出巢狀結構

		if err := json.Unmarshal([]byte(line), &m); err != nil {
			return err
		}

		fmt.Printf("[type][%v][url][%v]", m["type"], m[""])

	}

	return nil
}
