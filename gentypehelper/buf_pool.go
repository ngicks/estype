package gentypehelper

import (
	"bytes"
	"sync"
)

const (
	IdGetBuf = "GetBuf"
	IdPutBuf = "PutBuf"
)

var bufPool = &sync.Pool{
	New: func() any {
		return new(bytes.Buffer)
	},
}

func GetBuf() *bytes.Buffer {
	return bufPool.Get().(*bytes.Buffer)
}

func PutBuf(v *bytes.Buffer) {
	v.Reset()
	bufPool.Put(v)
}
