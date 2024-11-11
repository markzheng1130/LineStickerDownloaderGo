package stickertype

type Type string

const (
	UnDefined  Type = "undefined"
	Static     Type = "static-sticker"
	Animation  Type = "animation-sticker"
	PopupSound Type = "popup_sound-sticker"
)
