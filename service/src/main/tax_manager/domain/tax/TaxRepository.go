package tax

type TaxRepository interface {
	Save(tax Tax)
	FindTaxByMunicipalityIdAndTaxType(id int64, taxType TaxType) ([]Tax)
	FindTaxByMunicipalityIdAndTaxId(municipalityId int64, taxId int64) (*Tax)
	takeFirst(taxes []Tax) (*Tax)
	FindTaxByMunicipalityId(id int64) ([]Tax)
	DeleteAll()
	Delete(tax Tax)
}