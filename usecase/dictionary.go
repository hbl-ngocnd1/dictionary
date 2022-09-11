package usecase

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/hbl-ngocnd1/dictionary/models"
	"github.com/hbl-ngocnd1/dictionary/services"
)

var PermissionDeniedErr = errors.New("usecase: permission denied")
var InvalidErr = errors.New("usecase: invalid")

type dictUseCase struct {
	translateService  services.TranslateService
	dictionaryService services.DictionaryService
	cacheData         map[string][]models.Word
	cacheWonderWord   [][]models.Data
	mu                sync.Mutex
}

func NewDictUseCase() *dictUseCase {
	return &dictUseCase{
		translateService:  services.NewTranslate(),
		dictionaryService: services.NewDictionary(),
	}
}

type DictUseCase interface {
	GetDict(context.Context, int, int, string, string, string, models.MakeData) ([]models.Word, error)
	GetDetail(context.Context, string, int) (*string, error)
	GetITJapanWonderWork(context.Context) ([][]models.Data, error)
}

func (u *dictUseCase) GetDict(ctx context.Context, start, pageSize int, notCache, level, pwd string, fn models.MakeData) ([]models.Word, error) {
	if notCache != "true" && u.cacheData != nil && u.cacheData[level] != nil && len(u.cacheData[level]) > 0 {
		log.Println("use data from cache")
		if start > len(u.cacheData[level]) {
			start = len(u.cacheData[level])
		}
		end := start + pageSize
		if end > len(u.cacheData[level]) {
			end = len(u.cacheData[level])
		}
		return u.cacheData[level][start:end], nil
	}
	if notCache == "true" {
		if len(strings.TrimSpace(pwd)) == 0 || pwd != os.Getenv("SYNC_PASS") {
			return nil, PermissionDeniedErr
		}
	}
	log.Println("use data from source")
	url := fmt.Sprintf("https://japanesetest4you.com/jlpt-%s-vocabulary-list/", level)
	data, err := u.dictionaryService.GetDictionary(ctx, url, fn)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	words := make([]models.Word, len(data))
	for i, d := range data {
		if word, ok := d.(*models.Word); ok {
			words[i] = *word
		}
	}
	words = services.CompositeWordData(words, u.translateService.TranslateData(ctx, services.MakeTransDataFromWord(words)))
	u.mu.Lock()
	if u.cacheData == nil {
		u.cacheData = make(map[string][]models.Word)
	}
	u.cacheData[level] = words
	u.mu.Unlock()
	if start > len(words) {
		start = len(words)
	}
	end := start + pageSize
	if end > len(words) {
		end = len(words)
	}
	return words[start:end], nil
}

func (u *dictUseCase) GetDetail(ctx context.Context, level string, index int) (*string, error) {
	if u.cacheData[level] == nil || index >= len(u.cacheData[level]) {
		return nil, InvalidErr
	}
	detailURL := u.cacheData[level][index].Link
	if strings.TrimSpace(detailURL) == "" {
		return nil, nil
	}
	data, err := u.dictionaryService.GetDetail(ctx, detailURL, index)
	if err != nil {
		return nil, err
	}
	u.cacheData[level][index].Detail = data
	return &data, nil
}

func (u *dictUseCase) GetITJapanWonderWork(ctx context.Context) ([][]models.Data, error) {
	if u.cacheWonderWord == nil && len(u.cacheWonderWord) > 0 {
		return u.cacheWonderWord, nil
	}
	data, err := u.dictionaryService.GetITJapanWonderWork(ctx, "https://qiita.com/t_nakayama0714/items/478a8ed3a9ae143ad854")
	if err != nil {
		return nil, err
	}
	for idx := range data {
		data[idx] = services.CompositeWonderWordData(data[idx], u.translateService.TranslateData(ctx, services.MakeTransDataFromWonderWord(data[idx])))
	}
	u.cacheWonderWord = data
	return data, nil
}
