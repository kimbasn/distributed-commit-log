package log

type Config struct {
	Segment struct {
		MaxStoreBytes  uint64 // max .store file size before rotation
		MaxIndexBytes  uint64 // max .index file size before rotation
		InitialOffset  uint64 // logical offset to start this segment at
	}
}
