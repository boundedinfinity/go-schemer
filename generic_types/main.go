package generic_types

type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type Floating interface {
	~float32 | ~float64
}

type Numeric interface {
	Signed | Unsigned | Floating
}
