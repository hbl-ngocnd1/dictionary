package services

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func InitTest() {
	if os.Getenv("CI") == "true" {
		return
	}
	err := godotenv.Load("../.env_test")
	if err != nil {
		panic(err)
	}
}

func TestTranslateService_translateToVN(t *testing.T) {
	InitTest()
	ctx := context.Background()
	input := []string{"listed in public documentation", "test", "alone"}
	res := translateToVN(ctx, input)
	assert.Equal(t, len(input), len(res))
}

func TestTranslateService_TranslateData(t *testing.T) {
	InitTest()
	ctx := context.Background()
	data := []TransData{
		{
			Index: 3,
			Word:  "cache",
		},
		{
			Index: 4,
			Word:  "computer",
		},
	}
	res := translateData(ctx, data)
	fmt.Println(res)
	assert.Equal(t, len(data), len(res))
}
