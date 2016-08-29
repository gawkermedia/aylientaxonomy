package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"strings"
)

type Link struct {
	Rel  string
	Link string
}

type Term struct {
	Id       string
	Taxonomy string
	Label    string
	Links    []Link
}

func UrlToId(input string) (string, error) {
	if input == "" {
		return input, errors.New("The URL is empty")
	}
	u, err := url.Parse(input)
	if err != nil {
		return "", err
	}
	parts := strings.Split(u.Path, "/")
	if len(parts) < 1 {
		return "", errors.New("Missing Id")
	}
	return parts[len(parts)-1], err
}

func GetParentLinkOrSelf(links []Link) string {
	if len(links) > 1 {
		for _, link := range links {
			if link.Rel == "parent" {
				return link.Link
			}
		}
	}
	return links[0].Link
}

// recursive function
func LookupForMainCategory(keys map[string]Term, parentId string) string {
	id, err := UrlToId(GetParentLinkOrSelf(keys[parentId].Links))
	if err != nil {
		log.Fatal(err)
	}
	if id == parentId {
		return id
	} else {
		return LookupForMainCategory(keys, id)
	}
}

func main() {
	var terms []Term
	tableName := "iptc_subjectcode"
	filename := tableName + ".json"
	keys := make(map[string]Term)
	buff, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(buff, &terms)
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range terms {
		keys[item.Id] = item
	}
	for id, item := range keys {
		parentId, err := UrlToId(GetParentLinkOrSelf(item.Links))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("INSERT INTO %s VALUES (\"%s\", \"%s\", \"%s\", \"%s\");\n", tableName, id, item.Label, parentId, LookupForMainCategory(keys, parentId))
	}
}
