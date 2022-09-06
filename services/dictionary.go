package services

import (
	"bytes"
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
	"time"

	"github.com/hbl-ngocnd1/dictionary/helpers"
	"github.com/hbl-ngocnd1/dictionary/models"
	"golang.org/x/net/html"
)

type dictionaryService struct {
}

func NewDictionary() *dictionaryService {
	return &dictionaryService{}
}

type DictionaryService interface {
	GetDictionary(context.Context, string) ([]models.Word, error)
	GetGrammar(context.Context, string) ([]models.Word, error)
	GetDetail(context.Context, string, int) (string, error)
	GetITJapanWonderWork(context.Context, string) ([][]models.WonderWord, error)
}

func (d *dictionaryService) GetDictionary(ctx context.Context, url string) ([]models.Word, error) {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	//log.Println(string(body))
	elms := strings.Split(url, "/")
	fileName := elms[len(elms)-2]
	doc, err := html.Parse(bytes.NewReader(body))
	if err != nil {
		log.Fatal(err)
	}
	tag := helpers.GetElementByClass(doc, "entry clearfix")
	if tag == nil {
		body, err = os.ReadFile(path.Join("cache", fileName))
		if err != nil {
			return nil, nil
		}
		doc, err = html.Parse(bytes.NewReader(body))
		if err != nil {
			log.Fatal(err)
		}
		tag = helpers.GetElementByClass(doc, "entry clearfix")
	}
	if tag == nil {
		log.Println("can't get entry clearfix")
		return nil, nil
	}
	targets := helpers.GetListElementByTag(tag, "p")
	if len(targets) == 0 {
		log.Println("can't get entry clearfix p")
		return nil, nil
	}
	_ = os.WriteFile(path.Join("cache", fileName), body, 0666)
	if len(targets) > 2 {
		targets = targets[2:]
	}
	c := make(chan models.Word)
	defer close(c)
	cLen := 0
	for i, target := range targets {
		id := i
		if os.Getenv("DEBUG") == "true" && i > 40 {
			break
		}
		cLen++
		tar := target
		go func(c chan models.Word) {
			child := tar.FirstChild
			if child == nil {
				child = tar
			}
			var detail string
			var errDetail error
			detailURL, ok := helpers.GetAttribute(child, "href")
			if ok {
				detail, errDetail = d.getDetail(ctx, detailURL, id)
				if errDetail != nil {
					log.Println(errDetail)
				}
			}
			w := models.MakeWord(child, detailURL, detail, id)
			if w == nil {
				return
			}
			c <- *w
		}(c)
	}
	data := make([]models.Word, 0, cLen)
	mapResult := make(map[int]*models.Word)
	maxIdx := 0
	for i := 0; i < cLen; i++ {
		info := <-c
		maxIdx = i
		mapResult[info.Index] = &info
	}
	for i := 0; i <= maxIdx; i++ {
		if mapResult[i] == nil {
			continue
		}
		data = append(data, *mapResult[i])
	}
	log.Println("clone done")
	return data, nil
}

func (d *dictionaryService) GetGrammar(ctx context.Context, url string) ([]models.Word, error) {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	//log.Println(string(body))
	elms := strings.Split(url, "/")
	fileName := elms[len(elms)-2]
	doc, err := html.Parse(bytes.NewReader(body))
	if err != nil {
		log.Fatal(err)
	}
	tag := helpers.GetElementByClass(doc, "entry clearfix")
	if tag == nil {
		body, err = os.ReadFile(path.Join("cache", fileName))
		if err != nil {
			return nil, nil
		}
		doc, err = html.Parse(bytes.NewReader(body))
		if err != nil {
			log.Fatal(err)
		}
		tag = helpers.GetElementByClass(doc, "entry clearfix")
	}
	if tag == nil {
		log.Println("can't get entry clearfix")
		return nil, nil
	}
	targets := helpers.GetListElementByTag(tag, "p")
	if len(targets) == 0 {
		log.Println("can't get entry clearfix p")
		return nil, nil
	}
	_ = os.WriteFile(path.Join("cache", fileName), body, 0666)
	if len(targets) > 2 {
		targets = targets[2:]
	}
	c := make(chan models.Word)
	defer close(c)
	cLen := 0
	for i, target := range targets {
		id := i
		if os.Getenv("DEBUG") == "true" && i > 40 {
			break
		}
		cLen++
		tar := target
		go func(c chan models.Word) {
			child := tar.FirstChild
			if child == nil {
				child = tar
			}
			var detail string
			var errDetail error
			detailURL, ok := helpers.GetAttribute(child, "href")
			if ok {
				detail, errDetail = d.getDetail(ctx, detailURL, id)
				if errDetail != nil {
					log.Println(errDetail)
				}
			}
			w := models.MakeGrammarWord(child, detailURL, detail, id)
			if w == nil {
				return
			}
			c <- *w
		}(c)
	}
	data := make([]models.Word, 0, cLen)
	mapResult := make(map[int]*models.Word)
	maxIdx := 0
	for i := 0; i < cLen; i++ {
		info := <-c
		maxIdx = i
		mapResult[info.Index] = &info
	}
	for i := 0; i <= maxIdx; i++ {
		if mapResult[i] == nil {
			continue
		}
		data = append(data, *mapResult[i])
	}
	log.Println("clone done")
	return data, nil
}

func (d *dictionaryService) GetDetail(ctx context.Context, url string, i int) (string, error) {
	return d.getDetail(ctx, url, i)
}

func (d *dictionaryService) getDetail(ctx context.Context, url string, i int) (string, error) {
	log.Println("start with goroutine ", i)
	defer log.Println("end with goroutine ", i)
	ctx, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		log.Println(err)
		return "", err
	}
	client := http.DefaultClient

	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return "", err
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	elms := strings.Split(url, "/")
	fileName := elms[len(elms)-2]

	doc, err := html.Parse(bytes.NewReader(body))
	if err != nil {
		return "", err
	}
	var data []string
	tag := helpers.GetElementByClass(doc, "entry clearfix")
	if tag == nil {
		body, err = os.ReadFile(path.Join("cache", fileName))
		if err != nil {
			return "", nil
		}
		doc, err = html.Parse(bytes.NewReader(body))
		if err != nil {
			log.Fatal(err)
		}
		tag = helpers.GetElementByClass(doc, "entry clearfix")
	}
	if tag != nil {
		_ = os.WriteFile(path.Join("cache", fileName), body, 0666)
		nodes := helpers.GetListElementByTag(tag, "p")
		data = []string{helpers.RenderNode(nodes[1])}
		for _, node := range nodes[3:] {
			if node.FirstChild != nil && node.FirstChild.Data == "img" {
				continue
			}
			data = append(data, helpers.RenderNode(node))
		}
	}
	return strings.Join(data, ""), nil
}

func (d *dictionaryService) GetITJapanWonderWork(ctx context.Context, url string) ([][]models.WonderWord, error) {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	doc, err := html.Parse(bytes.NewReader(body))
	if err != nil {
		log.Fatal(err)
	}
	contentDiv := helpers.GetElementById(doc, "personal-public-article-body")
	nextDiv := helpers.GetListElementByTag(contentDiv, "div")
	tables := helpers.GetListElementByTag(nextDiv[0], "table")
	data := make([][]models.WonderWord, len(tables))
	for idx1, table := range tables {
		tbody := helpers.GetListElementByTag(table, "tbody")
		trs := helpers.GetListElementByTag(tbody[0], "tr")
		data[idx1] = make([]models.WonderWord, len(trs))
		for idx2, tr := range trs {
			work := models.MakeWonderWork(tr, idx2)
			if work != nil {
				data[idx1][idx2] = *work
			}
		}
	}
	return data, nil
}
