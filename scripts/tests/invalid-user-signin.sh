#!/bin/bash

EMAIL="VsDb3ZLlwJ@mail.xy"
PASSWORD="in12212valid"

curl -c ./cookie.txt --insecure --header "Content-Type: application/json" \
  --request POST \
  --data '{"email":"'"$EMAIL"'","password":"'"$PASSWORD"'"}' \
 https://ticketing.dev/api/users/signin | jq '.'
