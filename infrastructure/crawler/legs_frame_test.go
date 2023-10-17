package crawler_test

import (
	"testing"

	"github.com/capybara-alt/my-assemble/infrastructure/crawler"
	"github.com/capybara-alt/my-assemble/model"
	"github.com/stretchr/testify/assert"
)

func TestLegsFrameCrawler(t *testing.T) {
	test := struct {
		name string
		want model.Want[model.CrawlResultJSON]
	}{
		name: "[正常系]脚部ページのクロールテスト",
		want: model.Want[model.CrawlResultJSON]{
			Value: model.CrawlResultJSON{
				string(model.LEGS_UNIT_TYPE): {
					"二脚パーツ": {
						"AL-J-121 BASHO": {
							"メーカー":     "BAWS",
							"価格":       "141,000",
							"レギュレーション": "1.03.1",
							"AP":       "4160",
							"耐弾防御":     "362",
							"耐EN防御":    "325",
							"耐爆防御":     "398",
							"姿勢安定性能":   "824 (+50)",
							"積載上限":     "62600",
							"水平跳躍性能":   "132 (+7)",
							"垂直跳躍性能":   "25",
							"重量":       "20520",
							"EN負荷":     "300",
						},
						"AL-J-121/RC JAILBREAK": {
							"メーカー":     "BAWS",
							"価格":       "-",
							"レギュレーション": "1.03.1",
							"AP":       "2000",
							"耐弾防御":     "351",
							"耐EN防御":    "315",
							"耐爆防御":     "388",
							"姿勢安定性能":   "658 (+50)",
							"積載上限":     "62600",
							"水平跳躍性能":   "132 (+7)",
							"垂直跳躍性能":   "25",
							"重量":       "18560",
							"EN負荷":     "300",
						},
						"LG-011 MELANDER": {
							"メーカー":     "ベイラム",
							"価格":       "175,000",
							"レギュレーション": "1.03.1",
							"AP":       "4150",
							"耐弾防御":     "369",
							"耐EN防御":    "340",
							"耐爆防御":     "361",
							"姿勢安定性能":   "843 (+100)",
							"積載上限":     "60520",
							"水平跳躍性能":   "107 (+20)",
							"垂直跳躍性能":   "22",
							"重量":       "18700",
							"EN負荷":     "365",
						},
						"LG-012 MELANDER C3": {
							"メーカー":     "ベイラム",
							"価格":       "-",
							"レギュレーション": "1.03.1",
							"AP":       "3880",
							"耐弾防御":     "363",
							"耐EN防御":    "339",
							"耐爆防御":     "357",
							"姿勢安定性能":   "835 (+100)",
							"積載上限":     "55440",
							"水平跳躍性能":   "118 (+20)",
							"垂直跳躍性能":   "26",
							"重量":       "17210",
							"EN負荷":     "355",
						},
						"DF-LG-08 TIAN-QIANG": {
							"メーカー":     "大豊",
							"価格":       "350,000",
							"レギュレーション": "1.03.1",
							"AP":       "5300(+1100)",
							"耐弾防御":     "414",
							"耐EN防御":    "382",
							"耐爆防御":     "395",
							"姿勢安定性能":   "925 (+100)",
							"積載上限":     "82600",
							"水平跳躍性能":   "90 (+30)",
							"垂直跳躍性能":   "20",
							"重量":       "23600",
							"EN負荷":     "400",
						},
						"VP-422": {
							"メーカー":     "アーキバス",
							"価格":       "313,000",
							"レギュレーション": "1.03.1",
							"AP":       "3960",
							"耐弾防御":     "352",
							"耐EN防御":    "379",
							"耐爆防御":     "334",
							"姿勢安定性能":   "830 (+100)",
							"積載上限":     "58620",
							"水平跳躍性能":   "112 (+20)",
							"垂直跳躍性能":   "23",
							"重量":       "17900",
							"EN負荷":     "387",
						},
						"NACHTREIHER/42E": {
							"メーカー":   "シュナイダー",
							"価格":     "243,000",
							"AP":     "3500",
							"耐弾防御":   "295",
							"耐EN防御":  "330",
							"耐爆防御":   "298",
							"姿勢安定性能": "711",
							"積載上限":   "48650",
							"水平跳躍性能": "228",
							"垂直跳躍性能": "52",
							"重量":     "14030",
							"EN負荷":   "462",
						},
						"VE-42A": {
							"メーカー":     "アーキバス",
							"価格":       "504,000",
							"レギュレーション": "1.03.1",
							"AP":       "6000 (+800)",
							"耐弾防御":     "397",
							"耐EN防御":    "453",
							"耐爆防御":     "394",
							"姿勢安定性能":   "977 (+100)",
							"積載上限":     "85700",
							"水平跳躍性能":   "56 (+26)",
							"垂直跳躍性能":   "14",
							"重量":       "28950",
							"EN負荷":     "465",
						},
						"2C-2000 CRAWLER": {
							"メーカー":     "RaD",
							"価格":       "-",
							"レギュレーション": "1.03.1",
							"AP":       "3650",
							"耐弾防御":     "326",
							"耐EN防御":    "322",
							"耐爆防御":     "337",
							"姿勢安定性能":   "799 (+50)",
							"積載上限":     "53700 (+2500)",
							"水平跳躍性能":   "100 (+10)",
							"垂直跳躍性能":   "24",
							"重量":       "16300",
							"EN負荷":     "280",
						},
						"2C-3000 WRECKER": {
							"メーカー":     "RaD",
							"価格":       "139,000",
							"レギュレーション": "1.30.1",
							"AP":       "5220 (+1000)",
							"耐弾防御":     "350",
							"耐EN防御":    "312",
							"耐爆防御":     "383",
							"姿勢安定性能":   "1003 (+60)",
							"積載上限":     "68900",
							"水平跳躍性能":   "76",
							"垂直跳躍性能":   "17",
							"重量":       "21680",
							"EN負荷":     "680",
						},
						"2S-5000 DESSERT": {
							"メーカー":     "RaD",
							"価格":       "439,000",
							"レギュレーション": "1.30.1",
							"AP":       "5450 (+1040)",
							"耐弾防御":     "396",
							"耐EN防御":    "408",
							"耐爆防御":     "382",
							"姿勢安定性能":   "997 (+100)",
							"積載上限":     "77100",
							"水平跳躍性能":   "80 (+30)",
							"垂直跳躍性能":   "19",
							"重量":       "25880",
							"EN負荷":     "420",
						},
						"EL-TL-10 FIRMEZA": {
							"メーカー":     "エルカノ",
							"価格":       "400,000",
							"レギュレーション": "1.30.1",
							"AP":       "3600",
							"耐弾防御":     "328",
							"耐EN防御":    "266",
							"耐爆防御":     "270",
							"姿勢安定性能":   "737",
							"積載上限":     "52100",
							"水平跳躍性能":   "120 (+5)",
							"垂直跳躍性能":   "28",
							"重量":       "11200",
							"EN負荷":     "378",
						},
						"EL-PL-00 ALBA": {
							"メーカー":     "エルカノ",
							"価格":       "469,000",
							"レギュレーション": "1.30.1",
							"AP":       "3850",
							"耐弾防御":     "316",
							"耐EN防御":    "316",
							"耐爆防御":     "316",
							"姿勢安定性能":   "809 (+85)",
							"積載上限":     "50100",
							"水平跳躍性能":   "95 (+15)",
							"垂直跳躍性能":   "33",
							"重量":       "13150",
							"EN負荷":     "360",
						},
						"06-041 MIND ALPHA": {
							"メーカー":     "オールマインド",
							"価格":       "272,000",
							"レギュレーション": "1.30.1",
							"AP":       "4360",
							"耐弾防御":     "370",
							"耐EN防御":    "390",
							"耐爆防御":     "356",
							"姿勢安定性能":   "894 (+100)",
							"積載上限":     "63810",
							"水平跳躍性能":   "103 (+20)",
							"垂直跳躍性能":   "22",
							"重量":       "22110",
							"EN負荷":     "432",
						},
						"IA-C01L: EPHEMERA": {
							"メーカー":     "ルビコン調査技研",
							"価格":       "521,000",
							"レギュレーション": "1.03.1",
							"AP":       "3800",
							"耐弾防御":     "297",
							"耐EN防御":    "352",
							"耐爆防御":     "352",
							"姿勢安定性能":   "805 (+50)",
							"積載上限":     "55050",
							"水平跳躍性能":   "109 (+10)",
							"垂直跳躍性能":   "27",
							"重量":       "15200",
							"EN負荷":     "398",
						},
						"IB-C03L: HAL 826": {
							"メーカー":     "ルビコン調査技研",
							"価格":       "563,000",
							"レギュレーション": "1.30.1",
							"AP":       "4000",
							"耐弾防御":     "359",
							"耐EN防御":    "380",
							"耐爆防御":     "351",
							"姿勢安定性能":   "906 (+100)",
							"積載上限":     "64900",
							"水平跳躍性能":   "115 (+15)",
							"垂直跳躍性能":   "27",
							"重量":       "20890",
							"EN負荷":     "385",
						},
					},
				},
			},
			ErrMsg: "",
		},
	}
	c := crawler.NewLegsFrame()
	t.Run(test.name, func(t *testing.T) {
		results, err := c.Fetch()
		assert.Equal(t, err, nil)
		assert.Equal(t, test.want.Value, results)
	})
}