package process

import (
	"net/http"
	"os/exec"
	"sync"
	"time"

	log "qiniupkg.com/x/log.v7"

	"github.com/songjiayang/wdog/command"
	"github.com/songjiayang/wdog/config"
)

type Process struct {
	cfg            *config.Process
	checkInterval  time.Duration
	reloadInterval time.Duration

	mux *sync.Mutex
}

func NewProcess(cfg *config.Process) *Process {
	var mux sync.Mutex
	return &Process{
		cfg:            cfg,
		checkInterval:  parseDuration(cfg.CheckInterval),
		reloadInterval: parseDuration(cfg.ReloadInterval),

		mux: &mux,
	}
}

func (p *Process) Run() {
	defer func() {
		if r := recover(); r != nil {
			log.Errorf("process.Run(%s): %v", p.cfg.Name, r)
			p.check()
		}
	}()

	// check interval
	go func() {
		for {
			p.mux.Lock()
			p.check()
			p.mux.Unlock()

			time.Sleep(p.checkInterval)
		}
	}()

	// reload interval
	go func() {
		for {
			p.mux.Lock()
			p.reload(p.pids())
			p.mux.Unlock()

			time.Sleep(p.reloadInterval)
		}
	}()
}

func (p *Process) start() {
	go command.Run(exec.Command(p.cfg.RCmd))
}

func (p *Process) check() {
	pids := p.pids()

	// is stop
	if isEmpty(pids) {
		p.start()
		return
	}

	if p.isHalt() {
		p.reload(pids)
	}
}

func (p *Process) reload(pids []string) {
	if !isEmpty(pids) {
		for _, pid := range pids {
			command.Run(exec.Command("tskill.exe", pid))
		}
	}

	p.start()
}

func (p *Process) isHalt() (ok bool) {
	req, err := http.NewRequest(http.MethodGet, p.cfg.Endpoint, nil)
	if err != nil {
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode/100 == 5 {
		return true
	}

	return
}

func (p *Process) pids() []string {
	result, err := command.Run(exec.Command("qprocess.exe"))
	if err != nil {
		return nil
	}

	return findPids(result, p.cfg.Name)
}
