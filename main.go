package main

// getCountryInfo takes as parameters a string corresponding to a country name, or subregion (eg: north america... )
// it returns an array that contains a list of country
//  ./constants/values.go >> BASEURL
// Country struct >> ./models/country.go

func main() {
	getCountryInfo("north america")
	getCountryInfo("europe")
}

// TAFITA Julis
