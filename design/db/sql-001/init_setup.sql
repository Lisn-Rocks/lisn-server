-- To apply these changes, run:
-- psql -f init_setup.sql Lisn



CREATE TABLE artists (
	artistid	SERIAL		PRIMARY KEY,
	name		TEXT
);


-- Song files and album cover images are stored in the filesystem in folders
-- specified in config. Their filename is of format <id>.<extension>

CREATE TABLE albums (
	albumid		SERIAL		PRIMARY KEY,
	name		TEXT,
	artistid	INTEGER		REFERENCES artists,
	extension	TEXT
);


CREATE TABLE songs (
	songid 		SERIAL 		PRIMARY KEY,
	name		TEXT,
	artistid	INTEGER		REFERENCES artists,
	genre		TEXT,
	albumid		INTEGER		REFERENCES albums,
	extension	TEXT
);
