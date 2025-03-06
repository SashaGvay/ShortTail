package models

import (
	"fmt"
)

var (
	ErrNotFound = fmt.Errorf("not found")
)

type (
	URL struct {
		Original string `json:"original,omitempty"`
		Alias    string `json:"alias,omitempty"`
		QrCode   string `json:"qrcode,omitempty"`
	}
)
