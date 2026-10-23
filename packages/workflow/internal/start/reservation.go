package start

import (
	"encoding/json"
	"github.com/lurkshark/fireweed/packages/workflow/model"
	"os"
)

func (g localGit) claimBlob(claim model.Claim) (model.Claim, error) {
	claim.Object = ""
	raw, e := json.Marshal(claim)
	if e != nil {
		return claim, problem("cannot encode reservation: %v", e)
	}
	f, e := os.CreateTemp(g.common, "fireweed-start-claim-")
	if e != nil {
		return claim, problem("cannot create claim input: %v", e)
	}
	defer os.Remove(f.Name())
	if _, e = f.Write(raw); e != nil {
		f.Close()
		return claim, problem("cannot write claim input: %v", e)
	}
	if e = f.Close(); e != nil {
		return claim, problem("cannot close claim input: %v", e)
	}
	claim.Object, e = g.text("hash-object", "-w", "--", f.Name())
	return claim, e
}
func (g localGit) reserve(claim model.Claim) (model.Claim, error) {
	next, e := g.claimBlob(claim)
	if e != nil {
		return next, e
	}
	_, e = g.raw("update-ref", "--no-deref", "refs/fireweed/claims/"+claim.Change, next.Object, "")
	return next, e
}
func (g localGit) refresh(previous model.Claim, state string) (model.Claim, error) {
	next := previous
	next.State = state
	next, e := g.claimBlob(next)
	if e != nil {
		return next, e
	}
	_, e = g.raw("update-ref", "--no-deref", "refs/fireweed/claims/"+previous.Change, next.Object, previous.Object)
	return next, e
}
