package tax

import (
	_ "github.com/ziutek/mymysql/godrv"
	"main/tax_manager/datasource"
	"database/sql"
	"time"
	"main/tax_manager"
	"log"
)

type mySQLTaxRepository struct {
	database datasource.Database
}

func NewMySQLTaxRepository() (TaxRepository) {
	return mySQLTaxRepository{}
}

func (this mySQLTaxRepository) IsExistingTax(tax Tax) (bool) {
	existingTaxes := this.FindTaxByMunicipalityId(tax.MunicipalityId)
	for _, existingTax := range existingTaxes {
		log.Println("Comparing ", tax, " to ",existingTax)
		if existingTax.TaxType == tax.TaxType && existingTax.From == tax.From && existingTax.To == tax.To {
			return true
		}
	}
	return false
}

func (this mySQLTaxRepository) Save(tax Tax) {
	this.database.Execute("INSERT `TAXES` SET `MUNICIPALITY_ID`=?,`FROM`=?,`TO`=?,`TAX_TYPE`=?,`VALUE`=?",
		tax.MunicipalityId,
		tax.From.Format(tax_manager.DEFAULT_DATE_FORMAT),
		tax.To.Format(tax_manager.DEFAULT_DATE_FORMAT),
		string(tax.TaxType),
		tax.Value)
}

func (this mySQLTaxRepository) FindTaxByMunicipalityIdAndTaxType(id int64, taxType TaxType) ([]Tax) {
	result := this.database.Query("SELECT * FROM `TAXES` WHERE `MUNICIPALITY_ID`=? AND `TAX_TYPE`=?", id, string(taxType))
	return mapTo(result)
}

func (this mySQLTaxRepository) FindTaxByMunicipalityIdAndTaxId(municipalityId int64, taxId int64) (*Tax) {
	result := this.database.Query("SELECT * FROM `TAXES` WHERE `MUNICIPALITY_ID`=? AND `ID`=?", municipalityId, taxId)
	return this.takeFirst(mapTo(result))
}

func (mySQLTaxRepository) takeFirst(taxes []Tax) (*Tax) {
	if len(taxes) == 1 {
		return &taxes[0]
	} else {
		return nil
	}
}

func (this mySQLTaxRepository) FindTaxByMunicipalityId(id int64) ([]Tax) {
	result := this.database.Query("SELECT * FROM `TAXES` WHERE `MUNICIPALITY_ID`=?", id)
	return mapTo(result)
}

func (this mySQLTaxRepository) DeleteAll() {
	this.database.Query("DELETE FROM `TAXES`")
}

func (this mySQLTaxRepository) Delete(tax Tax) {
	this.database.Query("DELETE FROM `TAXES` WHERE `ID`=?", tax.Id)
}

func mapTo(result *sql.Rows) ([]Tax) {
	foundTaxes := []Tax{}
	for result.Next() {
		var id int64
		var municipalityId int64
		var from string
		var to string
		var taxType string
		var value float64
		result.Scan(&id, &municipalityId, &from, &to, &taxType, &value)

		parsedFrom, _ := time.Parse(tax_manager.DEFAULT_DATE_FORMAT, from)
		parsedTo, _ := time.Parse(tax_manager.DEFAULT_DATE_FORMAT, to)
		foundTaxes = append(foundTaxes, Tax{
			Id:             id,
			MunicipalityId: municipalityId,
			From:           parsedFrom,
			To:             parsedTo,
			TaxType:        FindTaxTypeByValue(taxType),
			Value:          value})
	}
	return foundTaxes
}
