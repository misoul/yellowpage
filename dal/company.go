package dal

import (
	"fmt"
	"strings"
)

//go:generate mockery -name=CompanyService -outpkg=mocks
type CompanyService interface {
	Get(id uint64) Company
	Search(keywords string) []Company
	Update(company Company) Company
	Finalize()
}

type Company struct {
	ID          uint64 `json:"id"`
	Name        string `json:"name"`
	Industries  string `json:"industries"`
	Website     string `json:"website"`
	FoundDate   string `json:"foundDate"`
	StockCode   string `json:"stockCode"`
	Desc        string `json:"desc"`
}

func (c *Company) MatchKeywords(keywords []string) bool {
	return strings.Contains(c.Name, keywords[0]) || strings.Contains(c.Desc, keywords[0]) //TODO: ghetto!
}

func (c *Company) String() string {
	return fmt.Sprintf("Company[%d,%s,%s,%s,%s,%s,%s]",c.ID,c.Name,c.Industries,c.Website,c.FoundDate,c.StockCode,c.Desc)
}
