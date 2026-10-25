// Package supervision defines the private protocol for an isolated Linux reaper.
package supervision

const (
	Version   = 1
	ImageFD   = 3
	RequestFD = 4
	ReplyFD   = 5
	CancelFD  = 6
)

type Request struct {
	Version           int
	Nonce, Executable string
	Args, Env         []string
	Combined          bool
}

type Reply struct {
	Version                                      int
	Nonce                                        string
	Ready, Final, HasRoot, Canceled, LaunchError bool
	Exit, Errno                                  int
	Error                                        string
}
