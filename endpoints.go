package gofair

// Endpoints contains all the Betfair Exchange API endpoints.
var Endpoints = struct {
	Login,
	Identity,
	Betting,
	Account,
	Navigation string
}{
	Login:      "https://identitysso-cert.betfair.it/api/",
	Identity:   "https://identitysso.betfair.com/api/",
	Betting:    "https://api.betfair.com/exchange/betting/rest/v1.0/",
	Account:    "https://api.betfair.com/exchange/account/rest/v1.0/",
	Navigation: "https://api.betfair.com/exchange/betting/rest/v1/en/navigation/menu.json",
}
