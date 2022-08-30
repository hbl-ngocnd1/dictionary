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
	TranslateData(context.Context, []TransData) []TransData
}

func (t *translateService) TranslateData(ctx context.Context, data []TransData) []TransData {
	return translateData(ctx, data)
}

var BucketSize = 100

type TransData struct {
	Index int
	Word  string
	Mean  string
}

func MakeTransDataFromWord(inputs []models.Word) []TransData {
	outs := make([]TransData, len(inputs))
	for i, _ := range inputs {
		outs[i].Index = inputs[i].Index
		outs[i].Word = inputs[i].MeanEng
	}
	return outs
}

func CompositeWordData(data []models.Word, trans []TransData) []models.Word {
	mapTrans := make(map[int]TransData, len(trans))
	for i, _ := range trans {
		mapTrans[trans[i].Index] = trans[i]
	}
	for i, _ := range data {
		if tran, ok := mapTrans[data[i].Index]; ok {
			data[i].MeanVN = tran.Mean
		}
	}
	return data
}

func MakeTransDataFromWonderWord(inputs []models.WonderWord) []TransData {
	outs := make([]TransData, len(inputs))
	for i, _ := range inputs {
		outs[i].Index = inputs[i].Index
		outs[i].Word = inputs[i].Term
	}
	return outs
}

func CompositeWonderWordData(data []models.WonderWord, trans []TransData) []models.WonderWord {
	mapTrans := make(map[int]TransData, len(trans))
	for i, _ := range trans {
		mapTrans[trans[i].Index] = trans[i]
	}
	for i, _ := range data {
		if tran, ok := mapTrans[data[i].Index]; ok {
			data[i].Mean = tran.Mean
		}
	}
	return data
}

func translateData(ctx context.Context, data []TransData) []TransData {
	if os.Getenv("DEBUG") == "true" {
		BucketSize = 10
	}
	mapData := make(map[int]TransData, len(data))
	minIdx := 0
	maxIdx := 0
	for _, w := range data {
		if w.Index < minIdx {
			minIdx = w.Index
		}
		if w.Index > maxIdx {
			maxIdx = w.Index
		}
		mapData[w.Index] = w
	}

	trans := make([]string, maxIdx+1)
	for i := minIdx; i <= maxIdx; i++ {
		if d, ok := mapData[i]; ok {
			trans[i] = d.Word
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
	for i, vn := range translated {
		if v, ok := mapData[i]; ok {
			v.Mean = vn
			mapData[i] = v
		}
	}

	result := make([]TransData, 0, len(mapData))
	for i := minIdx; i <= maxIdx; i++ {
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
