package context

type Store interface {
	Cancel()
	Fetch() string
}
