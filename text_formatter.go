package textfmt

type TextFormatter[T Options] interface {
	Format(T, string) (string, error)
}
