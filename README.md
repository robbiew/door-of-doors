# door-of-doors

This is a Linux-based BBS door that aggregates 3 popular door servers into a single menu system: Gold Mine, BBS Link and Door Party. It's a convenient way to quickly add 200+ door games to your BBS -- no need to create menus or add each door one by one. ANSI art files are included for customization. I've started to add door descriptions, but there's a LOT to document here.

It's been tested on Linux amd64 and armv7 (e.g. Raspberry Pi).

# How it works

Direct launch door codes for each server are stored in a sqlite3 database. When a user selects a door, it launches the external connection shell script and uses telnet/rlogin to pass the site credentials and connect the user to the door server.

# Installation

You can grab a full release -- just download the zip file -- or build yourself, if you've got Go installed.
Note, if you build yourself, the contents of /assets should be added to the root of the door directory.
Place your connections scripts in the /servers directory.
Door Party requires the door-party-connector app to be configured and running.

Instructions are in the assets/config.ini. You'll need to edit this file in order for it to work.
There's a sample 'launch.sh' in /assets as well.

Note, Door-of-Doors requires that you are a member (you have the credentials issued by the door server owners).

# To-Do

- sort and filter
- save to favorites
- some sort of grafitti wall
