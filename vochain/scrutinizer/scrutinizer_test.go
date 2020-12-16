package scrutinizer

import (
	"testing"

	"github.com/tendermint/go-amino"
	cryptoamino "github.com/tendermint/tendermint/crypto/encoding/amino"
	"gitlab.com/vocdoni/go-dvote/log"
	"gitlab.com/vocdoni/go-dvote/util"
	"gitlab.com/vocdoni/go-dvote/vochain"
)

func TestProcessList(t *testing.T) {
	testProcessList(t, 10)
	testProcessList(t, 20)
	testProcessList(t, 155)
}

func testProcessList(t *testing.T, procsCount int) {
	log.Init("info", "stdout")
	cdc := amino.NewCodec()
	cryptoamino.RegisterAmino(cdc)

	state, err := vochain.NewState(t.TempDir(), cdc)
	if err != nil {
		t.Fatal(err)
	}

	sc, err := NewScrutinizer(t.TempDir(), state)
	if err != nil {
		t.Fatal(err)
	}

	// Add 10 entities and process for storing random content
	for i := 0; i < 10; i++ {
		sc.addEntity(util.RandomHex(20), util.RandomHex(32))
	}

	// For a entity, add 25 processes (this will be the queried entity)
	eidTest := util.RandomHex(20)
	for i := 0; i < procsCount; i++ {
		sc.addEntity(eidTest, util.RandomHex(32))
	}

	procs := make(map[string]bool)
	last := ""
	var list []string
	for len(procs) < procsCount {
		list, err = sc.ProcessList(eidTest, last, 10)
		if err != nil {
			t.Fatal(err)
		}
		if len(list) < 1 {
			t.Log("list is empty")
			break
		}
		for _, p := range list {
			if procs[string(p)] {
				t.Fatalf("found duplicated entity: %x", p)
			}
			procs[string(p)] = true
		}
		last = list[len(list)-1]
	}
	if len(procs) != procsCount {
		t.Fatalf("expected %d processes, got %d", procsCount, len(procs))
	}
}
