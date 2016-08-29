#!/bin/bash

filename="iptc_subjectcode.json"
if [ ! -f $filename ]; then
	curl https://api.aylien.com/api/v1/classify/taxonomy/iptc-subjectcode > $filename
fi
go run importTaxonomy.go > iptc_subjectcode.sql
