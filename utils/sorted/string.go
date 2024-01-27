package sorted

type StringFunc func(i, j string) bool

var String StringFunc = StringDesc

func StringDesc(i, j string) bool {
	return i > j
}

func StringAsc(i, j string) bool {
	return i < j
}
