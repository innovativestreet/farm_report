package main

import (
	"absolutetech/farm_report/global-lib/envutils"
	"flag"
	"os"
)

type flags struct {
	environment      envutils.Env
	structureLogDir  string
	structureLogFile string
	debug            string
}

func getFlags() *flags {
	flg := &flags{
		environment: envutils.Testing,
	}
	fs := flag.NewFlagSet("", flag.ExitOnError)
	envutils.SetFlag(fs, &flg.environment)
	fs.StringVar(&flg.structureLogDir, "structure-dir", flg.structureLogDir, "Structure Log Directory")
	fs.StringVar(&flg.structureLogFile, "structure-file", flg.structureLogFile, "Structure Log File")
	fs.StringVar(&flg.debug, "debug", flg.debug, "Debug address")
	_ = fs.Parse(os.Args[1:]) // Ignore error, because it exits on error
	return flg
}
