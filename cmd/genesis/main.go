package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"regexp"
	"sort"
	"strings"
	"time"

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
		Compilekit:   getopt.StringLong("compile-kit", 0, "Create a distributable kit archive from dev."),
		Decompilekit: getopt.StringLong("decompile-kit", 0, "Unpack a kit to dev."),
		Describe:     getopt.StringLong("describe", 0, "", "Describe a Concourse pipeline, in words."),
		Download:     getopt.StringLong("download", 0, "", "Download a Genesis Kit from the Internet."),
		Graph:        getopt.StringLong("graph", 0, "", "Draw a Concourse pipeline."),
		Init:         getopt.StringLong("init", 0, "", "Initialize a new Genesis deployment."),
		Lookup:       getopt.StringLong("lookup", 0, "", "Find a key set in environment manifests."),
		Manifest:     getopt.StringLong("manifest", 0, "", "Generate a redacted BOSH deployment manifest for an environment."),
		New:          getopt.StringLong("new", 0, "", "Create a new Genesis deployment environment."),
		Ping:         getopt.StringLong("ping", 0, "", "See if the genesis binary is a real thing."),
		Repipe:       getopt.StringLong("repipe", 0, "", "Configure a Concourse pipeline for automating deployments."),
		Secrets:      getopt.StringLong("secrets", 0, "", "Re-generate // rotate credentials (passwords, keys, etc.)."),
		Summary:      getopt.BoolLong("summary", 0, "Print a summary of defined environments."),
		Version:      getopt.BoolLong("version", 0, "Print the version of Genesis."),
		Yamls:        getopt.StringLong("yamls", 0, "Print a list of the YAML files used for a single environment."),
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

	if *options.Version {
		if Version == "" {
			fmt.Printf("genesis (development)%s\n", Version)
		} else {
			fmt.Printf("genesis v%s\n", Version)
		}
		os.Exit(0)
	}

	c := NewCommand().With(options)

	c.HelpGroup("INFO:")
	c.Dispatch("help", "Get detailed help with a specific command",
		func(opts Options, args []string, help bool) error {
			if len(args) == 0 {
				buf := bytes.Buffer{}
				getopt.PrintUsage(&buf)
				ansi.Fprintf(os.Stderr, strings.Split(buf.String(), "\n")[0]+"\n")
				ansi.Fprintf(os, Stderr, "genesis v%s\n", Version)
				ansi.Fprintf(os.Stderr, "USAGE: genesis [OPTIONS] COMMAND [MORE OPTIONS]\n")
				ansi.Fprintf(os.Stderr, "\n  OPTIONS\n")
				ansi.Fprintf(os.Stderr, "    -h, --help       Show this help screen.\n")
				ansi.Fprintf(os.Stderr, "    -D, --debug      Enable debugging, printing helpful message about what\n")
				ansi.Fprintf(os.Stderr, "                     Genesis is doing, to standard error.\n")
				ansi.Fprintf(os.Stderr, "    -T, --trace      Even more debugging, including debugging inside called\n")
				ansi.Fprintf(os.Stderr, "                     tools (like spruce and bosh).\n")
				ansi.Fprintf(os.Stderr, "    -C, --cwd        Effective working directory.  Defaults to '.'\n")
				ansi.Fprintf(os.Stderr, "    -y, --yes        Answer 'yes' to all question, automatically.\n")
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
				ansi.Fprintf(os.Stderr, "\n  See `genesis COMMAND -h' for more specific, per-command usage information.\n")
				return nil
			} else if args[0] == "help" {
				ansi.Fprintf(os.Stderr, "@R{This is getting a bit too meta, don't you think?}\n")
				return nil
			}
			return c.Help(args...)
		})

	c.Alias("help", "usage")

	/* genesis compile-kit */
	c.Dispatch("compile-kit", "Create a distributable kit archive from dev.",
		func(opts Options, args []string, help bool) error {
			if help {
				ansi.Fprintf(os.Stdout, "genesis v%s\n", Version)
				ansi.Fprintf(os.Stdout, "USAGE: genesis compile-kit -n NAME -v VERSION\n")
				ansi.Fprintf(os.Stdout, "\n")
				ansi.Fprintf(os.Stdout, "OPTIONS\n")
				ansi.Fprintf(os.Stdout, "  -n, --name      Name of the kit archive.\n")
				ansi.Fprintf(os.Stdout, "  -v, --version   Version to package.\n")
				return nil
			}

			if *opts.Compilekit {
				ansi.Fprintf(os.Stdout, "Really just needed to use opts.")
			}

			ansi.Fprintf(os.Stdout, "This is genesis compile-kit.")
			return nil
		})

	/* genesis decompile-kit */
	c.Dispatch("decompile-kit", "Unpack a kit archive to dev.",
		func(opts Options, args []string, help bool) error {
			if help {
				ansi.Fprintf(os.Stdout, "genesis v%s\n", Version)
				ansi.Fprintf(os.Stdout, "USAGE: genesis decompile-kit [NAME/VERSION | path/to/kit.tar.gz]\n")
				ansi.Fprintf(os.Stdout, "\n")
				ansi.Fprintf(os.Stdout, "OPTIONS\n")
				ansi.Fprintf(os.Stdout, "  -f, --force  Overwrite dev/, if it exists.\n")
				return nil
			}

			if *opts.Decompilekit {
				ansi.Fprintf(os.Stdout, "Really just needed to use opts.")
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

			if *opts.Describe {
				ansi.Fprintf(os.Stdout, "Really just needed to use opts.")
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

			if *opts.Download {
				ansi.Fprintf(os.Stdout, "Really just needed to use opts.")
			}

			ansi.Fprintf(os.Stdout, "This is genesis download.")
			return nil
		})

	/* genesis graph */
	c.Dispatch("graph", "Draw a Concourse pipeline.",
		func(opts Options, args []string, help bool) error {
			if help {
				ansi.Fprintf(os.Stdout, "genesis v%s\n", Version)
				ansi.Fprintf(os.Stdout, "USAGE: genesis graph [pipeline-layout]\n")
				ansi.Fprintf(os.Stdout, "\n")
				ansi.Fprintf(os.Stdout, "OPTIONS\n")
				ansi.Fprintf(os.Stdout, "  -c, --config     Path to the pipeline configuration file, which specifies\n")
				ansi.Fprintf(os.Stdout, "                   Git parameters, notification settings, pipeline layouts,\n")
				ansi.Fprintf(os.Stdout, "                   etc.  Defaults to 'ci.yml'\n")
				return nil
			}

			if *opts.Graph {
				ansi.Fprintf(os.Stdout, "Really just needed to use opts.")
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

			if *opts.Init {
				ansi.Fprintf(os.Stdout, "Really just needed to use opts.")
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

			if *opts.Lookup {
				ansi.Fprintf(os.Stdout, "Really just needed to use opts.")
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

			if *opts.Manifest {
				ansi.Fprintf(os.Stdout, "Really just needed to use opts.")
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

			if *opts.New {
				ansi.Fprintf(os.Stdout, "Really just needed to use opts.")
			}

			ansi.Fprintf(os.Stdout, "This is genesis new.")
			return nil
		})

	/* genesis ping */
	c.Dispatch("ping", "See if the genesis binary is a real thing.",
		func(opts Options, args []string, help bool) error {
			if help {
				ansi.Fprintf(os.Stdout, "genesis v%s\n", Version)
				ansi.Fprintf(os.Stdout, "USAGE: genesis ping\n")
				ansi.Fprintf(os.Stdout, "\n")
				ansi.Fprintf(os.Stdout, "OPTIONS\n")
				return nil
			}

			if *opts.Ping {
				ansi.Fprintf(os.Stdout, "Really just needed to use opts.")
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

			if *opts.Repipe {
				ansi.Fprintf(os.Stdout, "Really just needed to use opts.")
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

			if *opts.Secrets {
				ansi.Fprintf(os.Stdout, "Really just needed to use opts.")
			}

			ansi.Fprintf(os.Stdout, "This is genesis secrets.")
			return nil
		})

	/* genesis summary */
	c.Dispatch("summary", "Print a summary of defined environments.",
		func(opts Options, args []string, help bool) error {
			if help {
				ansi.Frpintf(os.Stdout, "genesis v$VERSION\n")
				ansi.Frpintf(os.Stdout, "USAGE: genesis summary\n\n")
				ansi.Frpintf(os.Stdout, "OPTIONS\n")
				ansi.Frpintf(os.Stdout, "$GLOBAL_USAGE\n")
				return nil
			}

			if *opts.Summary {
				ansi.Fprintf(os.Stdout, "Really just needed to use opts.")
			}

			ansi.Fprintf(os.Stdout, "This is genesis summary.")
			return nil
		})

	/* genesis version */
	c.Dispatch("version", "Print the version of Genesis.",
		func(opts Options, args []string, help bool) error {
			if help {
				ansi.Fprintf(os.Stdout, "genesis v%s\n", Version)
				ansi.Fprintf(os.Stdout, "USAGE: genesis version\n\n")
				ansi.Fprintf(os.Stdout, "OPTIONS\n")
				return nil
			}

			if *opts.Version {
				ansi.Fprintf(os.Stdout, "Really just needed to use opts.")
			}

			ansi.Fprintf(os.Stdout, "This is genesis version.")
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

			if *opts.Yamls {
				ansi.Fprintf(os.Stdout, "Really just needed to use opts.")
			}

			ansi.Fprintf(os.Stdout, "This is genesis yamls.")
			return nil
		})
}
