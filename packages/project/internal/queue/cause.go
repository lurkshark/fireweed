package queue

import "github.com/lurkshark/fireweed/packages/contracts"

func errorCode(err error, fallback string) string {
	if code, _, ok := contracts.ErrorCode(err); ok {
		return code
	}
	return fallback
}
