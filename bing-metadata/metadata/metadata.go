package metadata

import (
	"archive/zip"
	"encoding/xml"
	"strings"
)

type OfficeCoreProperty struct {
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

func NewProperties(reader *zip.Reader) (*OfficeCoreProperty, *OfficeAppProperty, error) {
	var coreProperty OfficeCoreProperty
	var appProperty OfficeAppProperty
	for _, file := range reader.File {
		switch file.Name {
		case "docProps/core.xml":
			if err := process(file, &coreProperty); err != nil {
				return nil, nil, err
			}
		case "docProps/app.xml":
			if err := process(file, &appProperty); err != nil {
				return nil, nil, err
			}
		default:
			continue
		}
	}
	return &coreProperty, &appProperty, nil
}

func process(file *zip.File, prop interface{}) error {
	rc, err := file.Open()
	if err != nil {
		return err
	}
	defer rc.Close()
	if err := xml.NewDecoder(rc).Decode(&prop); err != nil {
		return err
	}
	return nil
}
