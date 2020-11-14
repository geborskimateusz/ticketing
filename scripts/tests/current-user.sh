#!/bin/bash
curl -b ./cookie.txt --insecure --header "Content-Type: application/json" \
  --request GET \
   https://ticketing.dev/api/users/currentuser

