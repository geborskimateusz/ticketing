#!/bin/bash
curl -c cookie.txt --insecure --header "Content-Type: application/json" \
  --request POST \
   https://ticketing.dev/api/users/signout

