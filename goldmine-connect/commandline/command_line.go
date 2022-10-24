package commandline

import "time"

// CommandLine type represents options read from command line arguments.
type CommandLine struct {
	host    string
	port    uint64
	name    string
	tag     string
	xtrn    string
	timeout time.Duration
}

// Host method returns a given host.
func (c *CommandLine) Host() string {
	return c.host
}

// Port method returns a given port.
func (c *CommandLine) Port() uint64 {
	return c.port
}

// Name method returns a BBS handle.
func (c *CommandLine) Name() string {
	return c.name
}

// From method returns a given source BBS name.
func (c *CommandLine) Tag() string {
	return c.tag
}

// Xtrn method returns a given xtrn id.
func (c *CommandLine) Xtrn() string {
	return c.xtrn
}

// Timeout method returns a given server response timeout after EOF of input file.
func (c *CommandLine) Timeout() time.Duration {
	return c.timeout
}
