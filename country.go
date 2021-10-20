package main

import (
	"CountryAPI/constants"
	"CountryAPI/models"
	"CountryAPI/utils"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func _apiCall(_api string) ([]byte, error) {
	url := fmt.Sprintf(env.BaseURL, _api)
	res, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		e := errors.New(fmt.Sprintf("Unexpected API status code %s", res.Status))
		return []byte{}, e
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []byte{}, err
	}
	return body, nil
}

func CountriesByName(name string) ([]models.TmpCountry, error) {
	resData, err := _apiCall(fmt.Sprintf("subregion/%s", name))
	var c []models.TmpCountry
	err = json.Unmarshal(resData, &c)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func getCountryInfo(name string){
	providedData, err := CountriesByName(name)
	if err == nil {
		fmt.Printf("******************** ABOUT %s  ******************\n", name)
		for i := 0; i < len(providedData); i++ {
			c := providedData[i]
			res := models.Country{
				Name : c.Name.Official,
				Cca2:  c.Cca2,
				Cca3:  c.Cca3,
				Region: c.Region,
				SubRegion: c.SubRegion,
				Currency: c.Currencies.USD.Name,
				CurrencySymbol: c.Currencies.USD.Symbol,
			}
			utils.PrintCountry(res)
		}
		fmt.Println("********************************************************")
	} else {
		utils.PrintError()
	}
}