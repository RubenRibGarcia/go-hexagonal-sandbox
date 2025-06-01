package eventbus

type Publisher interface {
	Publish(v any) error
}
