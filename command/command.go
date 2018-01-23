package command

import (
	"io/ioutil"
	"os/exec"

	log "qiniupkg.com/x/log.v7"
)

func Run(cmd *exec.Cmd) (result string, err error) {
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Errorf("cmd.StdoutPipe(): %v ", err)
		return
	}
	defer stdout.Close()

	if err = cmd.Start(); err != nil {
		log.Errorf("cmd.Start(): %v ", err)
		return
	}

	defer func() {
		if err = cmd.Wait(); err != nil {
			log.Errorf("cmd.Wait(%s): %v ", cmd.Args, err)
		}
	}()

	opBytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		log.Errorf("ioutil.ReadAll(): %v ", err)
		return
	}

	result = string(opBytes)
	return
}
