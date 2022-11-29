// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/example/example.proto

package example

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "github.com/gogo/protobuf/gogoproto"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *AddRequest) Validate() error {
	if this.Name == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must not be an empty string`, this.Name))
	}
	if !(len(this.Name) < 75) {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must have a length smaller than '75'`, this.Name))
	}
	return nil
}
func (this *GetRequest) Validate() error {
	return nil
}
func (this *CommonResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *Empty) Validate() error {
	return nil
}
