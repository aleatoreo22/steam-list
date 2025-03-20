#!/bin/bash
cd steam-list-api
go get
go build
steam-list-api.com&
cd ../steam-list-ui/frontend
npm i
npm run preview&
