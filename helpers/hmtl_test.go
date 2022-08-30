package helpers

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/net/html"
)

func TestInnerText(t *testing.T) {
	patterns := []struct {
		description string
		outerText   string
		tag         string
		expect      string
	}{
		{
			description: "p tag",
			outerText:   `<p style="text-align: left" data-sourcepos="21:19-21:162">データを送信すること。送信先に対してデータが正しい形式なのかに気を遣っている。引数は渡しちゃう。</p>`,
			tag:         "p",
			expect:      `データを送信すること。送信先に対してデータが正しい形式なのかに気を遣っている。引数は渡しちゃう。`,
		},
		{
			description: "div tag",
			outerText:   `<div style="text-align: left" data-sourcepos="21:19-21:162">データを送信すること。送信先に対してデータが正しい形式なのかに気を遣っている。引数は渡しちゃう。</div>`,
			tag:         "div",
			expect:      `データを送信すること。送信先に対してデータが正しい形式なのかに気を遣っている。引数は渡しちゃう。`,
		},
	}
	for i, p := range patterns {
		t.Run(fmt.Sprintf("%d-%s", i, p.description), func(t *testing.T) {
			node, err := html.Parse(strings.NewReader(p.outerText))
			require.NoError(t, err)
			actual := InnerText(node, p.tag)
			assert.Equal(t, p.expect, actual)
		})
	}
}
