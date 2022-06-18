package cbrru

//  https://www.onlinetool.io/xmltogo/

import (
	"encoding/xml"
)

//   https://cbr.ru/development/SXML/

type OneDay struct {
	XMLName xml.Name `xml:"ValCurs"`
	Text    string   `xml:",chardata"`
	Date    string   `xml:"Date,attr"`
	Name    string   `xml:"name,attr"`
	Valute  []struct {
		Text     string `xml:",chardata"`
		ID       string `xml:"ID,attr"`
		NumCode  string `xml:"NumCode"`
		CharCode string `xml:"CharCode"`
		Nominal  string `xml:"Nominal"`
		Name     string `xml:"Name"`
		Value    string `xml:"Value"`
	} `xml:"Valute"`
}

type RangeOfDays struct {
	XMLName    xml.Name `xml:"ValCurs"`
	Text       string   `xml:",chardata"`
	ID         string   `xml:"ID,attr"`
	DateRange1 string   `xml:"DateRange1,attr"`
	DateRange2 string   `xml:"DateRange2,attr"`
	Name       string   `xml:"name,attr"`
	Record     []struct {
		Text    string `xml:",chardata"`
		Date    string `xml:"Date,attr"`
		ID      string `xml:"Id,attr"`
		Nominal string `xml:"Nominal"`
		Value   string `xml:"Value"`
	} `xml:"Record"`
}
