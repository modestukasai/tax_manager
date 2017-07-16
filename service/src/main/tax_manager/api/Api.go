package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"log"
	"main/tax_manager/factory"
)

func Initialize(factory factory.ApplicationFactory) {
	router := httprouter.New()
	router.GET("/", GetIndex)
	router.GET("/municipalities", GetAllMunicipalities(factory))
	router.POST("/municipalities", SaveNewMunicipality)
	router.GET("/municipalities/:id", GetMunicipalityById)
	router.DELETE("/municipalities/:id", DeleteMunicipalityById)
	router.GET("/municipalities/:id/taxes", GetAllTaxes)
	router.POST("/municipalities/:id/taxes", SaveNewMunicipalityTax)
	router.GET("/municipalities/:id/taxes/:tax-id", GetTaxById)
	router.DELETE("/municipalities/:id/taxes/:tax-id", DeleteTaxById)
	router.GET("/calculate-tax", CalculateTax)

	log.Fatal(http.ListenAndServe(":8080", router))
}
