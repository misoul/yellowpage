package mysql

import (
	//"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strings"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/misoul/yellowpage/dal"
	"fmt"
	"time"
)

type CompanyMySql struct {
	companies []dal.Company
	db *gorm.DB
}

func InitCompany(dbUrl string) (*CompanyMySql, error) {
	db, err := gorm.Open("mysql", dbUrl)
	if err != nil {
		log.Fatalf("Failed to create SQL.DB: %s", err.Error())  // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	db.AutoMigrate(&dal.Company{})

	// Cache all the existing rows in db
	companies, err := getAllCompanies(db)
	if err != nil {
		log.Fatal("Failed to prefetch companies: %s", err)
	}
	log.Printf("Preloaded %d rows from database/companies",len(companies))

	//insertRecords(db)

	return &CompanyMySql{db: db, companies: companies}, err
}

func insertRecords(db *gorm.DB) { //TODO: replace this with 'goose'
	time1 := time.Date(1995,1,1,0,0,0,0,time.Local)
	time2 := time.Date(1985,1,1,0,0,0,0,time.Local)
	time3 := time.Date(1990,1,1,0,0,0,0,time.Local)
	user1 := dal.Company{Name:"Google Inc", Industries:"IT", Website:"google.com", FoundDate:time1, StockCode:"NASDAQ:GOOGL", Desc:"Search, Ads & Beyond"}
	user2 := dal.Company{Name:"Tribeco", Industries:"Food & Beverages", Website:"tribeco.com", FoundDate:time2, StockCode:"HCMX:TRIB", Desc:"Giày"}
	user3 := dal.Company{Name:"Bitis", Industries:"Clothing", Website:"bitis.com", FoundDate:time3, StockCode:"HCMX:BITI", Desc:"Giày"}
	db.Create(&user1)
	db.Create(&user2)
	db.Create(&user3)
	fmt.Println(user1)
	//TODO: Unable to store string "Nước giải khát" for some reason
}

func (cin CompanyMySql) Finalize() {
	log.Println("Closing up CompanyMySql: ", cin)
	cin.db.Close()
}

func (cin CompanyMySql) Get(id uint64) dal.Company {
	var company dal.Company
	cin.db.First(&company, id)
	return company
}

func (cin CompanyMySql) Update(company dal.Company) dal.Company {
	return dal.Company{} //TODO
}

func (cin CompanyMySql) Search(keywords []string) []dal.Company {
	result := cin.companies
	if keywords != nil {
		result = dal.FilterCompany(cin.companies, func(c dal.Company) bool {
			return strings.Contains(c.Name, keywords[0]) || strings.Contains(c.Desc, keywords[0])
		} ) //TODO: ghetto!
	}
	return result
}

func getAllCompanies(db *gorm.DB) ([]dal.Company, error) {
	var companies []dal.Company

	db.Find(&companies)

	return companies, nil
}
