// Code generated by go-swagger; DO NOT EDIT.

package firewall_management

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

// AggregatePolicyRulesReader is a Reader for the AggregatePolicyRules structure.
type AggregatePolicyRulesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AggregatePolicyRulesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAggregatePolicyRulesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAggregatePolicyRulesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAggregatePolicyRulesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewAggregatePolicyRulesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAggregatePolicyRulesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /fwmgr/aggregates/policy-rules/GET/v1] aggregate-policy-rules", response, response.Code())
	}
}

// NewAggregatePolicyRulesOK creates a AggregatePolicyRulesOK with default headers values
func NewAggregatePolicyRulesOK() *AggregatePolicyRulesOK {
	return &AggregatePolicyRulesOK{}
}

/*
AggregatePolicyRulesOK describes a response with status code 200, with default header values.

OK
*/
type AggregatePolicyRulesOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.FwmgrAPIAggregatesResponse
}

// IsSuccess returns true when this aggregate policy rules o k response has a 2xx status code
func (o *AggregatePolicyRulesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this aggregate policy rules o k response has a 3xx status code
func (o *AggregatePolicyRulesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this aggregate policy rules o k response has a 4xx status code
func (o *AggregatePolicyRulesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this aggregate policy rules o k response has a 5xx status code
func (o *AggregatePolicyRulesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this aggregate policy rules o k response a status code equal to that given
func (o *AggregatePolicyRulesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the aggregate policy rules o k response
func (o *AggregatePolicyRulesOK) Code() int {
	return 200
}

func (o *AggregatePolicyRulesOK) Error() string {
	return fmt.Sprintf("[POST /fwmgr/aggregates/policy-rules/GET/v1][%d] aggregatePolicyRulesOK  %+v", 200, o.Payload)
}

func (o *AggregatePolicyRulesOK) String() string {
	return fmt.Sprintf("[POST /fwmgr/aggregates/policy-rules/GET/v1][%d] aggregatePolicyRulesOK  %+v", 200, o.Payload)
}

func (o *AggregatePolicyRulesOK) GetPayload() *models.FwmgrAPIAggregatesResponse {
	return o.Payload
}

func (o *AggregatePolicyRulesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.FwmgrAPIAggregatesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAggregatePolicyRulesBadRequest creates a AggregatePolicyRulesBadRequest with default headers values
func NewAggregatePolicyRulesBadRequest() *AggregatePolicyRulesBadRequest {
	return &AggregatePolicyRulesBadRequest{}
}

/*
AggregatePolicyRulesBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AggregatePolicyRulesBadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.FwmgrMsaspecResponseFields
}

// IsSuccess returns true when this aggregate policy rules bad request response has a 2xx status code
func (o *AggregatePolicyRulesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this aggregate policy rules bad request response has a 3xx status code
func (o *AggregatePolicyRulesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this aggregate policy rules bad request response has a 4xx status code
func (o *AggregatePolicyRulesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this aggregate policy rules bad request response has a 5xx status code
func (o *AggregatePolicyRulesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this aggregate policy rules bad request response a status code equal to that given
func (o *AggregatePolicyRulesBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the aggregate policy rules bad request response
func (o *AggregatePolicyRulesBadRequest) Code() int {
	return 400
}

func (o *AggregatePolicyRulesBadRequest) Error() string {
	return fmt.Sprintf("[POST /fwmgr/aggregates/policy-rules/GET/v1][%d] aggregatePolicyRulesBadRequest  %+v", 400, o.Payload)
}

func (o *AggregatePolicyRulesBadRequest) String() string {
	return fmt.Sprintf("[POST /fwmgr/aggregates/policy-rules/GET/v1][%d] aggregatePolicyRulesBadRequest  %+v", 400, o.Payload)
}

func (o *AggregatePolicyRulesBadRequest) GetPayload() *models.FwmgrMsaspecResponseFields {
	return o.Payload
}

func (o *AggregatePolicyRulesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.FwmgrMsaspecResponseFields)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAggregatePolicyRulesForbidden creates a AggregatePolicyRulesForbidden with default headers values
func NewAggregatePolicyRulesForbidden() *AggregatePolicyRulesForbidden {
	return &AggregatePolicyRulesForbidden{}
}

/*
AggregatePolicyRulesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type AggregatePolicyRulesForbidden struct {

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

// IsSuccess returns true when this aggregate policy rules forbidden response has a 2xx status code
func (o *AggregatePolicyRulesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this aggregate policy rules forbidden response has a 3xx status code
func (o *AggregatePolicyRulesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this aggregate policy rules forbidden response has a 4xx status code
func (o *AggregatePolicyRulesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this aggregate policy rules forbidden response has a 5xx status code
func (o *AggregatePolicyRulesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this aggregate policy rules forbidden response a status code equal to that given
func (o *AggregatePolicyRulesForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the aggregate policy rules forbidden response
func (o *AggregatePolicyRulesForbidden) Code() int {
	return 403
}

func (o *AggregatePolicyRulesForbidden) Error() string {
	return fmt.Sprintf("[POST /fwmgr/aggregates/policy-rules/GET/v1][%d] aggregatePolicyRulesForbidden  %+v", 403, o.Payload)
}

func (o *AggregatePolicyRulesForbidden) String() string {
	return fmt.Sprintf("[POST /fwmgr/aggregates/policy-rules/GET/v1][%d] aggregatePolicyRulesForbidden  %+v", 403, o.Payload)
}

func (o *AggregatePolicyRulesForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *AggregatePolicyRulesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewAggregatePolicyRulesTooManyRequests creates a AggregatePolicyRulesTooManyRequests with default headers values
func NewAggregatePolicyRulesTooManyRequests() *AggregatePolicyRulesTooManyRequests {
	return &AggregatePolicyRulesTooManyRequests{}
}

/*
AggregatePolicyRulesTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type AggregatePolicyRulesTooManyRequests struct {

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

// IsSuccess returns true when this aggregate policy rules too many requests response has a 2xx status code
func (o *AggregatePolicyRulesTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this aggregate policy rules too many requests response has a 3xx status code
func (o *AggregatePolicyRulesTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this aggregate policy rules too many requests response has a 4xx status code
func (o *AggregatePolicyRulesTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this aggregate policy rules too many requests response has a 5xx status code
func (o *AggregatePolicyRulesTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this aggregate policy rules too many requests response a status code equal to that given
func (o *AggregatePolicyRulesTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the aggregate policy rules too many requests response
func (o *AggregatePolicyRulesTooManyRequests) Code() int {
	return 429
}

func (o *AggregatePolicyRulesTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /fwmgr/aggregates/policy-rules/GET/v1][%d] aggregatePolicyRulesTooManyRequests  %+v", 429, o.Payload)
}

func (o *AggregatePolicyRulesTooManyRequests) String() string {
	return fmt.Sprintf("[POST /fwmgr/aggregates/policy-rules/GET/v1][%d] aggregatePolicyRulesTooManyRequests  %+v", 429, o.Payload)
}

func (o *AggregatePolicyRulesTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *AggregatePolicyRulesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewAggregatePolicyRulesInternalServerError creates a AggregatePolicyRulesInternalServerError with default headers values
func NewAggregatePolicyRulesInternalServerError() *AggregatePolicyRulesInternalServerError {
	return &AggregatePolicyRulesInternalServerError{}
}

/*
AggregatePolicyRulesInternalServerError describes a response with status code 500, with default header values.

Unexpected Error
*/
type AggregatePolicyRulesInternalServerError struct {

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

// IsSuccess returns true when this aggregate policy rules internal server error response has a 2xx status code
func (o *AggregatePolicyRulesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this aggregate policy rules internal server error response has a 3xx status code
func (o *AggregatePolicyRulesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this aggregate policy rules internal server error response has a 4xx status code
func (o *AggregatePolicyRulesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this aggregate policy rules internal server error response has a 5xx status code
func (o *AggregatePolicyRulesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this aggregate policy rules internal server error response a status code equal to that given
func (o *AggregatePolicyRulesInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the aggregate policy rules internal server error response
func (o *AggregatePolicyRulesInternalServerError) Code() int {
	return 500
}

func (o *AggregatePolicyRulesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /fwmgr/aggregates/policy-rules/GET/v1][%d] aggregatePolicyRulesInternalServerError  %+v", 500, o.Payload)
}

func (o *AggregatePolicyRulesInternalServerError) String() string {
	return fmt.Sprintf("[POST /fwmgr/aggregates/policy-rules/GET/v1][%d] aggregatePolicyRulesInternalServerError  %+v", 500, o.Payload)
}

func (o *AggregatePolicyRulesInternalServerError) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *AggregatePolicyRulesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
