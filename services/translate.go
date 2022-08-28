package services

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"cloud.google.com/go/translate"
	"github.com/hbl-ngocnd1/dictionary/models"
	"golang.org/x/text/language"
	"google.golang.org/api/option"
)

type translateService struct {
}

func NewTranslate() *translateService {
	return &translateService{}
}

type TranslateService interface {
	TranslateData(context.Context, []models.Word) []models.Word
}

func (t *translateService) TranslateData(ctx context.Context, data []models.Word) []models.Word {
	return translateData(ctx, data)
}

var BucketSize = 100

func translateData(ctx context.Context, data []models.Word) []models.Word {
	if os.Getenv("DEBUG") == "true" {
		BucketSize = 10
	}
	mapData := make(map[int]models.Word, len(data))
	maxIdx := 0
	for _, w := range data {
		if w.Index > maxIdx {
			maxIdx = w.Index
		}
		mapData[w.Index] = w
	}
	trans := make([]string, maxIdx+1)
	for i := 0; i <= maxIdx; i++ {
		if d, ok := mapData[i]; ok {
			trans[i] = d.MeanEng
		}
	}
	translated := make([]string, 0, len(trans))
	type bulkTransData struct {
		index int
		data  []string
	}
	c := make(chan bulkTransData)
	defer close(c)
	bulkLen := 0
	for i := 0; i < len(trans); {
		e := i + BucketSize
		if e > len(trans) {
			e = len(trans)
		}
		start := i
		end := e
		bulkLen++
		index := bulkLen
		go func(c chan bulkTransData) {
			bulk := translateToVN(ctx, trans[start:end])
			c <- bulkTransData{index: index, data: bulk}
		}(c)
		i = e
	}
	transMap := make(map[int][]string, bulkLen)
	for i := 1; i <= bulkLen; i++ {
		info := <-c
		transMap[info.index] = info.data
	}
	for i := 1; i <= bulkLen; i++ {
		if v, ok := transMap[i]; ok {
			translated = append(translated, v...)
		}
	}
	start := 0
	for i, vn := range translated {
		if v, ok := mapData[i]; ok {
			if start == 0 {
				start = i
			}
			v.MeanVN = vn
			mapData[i] = v
		}
	}
	result := make([]models.Word, 0, len(mapData))
	for i := start; i < len(mapData)+start; i++ {
		if w, ok := mapData[i]; ok {
			result = append(result, w)
		}
	}
	return result
}

func translateToVN(ctx context.Context, texts []string) []string {
	log.Println("start translate")
	defer log.Println("end translate")
	apiKey := os.Getenv("GOOGLE_APPLICATION_API_KEY")
	ctx, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()
	lang, _ := language.Parse("vi")
	client, err := translate.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Println(err)
	}
	defer client.Close()
	resp, err := client.Translate(ctx, texts, lang, nil)
	if err != nil {
		log.Println(fmt.Errorf("translate: %v", err))
		return texts
	}
	if len(resp) == 0 {
		log.Println(fmt.Errorf("translate returned empty response to text: %v", texts))
		return texts
	}
	result := make([]string, len(resp))
	for i, res := range resp {
		result[i] = res.Text
	}
	return result
}
