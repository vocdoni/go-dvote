package processarchive

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	"go.vocdoni.io/dvote/data"
	"go.vocdoni.io/dvote/log"
	"go.vocdoni.io/dvote/types"
	"go.vocdoni.io/dvote/vochain"
	"go.vocdoni.io/proto/build/go/models"
)

type ProcessArchive struct {
	vochain     *vochain.BaseApplication
	ipfs        *data.IPFSHandle
	storage     *jsonStorage
	pprocs      []*Process
	publish     chan (bool)
	lastUpdate  time.Time
	close       chan (bool)
	publishLock sync.Mutex
}

type Process struct {
	Process *models.Process
	Votes   uint32
}

type jsonStorage struct {
	datadir string
	lock    sync.RWMutex
}

func NewJsonStorage(datadir string) (*jsonStorage, error) {
	err := os.MkdirAll(datadir, 0o750)
	if err != nil {
		return nil, err
	}
	return &jsonStorage{datadir: datadir}, nil
}

func (js *jsonStorage) AddProcess(p *Process) error {
	if p == nil || p.Process == nil || len(p.Process.ProcessId) != types.ProcessIDsize {
		return fmt.Errorf("process not valid")
	}
	data, err := json.Marshal(p)
	if err != nil {
		return err
	}
	js.lock.Lock()
	defer js.lock.Unlock()
	return os.WriteFile(filepath.Join(js.datadir, fmt.Sprintf("%x", p.Process.ProcessId)), data, 0o644)
}

// NewProcessArchive creates a new instance of the process archiver.
// It will subscribe to Vochain events and perform the process archival.
// JSON files (one per process) will be stored within datadir.
// The key parameter must be either a valid IPFS base64 encoded private key
// or empty (a new key will be generated).
// If ipfs is nil, only JSON archive storage will be performed.
func NewProcessArchive(v *vochain.BaseApplication, ipfs *data.IPFSHandle,
	datadir, key string) (*ProcessArchive, error) {
	js, err := NewJsonStorage(datadir)
	if err != nil {
		return nil, fmt.Errorf("could not create process archive: %w", err)
	}
	ir := &ProcessArchive{
		vochain: v,
		ipfs:    ipfs,
		storage: js,
		publish: make(chan (bool), 10),
		close:   make(chan (bool), 1),
	}
	if ipfs != nil {
		if err := ir.AddKey(key); err != nil {
			return nil, err
		}
		if pk, err := ir.GetKey(); err != nil {
			return nil, err
		} else {
			log.Infof("using IPNS privkey: %s", pk)
		}
		go ir.Publish()
	}
	v.State.AddEventListener(ir)
	return ir, nil
}

func (i *ProcessArchive) Rollback() {
	i.pprocs = []*Process{}
}

func (i *ProcessArchive) Commit(height uint32) error {
	for _, p := range i.pprocs {
		if err := i.storage.AddProcess(p); err != nil {
			log.Errorf("cannot add json process: %v", err)
			continue
		}
		log.Infof("stored json process %x", p.Process.ProcessId)
	}
	// publish to IPFS if there is a new process with results
	if len(i.pprocs) > 0 && i.ipfs != nil {
		log.Debugf("sending archive publish signal for height %d", height)
		i.publish <- true
	}
	return nil
}

func (i *ProcessArchive) OnProcessResults(pid []byte,
	results []*models.QuestionResult, txindex int32) error {
	process, err := i.vochain.State.Process(pid, false)
	if err != nil {
		return fmt.Errorf("cannot get process %x info: %w", pid, err)
	}
	process.Results = &models.ProcessResult{Votes: results, ProcessId: pid}
	i.pprocs = append(i.pprocs, &Process{
		Votes:   i.vochain.State.CountVotes(pid, false),
		Process: process,
	})
	return nil
}

func (i *ProcessArchive) Close() {
	i.close <- true
}

// NOT USED but required for implementing the interface
func (i *ProcessArchive) OnCancel(pid []byte, txindex int32)                       {}
func (i *ProcessArchive) OnVote(v *models.Vote, txindex int32)                     {}
func (i *ProcessArchive) OnProcessKeys(pid []byte, pub, com string, txindex int32) {}
func (i *ProcessArchive) OnRevealKeys(pid []byte, priv, rev string, txindex int32) {}
func (i *ProcessArchive) OnProcessStatusChange(pid []byte,
	status models.ProcessStatus, txindex int32) {
}
func (i *ProcessArchive) OnProcess(pid, eid []byte, censusRoot, censusURI string, txindex int32) {}
