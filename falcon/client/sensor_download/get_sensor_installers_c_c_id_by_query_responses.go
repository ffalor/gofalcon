// Code generated by go-swagger; DO NOT EDIT.

package sensor_download

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

// GetSensorInstallersCCIDByQueryReader is a Reader for the GetSensorInstallersCCIDByQuery structure.
type GetSensorInstallersCCIDByQueryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSensorInstallersCCIDByQueryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSensorInstallersCCIDByQueryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetSensorInstallersCCIDByQueryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetSensorInstallersCCIDByQueryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetSensorInstallersCCIDByQueryTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSensorInstallersCCIDByQueryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /sensors/queries/installers/ccid/v1] GetSensorInstallersCCIDByQuery", response, response.Code())
	}
}

// NewGetSensorInstallersCCIDByQueryOK creates a GetSensorInstallersCCIDByQueryOK with default headers values
func NewGetSensorInstallersCCIDByQueryOK() *GetSensorInstallersCCIDByQueryOK {
	return &GetSensorInstallersCCIDByQueryOK{}
}

/*
GetSensorInstallersCCIDByQueryOK describes a response with status code 200, with default header values.

OK
*/
type GetSensorInstallersCCIDByQueryOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaspecQueryResponse
}

// IsSuccess returns true when this get sensor installers c c Id by query o k response has a 2xx status code
func (o *GetSensorInstallersCCIDByQueryOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get sensor installers c c Id by query o k response has a 3xx status code
func (o *GetSensorInstallersCCIDByQueryOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get sensor installers c c Id by query o k response has a 4xx status code
func (o *GetSensorInstallersCCIDByQueryOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get sensor installers c c Id by query o k response has a 5xx status code
func (o *GetSensorInstallersCCIDByQueryOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get sensor installers c c Id by query o k response a status code equal to that given
func (o *GetSensorInstallersCCIDByQueryOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get sensor installers c c Id by query o k response
func (o *GetSensorInstallersCCIDByQueryOK) Code() int {
	return 200
}

func (o *GetSensorInstallersCCIDByQueryOK) Error() string {
	return fmt.Sprintf("[GET /sensors/queries/installers/ccid/v1][%d] getSensorInstallersCCIdByQueryOK  %+v", 200, o.Payload)
}

func (o *GetSensorInstallersCCIDByQueryOK) String() string {
	return fmt.Sprintf("[GET /sensors/queries/installers/ccid/v1][%d] getSensorInstallersCCIdByQueryOK  %+v", 200, o.Payload)
}

func (o *GetSensorInstallersCCIDByQueryOK) GetPayload() *models.MsaspecQueryResponse {
	return o.Payload
}

func (o *GetSensorInstallersCCIDByQueryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

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

	o.Payload = new(models.MsaspecQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSensorInstallersCCIDByQueryBadRequest creates a GetSensorInstallersCCIDByQueryBadRequest with default headers values
func NewGetSensorInstallersCCIDByQueryBadRequest() *GetSensorInstallersCCIDByQueryBadRequest {
	return &GetSensorInstallersCCIDByQueryBadRequest{}
}

/*
GetSensorInstallersCCIDByQueryBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetSensorInstallersCCIDByQueryBadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaspecQueryResponse
}

// IsSuccess returns true when this get sensor installers c c Id by query bad request response has a 2xx status code
func (o *GetSensorInstallersCCIDByQueryBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get sensor installers c c Id by query bad request response has a 3xx status code
func (o *GetSensorInstallersCCIDByQueryBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get sensor installers c c Id by query bad request response has a 4xx status code
func (o *GetSensorInstallersCCIDByQueryBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get sensor installers c c Id by query bad request response has a 5xx status code
func (o *GetSensorInstallersCCIDByQueryBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get sensor installers c c Id by query bad request response a status code equal to that given
func (o *GetSensorInstallersCCIDByQueryBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get sensor installers c c Id by query bad request response
func (o *GetSensorInstallersCCIDByQueryBadRequest) Code() int {
	return 400
}

func (o *GetSensorInstallersCCIDByQueryBadRequest) Error() string {
	return fmt.Sprintf("[GET /sensors/queries/installers/ccid/v1][%d] getSensorInstallersCCIdByQueryBadRequest  %+v", 400, o.Payload)
}

func (o *GetSensorInstallersCCIDByQueryBadRequest) String() string {
	return fmt.Sprintf("[GET /sensors/queries/installers/ccid/v1][%d] getSensorInstallersCCIdByQueryBadRequest  %+v", 400, o.Payload)
}

func (o *GetSensorInstallersCCIDByQueryBadRequest) GetPayload() *models.MsaspecQueryResponse {
	return o.Payload
}

func (o *GetSensorInstallersCCIDByQueryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

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

	o.Payload = new(models.MsaspecQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSensorInstallersCCIDByQueryForbidden creates a GetSensorInstallersCCIDByQueryForbidden with default headers values
func NewGetSensorInstallersCCIDByQueryForbidden() *GetSensorInstallersCCIDByQueryForbidden {
	return &GetSensorInstallersCCIDByQueryForbidden{}
}

/*
GetSensorInstallersCCIDByQueryForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetSensorInstallersCCIDByQueryForbidden struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

// IsSuccess returns true when this get sensor installers c c Id by query forbidden response has a 2xx status code
func (o *GetSensorInstallersCCIDByQueryForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get sensor installers c c Id by query forbidden response has a 3xx status code
func (o *GetSensorInstallersCCIDByQueryForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get sensor installers c c Id by query forbidden response has a 4xx status code
func (o *GetSensorInstallersCCIDByQueryForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get sensor installers c c Id by query forbidden response has a 5xx status code
func (o *GetSensorInstallersCCIDByQueryForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get sensor installers c c Id by query forbidden response a status code equal to that given
func (o *GetSensorInstallersCCIDByQueryForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get sensor installers c c Id by query forbidden response
func (o *GetSensorInstallersCCIDByQueryForbidden) Code() int {
	return 403
}

func (o *GetSensorInstallersCCIDByQueryForbidden) Error() string {
	return fmt.Sprintf("[GET /sensors/queries/installers/ccid/v1][%d] getSensorInstallersCCIdByQueryForbidden  %+v", 403, o.Payload)
}

func (o *GetSensorInstallersCCIDByQueryForbidden) String() string {
	return fmt.Sprintf("[GET /sensors/queries/installers/ccid/v1][%d] getSensorInstallersCCIdByQueryForbidden  %+v", 403, o.Payload)
}

func (o *GetSensorInstallersCCIDByQueryForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetSensorInstallersCCIDByQueryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

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

	o.Payload = new(models.MsaReplyMetaOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSensorInstallersCCIDByQueryTooManyRequests creates a GetSensorInstallersCCIDByQueryTooManyRequests with default headers values
func NewGetSensorInstallersCCIDByQueryTooManyRequests() *GetSensorInstallersCCIDByQueryTooManyRequests {
	return &GetSensorInstallersCCIDByQueryTooManyRequests{}
}

/*
GetSensorInstallersCCIDByQueryTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetSensorInstallersCCIDByQueryTooManyRequests struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

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

// IsSuccess returns true when this get sensor installers c c Id by query too many requests response has a 2xx status code
func (o *GetSensorInstallersCCIDByQueryTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get sensor installers c c Id by query too many requests response has a 3xx status code
func (o *GetSensorInstallersCCIDByQueryTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get sensor installers c c Id by query too many requests response has a 4xx status code
func (o *GetSensorInstallersCCIDByQueryTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get sensor installers c c Id by query too many requests response has a 5xx status code
func (o *GetSensorInstallersCCIDByQueryTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get sensor installers c c Id by query too many requests response a status code equal to that given
func (o *GetSensorInstallersCCIDByQueryTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the get sensor installers c c Id by query too many requests response
func (o *GetSensorInstallersCCIDByQueryTooManyRequests) Code() int {
	return 429
}

func (o *GetSensorInstallersCCIDByQueryTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /sensors/queries/installers/ccid/v1][%d] getSensorInstallersCCIdByQueryTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetSensorInstallersCCIDByQueryTooManyRequests) String() string {
	return fmt.Sprintf("[GET /sensors/queries/installers/ccid/v1][%d] getSensorInstallersCCIdByQueryTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetSensorInstallersCCIDByQueryTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetSensorInstallersCCIDByQueryTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

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

// NewGetSensorInstallersCCIDByQueryInternalServerError creates a GetSensorInstallersCCIDByQueryInternalServerError with default headers values
func NewGetSensorInstallersCCIDByQueryInternalServerError() *GetSensorInstallersCCIDByQueryInternalServerError {
	return &GetSensorInstallersCCIDByQueryInternalServerError{}
}

/*
GetSensorInstallersCCIDByQueryInternalServerError describes a response with status code 500, with default header values.

Unexpected Error
*/
type GetSensorInstallersCCIDByQueryInternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

// IsSuccess returns true when this get sensor installers c c Id by query internal server error response has a 2xx status code
func (o *GetSensorInstallersCCIDByQueryInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get sensor installers c c Id by query internal server error response has a 3xx status code
func (o *GetSensorInstallersCCIDByQueryInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get sensor installers c c Id by query internal server error response has a 4xx status code
func (o *GetSensorInstallersCCIDByQueryInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get sensor installers c c Id by query internal server error response has a 5xx status code
func (o *GetSensorInstallersCCIDByQueryInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get sensor installers c c Id by query internal server error response a status code equal to that given
func (o *GetSensorInstallersCCIDByQueryInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get sensor installers c c Id by query internal server error response
func (o *GetSensorInstallersCCIDByQueryInternalServerError) Code() int {
	return 500
}

func (o *GetSensorInstallersCCIDByQueryInternalServerError) Error() string {
	return fmt.Sprintf("[GET /sensors/queries/installers/ccid/v1][%d] getSensorInstallersCCIdByQueryInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSensorInstallersCCIDByQueryInternalServerError) String() string {
	return fmt.Sprintf("[GET /sensors/queries/installers/ccid/v1][%d] getSensorInstallersCCIdByQueryInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSensorInstallersCCIDByQueryInternalServerError) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetSensorInstallersCCIDByQueryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

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

	o.Payload = new(models.MsaReplyMetaOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
