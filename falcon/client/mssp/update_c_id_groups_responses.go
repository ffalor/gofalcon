// Code generated by go-swagger; DO NOT EDIT.

package mssp

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/crowdstrike/gofalcon/falcon/models"
)

// UpdateCIDGroupsReader is a Reader for the UpdateCIDGroups structure.
type UpdateCIDGroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateCIDGroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateCIDGroupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 207:
		result := NewUpdateCIDGroupsMultiStatus()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateCIDGroupsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateCIDGroupsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewUpdateCIDGroupsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdateCIDGroupsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateCIDGroupsOK creates a UpdateCIDGroupsOK with default headers values
func NewUpdateCIDGroupsOK() *UpdateCIDGroupsOK {
	return &UpdateCIDGroupsOK{}
}

/* UpdateCIDGroupsOK describes a response with status code 200, with default header values.

OK
*/
type UpdateCIDGroupsOK struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainCIDGroupsResponseV1
}

func (o *UpdateCIDGroupsOK) Error() string {
	return fmt.Sprintf("[PATCH /mssp/entities/cid-groups/v1][%d] updateCIdGroupsOK  %+v", 200, o.Payload)
}
func (o *UpdateCIDGroupsOK) GetPayload() *models.DomainCIDGroupsResponseV1 {
	return o.Payload
}

func (o *UpdateCIDGroupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.DomainCIDGroupsResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCIDGroupsMultiStatus creates a UpdateCIDGroupsMultiStatus with default headers values
func NewUpdateCIDGroupsMultiStatus() *UpdateCIDGroupsMultiStatus {
	return &UpdateCIDGroupsMultiStatus{}
}

/* UpdateCIDGroupsMultiStatus describes a response with status code 207, with default header values.

Multi-Status
*/
type UpdateCIDGroupsMultiStatus struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainCIDGroupsResponseV1
}

func (o *UpdateCIDGroupsMultiStatus) Error() string {
	return fmt.Sprintf("[PATCH /mssp/entities/cid-groups/v1][%d] updateCIdGroupsMultiStatus  %+v", 207, o.Payload)
}
func (o *UpdateCIDGroupsMultiStatus) GetPayload() *models.DomainCIDGroupsResponseV1 {
	return o.Payload
}

func (o *UpdateCIDGroupsMultiStatus) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.DomainCIDGroupsResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCIDGroupsBadRequest creates a UpdateCIDGroupsBadRequest with default headers values
func NewUpdateCIDGroupsBadRequest() *UpdateCIDGroupsBadRequest {
	return &UpdateCIDGroupsBadRequest{}
}

/* UpdateCIDGroupsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UpdateCIDGroupsBadRequest struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaErrorsOnly
}

func (o *UpdateCIDGroupsBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /mssp/entities/cid-groups/v1][%d] updateCIdGroupsBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateCIDGroupsBadRequest) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *UpdateCIDGroupsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.MsaErrorsOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCIDGroupsForbidden creates a UpdateCIDGroupsForbidden with default headers values
func NewUpdateCIDGroupsForbidden() *UpdateCIDGroupsForbidden {
	return &UpdateCIDGroupsForbidden{}
}

/* UpdateCIDGroupsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UpdateCIDGroupsForbidden struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaErrorsOnly
}

func (o *UpdateCIDGroupsForbidden) Error() string {
	return fmt.Sprintf("[PATCH /mssp/entities/cid-groups/v1][%d] updateCIdGroupsForbidden  %+v", 403, o.Payload)
}
func (o *UpdateCIDGroupsForbidden) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *UpdateCIDGroupsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.MsaErrorsOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCIDGroupsTooManyRequests creates a UpdateCIDGroupsTooManyRequests with default headers values
func NewUpdateCIDGroupsTooManyRequests() *UpdateCIDGroupsTooManyRequests {
	return &UpdateCIDGroupsTooManyRequests{}
}

/* UpdateCIDGroupsTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type UpdateCIDGroupsTooManyRequests struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	/* Too many requests, retry after this time (as milliseconds since epoch)
	 */
	XRateLimitRetryAfter int64

	Payload *models.MsaReplyMetaOnly
}

func (o *UpdateCIDGroupsTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /mssp/entities/cid-groups/v1][%d] updateCIdGroupsTooManyRequests  %+v", 429, o.Payload)
}
func (o *UpdateCIDGroupsTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *UpdateCIDGroupsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	// hydrates response header X-RateLimit-RetryAfter
	hdrXRateLimitRetryAfter := response.GetHeader("X-RateLimit-RetryAfter")

	if hdrXRateLimitRetryAfter != "" {
		valxRateLimitRetryAfter, err := swag.ConvertInt64(hdrXRateLimitRetryAfter)
		if err != nil {
			return errors.InvalidType("X-RateLimit-RetryAfter", "header", "int64", hdrXRateLimitRetryAfter)
		}
		o.XRateLimitRetryAfter = valxRateLimitRetryAfter
	}

	o.Payload = new(models.MsaReplyMetaOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCIDGroupsDefault creates a UpdateCIDGroupsDefault with default headers values
func NewUpdateCIDGroupsDefault(code int) *UpdateCIDGroupsDefault {
	return &UpdateCIDGroupsDefault{
		_statusCode: code,
	}
}

/* UpdateCIDGroupsDefault describes a response with status code -1, with default header values.

OK
*/
type UpdateCIDGroupsDefault struct {
	_statusCode int

	Payload *models.DomainCIDGroupsResponseV1
}

// Code gets the status code for the update c ID groups default response
func (o *UpdateCIDGroupsDefault) Code() int {
	return o._statusCode
}

func (o *UpdateCIDGroupsDefault) Error() string {
	return fmt.Sprintf("[PATCH /mssp/entities/cid-groups/v1][%d] updateCIDGroups default  %+v", o._statusCode, o.Payload)
}
func (o *UpdateCIDGroupsDefault) GetPayload() *models.DomainCIDGroupsResponseV1 {
	return o.Payload
}

func (o *UpdateCIDGroupsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DomainCIDGroupsResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}