#!/bin/sh

## Example startup script for Mystic BBS (pi)
##
## 1. Extract files to a location, like /mystic/doors/dod
## 2. Goto that dir and type 'chmod +x dod-linux-armv6 launch.sh' 
## 3. Add to Door Menu in ./mystic -cfg  (Editors > Menu Editor > default > Doors):
##
##      Command     (DD) Exec External Program                       
##      Data        /mystic/doors/dod/launch.sh %N

cd ~/mystic/doors/dod
./dod-linux-armv6 -path ~/mystic/temp$1/
