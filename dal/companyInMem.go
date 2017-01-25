package dal

import (
	"strings"
	"io/ioutil"
	"net/http"
	"fmt"
	"encoding/json"
	"log"
	"os"
)

const companiesFile = "./server/data/companies.json"

type CompanyInMem struct {
	companies []Company
}

func InitDB() (*CompanyInMem, error) {
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

	var companies []Company
	if err := json.Unmarshal(data, &companies); err != nil {
		errMsg := fmt.Sprintf("Unable to Unmarshal companies from data file (%s): %s", companiesFile, err)
		log.Fatal(errMsg, http.StatusInternalServerError)
		return nil, err
	}

	return &CompanyInMem{companies:companies}, nil
}

func (cin CompanyInMem) Get(id uint64) Company {
	return Company{} //TODO
}

func (cin CompanyInMem) Search(keywords []string) []Company {
	result := cin.companies
	if keywords != nil {
		result = Filter(cin.companies, func(c Company) bool {
			return strings.Contains(c.Name, keywords[0]) || strings.Contains(c.Desc, keywords[0])
		} ) //TODO: ghetto!
	}
	return result
}

//TODO: there should be a library for this already
func Filter(s []Company, fn func(Company) bool) []Company {
	var p []Company // == nil
	for _, v := range s {
		if fn(v) {
			p = append(p, v)
		}
	}
	return p
}
