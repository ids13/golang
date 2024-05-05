package main

import (
	"bytes"
	"encoding/base64"
	"encoding/csv"
	"encoding/hex"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/url"
)

func main() {
	// JSON Encoding and Decoding
	type Person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	person := Person{Name: "John", Age: 30}
	// Encoding to JSON
	jsonData, _ := json.Marshal(person)
	fmt.Println("JSON Encoding:", string(jsonData))

	// Decoding from JSON
	var newPerson Person
	json.Unmarshal(jsonData, &newPerson)
	fmt.Println("JSON Decoding - Name:", newPerson.Name)
	fmt.Println("JSON Decoding - Age:", newPerson.Age)

	// XML Encoding and Decoding
	type Address struct {
		City  string `xml:"city"`
		State string `xml:"state"`
	}
	address := Address{City: "New York", State: "NY"}
	// Encoding to XML
	xmlData, _ := xml.Marshal(address)
	fmt.Println("XML Encoding:", string(xmlData))

	// Decoding from XML
	var newAddress Address
	xml.Unmarshal(xmlData, &newAddress)
	fmt.Println("XML Decoding - City:", newAddress.City)
	fmt.Println("XML Decoding - State:", newAddress.State)

	// CSV Data as a slice of slices
	csvData := [][]string{{"Name", "Age"}, {"John", "30"}, {"Jane", "25"}}

	// CSV Encoding
	var csvBuffer bytes.Buffer
	csvWriter := csv.NewWriter(&csvBuffer)
	csvWriter.WriteAll(csvData)
	csvWriter.Flush()
	fmt.Println("CSV Writing:")
	fmt.Println(csvBuffer.String())

	// CSV Decoding
	csvReader := csv.NewReader(&csvBuffer)
	records, _ := csvReader.ReadAll()
	fmt.Println("CSV Reading:", records)

	// Encoding and Decoding Base64
	data := "Hello, world!"
	// Encoding to Base64
	encodedData := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println("Base64 Encoding:", encodedData)

	// Decoding from Base64
	decodedData, _ := base64.StdEncoding.DecodeString(encodedData)
	fmt.Println("Base64 Decoding:", string(decodedData))

	// Encoding and Decoding Hex
	hexData := []byte("hello")
	// Encoding to Hex
	encodedHex := hex.EncodeToString(hexData)
	fmt.Println("Hex Encoding:", encodedHex)

	// Decoding from Hex
	decodedHex, _ := hex.DecodeString(encodedHex)
	fmt.Println("Hex Decoding:", string(decodedHex))

	// Encoding and Decoding URL
	urlString := "https://www.example.com/?q=hello world"
	// Encoding to URL
	encodedURL := url.QueryEscape(urlString)
	fmt.Println("URL Encoding:", encodedURL)

	// Decoding from URL
	decodedURL, _ := url.QueryUnescape(encodedURL)
	fmt.Println("URL Decoding:", decodedURL)
}
