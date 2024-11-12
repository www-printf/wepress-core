package errors

type Error struct {
	LogError  error
	HTTPError error
}
