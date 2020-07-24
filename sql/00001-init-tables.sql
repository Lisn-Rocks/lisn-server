-- To apply these changes, run:
-- psql -f 001-init_setup.sql <database name>

-- Song files and album cover images are stored in the filesystem in folders
-- specified in config. Their filename is of format <id><extension>

create table artists
(
	artistid serial not null,
	artist text not null
);

create unique index artists_artistid_uindex
	on artists (artistid);

alter table artists
	add constraint artists_pk
		primary key (artistid);

create table albums
(
	albumid serial not null,
	album text not null,
	artistid int not null
		constraint artist
			references artists
				on delete cascade,
	coverext text default '.jpg' not null
);

create unique index albums_albumid_uindex
	on albums (albumid);

alter table albums
	add constraint albums_pk
		primary key (albumid);

create table songs
(
	songid serial not null,
	song text not null,
	albumid int not null
		constraint album
			references albums
				on delete cascade,
	audioext text default '.mp3' not null
);

create unique index songs_songid_uindex
	on songs (songid);

alter table songs
	add constraint songs_pk
		primary key (songid);

create table feats
(
	songid int not null
		constraint song
			references songs
				on delete cascade,
	featid int not null
		constraint feat
			references artists
				on delete cascade
);

create table genres
(
	albumid int not null
		constraint album
			references albums
				on delete cascade,
	genre text not null
);
