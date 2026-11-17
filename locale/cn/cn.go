// Package cn provides the China locale for seedfaker.
package cn

import "github.com/saphemmy/seedfaker"

// Locale is China: +86 phones, 6-digit postal codes.
var Locale = seedfaker.Locale{
	Code:            "zh_CN",
	Country:         "China",
	FirstNames:      []string{"Wei", "Fang", "Jing", "Li", "Na", "Min", "Lei", "Yan", "Tao", "Jun", "Xin", "Hui", "Yong", "Juan", "Hao"},
	LastNames:       []string{"Wang", "Li", "Zhang", "Liu", "Chen", "Yang", "Huang", "Zhao", "Wu", "Zhou", "Xu", "Sun", "Ma", "Zhu"},
	Titles:          []string{"Mr", "Mrs", "Ms", "Dr"},
	PhoneFormats:    []string{"+86 1## #### ####"},
	Cities:          []string{"Beijing", "Shanghai", "Guangzhou", "Shenzhen", "Chengdu", "Hangzhou", "Wuhan", "Xian", "Nanjing", "Tianjin"},
	States:          []string{"Beijing", "Shanghai", "Guangdong", "Sichuan", "Zhejiang", "Hubei", "Shaanxi", "Jiangsu", "Tianjin", "Shandong"},
	Streets:         []string{"Nanjing Road", "Wangfujing Street", "Zhongshan Road", "Jianguo Road", "Changan Avenue"},
	PostcodeFormats: []string{"######"},
}

func init() { seedfaker.Register(Locale) }
