package railwayIpc

import "github.com/golang/protobuf/proto"

type Handler interface {
	handleRequest(message proto.Message)
}

type Message interface {
	GetUserUuid() string
	GetCorrelationId() string
	GetUuid() string
	GetContext() map[string]string
	Reset()
	String() string
	ProtoMessage()
}
type ConsumerSpec struct {
	exchange string
	queue    string
	handlers map[proto.Message]Handler
}

//Payload work
//  deps will need protobuf library
//  Will need some protobufs for testing
//  Decode message from wire
//  Json decode {type:..., encoded_message:...}
//  Decode message with type
//  encoded protobuf message into json format
//Rabbit work
//  Connect to connection, channel
//  declare fanout exchange, and queue
//Consumer patterns
//  type ConsumerSpec {
//    exchange string
//    queue string
//    handlers map[Protobuf]func
//  }
//  go startConsumer(Some spec)
//Persistence work
//  save messages that come in
//  look them up by uuid, skip if handled
//*/
