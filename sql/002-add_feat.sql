-- Instead of including featured artists into the song name like this:
-- "The End (feat. GLC, Chip Tha Ripper, Nicole Wray)"
-- we are going to use an array of integers that reference artisid of those
-- featured artists.
--
-- The only problem is that PostgreSQL cannot enforce this constraint, so we
-- will have to check and enforce it ourselves.

-- To apply these changes, run:
-- psql -f 002-add_feat.sql <database name>

ALTER TABLE songs ADD feat INTEGER ARRAY;
