// Code generated by go-swagger; DO NOT EDIT.

package spotlight_vulnerabilities

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

// CombinedQueryVulnerabilitiesReader is a Reader for the CombinedQueryVulnerabilities structure.
type CombinedQueryVulnerabilitiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CombinedQueryVulnerabilitiesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCombinedQueryVulnerabilitiesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCombinedQueryVulnerabilitiesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCombinedQueryVulnerabilitiesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCombinedQueryVulnerabilitiesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCombinedQueryVulnerabilitiesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCombinedQueryVulnerabilitiesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewCombinedQueryVulnerabilitiesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /spotlight/combined/vulnerabilities/v1] combinedQueryVulnerabilities", response, response.Code())
	}
}

// NewCombinedQueryVulnerabilitiesOK creates a CombinedQueryVulnerabilitiesOK with default headers values
func NewCombinedQueryVulnerabilitiesOK() *CombinedQueryVulnerabilitiesOK {
	return &CombinedQueryVulnerabilitiesOK{}
}

/*
CombinedQueryVulnerabilitiesOK describes a response with status code 200, with default header values.

OK
*/
type CombinedQueryVulnerabilitiesOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainSPAPICombinedVulnerabilitiesResponse
}

// IsSuccess returns true when this combined query vulnerabilities o k response has a 2xx status code
func (o *CombinedQueryVulnerabilitiesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this combined query vulnerabilities o k response has a 3xx status code
func (o *CombinedQueryVulnerabilitiesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this combined query vulnerabilities o k response has a 4xx status code
func (o *CombinedQueryVulnerabilitiesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this combined query vulnerabilities o k response has a 5xx status code
func (o *CombinedQueryVulnerabilitiesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this combined query vulnerabilities o k response a status code equal to that given
func (o *CombinedQueryVulnerabilitiesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the combined query vulnerabilities o k response
func (o *CombinedQueryVulnerabilitiesOK) Code() int {
	return 200
}

func (o *CombinedQueryVulnerabilitiesOK) Error() string {
	return fmt.Sprintf("[GET /spotlight/combined/vulnerabilities/v1][%d] combinedQueryVulnerabilitiesOK  %+v", 200, o.Payload)
}

func (o *CombinedQueryVulnerabilitiesOK) String() string {
	return fmt.Sprintf("[GET /spotlight/combined/vulnerabilities/v1][%d] combinedQueryVulnerabilitiesOK  %+v", 200, o.Payload)
}

func (o *CombinedQueryVulnerabilitiesOK) GetPayload() *models.DomainSPAPICombinedVulnerabilitiesResponse {
	return o.Payload
}

func (o *CombinedQueryVulnerabilitiesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainSPAPICombinedVulnerabilitiesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCombinedQueryVulnerabilitiesBadRequest creates a CombinedQueryVulnerabilitiesBadRequest with default headers values
func NewCombinedQueryVulnerabilitiesBadRequest() *CombinedQueryVulnerabilitiesBadRequest {
	return &CombinedQueryVulnerabilitiesBadRequest{}
}

/*
CombinedQueryVulnerabilitiesBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CombinedQueryVulnerabilitiesBadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainSPAPICombinedVulnerabilitiesResponse
}

// IsSuccess returns true when this combined query vulnerabilities bad request response has a 2xx status code
func (o *CombinedQueryVulnerabilitiesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this combined query vulnerabilities bad request response has a 3xx status code
func (o *CombinedQueryVulnerabilitiesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this combined query vulnerabilities bad request response has a 4xx status code
func (o *CombinedQueryVulnerabilitiesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this combined query vulnerabilities bad request response has a 5xx status code
func (o *CombinedQueryVulnerabilitiesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this combined query vulnerabilities bad request response a status code equal to that given
func (o *CombinedQueryVulnerabilitiesBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the combined query vulnerabilities bad request response
func (o *CombinedQueryVulnerabilitiesBadRequest) Code() int {
	return 400
}

func (o *CombinedQueryVulnerabilitiesBadRequest) Error() string {
	return fmt.Sprintf("[GET /spotlight/combined/vulnerabilities/v1][%d] combinedQueryVulnerabilitiesBadRequest  %+v", 400, o.Payload)
}

func (o *CombinedQueryVulnerabilitiesBadRequest) String() string {
	return fmt.Sprintf("[GET /spotlight/combined/vulnerabilities/v1][%d] combinedQueryVulnerabilitiesBadRequest  %+v", 400, o.Payload)
}

func (o *CombinedQueryVulnerabilitiesBadRequest) GetPayload() *models.DomainSPAPICombinedVulnerabilitiesResponse {
	return o.Payload
}

func (o *CombinedQueryVulnerabilitiesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainSPAPICombinedVulnerabilitiesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCombinedQueryVulnerabilitiesForbidden creates a CombinedQueryVulnerabilitiesForbidden with default headers values
func NewCombinedQueryVulnerabilitiesForbidden() *CombinedQueryVulnerabilitiesForbidden {
	return &CombinedQueryVulnerabilitiesForbidden{}
}

/*
CombinedQueryVulnerabilitiesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CombinedQueryVulnerabilitiesForbidden struct {

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

// IsSuccess returns true when this combined query vulnerabilities forbidden response has a 2xx status code
func (o *CombinedQueryVulnerabilitiesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this combined query vulnerabilities forbidden response has a 3xx status code
func (o *CombinedQueryVulnerabilitiesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this combined query vulnerabilities forbidden response has a 4xx status code
func (o *CombinedQueryVulnerabilitiesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this combined query vulnerabilities forbidden response has a 5xx status code
func (o *CombinedQueryVulnerabilitiesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this combined query vulnerabilities forbidden response a status code equal to that given
func (o *CombinedQueryVulnerabilitiesForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the combined query vulnerabilities forbidden response
func (o *CombinedQueryVulnerabilitiesForbidden) Code() int {
	return 403
}

func (o *CombinedQueryVulnerabilitiesForbidden) Error() string {
	return fmt.Sprintf("[GET /spotlight/combined/vulnerabilities/v1][%d] combinedQueryVulnerabilitiesForbidden  %+v", 403, o.Payload)
}

func (o *CombinedQueryVulnerabilitiesForbidden) String() string {
	return fmt.Sprintf("[GET /spotlight/combined/vulnerabilities/v1][%d] combinedQueryVulnerabilitiesForbidden  %+v", 403, o.Payload)
}

func (o *CombinedQueryVulnerabilitiesForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *CombinedQueryVulnerabilitiesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCombinedQueryVulnerabilitiesNotFound creates a CombinedQueryVulnerabilitiesNotFound with default headers values
func NewCombinedQueryVulnerabilitiesNotFound() *CombinedQueryVulnerabilitiesNotFound {
	return &CombinedQueryVulnerabilitiesNotFound{}
}

/*
CombinedQueryVulnerabilitiesNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CombinedQueryVulnerabilitiesNotFound struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainSPAPICombinedVulnerabilitiesResponse
}

// IsSuccess returns true when this combined query vulnerabilities not found response has a 2xx status code
func (o *CombinedQueryVulnerabilitiesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this combined query vulnerabilities not found response has a 3xx status code
func (o *CombinedQueryVulnerabilitiesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this combined query vulnerabilities not found response has a 4xx status code
func (o *CombinedQueryVulnerabilitiesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this combined query vulnerabilities not found response has a 5xx status code
func (o *CombinedQueryVulnerabilitiesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this combined query vulnerabilities not found response a status code equal to that given
func (o *CombinedQueryVulnerabilitiesNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the combined query vulnerabilities not found response
func (o *CombinedQueryVulnerabilitiesNotFound) Code() int {
	return 404
}

func (o *CombinedQueryVulnerabilitiesNotFound) Error() string {
	return fmt.Sprintf("[GET /spotlight/combined/vulnerabilities/v1][%d] combinedQueryVulnerabilitiesNotFound  %+v", 404, o.Payload)
}

func (o *CombinedQueryVulnerabilitiesNotFound) String() string {
	return fmt.Sprintf("[GET /spotlight/combined/vulnerabilities/v1][%d] combinedQueryVulnerabilitiesNotFound  %+v", 404, o.Payload)
}

func (o *CombinedQueryVulnerabilitiesNotFound) GetPayload() *models.DomainSPAPICombinedVulnerabilitiesResponse {
	return o.Payload
}

func (o *CombinedQueryVulnerabilitiesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainSPAPICombinedVulnerabilitiesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCombinedQueryVulnerabilitiesTooManyRequests creates a CombinedQueryVulnerabilitiesTooManyRequests with default headers values
func NewCombinedQueryVulnerabilitiesTooManyRequests() *CombinedQueryVulnerabilitiesTooManyRequests {
	return &CombinedQueryVulnerabilitiesTooManyRequests{}
}

/*
CombinedQueryVulnerabilitiesTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type CombinedQueryVulnerabilitiesTooManyRequests struct {

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

// IsSuccess returns true when this combined query vulnerabilities too many requests response has a 2xx status code
func (o *CombinedQueryVulnerabilitiesTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this combined query vulnerabilities too many requests response has a 3xx status code
func (o *CombinedQueryVulnerabilitiesTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this combined query vulnerabilities too many requests response has a 4xx status code
func (o *CombinedQueryVulnerabilitiesTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this combined query vulnerabilities too many requests response has a 5xx status code
func (o *CombinedQueryVulnerabilitiesTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this combined query vulnerabilities too many requests response a status code equal to that given
func (o *CombinedQueryVulnerabilitiesTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the combined query vulnerabilities too many requests response
func (o *CombinedQueryVulnerabilitiesTooManyRequests) Code() int {
	return 429
}

func (o *CombinedQueryVulnerabilitiesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /spotlight/combined/vulnerabilities/v1][%d] combinedQueryVulnerabilitiesTooManyRequests  %+v", 429, o.Payload)
}

func (o *CombinedQueryVulnerabilitiesTooManyRequests) String() string {
	return fmt.Sprintf("[GET /spotlight/combined/vulnerabilities/v1][%d] combinedQueryVulnerabilitiesTooManyRequests  %+v", 429, o.Payload)
}

func (o *CombinedQueryVulnerabilitiesTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *CombinedQueryVulnerabilitiesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCombinedQueryVulnerabilitiesInternalServerError creates a CombinedQueryVulnerabilitiesInternalServerError with default headers values
func NewCombinedQueryVulnerabilitiesInternalServerError() *CombinedQueryVulnerabilitiesInternalServerError {
	return &CombinedQueryVulnerabilitiesInternalServerError{}
}

/*
CombinedQueryVulnerabilitiesInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type CombinedQueryVulnerabilitiesInternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainSPAPICombinedVulnerabilitiesResponse
}

// IsSuccess returns true when this combined query vulnerabilities internal server error response has a 2xx status code
func (o *CombinedQueryVulnerabilitiesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this combined query vulnerabilities internal server error response has a 3xx status code
func (o *CombinedQueryVulnerabilitiesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this combined query vulnerabilities internal server error response has a 4xx status code
func (o *CombinedQueryVulnerabilitiesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this combined query vulnerabilities internal server error response has a 5xx status code
func (o *CombinedQueryVulnerabilitiesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this combined query vulnerabilities internal server error response a status code equal to that given
func (o *CombinedQueryVulnerabilitiesInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the combined query vulnerabilities internal server error response
func (o *CombinedQueryVulnerabilitiesInternalServerError) Code() int {
	return 500
}

func (o *CombinedQueryVulnerabilitiesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /spotlight/combined/vulnerabilities/v1][%d] combinedQueryVulnerabilitiesInternalServerError  %+v", 500, o.Payload)
}

func (o *CombinedQueryVulnerabilitiesInternalServerError) String() string {
	return fmt.Sprintf("[GET /spotlight/combined/vulnerabilities/v1][%d] combinedQueryVulnerabilitiesInternalServerError  %+v", 500, o.Payload)
}

func (o *CombinedQueryVulnerabilitiesInternalServerError) GetPayload() *models.DomainSPAPICombinedVulnerabilitiesResponse {
	return o.Payload
}

func (o *CombinedQueryVulnerabilitiesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainSPAPICombinedVulnerabilitiesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCombinedQueryVulnerabilitiesServiceUnavailable creates a CombinedQueryVulnerabilitiesServiceUnavailable with default headers values
func NewCombinedQueryVulnerabilitiesServiceUnavailable() *CombinedQueryVulnerabilitiesServiceUnavailable {
	return &CombinedQueryVulnerabilitiesServiceUnavailable{}
}

/*
CombinedQueryVulnerabilitiesServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable
*/
type CombinedQueryVulnerabilitiesServiceUnavailable struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainSPAPICombinedVulnerabilitiesResponse
}

// IsSuccess returns true when this combined query vulnerabilities service unavailable response has a 2xx status code
func (o *CombinedQueryVulnerabilitiesServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this combined query vulnerabilities service unavailable response has a 3xx status code
func (o *CombinedQueryVulnerabilitiesServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this combined query vulnerabilities service unavailable response has a 4xx status code
func (o *CombinedQueryVulnerabilitiesServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this combined query vulnerabilities service unavailable response has a 5xx status code
func (o *CombinedQueryVulnerabilitiesServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this combined query vulnerabilities service unavailable response a status code equal to that given
func (o *CombinedQueryVulnerabilitiesServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

// Code gets the status code for the combined query vulnerabilities service unavailable response
func (o *CombinedQueryVulnerabilitiesServiceUnavailable) Code() int {
	return 503
}

func (o *CombinedQueryVulnerabilitiesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /spotlight/combined/vulnerabilities/v1][%d] combinedQueryVulnerabilitiesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *CombinedQueryVulnerabilitiesServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /spotlight/combined/vulnerabilities/v1][%d] combinedQueryVulnerabilitiesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *CombinedQueryVulnerabilitiesServiceUnavailable) GetPayload() *models.DomainSPAPICombinedVulnerabilitiesResponse {
	return o.Payload
}

func (o *CombinedQueryVulnerabilitiesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainSPAPICombinedVulnerabilitiesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
