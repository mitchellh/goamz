package sns

type Message struct {
	Topic   *Topic
	Message [8192]byte
	Subject string
}
