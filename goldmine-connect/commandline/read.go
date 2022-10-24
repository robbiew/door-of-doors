package commandline

import (
	"os"

	"gopkg.in/alecthomas/kingpin.v2"
)

// Read method returns valid options read from command line args.
func Read() *CommandLine {
	host := kingpin.Arg("host", "GoldMine host adddress").Required().String()
	port := kingpin.Arg("port", "Goldmine rlogin port").Required().Uint64()
	name := kingpin.Arg("name", "username").Required().String()
	tag := kingpin.Arg("tag", "BBS tag (no brackets)").Required().String()
	xtrn := kingpin.Arg("xtrn", "Gold Mine xtrn code").Required().String()
	timeout := kingpin.Flag("timeout", "Byte receiving timeout after the input EOF occurs").Short('t').Default("1s").Duration()

	kingpin.UsageTemplate(kingpin.CompactUsageTemplate).Version("1.0").Author("Forked by aLPHA from Marcin Tojek")
	kingpin.CommandLine.Name = "goldmine-connect"
	kingpin.CommandLine.Help = "Read bytes from stdin and pass them to the remote host."

	kingpin.Parse()

	return &CommandLine{
		host:    *host,
		port:    *port,
		name:    *name,
		tag:     *tag,
		xtrn:    *xtrn,
		timeout: *timeout,
	}
}

// SetCommandLineArgs method changes earlier set arguments (use only in debug).
func SetCommandLineArgs(customArguments ...string) {
	os.Args = os.Args[0:1] // leave only app path
	for _, customArgument := range customArguments {
		os.Args = append(os.Args, customArgument)
	}
}
