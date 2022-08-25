package services

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

//func TestDictionaryService_GetDictionary(t *testing.T) {
//	BucketSize = 2
//	ctx := context.Background()
//	data, err := NewDictionary().GetDictionary(ctx, "https://japanesetest4you.com/jlpt-n2-vocabulary-list/")
//	fmt.Printf("%+v", data)
//	assert.Equal(t, nil, err)
//}

func TestDictHandler_getDetail(t *testing.T) {
	ctx := context.Background()
	detail, err := NewDictionary().getDetail(ctx, "https://japanesetest4you.com/flashcard/%e8%b5%a4%e5%ad%97-akaji/", 1)
	fmt.Println(detail)
	assert.Equal(t, nil, err)
}
