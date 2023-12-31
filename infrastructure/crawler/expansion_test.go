package crawler_test

import (
	"testing"

	"github.com/capybara-alt/my-assemble/infrastructure/crawler"
	"github.com/capybara-alt/my-assemble/model"
	"github.com/stretchr/testify/assert"
)

func TestExpansionCrawler(t *testing.T) {
	test := struct {
		name string
		want model.Want[model.CrawlResultJSON]
	}{
		name: "[正常系]拡張機能ページのクロールテスト",
		want: model.Want[model.CrawlResultJSON]{
			Value: model.CrawlResultJSON{
				string(model.EXPANSION_TYPE): {
					"EXPANSION": {
						"ASSAULT ARMOR": {
							"攻撃力":  "1500",
							"衝撃力":  "2000",
							"衝撃残留": "1380",
							"爆発範囲": "60",
							"効果範囲": "200",
							"直撃補正": "230",
						},
						"PULSE ARMOR": {
							"耐久性能": "3300",
							"持続時間": "10",
						},
						"PULSE PROTECTION": {
							"レギュレーション": "1.03.1",
							"耐久性能":     "7000 (+3000)",
							"持続時間":     "25",
						},
						"TERMINAL ARMOR": {
							"レギュレーション": "1.03.1",
							"耐久性能":     "20000",
							"持続時間":     "5 (+3)",
						},
					},
				},
			},
			ErrMsg: "",
		},
	}
	c := crawler.NewExpansion()
	t.Run(test.name, func(t *testing.T) {
		results, err := c.Fetch()
		assert.Equal(t, err, nil)
		assert.Equal(t, test.want.Value, results)
	})
}
