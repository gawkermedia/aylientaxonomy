
CREATE TABLE iptc_subjectcode (
	id CHAR(8) PRIMARY KEY,
	label VARCHAR(60) NOT NULL,
	parent CHAR(8) NOT NULL,
	mainCategory CHAR(8) NOT NULL
);


