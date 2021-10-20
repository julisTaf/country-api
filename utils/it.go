package utils

import (
	"CountryAPI/models"
	"fmt"
)

func PrintError()  {
	fmt.Printf("##############################################################\n")
	fmt.Printf("##### ERROR: perhaps you inputed an unknown coutry cat?  #####\n")
	fmt.Printf("##############################################################\n")
}

func PrintCountry(res models.Country){
	fmt.Println(fmt.Sprintf("" +
		"\nCountry: %s \n" +
		"Region: %s \n" +
		"SubRegion: %s \n" +
		"Cca2: %s \n" +
		"Cca3: %s \n" +
		"Currency: %s \n" +
		"Currency Symbol: %s \n",
		res.Name,
		res.Region,
		res.SubRegion,
		res.Cca2,
		res.Cca3,
		res.Currency,
		res.CurrencySymbol))
}