package services

import (
	"context"
	"os"
	"testing"

	"github.com/hbl-ngocnd1/dictionary/models"
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
	data := []models.Word{
		{
			Index:   3,
			MeanEng: "cache",
		},
		{
			Index:   4,
			MeanEng: "computer",
		},
	}
	res := translateData(ctx, data)
	assert.Equal(t, len(data), len(res))
}
