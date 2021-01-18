#!/bin/bash

EMAIL=$(source ./random_email_generator.sh)
PASSWORD="validPassword"
echo "${EMAIL},${PASSWORD}" > test-credentials.txt

curl -c ./cookie.txt --insecure --header "Content-Type: application/json" \
  --request POST \
  --data '{"email":"'"$EMAIL"'","password":"'"$PASSWORD"'"}' \
 https://ticketing.dev/api/users/signup | jq '.'
