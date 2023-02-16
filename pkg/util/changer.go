package util

func ChangeTypeToTypePointer[T comparable](typ T) *T {
	if !(typ == *new(T)) {
		return &typ
	}

	return nil
}
