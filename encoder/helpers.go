package encoder

import "encoding/base64"

func toPtr[T any](constVar T) *T { return &constVar }

func toPtrOrNil[T comparable](comparable T) *T {
	var zero T
	if comparable == zero {
		return nil
	}
	return &comparable
}

func urlEncodeBytesPtrOrNil(b []byte) *string {
	if b == nil || len(b) == 0 || isZeros(b) {
		return nil
	}
	return toPtr(base64.RawURLEncoding.EncodeToString(b))
}

func isZeros(b []byte) bool {
	for i := 0; i < len(b); i++ {
		if b[i] != 0 {
			return false
		}
	}
	return true
}
