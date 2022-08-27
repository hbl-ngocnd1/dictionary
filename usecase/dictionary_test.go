package usecase

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/hbl-ngocnd1/dictionary/models"
	"github.com/hbl-ngocnd1/dictionary/services"
	"github.com/hbl-ngocnd1/dictionary/services/mock_services"
	"github.com/stretchr/testify/assert"
)

func Test_GetDict(t *testing.T) {
	patterns := []struct {
		description             string
		start                   int
		pageSize                int
		notCache                string
		level                   string
		pwd                     string
		newMockDictService      func(ctrl *gomock.Controller) services.DictionaryService
		newMockTranslateService func(ctrl *gomock.Controller) services.TranslateService
		expect                  []models.Word
		err                     error
	}{
		{
			description: "success",
			start:       0,
			pageSize:    1,
			notCache:    "true",
			level:       "n1",
			pwd:         "sync_pass",
			newMockDictService: func(ctrl *gomock.Controller) services.DictionaryService {
				mock := mock_services.NewMockDictionaryService(ctrl)
				mock.EXPECT().GetDictionary(gomock.Any(), gomock.Eq("https://japanesetest4you.com/jlpt-n1-vocabulary-list/")).Return([]models.Word{
					{}, {}, {}, {},
				}, nil)
				return mock
			},
			newMockTranslateService: func(ctrl *gomock.Controller) services.TranslateService {
				mock := mock_services.NewMockTranslateService(ctrl)
				mock.EXPECT().TranslateData(gomock.Any(), gomock.Any()).Return([]models.Word{
					{}, {}, {}, {},
				})
				return mock
			},
			expect: []models.Word{
				{},
			},
			err: nil,
		},
		{
			description: "success over size",
			start:       0,
			pageSize:    8,
			notCache:    "true",
			level:       "n1",
			pwd:         "sync_pass",
			newMockDictService: func(ctrl *gomock.Controller) services.DictionaryService {
				mock := mock_services.NewMockDictionaryService(ctrl)
				mock.EXPECT().GetDictionary(gomock.Any(), gomock.Eq("https://japanesetest4you.com/jlpt-n1-vocabulary-list/")).Return([]models.Word{
					{}, {}, {}, {},
				}, nil)
				return mock
			},
			newMockTranslateService: func(ctrl *gomock.Controller) services.TranslateService {
				mock := mock_services.NewMockTranslateService(ctrl)
				mock.EXPECT().TranslateData(gomock.Any(), gomock.Any()).Return([]models.Word{
					{}, {}, {}, {},
				})
				return mock
			},
			expect: []models.Word{
				{}, {}, {}, {},
			},
			err: nil,
		},
		{
			description: "success with notCache and start over size",
			start:       8,
			pageSize:    1,
			notCache:    "true",
			level:       "n1",
			pwd:         "sync_pass",
			newMockDictService: func(ctrl *gomock.Controller) services.DictionaryService {
				mock := mock_services.NewMockDictionaryService(ctrl)
				mock.EXPECT().GetDictionary(gomock.Any(), gomock.Eq("https://japanesetest4you.com/jlpt-n1-vocabulary-list/")).Return([]models.Word{
					{},
				}, nil)
				return mock
			},
			newMockTranslateService: func(ctrl *gomock.Controller) services.TranslateService {
				mock := mock_services.NewMockTranslateService(ctrl)
				mock.EXPECT().TranslateData(gomock.Any(), gomock.Any()).Return([]models.Word{
					{},
				})
				return mock
			},
			expect: []models.Word{},
			err:    nil,
		},
		{
			description: "Success with cache and start over sized",
			start:       5,
			pageSize:    8,
			notCache:    "false",
			level:       "n1",
			pwd:         "sync_pass",
			expect:      []models.Word{},
			err:         nil,
		},
		{
			description: "Success with cacheData is nil",
			start:       0,
			pageSize:    1,
			notCache:    "true",
			level:       "n1",
			pwd:         "sync_pass",
			newMockDictService: func(ctrl *gomock.Controller) services.DictionaryService {
				mock := mock_services.NewMockDictionaryService(ctrl)
				mock.EXPECT().GetDictionary(gomock.Any(), gomock.Eq("https://japanesetest4you.com/jlpt-n1-vocabulary-list/")).Return([]models.Word{
					{}, {}, {}, {},
				}, nil)
				return mock
			},
			newMockTranslateService: func(ctrl *gomock.Controller) services.TranslateService {
				mock := mock_services.NewMockTranslateService(ctrl)
				mock.EXPECT().TranslateData(gomock.Any(), gomock.Any()).Return([]models.Word{
					{}, {}, {}, {},
				})
				return mock
			},
			expect: []models.Word{
				{},
			},
			err: nil,
		},
		{
			description: "permission denied",
			start:       0,
			pageSize:    8,
			notCache:    "true",
			level:       "n1",
			pwd:         "1111",
			expect:      nil,
			err:         PermissionDeniedErr,
		},
		{
			description: "Success use Cache",
			start:       0,
			pageSize:    1,
			notCache:    "false",
			level:       "n1",
			expect:      []models.Word{{}},
			err:         nil,
		},
		{
			description: "Success use Cache over size",
			start:       0,
			pageSize:    8,
			notCache:    "false",
			level:       "n1",
			expect:      []models.Word{{}, {}, {}, {}},
			err:         nil,
		},
		{
			description: "Fail on GetDictionary",
			start:       0,
			pageSize:    1,
			notCache:    "true",
			level:       "n1",
			pwd:         "sync_pass",
			newMockDictService: func(ctrl *gomock.Controller) services.DictionaryService {
				mock := mock_services.NewMockDictionaryService(ctrl)
				mock.EXPECT().GetDictionary(gomock.Any(), gomock.Eq("https://japanesetest4you.com/jlpt-n1-vocabulary-list/")).Return([]models.Word{
					{}, {}, {}, {},
				}, InvalidErr)
				return mock
			},
			newMockTranslateService: func(ctrl *gomock.Controller) services.TranslateService {
				mock := mock_services.NewMockTranslateService(ctrl)
				return mock
			},
			expect: nil,
			err:    InvalidErr,
		},
	}

	for i, p := range patterns {
		t.Run(fmt.Sprintf("%d:%s", i, p.description), func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			var mockDict services.DictionaryService
			if p.newMockDictService != nil {
				mockDict = p.newMockDictService(ctrl)
			}
			var mockTrans services.TranslateService
			if p.newMockDictService != nil {
				mockTrans = p.newMockTranslateService(ctrl)
			}
			var uc dictUseCase
			if p.notCache == "false" {
				uc = dictUseCase{
					dictionaryService: mockDict,
					translateService:  mockTrans,
					cacheData: map[string][]models.Word{
						"n1": {
							{},
							{},
							{},
							{},
						},
					},
				}
			} else {
				uc = dictUseCase{
					dictionaryService: mockDict,
					translateService:  mockTrans,
					cacheData:         nil,
				}
			}
			ctx := context.Background()
			os.Setenv("SYNC_PASS", "sync_pass")
			actual, err := uc.GetDict(ctx, p.start, p.pageSize, p.notCache, p.level, p.pwd)
			assert.Equal(t, p.expect, actual)
			assert.Equal(t, p.err, err)
		})
	}
}

func Test_GetDetail(t *testing.T) {
	patterns := []struct {
		description              string
		level                    string
		index                    int
		newMockDictionaryService func(ctrl *gomock.Controller) services.DictionaryService
		expect                   *string
		err                      error
	}{
		{
			description: "success",
			level:       "n1",
			index:       0,
			newMockDictionaryService: func(ctrl *gomock.Controller) services.DictionaryService {
				mock := mock_services.NewMockDictionaryService(ctrl)
				mock.EXPECT().GetDetail(gomock.Any(), gomock.Any(), gomock.Any()).Return("data", nil)
				return mock
			},
			expect: makeTestString("data"),
			err:    nil,
		},
		{
			description: "Fail on invalid level",
			level:       "error",
			index:       0,
			expect:      nil,
			err:         InvalidErr,
		},
		{
			description: "Fail on invalid index",
			level:       "n1",
			index:       99,
			expect:      nil,
			err:         InvalidErr,
		},
		{
			description: "Success null Link",
			level:       "nil link",
			index:       0,
			expect:      nil,
			err:         nil,
		},
		{
			description: "Fail GetDetail",
			level:       "n1",
			index:       0,
			newMockDictionaryService: func(ctrl *gomock.Controller) services.DictionaryService {
				mock := mock_services.NewMockDictionaryService(ctrl)
				mock.EXPECT().GetDetail(gomock.Any(), gomock.Any(), gomock.Any()).Return("", InvalidErr)
				return mock
			},
			expect: nil,
			err:    InvalidErr,
		},
	}

	for i, p := range patterns {
		t.Run(fmt.Sprintf("%d:%s", i, p.description), func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			var mockDict services.DictionaryService
			if p.newMockDictionaryService != nil {
				mockDict = p.newMockDictionaryService(ctrl)
			}
			uc := dictUseCase{
				dictionaryService: mockDict,
				cacheData: map[string][]models.Word{
					"n1": {
						models.Word{Link: "n1"},
					},
					"nil link": {
						models.Word{Link: ""},
					},
				},
			}
			ctx := context.Background()
			actual, err := uc.GetDetail(ctx, p.level, p.index)
			assert.Equal(t, p.expect, actual)
			assert.Equal(t, p.err, err)
		})
	}
}

func makeTestString(s string) *string {
	return &s
}
