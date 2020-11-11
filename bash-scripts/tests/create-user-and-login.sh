#!/bin/bash

echo "Creating an user."
./user-signup.sh
echo "Signin.."
./user-signin.sh
echo "Checking current user."
./current-user.sh

