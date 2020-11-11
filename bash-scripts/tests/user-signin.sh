#!/bin/bash
curl -c ./cookie.txt --insecure --header "Content-Type: application/json" \
  --request POST \
  --data '{"email":"valid@email.com","password":"validPassword"}' \
  https://we-creators.dev/api/users/signin 
