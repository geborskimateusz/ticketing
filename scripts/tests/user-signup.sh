#!/bin/bash

EMAIL=$(source ./random_email_generator.sh)
echo "$EMAIL"
curl -c ./cookie.txt --insecure --header "Content-Type: application/json" \
  --request POST \
  --data '{"email":"'"$EMAIL"'","password":"validPassword"}' \
 https://ticketing.dev/api/users/signup | jq '.'
