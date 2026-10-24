package evals

import (
	"context"
	w "github.com/lurkshark/fireweed/packages/workflow"
	"os"
	"reflect"
	"strings"
	"testing"
	"time"
)

func Test_APP028VerifyReadonlyControls(t *testing.T) {
	for _, mode := range []string{"success-control", "structured-rejection-control", "fatal-control"} {
		t.Run(mode, func(t *testing.T) {
			f := cliVerifyFixture(t, "impl", "sha1", "all", false, false)
			if mode == "structured-rejection-control" {
				projectPut(t, f.Root, "units/a/evals/case.txt", "forbidden")
			}
			var reset func()
			if mode == "fatal-control" {
				reset = cliVerifyDiscoveryFault(t, f)
			}
			want, publicErr := w.Verify(context.Background(), f.Where, "current")
			reports := cliVerifyReports(t, f)
			cliVerifyClear(t, f)
			if mode == "fatal-control" {
				if publicErr == nil || !strings.HasPrefix(publicErr.Error(), "C01:") {
					t.Fatal("observer public fatal premise", want, publicErr)
				}
			} else if publicErr != nil {
				t.Fatal(publicErr)
			}
			if mode == "success-control" && (cliVerifyExit(want) != 0 || len(want.Coverage) != 3) {
				t.Fatal("observer success premise", want)
			}
			if mode == "structured-rejection-control" && (cliVerifyExit(want) != 1 || len(want.Coverage) != 0) {
				t.Fatal("observer rejection premise", want)
			}
			bin, e := os.Executable()
			if e != nil {
				t.Fatal(e)
			}
			for repeat := 0; repeat < 2; repeat++ {
				if reset != nil {
					reset()
				}
				env := os.Environ()
				cliVerifyReadonly(t, f, func() {
					ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
					defer cancel()
					began := time.Now()
					got, e := w.Verify(ctx, f.Where, "current")
					if ctx.Err() != nil || time.Since(began) >= 5*time.Second {
						t.Fatal("observer prevents bounded public Verify", repeat, e, ctx.Err())
					}
					if (e != nil) != (publicErr != nil) {
						t.Fatal("observer changed public fatal selection", e, publicErr)
					}
					if e != nil && e.Error() != publicErr.Error() {
						t.Fatal("observer changed full public fatal message", e, publicErr)
					}
					actual := cliVerifyReports(t, f)
					if !reflect.DeepEqual(cliVerifyWire(t, got, actual), cliVerifyWire(t, want, reports)) {
						t.Error("observer changed full public evidence", repeat, got, want)
					}
					cliVerifyObserveVerification(t, f, want, bin)
				})
				if !reflect.DeepEqual(env, os.Environ()) {
					t.Fatal("readonly observer retained its environment", repeat)
				}
				cliVerifyClear(t, f)
			}
		})
	}
}
