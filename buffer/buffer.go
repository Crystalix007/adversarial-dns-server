package buffer

type Buffer struct {
	buffer []byte
}

func NewFromBytes(b []byte) Buffer {
	return Buffer{
		buffer: b,
	}
}

func (b *Buffer) Prepend(p []byte) {
	b.buffer = append(p, b.buffer...)
}

func (b *Buffer) Dequeue(elemCount int) []byte {
	if elemCount < 0 {
		elemCount = len(b.buffer) + elemCount
	} else if elemCount > len(b.buffer) {
		elemCount = len(b.buffer)
	}

	r := b.buffer[:elemCount]
	b.buffer = b.buffer[elemCount:]

	return r
}

func (b *Buffer) Length() int {
	return len(b.buffer)
}
