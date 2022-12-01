[![dod1.png](https://i.postimg.cc/bdXzNVQ4/dod1.png)](https://postimg.cc/4H5CW521)

# door-of-doors

This is a Linux-based BBS door that aggregates 3 popular door servers into a single menu system that runs as an external door, utilizing door32.sys. It currently includes Gold Mine (my own door server), BBS Link and Door Party. It's a convenient way to quickly add 200+ door games to a BBS -- no need to create menus or add each door one by one. ANSI art files are included for menu customization. I've started to add door descriptions, but there's a LOT to document here, so it's currently mostly sparse.

It also contains an ANSI screen to show most popular doors, and a user log.

It's been tested on Linux amd64 and armv7 (e.g. Raspberry Pi). 

# How it works

Direct launch door codes for each server are stored in a sqlite3 database. When a user selects a door, it launches the *external connection shell script* provided by each door server and uses standard telnet/rlogin to pass the site/user credentials to connect the user to the door server.

# Installation

You can grab a full release -- just download the zip file for your platform from RELEASES -- or build it yourself, if you're handy with Go.

Note, if you build yourself, the contents of /assets should be added to the root of the door directory.

The Makefile will detect your platform and "build down" from there -- e.g., if you are on ARM64, it'll generate 32-bit and 64-bit versions. For Pi, it can handle armv6-armv8.

Place your door server connection scripts in the /servers directory, make sure paths are set in config.ini Note, Door Party requires the "door-party-connector" app to be configured and running.

Instructions are in assets/config.ini. You'll need to edit this file in order for it to all work. 
There's a sample 'launch.sh' in /assets as well that shows how you might launch from a BBS, like Mystic.

Note, Door-of-Doors requires that you are a member of each door server (e.g. you have the credentials issued by the door server owners). You can also toggle individual door servers off/on if you don't use a particular one.

I've also included a toggle for "Adult" type gamnes, as I'm sure some folks won't want some of those on their BBS.

# To-Do

- sort and filter doors
- save to some sort of favorites list
- maybe some sort of grafitti wall
- toggle the Stats screen
- implement a better timer/time-out
