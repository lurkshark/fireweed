package evals

import (
	"context"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"testing"
	"time"
)

func Test_AST001_AST002SandboxIndependence(t *testing.T) {
	if runtime.GOOS != "darwin" {
		t.Skip("native file/network denial probe is verified on macOS; ordinary embedded content checks remain portable")
	}
	root, err := filepath.Abs("../../..")
	if err != nil {
		t.Fatal(err)
	}
	temp := t.TempDir()
	source := filepath.Join(temp, "probe.go")
	binary := filepath.Join(temp, "probe")
	program := `package main
import("fmt";"net";"os";"time"; a "github.com/lurkshark/fireweed/packages/assets")
func main(){
 if len(os.Args)>1&&os.Args[1]=="--guard"{
  if _,err:=os.ReadFile(os.Args[2]);err==nil{os.Exit(3)}
  if _,err:=os.ReadFile(os.Args[4]);err==nil{os.Exit(9)}
  if conn,err:=net.DialTimeout("tcp",os.Args[3],time.Second);err==nil{conn.Close();os.Exit(4)}
  fmt.Println("DENIAL_GUARD_OK");return
 }
 if a.Info().Version!="0.4.0"{os.Exit(5)}
 for _,name:=range a.Playbooks(){text,err:=a.Playbook(name);if err!=nil||text.Text==""||text.Hash==""{os.Exit(6)}}
 if len(a.Playbooks())!=5{os.Exit(7)}
 for _,code:=range a.Codes(){if text,err:=a.Explain(code);err!=nil||text==""{os.Exit(8)}}
 fmt.Println("EMBEDDED_ASSETS_OK")
}`
	if err := os.WriteFile(source, []byte(program), 0644); err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	build := exec.CommandContext(ctx, "go", "build", "-o", binary, source)
	build.Dir = root
	build.Env = append(os.Environ(), "GOPROXY=off", "GOTOOLCHAIN=local")
	if output, err := build.CombinedOutput(); err != nil {
		t.Fatalf("build isolated public-API probe: %v %s", err, output)
	}
	root, err = filepath.EvalSymlinks(root)
	if err != nil {
		t.Fatal(err)
	}
	binary, err = filepath.EvalSymlinks(binary)
	if err != nil {
		t.Fatal(err)
	}
	sentinel := filepath.Join(temp, "forbidden-source-data")
	if err := os.WriteFile(sentinel, []byte("private data"), 0644); err != nil {
		t.Fatal(err)
	}
	sentinel, err = filepath.EvalSymlinks(sentinel)
	if err != nil {
		t.Fatal(err)
	}
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer listener.Close()
	// Deny all workspace sources (including any sibling bootstrap), a private
	// file-read control, and network access. OS loader/cache reads remain usable.
	profile := "(version 1)(allow default)(deny file-read-data (subpath " + strconv.Quote(filepath.Dir(root)) + ") (literal " + strconv.Quote(sentinel) + "))(deny network*)"
	unrestricted := exec.CommandContext(ctx, binary, "--guard", sentinel, listener.Addr().String(), filepath.Join(root, "go.mod"))
	if err := unrestricted.Run(); err == nil {
		t.Fatal("negative guard control did not detect permitted file access")
	}
	guarded := exec.CommandContext(ctx, "/usr/bin/sandbox-exec", "-p", profile, binary, "--guard", sentinel, listener.Addr().String(), filepath.Join(root, "go.mod"))
	guarded.Dir = temp
	if output, err := guarded.CombinedOutput(); err != nil || !strings.Contains(string(output), "DENIAL_GUARD_OK") {
		t.Fatalf("native deny policy unavailable or ineffective: %v %s", err, output)
	}
	command := exec.CommandContext(ctx, "/usr/bin/sandbox-exec", "-p", profile, binary)
	command.Dir = temp
	if output, err := command.CombinedOutput(); err != nil || !strings.Contains(string(output), "EMBEDDED_ASSETS_OK") {
		t.Fatalf("assets depend on source/files/network: %v %s", err, output)
	}
}
