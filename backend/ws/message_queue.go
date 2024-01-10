package ws

type MessageQueue struct {
	queue chan []byte
}

func NewMessageQueue(size int) *MessageQueue {
	return &MessageQueue{
		queue: make(chan []byte, size),
	}
}

func (mq *MessageQueue) Enqueue(message []byte) {
	mq.queue <- message
}

func (mq *MessageQueue) StartProcessing(hub *Hub) {
	go func() {
		for message := range mq.queue {
			hub.broadcast <- message
			// Additional logic for delay or throttling will be added here
		}
	}()
}
