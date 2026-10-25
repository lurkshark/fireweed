package evals

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"testing"
	"time"

	a "github.com/lurkshark/fireweed/packages/assets"
	p "github.com/lurkshark/fireweed/packages/project"
	w "github.com/lurkshark/fireweed/packages/workflow"
)

const startFaultGuard = `#include <sys/types.h>
#include <sys/mman.h>
#include <sys/proc_info.h>
#include <libproc.h>
#include <crt_externs.h>
#include <mach-o/dyld.h>
#include <mach-o/loader.h>
#include <mach-o/nlist.h>
#include <arpa/inet.h>
#include <sys/socket.h>
#include <sys/stat.h>
#include <sys/wait.h>
#include <spawn.h>
#include <unistd.h>
#include <fcntl.h>
#include <stdarg.h>
#include <stdint.h>
#include <string.h>
#include <stdio.h>
#include <errno.h>
#include <stdlib.h>
#include <signal.h>
#include <dirent.h>
static int armed=0,fired=0,deletions=0;
static pid_t owner=0;
static void stopped(int signal_number){_exit(42);}
static void violation(void){if(owner&&getpid()!=owner)kill(owner,SIGUSR2);_exit(42);}
static const char *root=0,*change=0,*context_file=0,*private_file=0;
static char failed_delete[4096],recreated[4096];
// The guard follows native supervision without permitting unguarded Go self
// execution. Policy crosses exec through test-owned POSIX shared memory, never
// an argument convention belonging to the product implementation.
#define START_POLICY_MAGIC UINT64_C(0x6677737461727431)
struct start_policy { uint64_t magic;pid_t owner;int mode;uint64_t sec,usec;char source[4096],root[4096],change[4096],context[4096],private_path[4096],address[4096]; };
static struct start_policy *owned_policy=0,*inherited_policy=0;
static char source_executable[4096],policy_name[64];
static int supervisor=0;
static const char*network_address=0;
static void native_calibration_entry(int argc,char**argv);
static void install_go_entry_guard(void);
static void policy_path(char*out,size_t n,pid_t pid){snprintf(out,n,"/fwstartguard-%d",pid);}
static int native_identity(pid_t pid,struct proc_bsdinfo*info){return proc_pidinfo(pid,PROC_PIDTBSDINFO,0,info,sizeof(*info))==(int)sizeof(*info);}
static void source_path(void){if(!source_executable[0]&&proc_pidpath(getpid(),source_executable,sizeof(source_executable))<=0)_exit(85);}
#define START_POLICY_ENV "__FIREWEED_TEST_START_POLICY_V1"
static void inherit_policy(void){
 source_path();const char*raw=getenv(START_POLICY_ENV);
 if(!raw){char parent_path[4096];pid_t parent=getppid();if(proc_pidpath(parent,parent_path,sizeof(parent_path))>0&&!strcmp(parent_path,source_executable)){owner=parent;violation();}return;}
 int parent=0,used=0;unsigned long long sec=0,usec=0;
 if(sscanf(raw,"%d:%llx:%llx%n",&parent,&sec,&usec,&used)!=3||raw[used]||parent<=1){_exit(42);}
 // This is observer-only transport. Remove it before any native supervisor
 // logic or Go initialization sees its original environment.
 if(unsetenv(START_POLICY_ENV)){owner=parent;violation();}
 owner=parent;char name[64];policy_path(name,sizeof(name),parent);int fd=shm_open(name,O_RDONLY,0);if(fd<0)violation();
 struct stat st;int good=fstat(fd,&st)==0&&st.st_size>=(off_t)sizeof(struct start_policy);struct start_policy*policy=good?mmap(0,sizeof(*policy),PROT_READ,MAP_SHARED,fd,0):MAP_FAILED;close(fd);
 struct proc_bsdinfo info;
 if(policy==MAP_FAILED||policy->magic!=START_POLICY_MAGIC||policy->owner!=parent||policy->mode<1||policy->mode>8||strcmp(policy->source,source_executable)||policy->sec!=sec||policy->usec!=usec||!native_identity(parent,&info)||info.pbi_start_tvsec!=policy->sec||info.pbi_start_tvusec!=policy->usec)violation();
 inherited_policy=policy;armed=policy->mode;supervisor=1;root=policy->root;change=policy->change;context_file=policy->context;private_file=policy->private_path;network_address=policy->address;signal(SIGUSR2,stopped);
}
// The parent environment and supplied strings are never mutated. Only a
// source-executable child receives the oracle token; its constructor removes
// that token before product code runs. Git/other targets retain their argv/env.
static char**supervisor_environment(const char*path,char*const env[],char marker[160]){
 if(!armed||strcmp(path,source_executable))return 0;
 const struct start_policy*policy=owned_policy?owned_policy:inherited_policy;if(!policy)violation();
 size_t n=0;while(env&&env[n]){if(!strncmp(env[n],START_POLICY_ENV "=",sizeof(START_POLICY_ENV)))violation();if(++n>(1U<<20))violation();}
 char**copy=calloc(n+2,sizeof(char*));if(!copy)violation();for(size_t i=0;i<n;i++)copy[i]=env[i];
 snprintf(marker,160,START_POLICY_ENV "=%d:%llx:%llx",policy->owner,(unsigned long long)policy->sec,(unsigned long long)policy->usec);copy[n]=marker;return copy;
}
__attribute__((constructor)) static void native_supervisor_entry(void){
 inherit_policy();install_go_entry_guard();int argc=*_NSGetArgc();char**argv=*_NSGetArgv();
 native_calibration_entry(argc,argv);
}
void fw_start_go_entered(void){if(supervisor)violation();}
// Go calls this native notification after runtime initialization and BEFORE
// initializing imported/application packages (runtime/proc.go). Locate only
// this owned probe's writable slot in its loaded Mach-O symbol table; no file
// or product implementation is consulted. Absence fails the guard closed.
static void (*original_go_notify)(void*);
static void guarded_go_notify(void*argument){fw_start_go_entered();original_go_notify(argument);}
static void **go_notify_slot(void){
 const struct mach_header_64*h=(const struct mach_header_64*)_dyld_get_image_header(0);if(!h||h->magic!=MH_MAGIC_64)return 0;
 intptr_t slide=_dyld_get_image_vmaddr_slide(0);const struct load_command*command=(const struct load_command*)(h+1);const struct symtab_command*symbols=0;uintptr_t linkedit=0;
 for(uint32_t i=0;i<h->ncmds;i++){if(command->cmd==LC_SEGMENT_64){const struct segment_command_64*segment=(const struct segment_command_64*)command;if(!strcmp(segment->segname,"__LINKEDIT"))linkedit=(uintptr_t)(slide+segment->vmaddr-segment->fileoff);}if(command->cmd==LC_SYMTAB)symbols=(const struct symtab_command*)command;command=(const struct load_command*)((const char*)command+command->cmdsize);}
 if(!symbols||!linkedit||symbols->nsyms>(1U<<22))return 0;const struct nlist_64*list=(const struct nlist_64*)(linkedit+symbols->symoff);const char*names=(const char*)(linkedit+symbols->stroff);
 for(uint32_t i=0;i<symbols->nsyms;i++)if(list[i].n_un.n_strx<symbols->strsize&&!strcmp(names+list[i].n_un.n_strx,"__cgo_notify_runtime_init_done"))return(void**)(slide+list[i].n_value);return 0;
}
static void install_go_entry_guard(void){void**slot=go_notify_slot();if(!slot)_exit(88);original_go_notify=(void(*)(void*))*slot;if(!original_go_notify)_exit(88);*slot=(void*)guarded_go_notify;}
int fw_start_cleanup(int pid){
 char name[64];policy_path(name,sizeof(name),pid);int fd=shm_open(name,O_RDONLY,0);if(fd<0)return errno==ENOENT?0:1;
 struct stat st;int good=fstat(fd,&st)==0&&st.st_size>=(off_t)sizeof(struct start_policy);struct start_policy*policy=good?mmap(0,sizeof(*policy),PROT_READ,MAP_SHARED,fd,0):MAP_FAILED;close(fd);
 source_path();if(policy==MAP_FAILED)return 2;int owned=policy->magic==START_POLICY_MAGIC&&policy->owner==pid&&!strcmp(policy->source,source_executable);munmap(policy,sizeof(*policy));if(!owned)return 3;return shm_unlink(name)==0||errno==ENOENT?0:4;
}
void fw_start_arm(int m,const char*r,const char*c,const char*x,const char*p,const char*address){
 armed=0;
 if(!m){if(owned_policy){munmap(owned_policy,sizeof(*owned_policy));owned_policy=0;shm_unlink(policy_name);}root=r;change=c;context_file=x;private_file=p;network_address=address;return;}
 source_path();owner=getpid();policy_path(policy_name,sizeof(policy_name),owner);
 int fd=shm_open(policy_name,O_RDWR|O_CREAT|O_EXCL,0600);if(fd<0)_exit(86);if(ftruncate(fd,sizeof(struct start_policy))){close(fd);shm_unlink(policy_name);_exit(86);}
 owned_policy=mmap(0,sizeof(*owned_policy),PROT_READ|PROT_WRITE,MAP_SHARED,fd,0);close(fd);if(owned_policy==MAP_FAILED){owned_policy=0;shm_unlink(policy_name);_exit(86);}
 struct proc_bsdinfo info;if(!native_identity(owner,&info))_exit(86);
 owned_policy->magic=START_POLICY_MAGIC;owned_policy->owner=owner;owned_policy->mode=m;owned_policy->sec=info.pbi_start_tvsec;owned_policy->usec=info.pbi_start_tvusec;
 const char*values[6]={source_executable,r,c,x,p,address};char*destinations[6]={owned_policy->source,owned_policy->root,owned_policy->change,owned_policy->context,owned_policy->private_path,owned_policy->address};
 for(int i=0;i<6;i++){const char*value=values[i]?values[i]:"";if(strlen(value)>=4096)_exit(86);snprintf(destinations[i],4096,"%s",value);}
 root=r;change=c;context_file=x;private_file=p;network_address=address;armed=m;signal(SIGUSR2,stopped);
}
int fw_start_fired(void){return fired;}
const char*fw_start_recreated(void){return recreated;}
static int exact(const char*a,const char*b){return a&&b&&strcmp(a,b)==0;}
static int within(const char*p,const char*d){return p&&d&&strncmp(p,d,strlen(d))==0&&(p[strlen(d)]==0||p[strlen(d)]=='/');}
static const char*fullpath(int fd,const char*p,char*b){
 if(!p)return p;
 char raw[4096],d[4096];
 if(p[0]=='/'){if(snprintf(raw,sizeof(raw),"%s",p)>=(int)sizeof(raw))return p;}
 else{int ok=fd==AT_FDCWD?getcwd(d,sizeof(d))!=0:fcntl(fd,F_GETPATH,d)==0;if(!ok)return p;if(snprintf(raw,sizeof(raw),"%s/%s",d,p)>=(int)sizeof(raw))return p;}
 size_t used=1;b[0]='/';b[1]=0;
 for(char*q=raw;*q;){while(*q=='/')q++;char*part=q;while(*q&&*q!='/')q++;size_t n=q-part;if(!n)break;
  if(n==1&&part[0]=='.')continue;
  if(n==2&&part[0]=='.'&&part[1]=='.'){if(used>1){while(used>1&&b[used-1]!='/')used--;if(used>1)used--;b[used]=0;}continue;}
  if(used>1)b[used++]='/';memcpy(b+used,part,n);used+=n;b[used]=0;
 }
 return b;
}
static void inspection(const char*p){if(armed==6)violation();}
static void mutation(const char*p){char path[4096];p=fullpath(AT_FDCWD,p,path);inspection(p);
 if((armed==7||armed==8)&&p){
  if(p[0]!='/')violation();
  if(within(p,root)){const char*q=p+strlen(root);if(strncmp(q,"/.git/",6)!=0&&strcmp(q,"/.git")!=0)violation();}
  if(within(p,context_file))violation();
 }
 if(armed&&within(p,private_file))violation();
}
static int opening(const char*p,int flags){
 if(!armed)return 0;
 char path[4096];p=fullpath(AT_FDCWD,p,path);inspection(p);
 if(flags&(O_WRONLY|O_RDWR)){mutation(p);if(armed==2&&exact(p,change)){fired=1;errno=EACCES;return -1;}}
 if((armed==5&&exact(p,context_file))||(armed==8&&exact(p,change))){fired=1;errno=EACCES;return -1;}
 return 0;
}
static int fw_open(const char*p,int f,...){if(opening(p,f)<0)return -1;mode_t m=0;if(f&O_CREAT){va_list a;va_start(a,f);m=va_arg(a,int);va_end(a);}return open(p,f,m);}
static int fw_openat(int d,const char*p,int f,...){char b[4096];const char*q=fullpath(d,p,b);if(opening(q,f)<0)return -1;mode_t m=0;if(f&O_CREAT){va_list a;va_start(a,f);m=va_arg(a,int);va_end(a);}return openat(d,p,f,m);}
static ssize_t fw_read(int d,void*b,size_t n){if(armed&&d==0)violation();return read(d,b,n);}
static ssize_t fw_write(int d,const void*b,size_t n){if(armed&&!supervisor&&(d==1||d==2))violation();return write(d,b,n);}
static int fw_socket(int d,int t,int p){if(armed)violation();return socket(d,t,p);}
static int fw_connect(int s,const struct sockaddr*a,socklen_t n){if(armed)violation();return connect(s,a,n);}
static void execution(const char*p,int replacing){if(armed){if(armed==6)violation();if(p&&!strcmp(p,source_executable)){if(replacing&&getpid()==owner)violation();return;}const char*b=strrchr(p,'/');b=b?b+1:p;if(strcmp(b,"git")!=0)violation();}}
static int spawn_replaces(const posix_spawnattr_t*t){short flags=0;if(t&&posix_spawnattr_getflags(t,&flags)){if(armed)violation();return 0;}return (flags&POSIX_SPAWN_SETEXEC)!=0;}
static int fw_execve(const char*p,char*const a[],char*const e[]){execution(p,1);char marker[160];char**env=supervisor_environment(p,e,marker);int result=execve(p,a,env?env:e);int saved=errno;free(env);errno=saved;return result;}
static int fw_posix_spawn(pid_t*p,const char*s,const posix_spawn_file_actions_t*f,const posix_spawnattr_t*t,char*const a[],char*const e[]){execution(s,spawn_replaces(t));char marker[160];char**env=supervisor_environment(s,e,marker);int result=posix_spawn(p,s,f,t,a,env?env:e);free(env);return result;}
static int fw_posix_spawnp(pid_t*p,const char*s,const posix_spawn_file_actions_t*f,const posix_spawnattr_t*t,char*const a[],char*const e[]){execution(s,spawn_replaces(t));char marker[160];char**env=supervisor_environment(s,e,marker);int result=posix_spawnp(p,s,f,t,a,env?env:e);free(env);return result;}
int fw_start_replace_control(void){posix_spawnattr_t attr;if(posix_spawnattr_init(&attr))return 1;posix_spawnattr_setflags(&attr,POSIX_SPAWN_SETEXEC);char*args[]={source_executable,"__fw_start_native_control__",0};extern char**environ;pid_t child;return fw_posix_spawn(&child,source_executable,0,&attr,args,environ);}
static int fw_rename(const char*p,const char*q){char path[4096];q=fullpath(AT_FDCWD,q,path);mutation(p);mutation(q);if(armed==2&&exact(q,change)){fired=1;errno=EACCES;return -1;}return rename(p,q);}
static int disposable(const char*p){return root&&p&&within(p,root)&&strstr(p,"/units/a/impl/")&&(strstr(p,"/code.txt")||strstr(p,"/second.txt"));}
static int deletion(const char*p){
 mutation(p);
 if((armed==3||armed==4)&&disposable(p)){
  if(failed_delete[0]&&exact(p,failed_delete)){errno=EACCES;return -1;}
  if(deletions){snprintf(failed_delete,sizeof(failed_delete),"%s",p);fired=1;errno=EACCES;return -1;}
 }
 return 0;
}
static void deleted(const char*p,int result){
 if(result==0&&(armed==3||armed==4)&&disposable(p)){
  deletions++;
  if(armed==4){int f=open(p,O_WRONLY|O_CREAT|O_EXCL,0640);if(f<0)_exit(86);const char*v="Concurrent winner bytes.\n";if(write(f,v,strlen(v))!=(ssize_t)strlen(v))_exit(87);fchmod(f,0640);close(f);snprintf(recreated,sizeof(recreated),"%s",p);}
 }
}
static int fw_unlink(const char*p){char path[4096];p=fullpath(AT_FDCWD,p,path);if(deletion(p)<0)return -1;int r=unlink(p);deleted(p,r);return r;}
static int fw_unlinkat(int d,const char*p,int f){char b[4096];const char*q=fullpath(d,p,b);if(deletion(q)<0)return -1;int r=unlinkat(d,p,f);deleted(q,r);return r;}
static char*fw_getcwd(char*b,size_t n){if(armed==6)violation();return getcwd(b,n);}
static int fw_stat(const char*p,struct stat*b){inspection(p);return stat(p,b);}
static int fw_lstat(const char*p,struct stat*b){inspection(p);return lstat(p,b);}
static DIR*fw_opendir(const char*p){inspection(p);return opendir(p);}
static ssize_t fw_readlink(const char*p,char*b,size_t n){inspection(p);return readlink(p,b,n);}
static int fw_access(const char*p,int f){inspection(p);return access(p,f);}
static int fw_rmdir(const char*p){mutation(p);return rmdir(p);}
static int fw_mkdir(const char*p,mode_t m){mutation(p);return mkdir(p,m);}
static int fw_chmod(const char*p,mode_t m){mutation(p);return chmod(p,m);}
static void native_calibration_entry(int argc,char**argv){
 if((argc!=2&&argc!=3)||strncmp(argv[1],"__fw_start_native_",18))return;
 if(!supervisor)_exit(89);
 if(!strcmp(argv[1],"__fw_start_native_control__"))_exit(0);
 if(!strcmp(argv[1],"__fw_start_native_environment__")){extern char**environ;int count=0;for(char**p=environ;*p;p++)count++;if(count!=2||strcmp(environ[0],"EXACT=a=b")||strcmp(environ[1],"EMPTY=")||getenv(START_POLICY_ENV))_exit(92);_exit(0);}
 if(!strcmp(argv[1],"__fw_start_native_chain__")){
  int depth=argc==3?atoi(argv[2]):-1;if(depth<0||depth>40)_exit(93);char next[16];snprintf(next,sizeof(next),"%d",depth-1);char*args[]={source_executable,depth?"__fw_start_native_chain__":"__fw_start_imported_go_control__",depth?next:0,0};pid_t child;extern char**environ;int launched=fw_posix_spawn(&child,source_executable,0,0,args,environ);if(launched)_exit(94);int status;while(waitpid(child,&status,0)<0&&errno==EINTR){};_exit(70);
 }
 if(!strcmp(argv[1],"__fw_start_native_output__")){const char*out="native-forward-out\n",*err="native-forward-err\n";fw_write(1,out,strlen(out));fw_write(2,err,strlen(err));_exit(0);}
 if(!strcmp(argv[1],"__fw_start_native_stdin__")){char byte;fw_read(0,&byte,1);}
 if(!strcmp(argv[1],"__fw_start_native_socket__"))fw_socket(AF_INET,SOCK_STREAM,0);
 if(!strcmp(argv[1],"__fw_start_native_connect__")){
  int port=0;if(!network_address||sscanf(network_address,"127.0.0.1:%d",&port)!=1||port<1||port>65535)_exit(91);
  // Call the underlying socket only to isolate a valid connect-denial control.
  int fd=socket(AF_INET,SOCK_STREAM,0);if(fd<0)_exit(91);struct sockaddr_in address={0};address.sin_len=sizeof(address);address.sin_family=AF_INET;address.sin_port=htons((uint16_t)port);address.sin_addr.s_addr=htonl(INADDR_LOOPBACK);fw_connect(fd,(struct sockaddr*)&address,sizeof(address));close(fd);
 }
 if(!strcmp(argv[1],"__fw_start_native_write__"))fw_open(private_file,O_WRONLY|O_TRUNC);
 if(!strcmp(argv[1],"__fw_start_native_chmod__"))fw_chmod(private_file,0777);
 if(!strcmp(argv[1],"__fw_start_native_unlink__"))fw_unlink(private_file);
 if(!strcmp(argv[1],"__fw_start_native_process__")){pid_t pid;char*args[]={"/usr/bin/true",0};extern char**environ;fw_posix_spawn(&pid,args[0],0,0,args,environ);}
 _exit(70);
}
#define I(r,o) __attribute__((used)) static struct{const void*replacement;const void*original;} fw_interpose_##o __attribute__((section("__DATA,__interpose")))={(const void*)(uintptr_t)&r,(const void*)(uintptr_t)&o};
I(fw_getcwd,getcwd) I(fw_stat,stat) I(fw_lstat,lstat) I(fw_opendir,opendir) I(fw_readlink,readlink) I(fw_access,access)
I(fw_open,open) I(fw_openat,openat) I(fw_read,read) I(fw_write,write) I(fw_socket,socket) I(fw_connect,connect)
I(fw_execve,execve) I(fw_posix_spawn,posix_spawn) I(fw_posix_spawnp,posix_spawnp) I(fw_rename,rename)
I(fw_unlink,unlink) I(fw_unlinkat,unlinkat) I(fw_rmdir,rmdir) I(fw_mkdir,mkdir) I(fw_chmod,chmod)
`

type startFaultInput struct {
	Root, Change, Context, Private, Address, Mode string
	Fault                                         int
}
type startFaultOutput struct {
	Result    w.StartResult
	Object    string
	Error     string
	Fired     bool
	Recreated string
	Control   bool
	StateSame bool
}

func Test_FLOW032_FLOW037_FLOW038_FLOW039_FLOW040StartNativeFailures(t *testing.T) {
	if runtime.GOOS != "darwin" {
		t.Skip("macOS armed native fault/interaction controls; portable start assertions run separately")
	}
	build, err := filepath.EvalSymlinks(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	worktreePut(t, build, "guard.c", startFaultGuard)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()
	lib := filepath.Join(build, "libstartfault.dylib")
	cmd := exec.CommandContext(ctx, "cc", "-dynamiclib", "-o", lib, filepath.Join(build, "guard.c"), "-Wl,-install_name,"+lib)
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("native fault guard compile: %v %s", err, out)
	}
	program := fmt.Sprintf(`package main
/*
#cgo LDFLAGS: -L%s -lstartfault
void fw_start_arm(int m,const char*r,const char*c,const char*x,const char*p,const char*address);
void fw_start_go_entered(void);
int fw_start_cleanup(int pid);
int fw_start_replace_control(void);
int fw_start_fired(void);
const char*fw_start_recreated(void);
*/
import "C"
import("encoding/json";"os";"os/exec";"net";"time";"fmt";"reflect";"strings";"strconv";"syscall";_ "fireweedstartprobe/early";w "github.com/lurkshark/fireweed/packages/workflow")
type input struct{Root,Change,Context,Private,Address,Mode string;Fault int}
type output struct{Result w.StartResult;Object string;Error string;Fired bool;Recreated string;Control,StateSame bool}
func init(){C.fw_start_go_entered()}
func main(){
 if len(os.Args)==3&&os.Args[1]=="--cleanup-start-guard"{pid,e:=strconv.Atoi(os.Args[2]);if e!=nil{os.Exit(90)};os.Exit(int(C.fw_start_cleanup(C.int(pid))))}
 raw,e:=os.ReadFile(os.Args[1]);if e!=nil{os.Exit(80)};var in input;if json.Unmarshal(raw,&in)!=nil{os.Exit(81)}
 cwd,e:=os.Getwd();if e!=nil{os.Exit(82)};environment:=os.Environ()
 C.fw_start_arm(C.int(in.Fault),C.CString(in.Root),C.CString(in.Change),C.CString(in.Context),C.CString(in.Private),C.CString(in.Address))
 out:=output{}
 relative:=func(p string)string{return "./"+strings.TrimPrefix(p,in.Root+"/")}
 switch in.Mode{
 case "token-transport-control":out.Result.Claim.Object="native-token-transport"
 case "invalid-relative-read-control":os.ReadFile("fireweed.yaml")
 case "invalid-relative-stat-control":os.Stat("fireweed.yaml")
 case "precondition-read-control":_,e=os.ReadFile(in.Change);out.Control=e!=nil
 case "precondition-write-control":os.WriteFile(in.Change,[]byte("transient bad mutation"),0600)
 case "precondition-claim-control":os.WriteFile(in.Context+"/current",[]byte("transient bad claim"),0600)
 case "invalid-read-control":os.ReadFile(in.Change)
 case "invalid-stat-control":os.Stat(in.Change)
 case "invalid-git-control":exec.Command("git","--version").Run()
 case "invalid-subject":out.Result,e=w.Start(in.Root,"INVALID/slug");if e!=nil{out.Error=e.Error()}
 case "stdin-control":os.Stdin.Read(make([]byte,1))
 case "prompt-control":fmt.Fprint(os.Stdout,"prompt")
 case "stderr-control":fmt.Fprint(os.Stderr,"prompt")
 case "network-control":net.DialTimeout("tcp",in.Address,time.Second)
 case "process-control":exec.Command("/usr/bin/true").Run()
 case "native-supervision-control","invalid-native-supervision-control":binary,_:=os.Executable();e=exec.Command(binary,"__fw_start_native_control__").Run();out.Control=e==nil
 case "native-environment-control":binary,_:=os.Executable();command:=exec.Command(binary,"__fw_start_native_environment__");command.Env=[]string{"EXACT=a=b","EMPTY="};e=command.Run();out.Control=e==nil
 case "native-depth-control":binary,_:=os.Executable();exec.Command(binary,"__fw_start_native_chain__","33").Run()
 case "native-output-control":binary,_:=os.Executable();b,e:=exec.Command(binary,"__fw_start_native_output__").CombinedOutput();out.Control=e==nil&&string(b)=="native-forward-out\nnative-forward-err\n"
 case "go-imported-init-control":binary,_:=os.Executable();exec.Command(binary,"__fw_start_imported_go_control__").Run()
 case "native-stdin-control","native-socket-control","native-connect-control","native-write-control","native-chmod-control","native-unlink-control","native-process-control":binary,_:=os.Executable();name:=strings.TrimSuffix(strings.TrimPrefix(in.Mode,"native-"),"-control");exec.Command(binary,"__fw_start_native_"+name+"__").Run()
 case "go-self-control":binary,_:=os.Executable();exec.Command(binary,"__fw_start_go_control__").Run()
 case "self-exec-control":binary,_:=os.Executable();syscall.Exec(binary,[]string{binary,"__fw_start_native_control__"},os.Environ())
 case "self-spawn-replace-control":C.fw_start_replace_control()
 case "native-copy-control":exec.Command("probe","__fw_start_native_control__").Run()
 case "hook-control":exec.Command("start-hook-trap").Run()
 case "agent-control":exec.Command("start-agent-trap").Run()
 case "private-write-control":os.WriteFile(in.Private,[]byte("bad"),0644)
 case "private-chmod-control":os.Chmod(in.Private,0777)
 case "private-unlink-control":os.Remove(in.Private)
 case "git-control":_,e=exec.Command("git","--version").CombinedOutput();out.Control=e==nil
 case "relative-write-fault-control":e=os.WriteFile(relative(in.Change),[]byte("bad"),0600);out.Control=e!=nil
 case "relative-rename-fault-control":
  tmp:=in.Root+"/rename-control";e=os.WriteFile(tmp,[]byte("temporary source"),0600);if e==nil{e=os.Rename(relative(tmp),relative(in.Change));out.Control=e!=nil;os.Remove(tmp)}
 case "relative-context-fault-control":_,e=os.ReadFile(relative(in.Context));out.Control=e!=nil
 case "fd-context-fault-control":
  root,e:=os.OpenRoot(in.Root);if e==nil{f,e:=root.Open(strings.TrimPrefix(in.Context,in.Root+"/"));out.Control=e!=nil;if f!=nil{f.Close()};root.Close()}
 case "relative-delete-fault-control":
  a:=os.Remove("./units/a/impl/code.txt");b:=os.Remove("./units/a/impl/second.txt");out.Control=a==nil&&b!=nil
 case "dotdot-context-fault-control":_,e=os.ReadFile("./units/a/../../"+strings.TrimPrefix(in.Context,in.Root+"/"));out.Control=e!=nil
 case "write-fault-control":e=os.WriteFile(in.Change,[]byte("bad"),0600);out.Control=e!=nil
 case "context-fault-control":_,e=os.ReadFile(in.Context);out.Control=e!=nil
 case "delete-fault-control":
  a:=os.Remove(in.Root+"/units/a/impl/code.txt");b:=os.Remove(in.Root+"/units/a/impl/second.txt");out.Control=a==nil&&b!=nil
 default:
  slug:="current";if in.Mode=="precondition-missing"{slug="missing"};out.Result,e=w.Start(in.Root,slug);if e!=nil{out.Error=e.Error()}
 }
 out.Object=out.Result.Claim.Object;out.Fired=C.fw_start_fired()!=0;out.Recreated=C.GoString(C.fw_start_recreated());C.fw_start_arm(0,nil,nil,nil,nil,nil)
 afterCwd,e:=os.Getwd();out.StateSame=e==nil&&cwd==afterCwd&&reflect.DeepEqual(environment,os.Environ())
 switch in.Mode{case "invalid-relative-read-control","invalid-relative-stat-control","precondition-write-control","precondition-claim-control","invalid-read-control","invalid-stat-control","invalid-git-control","stdin-control","prompt-control","stderr-control","network-control","process-control","native-depth-control","go-imported-init-control","native-stdin-control","native-socket-control","native-connect-control","native-write-control","native-chmod-control","native-unlink-control","native-process-control","go-self-control","self-exec-control","self-spawn-replace-control","native-copy-control","invalid-native-supervision-control","hook-control","agent-control","private-write-control","private-chmod-control","private-unlink-control":os.Exit(70)}
 json.NewEncoder(os.Stdout).Encode(out)
}`, build)
	worktreePut(t, build, "probe.go", program)
	source, err := filepath.Abs("../../..")
	if err != nil {
		t.Fatal(err)
	}
	worktreePut(t, build, "go.mod", fmt.Sprintf("module fireweedstartprobe\n\ngo 1.23\n\nrequire github.com/lurkshark/fireweed v0.0.0\nreplace github.com/lurkshark/fireweed => %q\n", source))
	worktreePut(t, build, "early/early.go", `package early
import "os"
func init(){if len(os.Args)==2&&os.Args[1]=="__fw_start_imported_go_control__"{os.Exit(0)}}
`)
	binary := filepath.Join(build, "probe")
	cmd = exec.CommandContext(ctx, "go", "build", "-mod=mod", "-o", binary, filepath.Join(build, "probe.go"))
	cmd.Dir = build
	cmd.Env = append(os.Environ(), "CGO_ENABLED=1", "GOPROXY=off", "GOTOOLCHAIN=local")
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("public start native probe compile: %v %s", err, out)
	}
	trapDir := filepath.Join(build, "traps")
	if err := os.Mkdir(trapDir, 0755); err != nil {
		t.Fatal(err)
	}
	executableBytes, err := os.ReadFile(binary)
	if err != nil {
		t.Fatal(err)
	}
	for _, name := range []string{"start-hook-trap", "start-agent-trap", "fireweed-bootstrap", "probe"} {
		if err := os.WriteFile(filepath.Join(trapDir, name), executableBytes, 0755); err != nil {
			t.Fatal(err)
		}
	}
	// The transport key belongs only to this observer. Remove an ambient
	// collision during fixture setup; the product caller environment remains
	// unchanged throughout each measured operation.
	probeEnvironment := []string{}
	for _, entry := range os.Environ() {
		if !strings.HasPrefix(entry, "__FIREWEED_TEST_START_POLICY_V1=") {
			probeEnvironment = append(probeEnvironment, entry)
		}
	}
	probeEnvironment = append(probeEnvironment, "PATH="+trapDir+string(os.PathListSeparator)+os.Getenv("PATH"))
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer listener.Close()
	run := func(t *testing.T, in startFaultInput) (startFaultOutput, int) {
		t.Helper()
		path := filepath.Join(t.TempDir(), "input.json")
		b, err := json.Marshal(in)
		if err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(path, b, 0600); err != nil {
			t.Fatal(err)
		}
		child := exec.CommandContext(ctx, binary, path)
		child.Dir = in.Root
		child.Env = append([]string(nil), probeEnvironment...)
		out, err := child.CombinedOutput()
		if child.Process != nil {
			cleanupContext, cleanupCancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cleanupCancel()
			cleanup := exec.CommandContext(cleanupContext, binary, "--cleanup-start-guard", fmt.Sprint(child.Process.Pid))
			cleanup.Env = append([]string(nil), probeEnvironment...)
			if cleanupOut, cleanupErr := cleanup.CombinedOutput(); cleanupErr != nil {
				t.Fatalf("native guard policy cleanup: %v %s", cleanupErr, cleanupOut)
			}
		}
		code := 0
		if err != nil {
			var ok bool
			var exit *exec.ExitError
			if exit, ok = err.(*exec.ExitError); ok {
				code = exit.ExitCode()
			} else {
				t.Fatal("native probe execution", err)
			}
		}
		if code != 0 {
			return startFaultOutput{}, code
		}
		var got startFaultOutput
		if err := json.Unmarshal(out, &got); err != nil {
			t.Fatalf("native probe JSON: %v %q", err, out)
		}
		got.Result.Claim.Object = got.Object
		return got, code
	}
	fixture := func(t *testing.T, kind string) (startFixtureData, startFaultInput) {
		t.Helper()
		f := startFixture(t, "sha1", kind, "")
		worktreePut(t, f.Root, "units/a/impl/second.txt", "Second original disposable bytes.\n")
		phase := "evals"
		if kind == "regenerate" {
			phase = "regenerate"
		}
		book, err := a.Playbook(phase)
		if err != nil {
			t.Fatal(err)
		}
		target := ".fireweed/playbooks/" + phase + ".md"
		worktreePut(t, f.Root, target, "---\nbased_on: '"+book.Hash+"'\n---\n\nExact selected phase instructions.\n")
		startCommit(t, f.Root, "Fault fixture retained sources")
		selected, err := p.ResolvePlaybook(f.Root, phase)
		if err != nil || !selected.Overridden {
			t.Fatal("public override fixture premise", selected, err)
		}
		return f, startFaultInput{Root: f.Root, Change: filepath.Join(f.Root, f.Path), Context: filepath.Join(f.Root, target), Private: filepath.Join(f.Root, "private/secret.txt"), Address: listener.Addr().String(), Fault: 1}
	}
	t.Run("token-transport-control", func(t *testing.T) {
		_, in := fixture(t, "spec")
		in.Mode = "token-transport-control"
		in.Fault = 0
		got, code := run(t, in)
		if code != 0 || got.Object != "native-token-transport" || got.Result.Claim.Object != got.Object || !got.StateSame {
			t.Fatal("native claim token transport is not exact", code, got)
		}
	})
	for _, mode := range []string{"stdin-control", "prompt-control", "stderr-control", "network-control", "process-control", "native-depth-control", "go-imported-init-control", "native-stdin-control", "native-socket-control", "native-connect-control", "native-write-control", "native-chmod-control", "native-unlink-control", "native-process-control", "go-self-control", "self-exec-control", "self-spawn-replace-control", "native-copy-control", "hook-control", "agent-control", "private-write-control", "private-chmod-control", "private-unlink-control"} {
		t.Run(mode, func(t *testing.T) {
			_, in := fixture(t, "spec")
			in.Mode = mode
			if _, code := run(t, in); code != 42 {
				t.Errorf("live ignored-error control exit %d want42", code)
			}
		})
	}
	for _, control := range []struct {
		Mode  string
		Fault int
	}{{"git-control", 1}, {"native-supervision-control", 1}, {"native-output-control", 1}, {"native-environment-control", 1}, {"relative-write-fault-control", 2}, {"relative-rename-fault-control", 2}, {"relative-context-fault-control", 5}, {"dotdot-context-fault-control", 5}, {"fd-context-fault-control", 5}, {"relative-delete-fault-control", 3}, {"relative-delete-fault-control", 4}, {"write-fault-control", 2}, {"context-fault-control", 5}, {"delete-fault-control", 3}, {"delete-fault-control", 4}} {
		t.Run(fmt.Sprintf("%s-%d", control.Mode, control.Fault), func(t *testing.T) {
			_, in := fixture(t, "regenerate")
			in.Mode = control.Mode
			in.Fault = control.Fault
			got, code := run(t, in)
			if code != 0 || !got.Control || !got.StateSame || control.Fault != 1 && !got.Fired {
				t.Errorf("native fault positive control exit%d %+v", code, got)
			}
		})
	}
	for _, mode := range []string{"invalid-read-control", "invalid-stat-control", "invalid-git-control", "invalid-native-supervision-control", "invalid-relative-read-control", "invalid-relative-stat-control", "invalid-subject"} {
		t.Run(mode, func(t *testing.T) {
			f, in := fixture(t, "spec")
			in.Mode = mode
			in.Fault = 6
			checkStorage := startStorage(t, f.Root, "current")
			defer checkStorage()
			before := startTree(t, f.Root)
			git := startGitSnapshot(t, f.Root)
			got, code := run(t, in)
			if mode != "invalid-subject" {
				if code != 42 {
					t.Error("invalid-selector access guard not live", mode, code)
				}
				return
			}
			if code != 0 || !strings.Contains(got.Error, "C05") || !policyCorrective(got.Error) || !reflect.DeepEqual(got.Result, w.StartResult{}) || !got.StateSame {
				t.Error("invalid selector performed repository I/O or returned invalid error", code, got)
			}
			if !reflect.DeepEqual(before, startTree(t, f.Root)) || !reflect.DeepEqual(git, startGitSnapshot(t, f.Root)) {
				t.Error("invalid selector changed repository")
			}
		})
	}
	proxyBinary := startProxyBuild(t)
	for _, mode := range []string{"precondition-write-control", "precondition-claim-control", "precondition-read-control", "precondition-dirty", "precondition-incomplete", "precondition-missing", "precondition-ambiguous", "precondition-config", "precondition-queue", "precondition-source", "precondition-blocked"} {
		t.Run(mode, func(t *testing.T) {
			f, in := fixture(t, "spec")
			in.Fault = 7
			in.Mode = mode
			in.Context = filepath.Join(claimCommon(t, f.Root), "refs/fireweed/claims")
			code := ""
			switch mode {
			case "precondition-dirty":
				worktreePut(t, f.Root, "retained", "dirty")
				code = "P01"
			case "precondition-incomplete":
				worktreePut(t, f.Root, f.Path, strings.Replace(f.Raw, "module: alpha\n", "module: alpha\nrun: incomplete\n", 1))
				startCommit(t, f.Root, "Incomplete started state")
			case "precondition-source", "precondition-read-control":
				in.Fault = 8
				code = ""
			case "precondition-ambiguous":
				worktreePut(t, f.Root, "changes/0002-current.md", f.Raw)
				startCommit(t, f.Root, "Ambiguous selected slug")
			case "precondition-queue":
				worktreePut(t, f.Root, f.Path, strings.ReplaceAll(f.Raw, "new:feature", "ALP-099"))
				startCommit(t, f.Root, "Invalid current queue metadata")
			case "precondition-config":
				worktreePut(t, f.Root, "fireweed.yaml", "version: nonsense\n")
				startCommit(t, f.Root, "Invalid static metadata")
				code = ""
			case "precondition-blocked":
				worktreePut(t, f.Root, f.Path, strings.Replace(f.Raw, "module: alpha\n", "module: alpha\nblocked: hold\n", 1))
				startCommit(t, f.Root, "Blocked precondition")
				code = "P03"
			}
			checkNoClaim := startPreconditionProxy(t, proxyBinary, f)
			defer checkNoClaim()
			checkStorage := startStorage(t, f.Root, "current")
			defer checkStorage()
			before := startTree(t, f.Root)
			git := startGitSnapshot(t, f.Root)
			got, status := run(t, in)
			if mode == "precondition-read-control" {
				if status != 0 || !got.Control || !got.Fired {
					t.Error("precondition source-read fault not live", status, got)
				}
				return
			}
			if strings.HasSuffix(mode, "-control") {
				if status != 42 {
					t.Error("transient mutation guard not live", mode, status)
				}
				return
			}
			if mode == "precondition-source" && !got.Fired {
				t.Error("Start did not reach required initial source read")
			}
			if status != 0 || !strings.Contains(got.Error, code) || !policyCorrective(got.Error) || !reflect.DeepEqual(got.Result, w.StartResult{}) || !got.StateSame {
				t.Error("precondition performed transient mutation or returned invalid rejection", mode, status, got)
			}
			if !reflect.DeepEqual(before, startTree(t, f.Root)) || !reflect.DeepEqual(git, startGitSnapshot(t, f.Root)) {
				t.Error("precondition changed final repository state")
			}
		})
	}

	for _, test := range []struct {
		Name, Kind string
		Fault      int
	}{{"success", "spec", 1}, {"source-read", "spec", 5}, {"interface-read", "spec", 5}, {"drill-read", "regenerate", 5}, {"selected-write", "spec", 2}, {"context-read", "spec", 5}, {"regenerate-context", "regenerate", 5}, {"second-delete", "regenerate", 3}, {"symlink-delete", "regenerate", 3}, {"recreated-delete", "regenerate", 4}} {
		t.Run(test.Name, func(t *testing.T) {
			f, in := fixture(t, test.Kind)
			in.Mode = "subject"
			if test.Name == "symlink-delete" {
				outside := filepath.Join(t.TempDir(), "target")
				if err := os.WriteFile(outside, []byte("Untouched external symlink target.\n"), 0640); err != nil {
					t.Fatal(err)
				}
				for _, leaf := range []string{"code.txt", "second.txt"} {
					path := filepath.Join(f.Root, "units/a/impl", leaf)
					if err := os.Remove(path); err != nil {
						t.Fatal(err)
					}
					if err := os.Symlink(outside, path); err != nil {
						t.Fatal(err)
					}
				}
				startCommit(t, f.Root, "Tracked disposable leaf symlinks for rollback")
				external := startTree(t, filepath.Dir(outside))
				defer func() {
					if !reflect.DeepEqual(external, startTree(t, filepath.Dir(outside))) {
						t.Error("rollback touched external symlink target")
					}
				}()
			}
			if test.Name == "context-read" {
				if err := os.Chmod(filepath.Join(f.Root, f.Path), 0751); err != nil {
					t.Fatal(err)
				}
				f.Mode = 0751
				startCommit(t, f.Root, "Executable original change mode for rollback")
			}
			switch test.Name {
			case "source-read":
				in.Context = in.Change
			case "interface-read":
				in.Context = filepath.Join(f.Root, "units/a/api.txt")
			case "drill-read":
				startReport(t, f, "pass", claimBase(t, f.Root))
				in.Context = filepath.Join(f.Root, ".fireweed/drills/alpha.md")
			}
			in.Fault = test.Fault
			checkStorage := startStorage(t, f.Root, "current")
			defer checkStorage()
			before := startTree(t, f.Root)
			git := startGitSnapshot(t, f.Root)
			got, code := run(t, in)
			if code != 0 {
				t.Fatalf("subject attempted forbidden interaction or crashed: exit%d", code)
			}
			if !got.StateSame {
				t.Error("fault/success changed the subject process cwd/environment")
			}
			if test.Fault == 1 || test.Name == "drill-read" {
				if got.Error != "" || got.Result.Claim.Object == "" {
					t.Fatalf("valid guarded Start: %+v", got)
				}
				loaded, err := w.LoadClaim(f.Root, "current")
				if err != nil || loaded != got.Result.Claim {
					t.Error("guarded claim mismatch", loaded, err)
				}
				delete(before, f.Path)
				if test.Name == "drill-read" {
					if !got.Fired {
						t.Error("report read fault not reached")
					}
					warning := false
					for _, d := range got.Result.Diagnostics {
						if d.Code == "D01" && d.Severity == "warning" {
							warning = true
						}
					}
					if !warning {
						t.Error("unreadable passing report did not warn", got.Result.Diagnostics)
					}
					delete(before, "units/a/impl/code.txt")
					delete(before, "units/a/impl/second.txt")
				}
				after := startTree(t, f.Root)
				delete(after, f.Path)
				if !reflect.DeepEqual(before, after) {
					t.Error("guarded success changed unrelated files")
				}
				startUnrelated(t, f.Root, git, "current")
				return
			}
			if !got.Fired {
				t.Error("start did not reach its required public write/deletion/context-read boundary")
			}
			if got.Error == "" || !policyCorrective(got.Error) || !reflect.DeepEqual(got.Result, w.StartResult{}) {
				t.Errorf("fault did not reject atomically: %+v", got)
			}
			after := startTree(t, f.Root)
			if test.Fault == 4 {
				if got.Recreated == "" {
					t.Fatal("no externally recreated deletion premise")
				}
				rel, err := filepath.Rel(f.Root, got.Recreated)
				if err != nil {
					t.Fatal(err)
				}
				rel = filepath.ToSlash(rel)
				node, ok := after[rel]
				if !ok || node.Bytes != "Concurrent winner bytes.\n" || node.Mode.Perm() != 0640 {
					t.Error("overwrote concurrent recreated path", node)
				}
				before[rel] = node
				text := strings.ToLower(got.Error)
				if !strings.Contains(text, "recover") && !strings.Contains(text, "restore") && !strings.Contains(text, "ownership") && !strings.Contains(text, "preserv") {
					t.Error("lost incomplete-recovery detail", got.Error)
				}
			}
			if !reflect.DeepEqual(before, after) {
				t.Error("fault rollback changed exact original files/type/modes")
			}
			if !reflect.DeepEqual(git, startGitSnapshot(t, f.Root)) {
				t.Error("fault rollback changed claim/refs/index/HEAD/parent state")
			}
		})
	}
}
