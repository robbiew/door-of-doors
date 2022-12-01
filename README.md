# door-of-doors

This is a Linux-based BBS door that aggregates 3 popular door servers into a single menu system: Gold Mine, BBS Link and Door Party. 
It's been tested on Linux amd64 and armv7 (e.g. Raspberry Pi).

# How it works
Direct launch door codes for each server are stored in a sqlite3 database. when a user selects a door, it launches the external connection shell script and uses telnet/rlogin to pass the site credentials and connect the user to the door server. 


You can grab a release, or build yourself.

Instructions are in the assts/config.ini.

Note, Door-of-Doors requires that you are a member (you have the credentials issued by the door server owners). 
