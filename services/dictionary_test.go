package services

import (
	"context"
	"fmt"
	"github.com/hbl-ngocnd1/dictionary/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDictionaryService_GetDictionary(t *testing.T) {
	InitTest()
	BucketSize = 2
	ctx := context.Background()
	_, err := NewDictionary().GetDictionary(ctx, "https://japanesetest4you.com/jlpt-n2-vocabulary-list/", models.MakeWord)
	assert.Equal(t, nil, err)
}

func TestDictHandler_getDetail(t *testing.T) {
	ctx := context.Background()
	detail, err := NewDictionary().getDetail(ctx, "https://japanesetest4you.com/flashcard/%e8%b5%a4%e5%ad%97-akaji/", 1)
	assert.NotNil(t, detail)
	assert.Equal(t, nil, err)
}

func TestDictHandler_GetITJapanWonderWork(t *testing.T) {
	ctx := context.Background()
	data, err := NewDictionary().GetITJapanWonderWork(ctx, "https://qiita.com/t_nakayama0714/items/478a8ed3a9ae143ad854", models.MakeWonderWork)
	fmt.Println(data)
	assert.Equal(t, nil, err)
}
