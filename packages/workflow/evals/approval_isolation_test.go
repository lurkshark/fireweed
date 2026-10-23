package evals

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"

	p "github.com/lurkshark/fireweed/packages/project"
	w "github.com/lurkshark/fireweed/packages/workflow"
)

func Test_FLOW020_FLOW022_FLOW023ApprovalNativeIsolation(t *testing.T) {
	if runtime.GOOS != "darwin" {
		t.Skip("macOS native read/ignored-error controls; portable approval tests run elsewhere")
	}
	temp := stateCanonicalDirectory(t, t.TempDir())
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Minute)
	defer cancel()
	guard := `#include <sys/types.h>
#include <sys/stat.h>
#include <sys/socket.h>
#include <sys/attr.h>
#include <fcntl.h>
#include <unistd.h>
#include <stdarg.h>
#include <stdint.h>
#include <string.h>
#include <dirent.h>
#include <stdio.h>
#include <stdlib.h>
#include <limits.h>
static int armed=0;static const char*root=0;static const char*git=0;static const char*extra=0;
void fw_approval_arm(int m,const char*r,const char*g,const char*b){armed=m;root=r;git=g;extra=b;}
static int prefix(const char*p,const char*q,size_t n){return p&&q&&n&&strncmp(p,q,n)==0&&(p[n]==0||p[n]=='/');}
static void path_for(int d,const char*p,char*out){
 char abs[PATH_MAX],base[PATH_MAX],probe[PATH_MAX],tail[PATH_MAX]="",resolved[PATH_MAX];
 if(p[0]=='/'){if(snprintf(abs,sizeof(abs),"%s",p)>=(int)sizeof(abs))_exit(42);}
 else {if(d==AT_FDCWD){if(!getcwd(base,sizeof(base)))_exit(42);}else{if(fcntl(d,F_GETPATH,base)<0)_exit(42);}
  if(snprintf(abs,sizeof(abs),"%s/%s",base,p)>=(int)sizeof(abs))_exit(42);}
 if(!realpath(abs,resolved)){
  snprintf(probe,sizeof(probe),"%s",abs);
  while(!realpath(probe,resolved)){
   char*slash=strrchr(probe,'/');if(!slash)_exit(42);
   char next[PATH_MAX];if(snprintf(next,sizeof(next),"/%s%s",slash+1,tail)>=(int)sizeof(next))_exit(42);
   snprintf(tail,sizeof(tail),"%s",next);if(slash==probe)probe[1]=0;else *slash=0;
  }
  if(snprintf(abs,sizeof(abs),"%s%s",resolved,tail)>=(int)sizeof(abs))_exit(42);
 }else{snprintf(abs,sizeof(abs),"%s",resolved);}
 // Normalize missing suffix components as well as existing aliases.
 char*save=0;size_t n=0;out[n++]='/';out[n]=0;
 for(char*v=strtok_r(abs,"/",&save);v;v=strtok_r(0,"/",&save)){
  if(!strcmp(v,"."))continue;
  if(!strcmp(v,"..")){while(n>1&&out[n-1]!='/')n--;if(n>1)n--;out[n]=0;continue;}
  if(n>1)out[n++]='/';size_t len=strlen(v);if(n+len>=PATH_MAX)_exit(42);memcpy(out+n,v,len+1);n+=len;
 }
}
static void content_at(int d,const char*p){if(armed==2)_exit(42);if(!armed||!p)return;char path[PATH_MAX];path_for(d,p,path);p=path;char meta[PATH_MAX];snprintf(meta,sizeof(meta),"%s/.git",root?root:"");if(root&&prefix(p,root,strlen(root))&&strcmp(p,meta)&&!(git&&prefix(p,git,strlen(git))))_exit(42);const char*c=extra;while(c&&*c){const char*e=strchr(c,'\n');size_t n=e?(size_t)(e-c):strlen(c);if(prefix(p,c,n))_exit(42);if(!e)break;c=e+1;}}
static void content(const char*p){content_at(AT_FDCWD,p);}
int fw_approval_prepare(const char*p){char dir[PATH_MAX];snprintf(dir,sizeof(dir),"%s",p);char*slash=strrchr(dir,'/');if(!slash)return -1;*slash=0;return open(dir,O_RDONLY|O_DIRECTORY);}

static void mutation(void){if(armed)_exit(42);}
static int fw_open(const char*p,int f,...){content(p);if(f&(O_WRONLY|O_RDWR|O_CREAT|O_TRUNC|O_APPEND))mutation();mode_t m=0;if(f&O_CREAT){va_list a;va_start(a,f);m=va_arg(a,int);va_end(a);}return open(p,f,m);}
static int fw_openat(int d,const char*p,int f,...){content_at(d,p);if(f&(O_WRONLY|O_RDWR|O_CREAT|O_TRUNC|O_APPEND))mutation();mode_t m=0;if(f&O_CREAT){va_list a;va_start(a,f);m=va_arg(a,int);va_end(a);}return openat(d,p,f,m);}
static ssize_t fw_read(int d,void*b,size_t n){if(armed==2||(armed&&d==0))_exit(42);return read(d,b,n);}
static ssize_t fw_write(int d,const void*b,size_t n){mutation();return write(d,b,n);}
static int fw_socket(int d,int t,int p){mutation();return socket(d,t,p);}
static int fw_connect(int s,const struct sockaddr*a,socklen_t n){mutation();return connect(s,a,n);}
static int fw_mkdir(const char*p,mode_t m){mutation();return mkdir(p,m);}
static int fw_unlink(const char*p){mutation();return unlink(p);}
static int fw_rmdir(const char*p){mutation();return rmdir(p);}
static int fw_rename(const char*p,const char*q){mutation();return rename(p,q);}
static int fw_symlink(const char*p,const char*q){mutation();return symlink(p,q);}
static int fw_chmod(const char*p,mode_t m){mutation();return chmod(p,m);}
static DIR*fw_opendir(const char*p){content(p);return opendir(p);}
#define I(r,o) __attribute__((used)) static struct{const void*replacement;const void*original;} bi_##o __attribute__((section("__DATA,__interpose")))={(const void*)(uintptr_t)&r,(const void*)(uintptr_t)&o};
I(fw_open,open) I(fw_openat,openat) I(fw_read,read) I(fw_write,write) I(fw_socket,socket) I(fw_connect,connect) I(fw_mkdir,mkdir) I(fw_unlink,unlink) I(fw_rmdir,rmdir) I(fw_rename,rename) I(fw_symlink,symlink) I(fw_chmod,chmod) I(fw_opendir,opendir)
static int fw_stat(const char*p,struct stat*b){if(armed==2)_exit(42);return stat(p,b);}
static int fw_lstat(const char*p,struct stat*b){if(armed==2)_exit(42);return lstat(p,b);}
static char*fw_getcwd(char*b,size_t n){if(armed==2)_exit(42);return getcwd(b,n);}
static ssize_t fw_readlink(const char*p,char*b,size_t n){if(armed==2)_exit(42);return readlink(p,b,n);}
I(fw_stat,stat) I(fw_lstat,lstat) I(fw_getcwd,getcwd) I(fw_readlink,readlink)
`
	worktreePut(t, temp, "guard.c", guard)
	lib := filepath.Join(temp, "libapprovalguard.dylib")
	cmd := exec.CommandContext(ctx, "cc", "-dynamiclib", "-o", lib, filepath.Join(temp, "guard.c"), "-Wl,-install_name,"+lib)
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("approval native guard build %v %s", err, out)
	}
	program := fmt.Sprintf(`package main
/*
#cgo LDFLAGS: -L%s -lapprovalguard
void fw_approval_arm(int m,const char*r,const char*g,const char*b);
int fw_approval_prepare(const char*p);
#include <fcntl.h>
#include <unistd.h>
#include <string.h>
static int approval_readat(int d,const char*p){const char*base=strrchr(p,'/');int f=openat(d,base?base+1:p,O_RDONLY);if(f<0)return 0;close(f);return 1;}
*/
import "C"
import("encoding/json";"os";"fmt";"strings";"reflect";"net";"time";p "github.com/lurkshark/fireweed/packages/project";w "github.com/lurkshark/fireweed/packages/workflow")
type record struct {Root,Start,Slug,Git,Blocked,Error,Object string;Expected p.Snapshot;Claim w.Claim;Current bool}
func main(){var in record;raw,err:=os.ReadFile(os.Args[2]);if err!=nil||json.Unmarshal(raw,&in)!=nil {os.Exit(80)}
 mode:=os.Args[1];target:=os.Args[3]
 if in.Claim.Object!="" {os.Exit(88)};in.Claim.Object=in.Object
 if mode=="fixture-control" {if in.Root==""||in.Expected.Modules==nil||in.Expected.Diagnostics==nil||in.Object!=target||((in.Claim.Change=="")!=(in.Object=="")) {os.Exit(81)};return}
 fd:=-1;if mode=="readat-control"||mode=="allowed-readat" {fd=int(C.fw_approval_prepare(C.CString(target)));if fd<0 {os.Exit(82)}}
 root:=C.CString(in.Root);if in.Current {root=nil}
 arm:=1;if mode=="no-io"||mode=="metadata-control" {arm=2}
 C.fw_approval_arm(C.int(arm),root,C.CString(in.Git),C.CString(in.Blocked))
 switch mode {case "read-control":os.ReadFile(target);os.Exit(75);case "readat-control":C.approval_readat(C.int(fd),C.CString(target));os.Exit(76);case "allowed-readat":ok:=C.approval_readat(C.int(fd),C.CString(target));C.fw_approval_arm(0,nil,nil,nil);if ok!=1 {os.Exit(77)};return;case "network-control":net.DialTimeout("tcp",target,time.Second);os.Exit(78);case "stdin-control":b:=make([]byte,1);os.Stdin.Read(b);os.Exit(79);case "prompt-control":fmt.Print("prompt");os.Exit(83);case "write-control":os.WriteFile(target,[]byte("forbidden"),0600);os.Exit(84);case "metadata-control":os.Stat(target);os.Exit(85)}
 if mode=="source-control" {var source p.Snapshot;if in.Current {source,err=p.Load(in.Start)}else{source,err=p.LoadAt(in.Start,in.Claim.Base)};C.fw_approval_arm(0,nil,nil,nil);if err!=nil||source.Root!=in.Root||len(source.Modules)==0 {fmt.Fprint(os.Stderr,err);os.Exit(86)};return}
 if mode=="claim-control" {cl,err:=w.LoadClaim(in.Start,"current");C.fw_approval_arm(0,nil,nil,nil);if err!=nil||cl!=in.Claim {fmt.Fprint(os.Stderr,err);os.Exit(87)};return}
 got,cl,err:=w.LoadApproval(in.Start,in.Slug)
 C.fw_approval_arm(0,nil,nil,nil)
 if in.Error!="" {if err==nil||!strings.Contains(err.Error(),in.Error)||!reflect.DeepEqual(got,p.Snapshot{})||cl!=(w.Claim{}) {fmt.Fprintf(os.Stderr,"wrong rejection %%v %%+v %%+v",err,got,cl);os.Exit(10)};return}
 if err!=nil||cl!=in.Claim||!reflect.DeepEqual(got,in.Expected) {fmt.Fprintf(os.Stderr,"wrong approved source %%v %%+v",err,cl);os.Exit(11)}
}`, temp)
	worktreePut(t, temp, "probe.go", program)
	sourceRoot, err := filepath.Abs("../../..")
	if err != nil {
		t.Fatal(err)
	}
	binary := filepath.Join(temp, "probe")
	cmd = exec.CommandContext(ctx, "go", "build", "-o", binary, filepath.Join(temp, "probe.go"))
	cmd.Dir = sourceRoot
	cmd.Env = append(os.Environ(), "CGO_ENABLED=1", "GOPROXY=off", "GOTOOLCHAIN=local")
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("approval public probe build %v %s", err, out)
	}
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer listener.Close()
	fixture := approvalFixture(t, "impl", "sha1", true)
	input := func(f approvalFixtureData, current bool, code, slug string, extra []string) string {
		t.Helper()
		value := map[string]any{"Root": f.Root, "Start": f.Root, "Slug": slug, "Git": claimCommon(t, f.Root), "Blocked": strings.Join(extra, "\n"), "Expected": f.Want, "Claim": f.Claim, "Object": f.Claim.Object, "Error": code, "Current": current}
		raw, err := json.Marshal(value)
		if err != nil {
			t.Fatal(err)
		}
		path := filepath.Join(temp, "record.json")
		if err = os.WriteFile(path, raw, 0644); err != nil {
			t.Fatal(err)
		}
		return path
	}
	run := func(root, mode, record, target string) ([]byte, error) {
		cmd := exec.CommandContext(ctx, binary, mode, record, target)
		cmd.Dir = root
		return cmd.CombinedOutput()
	}
	record := input(fixture, false, "", "current", nil)
	if out, err := run(fixture.Root, "fixture-control", record, fixture.Claim.Object); err != nil {
		t.Fatalf("serialized fixture control %v %s", err, out)
	}
	for _, path := range []string{"fireweed.yaml", "units/a/SPEC.md", "changes/0002-current.md", "changes/archive/0009-earlier.md", "SYSTEM.md", "units/a/e/current.txt", "units/a/impl/private.txt", "private/secret.txt", "shared/contract.txt", ".fireweed/playbooks/impl.md"} {
		t.Run("ignored-read/"+path, func(t *testing.T) {
			for _, target := range []string{filepath.Join(fixture.Root, path), path, "./" + path} {
				out, err := run(fixture.Root, "read-control", record, target)
				exit, ok := err.(*exec.ExitError)
				if !ok || exit.ExitCode() != 42 || len(out) != 0 {
					t.Fatalf("ignored native read control %q %v %s", target, err, out)
				}
			}
		})
	}
	t.Run("fd-read-controls", func(t *testing.T) {
		out, err := run(fixture.Root, "readat-control", record, filepath.Join(fixture.Root, "units/a/e/current.txt"))
		exit, ok := err.(*exec.ExitError)
		if !ok || exit.ExitCode() != 42 || len(out) != 0 {
			t.Fatalf("fd ignored read control %v %s", err, out)
		}
		if out, err = run(fixture.Root, "allowed-readat", record, filepath.Join(claimCommon(t, fixture.Root), "config")); err != nil {
			t.Fatalf("allowed Git fd control %v %s", err, out)
		}
	})
	for _, mode := range []string{"network-control", "stdin-control", "prompt-control", "write-control", "metadata-control"} {
		t.Run(mode, func(t *testing.T) {
			target := filepath.Join(temp, "forbidden")
			if mode == "network-control" {
				target = listener.Addr().String()
			}
			out, err := run(fixture.Root, mode, record, target)
			exit, ok := err.(*exec.ExitError)
			if !ok || exit.ExitCode() != 42 || len(out) != 0 {
				t.Fatalf("native control %s: %v %s", mode, err, out)
			}
		})
	}
	for _, kind := range []string{"claimed", "unclaimed", "linked", "missing-current", "foreign", "reservation", "malformed", "invalid-selector", "current-escape", "baseline-escape"} {
		t.Run(kind, func(t *testing.T) {
			f := approvalFixture(t, "impl", "sha1", true)
			current, code, slug, mode := false, "", "current", "check"
			extra := []string{}
			outside := stateCanonicalDirectory(t, t.TempDir())
			worktreePut(t, outside, "private", "outside must never open")
			if kind == "linked" {
				f.Root = claimPeer(t, f.Root)
				f.Claim.Worktree = f.Root
				f.Claim.Object = claimRaw(t, f.Root, "current", claimPayload(t, f.Claim))
				f.Want.Root = f.Root
			}
			switch kind {
			case "unclaimed", "current-escape":
				current = true
				claimGit(t, f.Root, "update-ref", "-d", "refs/fireweed/claims/current", f.Claim.Object)
				if kind == "current-escape" {
					if err := os.Remove(filepath.Join(f.Root, "units/a/SPEC.md")); err != nil {
						t.Fatal(err)
					}
					if err := os.Symlink(filepath.Join(outside, "private"), filepath.Join(f.Root, "units/a/SPEC.md")); err != nil {
						t.Fatal(err)
					}
					code = "C05"
				}
				var err error
				f.Want, err = p.Load(f.Root)
				if err != nil {
					t.Fatal(err)
				}
				f.Claim = w.Claim{}
				for _, path := range []string{"SYSTEM.md", "private/secret.txt", "shared/contract.txt", "units/a/e/current.txt", "units/a/impl/private.txt", "units/a/pub/contract ' λ.txt", ".fireweed/playbooks/impl.md"} {
					extra = append(extra, filepath.Join(f.Root, path))
				}
			case "missing-current":
				if err := os.RemoveAll(filepath.Join(f.Root, "changes")); err != nil {
					t.Fatal(err)
				}
			case "foreign":
				f.Claim.Worktree = outside
				f.Claim.Object = claimRaw(t, f.Root, "current", claimPayload(t, f.Claim))
				code = "P02"
			case "reservation":
				f.Claim.State = ""
				f.Claim.Object = claimRaw(t, f.Root, "current", claimPayload(t, f.Claim))
				code = "P02"
			case "malformed":
				claimRaw(t, f.Root, "current", "invalid")
				code = "P02"
			case "invalid-selector":
				slug = "../bad"
				code = "C05"
				mode = "no-io"
			case "baseline-escape":
				if err := os.Remove(filepath.Join(f.Root, "units/a/SPEC.md")); err != nil {
					t.Fatal(err)
				}
				if err := os.Symlink(filepath.Join(outside, "private"), filepath.Join(f.Root, "units/a/SPEC.md")); err != nil {
					t.Fatal(err)
				}
				f.Claim.Base = approvalCommit(t, f.Root)
				f.Claim.Object = claimRaw(t, f.Root, "current", claimPayload(t, f.Claim))
				code = "C05"
			}
			extra = append(extra, outside)
			record := input(f, current, code, slug, extra)
			// Canonical outside and missing endpoints are intercepted even when ignored.
			for _, target := range []string{filepath.Join(outside, "private"), filepath.Join(outside, "missing")} {
				out, err := run(f.Root, "read-control", record, target)
				exit, ok := err.(*exec.ExitError)
				if !ok || exit.ExitCode() != 42 || len(out) != 0 {
					t.Fatalf("outside native control %v %s", err, out)
				}
			}

			fixtureBase := f.Claim.Base
			if fixtureBase == "" {
				fixtureBase = claimBase(t, f.Root)
			}

			// Inspect object bodies only by approved identity, never by current revision/path.
			real := protocolRealGit(t)
			tree := protocolGit(t, real, f.Root, "", "--no-replace-objects", "ls-tree", "-r", "-z", fixtureBase)
			allowed := []string{fixtureBase}
			if !current {
				allowed = append(allowed, strings.TrimSpace(protocolGit(t, real, f.Root, "", "rev-parse", "refs/fireweed/claims/current")))
			}
			for _, line := range strings.Split(strings.TrimSuffix(tree, "\x00"), "\x00") {
				header, path, ok := strings.Cut(line, "\t")
				fields := strings.Fields(header)
				if !ok || len(fields) != 3 {
					t.Fatal("Git tree fixture record", line)
				}
				if path == "fireweed.yaml" || (strings.HasPrefix(path, "units/") && strings.Count(path, "/") == 2 && strings.HasSuffix(path, "/SPEC.md")) || (strings.HasPrefix(path, "changes/") && strings.Count(path, "/") == 1) || (strings.HasPrefix(path, "changes/archive/") && strings.Count(path, "/") == 2) {
					allowed = append(allowed, fields[2])
				}
			}
			private := strings.TrimSpace(protocolGit(t, real, f.Root, "", "rev-parse", fixtureBase+":units/a/impl/private.txt"))
			loose := filepath.Join(claimCommon(t, f.Root), "objects", private[:2], private[2:])
			if raw, err := os.ReadFile(loose); err != nil || len(raw) == 0 {
				t.Fatal("private loose object fixture", err)
			}
			extra = append(extra, loose)
			record = input(f, current, code, slug, extra)
			if out, err := run(f.Root, "fixture-control", record, f.Claim.Object); err != nil {
				t.Fatalf("explicit expected token serialization control %v %s", err, out)
			}
			for _, readMode := range []string{"read-control", "readat-control"} {
				out, err := run(f.Root, readMode, record, loose)
				exit, ok := err.(*exec.ExitError)
				if !ok || exit.ExitCode() != 42 || len(out) != 0 {
					t.Fatalf("private loose object ignored read control %v %s", err, out)
				}
			}

			dir := t.TempDir()
			marker := filepath.Join(dir, "forbidden")
			reject := "printf attempted >> " + claimQuote(marker) + "; exit 93"
			script := "#!/bin/sh\nfor a in \"$@\"; do case \"$a\" in fetch|push|clone|ls-remote|commit|branch|checkout|switch|reset|restore|add|worktree|update-ref|hash-object|show|archive|refs/fireweed/claims|refs/fireweed/claims/) " + reject + ";; esac; done\nbody=0; tree=0; last=''; for a in \"$@\"; do [ \"$a\" = cat-file ] && body=1; [ \"$a\" = ls-tree ] && tree=1; last=$a; done\nif [ \"$body\" = 1 ]; then case \"$last\" in " + strings.Join(allowed, "|") + ") ;; *) " + reject + ";; esac; fi\nif [ \"$tree\" = 1 ] && [ \"$last\" != " + claimQuote(fixtureBase) + " ]; then " + reject + "; fi\nexec " + claimQuote(real) + " \"$@\"\n"
			if kind == "invalid-selector" {
				script = "#!/bin/sh\n" + reject + "\n"
			}
			worktreePut(t, dir, "git", script)
			if err := os.Chmod(filepath.Join(dir, "git"), 0755); err != nil {
				t.Fatal(err)
			}
			out, _, err := protocolCommand(t, filepath.Join(dir, "git"), f.Root, "", "cat-file", "blob", private)
			if err == nil || out != "" {
				t.Fatal("unselected body denial control", err, out)
			}
			if _, err := os.Stat(marker); err != nil {
				t.Fatal("body denial control not reached")
			}
			if out, _, err := protocolCommand(t, filepath.Join(dir, "git"), f.Root, "", "for-each-ref", "--format=%(refname)", "refs/fireweed/claims/"); err == nil || out != "" {
				t.Fatal("claim inventory denial control", err, out)
			}

			if err := os.Remove(marker); err != nil {
				t.Fatal(err)
			}
			t.Setenv("PATH", dir+string(os.PathListSeparator)+os.Getenv("PATH"))
			defer func() {
				if _, err := os.Stat(marker); !os.IsNotExist(err) {
					t.Fatal("approval attempted a forbidden Git body/revision/command")
				}
			}()
			if kind != "invalid-selector" {
				if out, err := run(f.Root, "source-control", record, ""); err != nil {
					t.Fatalf("allowed public source control %v %s", err, out)
				}
				if !current && kind != "malformed" {
					if out, err := run(f.Root, "claim-control", record, ""); err != nil {
						t.Fatalf("allowed exact claim control %v %s", err, out)
					}
				}
			}
			if out, err := run(f.Root, mode, record, ""); err != nil {
				t.Fatalf("approval native %s %v %s", kind, err, out)
			}
		})
	}
}
