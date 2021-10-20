package models

// This structure was used to refine the data chosen from the api end point
// Indeed, dealing with third-party APIs nested data (from a JSON object) is not that comfortable.
// Despite the NJSON library, I preferred using JSON unmarshaling to a struct that contains another struct as
//part of this test.
// Using the standard JSON package we are going to unmarshal Json from http://rescountries.com/ then restructure
//it like the following :

type TmpCountry struct {
	Name struct{
		Common string `json:"common"`
		Official string `json:"official"`
		NativeName map[string]interface{}
	}
	Currencies struct{
		USD struct{
			Name string `json:"name"`
			Symbol string `json:"symbol"`
		}
	}
	Cca2 string `json:"cca2"`
	Cca3 string `json:"cca3"`
	Region string `json:"region"`
	SubRegion string `json:"subregion"`
}