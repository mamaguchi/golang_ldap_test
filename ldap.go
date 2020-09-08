package main

import (
	"github.com/go-ldap/ldap/v3"
        "log"
        "fmt"
)

func delete() {
    l, err := ldap.DialURL("ldap://127.0.0.1:389")
    if err != nil {
	    log.Fatal(err)
    }
    defer l.Close()

    err = l.Bind("cn=Directory Manager", "88motherfaker88")
    if err != nil {
	    log.Fatal(err)
    }

    del := ldap.NewDelRequest("clinicName=kk_maran,ou=clinic,ou=primhc,ou=district,ou=state,ou=kkm,dc=example,dc=com", nil)

    err = l.Del(del)
    if err != nil {
	    log.Fatal(err)
    }
}

func update() {
    l, err := ldap.DialURL("ldap://127.0.0.1:389")
    if err != nil {
	    log.Fatal(err)
    }
    defer l.Close()

    err = l.Bind("cn=Directory Manager", "88motherfaker88")
    if err != nil {
	    log.Fatal(err)
    }

    modify := ldap.NewModifyRequest("clinicName=kk_maran,ou=clinic,ou=primhc,ou=district,ou=state,ou=kkm,dc=example,dc=com", nil)
    modify.Replace("clinicState", []string{"selangor"})

    err = l.Modify(modify)
    if err != nil {
	    log.Fatal(err)
    }
}

func add() {
    l, err := ldap.DialURL("ldap://127.0.0.1:389")
    if err != nil {
	    log.Fatal(err)
    }
    defer l.Close()

    err = l.Bind("cn=Directory Manager", "88motherfaker88")
    if err != nil {
	    log.Fatal(err)
    }

    addRequest := ldap.NewAddRequest(
	    "clinicName=kk_maran,ou=clinic,ou=primhc,ou=district,ou=state,ou=kkm,dc=example,dc=com",
	    nil,
    )
    addRequest.Attribute("objectClass", []string{"top"})
    addRequest.Attribute("objectClass", []string{"mytcaClinic"})
    addRequest.Attribute("clinicName", []string{"kk_maran"})
    addRequest.Attribute("clinicState", []string{"pahang"})
    addRequest.Attribute("clinicDistrict", []string{"maran"})

    err = l.Add(addRequest)
    if err != nil {
	    log.Fatal(err)
    }
}

func search() {
    l, err := ldap.DialURL("ldap://127.0.0.1:389")
    if err != nil {
            log.Fatal(err)
    }
    defer l.Close()

    searchRequest := ldap.NewSearchRequest(
	    "dc=example,dc=com",
	    ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0,0, false,
	    "(ou=patient)",
	    []string{"ou"},
	    nil,
    )

    sr, err := l.Search(searchRequest)
    if err != nil {
	    log.Fatal(err)
    }

    for _, entry := range sr.Entries {
	    fmt.Printf("%s: %v\n", entry.DN, entry.GetAttributeValue("ou"))
    }
}

func main() {
	//search()
	//add()
	//update()
	delete()
}


