package env

// BaseURL: in order to reduce the manipulation of data from the request, I filtered the data obtained by name,
// currencies,cca2,cca3,region,subregion.
// The following baseUrl should however according to the documentation of http://restcountries.com give us
// a json of format equivalent to the STRUCT tmpcountry defined in ./models/tmpcountry

const BaseURL = "https://restcountries.com/v3.1/%s?fields=name,currencies,cca2,cca3,region,subregion"