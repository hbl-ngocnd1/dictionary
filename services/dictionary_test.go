package services

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDictionaryService_GetDictionary(t *testing.T) {
	InitTest()
	BucketSize = 2
	ctx := context.Background()
	_, err := NewDictionary().GetDictionary(ctx, "https://japanesetest4you.com/jlpt-n2-vocabulary-list/")
	assert.Equal(t, nil, err)
}

func TestDictHandler_getDetail(t *testing.T) {
	ctx := context.Background()
	detail, err := NewDictionary().getDetail(ctx, "https://japanesetest4you.com/flashcard/%e8%b5%a4%e5%ad%97-akaji/", 1)
	assert.NotNil(t, detail)
	assert.Equal(t, nil, err)
}
