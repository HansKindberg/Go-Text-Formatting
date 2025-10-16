package textfmt

import (
	"io"
)

type TextFormatter[T Options] interface {
	Format(T, io.Reader, io.Writer) error
}
