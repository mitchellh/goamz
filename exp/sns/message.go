package sns

type Message struct {
	Topic   *Topic
	Message string
	Subject string
}
