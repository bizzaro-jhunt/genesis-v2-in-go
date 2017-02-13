package main

import (
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	fmt "github.com/starkandwayne/goutils/ansi"
)

var (
	SpruceMinimumVersion = "1.8.9"
	SafeMinimumVersion   = "0.0.29"
	GitMinimumVersion    = "1.8.0"
)

func uinteger(s string) uint {
	u, _ := strconv.ParseUint(s, 10, 0)
	return uint(u)
}

func semver(s string) ([]uint, error) {
	var (
		vXYZR = regexp.MustCompile(`^(\d+)\.(\d+)\.(\d+)\.(\d+)$`)
		vXYZ  = regexp.MustCompile(`^(\d+)\.(\d+)\.(\d+)$`)
		vXY   = regexp.MustCompile(`^(\d+)\.(\d+)$`)
		vX    = regexp.MustCompile(`^(\d+)$`)
	)

	if m := vXYZR.FindStringSubmatch(s); len(m) > 0 {
		return []uint{
			uinteger(m[1]),
			uinteger(m[2]),
			uinteger(m[3]),
			uinteger(m[4]),
		}, nil
	}
	if m := vXYZ.FindStringSubmatch(s); len(m) > 0 {
		return []uint{
			uinteger(m[1]),
			uinteger(m[2]),
			uinteger(m[3]),
			0,
		}, nil
	}
	if m := vXY.FindStringSubmatch(s); len(m) > 0 {
		return []uint{
			uinteger(m[1]),
			uinteger(m[2]),
			0,
			0,
		}, nil
	}
	if m := vX.FindStringSubmatch(s); len(m) > 0 {
		return []uint{
			uinteger(m[1]),
			0,
			0,
			0,
		}, nil
	}

	return nil, fmt.Errorf("'%s' does not look like a valid version string", s)
}

func newEnough(a, b string) (bool, error) {
	vA, err := semver(a)
	if err != nil {
		return false, err
	}

	vB, err := semver(b)
	if err != nil {
		return false, err
	}

	for i := range vA {
		if vA[i] > vB[i] {
			return true, nil
		}
		if vA[i] < vB[i] {
			return false, nil
		}
	}

	/* exact same version!  what're the odds? */
	return true, nil
}

func checkPrerequisites() {
	var (
		s      string
		b      []byte
		err    error
		failed = false
	)

	// check for a new enough Spruce
	b, err = exec.Command("/bin/sh", "-c", "spruce -v 2>/dev/null").Output()
	if err != nil {
		failed = true
		fmt.Fprintf(os.Stderr, "@R{!!! Missing `spruce' - install Spruce from https://github.com/geofffranks/spruce/releases}\n")
		fmt.Fprintf(os.Stderr, "    (`spruce -v` said: %s)\n", err)
	} else {
		s = strings.TrimSuffix(string(b), "\n")
		m := regexp.MustCompile(`(?i)version\s+(\S+)`).FindStringSubmatch(s)
		if len(m) != 2 {
			failed = true
			fmt.Fprintf(os.Stderr, "@R{!!! Your `spruce' binary seems to be corrupt; running `spruce -v' resulted in}\n")
			fmt.Fprintf(os.Stderr, "    '%s'\n", s)
			fmt.Fprintf(os.Stderr, "    Please re-install Spruce from https://github.com/geofffranks/spruce/releases\n")
		} else {
			ok, err := newEnough(m[1], SpruceMinimumVersion)
			if err != nil {
				failed = true
				fmt.Fprintf(os.Stderr, "@R{!!! Your `spruce' binary seems to be corrupt; running `spruce -v' resulted in}\n")
				fmt.Fprintf(os.Stderr, "    '%s' (%s)\n", s, err)
				fmt.Fprintf(os.Stderr, "    Please re-install Spruce from https://github.com/geofffranks/spruce/releases\n")
			} else if !ok {
				failed = true
				fmt.Fprintf(os.Stderr, "@R{!!! Spruce v%s is installed, but Genesis requires at least v%s}\n", m[1], SpruceMinimumVersion)
				fmt.Fprintf(os.Stderr, "    Please upgrade your Spruce, via https://github.com/geofffranks/spruce/releases\n")
			}
		}
	}

	// check for a new enough Safe
	b, err = exec.Command("/bin/sh", "-c", "safe -v 2>&1").Output()
	if err != nil {
		failed = true
		fmt.Fprintf(os.Stderr, "@R{!!! Missing `safe' - install Spruce from https://github.com/starkandwayne/safe/releases}\n")
		fmt.Fprintf(os.Stderr, "    (`safe -v` said: %s)\n", err)
	} else {
		s = strings.TrimSuffix(string(b), "\n")
		m := regexp.MustCompile(`(?i)^safe\s+v(\S+)`).FindStringSubmatch(s)
		if len(m) != 2 {
			failed = true
			fmt.Fprintf(os.Stderr, "@R{!!! Your `safe' binary seems to be corrupt; running `safe -v' resulted in}\n")
			fmt.Fprintf(os.Stderr, "    '%s'\n", s)
			fmt.Fprintf(os.Stderr, "    Please re-install Safe from https://github.com/starkandwayne/safe/releases\n")
		} else {
			ok, err := newEnough(m[1], SafeMinimumVersion)
			if err != nil {
				failed = true
				fmt.Fprintf(os.Stderr, "@R{!!! Your `safe' binary seems to be corrupt; running `safe -v' resulted in}\n")
				fmt.Fprintf(os.Stderr, "    '%s' (%s)\n", s, err)
				fmt.Fprintf(os.Stderr, "    Please re-install Safe from https://github.com/starkandwayne/safe/releases\n")
			} else if !ok {
				failed = true
				fmt.Fprintf(os.Stderr, "@R{!!! Safe v%s is installed, but Genesis requires at least v%s}\n", m[1], SafeMinimumVersion)
				fmt.Fprintf(os.Stderr, "    Please upgrade your Safe, via https://github.com/starkandwayne/safe/releases\n")
			}
		}
	}

	// check for a new enough Git
	b, err = exec.Command("/bin/sh", "-c", "git --version 2>/dev/null").Output()
	if err != nil {
		failed = true
		fmt.Fprintf(os.Stderr, "@R{!!! Missing `git' - install git via your platform package manager}\n")
		fmt.Fprintf(os.Stderr, "    (`git --version` said: %s)\n", err)
	} else {
		s = strings.TrimSuffix(string(b), "\n")
		m := regexp.MustCompile(`(?i)version\s+(\S+)`).FindStringSubmatch(s)
		if len(m) != 2 {
			failed = true
			fmt.Fprintf(os.Stderr, "@R{!!! Your `git' binary seems to be corrupt; running `git -v' resulted in}\n")
			fmt.Fprintf(os.Stderr, "    '%s'\n", s)
			fmt.Fprintf(os.Stderr, "    Please re-install git (via your platform package manager)\n")
		} else {
			ok, err := newEnough(m[1], GitMinimumVersion)
			if err != nil {
				failed = true
				fmt.Fprintf(os.Stderr, "@R{!!! Your `git' binary seems to be corrupt; running `git --version' resulted in}\n")
				fmt.Fprintf(os.Stderr, "    '%s' (%s)\n", s, err)
				fmt.Fprintf(os.Stderr, "    Please re-install git (via your platform package manager)\n")
			} else if !ok {
				failed = true
				fmt.Fprintf(os.Stderr, "@R{!!! Git v%s is installed, but Genesis requires at least v%s}\n", m[1], GitMinimumVersion)
				fmt.Fprintf(os.Stderr, "    Please upgrade git (via your platform package manager)\n")
			}
		}
	}

	// FIXME: check for a new enough BOSH v2 CLI

	if failed {
		fmt.Fprintf(os.Stderr, "\n@R{GENESIS PREREQ CHECKS FAILED!!}\n")
		fmt.Fprintf(os.Stderr, "@R{Your system does not look like it is ready for Genesis.}\n")
		os.Exit(2)
	}
}
