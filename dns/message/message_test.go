package message_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/Crystalix007/adversarial-dns-server/buffer"
	"github.com/Crystalix007/adversarial-dns-server/dns/message"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestADecoding(t *testing.T) {
	t.Parallel()

	m, err := message.Decode(loadTestcase(t, "A-test.example.com"))

	require.NoError(t, err)
	assert.Equal(t, uint16(0xd31a), m.Header.ID.ID)
	assert.Equal(t, 1, len(m.Question))
	question := m.Question[0]
	// 4 components, because of the root element
	// "A-test" "example" "com" ""
	assert.Equal(t, 4, len(question.QName))
	qnames := question.QName
	assert.Equal(t, []byte("A-test"), qnames[0].Label)
	assert.Equal(t, []byte("example"), qnames[1].Label)
	assert.Equal(t, []byte("com"), qnames[2].Label)
	assert.Equal(t, []byte(""), qnames[3].Label)
}

func loadTestcase(t *testing.T, name string) buffer.Buffer {
	path := filepath.Join("testdata", name)
	b, err := os.ReadFile(path)

	require.NoErrorf(t, err, "no error loading testcase '%s': %w", name, err)

	return buffer.NewFromBytes(b)
}
