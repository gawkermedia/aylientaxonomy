# Aylien Taxonomy importer

It downloads the Aylien's taxonomy database from
https://api.aylien.com/api/v1/classify/taxonomy/iptc-subjectcode and generates
`INSERT INTO` SQL commands from it. You can see the table definition in the
`iptc-subjectcode-ddl.sql` file.

## How to run

`$ ./run.sh`


