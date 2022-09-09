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

type WonderWord struct {
	Index       int    `json:"index"`
	Term        string `json:"term"`
	Reading     string `json:"reading"`
	Explanation string `json:"explanation"`
	Example     string `json:"example"`
	Mean        string `json:"mean"`
}

type Data interface {
	GetText() string
	GetMean() string
	GetDetail() string
}

func (w *Word) GetText() string {
	return w.Text
}

func (w *Word) GetMean() string {
	return w.Text
}

func (w *Word) GetDetail() string {
	return w.Text
}

func (w *WonderWord) GetText() string {
	return w.Term
}

func (w *WonderWord) GetMean() string {
	return w.Mean
}

func (w *WonderWord) GetDetail() string {
	return w.Explanation
}

type MakeData func(c *html.Node, idx int, option ...string) Data

func MakeWord(c *html.Node, index int, options ...string) Data {
	if c.FirstChild == nil {
		c = c.Parent
	}
	idx := strings.Index(c.FirstChild.Data, ":")
	if idx < 0 {
		return &Word{
			Index:  index,
			Text:   c.FirstChild.Data,
			Detail: options[0],
			Link:   options[1],
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
		Detail:   options[0],
		Link:     options[1],
	}
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
			"",
		}
	}
	return nil
}
