package metadata

import "strings"

type OfficeCodeProperty struct {
	XMLName        string `xml:"coreProperties"`
	Creator        string `xml:"creator"`
	LastModifiedBy string `xml:"lastModifiedBy"`
}

type OfficeAppProperty struct {
	XMLName     string `xml:"Properties"`
	Application string `xml:"Application"`
	Company     string `xml:"Company"`
	AppVersion  string `xml:"AppVersion"`
}

var OfficeVersions = map[string]string{
	"16": "2016",
	"15": "2013",
	"14": "2010",
	"12": "2007",
	"11": "2003",
}

func (appProperty *OfficeAppProperty) GetMajorVersion() string {
	tokens := strings.Split(appProperty.AppVersion, ".")
	if len(tokens) < 2 {
		return "Unkown"
	}
	if version, ok := OfficeVersions[tokens[0]]; ok {
		return version
	}
	return "Unknown"
}
