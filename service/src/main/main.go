package main

import (
	"main/tax_manager/api"
	"main/tax_manager/datasource"
	"main/tax_manager/file"
)

func init() {
	datasource.Database{}.CheckConnection()
}

func main() {
	file.PopulateDataFromFile{}.Populate("tax_file.csv")
	api.Initialize()
}
