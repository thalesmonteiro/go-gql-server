#!/bin/sh
srcPath="cmd"
pkgFile="main.go"
app="gql-server"
src="$srcPath/$app/$pkgFile"

printf "\nStart running: $app\n"
# Set all ENV vars for the server to run
time go run $src
# This should unset all the ENV vars, just in case.
printf "\nStopped running: $app\n\n"