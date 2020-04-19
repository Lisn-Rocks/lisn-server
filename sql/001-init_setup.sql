-- To apply these changes, run:
-- psql -f 001-init_setup.sql <database name>


CREATE TABLE artists (
    artistid    SERIAL      PRIMARY KEY,
    name        TEXT        NOT NULL        UNIQUE
);


-- Song files and album cover images are stored in the filesystem in folders
-- specified in config. Their filename is of format <id><extension>

CREATE TABLE albums (
    albumid     SERIAL      PRIMARY KEY,
    name        TEXT        NOT NULL,
    artistid    INTEGER     NOT NULL        REFERENCES artists,
    extension   TEXT        NOT NULL
);


CREATE TABLE songs (
    songid      SERIAL      PRIMARY KEY,
    name        TEXT        NOT NULL,
    artistid    INTEGER     NOT NULL        REFERENCES artists,
    genre       TEXT        NOT NULL,
    albumid     INTEGER     NOT NULL        REFERENCES albums,
    extension   TEXT        NOT NULL
);
