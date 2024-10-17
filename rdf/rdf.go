package rdf

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

type RDF struct {
	XMLName xml.Name `xml:"RDF"`
	Project Project  `xml:"Project"`
}

type Project struct {
	XMLName       xml.Name `xml:"Project"`
	About         string   `xml:"about,attr"`
	Name          string   `xml:"name"`
	ShortDesc     string   `xml:"shortdesc"`
	Description   string   `xml:"description"`
	Homepage      string   `xml:"homepage"`
	BugDatabase   string   `xml:"bug-database"`
	Specification string   `xml:"Specification"`
	License       string   `xml:"license"`
	Maintainer    Person   `xml:"maintainer>Person"`
	Documenter    Person   `xml:"documenter>Person"`
}

type Person struct {
	XMLName  xml.Name `xml:"Person"`
	About    string   `xml:"about,attr"`
	Name     string   `xml:"name"`
	Homepage string   `xml:"homepage"`
	Mbox     string   `xml:"mbox"`
}

func PrintRdfInfo() {
	xmlFile, err := os.Open("doc/mock/doap.rdf")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer xmlFile.Close()

	byteValue, _ := io.ReadAll(xmlFile)

	var rdf RDF
	err = xml.Unmarshal(byteValue, &rdf)
	if err != nil {
		fmt.Println("Error unmarshalling XML:", err)
		return
	}

	fmt.Printf("Project Name: %s\n", rdf.Project.Name)
	fmt.Printf("Short Description: %s\n", rdf.Project.ShortDesc)
	fmt.Printf("Description: %s\n", rdf.Project.Description)
	fmt.Printf("Homepage: %s\n", rdf.Project.Homepage)
	fmt.Printf("Bug Database: %s\n", rdf.Project.BugDatabase)
	fmt.Printf("Specification: %s\n", rdf.Project.Specification)
	fmt.Printf("License: %s\n", rdf.Project.License)
	fmt.Printf("Maintainer Name: %s\n", rdf.Project.Maintainer.Name)
	fmt.Printf("Maintainer Homepage: %s\n", rdf.Project.Maintainer.Homepage)
	fmt.Printf("Maintainer Mbox: %s\n", rdf.Project.Maintainer.Mbox)
	fmt.Printf("Documenter Name: %s\n", rdf.Project.Documenter.Name)
}
