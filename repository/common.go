package repository

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
	"github.com/capybara-alt/my-assemble/config"
	"github.com/capybara-alt/my-assemble/model"
	"github.com/gocolly/colly"
)

func Crawl(targetURL config.CrawlTargetPage, unitType model.SecondaryUnitType) (model.CrawlResultJSON, error) {
	sunitType := string(unitType)
	collector := colly.NewCollector()
	unitList := make(model.CrawlResultJSON)
	unitList[sunitType] = make(model.CategoryGroupedJSON)
	collector.OnHTML("div#wikibody", func(root *colly.HTMLElement) {
		root.ForEach("div.plugin_contents > ul > li > a", func(i int, h *colly.HTMLElement) {
			if h.DOM.NextAllFiltered("ul").Length() != 0 {
				unitList[sunitType][h.Text] = make(model.NameGroupedJSON)
			}
			h.DOM.NextAllFiltered("ul").Each(func(i int, h2list *goquery.Selection) {
				h2list.ChildrenFiltered("li").ChildrenFiltered("a").Each(func(i int, h3list *goquery.Selection) {
					id, ok := h3list.Attr("href")
					if !ok {
						return
					}

					h3WithId := fmt.Sprintf("h3%s", id)
					h3 := root.DOM.ChildrenFiltered(h3WithId)
					table := h3.NextAllFiltered("table:not([class=\"atwiki_plugin_region\"])")
					if !isTargetTable(table) {
						return
					}
					unitList[sunitType][h.Text][h3.Text()] = getUnitInfo(table.First())
				})
			})
		})
	})

	err := collector.Visit(string(targetURL))
	return unitList, err
}

func isTargetTable(s *goquery.Selection) bool {
	filterTableStyle := "margin:0;padding:0;border:none;background-color:transparent;"
	filterTargetTable := fmt.Sprintf("tr[style=\"%s\"]", filterTableStyle)
	return s.ChildrenFiltered("tbody").ChildrenFiltered(filterTargetTable).Text() == ""
}

func getUnitInfo(s *goquery.Selection) model.UnitInfoJSON {
	info := make(model.UnitInfoJSON)
	s.ChildrenFiltered("tbody").ChildrenFiltered("tr").Each(func(i int, s *goquery.Selection) {
		header := s.ChildrenFiltered("th").First().Text()
		column := s.ChildrenFiltered("td").First().Text()
		info[header] = column
	})

	return info
}
