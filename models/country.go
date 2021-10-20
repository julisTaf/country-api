package models

//This is the desired structure for the final result.
//type Countries struct{}

type Country struct {
	Name string // official name in english
	Cca2 string // country code on 2 digits
	Cca3 string // country code on 3 digits
	Region string
	SubRegion string
	Currency string
	CurrencySymbol string
}