package dal

type Company struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Industries  int32  `json:"industries"`
	Website     string `json:"website"`
	FoundDate   string `json:"foundDate"`
	StockCode   string `json:"stockCode"`
	Desc        string `json:"desc"`
}
