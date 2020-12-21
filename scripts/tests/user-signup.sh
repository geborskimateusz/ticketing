#!/bin/bash
curl --insecure --header "Content-Type: application/json" \
  --request POST \
  --data '{"email":"new@email.com","password":"validPassword"}' \
 http:/ticketing.dev/api/users/signup 

 #http://localhost:8080/api/users/signup
