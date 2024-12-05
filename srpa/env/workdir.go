package env

import (
	"flag"
	"fmt"
	"os"
)

func SetWorkDir() {
	env := os.Getenv("SRPA_HOME")

	if env != "" {
		os.Chdir(env)
		return
	}

	workdir := flag.String("workdir", "", "spacify work directory")

	flag.Parse()

	if *workdir != "" {
		os.Chdir(*workdir)
	}

	wd, err := os.Getwd()
	if err == nil {
		fmt.Println("Current workdir is " + wd)
	}
}
