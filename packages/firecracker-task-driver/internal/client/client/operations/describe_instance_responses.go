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

// DescribeInstanceReader is a Reader for the DescribeInstance structure.
type DescribeInstanceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DescribeInstanceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDescribeInstanceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDescribeInstanceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDescribeInstanceOK creates a DescribeInstanceOK with default headers values
func NewDescribeInstanceOK() *DescribeInstanceOK {
	return &DescribeInstanceOK{}
}

/* DescribeInstanceOK describes a response with status code 200, with default header values.

The instance information
*/
type DescribeInstanceOK struct {
	Payload *models.InstanceInfo
}

func (o *DescribeInstanceOK) Error() string {
	return fmt.Sprintf("[GET /][%d] describeInstanceOK  %+v", 200, o.Payload)
}
func (o *DescribeInstanceOK) GetPayload() *models.InstanceInfo {
	return o.Payload
}

func (o *DescribeInstanceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InstanceInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDescribeInstanceDefault creates a DescribeInstanceDefault with default headers values
func NewDescribeInstanceDefault(code int) *DescribeInstanceDefault {
	return &DescribeInstanceDefault{
		_statusCode: code,
	}
}

/* DescribeInstanceDefault describes a response with status code -1, with default header values.

Internal Server Error
*/
type DescribeInstanceDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the describe instance default response
func (o *DescribeInstanceDefault) Code() int {
	return o._statusCode
}

func (o *DescribeInstanceDefault) Error() string {
	return fmt.Sprintf("[GET /][%d] describeInstance default  %+v", o._statusCode, o.Payload)
}
func (o *DescribeInstanceDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DescribeInstanceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
