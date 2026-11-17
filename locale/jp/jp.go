// Package jp provides the Japan locale for seedfaker.
package jp

import "github.com/saphemmy/seedfaker"

// Locale is Japan: +81 phones, 7-digit postal codes (###-####).
var Locale = seedfaker.Locale{
	Code:            "ja_JP",
	Country:         "Japan",
	FirstNames:      []string{"Haruto", "Yui", "Sota", "Aoi", "Yuto", "Hina", "Haruki", "Sakura", "Riku", "Mei", "Ren", "Yuna", "Kaito", "Akari", "Sora"},
	LastNames:       []string{"Sato", "Suzuki", "Takahashi", "Tanaka", "Watanabe", "Ito", "Yamamoto", "Nakamura", "Kobayashi", "Kato", "Yoshida", "Yamada"},
	Titles:          []string{"Mr", "Mrs", "Ms", "Dr"},
	PhoneFormats:    []string{"+81 90 #### ####", "+81 80 #### ####"},
	Cities:          []string{"Tokyo", "Yokohama", "Osaka", "Nagoya", "Sapporo", "Fukuoka", "Kobe", "Kyoto", "Kawasaki", "Saitama"},
	States:          []string{"Tokyo", "Kanagawa", "Osaka", "Aichi", "Hokkaido", "Fukuoka", "Hyogo", "Kyoto", "Saitama", "Chiba"},
	Streets:         []string{"Chuo-dori", "Omotesando", "Ginza-dori", "Takeshita-dori", "Nakamise-dori"},
	PostcodeFormats: []string{"###-####"},
}

func init() { seedfaker.Register(Locale) }
