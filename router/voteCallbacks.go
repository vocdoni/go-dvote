package router

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"

	models "github.com/vocdoni/dvote-protobuf/build/go/models"
	"go.vocdoni.io/dvote/log"
	"go.vocdoni.io/dvote/types"
	"go.vocdoni.io/dvote/util"
	"go.vocdoni.io/dvote/vochain/scrutinizer"
	"google.golang.org/protobuf/proto"
)

const MaxListSize = 256
const MaxListIterations = int64(64)

func (r *Router) submitRawTx(request routerRequest) {
	res, err := r.vocapp.SendTX(request.Payload)
	if err != nil {
		r.sendError(request, err.Error())
		return
	}
	if res == nil {
		r.sendError(request, "no reply from Vochain")
		return
	}
	if res.Code != 0 {
		r.sendError(request, string(res.Data))
		return
	}
	log.Infof("broadcasting tx hash:%s", res.Hash)
	var response types.MetaResponse
	response.Payload = fmt.Sprintf("%x", res.Data) // return nullifier or other info
	request.Send(r.buildReply(request, &response))
}

func (r *Router) submitEnvelope(request routerRequest) {
	var err error
	if request.Payload == nil {
		r.sendError(request, "payload is empty")
		return
	}
	// Decode vote envelope
	tx := &models.VoteEnvelope{}
	if err := proto.Unmarshal(request.Payload, tx); err != nil {
		r.sendError(request, fmt.Sprintf("cannot unmarshal payload: (%s)", err))
		return
	}

	// Prepare Vote transaction
	vtx := models.Tx{Payload: &models.Tx_Vote{Vote: tx}}

	// Signature if available
	if request.Signature != "" {
		vtx.Signature, err = hex.DecodeString(util.TrimHex(request.Signature))
		if err != nil {
			r.sendError(request, fmt.Sprintf("cannot unmarshal signature: (%s)", err))
			return
		}
	}

	// Encode and forward the transaction to the Vochain mempool
	txBytes, err := proto.Marshal(&vtx)
	if err != nil {
		r.sendError(request, fmt.Sprintf("cannot marshal vote transaction: (%s)", err))
		return
	}

	res, err := r.vocapp.SendTX(txBytes)
	if err != nil || res == nil {
		r.sendError(request, fmt.Sprintf("cannot broadcast transaction: (%s)", err))
		return
	}

	// Get mempool checkTx reply
	if res.Code != 0 {
		r.sendError(request, string(res.Data))
		return
	}
	log.Infof("broadcasting vochain tx hash:%s code:%d", res.Hash, res.Code)
	var response types.MetaResponse
	response.Nullifier = fmt.Sprintf("%x", res.Data)
	request.Send(r.buildReply(request, &response))
}

func (r *Router) getEnvelopeStatus(request routerRequest) {
	// check pid
	request.ProcessID = util.TrimHex(request.ProcessID)
	if !util.IsHexEncodedStringWithLength(request.ProcessID, types.ProcessIDsize) {
		r.sendError(request, "cannot get envelope status: (malformed processId)")
		return
	}
	// check nullifier
	request.Nullifier = util.TrimHex(request.Nullifier)
	if !util.IsHexEncodedStringWithLength(request.Nullifier, types.VoteNullifierSize) {
		r.sendError(request, "cannot get envelope status: (malformed nullifier)")
		return
	}

	// Check envelope status and send reply
	var response types.MetaResponse
	response.Registered = types.False

	pid, err := hex.DecodeString(request.ProcessID)
	if err != nil {
		r.sendError(request, "cannot decode processID")
		return
	}
	nullifier, err := hex.DecodeString(request.Nullifier)
	if err != nil {
		r.sendError(request, fmt.Sprintf("cannot decode nullifier: (%s)", err))
		return
	}
	e, err := r.vocapp.State.Envelope(pid, nullifier, true)
	// Warning, error is ignored. We should find a better way to check the envelope status
	if err == nil && e != nil {
		response.Registered = types.True
		response.Nullifier = fmt.Sprintf("%x", e.Nullifier)
		response.Height = &e.Height
		block := r.vocapp.Node.BlockStore().LoadBlock(int64(e.Height))
		if block == nil {
			r.sendError(request, "failed getting envelope block timestamp")
			return
		}
		response.BlockTimestamp = int32(block.Time.Unix())
	}
	request.Send(r.buildReply(request, &response))
}

func (r *Router) getEnvelope(request routerRequest) {
	// check pid
	request.ProcessID = util.TrimHex(request.ProcessID)
	if !util.IsHexEncodedStringWithLength(request.ProcessID, types.ProcessIDsize) {
		r.sendError(request, "cannot get envelope: (malformed processId)")
		return
	}
	// check nullifier
	request.Nullifier = util.TrimHex(request.Nullifier)
	if !util.IsHexEncodedStringWithLength(request.Nullifier, types.VoteNullifierSize) {
		r.sendError(request, "cannot get envelope: (malformed nullifier)")
		return
	}
	pid, err := hex.DecodeString(request.ProcessID)
	if err != nil {
		r.sendError(request, "cannot decode processID")
		return
	}
	nullifier, err := hex.DecodeString(util.TrimHex(request.Nullifier))
	if err != nil {
		r.sendError(request, fmt.Sprintf("cannot decode nullifier: (%s)", err))
		return
	}
	envelope, err := r.vocapp.State.Envelope(pid, nullifier, true)
	if err != nil {
		r.sendError(request, fmt.Sprintf("cannot get envelope: (%s)", err))
		return
	}
	var response types.MetaResponse
	response.Registered = types.True
	response.Payload = base64.StdEncoding.EncodeToString(envelope.VotePackage)
	request.Send(r.buildReply(request, &response))
}

func (r *Router) getEnvelopeHeight(request routerRequest) {
	// check pid
	request.ProcessID = util.TrimHex(request.ProcessID)
	if !util.IsHexEncodedStringWithLength(request.ProcessID, types.ProcessIDsize) {
		r.sendError(request, "cannot get envelope height: (malformed processId)")
		return
	}
	pid, err := hex.DecodeString(request.ProcessID)
	if err != nil {
		r.sendError(request, "cannot decode processID")
		return
	}
	votes := r.vocapp.State.CountVotes(pid, true)
	var response types.MetaResponse
	response.Height = new(uint32)
	*response.Height = votes
	request.Send(r.buildReply(request, &response))
}

func (r *Router) getBlockHeight(request routerRequest) {
	var response types.MetaResponse
	h := uint32(r.vocapp.State.Header(true).Height)
	response.Height = &h
	response.BlockTimestamp = int32(r.vocapp.State.Header(true).Timestamp)
	request.Send(r.buildReply(request, &response))
}

func (r *Router) getProcessList(request routerRequest) {
	// check/sanitize eid and fromId
	request.EntityId = util.TrimHex(request.EntityId)
	if !util.IsHexEncodedStringWithLength(request.EntityId, types.EntityIDsize) &&
		!util.IsHexEncodedStringWithLength(request.EntityId, types.EntityIDsizeV2) {
		r.sendError(request, "cannot get process list: (malformed entityId)")
		return
	}
	if len(request.FromID) > 0 {
		request.FromID = util.TrimHex(request.FromID)
		if !util.IsHexEncodedStringWithLength(request.FromID, types.ProcessIDsize) {
			r.sendError(request, "cannot get process list: (malformed fromId)")
			return
		}
	}
	eid, err := hex.DecodeString(request.EntityId)
	if err != nil {
		r.sendError(request, "cannot decode entityID")
		return
	}
	fromID := []byte{}
	if request.FromID != "" {
		fromID, err = hex.DecodeString(request.FromID)
		if err != nil {
			r.sendError(request, "cannot decode fromID")
			return
		}
	}

	var response types.MetaResponse
	max := request.ListSize
	if max > MaxListIterations || max <= 0 {
		max = MaxListIterations
	}
	processList, err := r.Scrutinizer.ProcessList(eid, fromID, max)
	if err != nil {
		r.sendError(request, fmt.Sprintf("cannot get entity process list: (%s)", err))
		return
	}
	for _, p := range processList {
		response.ProcessList = append(response.ProcessList, fmt.Sprintf("%x", p))
	}
	if len(response.ProcessList) == 0 {
		response.Message = "entity does not exist or has not yet created a process"
		request.Send(r.buildReply(request, &response))
		return
	}

	response.Size = new(int64)
	*response.Size = int64(len(response.ProcessList))
	request.Send(r.buildReply(request, &response))
}

func (r *Router) getProcessCount(request routerRequest) {
	var response types.MetaResponse
	response.Size = new(int64)
	count := r.vocapp.State.CountProcesses(true)
	*response.Size = count
	request.Send(r.buildReply(request, &response))
}

func (r *Router) getScrutinizerEntityCount(request routerRequest) {
	var response types.MetaResponse
	response.Size = new(int64)
	*response.Size = r.Scrutinizer.EntityCount()
	request.Send(r.buildReply(request, &response))
}

func (r *Router) getProcessKeys(request routerRequest) {
	// check pid
	request.ProcessID = util.TrimHex(request.ProcessID)
	if !util.IsHexEncodedStringWithLength(request.ProcessID, types.ProcessIDsize) {
		r.sendError(request, "cannot get envelope height: (malformed processId)")
		return
	}
	pid, err := hex.DecodeString(request.ProcessID)
	if err != nil {
		r.sendError(request, "cannot decode processID")
		return
	}
	process, err := r.vocapp.State.Process(pid, true)
	if err != nil {
		r.sendError(request, fmt.Sprintf("cannot get process encryption public keys: (%s)", err))
		return
	}
	var response types.MetaResponse
	var pubs, privs, coms, revs []types.Key
	for idx, pubk := range process.EncryptionPublicKeys {
		if len(pubk) > 0 {
			pubs = append(pubs, types.Key{Idx: idx, Key: pubk})
		}
	}
	for idx, privk := range process.EncryptionPrivateKeys {
		if len(privk) > 0 {
			privs = append(privs, types.Key{Idx: idx, Key: privk})
		}
	}
	for idx, comk := range process.CommitmentKeys {
		if len(comk) > 0 {
			coms = append(coms, types.Key{Idx: idx, Key: comk})
		}
	}
	for idx, revk := range process.RevealKeys {
		if len(revk) > 0 {
			revs = append(revs, types.Key{Idx: idx, Key: revk})
		}
	}
	response.EncryptionPublicKeys = pubs
	response.EncryptionPrivKeys = privs
	response.CommitmentKeys = coms
	response.RevealKeys = revs
	request.Send(r.buildReply(request, &response))
}

func (r *Router) getEnvelopeList(request routerRequest) {
	request.ProcessID = util.TrimHex(request.ProcessID)
	if !util.IsHexEncodedStringWithLength(request.ProcessID, types.ProcessIDsize) {
		r.sendError(request, "cannot get envelope list: (malformed processId)")
		return
	}
	if request.ListSize > MaxListSize {
		r.sendError(request, fmt.Sprintf("listSize overflow, maximum is %d", MaxListSize))
		return
	}
	if request.ListSize == 0 {
		request.ListSize = 64
	}
	pid, err := hex.DecodeString(request.ProcessID)
	if err != nil {
		r.sendError(request, "cannot decode processID")
		return
	}
	nullifiers := r.vocapp.State.EnvelopeList(pid, request.From, request.ListSize, true)
	var response types.MetaResponse
	strnull := []string{}
	for _, n := range nullifiers {
		strnull = append(strnull, fmt.Sprintf("%x", n))
	}
	response.Nullifiers = &strnull
	request.Send(r.buildReply(request, &response))
}

func (r *Router) getResults(request routerRequest) {
	var err error
	request.ProcessID = util.TrimHex(request.ProcessID)
	if !util.IsHexEncodedStringWithLength(request.ProcessID, types.ProcessIDsize) {
		r.sendError(request, "cannot get results: (malformed processId)")
		return
	}

	var response types.MetaResponse

	// Get process info
	pid, err := hex.DecodeString(request.ProcessID)
	if err != nil {
		r.sendError(request, "cannot decode processID")
		return
	}
	procInfo, err := r.Scrutinizer.ProcessInfo(pid)
	if err != nil {
		log.Warn(err)
		r.sendError(request, err.Error())
		return
	}

	if procInfo.EnvelopeType.Anonymous {
		response.Type = "anonymous"
	} else {
		response.Type = "poll"
	}
	if procInfo.EnvelopeType.EncryptedVotes {
		response.Type = response.Type + " encrypted"
	} else {
		response.Type = response.Type + " open"
	}
	if procInfo.EnvelopeType.Serial {
		response.Type = response.Type + " serial"
	} else {
		response.Type = response.Type + " single"
	}
	response.State = procInfo.Status.String()

	// Get results info
	vr, err := r.Scrutinizer.VoteResult(pid)
	if err != nil && err != scrutinizer.ErrNoResultsYet {
		r.sendError(request, err.Error())
		return
	}
	if err == scrutinizer.ErrNoResultsYet {
		response.Message = scrutinizer.ErrNoResultsYet.Error()
		request.Send(r.buildReply(request, &response))
		return
	}
	if vr == nil {
		r.sendError(request, "unknown problem fetching results")
		return
	}
	response.Results = r.Scrutinizer.GetFriendlyResults(vr)

	// Get number of votes
	votes := r.vocapp.State.CountVotes(pid, true)
	response.Height = new(uint32)
	*response.Height = votes

	request.Send(r.buildReply(request, &response))
}

// finished processes
func (r *Router) getProcListResults(request routerRequest) {
	var response types.MetaResponse
	if request.ListSize > MaxListIterations || request.ListSize <= 0 {
		request.ListSize = MaxListIterations
	}
	var err error
	response.ProcessIDs, err = r.Scrutinizer.ProcessListWithResults(request.ListSize, request.FromID)
	if err != nil {
		r.sendError(request, "cannot decode fromID")
		return
	}
	request.Send(r.buildReply(request, &response))
}

// live processes
func (r *Router) getProcListLiveResults(request routerRequest) {
	var response types.MetaResponse
	if request.ListSize > MaxListIterations || request.ListSize <= 0 {
		request.ListSize = MaxListIterations
	}
	var err error
	response.ProcessIDs, err = r.Scrutinizer.ProcessListWithLiveResults(request.ListSize, request.FromID)
	if err != nil {
		r.sendError(request, err.Error())
		return
	}

	request.Send(r.buildReply(request, &response))
}

// known entities
func (r *Router) getScrutinizerEntities(request routerRequest) {
	var response types.MetaResponse
	if request.ListSize > MaxListIterations || request.ListSize <= 0 {
		request.ListSize = MaxListIterations
	}
	var err error
	response.EntityIDs, err = r.Scrutinizer.EntityList(request.ListSize, request.FromID)
	if err != nil {
		r.sendError(request, err.Error())
		return
	}

	request.Send(r.buildReply(request, &response))
}

func (r *Router) getBlockStatus(request routerRequest) {
	var response types.MetaResponse
	h := uint32(r.vocapp.State.Header(true).Height)
	response.BlockTime = r.vocinfo.BlockTimes()
	response.Height = &h
	response.BlockTimestamp = int32(r.vocapp.State.Header(true).Timestamp)
	request.Send(r.buildReply(request, &response))
}
