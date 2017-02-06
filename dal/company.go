package dal

import (
	"fmt"
	"strings"
	"github.com/jinzhu/gorm"
	"time"
)

//go:generate mockery -name=CompanyService -outpkg=mocks
type CompanyService interface {
	Get(id uint64) (Company, error)
	Create(company Company) (Company, error)
	Search(keywords string) ([]Company, error)
	Update(company Company) (Company, error)
	Finalize()
}

type Company struct {
	Name        string `json:"name"`
	Industries  string `json:"industries"`
	Website     string `json:"website"`
	FoundDate   time.Time `json:"foundDate"`
	StockCode   string `json:"stockCode"`
	Desc        string `json:"desc"`
	gorm.Model
}

func (c *Company) MatchKeywords(keywords []string) bool {
	return strings.Contains(c.Name, keywords[0]) || strings.Contains(c.Desc, keywords[0]) //TODO: ghetto!
}

func (c Company) String() string {
	return fmt.Sprintf("Company[%d,%s,%s,%s,%s,%s,%s]", c.ID, c.Name, c.Industries, c.Website, c.FoundDate, c.StockCode, c.Desc)
}
