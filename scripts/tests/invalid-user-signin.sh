#!/bin/bash

EMAIL="johndoe"
PASSWORD="invalid"

curl -c ./cookie.txt --insecure --header "Content-Type: application/json" \
  --request POST \
  --data '{"email":"'"$EMAIL"'","password":"'"$PASSWORD"'"}' \
 https://ticketing.dev/api/users/signin | jq '.'
