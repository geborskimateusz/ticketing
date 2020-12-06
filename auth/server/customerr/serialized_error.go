package customerr

type SerializedError interface {
	Error() string
	SerializeErrors() []string
}
