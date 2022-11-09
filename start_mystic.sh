#!/bin/sh

## Example startup script for Mystic BBS
##
## 1. Extract files to a location, like /mystic/doors/door-of-doors
## 2. Goto that dir and type 'chmod +x door-of-doors start_mystic.sh' 
## 3. Setup Door Menu in ./mystic -cfg  (Editors > Menu Editor > default > Doors):
##
##      Command     (DD) Exec External Program                       
##      Data        /mystic/doors/door-of-doors/start_mystic.sh %N


cd /mystic/doors/door-of-doors
./door-of-doors -path /mystic/temp$1/