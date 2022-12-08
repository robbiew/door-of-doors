<!-- [![dod1.png](https://i.postimg.cc/bdXzNVQ4/dod1.png)](https://postimg.cc/4H5CW521) -->

# door-of-doors

A Linux-based BBS door that aggregates 3 popular door servers into a single menu system. It has lightbar driven menus, customizable ANSI art and a built-in statistics screen. An ANSI capable terminal program (like SyncTerm, Maigterm or Netrunner), at 80x25 screen size, is required.

It's been tested on Linux i386, amd64, armv6 and armv7 (e.g. Raspberry Pi). Windows OS is not supported.

## Door Servers:
- [Gold Mine]()
- [BBSLink]()
- [Door Party]()

# How it works

Direct launch door codes for each server are stored in a sqlite3 database. When a user selects a door, it launches the _external connection shell script_ provided by each door server and uses standard telnet/rlogin to pass the site/user credentials to connect the user to the door server.

# Installation

You can grab a full release -- just download the zip file for your platform from RELEASES -- or build it yourself, if you're handy with Go.

# Build from Go source

Note, if you build yourself, the contents of /release should be added to the root of the door directory.

The Makefile will detect your platform and "build down" from there -- e.g., if you are on ARM64, it'll generate 32-bit and 64-bit versions. For Pi, it can handle armv6, armv7 and armv8. Simple select the version you want to used.

go-sqlite3 is cgo package. If you want to build your app using go-sqlite3, you need gcc. However, after you have built and installed go-sqlite3 with `go install github.com/mattn/go-sqlite3` (which requires gcc), you can build your app without relying on gcc in future.

# Door Servers
Note, Door-of-Doors requires that you are a member of each door server (e.g. you have the credentials issued by the door server owners).

Place your BBSLink and Door Party connection scripts in the /servers directory, making sure the correct paths are set in config.ini. Gold Mine's conenction is script is provided in the /servers/release/goldmine folder

Note, Door Party requires the "door-party-connector" app to be configured and running.

# Config

There are instructions are in /release/config.ini. You'll need to edit this file in order for it to all work. There's also a sample 'launch.sh' in /release a that shows how you might launch this from a BBS, like Mystic.

I've also included a toggle for "Adult" type gamnes, as maybe some folks won't want some of those racy doors on their BBS.

# To-Do

- sort and filter doors?
- save to some sort of favorites list?
- maybe some sort of grafitti wall?
- implement a better timer/time-out
