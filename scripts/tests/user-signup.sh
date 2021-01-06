#!/bin/bash

EMAIL=$(source ./random_email_generator.sh)
echo "$EMAIL"
curl -c, --cookie-jar ./cookie.txt --insecure --header "Content-Type: application/json" \
  --request POST \
  --data '{"email":"'"$EMAIL"'","password":"validPassword"}' \
 https:/ticketing.dev/api/users/signup 

 #http://localhost:8080/api/users/signup
