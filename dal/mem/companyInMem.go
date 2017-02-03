package mem

import (
	"io/ioutil"
	"net/http"
	"fmt"
	"encoding/json"
	"log"
	"os"

	"github.com/misoul/yellowpage/dal"
)

const companiesFile = "./server/data/companies.json"

type CompanyInMem struct {
	companies []dal.Company
}

func InitCompany() (*CompanyInMem, error) {
	_, err := os.Stat(companiesFile)
	if err != nil {
		errMsg := fmt.Sprintf("Unable to stat the data file (%s): %s", companiesFile, err)
		log.Fatal(errMsg, http.StatusInternalServerError)
		return nil, err
	}

	data, err := ioutil.ReadFile(companiesFile) //TODO: not load file for each REST request
	if err != nil {
		errMsg := fmt.Sprintf("Unable to read the data file (%s): %s", companiesFile, err)
		log.Fatal(errMsg, http.StatusInternalServerError)
		return nil, err
	}

	var companies []dal.Company
	if err := json.Unmarshal(data, &companies); err != nil {
		errMsg := fmt.Sprintf("Unable to Unmarshal companies from data file (%s): %s", companiesFile, err)
		log.Fatal(errMsg, http.StatusInternalServerError)
		return nil, err
	}

	return &CompanyInMem{companies:companies}, nil
}

func (cin CompanyInMem) Finalize() {
	log.Println("Closing up CompanyInMem: ", cin)
}

func (cin CompanyInMem) Get(id uint64) dal.Company {
	return dal.Company{} //TODO
}

func (cin CompanyInMem) Update(company dal.Company) dal.Company {
	return dal.Company{} //TODO
}

func (cin CompanyInMem) Search(keywords []string) []dal.Company {
	result := cin.companies
	if keywords != nil {
		result = dal.FilterCompany(cin.companies, func(c dal.Company) bool {
			return c.MatchKeywords(keywords)
		})
	}
	return result
}
