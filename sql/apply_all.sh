#!/usr/bin/bash

#
# Use this bash script to apply all SQL migrations in this folder to a newly
# created PostgreSQL database.
#
# NOTICE:
# You are supposed to create a PostgreSQL database first.
#
# To create database use the following command:
#     CREATE DATABASE <database name>
#
# Run this script as follows (replace <database name> with the name of the
# datbase you have created):
#     ./apply_all <database name>
#
# NOTICE:
# If this file does not execute as expected, make sure that it has
# executable privileges.
#
# As a quickfix, run the following:
#     chmod +x apply_all
#     ./apply_all <database name>
#
# Alternatively, try:
#     bash apply_all <database name>
#

DBNAME=$1

for file in *.sql; do
    psql -f $file $DBNAME
done
