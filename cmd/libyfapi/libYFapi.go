package main

import (
	"C"
	"encoding/json"
	"fmt"

	yahoofinanceapi "github.com/iolalla/yahoo-finance-api"
)

//export GetHistory
func GetHistory(symbol *C.char) *C.char {
	goSymbol := C.GoString(symbol)
	ticker := yahoofinanceapi.NewTicker(goSymbol)
	// Using default query parameters
	data, err := ticker.History(yahoofinanceapi.HistoryQuery{})
	if err != nil {
		return C.CString(fmt.Sprintf("Error: %v", err))
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return C.CString(fmt.Sprintf("Error parsing JSON: %v", err))
	}
	return C.CString(string(jsonData))
}

//export GetInfo
func GetInfo(symbol *C.char) *C.char {
	goSymbol := C.GoString(symbol)
	ticker := yahoofinanceapi.NewTicker(goSymbol)
	data, err := ticker.Info()
	if err != nil {
		return C.CString(fmt.Sprintf("Error: %v", err))
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return C.CString(fmt.Sprintf("Error parsing JSON: %v", err))
	}
	return C.CString(string(jsonData))
}

//export GetQuote
func GetQuote(symbol *C.char) *C.char {
	goSymbol := C.GoString(symbol)
	ticker := yahoofinanceapi.NewTicker(goSymbol)
	data, err := ticker.Quote()
	if err != nil {
		return C.CString(fmt.Sprintf("Error: %v", err))
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return C.CString(fmt.Sprintf("Error parsing JSON: %v", err))
	}
	return C.CString(string(jsonData))
}

//export GetOptionChain
func GetOptionChain(symbol *C.char) *C.char {
	goSymbol := C.GoString(symbol)
	ticker := yahoofinanceapi.NewTicker(goSymbol)
	data := ticker.OptionChain()
	jsonData, err := json.Marshal(data)
	if err != nil {
		return C.CString(fmt.Sprintf("Error parsing JSON: %v", err))
	}
	return C.CString(string(jsonData))
}

//export GetOptionChainByExpiration
func GetOptionChainByExpiration(symbol *C.char, expiration *C.char) *C.char {
	goSymbol := C.GoString(symbol)
	goExpiration := C.GoString(expiration)
	ticker := yahoofinanceapi.NewTicker(goSymbol)
	data := ticker.OptionChainByExpiration(goExpiration)
	jsonData, err := json.Marshal(data)
	if err != nil {
		return C.CString(fmt.Sprintf("Error parsing JSON: %v", err))
	}
	return C.CString(string(jsonData))
}

//export GetExpirationDates
func GetExpirationDates(symbol *C.char) *C.char {
	goSymbol := C.GoString(symbol)
	ticker := yahoofinanceapi.NewTicker(goSymbol)
	data := ticker.ExpirationDates()
	jsonData, err := json.Marshal(data)
	if err != nil {
		return C.CString(fmt.Sprintf("Error parsing JSON: %v", err))
	}
	return C.CString(string(jsonData))
}

func main() {}
