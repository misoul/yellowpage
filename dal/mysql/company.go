package mysql

import (
	"log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/misoul/yellowpage/dal"
	"fmt"
	"time"
)

type CompanyMySql struct {
	db *gorm.DB
}

func InitCompany(dbUrl string) (*CompanyMySql, error) {
	db, err := gorm.Open("mysql", dbUrl)
	if err != nil {
		log.Fatalf("Failed to create SQL.DB: %s", err.Error())  // Just for example purpose. You should use proper error handling instead of panic
	}

	db.LogMode(true)
	db.AutoMigrate(&dal.Company{})

	return &CompanyMySql{db: db}, err
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

func (cin CompanyMySql) Get(id uint64) (dal.Company, error) {
	var company dal.Company
	err := cin.db.First(&company, id).Error
	if err != nil {
		log.Printf("Failed to get company %d: %s", id, err)
	}
	return company, err
}

func (cin CompanyMySql) Create(company dal.Company) (dal.Company, error) {
	err := cin.db.Create(&company).Error
	if err != nil {
		log.Printf("Failed to create company [%s]\n", company)
		log.Println(err)
	}
	return company, err
}

func (cin CompanyMySql) Update(company dal.Company) (dal.Company, error) {
	err := cin.db.Save(&company).Error
	if err != nil {
		log.Printf("Failed to update company [%s]\n", company)
		log.Println(err)
	}
	return company, err
}

func (cin CompanyMySql) Search(keywords []string) ([]dal.Company, error) {
	var companies []dal.Company
	if len(keywords) > 0 && len(keywords[0]) > 0 {
		search := "%" + keywords[0] + "%"
		err := cin.db.Where("`name` LIKE ? OR `desc` LIKE ?", search, search).Find(&companies).Error
		if err != nil {
			log.Printf("Failed to find %s: %s", keywords[0], err)
			return nil, err
		}
	} else {
		companies, _ = getAllCompanies(cin.db)
	}

	return companies, nil
}

func getAllCompanies(db *gorm.DB) ([]dal.Company, error) {
	var companies []dal.Company
	err := db.Find(&companies).Error
	if err != nil {
		log.Printf("Failed to query all: %s", err)
		return nil, err
	}
	return companies, nil
}
