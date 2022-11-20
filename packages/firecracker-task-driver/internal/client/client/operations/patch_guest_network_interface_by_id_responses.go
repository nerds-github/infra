// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/devbookhq/packages/firecracker-task-driver/internal/client/models"
)

// PatchGuestNetworkInterfaceByIDReader is a Reader for the PatchGuestNetworkInterfaceByID structure.
type PatchGuestNetworkInterfaceByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchGuestNetworkInterfaceByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPatchGuestNetworkInterfaceByIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchGuestNetworkInterfaceByIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPatchGuestNetworkInterfaceByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchGuestNetworkInterfaceByIDNoContent creates a PatchGuestNetworkInterfaceByIDNoContent with default headers values
func NewPatchGuestNetworkInterfaceByIDNoContent() *PatchGuestNetworkInterfaceByIDNoContent {
	return &PatchGuestNetworkInterfaceByIDNoContent{}
}

/* PatchGuestNetworkInterfaceByIDNoContent describes a response with status code 204, with default header values.

Network interface updated
*/
type PatchGuestNetworkInterfaceByIDNoContent struct {
}

func (o *PatchGuestNetworkInterfaceByIDNoContent) Error() string {
	return fmt.Sprintf("[PATCH /network-interfaces/{iface_id}][%d] patchGuestNetworkInterfaceByIdNoContent ", 204)
}

func (o *PatchGuestNetworkInterfaceByIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchGuestNetworkInterfaceByIDBadRequest creates a PatchGuestNetworkInterfaceByIDBadRequest with default headers values
func NewPatchGuestNetworkInterfaceByIDBadRequest() *PatchGuestNetworkInterfaceByIDBadRequest {
	return &PatchGuestNetworkInterfaceByIDBadRequest{}
}

/* PatchGuestNetworkInterfaceByIDBadRequest describes a response with status code 400, with default header values.

Network interface cannot be updated due to bad input
*/
type PatchGuestNetworkInterfaceByIDBadRequest struct {
	Payload *models.Error
}

func (o *PatchGuestNetworkInterfaceByIDBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /network-interfaces/{iface_id}][%d] patchGuestNetworkInterfaceByIdBadRequest  %+v", 400, o.Payload)
}
func (o *PatchGuestNetworkInterfaceByIDBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PatchGuestNetworkInterfaceByIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchGuestNetworkInterfaceByIDDefault creates a PatchGuestNetworkInterfaceByIDDefault with default headers values
func NewPatchGuestNetworkInterfaceByIDDefault(code int) *PatchGuestNetworkInterfaceByIDDefault {
	return &PatchGuestNetworkInterfaceByIDDefault{
		_statusCode: code,
	}
}

/* PatchGuestNetworkInterfaceByIDDefault describes a response with status code -1, with default header values.

Internal server error
*/
type PatchGuestNetworkInterfaceByIDDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the patch guest network interface by ID default response
func (o *PatchGuestNetworkInterfaceByIDDefault) Code() int {
	return o._statusCode
}

func (o *PatchGuestNetworkInterfaceByIDDefault) Error() string {
	return fmt.Sprintf("[PATCH /network-interfaces/{iface_id}][%d] patchGuestNetworkInterfaceByID default  %+v", o._statusCode, o.Payload)
}
func (o *PatchGuestNetworkInterfaceByIDDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PatchGuestNetworkInterfaceByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
