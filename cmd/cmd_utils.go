package cmd

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v3"
)

func ReadXML(file string, textData *[]byte) error {
	var err error
	xmlFile, err := os.Open(file)
	if err != nil {
		return err
	}

	// defer the closing of our xmlFile so that we can parse it later on
	defer xmlFile.Close()
	*textData, err = ioutil.ReadAll(xmlFile)
	if err != nil {
		return err
	}
	return nil
}

func ParseXML(data []byte, boxList *List) error {
	if err := xml.Unmarshal(data, &boxList); err != nil {
		return err
	}
	return nil
}

func GetJSON(boxList *List, outData *[]byte) error {
	var err error
	*outData, err = json.MarshalIndent((*boxList).Items, "", "  ")
	if err != nil {
		return err
	}
	return nil
}

func GetYAML(boxList *List, outData *[]byte) error {
	var err error
	*outData, err = yaml.Marshal((*boxList).Items)
	if err != nil {
		return err
	}
	return nil
}

func GetXML(boxList *List, outData *[]byte) error {
	var err error
	*outData, err = xml.MarshalIndent(*boxList, "", "  ")
	// prepend the xml info
	*outData = append([]byte{1}, []byte("<?xml version=\"1.0\" encoding=\"UTF-8\" standalone=\"yes\"?>")...)
	if err != nil {
		return err
	}
	return nil
}

func WriteToFile(outData *[]byte, file string) error {
	err := os.WriteFile(file, *outData, 0666)
	if err != nil {
		return err
	}
	return nil
}
