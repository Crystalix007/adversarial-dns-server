package message_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/Crystalix007/adversarial-dns-server/dns/message"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestADecoding(t *testing.T) {
	t.Parallel()

	message.Decode(loadTestcase(t, "A-test.example.com"))

	assert.Equal(t)
}

func loadTestcase(t *testing.T, name string) []byte {
	path := filepath.Join("testdata", name)
	b, err := os.ReadFile(path)

	require.NoErrorf(t, err, "no error loading testcase '%s': %w", name, err)

	return b
}
