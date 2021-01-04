#!/bin/bash

BODY="$(cat /dev/urandom | tr -dc 'a-zA-Z0-9' | fold -w 10 | head -n 1)"
MAIL="${BODY}@mail.xy"
echo "$MAIL"

