package mem

import (
	"testing"
	"github.com/stretchr/testify/assert"

	"github.com/misoul/yellowpage/dal"
	"github.com/misoul/yellowpage/dal/mocks"
)

var nilCompany = dal.Company{}
var testCompany1 = dal.Company {ID:1, Name:"11", Industries:"", Website:"", FoundDate:"", StockCode:"", Desc:""}
var testCompany2 = dal.Company {ID:2, Name:"22", Industries:"", Website:"", FoundDate:"", StockCode:"", Desc:""}
var testCompany3 = dal.Company {ID:22, Name:"33", Industries:"", Website:"", FoundDate:"", StockCode:"", Desc:"2222"}
var testCompanies = []dal.Company {testCompany1, testCompany2, testCompany3}

func TestCompanyInMem_Get(t *testing.T) {
	cin := CompanyInMem{}
	assert.Equal(t, nilCompany, cin.Get(1), "should be equal")
}

func TestCompanyInMem_Update(t *testing.T) {
	cin := CompanyInMem{}
	assert.Equal(t, nilCompany, cin.Update(nilCompany), "should be equal")
}

func TestCompanyInMem_Search(t *testing.T) {
	cin := CompanyInMem{testCompanies}

	assert.Equal(t, []dal.Company{testCompany1}, cin.Search([]string{"1"}), "should find 1 company")
	assert.Equal(t, []dal.Company{testCompany2, testCompany3}, cin.Search([]string{"22"}), "should find 2 companies")
	assert.Equal(t, []dal.Company{}, cin.Search([]string{"1111"}), "should return empty list (not a nil)")
}

func TestMockCompanyService(t *testing.T) {
	m := &mocks.CompanyService{}
	m.On("Get", uint64(1)).Return(testCompany1)
	m.On("Get", uint64(2)).Return(testCompany2)

	assert.Equal(t, testCompany1, m.Get(uint64(1)))
	assert.Equal(t, testCompany2, m.Get(uint64(2)))

	m.AssertNumberOfCalls(t, "Get", 2)
}
