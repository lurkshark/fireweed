package evals

import (
	"bytes"
	"context"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"testing"
	"time"

	p "github.com/lurkshark/fireweed/packages/project"
	w "github.com/lurkshark/fireweed/packages/workflow"
)

// Both seams observe public Git data. The native seam observes actual loose
// object reads and the subsequent ref inventory. The Git seam observes source
// bytes on stdout and gates the next invocation or streaming batch request.
// Neither seam identifies a product function or mandates a body-read command.
// The mmap is test-owned transport outside the fixture, never product state.
const heldRaceGuard = `#include <sys/types.h>
#include <sys/mman.h>
#include <sys/stat.h>
#include <fcntl.h>
#include <unistd.h>
#include <dirent.h>
#include <stdarg.h>
#include <stdint.h>
#include <stdlib.h>
#include <string.h>
#include <time.h>
#include <stdatomic.h>
#include <errno.h>
#include <spawn.h>
#include <dlfcn.h>
static _Atomic(uint64_t)*state;
static char objects[32][4096],common[4096];static int count;
static char original_git[4096],engine_git[4096],traced_git[4096];
static char native_policy[4096],library_path[4096],git_exec_path[4096];static _Thread_local int resolving;
static int within(const char*p,const char*q){size_t n=strlen(q);return p&&!strncmp(p,q,n)&&(p[n]==0||p[n]=='/');}
static void gate(void){
 if(!state||atomic_load(state)!=atomic_load(state+4))return;
 uint64_t zero=0;atomic_compare_exchange_strong(state+1,&zero,1);
 struct timespec pause={0,1000000};int tries=0;while(!atomic_load(state+2)){if(++tries>10000)_exit(94);nanosleep(&pause,0);}
}
static int ownership(const char*p){if(!state||!p)return 0;char refs[4096],packed[4096];snprintf(refs,sizeof(refs),"%s/refs/fireweed/claims",common);snprintf(packed,sizeof(packed),"%s/packed-refs",common);return within(p,refs)||!strcmp(p,packed);}
static void inventory(const char*p){if(!state||!p||resolving)return;char canonical[4096];resolving=1;if(realpath(p,canonical))p=canonical;resolving=0;if(ownership(p))gate();}
static void descriptor_inventory(int fd){if(!state)return;char path[4096];if(fcntl(fd,F_GETPATH,path)>=0)inventory(path);}
static void descriptor(int fd){if(!state)return;char path[4096];if(fcntl(fd,F_GETPATH,path)<0)return;if(ownership(path)&&atomic_load(state+1))atomic_fetch_add(state+3,1);for(int i=0;i<count;i++)if(!strcmp(path,objects[i]))atomic_fetch_or(state,((uint64_t)1)<<i);}
void held_race_arm(const char*file,const char*git,const char*paths){
 int fd=open(file,O_RDWR);if(fd<0)_exit(80);state=mmap(0,48,PROT_READ|PROT_WRITE,MAP_SHARED,fd,0);close(fd);if(state==MAP_FAILED)_exit(81);count=0;strlcpy(common,git,sizeof(common));char*copy=strdup(paths),*save=0;for(char*p=strtok_r(copy,"\n",&save);p;p=strtok_r(0,"\n",&save)){if(count>=32)_exit(82);strlcpy(objects[count++],p,sizeof(objects[0]));}free(copy);
}
// Native Git runs from an unsigned test-owned copy so macOS does not strip
// instrumentation from a protected system image. Its argv and Git exec path
// are retained. Absolute paths and compressed stdout bypassing the PATH shim
// therefore still reach the same actual object/ref observer.
__attribute__((constructor)) static void race_initialize(void){const char*policy=getenv("HELD_RACE_NATIVE_POLICY");if(!policy)return;strlcpy(native_policy,policy,sizeof(native_policy));Dl_info image;if(dladdr((void*)&race_initialize,&image)&&image.dli_fname)strlcpy(library_path,image.dli_fname,sizeof(library_path));const char*exec=getenv("GIT_EXEC_PATH");if(exec)strlcpy(git_exec_path,exec,sizeof(git_exec_path));FILE*f=fopen(policy,"r");if(!f)_exit(91);char file[4096],git[4096],paths[131072],line[4096];paths[0]=0;if(!fgets(file,sizeof(file),f)||!fgets(git,sizeof(git),f)||!fgets(original_git,sizeof(original_git),f)||!fgets(engine_git,sizeof(engine_git),f)||!fgets(traced_git,sizeof(traced_git),f))_exit(92);file[strcspn(file,"\n")]=0;git[strcspn(git,"\n")]=0;original_git[strcspn(original_git,"\n")]=0;engine_git[strcspn(engine_git,"\n")]=0;traced_git[strcspn(traced_git,"\n")]=0;while(fgets(line,sizeof(line),f))strlcat(paths,line,sizeof(paths));fclose(f);held_race_arm(file,git,paths);}
static const char*image(const char*p){char canonical[4096];if(original_git[0]&&p&&realpath(p,canonical)&&(!strcmp(canonical,original_git)||!strcmp(canonical,engine_git)))return traced_git;return p;}
static char**transport(char*const env[]){size_t n=0;if(env)while(env[n])n++;char**next=calloc(n+4,sizeof(char*));size_t j=0;int has_exec=0;for(size_t i=0;i<n;i++){if(!strncmp(env[i],"DYLD_INSERT_LIBRARIES=",22)||!strncmp(env[i],"HELD_RACE_NATIVE_POLICY=",24))continue;if(!strncmp(env[i],"GIT_EXEC_PATH=",14))has_exec=1;next[j++]=strdup(env[i]);}asprintf(next+j++,"DYLD_INSERT_LIBRARIES=%s",library_path);asprintf(next+j++,"HELD_RACE_NATIVE_POLICY=%s",native_policy);if(!has_exec&&git_exec_path[0])asprintf(next+j++,"GIT_EXEC_PATH=%s",git_exec_path);return next;}
static void free_transport(char**env){for(size_t i=0;env[i];i++)free(env[i]);free(env);}
static int race_execve(const char*p,char*const argv[],char*const env[]){const char*target=image(p);if(target==p)return execve(p,argv,env);char**next=transport(env);int result=execve(target,argv,next);int code=errno;free_transport(next);errno=code;return result;}
static int race_spawn(pid_t*pid,const char*p,const posix_spawn_file_actions_t*actions,const posix_spawnattr_t*attr,char*const argv[],char*const env[]){const char*target=image(p);if(target==p)return posix_spawn(pid,p,actions,attr,argv,env);char**next=transport(env);int result=posix_spawn(pid,target,actions,attr,argv,next);free_transport(next);return result;}
static int race_spawnp(pid_t*pid,const char*p,const posix_spawn_file_actions_t*actions,const posix_spawnattr_t*attr,char*const argv[],char*const env[]){const char*target=image(p);if(target==p)return posix_spawnp(pid,p,actions,attr,argv,env);char**next=transport(env);int result=posix_spawn(pid,target,actions,attr,argv,next);free_transport(next);return result;}
void held_race_disarm(void){if(state){munmap((void*)state,48);state=0;}}
static int race_open(const char*p,int flags,...){inventory(p);mode_t mode=0;if(flags&O_CREAT){va_list a;va_start(a,flags);mode=va_arg(a,int);va_end(a);}return open(p,flags,mode);}
static int race_openat(int fd,const char*p,int flags,...){char path[4096],dir[4096];const char*q=p;if(p&&p[0]!='/'){if(fd==AT_FDCWD){if(getcwd(dir,sizeof(dir)))snprintf(path,sizeof(path),"%s/%s",dir,p);else path[0]=0;}else if(fcntl(fd,F_GETPATH,dir)>=0)snprintf(path,sizeof(path),"%s/%s",dir,p);else path[0]=0;q=path;}inventory(q);mode_t mode=0;if(flags&O_CREAT){va_list a;va_start(a,flags);mode=va_arg(a,int);va_end(a);}return openat(fd,p,flags,mode);}
static ssize_t race_read(int fd,void*b,size_t n){descriptor_inventory(fd);ssize_t got=read(fd,b,n);if(got>0)descriptor(fd);return got;}
static ssize_t race_pread(int fd,void*b,size_t n,off_t offset){descriptor_inventory(fd);ssize_t got=pread(fd,b,n,offset);if(got>0)descriptor(fd);return got;}
static void*race_mmap(void*a,size_t n,int protection,int flags,int fd,off_t offset){descriptor_inventory(fd);void*p=mmap(a,n,protection,flags,fd,offset);if(p!=MAP_FAILED&&n&&fd>=0)descriptor(fd);return p;}
static DIR*race_opendir(const char*p){inventory(p);return opendir(p);}
static struct dirent*race_readdir(DIR*d){char path[4096];if(fcntl(dirfd(d),F_GETPATH,path)>=0)inventory(path);struct dirent*entry=readdir(d);if(entry)descriptor(dirfd(d));return entry;}
#define I(r,o) __attribute__((used)) static struct{const void*replacement;const void*original;} link_##o __attribute__((section("__DATA,__interpose")))={(const void*)(uintptr_t)&r,(const void*)(uintptr_t)&o};
I(race_open,open) I(race_openat,openat) I(race_read,read) I(race_pread,pread) I(race_mmap,mmap) I(race_opendir,opendir) I(race_readdir,readdir) I(race_execve,execve) I(race_spawn,posix_spawn) I(race_spawnp,posix_spawnp)
`

const heldRaceGitProgram = `package main
import("encoding/json";"io";"os";"os/exec";"path/filepath";"strings";"sync";"sync/atomic";"syscall";"time";"unsafe")
type policy struct{Real,State,Token string;Markers []string}
var state []byte;var p policy;var stream sync.Mutex
func word(i int)*uint64{return (*uint64)(unsafe.Pointer(&state[i*8]))}
func gate(){if atomic.LoadUint64(word(0))!=atomic.LoadUint64(word(4)){return};atomic.CompareAndSwapUint64(word(1),0,1);until:=time.Now().Add(10*time.Second);for atomic.LoadUint64(word(2))==0{if time.Now().After(until){os.Exit(94)};time.Sleep(time.Millisecond)}}
type input struct{destination io.Writer}
func (v input)Write(b []byte)(int,error){stream.Lock();defer stream.Unlock();gate();return v.destination.Write(b)}
type output struct{tail string}
func (v *output)Write(b []byte)(int,error){
 stream.Lock();defer stream.Unlock();n,e:=os.Stdout.Write(b);text:=v.tail+string(b[:n]);for i,marker:=range p.Markers{if strings.Contains(text,marker){for{old:=atomic.LoadUint64(word(0));if atomic.CompareAndSwapUint64(word(0),old,old|uint64(1)<<i){break}}}}
 if atomic.LoadUint64(word(1))!=0&&(strings.Contains(text,"refs/fireweed/claims/")||strings.Contains(text,p.Token)){atomic.AddUint64(word(3),1)}
 if len(text)>8192{text=text[len(text)-8192:]};v.tail=text;return n,e
}
func main(){self,e:=os.Executable();if e!=nil{os.Exit(80)};raw,e:=os.ReadFile(filepath.Join(filepath.Dir(self),"race-policy.json"));if e!=nil||json.Unmarshal(raw,&p)!=nil{os.Exit(81)};f,e:=os.OpenFile(p.State,os.O_RDWR,0);if e!=nil{os.Exit(82)};state,e=syscall.Mmap(int(f.Fd()),0,48,syscall.PROT_READ|syscall.PROT_WRITE,syscall.MAP_SHARED);f.Close();if e!=nil{os.Exit(83)};defer syscall.Munmap(state)
 gate();command:=exec.Command(p.Real,os.Args[1:]...);stdin,e:=command.StdinPipe();if e!=nil{os.Exit(84)};command.Stdout=&output{};command.Stderr=os.Stderr;if e=command.Start();e!=nil{os.Exit(85)};go func(){io.Copy(input{stdin},os.Stdin);stdin.Close()}();if e=command.Wait();e!=nil{if x,ok:=e.(*exec.ExitError);ok{os.Exit(x.ExitCode())};os.Exit(86)}
}
`

type heldRaceInput struct {
	Root, Common, State, Objects, Mode, ExpectedToken, RealGit, Base, NativePolicy string
	Paths                                                                          []string
	ObjectIDs                                                                      []string
}

func heldBuildRaceProbe(t *testing.T) (string, string) {
	t.Helper()
	directory := stateCanonicalDirectory(t, t.TempDir())
	worktreePut(t, directory, "guard.c", "#include <stdio.h>\n"+heldRaceGuard)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()
	library := filepath.Join(directory, "libheldrace.dylib")
	command := exec.CommandContext(ctx, "cc", "-dynamiclib", "-o", library, filepath.Join(directory, "guard.c"), "-Wl,-install_name,"+library)
	if out, err := command.CombinedOutput(); err != nil {
		t.Fatal("race native observer compile", err, string(out))
	}
	program := fmt.Sprintf(`package main
/*
#cgo LDFLAGS: -L%s -lheldrace
void held_race_arm(const char*,const char*,const char*);
void held_race_disarm(void);
#include <spawn.h>
#include <sys/wait.h>
#include <fcntl.h>
static int control_spawnp(const char*git,const char*base,const char*object){
 char*empty[]={0};char*source[]={(char*)git,"cat-file","blob",(char*)object,0};char*archive[]={(char*)git,"archive","--format=tar.gz",(char*)base,0};char*refs[]={(char*)git,"for-each-ref","--format=%%(refname:strip=3)","refs/fireweed/claims/",0};char**commands[]={source,archive,refs};
 for(int i=0;i<3;i++){posix_spawn_file_actions_t actions;posix_spawn_file_actions_init(&actions);posix_spawn_file_actions_addopen(&actions,1,"/dev/null",O_WRONLY,0);pid_t pid;int e=posix_spawnp(&pid,git,&actions,0,commands[i],empty);posix_spawn_file_actions_destroy(&actions);if(e)return e;int status;if(waitpid(pid,&status,0)<0)return 200;if(!WIFEXITED(status)||WEXITSTATUS(status))return 201;}
 return 0;
}
*/
import "C"
import("bufio";"encoding/json";"fmt";"io";"os";"os/exec";"path/filepath";"strings";p "github.com/lurkshark/fireweed/packages/project";w "github.com/lurkshark/fireweed/packages/workflow")
type input struct{Root,Common,State,Objects,Mode,ExpectedToken,RealGit,Base,NativePolicy string;Paths,ObjectIDs []string}
type output struct{Snapshot p.Snapshot;Claim w.Claim;ClaimObject string;Present bool;Error string;StateSame bool}
func runGit(git string,args...string)error{command:=exec.Command(git,args...);command.Stderr=os.Stderr;e:=command.Run();if e!=nil{fmt.Fprintln(os.Stderr,"native control",git,args,e)};return e}
func runGitEnv(git string,environment []string,args...string)error{command:=exec.Command(git,args...);command.Env=environment;command.Stderr=os.Stderr;e:=command.Run();if e!=nil{fmt.Fprintln(os.Stderr,"native minimal-env control",git,args,e)};return e}
func main(){raw,e:=os.ReadFile(os.Args[1]);var in input;if e!=nil||json.Unmarshal(raw,&in)!=nil{os.Exit(80)};var result output
 C.held_race_arm(C.CString(in.State),C.CString(in.Common),C.CString(in.Objects))
 switch in.Mode{
 case "direct-control":for _,path:=range in.Paths{if _,e=os.ReadFile(path);e!=nil{os.Exit(81)}};if _,e=os.ReadFile(in.Common+"/refs/fireweed/claims/current");e!=nil{os.Exit(82)}
 case "relative-control","relative-dot-parent-control":for _,path:=range in.Paths{relative,err:=filepath.Rel(in.Root,path);if err!=nil{os.Exit(81)};if _,e=os.ReadFile(relative);e!=nil{os.Exit(81)}};path:=".git/refs/fireweed/claims/current";if in.Mode=="relative-dot-parent-control"{path="./units/../.git/refs/fireweed/claims/../claims/current"};if _,e=os.ReadFile(path);e!=nil{os.Exit(82)}
 case "minimal-env-control","empty-env-control":environment:=[]string{};if in.Mode=="minimal-env-control"{environment=[]string{"PATH=/usr/bin:/bin","LANG=C","GIT_OPTIONAL_LOCKS=0"}};if e=runGitEnv(in.RealGit,environment,"cat-file","blob",in.ObjectIDs[0]);e!=nil{os.Exit(83)};if e=runGitEnv(in.RealGit,environment,"archive","--format=tar.gz",in.Base);e!=nil{os.Exit(84)};if e=runGitEnv(in.RealGit,environment,"for-each-ref","--format=%%(refname:strip=3)","refs/fireweed/claims/");e!=nil{os.Exit(85)}
 case "spawnp-empty-env-control":if code:=C.control_spawnp(C.CString(in.RealGit),C.CString(in.Base),C.CString(in.ObjectIDs[0]));code!=0{fmt.Fprintln(os.Stderr,"native spawnp control",code);os.Exit(93)}
 case "approval-wire":result.Snapshot,result.Claim,e=w.LoadApproval(in.Root,"current");result.Present=e==nil;if e!=nil{result.Error=e.Error()}
 case "object-only-wire":result.Claim.Object="held-object-only-wire-control";result.Error="restore source for wire control"
 case "git-control":for _,object:=range in.ObjectIDs{if e=runGit("git","cat-file","blob",object);e!=nil{os.Exit(83)}};if e=runGit("git","for-each-ref","--format=%%(refname) %%(objectname)","refs/fireweed/claims/");e!=nil{os.Exit(84)}
 case "absolute-compressed-control":if e=runGit(in.RealGit,"cat-file","blob",in.ObjectIDs[0]);e!=nil{os.Exit(83)};if e=runGit(in.RealGit,"archive","--format=tar.gz",in.Base);e!=nil{os.Exit(84)};if e=runGit(in.RealGit,"for-each-ref","--format=%%(refname:strip=3)","refs/fireweed/claims/");e!=nil{os.Exit(85)}
 case "batch-control":command:=exec.Command("git","cat-file","--batch");stdin,_:=command.StdinPipe();stdout,_:=command.StdoutPipe();command.Stderr=os.Stderr;if e=command.Start();e!=nil{os.Exit(85)};reader:=bufio.NewReader(stdout);for _,object:=range append(append([]string{},in.ObjectIDs...),in.ObjectIDs[0]){fmt.Fprintln(stdin,object);line,err:=reader.ReadString('\n');if err!=nil{os.Exit(86)};var oid,kind string;var size int;if _,err=fmt.Sscan(line,&oid,&kind,&size);err!=nil{os.Exit(87)};if _,err=io.CopyN(io.Discard,reader,int64(size)+1);err!=nil{os.Exit(88)}};stdin.Close();if e=command.Wait();e!=nil{os.Exit(89)}
 default:result.Snapshot,result.Claim,result.Present,e=w.LoadHeldApproval(in.Root);if e!=nil{result.Error=e.Error()}
 }
 if in.Mode!=""&&!strings.HasSuffix(in.Mode,"wire"){var token []byte;if in.Mode=="direct-control"||strings.HasPrefix(in.Mode,"relative-"){token,e=os.ReadFile(in.Common+"/refs/fireweed/claims/current")}else{git:="git";if in.Mode=="absolute-compressed-control"||strings.HasSuffix(in.Mode,"env-control"){git=in.RealGit};token,e=exec.Command(git,"rev-parse","refs/fireweed/claims/current").Output()};if e!=nil||strings.TrimSpace(string(token))!=in.ExpectedToken{result.Error=fmt.Sprint("live ownership control mismatch: ",string(token)," ",e)}}
 C.held_race_disarm();result.ClaimObject=result.Claim.Object;if json.NewEncoder(os.Stdout).Encode(result)!=nil{os.Exit(90)}
}
`, directory)
	worktreePut(t, directory, "probe.go", program)
	source, err := filepath.Abs("../../..")
	if err != nil {
		t.Fatal(err)
	}
	worktreePut(t, directory, "go.mod", fmt.Sprintf("module heldraceprobe\n\ngo 1.23\nrequire github.com/lurkshark/fireweed v0.0.0\nreplace github.com/lurkshark/fireweed => %q\n", source))
	binaryPath := filepath.Join(directory, "probe")
	command = exec.CommandContext(ctx, "go", "build", "-mod=mod", "-o", binaryPath, filepath.Join(directory, "probe.go"))
	command.Dir = directory
	command.Env = append(os.Environ(), "CGO_ENABLED=1", "GOPROXY=off", "GOTOOLCHAIN=local")
	if out, err := command.CombinedOutput(); err != nil {
		t.Fatal("race public probe compile", err, string(out))
	}
	worktreePut(t, directory, "git.go", heldRaceGitProgram)
	command = exec.CommandContext(ctx, "go", "build", "-o", filepath.Join(directory, "git"), filepath.Join(directory, "git.go"))
	command.Env = append(os.Environ(), "GOPROXY=off", "GOTOOLCHAIN=local")
	if out, err := command.CombinedOutput(); err != nil {
		t.Fatal("race streaming Git observer compile", err, string(out))
	}
	real := protocolRealGit(t)
	real, err = filepath.EvalSymlinks(real)
	if err != nil {
		t.Fatal(err)
	}
	engine := heldNativeGitEngine(t, real)
	data, err := os.ReadFile(engine)
	if err != nil {
		t.Fatal(err)
	}
	traced := filepath.Join(directory, "native-git")
	if err := os.WriteFile(traced, data, 0755); err != nil {
		t.Fatal(err)
	}
	command = exec.CommandContext(ctx, "codesign", "--remove-signature", traced)
	if out, err := command.CombinedOutput(); err != nil {
		t.Fatal("prepare owned native Git observer image", err, string(out))
	}
	command = exec.CommandContext(ctx, "codesign", "--force", "--sign", "-", traced)
	if out, err := command.CombinedOutput(); err != nil {
		t.Fatal("sign owned native Git observer image", err, string(out))
	}
	return directory, binaryPath
}

func heldNativeGitEngine(t *testing.T, real string) string {
	t.Helper()
	if real != "/usr/bin/git" {
		return real
	}
	// /usr/bin/git is Apple's arm64e developer-tool launcher. Instrument
	// the selected developer Git engine, retaining the launcher's actual
	// --exec-path, rather than copying the unrelated launcher image.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	command := exec.CommandContext(ctx, "xcrun", "--find", "git")
	bytes, err := command.Output()
	if err != nil {
		t.Fatal("selected native Git engine", err)
	}
	engine, err := filepath.EvalSymlinks(strings.TrimSpace(string(bytes)))
	if err != nil {
		t.Fatal(err)
	}
	return engine
}

func heldRaceFixture(t *testing.T, directory, real, format string) (approvalFixtureData, heldRaceInput) {
	t.Helper()
	f := heldSourceFixture(t, "impl", format, false)
	paths := []string{"fireweed.yaml", "units/a/SPEC.md", "units/b/SPEC.md", "units/c/SPEC.md", "units/d/SPEC.md", "units/x/SPEC.md", "changes/0001-before.md", "changes/0003-later.md", "changes/archive/0009-earlier.md"}
	markers := []string{}
	for index, path := range paths {
		marker := fmt.Sprintf("held-source-%s-%02d", filepath.Base(t.TempDir()), index)
		markers = append(markers, marker)
		raw, err := os.ReadFile(filepath.Join(f.Root, path))
		if err != nil {
			t.Fatal(err)
		}
		text := string(raw)
		if path == "fireweed.yaml" {
			text = strings.Replace(text, "worker_writable: [shared/**]", "worker_writable: [shared/**, "+marker+"/**]", 1)
		} else {
			text += "\n\n## Notes\n\n" + marker + "\n"
		}
		if !strings.Contains(text, marker) {
			t.Fatal("fresh normative source nonce premise", path)
		}
		worktreePut(t, f.Root, path, text)
	}
	f.Claim.State += "\n\n## Notes\n\nheld-started-state-nonce\n"
	f.Claim.Base = approvalCommit(t, f.Root)
	f.Claim.Object = claimRaw(t, f.Root, "current", claimPayload(t, f.Claim))
	markers = append([]string{"held-started-state-nonce"}, markers...)
	var err error
	f.Want, _, err = w.LoadApproval(f.Root, "current")
	if err != nil {
		t.Fatal("fresh exact approval control", err)
	}
	common := stateCanonicalDirectory(t, claimCommon(t, f.Root))
	objects := []string{f.Claim.Object}
	for _, path := range paths {
		objects = append(objects, claimGit(t, f.Root, "rev-parse", f.Claim.Base+":"+path))
	}
	loose := []string{}
	for _, object := range objects {
		path := filepath.Join(common, "objects", object[:2], object[2:])
		if _, err := os.Stat(path); err != nil {
			t.Fatal("fresh loose source object premise", err)
		}
		loose = append(loose, path)
	}
	statePath := filepath.Join(t.TempDir(), "state")
	state := make([]byte, 48)
	binary.LittleEndian.PutUint64(state[32:], uint64(1)<<len(markers)-1)
	if err := os.WriteFile(statePath, state, 0600); err != nil {
		t.Fatal(err)
	}
	policy, err := json.Marshal(struct {
		Real, State, Token string
		Markers            []string
	}{real, statePath, f.Claim.Object, markers})
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(directory, "race-policy.json"), policy, 0600); err != nil {
		t.Fatal(err)
	}
	real, err = filepath.EvalSymlinks(real)
	if err != nil {
		t.Fatal(err)
	}
	nativePolicy := filepath.Join(t.TempDir(), "native-policy")
	if err := os.WriteFile(nativePolicy, []byte(strings.Join(append([]string{statePath, common, real, heldNativeGitEngine(t, real), filepath.Join(directory, "native-git")}, loose...), "\n")+"\n"), 0600); err != nil {
		t.Fatal(err)
	}
	heldPoisonCurrent(t, f.Root)
	return f, heldRaceInput{Root: f.Root, Common: common, State: statePath, Objects: strings.Join(loose, "\n"), Paths: loose, ObjectIDs: objects, RealGit: real, Base: f.Claim.Base, NativePolicy: nativePolicy}
}

func heldRaceState(t *testing.T, path string) []uint64 {
	t.Helper()
	bytes, err := os.ReadFile(path)
	if err != nil || len(bytes) != 48 {
		t.Fatal("race observer state", err, len(bytes))
	}
	values := []uint64{}
	for index := 0; index < 6; index++ {
		values = append(values, binary.LittleEndian.Uint64(bytes[index*8:]))
	}
	return values
}

func heldRaceRun(t *testing.T, directory, executable string, in heldRaceInput, update func()) (heldNativeOutput, bool) {
	t.Helper()
	input, err := json.Marshal(in)
	if err != nil {
		t.Fatal(err)
	}
	inputPath := filepath.Join(t.TempDir(), "input.json")
	if err := os.WriteFile(inputPath, input, 0600); err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 25*time.Second)
	defer cancel()
	command := exec.CommandContext(ctx, executable, inputPath)
	command.Dir = in.Root
	execPath := protocolGit(t, in.RealGit, in.Root, "", "--exec-path")
	command.Env = append(os.Environ(), "PATH="+directory+string(os.PathListSeparator)+os.Getenv("PATH"), "DYLD_INSERT_LIBRARIES="+filepath.Join(directory, "libheldrace.dylib"), "HELD_RACE_NATIVE_POLICY="+in.NativePolicy, "GIT_EXEC_PATH="+execPath)
	var stdout, stderr bytes.Buffer
	command.Stdout, command.Stderr = &stdout, &stderr
	if err := command.Start(); err != nil {
		t.Fatal(err)
	}
	done := make(chan error, 1)
	go func() { done <- command.Wait() }()
	ticker := time.NewTicker(2 * time.Millisecond)
	defer ticker.Stop()
	fired := false
	for {
		select {
		case <-ticker.C:
			state := heldRaceState(t, in.State)
			if state[1] != 0 && !fired {
				fired = true
				if state[0] != state[4] {
					t.Fatal("ownership barrier preceded complete actual source delivery", state)
				}
				if update != nil {
					update()
				}
				file, err := os.OpenFile(in.State, os.O_WRONLY, 0)
				if err != nil {
					t.Fatal(err)
				}
				_, err = file.WriteAt([]byte{1, 0, 0, 0, 0, 0, 0, 0}, 16)
				file.Close()
				if err != nil {
					t.Fatal(err)
				}
			}
		case err := <-done:
			if err != nil || ctx.Err() != nil {
				t.Fatal("bounded race probe", err, ctx.Err(), stderr.String(), stdout.String())
			}
			var result heldNativeOutput
			if err := json.Unmarshal(stdout.Bytes(), &result); err != nil {
				t.Fatal("race public result", err, stdout.String(), stderr.String())
			}
			result.Claim.Object = result.ClaimObject
			return result, fired
		}
	}
}

func Test_FLOW049_FLOW050HeldOwnershipRevalidation(t *testing.T) {
	if runtime.GOOS != "darwin" {
		t.Skip("macOS native data/ref observer; portable ownership/source cases run separately")
	}
	directory, executable := heldBuildRaceProbe(t)
	real := protocolRealGit(t)
	for _, mode := range []string{"approval-wire", "object-only-wire", "direct-control", "relative-control", "relative-dot-parent-control", "git-control", "batch-control", "absolute-compressed-control", "minimal-env-control", "empty-env-control", "spawnp-empty-env-control"} {
		t.Run(mode, func(t *testing.T) {
			f, in := heldRaceFixture(t, directory, real, "sha1")
			in.Mode = mode
			if strings.HasSuffix(mode, "wire") {
				result, _ := heldRaceRun(t, directory, executable, in, nil)
				if mode == "approval-wire" {
					if result.Error != "" || !result.Present || result.Claim != f.Claim || !reflect.DeepEqual(result.Snapshot, f.Want) {
						t.Fatal("race nonzero exact-token wire control", result.Error, result.Claim)
					}
				} else if result.Claim != (w.Claim{Object: "held-object-only-wire-control"}) || result.ClaimObject == "" {
					t.Fatal("race wire concealed an Object-only partial claim", result)
				}
				return
			}
			replacement := f.Claim
			replacement.State += "\nExternal observer control replacement.\n"
			in.ExpectedToken = strings.TrimSpace(protocolGit(t, real, f.Root, claimPayload(t, replacement), "hash-object", "-w", "--stdin"))
			result, fired := heldRaceRun(t, directory, executable, in, func() {
				protocolGit(t, real, f.Root, "", "update-ref", "refs/fireweed/claims/current", in.ExpectedToken, f.Claim.Object)
				if strings.TrimSpace(protocolGit(t, real, f.Root, "", "rev-parse", "refs/fireweed/claims/current")) != in.ExpectedToken {
					t.Fatal("observer control external CAS failed")
				}
			})
			if result.Error != "" {
				t.Fatal("post-barrier actual live ownership control", result.Error)
			}
			if !fired || heldRaceState(t, in.State)[0] != heldRaceState(t, in.State)[4] {
				t.Fatal("actual source and subsequent ownership/request observer control absent", fired, heldRaceState(t, in.State))
			}
		})
	}
	for _, format := range []string{"sha1", "sha256"} {
		for _, mode := range []string{"stable", "delete", "replace-local", "replace-foreign", "add-second-local"} {
			t.Run(format+"/"+mode, func(t *testing.T) {
				f, in := heldRaceFixture(t, directory, real, format)
				newObject := ""
				ref := "refs/fireweed/claims/current"
				if mode != "stable" && mode != "delete" {
					replacement := f.Claim
					replacement.State += "\nReplacement state.\n"
					if mode == "replace-foreign" {
						replacement.Worktree = heldPeer(t, f.Root)
					}
					if mode == "add-second-local" {
						replacement.Change = "second"
						ref = "refs/fireweed/claims/second"
					}
					newObject = strings.TrimSpace(protocolGit(t, real, f.Root, claimPayload(t, replacement), "hash-object", "-w", "--stdin"))
				}
				update := func() {
					switch mode {
					case "delete":
						protocolGit(t, real, f.Root, "", "update-ref", "-d", ref, f.Claim.Object)
					case "replace-local", "replace-foreign":
						protocolGit(t, real, f.Root, "", "update-ref", ref, newObject, f.Claim.Object)
					case "add-second-local":
						protocolGit(t, real, f.Root, "", "update-ref", ref, newObject, strings.Repeat("0", len(f.Claim.Object)))
					}
					if mode == "delete" {
						_, _, err := protocolCommand(t, real, f.Root, "", "show-ref", "--verify", ref)
						if err == nil {
							t.Fatal("external exact-token deletion did not occur")
						}
					} else if mode != "stable" && strings.TrimSpace(protocolGit(t, real, f.Root, "", "rev-parse", ref)) != newObject {
						t.Fatal("external replacement/inventory control")
					}
				}
				result, fired := heldRaceRun(t, directory, executable, in, update)
				if !fired {
					// An inert stub has not established a race. Its ordinary cold
					// source failure remains genuine, not an injection failure.
					if result.Error != "" || !result.Present || result.Claim != f.Claim || !reflect.DeepEqual(result.Snapshot, f.Want) {
						t.Error("available cold approval was not selected; no ownership race was exercised", result.Error)
					} else {
						state := heldRaceState(t, in.State)
						if state[0] != state[4] {
							t.Fatal("eval observer coverage is insufficient for this successful source strategy; no ownership-clause failure is established", state)
						} else {
							t.Error("complete observed fresh approval returned without subsequent live ownership validation", state)
						}
					}
					return
				}
				if mode == "stable" {
					if heldRaceState(t, in.State)[3] == 0 {
						t.Error("stable cold source lacked actual subsequent live ownership data", heldRaceState(t, in.State))
					}
					if result.Error != "" || !result.Present || result.Claim != f.Claim || !reflect.DeepEqual(result.Snapshot, f.Want) {
						t.Error("stable source after actual final observation", result.Error, result.Claim)
					}
				} else if result.Error == "" || !policyCorrective(result.Error) || result.Present || result.Claim != (w.Claim{}) || !reflect.DeepEqual(result.Snapshot, p.Snapshot{}) {
					t.Error("observed ownership change returned authority or lacked corrective rejection", result.Error, result.Present, result.Claim)
				}
				if mode == "delete" {
					if _, _, err := protocolCommand(t, real, f.Root, "", "show-ref", "--verify", ref); err == nil {
						t.Error("reader recreated externally deleted ownership")
					}
				} else if mode != "stable" && strings.TrimSpace(protocolGit(t, real, f.Root, "", "rev-parse", ref)) != newObject {
					t.Error("reader overwrote external ownership update")
				}
			})
		}
	}
}
