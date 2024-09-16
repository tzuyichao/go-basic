package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/go-ldap/ldap/v3"
)

func main() {
	ldapURL := flag.String("url", "ldap://yourldapserver:389", "LDAP Server URL")
	bindDN := flag.String("bindDN", "cn=admin,dc=example,dc=com", "Bind DN")
	password := flag.String("password", "", "Password for the Bind DN")
	searchBase := flag.String("base", "dc=example,dc=com", "Search base DN")
	query := flag.String("query", "(objectClass=person)", "LDAP search filter")

	flag.Parse()
	l, err := ldap.DialURL(*ldapURL)
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	err = l.Bind(*bindDN, *password)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully authenticated")
	searchRequest := ldap.NewSearchRequest(
		*searchBase,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		*query,
		[]string{"dn", "cn", "mial"},
		nil,
	)
	sr, err := l.Search(searchRequest)
	if err != nil {
		log.Fatal(err)
	}
	for _, entry := range sr.Entries {
		fmt.Printf("DN: %s\n", entry.DN)
		fmt.Printf("CN: %s\n", entry.GetAttributeValue("cn"))
		fmt.Printf("Email: %s\n", entry.GetAttributeValue("mail"))
	}
	fmt.Println("Done.")
}
