package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strings"

	"github.com/misoul/yellowpage/dal"
)

type CompanyMySql struct {
	companies []dal.Company
	db *sql.DB
}

func InitDB() (*CompanyMySql, error) {
	db, err := sql.Open("mysql", "root:yellowpage@tcp(192.168.99.100:3306)/testdb1")
	if err != nil {
		log.Fatalf("Failed to create SQL.DB: %s", err.Error())  // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	// Cache all the existing rows in db
	rows, err := db.Query("SELECT * FROM companies")
	if err != nil {
		log.Fatalf("Failed to create db.Query: %s", err.Error())
	}
	defer rows.Close()

	var (
		id uint64
		name,industries,website,foundDate,stockCode,desc string
		companies []dal.Company
	)
	for rows.Next() {
		err := rows.Scan(&id, &name, &industries, &website, &foundDate, &stockCode, &desc)
		if err != nil {
			log.Fatalf(         "Failed to get next now: %s", err)
		}
		companies = append(companies, dal.Company{id,name,industries,website,foundDate,stockCode,desc})
	}
	err = rows.Err()
	if err != nil {
		log.Fatal("rows has error: %s", err)
	}

	log.Printf("Preloaded %d rows from database/companies",len(companies))

	return &CompanyMySql{db: db, companies: companies}, err
}

func (cin CompanyMySql) Finalize() {
	log.Println("Closing up CompanyMySql: ", cin)
	cin.db.Close()
}

func (cin CompanyMySql) Get(id uint64) dal.Company {
	var (
		id2 uint64
		name,industries,website,foundDate,stockCode,desc string
	)

	err := cin.db.QueryRow("SELECT * FROM companies where id = ?", id).Scan(&id2, &name, &industries, &website, &foundDate, &stockCode, &desc)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("No row with id: ", id)
		}
		log.Fatalf("Failed to create QueryRow: %s", err.Error())
	}
	return dal.Company{id,name,industries,website,foundDate,stockCode,desc}
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
