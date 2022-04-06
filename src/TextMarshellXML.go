package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"strings"
)

type Size int

const (
	Unrecognized Size = iota
	Small
	Large
)

func (s *Size) UnmarshalText(text []byte) error {
	switch strings.ToLower(string(text)) {
	default:
		*s = Unrecognized
	case "small":
		*s = Small
	case "large":
		*s = Large
	}
	return nil
}

func (s Size) MarshellText() ([]byte, error) {
	var name string
	switch s {
	default:
		name = "unrecognized"
	case Small:
		name = "small"
	case Large:
		name = "large"
	}
	return []byte(name), nil
}

// https://knowledge-base.secureflag.com/vulnerabilities/xml_injection/xml_entity_expansion_go_lang.html
func main() {
	// blob := `
    // <!DOCTYPE d [<!ENTITY xxe SYSTEM "file:///Test/popup.ps1">]>
	// <sizes>
	// 	<size>regular</size>
	// 	<size>large</size>
	// 	<size>unrecognized</size>
	// 	<size>small</size>
	// 	<size>normal</size>
	// 	<size>small</size>
	// 	<size>large</size>
	// 	<size>&xxe;</size>
	// </sizes>
	// `
	blob := `
	<sizes>
		<size>regular</size>
		<size>large</size>
		<size>unrecognized</size>
		<size>small</size>
		<size>normal</size>
		<size>small</size>
		<size>large</size>
	</sizes>
	`
	var inventory struct {
		Sizes []Size `xml:"size"`
	}
	if err := xml.Unmarshal([]byte(blob), &inventory); err != nil {
		log.Fatal(err)
	}
	counts := make(map[Size]int)
	for _, size := range inventory.Sizes {
		counts[size] += 1
	}
	fmt.Printf("Inventory Counts:\n* Small:         %d\n* Large:         %d\n* Unrecognized:  %d", counts[Small], counts[Large], counts[Unrecognized])
}