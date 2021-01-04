#!/bin/bash

EMAIL=$(source ./random_email_generator.sh)
echo "$EMAIL"
curl --insecure --header "Content-Type: application/json" \
  --request POST \
  --data '{"email":'"$EMAIL"',"password":"validPassword"}' \
 http:/ticketing.dev/api/users/signup 

 #http://localhost:8080/api/users/signup
