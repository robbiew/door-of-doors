#! /bin/bash
#
# Modified by aLPHA for Door of Doors, October 2022
#
# Original script credits:
#
# *******************************************************
# * BBSlink.net InterBBS Door Server Connection Script	*
# *******************************************************
# *  Created by: Dennis Ayala			  	            *
# *  Virtual Realities BBS				                *
# *  telnet://bbs.virtualrealitiesbbs.com		        *
# *******************************************************
#
#  Version 1.0.4  19th December 2015
#
#  (C)2015 Dennis Ayala for BBSlink.net.
#
#  Thanks go to haliphax for cryptography advice.
#

host=$1		
syscode=$2				   
authcode=$3			 		
schemecode=$4			

version="1.0.4"
screenrows="24"
scripttype="bash"
scriptver="$version"

if [ "$#" -lt 3 ]; then
	doorcode="$5"
	usernumber="$6"
	# * Generate random 32 character alphanumeric string 	*
	# * (upper and lowercase)				*
	xkey=$(cat /dev/urandom | tr -dc 'a-zA-Z0-9' | fold -w 6 | head -n 1)
	# * Get Token from BBSLink				*
	token=$(curl -s "http://$host/token.php?key=$xkey")

	xauth=$(echo -n "$authcode$token" | md5sum | awk '{print $1}')
	xcode=$(echo -n "$schemecode$token" | md5sum | awk '{print $1}')

	result=$(curl -s -H "X-User: $usernumber" -H "X-System: $syscode" -H "X-Auth: $xauth" -H "X-Code: $xcode" -H "X-Rows: $screenrows" -H "X-Key: $xkey" -H "X-Door: $doorcode" -H "X-Token: $token" -H "X-Type: $scripttype" -H "X-Version: $scriptver" "http://$host/auth.php?key=$xkey")

	if [ "$result" == "complete" ]; then
		export TERM=linux
		clear
		telnet -E -K -8 $host
	else
		echo "Error: $result"
		exit 1
	fi
else
    echo ""
    echo "Usage: bbslink.sh [Door Code] [User Number]"
    echo ""
    echo ""
	echo "Mystic BBS Example:"
	echo "bbslink.sh lord %#"
    echo ""
    exit 1
fi
exit 0