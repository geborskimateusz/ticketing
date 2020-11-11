#!/bin/bash

POST_ID=$1

curl -b ./cookie.txt --insecure --header "Content-Type: application/json" \
  --request GET \
   https://we-creators.dev/api/items/${POST_ID}

