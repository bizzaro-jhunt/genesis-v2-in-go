package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	. "github.com/jhunt/genesis/command"
	"github.com/pborman/getopt"
	"github.com/starkandwayne/goutils/ansi"
)

func require(good bool, msg string) {
	if !good {
		fmt.Fprintf(os.Stderr, "USAGE: %s ...\n", msg)
		os.Exit(1)
	}
}

func readall(in io.Reader) (string, error) {
	b, err := ioutil.ReadAll(in)
	return string(b), err
}

var (
	debug = false
)

var Version = ""

func main() {
	options := Options{
		Cwd:     getopt.StringLong("cwd", 'C', ".", "Effective working directory. Defaults to '.'"),
		Debug:   getopt.BoolLong("debug", 'D', "Enable debugging to print helpful messages about what Genesis is doing (for developers)."),
		Help:    getopt.BoolLong("help", 'h', "Show the help"),
		Trace:   getopt.BoolLong("trace", 'T', "Even more debugging, including debugging inside called tools like 'spruce' and 'bosh'."),
		Verbose: getopt.BoolLong("verbose",'v', "Enable debugging to print helpful messages about what Genesis is doing (for operators)."),
		Yes:     getopt.BoolLong("yes", 'y', "Answer 'yes' to all questions automatically."),
	}

	var command []string
	var opts = getopt.CommandLine
	args := os.Args
	for {
		opts.Parse(args)
		if opts.NArgs() == 0 {
			break
		}
		command = append(command, opts.Arg(0))
		args = opts.Args()
	}

	if len(command) == 0 {
		command = []string{"help"}
	}

	debug = *options.Debug

	c := NewCommand().With(options)

	c.HelpGroup("INFO:")
	c.Dispatch("help", "Get detailed help with a specific command",
		func(opts Options, args []string, help bool) error {
			if len(args) == 0 {
				buf := bytes.Buffer{}
				getopt.PrintUsage(&buf)
				ansi.Fprintf(os.Stderr, strings.Split(buf.String(), "\n")[0]+"\n")
				if Version == "" {
					ansi.Fprintf(os.Stderr, "genesis @*{development version 🦄}\n")
				} else {
					ansi.Fprintf(os.Stderr, "genesis v%s\n", Version)
				}
				ansi.Fprintf(os.Stderr, "USAGE: genesis [OPTIONS] COMMAND [MORE OPTIONS]\n")
				ansi.Fprintf(os.Stderr, "\n  OPTIONS\n")
				ansi.Fprintf(os.Stderr, "    -h, --help       Show this help screen.\n")
				ansi.Fprintf(os.Stderr, "    -D, --debug      Enable debugging, printing helpful message about what\n")
				ansi.Fprintf(os.Stderr, "                     Genesis is doing, to standard error.\n")
				ansi.Fprintf(os.Stderr, "    -T, --trace      Even more debugging, including debugging inside called\n")
				ansi.Fprintf(os.Stderr, "                     tools (like spruce and bosh).\n")
				ansi.Fprintf(os.Stderr, "    -C, --cwd        Effective working directory.  Defaults to '.'\n")
				ansi.Fprintf(os.Stderr, "    -y, --yes        Answer 'yes' to all questions, automatically.\n")
				ansi.Fprintf(os.Stderr, "\n\n  COMMANDS\n")
				ansi.Fprintf(os.Stderr, "    compile-kit      Create a distributable kit archive from dev.\n")
				ansi.Fprintf(os.Stderr, "    decompile-kit    Unpack a kit archive to dev.\n")
				ansi.Fprintf(os.Stderr, "    describe         Describe a Concourse pipeline, in words.\n")
				ansi.Fprintf(os.Stderr, "    download         Download a Genesis Kit from the Internet.\n")
				ansi.Fprintf(os.Stderr, "    graph            Draw a Concourse pipeline.\n")
				ansi.Fprintf(os.Stderr, "    init             Initialize a new Genesis deployment.\n")
				ansi.Fprintf(os.Stderr, "    lookup           Find a key set in environment manifests.\n")
				ansi.Fprintf(os.Stderr, "    manifest         Generate a redacted BOSH deployment manifest for an environment.\n")
				ansi.Fprintf(os.Stderr, "    new              Create a new Genesis deployment environment.\n")
				ansi.Fprintf(os.Stderr, "    ping             See if the genesis binary is a real thing.\n")
				ansi.Fprintf(os.Stderr, "    repipe           Configure a Concourse pipeline for automating deployments.\n")
				ansi.Fprintf(os.Stderr, "    secrets          Re-generate // rotate credentials (passwords, keys, etc.).\n")
				ansi.Fprintf(os.Stderr, "    summary          Print a summary of defined environments.\n")
				ansi.Fprintf(os.Stderr, "    version          Print the version of genesis\n")
				ansi.Fprintf(os.Stderr, "    yamls            Print a list of the YAML files used for a single environment.\n")
				ansi.Fprintf(os.Stderr, "\n  See 'genesis COMMAND -h' for more specific, per-command usage information.\n")
				return nil
			} else if args[0] == "help" {
				ansi.Fprintf(os.Stderr, "@R{This is getting a bit too meta, don't you think?}\n")
				return nil
			}
			return c.Help(args...)
		})
	c.Alias("usage", "help")

	/* genesis compile-kit */
	c.Dispatch("compile-kit", "Create a distributable kit archive from dev.",
		func(opts Options, args []string, help bool) error {
			if help {
				ansi.Fprintf(os.Stdout, "genesis v%s\n", Version)
				ansi.Fprintf(os.Stdout, "USAGE: genesis compile-kit -n NAME -v VERSION\n\n")
				ansi.Fprintf(os.Stdout, "OPTIONS\n")
				ansi.Fprintf(os.Stdout, "  -n, --name      Name of the kit archive.\n")
				ansi.Fprintf(os.Stdout, "  -v, --version   Version to package.\n")
				return nil
			}

			ansi.Fprintf(os.Stdout, "This is genesis compile-kit.")
			return nil
		})

	/* genesis decompile-kit */
	c.Dispatch("decompile-kit", "Unpack a kit archive to dev.",
		func(opts Options, args []string, help bool) error {
			if help {
				ansi.Fprintf(os.Stdout, "genesis v%s\n", Version)
				ansi.Fprintf(os.Stdout, "USAGE: genesis decompile-kit [NAME/VERSION | path/to/kit.tar.gz]\n\n")
				ansi.Fprintf(os.Stdout, "OPTIONS\n")
				ansi.Fprintf(os.Stdout, "  -f, --force  Overwrite dev/, if it exists.\n")
				return nil
			}

			ansi.Fprintf(os.Stdout, "This is genesis decompile-kit.")
			return nil
		})

	/* genesis describe */
	c.Dispatch("describe", "Describe a Concourse pipeline with words.",
		func(opts Options, args []string, help bool) error {
			if help {
				ansi.Fprintf(os.Stdout, "genesis v%s\n", Version)
				ansi.Fprintf(os.Stdout, "USAGE: genesis describe [pipeline-layout]\n\n")
				ansi.Fprintf(os.Stdout, "OPTIONS\n")
				ansi.Fprintf(os.Stdout, "  -c, --config     Path to the pipeline configuration file, which specifies\n")
				ansi.Fprintf(os.Stdout, "                   Git parameters, notification settings, pipeline layouts,\n")
				ansi.Fprintf(os.Stdout, "                   etc.  Defaults to 'ci.yml'\n")
				return nil
			}

			ansi.Fprintf(os.Stdout, "This is genesis describe.")
			return nil
		})

	/* genesis download */
	c.Dispatch("download", "Download a Genesis Kit from the Internet.",
		func(opts Options, args []string, help bool) error {
			if help {
				ansi.Fprintf(os.Stdout, "genesis v%s\n", Version)
				ansi.Fprintf(os.Stdout, "USAGE: genesis download NAME[/VERSION] [...]\n\n")
				ansi.Fprintf(os.Stdout, "OPTIONS\n")
				return nil
			}

			ansi.Fprintf(os.Stdout, "This is genesis download.")
			return nil
		})

	/* genesis graph */
	c.Dispatch("graph", "Draw a Concourse pipeline.",
		func(opts Options, args []string, help bool) error {
			if help {
				ansi.Fprintf(os.Stdout, "genesis v%s\n", Version)
				ansi.Fprintf(os.Stdout, "USAGE: genesis graph [pipeline-layout]\n\n")
				ansi.Fprintf(os.Stdout, "OPTIONS\n")
				ansi.Fprintf(os.Stdout, "  -c, --config     Path to the pipeline configuration file, which specifies\n")
				ansi.Fprintf(os.Stdout, "                   Git parameters, notification settings, pipeline layouts,\n")
				ansi.Fprintf(os.Stdout, "                   etc.  Defaults to 'ci.yml'\n")
				return nil
			}

			ansi.Fprintf(os.Stdout, "This is genesis graph.")
			return nil
		})

	/* genesis init */
	c.Dispatch("init", "Initialize a new Genesis deployment.",
		func(opts Options, args []string, help bool) error {
			if help {
				ansi.Fprintf(os.Stdout, "genesis v%s\n", Version)
				ansi.Fprintf(os.Stdout, "USAGE: genesis init [-k KIT/VERSION] name\n\n")
				ansi.Fprintf(os.Stdout, "OPTIONS\n")
				ansi.Fprintf(os.Stdout, "  -k, --kit        Name (and optionally, version) of the Genesis Kit to\n")
				ansi.Fprintf(os.Stdout, "                   base these deployments on.  i.e.: shield/6.3.0\n")
				return nil
			}

			ansi.Fprintf(os.Stdout, "This is genesis init.")
			return nil
		})

	/* genesis lookup */
	c.Dispatch("lookup", "Find a key set in environment manifests.",
		func(opts Options, args []string, help bool) error {
			if help {
				ansi.Fprintf(os.Stdout, "genesis v%s\n", Version)
				ansi.Fprintf(os.Stdout, "USAGE: genesis lookup key env-name default-value\n\n")
				ansi.Fprintf(os.Stdout, "OPTIONS\n")
				return nil
			}

			ansi.Fprintf(os.Stdout, "This is genesis lookup.")
			return nil
		})

	/* genesis manifest */
	c.Dispatch("manifest", "Compile a deployment manifest.",
		func(opts Options, args []string, help bool) error {
			if help {
				ansi.Fprintf(os.Stdout, "genesis v%s\n", Version)
				ansi.Fprintf(os.Stdout, "USAGE: genesis manifest [--no-redact] [--cloud-config path.yml] deployment-env.yml\n\n")
				ansi.Fprintf(os.Stdout, "OPTIONS\n")
				ansi.Fprintf(os.Stdout, "$GLOBAL_USAGE\n\n")
				ansi.Fprintf(os.Stdout, "  -c, --cloud-config PATH    Path to your downloaded BOSH cloud-config\n\n")
				ansi.Fprintf(os.Stdout, "      --no-redact            Do not redact credentials in the manifest.\n")
				ansi.Fprintf(os.Stdout, "                             USE THIS OPTION WITH GREAT CARE AND CAUTION.\n")
				return nil
			}

			ansi.Fprintf(os.Stdout, "This is genesis manifest.")
			return nil
		})

	/* genesis new */
	c.Dispatch("new", "Create a new Genesis deployment environment.",
		func(opts Options, args []string, help bool) error {
			if help {
				ansi.Fprintf(os.Stdout, "genesis v%s\n", Version)
				ansi.Fprintf(os.Stdout, "USAGE: genesis new [--vault target] env-name[.yml]\n\n")
				ansi.Fprintf(os.Stdout, "OPTIONS\n")
				ansi.Fprintf(os.Stdout, "      --vault      The name of a `safe' target (a Vault) to store newly\n")
				ansi.Fprintf(os.Stdout, "                   generated credentials in.\n")
				return nil
			}

			ansi.Fprintf(os.Stdout, "This is genesis new.")
			return nil
		})

	/* genesis ping */
	c.Dispatch("ping", "See if the genesis binary is a real thing.",
		func(opts Options, args []string, help bool) error {
			if help {
				ansi.Fprintf(os.Stdout, "genesis v%s\n", Version)
				ansi.Fprintf(os.Stdout, "USAGE: genesis ping\n\n")
				ansi.Fprintf(os.Stdout, "OPTIONS\n")
				return nil
			}

			ansi.Fprintf(os.Stdout, "This is genesis ping.")
			return nil
		})

	/* genesis repipe */
	c.Dispatch("repipe", "Configure a Concourse pipeline for automating deployments.",
		func(opts Options, args []string, help bool) error {
			if help {
				ansi.Fprintf(os.Stdout, "genesis v%s\n", Version)
				ansi.Fprintf(os.Stdout, "USAGE: genesis repipe [pipeline-layout]\n\n")
				ansi.Fprintf(os.Stdout, "OPTIONS\n")
				ansi.Fprintf(os.Stdout, "  -t, --target     The name of your Concourse target (per `fly targets'),\n")
				ansi.Fprintf(os.Stdout, "                   if it differs from the pipeline layout name.\n\n")
				ansi.Fprintf(os.Stdout, "  -n, --dry-run    Generate the Concourse Pipeline configuration, but\n")
				ansi.Fprintf(os.Stdout, "                   refrain from actually deploying it to Concourse.\n")
				ansi.Fprintf(os.Stdout, "                   Instead, just print the YAML.\n\n")
				ansi.Fprintf(os.Stdout, "  -c, --config     Path to the pipeline configuration file, which specifies\n")
				ansi.Fprintf(os.Stdout, "                   Git parameters, notification settings, pipeline layouts,\n")
				ansi.Fprintf(os.Stdout, "                   etc.  Defaults to 'ci.yml'\n")
				return nil
			}

			ansi.Fprintf(os.Stdout, "This is genesis repipe.")
			return nil
		})

	/* genesis secrets */
	c.Dispatch("secrets", "Re-generate // rotate credentials (passwords, keys, etc.).",
		func(opts Options, args []string, help bool) error {
			if help {
				ansi.Fprintf(os.Stdout, "genesis v%s\n", Version)
				ansi.Fprintf(os.Stdout, "USAGE: genesis secrets [--rotate] [--vault target] deployment-env.yml\n\n")
				ansi.Fprintf(os.Stdout, "OPTIONS\n")
				ansi.Fprintf(os.Stdout, "      --rotate     Rotate credentials.  Any non-fixed credentials defined\n")
				ansi.Fprintf(os.Stdout, "                   by the kit will be regenerated in the Vault.\n")
				ansi.Fprintf(os.Stdout, "      --vault      The name of a `safe' target (a Vault) to store newly\n")
				ansi.Fprintf(os.Stdout, "                   generated credentials in.\n")
				return nil
			}

			ansi.Fprintf(os.Stdout, "This is genesis secrets.")
			return nil
		})

	/* genesis summary */
	c.Dispatch("summary", "Print a summary of defined environments.",
		func(opts Options, args []string, help bool) error {
			if help {
				ansi.Fprintf(os.Stdout, "genesis v$VERSION\n")
				ansi.Fprintf(os.Stdout, "USAGE: genesis summary\n\n")
				ansi.Fprintf(os.Stdout, "OPTIONS\n")
				ansi.Fprintf(os.Stdout, "$GLOBAL_USAGE\n")
				return nil
			}

			ansi.Fprintf(os.Stdout, "This is genesis summary.")
			return nil
		})

	/* genesis version */
	c.Dispatch("version", "Print the current version of Genesis.",
		func (opts Options, args []string, help bool) error {
			if Version == "" {
				ansi.Fprintf(os.Stdout, "genesis @C{(development release)}\n")
			} else {
				ansi.Fprintf(os.Stdout, "genesis v%s\n", Version)
			}
			return nil
		})

	/* genesis yamls */
	c.Dispatch("yamls", "Print a list of the YAML files used for a single environment.",
		func(opts Options, args []string, help bool) error {
			if help {
				ansi.Fprintf(os.Stdout, "genesis v%s\n", Version)
				ansi.Fprintf(os.Stdout, "USAGE: genesis yamls deployment-env.yml\n\n")
				ansi.Fprintf(os.Stdout, "OPTIONS\n")
				return nil
			}

			ansi.Fprintf(os.Stdout, "This is genesis yamls.")
			return nil
		})

	err := c.Execute(os.Args[1:]...)
	if err != nil {
		ansi.Fprintf(os.Stderr, "@R{!!! %s}\n", err)
		os.Exit(1)
	}
}
