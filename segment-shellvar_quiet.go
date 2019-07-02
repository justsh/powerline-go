package main

import (
	"os"
)

func segmentShellVarQuiet(p *powerline) {
	shellVarName := *p.args.ShellVarQuiet
	varContent, varExists := os.LookupEnv(shellVarName)

	if varExists {
		if varContent != "" {
			p.appendSegment("shell-var-quiet", segment{
				content:    varContent,
				foreground: p.theme.ShellVarFg,
				background: p.theme.ShellVarBg,
			})
		} else {
			p.appendSegment("shell-var-quiet", segment{
				content:    shellVarName,
				foreground: p.theme.ShellVarFg,
				background: p.theme.ShellVarBg,
			})
		}
	}
}
