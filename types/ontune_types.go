package types

const (
	DATAKEY_CODE  = 0x00000001
	HOST_CODE     = 0x00000002
	LASTPERF_CODE = 0x00000003
	BASIC_CODE    = 0x00000004
	CPU_CODE      = 0x00000005
	MEM_CODE      = 0x00000006
	NET_CODE      = 0x00000007
	DISK_CODE     = 0x00000008
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

type DataCode struct {
	Code uint32 `json:"code"`
}