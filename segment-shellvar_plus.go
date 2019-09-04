package main

import (
	"os"
	"strconv"
)

func segmentShellVarPlus(p *powerline) {
	shellVarName := *p.args.ShellVarPlus

	shellVarNameDefault := "POWERLINE_SEGMENT_SHELLVAR"
	shellVarFromArg := true

	if shellVarName == "" {
		shellVarName = shellVarNameDefault
		shellVarFromArg = false
	}

	varQuiet, _ := os.LookupEnv("POWERLINE_SEGMENT_SHELLVAR_QUIET")

	varContent, varExists := os.LookupEnv(shellVarName)
	if varExists {
		fg := p.theme.ShellVarFg
		bg := p.theme.ShellVarFg
		var result uint64
		var err error

		// Try to load colors from the environment, if not,
		// fallback on theme colors for the builtin shell-var segment
		varFG, _ := os.LookupEnv("POWERLINE_SEGMENT_SHELLVAR_FG")
		result, err = strconv.ParseUint(varFG, 10, 8)
		if err == nil {
			fg = uint8(result)
		} else {
			fg = p.theme.ShellVarFg
		}
		varBG, _ := os.LookupEnv("POWERLINE_SEGMENT_SHELLVAR_BG")
		result, err = strconv.ParseUint(varBG, 10, 8)
		if err == nil {
			bg = uint8(result)
		} else {
			bg = p.theme.ShellVarBg
		}

		// Handle adding the segment, displaying a warning if
		// quiet mode is not enabled
		if varContent != "" {
			p.appendSegment("shell-var-plus", segment{
				content:    varContent,
				foreground: fg,
				background: bg,
			})
		} else {
			if shellVarFromArg && varQuiet != "1" {
				warn("Shell variable " + shellVarName + " does not exist.\n")
			} else {
				p.appendSegment("shell-var-plus", segment{
					content:    shellVarName,
					foreground: fg,
					background: bg,
				})
			}
		}
	} else {
		if shellVarFromArg && varQuiet != "1" {
			// If the shell variable is not set as a command line argument,
			// and doesn't exist in the environment, we are quiet to avoid
			// false-positive warnings.
			// Else if the user only configures the default environment
			// variable, we are also quiet by default.
			// When a user directly uses the command line argument,
			// we warn by default (to mimic -shell-var) unless explicitly
			// silenced by environment variable
			warn("Shell variable " + shellVarName + " does not exist.\n")
		} 
	}
}
