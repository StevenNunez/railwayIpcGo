package payload_test

import (
	payload "railwayIpc/railwayIpc"
	railway_ipc_sample "railwayIpc/support/protos"
	"testing"
)

func TestEncoding(t *testing.T) {
	messageUUID := "07a9a4ba-778d-11ea-ae8c-10e7c6097871"
	userUUID := "07a9a4b7-778d-11ea-ae8c-10e7c6097871"
	correlationId := "07a9a3dc-778d-11ea-ae8c-10e7c6097871"

	data := railway_ipc_sample.AddRequest_Data{
		Op1: 1,
		Op2: 2,
	}
	message := railway_ipc_sample.AddRequest{
		UserUuid:      userUUID,
		CorrelationId: correlationId,
		Uuid:          messageUUID,
		Context:       map[string]string{"long-term": "value"},
		Data:          &data,
	}
	result := payload.Encode(&message)

	expected := `{"encoded_message":"CiQwN2E5YTRiNy03NzhkLTExZWEtYWU4Yy0xMGU3YzYwOTc4NzESJDA3YTlhM2RjLTc3OGQtMTFlYS1hZThjLTEwZTdjNjA5Nzg3MRokMDdhOWE0YmEtNzc4ZC0xMWVhLWFlOGMtMTBlN2M2MDk3ODcxIhIKCWxvbmctdGVybRIFdmFsdWUqBAgBEAI=","type":"RailwayIpc::Sample::AddRequest"}`

	if result != expected  {
		t.Fatalf("expected %s to equal %s", result, expected)
	}
}

func TestDecoding(t *testing.T) {
	//messageUUID := "07a9a4ba-778d-11ea-ae8c-10e7c6097871"
	//userUUID := "07a9a4b7-778d-11ea-ae8c-10e7c6097871"
	//correlationId := "07a9a3dc-778d-11ea-ae8c-10e7c6097871"
	//
	//data := railway_ipc_sample.AddRequest_Data{
	//	Op1: 1,
	//	Op2: 2,
	//}
	//message := railway_ipc_sample.AddRequest{
	//	UserUuid:      userUUID,
	//	CorrelationId: correlationId,
	//	Uuid:          messageUUID,
	//	Context:       map[string]string{"long-term": "value"},
	//	Data:          &data,
	//}
	//result := payload.Encode(&message)

	payload_string := `{"encoded_message":"CiQwN2E5YTRiNy03NzhkLTExZWEtYWU4Yy0xMGU3YzYwOTc4NzESJDA3YTlhM2RjLTc3OGQtMTFlYS1hZThjLTEwZTdjNjA5Nzg3MRokMDdhOWE0YmEtNzc4ZC0xMWVhLWFlOGMtMTBlN2M2MDk3ODcxIhIKCWxvbmctdGVybRIFdmFsdWUqBAgBEAI=","type":"RailwayIpc::Sample::AddRequest"}`
	result := payload.Decode(payload_string)
	if result == nil {
		t.Fatalf("can't %v", result)
	}
}
