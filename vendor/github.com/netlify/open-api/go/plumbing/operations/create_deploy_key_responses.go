// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netlify/open-api/go/models"
)

// CreateDeployKeyReader is a Reader for the CreateDeployKey structure.
type CreateDeployKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateDeployKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateDeployKeyCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateDeployKeyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateDeployKeyCreated creates a CreateDeployKeyCreated with default headers values
func NewCreateDeployKeyCreated() *CreateDeployKeyCreated {
	return &CreateDeployKeyCreated{}
}

/*CreateDeployKeyCreated handles this case with default header values.

Created
*/
type CreateDeployKeyCreated struct {
	Payload *models.DeployKey
}

func (o *CreateDeployKeyCreated) Error() string {
	return fmt.Sprintf("[POST /deploy_keys][%d] createDeployKeyCreated  %+v", 201, o.Payload)
}

func (o *CreateDeployKeyCreated) GetPayload() *models.DeployKey {
	return o.Payload
}

func (o *CreateDeployKeyCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeployKey)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDeployKeyDefault creates a CreateDeployKeyDefault with default headers values
func NewCreateDeployKeyDefault(code int) *CreateDeployKeyDefault {
	return &CreateDeployKeyDefault{
		_statusCode: code,
	}
}

/*CreateDeployKeyDefault handles this case with default header values.

error
*/
type CreateDeployKeyDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the create deploy key default response
func (o *CreateDeployKeyDefault) Code() int {
	return o._statusCode
}

func (o *CreateDeployKeyDefault) Error() string {
	return fmt.Sprintf("[POST /deploy_keys][%d] createDeployKey default  %+v", o._statusCode, o.Payload)
}

func (o *CreateDeployKeyDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateDeployKeyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}