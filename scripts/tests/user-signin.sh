#!/bin/bash

IFS=,
CREDENTIALS=`cat test-credentials.txt`
read EMAIL PASSWORD <<< $CREDENTIALS

echo "Signin as: ";
echo $EMAIL; 
echo $PASSWORD; 

curl -c ./cookie.txt --insecure --header "Content-Type: application/json" \
  --request POST \
  --data '{"email":"'"$EMAIL"'","password":"'"$PASSWORD"'"}' \
 https://ticketing.dev/api/users/signin 
