#!/bin/bash
# Copyright (c) 2021 Red Hat, Inc.
# Copyright Contributors to the Open Cluster Management project

openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout ./opt/app-root/search-api/sslcert/searchapi.key -out ./opt/app-root/search-api/sslcert/searchapi.crt -config ./opt/app-root/search-api/sslcert/req.conf -extensions 'v3_req'
