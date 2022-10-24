# goldmine-connect

This is a telnet/rlogin program that connects to a Synchronet-based game server, called Gold Mine. Goldmine allows passwordless RLOGIN so pre-approved BBS sysops can pass theiir users through to Gold Mine. It is used primarily by the "Door of Doors" BBS program (also by robbiew/aLPHA).

What does it do? It read bytes from stdin and passes them to the remote Gold Mine host. The application works similarly to the old-school telnet application, but it lets you read bytes from standard input and wait for response.

### Usage

```
$ goldmine-connect --help
usage: goldmine-connect [<flags>] <host> <port> <name> <tag> <xtrn>

Read bytes from stdin and pass them to the remote host.

Flags:
  --help            Show help (also see --help-long and --help-man).
  -t, --timeout=1s  Byte receiving timeout after the input EOF occurs
  --version         Show application version.

Args:
  <host>  Gold Mine host adddress
  <port>  Gold Mine rlogin port
  <name>  Username to connect as
  <tag>   BBS tag (no brackets)
  <xtrn>  Gold Mine xtrn code
```
