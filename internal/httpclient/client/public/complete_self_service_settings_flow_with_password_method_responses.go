// Code generated by go-swagger; DO NOT EDIT.

package public

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ory/kratos/internal/httpclient/models"
)

// CompleteSelfServiceSettingsFlowWithPasswordMethodReader is a Reader for the CompleteSelfServiceSettingsFlowWithPasswordMethod structure.
type CompleteSelfServiceSettingsFlowWithPasswordMethodReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CompleteSelfServiceSettingsFlowWithPasswordMethodReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 302:
		result := NewCompleteSelfServiceSettingsFlowWithPasswordMethodFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCompleteSelfServiceSettingsFlowWithPasswordMethodInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCompleteSelfServiceSettingsFlowWithPasswordMethodFound creates a CompleteSelfServiceSettingsFlowWithPasswordMethodFound with default headers values
func NewCompleteSelfServiceSettingsFlowWithPasswordMethodFound() *CompleteSelfServiceSettingsFlowWithPasswordMethodFound {
	return &CompleteSelfServiceSettingsFlowWithPasswordMethodFound{}
}

/*CompleteSelfServiceSettingsFlowWithPasswordMethodFound handles this case with default header values.

Empty responses are sent when, for example, resources are deleted. The HTTP status code for empty responses is
typically 201.
*/
type CompleteSelfServiceSettingsFlowWithPasswordMethodFound struct {
}

func (o *CompleteSelfServiceSettingsFlowWithPasswordMethodFound) Error() string {
	return fmt.Sprintf("[POST /self-service/settings/methods/password][%d] completeSelfServiceSettingsFlowWithPasswordMethodFound ", 302)
}

func (o *CompleteSelfServiceSettingsFlowWithPasswordMethodFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCompleteSelfServiceSettingsFlowWithPasswordMethodInternalServerError creates a CompleteSelfServiceSettingsFlowWithPasswordMethodInternalServerError with default headers values
func NewCompleteSelfServiceSettingsFlowWithPasswordMethodInternalServerError() *CompleteSelfServiceSettingsFlowWithPasswordMethodInternalServerError {
	return &CompleteSelfServiceSettingsFlowWithPasswordMethodInternalServerError{}
}

/*CompleteSelfServiceSettingsFlowWithPasswordMethodInternalServerError handles this case with default header values.

genericError
*/
type CompleteSelfServiceSettingsFlowWithPasswordMethodInternalServerError struct {
	Payload *models.GenericError
}

func (o *CompleteSelfServiceSettingsFlowWithPasswordMethodInternalServerError) Error() string {
	return fmt.Sprintf("[POST /self-service/settings/methods/password][%d] completeSelfServiceSettingsFlowWithPasswordMethodInternalServerError  %+v", 500, o.Payload)
}

func (o *CompleteSelfServiceSettingsFlowWithPasswordMethodInternalServerError) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *CompleteSelfServiceSettingsFlowWithPasswordMethodInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}