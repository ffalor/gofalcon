// Code generated by go-swagger; DO NOT EDIT.

package certificate_based_exclusions

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

// CbExclusionsQueryV1Reader is a Reader for the CbExclusionsQueryV1 structure.
type CbExclusionsQueryV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CbExclusionsQueryV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCbExclusionsQueryV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCbExclusionsQueryV1BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCbExclusionsQueryV1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCbExclusionsQueryV1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCbExclusionsQueryV1TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCbExclusionsQueryV1InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /exclusions/queries/cert-based-exclusions/v1] cb-exclusions.query.v1", response, response.Code())
	}
}

// NewCbExclusionsQueryV1OK creates a CbExclusionsQueryV1OK with default headers values
func NewCbExclusionsQueryV1OK() *CbExclusionsQueryV1OK {
	return &CbExclusionsQueryV1OK{}
}

/*
CbExclusionsQueryV1OK describes a response with status code 200, with default header values.

OK
*/
type CbExclusionsQueryV1OK struct {

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

// IsSuccess returns true when this cb exclusions query v1 o k response has a 2xx status code
func (o *CbExclusionsQueryV1OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cb exclusions query v1 o k response has a 3xx status code
func (o *CbExclusionsQueryV1OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cb exclusions query v1 o k response has a 4xx status code
func (o *CbExclusionsQueryV1OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cb exclusions query v1 o k response has a 5xx status code
func (o *CbExclusionsQueryV1OK) IsServerError() bool {
	return false
}

// IsCode returns true when this cb exclusions query v1 o k response a status code equal to that given
func (o *CbExclusionsQueryV1OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the cb exclusions query v1 o k response
func (o *CbExclusionsQueryV1OK) Code() int {
	return 200
}

func (o *CbExclusionsQueryV1OK) Error() string {
	return fmt.Sprintf("[GET /exclusions/queries/cert-based-exclusions/v1][%d] cbExclusionsQueryV1OK  %+v", 200, o.Payload)
}

func (o *CbExclusionsQueryV1OK) String() string {
	return fmt.Sprintf("[GET /exclusions/queries/cert-based-exclusions/v1][%d] cbExclusionsQueryV1OK  %+v", 200, o.Payload)
}

func (o *CbExclusionsQueryV1OK) GetPayload() *models.MsaspecQueryResponse {
	return o.Payload
}

func (o *CbExclusionsQueryV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCbExclusionsQueryV1BadRequest creates a CbExclusionsQueryV1BadRequest with default headers values
func NewCbExclusionsQueryV1BadRequest() *CbExclusionsQueryV1BadRequest {
	return &CbExclusionsQueryV1BadRequest{}
}

/*
CbExclusionsQueryV1BadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CbExclusionsQueryV1BadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaspecResponseFields
}

// IsSuccess returns true when this cb exclusions query v1 bad request response has a 2xx status code
func (o *CbExclusionsQueryV1BadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cb exclusions query v1 bad request response has a 3xx status code
func (o *CbExclusionsQueryV1BadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cb exclusions query v1 bad request response has a 4xx status code
func (o *CbExclusionsQueryV1BadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this cb exclusions query v1 bad request response has a 5xx status code
func (o *CbExclusionsQueryV1BadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this cb exclusions query v1 bad request response a status code equal to that given
func (o *CbExclusionsQueryV1BadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the cb exclusions query v1 bad request response
func (o *CbExclusionsQueryV1BadRequest) Code() int {
	return 400
}

func (o *CbExclusionsQueryV1BadRequest) Error() string {
	return fmt.Sprintf("[GET /exclusions/queries/cert-based-exclusions/v1][%d] cbExclusionsQueryV1BadRequest  %+v", 400, o.Payload)
}

func (o *CbExclusionsQueryV1BadRequest) String() string {
	return fmt.Sprintf("[GET /exclusions/queries/cert-based-exclusions/v1][%d] cbExclusionsQueryV1BadRequest  %+v", 400, o.Payload)
}

func (o *CbExclusionsQueryV1BadRequest) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *CbExclusionsQueryV1BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaspecResponseFields)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCbExclusionsQueryV1Unauthorized creates a CbExclusionsQueryV1Unauthorized with default headers values
func NewCbExclusionsQueryV1Unauthorized() *CbExclusionsQueryV1Unauthorized {
	return &CbExclusionsQueryV1Unauthorized{}
}

/*
CbExclusionsQueryV1Unauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CbExclusionsQueryV1Unauthorized struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaspecResponseFields
}

// IsSuccess returns true when this cb exclusions query v1 unauthorized response has a 2xx status code
func (o *CbExclusionsQueryV1Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cb exclusions query v1 unauthorized response has a 3xx status code
func (o *CbExclusionsQueryV1Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cb exclusions query v1 unauthorized response has a 4xx status code
func (o *CbExclusionsQueryV1Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this cb exclusions query v1 unauthorized response has a 5xx status code
func (o *CbExclusionsQueryV1Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this cb exclusions query v1 unauthorized response a status code equal to that given
func (o *CbExclusionsQueryV1Unauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the cb exclusions query v1 unauthorized response
func (o *CbExclusionsQueryV1Unauthorized) Code() int {
	return 401
}

func (o *CbExclusionsQueryV1Unauthorized) Error() string {
	return fmt.Sprintf("[GET /exclusions/queries/cert-based-exclusions/v1][%d] cbExclusionsQueryV1Unauthorized  %+v", 401, o.Payload)
}

func (o *CbExclusionsQueryV1Unauthorized) String() string {
	return fmt.Sprintf("[GET /exclusions/queries/cert-based-exclusions/v1][%d] cbExclusionsQueryV1Unauthorized  %+v", 401, o.Payload)
}

func (o *CbExclusionsQueryV1Unauthorized) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *CbExclusionsQueryV1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaspecResponseFields)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCbExclusionsQueryV1Forbidden creates a CbExclusionsQueryV1Forbidden with default headers values
func NewCbExclusionsQueryV1Forbidden() *CbExclusionsQueryV1Forbidden {
	return &CbExclusionsQueryV1Forbidden{}
}

/*
CbExclusionsQueryV1Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CbExclusionsQueryV1Forbidden struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaspecResponseFields
}

// IsSuccess returns true when this cb exclusions query v1 forbidden response has a 2xx status code
func (o *CbExclusionsQueryV1Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cb exclusions query v1 forbidden response has a 3xx status code
func (o *CbExclusionsQueryV1Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cb exclusions query v1 forbidden response has a 4xx status code
func (o *CbExclusionsQueryV1Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this cb exclusions query v1 forbidden response has a 5xx status code
func (o *CbExclusionsQueryV1Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this cb exclusions query v1 forbidden response a status code equal to that given
func (o *CbExclusionsQueryV1Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the cb exclusions query v1 forbidden response
func (o *CbExclusionsQueryV1Forbidden) Code() int {
	return 403
}

func (o *CbExclusionsQueryV1Forbidden) Error() string {
	return fmt.Sprintf("[GET /exclusions/queries/cert-based-exclusions/v1][%d] cbExclusionsQueryV1Forbidden  %+v", 403, o.Payload)
}

func (o *CbExclusionsQueryV1Forbidden) String() string {
	return fmt.Sprintf("[GET /exclusions/queries/cert-based-exclusions/v1][%d] cbExclusionsQueryV1Forbidden  %+v", 403, o.Payload)
}

func (o *CbExclusionsQueryV1Forbidden) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *CbExclusionsQueryV1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaspecResponseFields)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCbExclusionsQueryV1TooManyRequests creates a CbExclusionsQueryV1TooManyRequests with default headers values
func NewCbExclusionsQueryV1TooManyRequests() *CbExclusionsQueryV1TooManyRequests {
	return &CbExclusionsQueryV1TooManyRequests{}
}

/*
CbExclusionsQueryV1TooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type CbExclusionsQueryV1TooManyRequests struct {

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

// IsSuccess returns true when this cb exclusions query v1 too many requests response has a 2xx status code
func (o *CbExclusionsQueryV1TooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cb exclusions query v1 too many requests response has a 3xx status code
func (o *CbExclusionsQueryV1TooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cb exclusions query v1 too many requests response has a 4xx status code
func (o *CbExclusionsQueryV1TooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this cb exclusions query v1 too many requests response has a 5xx status code
func (o *CbExclusionsQueryV1TooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this cb exclusions query v1 too many requests response a status code equal to that given
func (o *CbExclusionsQueryV1TooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the cb exclusions query v1 too many requests response
func (o *CbExclusionsQueryV1TooManyRequests) Code() int {
	return 429
}

func (o *CbExclusionsQueryV1TooManyRequests) Error() string {
	return fmt.Sprintf("[GET /exclusions/queries/cert-based-exclusions/v1][%d] cbExclusionsQueryV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *CbExclusionsQueryV1TooManyRequests) String() string {
	return fmt.Sprintf("[GET /exclusions/queries/cert-based-exclusions/v1][%d] cbExclusionsQueryV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *CbExclusionsQueryV1TooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *CbExclusionsQueryV1TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCbExclusionsQueryV1InternalServerError creates a CbExclusionsQueryV1InternalServerError with default headers values
func NewCbExclusionsQueryV1InternalServerError() *CbExclusionsQueryV1InternalServerError {
	return &CbExclusionsQueryV1InternalServerError{}
}

/*
CbExclusionsQueryV1InternalServerError describes a response with status code 500, with default header values.

Unexpected Error
*/
type CbExclusionsQueryV1InternalServerError struct {

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

// IsSuccess returns true when this cb exclusions query v1 internal server error response has a 2xx status code
func (o *CbExclusionsQueryV1InternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cb exclusions query v1 internal server error response has a 3xx status code
func (o *CbExclusionsQueryV1InternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cb exclusions query v1 internal server error response has a 4xx status code
func (o *CbExclusionsQueryV1InternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this cb exclusions query v1 internal server error response has a 5xx status code
func (o *CbExclusionsQueryV1InternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this cb exclusions query v1 internal server error response a status code equal to that given
func (o *CbExclusionsQueryV1InternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the cb exclusions query v1 internal server error response
func (o *CbExclusionsQueryV1InternalServerError) Code() int {
	return 500
}

func (o *CbExclusionsQueryV1InternalServerError) Error() string {
	return fmt.Sprintf("[GET /exclusions/queries/cert-based-exclusions/v1][%d] cbExclusionsQueryV1InternalServerError  %+v", 500, o.Payload)
}

func (o *CbExclusionsQueryV1InternalServerError) String() string {
	return fmt.Sprintf("[GET /exclusions/queries/cert-based-exclusions/v1][%d] cbExclusionsQueryV1InternalServerError  %+v", 500, o.Payload)
}

func (o *CbExclusionsQueryV1InternalServerError) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *CbExclusionsQueryV1InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
