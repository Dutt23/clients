package metadata

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
