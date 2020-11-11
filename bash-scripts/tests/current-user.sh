#!/bin/bash
curl -b ./cookie.txt --insecure --header "Content-Type: application/json" \
  --request GET \
   https://we-creators.dev/api/users/currentuser

