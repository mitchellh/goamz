package sns

type Message struct {
	SNS     *SNS
	Topic   *Topic
	Message [8192]byte
	Subject string
}
