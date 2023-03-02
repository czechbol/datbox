package cmd

import "encoding/xml"

type List struct {
	XMLName      xml.Name `xml:"list" json:"-" yaml:"-"`
	XmlNS        string   `xml:"xmlns,attr,omitempty" json:"-" yaml:"-"`
	SourceSystem string   `xml:"source-system,attr,omitempty" json:"-" yaml:"-"`
	Items        []Box    `xml:"box,omitempty" json:"box,omitempty" yaml:"box,omitempty"`
}

type Box struct {
	XMLName   xml.Name  `xml:"box" json:"-" yaml:"-"`
	ID        string    `xml:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
	Type      string    `xml:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	Subtype   int       `xml:"subtype,omitempt" json:"subtype,omitempty" yaml:"subtype,omitempty"`
	Name      Name      `xml:"name" json:"name" yaml:"name"`
	Ico       string    `xml:"ico,omitempty" json:"ico,omitempty" yaml:"ico,omitempty"`
	Address   Address   `xml:"address" json:"address" yaml:"address"`
	Pdz       bool      `xml:"pdz,omitempty" json:"pdz,omitempty" yaml:"pdz,omitempty"`
	Ovm       bool      `xml:"ovm,omitempty" json:"ovm,omitempty" yaml:"ovm,omitempty"`
	Hierarchy Hierarchy `xml:"hierarchy" json:"hierarchy" yaml:"hierarchy"`
	IdOVM     string    `xml:"idOVM,omitempty" json:"idOVM,omitempty" yaml:"idOVM,omitempty"`
}

type Address struct {
	XMLName      xml.Name `xml:"address" json:"-" yaml:"-"`
	Code         int      `xml:"code,omitempty" json:"code,omitempty" yaml:"code,omitempty"`
	City         string   `xml:"city,omitempty" json:"city,omitempty" yaml:"city,omitempty"`
	District     string   `xml:"district,omitempty" json:"district,omitempty" yaml:"district,omitempty"`
	Street       string   `xml:"street,omitempty" json:"street,omitempty" yaml:"street,omitempty"`
	Cp           string   `xml:"cp,omitempty" json:"cp,omitempty" yaml:"cp,omitempty"`
	Co           string   `xml:"co,omitempty" json:"co,omitempty" yaml:"co,omitempty"`
	Ce           string   `xml:"ce,omitempty" json:"ce,omitempty" yaml:"ce,omitempty"`
	Zip          string   `xml:"zip,omitempty" json:"zip,omitempty" yaml:"zip,omitempty"`
	Region       string   `xml:"region,omitempty" json:"region,omitempty" yaml:"region,omitempty"`
	AddressPoint int      `xml:"addressPoint,omitempty" json:"addressPoint,omitempty" yaml:"addressPoint,omitempty"`
	State        string   `xml:"state,omitempty" json:"state,omitempty" yaml:"state,omitempty"`
	FullAddress  string   `xml:"fullAddress,omitempty" json:"fullAddress,omitempty" yaml:"fullAddress,omitempty"`
}

type Name struct {
	XMLName   xml.Name `xml:"name" json:"-" yaml:"-"`
	Person    Person   `xml:"person,omitempty" json:"person,omitempty" yaml:"person,omitempty"`
	TradeName string   `xml:"tradeName,omitempty" json:"tradeName,omitempty" yaml:"tradeName,omitempty"`
}

type Person struct {
	XMLName         xml.Name `xml:"person" json:"-" yaml:"-"`
	FirstName       string   `xml:"firstName" json:"firstName" yaml:"firstName"`
	LastName        string   `xml:"lastName" json:"lastName" yaml:"lastName"`
	MiddleName      string   `xml:"middleName,omitempty" yaml:"middleName,omitempty" yaml:"middleName,omitempty"`
	LastNameAtBirth string   `xml:"lastNameAtBirth,omitempty" yaml:"lastNameAtBirth,omitempty" yaml:"lastNameAtBirth,omitempty"`
}

type Hierarchy struct {
	XMLName  xml.Name `xml:"hierarchy" json:"-" yaml:"-"`
	IsMaster bool     `xml:"isMaster" json:"isMaster" yaml:"isMaster"`
	MasterId string   `xml:"masterId,omitempty" json:"masterId,omitempty" yaml:"masterId,omitempty"`
}
