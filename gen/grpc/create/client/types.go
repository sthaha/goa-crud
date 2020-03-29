// Code generated by goa v3.1.1, DO NOT EDIT.
//
// create gRPC client types
//
// Command:
// $ goa gen crud/design

package client

import (
	create "crud/gen/create"
	createpb "crud/gen/grpc/create/pb"
	"unicode/utf8"

	goa "goa.design/goa/v3/pkg"
)

// NewCreateRequest builds the gRPC request type from the payload of the
// "create" endpoint of the "create" service.
func NewCreateRequest(payload *create.Blog) *createpb.CreateRequest {
	message := &createpb.CreateRequest{}
	if payload.ID != nil {
		message.Id = *payload.ID
	}
	if payload.Name != nil {
		message.Name = *payload.Name
	}
	if payload.Comments != nil {
		message.Comments = make([]string, len(payload.Comments))
		for i, val := range payload.Comments {
			message.Comments[i] = val
		}
	}
	return message
}

// NewCreateResult builds the result type of the "create" endpoint of the
// "create" service from the gRPC response type.
func NewCreateResult(message *createpb.CreateResponse) *create.Blog {
	result := &create.Blog{}
	if message.Id != 0 {
		result.ID = &message.Id
	}
	if message.Name != "" {
		result.Name = &message.Name
	}
	if message.Comments != nil {
		result.Comments = make([]string, len(message.Comments))
		for i, val := range message.Comments {
			result.Comments[i] = val
		}
	}
	return result
}

// NewListRequest builds the gRPC request type from the payload of the "list"
// endpoint of the "create" service.
func NewListRequest() *createpb.ListRequest {
	message := &createpb.ListRequest{}
	return message
}

// NewListResult builds the result type of the "list" endpoint of the "create"
// service from the gRPC response type.
func NewListResult(message *createpb.ListResponse) []*create.Storedblog {
	result := make([]*create.Storedblog, len(message.Field))
	for i, val := range message.Field {
		result[i] = &create.Storedblog{
			ID:   val.Id,
			Name: val.Name,
		}
		if val.Comments != nil {
			result[i].Comments = make([]string, len(val.Comments))
			for j, val := range val.Comments {
				result[i].Comments[j] = val
			}
		}
	}
	return result
}

// NewRemoveRequest builds the gRPC request type from the payload of the
// "remove" endpoint of the "create" service.
func NewRemoveRequest(payload *create.RemovePayload) *createpb.RemoveRequest {
	message := &createpb.RemoveRequest{
		Id: payload.ID,
	}
	return message
}

// NewUpdateRequest builds the gRPC request type from the payload of the
// "update" endpoint of the "create" service.
func NewUpdateRequest(payload *create.UpdatePayload) *createpb.UpdateRequest {
	message := &createpb.UpdateRequest{
		Name: payload.Name,
	}
	if payload.ID != nil {
		message.Id = *payload.ID
	}
	if payload.Comments != nil {
		message.Comments = make([]string, len(payload.Comments))
		for i, val := range payload.Comments {
			message.Comments[i] = val
		}
	}
	return message
}

// ValidateCreateResponse runs the validations defined on CreateResponse.
func ValidateCreateResponse(message *createpb.CreateResponse) (err error) {
	if message.Name != "" {
		if utf8.RuneCountInString(message.Name) > 100 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("message.name", message.Name, utf8.RuneCountInString(message.Name), 100, false))
		}
	}
	if len(message.Comments) > 100 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("message.comments", message.Comments, len(message.Comments), 100, false))
	}
	return
}

// ValidateListResponse runs the validations defined on ListResponse.
func ValidateListResponse(message *createpb.ListResponse) (err error) {
	for _, e := range message.Field {
		if e != nil {
			if err2 := ValidateStoredblog(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateStoredblog runs the validations defined on Storedblog.
func ValidateStoredblog(message *createpb.Storedblog) (err error) {
	if utf8.RuneCountInString(message.Name) > 100 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("message.name", message.Name, utf8.RuneCountInString(message.Name), 100, false))
	}
	if len(message.Comments) > 100 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("message.comments", message.Comments, len(message.Comments), 100, false))
	}
	return
}