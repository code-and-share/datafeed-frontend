#!/bin/bash

#Postgresql profile
export DB_NAME_FILE="secrets/db_feed_name"
export DB_USER_FILE="secrets/db_feed_user"
export DB_PASS_FILE="secrets/db_feed_pass"
export DB_PSQL_HOST="0.0.0.0"
export DB_PSQL_PORT="5432"

# other
export PORT="9000"
export OWN_URL="http://0.0.0.0:9001"
export FILES_SOURCE="http://0.0.0.0:8080/files/"


echo "All set!"
