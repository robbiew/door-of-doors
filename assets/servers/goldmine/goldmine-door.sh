#!/bin/bash

# Script for connecting to Gold Mine rlogin
# Install rsh-redone-client: "sudo apt install rsh-redone-client"
# You shouldn't need to edit this script
# Makre sure config.ini is filled out properly

export TERM=linux

port=$5
host=$4
prefix="["$2"]"
user=${1//_/ }

TERM=$3 rlogin -p $port -l "$prefix$user" $host

