package payload

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/generator"
	"log"
	"railwayIpc"
	railway_ipc_sample "railwayIpc/support/protos"
	"regexp"
	"strings"
)

func Encode(message railwayIpc.Message) string {
	encodedMessage, err := proto.Marshal(message)
	if err != nil {
		log.Fatalf("Failed to encode %v.", message)
	}
	base64encoded := base64.StdEncoding.EncodeToString([]byte(encodedMessage))

	encodedMap := map[string]string{
		"type":            externalModuleName(message),
		"encoded_message": base64encoded,
	}

	result, err := json.Marshal(encodedMap)

	if err != nil {
		log.Fatalf("Failed to JSON encode %v.", message)
	}
	return string(result)
}

func Decode(payload string) railwayIpc.Message {
	encodedMap := map[string]string{
		"type":            "",
		"encoded_message": "",
	}
	err := json.Unmarshal([]byte(payload), &encodedMap)
	if err != nil {
		log.Fatalf("Failed to JSON decode %s", payload)
	}
	internalModuleName(encodedMap["type"])
	return &railway_ipc_sample.AddRequest{}
}

func externalModuleName(message railwayIpc.Message) string {
	rawModuleName := proto.MessageName(message)
	moduleSegment := strings.Split(rawModuleName, ".")
	for i, segment := range moduleSegment {
		moduleSegment[i] = generator.CamelCase(segment)
	}
	return strings.Join(moduleSegment, "::")
}

func internalModuleName(messageType string) railwayIpc.Message {
	splitType := strings.Split(messageType, "::")
	publicStructPos := len(splitType)-1
	needConversion := splitType[:publicStructPos]
	publicStruct := splitType[publicStructPos:]

	for i, segment := range needConversion {
		needConversion[i] = toSnakeCase(segment)
	}
	structName := strings.Join(append(needConversion, publicStruct...), ".")
	fmt.Println(structName)
	 proto.MessageType(structName)
	 h := proto.MessageType("railway_ipc.sample.AddRequest")
	 fmt.Printf("What are you? %+v\n", h.)
	return &railway_ipc_sample.AddRequest{}
}


var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")
func toSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}
