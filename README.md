<!-- [![dod1.png](https://i.postimg.cc/bdXzNVQ4/dod1.png)](https://postimg.cc/4H5CW521) -->

# door-of-doors

A Linux-based BBS door (for modern "old school" style BBSs like Mystic, Synchronet, ENiGMA½ or Talisman) that aggregates 3 popular door servers into a single menu system. It has lightbar driven menus, 100% customizable/replaceable ANSI art and a built-in statistics screen to show most popular doors and recent plays. An ANSI capable terminal program (like SyncTerm, Maigterm or Netrunner), at 80x25 screen size or larger, is required.

It's been tested on Linux i386, amd64, armv6 and armv7 (e.g. Raspberry Pi). Windows OS is not supported.

### Door Servers:
- [Gold Mine](http://goldminebbs.com)
- [BBSLink](https://bbslink.net/)
- [Door Party](http://wiki.throwbackbbs.com/doku.php?id=start)

# How it works

Direct launch door codes for each server are stored in a sqlite3 database. When a user selects a door, it launches the _external connection shell script_ provided by each door server and uses standard telnet/rlogin to pass the site/user credentials to connect the user to the door server.

# Installation options

### Option 1: Grab the release folder

Easiest way to get up and running. To download the latest, grab everything in the `release` folder.

### Option 2: Build from source

If you're a Go developer and want to contribute: fork this repo, make some changes, then submit a pull request! 

Note: the contents of `/release` should be added to the root of the door directory, as it contains all the static files necessary to run the door.

The included Makefile can detect your platform and "build down" from there -- e.g., if you are on ARM64, it'll generate 32-bit and 64-bit versions. For Pi, it can handle armv6, armv7 and armv8. Simple select the version you want to used. You don't have to use this.

Important: `go-sqlite3` is cgo package and you'll need gcc installed (e.g. `sudo apt install build-essentials`). However, after you have built and installed go-sqlite3 with `go install github.com/mattn/go-sqlite3` (which requires gcc), you can build your app without relying on gcc in future.

# Door Server Setup

Note, Door-of-Doors requires that you are a member of each door server (e.g. you have the credentials issued by the door server owners).

Place your BBSLink and Door Party connection scripts in the `/servers` directory, making sure the correct paths are set in config.ini. Gold Mine's conenction is script is provided in the `/servers/release/goldmine` folder

Note, Door Party requires the "door-party-connector" app to be configured and running.

# Configuration

There are instructions are in `/release/config.ini`. You'll need to edit this file in order for it to all work. There's also a sample 'launch.sh' in /release a that shows how you might launch this from a BBS, like Mystic.

I've also included a toggle for "Adult" type gamnes, as maybe some folks won't want some of those racy doors on their BBS.

# To-Do

- sort and filter doors?
- save to some sort of favorites list?
- maybe some sort of grafitti wall?
- implement a better timer/time-out
- wide-screen support
- loadable fonts
