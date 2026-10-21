package wasi_snapshot_preview1_test

import (
	"io"
	"testing"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/internal/testing/require"
	"github.com/tetratelabs/wazero/internal/wasip1"
)

func Test_fdWrite_BrokenPipe(t *testing.T) {
	reader, writer := io.Pipe()
	require.NoError(t, reader.Close())
	defer writer.Close()
	mod, r, _ := requireProxyModule(t, wazero.NewModuleConfig().WithStdout(writer))
	defer r.Close(testCtx)
	// A one-byte iovec points at the payload following the vector.
	require.True(t, mod.Memory().Write(0, []byte{8, 0, 0, 0, 1, 0, 0, 0, 'x'}))
	requireErrnoResult(t, wasip1.ErrnoPipe, mod, wasip1.FdWriteName, 1, 0, 1, 12)
}
