// Native entry for a per-call Darwin ownership anchor. This source is rebuilt
// into the architecture-specific syso files; it is never compiled at runtime.
#include <crt_externs.h>
#include <dlfcn.h>
#include <errno.h>
#include <fcntl.h>
#include <libproc.h>
#include <poll.h>
#include <signal.h>
#include <spawn.h>
#include <stdint.h>
#include <stdlib.h>
#include <string.h>
#include <sys/mman.h>
#include <sys/stat.h>
#include <sys/wait.h>
#include <unistd.h>

#define REQUEST_MAGIC UINT64_C(0x667770726f633031)
#define REPLY_MAGIC UINT64_C(0x66777265706c3031)
#define MAX_REQUEST (16U << 20)
static const char entry_marker[] = "__fireweed_process_anchor_v1__";
typedef pid_t (*responsible_fn)(pid_t);
static responsible_fn responsible;

static int write_all(int fd, const void *raw, size_t size) {
 const unsigned char *p = raw;
 while(size) { ssize_t n=write(fd,p,size);if(n<0&&errno==EINTR)continue;if(n<=0)return -1;p+=n;size-=(size_t)n; }
 return 0;
}
static void reply(uint64_t stage, int code, int error) {
 uint64_t words[5]={REPLY_MAGIC,stage,(uint64_t)(int64_t)code,(uint64_t)error,0};
 (void)write_all(5,words,sizeof(words));
}
static void finish(int code,int error) { reply(2,code,error);_exit(error||code<0?0:code); }
static void finish_signal(int number) {
 reply(2,-1,0);struct sigaction action={0};action.sa_handler=SIG_DFL;
 (void)sigaction(number,&action,0);sigset_t empty;sigemptyset(&empty);(void)sigprocmask(SIG_SETMASK,&empty,0);
 (void)kill(getpid(),number);_exit(128+number);
}
static int canceled(void) { struct pollfd p={4,POLLIN,0};return poll(&p,1,0)>0&&(p.revents&(POLLIN|POLLHUP|POLLERR|POLLNVAL)); }

struct lifetime { unsigned char uuid[16];uint64_t unique,parent;int32_t version,parent_version;uint64_t reserved[2]; };
struct identity { struct proc_bsdinfo bsd;struct lifetime life; };
struct owned { int pid;struct identity identity; };
static int inspect(int pid, struct identity *info) { memset(info,0,sizeof(*info));return proc_pidinfo(pid,18,1,info,sizeof(*info))==(int)sizeof(*info)&&info->bsd.pbi_pid==(uint32_t)pid; }
static int same(const struct identity *a,const struct identity *b) { return a->life.unique==b->life.unique&&a->bsd.pbi_start_tvsec==b->bsd.pbi_start_tvsec&&a->bsd.pbi_start_tvusec==b->bsd.pbi_start_tvusec; }
static int seen(struct owned *all,size_t count,int pid) { for(size_t i=0;i<count;i++)if(all[i].pid==pid)return 1;return 0; }
static int edge(struct owned *all,size_t count,const struct identity *info) { for(size_t i=0;i<count;i++)if(all[i].identity.life.unique&&all[i].identity.life.unique==info->life.parent)return 1;return 0; }
static void signal_owned(const struct owned *p,int signal) { struct identity now;if(inspect(p->pid,&now)&&same(&now,&p->identity)&&now.bsd.pbi_status!=5)(void)kill(p->pid,signal); }

static void stop_tree(pid_t root) {
 // The live anchor's responsibility identity survives arbitrary ordinary
 // parent exits, reaping, independent process groups and independent sessions.
 // Retain original-parent identities as well for nested native anchors.
 size_t capacity=64,count=0;struct owned *all=calloc(capacity,sizeof(*all));
 if(!all){(void)kill(root,SIGKILL);return;}
 pid_t self=getpid();
 // The root may already be reaped while descendants retain the streams.
 // Do not seed a later unrelated reuse of its numeric PID as owned.
 struct identity initial;if(inspect(root,&initial)&&responsible(root)==self)all[count++]=(struct owned){root,initial};
 for(int round=0;round<100;round++) {
  int size=proc_listallpids(0,0);if(size<1)break;size+=64;
  int *pids=calloc((size_t)size,sizeof(int));if(!pids)break;
  int n=proc_listallpids(pids,size*(int)sizeof(int));int added=0;
  for(int i=0;i<n;i++) {
   int pid=pids[i];if(pid<=1||pid==self||seen(all,count,pid))continue;
   struct identity info;if(!inspect(pid,&info)||info.bsd.pbi_status==5)continue;
   if(responsible(pid)!=self&&!edge(all,count,&info))continue;
   if(count==capacity){size_t next=capacity*2;struct owned *grown=realloc(all,next*sizeof(*all));if(!grown)break;all=grown;capacity=next;}
   all[count++]=(struct owned){pid,info};added=1;
  }
  free(pids);
  int pending=0;
  for(size_t i=0;i<count;i++) {
   struct identity now;if(!inspect(all[i].pid,&now)||!same(&now,&all[i].identity)||now.bsd.pbi_status==5)continue;
   if(now.bsd.pbi_status!=4){signal_owned(&all[i],SIGSTOP);pending=1;}
  }
  // Re-scan after stops are acknowledged, so a last fork cannot escape the
  // inventory merely because its parent was between enumeration and SIGSTOP.
  if(!added&&!pending)break;
  usleep(1000);
 }
 for(size_t i=count;i>0;i--)if(all[i-1].pid!=root)signal_owned(&all[i-1],SIGKILL);
 for(size_t i=0;i<count;i++)if(all[i].pid==root)signal_owned(&all[i],SIGKILL);
 // Do not let a signalled but still executing owned process outlive reply.
 for(int round=0;round<500;round++) {
  int alive=0;for(size_t i=0;i<count;i++){struct identity now;if(inspect(all[i].pid,&now)&&same(&now,&all[i].identity)&&now.bsd.pbi_status!=5){alive=1;signal_owned(&all[i],SIGKILL);}}
  if(!alive)break;usleep(1000);
 }
 free(all);
}

static char *next_string(char **cursor,char *end) {
 if(*cursor>=end)return 0;char *start=*cursor;char *zero=memchr(start,0,(size_t)(end-start));if(!zero)return 0;*cursor=zero+1;return start;
}
static void anchor(void) {
 struct sigaction ignore={0};ignore.sa_handler=SIG_IGN;sigaction(SIGPIPE,&ignore,0);
 struct stat st;if(fstat(3,&st)||st.st_size<64||st.st_size>MAX_REQUEST)finish(2,EINVAL);
 uint64_t *header=mmap(0,(size_t)st.st_size,PROT_READ,MAP_SHARED,3,0);if(header==MAP_FAILED)finish(2,errno);
 size_t length=(size_t)header[1],argc=(size_t)header[2],envc=(size_t)header[3];
 if(header[0]!=REQUEST_MAGIC||length<64||length>(size_t)st.st_size||argc<1||argc>65536||envc>65536)finish(2,EINVAL);
 char **args=calloc(argc+1,sizeof(char*)),**env=calloc(envc+1,sizeof(char*));if(!args||!env)finish(2,ENOMEM);
 char *cursor=(char*)header+64,*end=(char*)header+length,*path=next_string(&cursor,end);if(!path||!*path)finish(2,EINVAL);
 for(size_t i=0;i<argc;i++)if(!(args[i]=next_string(&cursor,end)))finish(2,EINVAL);
 for(size_t i=0;i<envc;i++)if(!(env[i]=next_string(&cursor,end)))finish(2,EINVAL);
 if(cursor!=end||!args[0][0])finish(2,EINVAL);
 (void)fcntl(3,F_SETFD,FD_CLOEXEC);(void)fcntl(4,F_SETFD,FD_CLOEXEC);(void)fcntl(5,F_SETFD,FD_CLOEXEC);
 responsible=(responsible_fn)dlsym(RTLD_DEFAULT,"responsibility_get_pid_responsible_for_pid");
 if(!responsible||responsible(getpid())!=getpid())finish(2,ENOTSUP);
 if(canceled())finish(-1,ECANCELED);
 int out[2],err[2];if(pipe(out))finish(2,errno);int combined=header[4]!=0;
 if(combined){err[0]=-1;err[1]=out[1];}else if(pipe(err)){close(out[0]);close(out[1]);finish(2,errno);}
 for(int i=0;i<2;i++){(void)fcntl(out[i],F_SETFD,FD_CLOEXEC);if(!combined)(void)fcntl(err[i],F_SETFD,FD_CLOEXEC);}
 posix_spawn_file_actions_t files;posix_spawnattr_t attr;int error=posix_spawn_file_actions_init(&files);if(error)finish(2,error);
 error=posix_spawnattr_init(&attr);if(error)finish(2,error);
 sigset_t empty,defaults;sigemptyset(&empty);sigfillset(&defaults);
 if(!(error=posix_spawnattr_setsigmask(&attr,&empty)))error=posix_spawnattr_setsigdefault(&attr,&defaults);
 if(!error)error=posix_spawnattr_setpgroup(&attr,0);
 if(!error)error=posix_spawnattr_setflags(&attr,POSIX_SPAWN_SETSIGMASK|POSIX_SPAWN_SETSIGDEF|POSIX_SPAWN_SETPGROUP|POSIX_SPAWN_CLOEXEC_DEFAULT);
 if(!error)error=posix_spawn_file_actions_adddup2(&files,out[1],1);
 if(!error)error=posix_spawn_file_actions_adddup2(&files,err[1],2);
 if(!error)error=posix_spawn_file_actions_addinherit_np(&files,0);
 pid_t root=0;if(!error)error=posix_spawn(&root,path,&files,&attr,args,env);
 posix_spawnattr_destroy(&attr);posix_spawn_file_actions_destroy(&files);close(out[1]);if(!combined)close(err[1]);
 if(error)finish(2,error);
 // The target has its original argv/env/cwd and no protocol descriptors.
 reply(1,0,0);free(args);free(env);munmap(header,(size_t)st.st_size);close(3);
 int root_done=0,status=0,open_streams=combined?1:2;
 struct pollfd watching[3]={{4,POLLIN,0},{out[0],POLLIN,0},{combined?-1:err[0],POLLIN,0}};
 while(!root_done||open_streams) {
  if(!root_done){pid_t got=waitpid(root,&status,WNOHANG);if(got==root)root_done=1;else if(got<0&&errno!=EINTR){stop_tree(root);finish(2,errno);}}
  int events=poll(watching,3,10);if(events<0&&errno==EINTR)continue;if(events<0){stop_tree(root);finish(2,errno);}
  if(watching[0].revents&(POLLIN|POLLHUP|POLLERR|POLLNVAL)) {
   stop_tree(root);if(!root_done)while(waitpid(root,&status,0)<0&&errno==EINTR){};
   // Drain already buffered bytes after termination; pipe EOF follows once
   // every owned stream holder is gone.
   watching[0].fd=-1;
   for(int i=1;i<3;i++)if(watching[i].fd>=0){int flags=fcntl(watching[i].fd,F_GETFL);if(flags>=0)(void)fcntl(watching[i].fd,F_SETFL,flags|O_NONBLOCK);char buffer[32768];for(;;){ssize_t n=read(watching[i].fd,buffer,sizeof(buffer));if(n>0)(void)write_all(i,buffer,(size_t)n);else if(n<0&&errno==EINTR)continue;else break;}close(watching[i].fd);}
   finish(-1,ECANCELED);
  }
  for(int i=1;i<3;i++)if(watching[i].fd>=0&&watching[i].revents&(POLLIN|POLLHUP|POLLERR)) {
   char buffer[32768];ssize_t n=read(watching[i].fd,buffer,sizeof(buffer));
   if(n>0){if(write_all(i,buffer,(size_t)n)){stop_tree(root);finish(2,EPIPE);}}
   else if(n==0){close(watching[i].fd);watching[i].fd=-1;open_streams--;}
   else if(errno!=EINTR){stop_tree(root);finish(2,errno);}
  }
 }
 if(WIFSIGNALED(status))finish_signal(WTERMSIG(status));
 finish(WIFEXITED(status)?WEXITSTATUS(status):-1,0);
}
static void early_entry(void) {
 int argc=*_NSGetArgc();char **argv=*_NSGetArgv();
 if(argc!=3||strcmp(argv[1],entry_marker)||strlen(argv[2])!=48)return;
 struct stat st;if(fstat(3,&st)||(st.st_mode&S_IFMT)!=0||st.st_size<64||st.st_size>MAX_REQUEST)return;
 unsigned char *header=mmap(0,64,PROT_READ,MAP_SHARED,3,0);if(header==MAP_FAILED)return;
 if(*(uint64_t*)header!=REQUEST_MAGIC){munmap(header,64);return;}
 static const char hex[]="0123456789abcdef";
 for(int i=0;i<24;i++)if(argv[2][i*2]!=hex[header[40+i]>>4]||argv[2][i*2+1]!=hex[header[40+i]&15]){munmap(header,64);return;}
 munmap(header,64);anchor();_exit(125);
}
// External cgo links retain this constructor. Internal non-cgo links use the
// pre-runtime _cgo_init slot instead; neither path runs Go initializers/main.
__attribute__((constructor)) static void constructor_entry(void){early_entry();}
void fireweed_early_entry(void*g,void*setg,void**tlsg,void**base){early_entry();}
static void early_tls(void**tlsg,void**base){early_entry();}
void (*x_cgo_inittls)(void**,void**)=early_tls;

int fireweed_spawn_anchor(const char *path,const char *cwd,const char *nonce,int request,int cancel,int result,int out,int err,int *pid) {
 int (*disclaim)(posix_spawnattr_t*,int)=dlsym(RTLD_DEFAULT,"responsibility_spawnattrs_setdisclaim");
 if(!disclaim||!dlsym(RTLD_DEFAULT,"responsibility_get_pid_responsible_for_pid"))return ENOTSUP;
 int source[5]={request,cancel,result,out,err},copy[5];for(int i=0;i<5;i++)copy[i]=-1;
 int error=0;for(int i=0;i<5;i++){copy[i]=fcntl(source[i],F_DUPFD_CLOEXEC,10);if(copy[i]<0){error=errno;goto end;}}
 posix_spawn_file_actions_t files;posix_spawnattr_t attr;error=posix_spawn_file_actions_init(&files);if(error)goto end;
 error=posix_spawnattr_init(&attr);if(error){posix_spawn_file_actions_destroy(&files);goto end;}
 if(!(error=disclaim(&attr,1)))error=posix_spawnattr_setpgroup(&attr,0);
 sigset_t empty,defaults;sigemptyset(&empty);sigfillset(&defaults);
 if(!error)error=posix_spawnattr_setsigmask(&attr,&empty);
 if(!error)error=posix_spawnattr_setsigdefault(&attr,&defaults);
 if(!error)error=posix_spawnattr_setflags(&attr,POSIX_SPAWN_SETPGROUP|POSIX_SPAWN_SETSIGMASK|POSIX_SPAWN_SETSIGDEF|POSIX_SPAWN_CLOEXEC_DEFAULT);
 int destinations[5]={3,4,5,1,2};for(int i=0;i<5&&!error;i++)error=posix_spawn_file_actions_adddup2(&files,copy[i],destinations[i]);
 if(!error)error=posix_spawn_file_actions_addopen(&files,0,"/dev/null",O_RDONLY,0);
 if(!error&&cwd&&*cwd)error=posix_spawn_file_actions_addchdir_np(&files,cwd);
 char *args[]={(char*)path,(char*)entry_marker,(char*)nonce,0},*env[]={0};
 if(!error)error=posix_spawn(pid,path,&files,&attr,args,env);
 posix_spawnattr_destroy(&attr);posix_spawn_file_actions_destroy(&files);
end:
 for(int i=0;i<5;i++)if(copy[i]>=0)close(copy[i]);return error;
}
