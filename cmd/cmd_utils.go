package cmd

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
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

func OutputJSON(boxList *List) error {
	bytes, err := json.MarshalIndent((*boxList).Items, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(bytes))
	return nil
}

func OutputYAML(boxList *List) error {
	bytes, err := yaml.Marshal((*boxList).Items)
	if err != nil {
		return err
	}
	fmt.Print(string(bytes))
	return nil
}

func OutputXML(boxList *List) error {
	fmt.Println("<?xml version=\"1.0\" encoding=\"UTF-8\" standalone=\"yes\"?>")
	bytes, err := xml.MarshalIndent(*boxList, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(bytes))
	return nil
}
