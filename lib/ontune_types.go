package lib

const (
	DATAKEY_CODE = 0x00000001
)

const (
	HOST_KEY     = 1
	LASTPERF_KEY = 2
	BASIC_KEY    = 4
	CPU_KEY      = 8
	MEM_KEY      = 16
	NET_KEY      = 32
	DISK_KEY     = 64
)

type DataKey struct {
	Code uint32  `json:"code"`
	Key  Bitmask `json:"key"`
}