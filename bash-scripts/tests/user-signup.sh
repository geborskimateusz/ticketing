#!/bin/bash
curl --insecure --header "Content-Type: application/json" \
  --request POST \
  --data '{"email":"valid@email.com","password":"validPassword"}' \
  http://localhost:8080/api/users/signup
  #https:/ticketing.dev/api/users/signup 
