package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"strings"

	. "github.com/jhunt/genesis/command"
	"github.com/pborman/getopt"
	fmt "github.com/starkandwayne/goutils/ansi"
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
		Verbose: getopt.BoolLong("verbose", 'v', "Enable debugging to print helpful messages about what Genesis is doing (for operators)."),
		Yes:     getopt.BoolLong("yes", 'y', "Answer 'yes' to all questions automatically."),
	}

	opts := getopt.CommandLine
	opts.Parse(os.Args)

	args := opts.Args()
	debug = *options.Debug

	c := NewCommand().With(options)

	c.HelpGroup("INFO:")
	c.Dispatch("help", "Get detailed help with a specific command",
		func(opts Options, args []string, help bool) error {
			if len(args) == 0 {
				buf := bytes.Buffer{}
				getopt.PrintUsage(&buf)
				fmt.Fprintf(os.Stderr, strings.Split(buf.String(), "\n")[0]+"\n")
				if Version == "" {
					fmt.Fprintf(os.Stderr, "genesis @*{development version 🦄}\n")
				} else {
					fmt.Fprintf(os.Stderr, "genesis v%s\n", Version)
				}
				fmt.Fprintf(os.Stderr, "USAGE: genesis [OPTIONS] COMMAND [MORE OPTIONS]\n")
				fmt.Fprintf(os.Stderr, "\n  OPTIONS\n")
				fmt.Fprintf(os.Stderr, "    -h, --help       Show this help screen.\n")
				fmt.Fprintf(os.Stderr, "    -D, --debug      Enable debugging, printing helpful message about what\n")
				fmt.Fprintf(os.Stderr, "                     Genesis is doing, to standard error.\n")
				fmt.Fprintf(os.Stderr, "    -T, --trace      Even more debugging, including debugging inside called\n")
				fmt.Fprintf(os.Stderr, "                     tools (like spruce and bosh).\n")
				fmt.Fprintf(os.Stderr, "    -C, --cwd        Effective working directory.  Defaults to '.'\n")
				fmt.Fprintf(os.Stderr, "    -y, --yes        Answer 'yes' to all questions, automatically.\n")
				fmt.Fprintf(os.Stderr, "\n\n  COMMANDS\n")
				fmt.Fprintf(os.Stderr, "    compile-kit      Create a distributable kit archive from dev.\n")
				fmt.Fprintf(os.Stderr, "    decompile-kit    Unpack a kit archive to dev.\n")
				fmt.Fprintf(os.Stderr, "    describe         Describe a Concourse pipeline, in words.\n")
				fmt.Fprintf(os.Stderr, "    download         Download a Genesis Kit from the Internet.\n")
				fmt.Fprintf(os.Stderr, "    graph            Draw a Concourse pipeline.\n")
				fmt.Fprintf(os.Stderr, "    init             Initialize a new Genesis deployment.\n")
				fmt.Fprintf(os.Stderr, "    lookup           Find a key set in environment manifests.\n")
				fmt.Fprintf(os.Stderr, "    manifest         Generate a redacted BOSH deployment manifest for an environment.\n")
				fmt.Fprintf(os.Stderr, "    new              Create a new Genesis deployment environment.\n")
				fmt.Fprintf(os.Stderr, "    ping             See if the genesis binary is a real thing.\n")
				fmt.Fprintf(os.Stderr, "    repipe           Configure a Concourse pipeline for automating deployments.\n")
				fmt.Fprintf(os.Stderr, "    secrets          Re-generate // rotate credentials (passwords, keys, etc.).\n")
				fmt.Fprintf(os.Stderr, "    summary          Print a summary of defined environments.\n")
				fmt.Fprintf(os.Stderr, "    version          Print the version of genesis\n")
				fmt.Fprintf(os.Stderr, "    yamls            Print a list of the YAML files used for a single environment.\n")
				fmt.Fprintf(os.Stderr, "\n  See 'genesis COMMAND -h' for more specific, per-command usage information.\n")
				return nil
			} else if args[0] == "help" {
				fmt.Fprintf(os.Stderr, "@R{This is getting a bit too meta, don't you think?}\n")
				return nil
			}
			return c.Help(args...)
		})
	c.Alias("usage", "help")

	/* genesis compile-kit */
	// FIXME: implement
	c.Dispatch("compile-kit", "Create a distributable kit archive from dev.",
		func(opts Options, args []string, help bool) error {
			if help {
				fmt.Printf("genesis v%s\n", Version)
				fmt.Printf("USAGE: genesis compile-kit -n NAME -v VERSION\n\n")
				fmt.Printf("OPTIONS\n")
				fmt.Printf("  -n, --name      Name of the kit archive.\n")
				fmt.Printf("  -v, --version   Version to package.\n")
				return nil
			}

			fmt.Printf("This is genesis compile-kit.")
			return nil
		})

	/* genesis decompile-kit */
	// FIXME: implement
	c.Dispatch("decompile-kit", "Unpack a kit archive to dev.",
		func(opts Options, args []string, help bool) error {
			if help {
				fmt.Printf("genesis v%s\n", Version)
				fmt.Printf("USAGE: genesis decompile-kit [NAME/VERSION | path/to/kit.tar.gz]\n\n")
				fmt.Printf("OPTIONS\n")
				fmt.Printf("  -f, --force  Overwrite dev/, if it exists.\n")
				return nil
			}

			fmt.Printf("This is genesis decompile-kit.")
			return nil
		})

	/* genesis describe */
	// FIXME: implement
	c.Dispatch("describe", "Describe a Concourse pipeline with words.",
		func(opts Options, args []string, help bool) error {
			if help {
				fmt.Printf("genesis v%s\n", Version)
				fmt.Printf("USAGE: genesis describe [pipeline-layout]\n\n")
				fmt.Printf("OPTIONS\n")
				fmt.Printf("  -c, --config     Path to the pipeline configuration file, which specifies\n")
				fmt.Printf("                   Git parameters, notification settings, pipeline layouts,\n")
				fmt.Printf("                   etc.  Defaults to 'ci.yml'\n")
				return nil
			}

			fmt.Printf("This is genesis describe.")
			return nil
		})

	/* genesis download */
	// FIXME: implement
	c.Dispatch("download", "Download a Genesis Kit from the Internet.",
		func(opts Options, args []string, help bool) error {
			if help {
				fmt.Printf("genesis v%s\n", Version)
				fmt.Printf("USAGE: genesis download NAME[/VERSION] [...]\n\n")
				fmt.Printf("OPTIONS\n")
				return nil
			}

			fmt.Printf("This is genesis download.")
			return nil
		})

	/* genesis graph */
	// FIXME: implement
	c.Dispatch("graph", "Draw a Concourse pipeline.",
		func(opts Options, args []string, help bool) error {
			if help {
				fmt.Printf("genesis v%s\n", Version)
				fmt.Printf("USAGE: genesis graph [pipeline-layout]\n\n")
				fmt.Printf("OPTIONS\n")
				fmt.Printf("  -c, --config     Path to the pipeline configuration file, which specifies\n")
				fmt.Printf("                   Git parameters, notification settings, pipeline layouts,\n")
				fmt.Printf("                   etc.  Defaults to 'ci.yml'\n")
				return nil
			}

			fmt.Printf("This is genesis graph.")
			return nil
		})

	/* genesis init */
	// FIXME: implement
	c.Dispatch("init", "Initialize a new Genesis deployment.",
		func(global Options, args []string, help bool) error {
			if help {
				fmt.Printf("genesis v%s\n", Version)
				fmt.Printf("USAGE: genesis init [-k KIT/VERSION] name\n\n")
				fmt.Printf("OPTIONS\n")
				fmt.Printf("  -k, --kit        Name (and optionally, version) of the Genesis Kit to\n")
				fmt.Printf("                   base these deployments on.  i.e.: shield/6.3.0\n")
				return nil
			}

			getopt.Reset()
			kitver := getopt.StringLong("kit", 'k', "", "Name (and optionally, version) of the Genesis Kit to base these deployments on")
			isdev := getopt.BoolLong("dev", 'd', "Whether or not to create a dev/ kit")

			opts := getopt.CommandLine
			args = append([]string{"init"}, args...)
			opts.Parse(args)
			args = opts.Args()

			if len(args) != 1 {
				fmt.Fprintf(os.Stderr, "@R{USAGE...}\n")
				os.Exit(3)
			}

			checkPrerequisites()

			name := strings.TrimSuffix(args[0], "-deployments")

			if isdev != nil && *isdev && kitver != nil {
				fmt.Fprintf(os.Stderr, "@R{You cannot specify both --kit and --dev, together.}\n")
				os.Exit(3)
			}
			if !validRepoName(name) {
				fmt.Fprintf(os.Stderr, "@R{Invalid Genesis repo name '%s'}\n", name)
				os.Exit(3)
			}

			if kitver != nil {
				fmt.Printf("using kit %s\n", *kitver)
			}
			if isdev != nil && *isdev {
				fmt.Printf("creating a deelopment deployment\n")
			}

			mkdirs(root, ".genesis", "bin")
			f, err := os.OpenFile(root + "/.genesis/config", 0, 0666)
			if err != nil {
				fmt.Fprintf(os.Stderr, "@R{!!! %s}\n", err)
				os.Exit(3)
			}
			fmt.Fprintf(f, "---\n")
			fmt.Fprintf(f, "genesis: %s\n", Version)
			f.Close()

			copySelf(root + "/.genesis/bin/genesis")
			if isdev && *isdev {
				mkdirs(root, "dev")
			}

			for _, option := range args {
				fmt.Printf(" - %s\n", option)
			}

			fmt.Printf("This is genesis init.\n")
			return nil
		})

	/* genesis lookup */
	// FIXME: implement
	c.Dispatch("lookup", "Find a key set in environment manifests.",
		func(opts Options, args []string, help bool) error {
			if help {
				fmt.Printf("genesis v%s\n", Version)
				fmt.Printf("USAGE: genesis lookup key env-name default-value\n\n")
				fmt.Printf("OPTIONS\n")
				return nil
			}

			fmt.Printf("This is genesis lookup.")
			return nil
		})

	/* genesis manifest */
	// FIXME: implement
	c.Dispatch("manifest", "Compile a deployment manifest.",
		func(opts Options, args []string, help bool) error {
			if help {
				fmt.Printf("genesis v%s\n", Version)
				fmt.Printf("USAGE: genesis manifest [--no-redact] [--cloud-config path.yml] deployment-env.yml\n\n")
				fmt.Printf("OPTIONS\n")
				fmt.Printf("$GLOBAL_USAGE\n\n")
				fmt.Printf("  -c, --cloud-config PATH    Path to your downloaded BOSH cloud-config\n\n")
				fmt.Printf("      --no-redact            Do not redact credentials in the manifest.\n")
				fmt.Printf("                             USE THIS OPTION WITH GREAT CARE AND CAUTION.\n")
				return nil
			}

			fmt.Printf("This is genesis manifest.")
			return nil
		})

	/* genesis new */
	// FIXME: implement
	c.Dispatch("new", "Create a new Genesis deployment environment.",
		func(opts Options, args []string, help bool) error {
			if help {
				fmt.Printf("genesis v%s\n", Version)
				fmt.Printf("USAGE: genesis new [--vault target] env-name[.yml]\n\n")
				fmt.Printf("OPTIONS\n")
				fmt.Printf("      --vault      The name of a `safe' target (a Vault) to store newly\n")
				fmt.Printf("                   generated credentials in.\n")
				return nil
			}

			fmt.Printf("This is genesis new.")
			return nil
		})

	/* genesis ping */
	c.Dispatch("ping", "See if the genesis binary is a real thing.",
		func(opts Options, args []string, help bool) error {
			if help {
				fmt.Printf("genesis v%s\n", Version)
				fmt.Printf("USAGE: genesis ping\n\n")
				fmt.Printf("OPTIONS\n")
				return nil
			}

			checkPrerequisites()
			fmt.Printf("PING\n")
			return nil
		})

	/* genesis repipe */
	// FIXME: implement
	c.Dispatch("repipe", "Configure a Concourse pipeline for automating deployments.",
		func(opts Options, args []string, help bool) error {
			if help {
				fmt.Printf("genesis v%s\n", Version)
				fmt.Printf("USAGE: genesis repipe [pipeline-layout]\n\n")
				fmt.Printf("OPTIONS\n")
				fmt.Printf("  -t, --target     The name of your Concourse target (per `fly targets'),\n")
				fmt.Printf("                   if it differs from the pipeline layout name.\n\n")
				fmt.Printf("  -n, --dry-run    Generate the Concourse Pipeline configuration, but\n")
				fmt.Printf("                   refrain from actually deploying it to Concourse.\n")
				fmt.Printf("                   Instead, just print the YAML.\n\n")
				fmt.Printf("  -c, --config     Path to the pipeline configuration file, which specifies\n")
				fmt.Printf("                   Git parameters, notification settings, pipeline layouts,\n")
				fmt.Printf("                   etc.  Defaults to 'ci.yml'\n")
				return nil
			}

			fmt.Printf("This is genesis repipe.")
			return nil
		})

	/* genesis secrets */
	// FIXME: implement
	c.Dispatch("secrets", "Re-generate // rotate credentials (passwords, keys, etc.).",
		func(opts Options, args []string, help bool) error {
			if help {
				fmt.Printf("genesis v%s\n", Version)
				fmt.Printf("USAGE: genesis secrets [--rotate] [--vault target] deployment-env.yml\n\n")
				fmt.Printf("OPTIONS\n")
				fmt.Printf("      --rotate     Rotate credentials.  Any non-fixed credentials defined\n")
				fmt.Printf("                   by the kit will be regenerated in the Vault.\n")
				fmt.Printf("      --vault      The name of a `safe' target (a Vault) to store newly\n")
				fmt.Printf("                   generated credentials in.\n")
				return nil
			}

			fmt.Printf("This is genesis secrets.")
			return nil
		})

	/* genesis summary */
	// FIXME: implement
	c.Dispatch("summary", "Print a summary of defined environments.",
		func(opts Options, args []string, help bool) error {
			if help {
				fmt.Printf("genesis v$VERSION\n")
				fmt.Printf("USAGE: genesis summary\n\n")
				fmt.Printf("OPTIONS\n")
				fmt.Printf("$GLOBAL_USAGE\n")
				return nil
			}

			fmt.Printf("This is genesis summary.")
			return nil
		})

	/* genesis version */
	// FIXME: implement
	c.Dispatch("version", "Print the current version of Genesis.",
		func(opts Options, args []string, help bool) error {
			if Version == "" {
				fmt.Printf("genesis @C{(development release)}\n")
			} else {
				fmt.Printf("genesis v%s\n", Version)
			}
			return nil
		})

	/* genesis yamls */
	// FIXME: implement
	c.Dispatch("yamls", "Print a list of the YAML files used for a single environment.",
		func(opts Options, args []string, help bool) error {
			if help {
				fmt.Printf("genesis v%s\n", Version)
				fmt.Printf("USAGE: genesis yamls deployment-env.yml\n\n")
				fmt.Printf("OPTIONS\n")
				return nil
			}

			fmt.Printf("This is genesis yamls.")
			return nil
		})

	err := c.Execute(args...)
	if err != nil {
		fmt.Fprintf(os.Stderr, "@R{!!! %s}\n", err)
		os.Exit(1)
	}
}
