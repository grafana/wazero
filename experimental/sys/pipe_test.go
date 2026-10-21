//go:build !(plan9 || aix)

package sys

import (
	"errors"
	"syscall"
	"testing"
)

func TestBrokenPipeErrno(t *testing.T) {
	if got := UnwrapOSError(syscall.EPIPE); got != EPIPE {
		t.Fatalf("errno = %v", got)
	}
	if !errors.Is(EPIPE, syscall.EPIPE) {
		t.Fatal("EPIPE lost syscall identity")
	}
	if EPIPE.Error() != "broken pipe" {
		t.Fatalf("message = %s", EPIPE.Error())
	}
}
