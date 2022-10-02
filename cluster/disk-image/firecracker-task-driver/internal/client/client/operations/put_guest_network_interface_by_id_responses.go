// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/devbookhq/firecracker-task-driver/internal/client/models"
)

// PutGuestNetworkInterfaceByIDReader is a Reader for the PutGuestNetworkInterfaceByID structure.
type PutGuestNetworkInterfaceByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutGuestNetworkInterfaceByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutGuestNetworkInterfaceByIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutGuestNetworkInterfaceByIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPutGuestNetworkInterfaceByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutGuestNetworkInterfaceByIDNoContent creates a PutGuestNetworkInterfaceByIDNoContent with default headers values
func NewPutGuestNetworkInterfaceByIDNoContent() *PutGuestNetworkInterfaceByIDNoContent {
	return &PutGuestNetworkInterfaceByIDNoContent{}
}

/* PutGuestNetworkInterfaceByIDNoContent describes a response with status code 204, with default header values.

Network interface created/updated
*/
type PutGuestNetworkInterfaceByIDNoContent struct {
}

func (o *PutGuestNetworkInterfaceByIDNoContent) Error() string {
	return fmt.Sprintf("[PUT /network-interfaces/{iface_id}][%d] putGuestNetworkInterfaceByIdNoContent ", 204)
}

func (o *PutGuestNetworkInterfaceByIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutGuestNetworkInterfaceByIDBadRequest creates a PutGuestNetworkInterfaceByIDBadRequest with default headers values
func NewPutGuestNetworkInterfaceByIDBadRequest() *PutGuestNetworkInterfaceByIDBadRequest {
	return &PutGuestNetworkInterfaceByIDBadRequest{}
}

/* PutGuestNetworkInterfaceByIDBadRequest describes a response with status code 400, with default header values.

Network interface cannot be created due to bad input
*/
type PutGuestNetworkInterfaceByIDBadRequest struct {
	Payload *models.Error
}

func (o *PutGuestNetworkInterfaceByIDBadRequest) Error() string {
	return fmt.Sprintf("[PUT /network-interfaces/{iface_id}][%d] putGuestNetworkInterfaceByIdBadRequest  %+v", 400, o.Payload)
}
func (o *PutGuestNetworkInterfaceByIDBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutGuestNetworkInterfaceByIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutGuestNetworkInterfaceByIDDefault creates a PutGuestNetworkInterfaceByIDDefault with default headers values
func NewPutGuestNetworkInterfaceByIDDefault(code int) *PutGuestNetworkInterfaceByIDDefault {
	return &PutGuestNetworkInterfaceByIDDefault{
		_statusCode: code,
	}
}

/* PutGuestNetworkInterfaceByIDDefault describes a response with status code -1, with default header values.

Internal server error
*/
type PutGuestNetworkInterfaceByIDDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the put guest network interface by ID default response
func (o *PutGuestNetworkInterfaceByIDDefault) Code() int {
	return o._statusCode
}

func (o *PutGuestNetworkInterfaceByIDDefault) Error() string {
	return fmt.Sprintf("[PUT /network-interfaces/{iface_id}][%d] putGuestNetworkInterfaceByID default  %+v", o._statusCode, o.Payload)
}
func (o *PutGuestNetworkInterfaceByIDDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutGuestNetworkInterfaceByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}