package dal

type Company struct {
	ID          uint64 `json:"id"`
	Name        string `json:"name"`
	Industries  string `json:"industries"`
	Website     string `json:"website"`
	FoundDate   string `json:"foundDate"`
	StockCode   string `json:"stockCode"`
	Desc        string `json:"desc"`
}

type CompanyService interface {
	Get(id uint64) Company
	Search(keywords string) []Company
	Update(company Company) Company
}

