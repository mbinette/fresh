package runner

import (
	"flag"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
)

var doGet bool

func init() {
	flag.BoolVar(&doGet, "get", false, "Execute go get before build")
}

func get() (string, bool) {

	if doGet {
		buildLog("Get dependencies...")

		cmd := exec.Command("go", "get")

		stderr, err := cmd.StderrPipe()
		if err != nil {
			fatal(err)
		}

		stdout, err := cmd.StdoutPipe()
		if err != nil {
			fatal(err)
		}

		err = cmd.Start()
		if err != nil {
			fatal(err)
		}

		io.Copy(os.Stdout, stdout)
		errBuf, _ := ioutil.ReadAll(stderr)

		err = cmd.Wait()
		if err != nil {
			return string(errBuf), false
		}
	}
	return "", true
}
