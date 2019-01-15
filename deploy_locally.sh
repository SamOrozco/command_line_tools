#!/usr/bin/env bash
toolDir=~/.orozco/
mkdir $toolDir
go build -o orozco main/main.go
cp orozco $toolDir