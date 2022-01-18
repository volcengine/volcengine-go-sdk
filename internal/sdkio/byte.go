package sdkio

import "io"

const (
	// Byte is 8 bits
	Byte int64 = 1
	// KibiByte (KiB) is 1024 Bytes
	KibiByte = Byte * 1024
	// MebiByte (MiB) is 1024 KiB
	MebiByte = KibiByte * 1024
	// GibiByte (GiB) is 1024 MiB
	GibiByte = MebiByte * 1024
)

const (
	SeekStart   = io.SeekStart   // seek relative to the origin of the file
	SeekCurrent = io.SeekCurrent // seek relative to the current offset
	SeekEnd     = io.SeekEnd     // seek relative to the end
)
