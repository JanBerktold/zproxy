package main

func NewConnectionHandler(sender Sender, receiver Receiver) (<-chan ZkRequest, chan<- ZkResponse) {
	requestChannel := make(chan ZkRequest)

	go func() {
		for true {
			req := <-requestChannel
			sender.Send(req)
		}
	}()

	return requestChannel, nil
}

type Sender interface {
	Send(request ZkRequest) error
}

type Receiver interface {
	Receive() (ZkResponse, error)
}

// Refer to https://github.com/apache/zookeeper/blob/master/src/zookeeper.jute
type ZkRequest interface {

}

type ZkResponse interface {

}
