package types

type Bitmask uint32

func (value Bitmask) IsSet(key Bitmask) bool {
	if value&key != 0 {
		return true
	}
	return false
}