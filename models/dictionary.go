package models

import (
	"strings"

	"github.com/hbl-ngocnd1/dictionary/helpers"
	"golang.org/x/net/html"
)

type Word struct {
	Index    int    `json:"index"`
	Text     string `json:"text"`
	Alphabet string `json:"alphabet"`
	MeanEng  string `json:"mean_eng"`
	MeanVN   string `json:"mean_vn"`
	Detail   string `json:"detail"`
	Link     string `json:"_"`
}

func MakeWord(c *html.Node, link, detail string, index int) *Word {
	if c.FirstChild == nil {
		c = c.Parent
	}
	idx := strings.Index(c.FirstChild.Data, ":")
	if idx < 0 {
		return &Word{
			Index:  index,
			Text:   c.FirstChild.Data,
			Detail: detail,
			Link:   link,
		}
	}
	mean := c.FirstChild.Data[idx+1:]
	arr := strings.Split(c.FirstChild.Data[:idx], " ")
	text := arr[0]
	var alphabet string
	if len(arr) > 1 {
		alphabet = strings.TrimRight(strings.TrimLeft(strings.Join(arr[1:], " "), "("), ")")
	}
	return &Word{
		Index:    index,
		Text:     text,
		Alphabet: alphabet,
		MeanEng:  mean,
		Detail:   detail,
		Link:     link,
	}
}

type WonderWord struct {
	Index       int    `json:"index"`
	Term        string `json:"term"`
	Reading     string `json:"reading"`
	Explanation string `json:"explanation"`
	Example     string `json:"example"`
}

func MakeWonderWork(tr *html.Node, idx int) *WonderWord {
	tds := helpers.GetListElementByTag(tr, "td")
	if len(tds) == 4 {
		return &WonderWord{
			idx,
			helpers.InnerText(tds[0], "td"),
			helpers.InnerText(tds[1], "td"),
			helpers.InnerText(tds[2], "td"),
			helpers.InnerText(tds[3], "td"),
		}
	}
	return nil
}
