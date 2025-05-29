package handler

import "net/http"

var HttpMessage = map[int]string{
	// Informational responses (1xx)
	http.StatusContinue:           "Request received, continue processing.",
	http.StatusSwitchingProtocols: "Protocol change requested, switching protocols.",
	http.StatusProcessing:         "Processing request, please wait.",
	http.StatusEarlyHints:         "Preliminary response, more information coming.",

	// Success (2xx)
	http.StatusOK:                   "Request successful.",
	http.StatusCreated:              "Resource successfully created.",
	http.StatusAccepted:             "Request accepted, processing in progress.",
	http.StatusNonAuthoritativeInfo: "Request successful, but information may be from another source.",
	http.StatusNoContent:            "Request successful, but no content to return.",
	http.StatusResetContent:         "Request successful, reset your view.",
	http.StatusPartialContent:       "Partial response returned.",
	http.StatusMultiStatus:          "Multiple status results available.",
	http.StatusAlreadyReported:      "Resource already reported previously.",
	http.StatusIMUsed:               "IM used successfully.",

	// Redirections (3xx)
	http.StatusMultipleChoices:   "Multiple options available, choose an alternative.",
	http.StatusMovedPermanently:  "Resource permanently moved to a new location.",
	http.StatusFound:             "Resource temporarily moved, check new location.",
	http.StatusSeeOther:          "See another resource for the response.",
	http.StatusNotModified:       "Resource not modified since last request.",
	http.StatusUseProxy:          "Request must be accessed via a proxy.",
	http.StatusTemporaryRedirect: "Resource temporarily moved, try later.",
	http.StatusPermanentRedirect: "Resource permanently moved, update your links.",

	// Client errors (4xx)
	http.StatusBadRequest:                   "Invalid request, please check the input.",
	http.StatusUnauthorized:                 "Authentication required for this request.",
	http.StatusPaymentRequired:              "Payment needed to access this resource.",
	http.StatusForbidden:                    "You do not have permission to access this resource.",
	http.StatusNotFound:                     "Requested resource not found.",
	http.StatusMethodNotAllowed:             "HTTP method not allowed for this resource.",
	http.StatusNotAcceptable:                "Requested format not supported.",
	http.StatusProxyAuthRequired:            "Proxy authentication required.",
	http.StatusRequestTimeout:               "Request took too long, please try again.",
	http.StatusConflict:                     "Request conflicts with the current state of the resource.",
	http.StatusGone:                         "Resource is no longer available.",
	http.StatusLengthRequired:               "Content length must be specified.",
	http.StatusPreconditionFailed:           "Preconditions failed, request cannot be processed.",
	http.StatusRequestEntityTooLarge:        "Request payload is too large.",
	http.StatusRequestURITooLong:            "Request URL is too long.",
	http.StatusUnsupportedMediaType:         "Media type not supported.",
	http.StatusRequestedRangeNotSatisfiable: "Requested range not available.",
	http.StatusExpectationFailed:            "Request expectations could not be met.",
	http.StatusTeapot:                       "I'm a teapot. ☕",
	http.StatusMisdirectedRequest:           "Request cannot be handled by this server.",
	http.StatusUnprocessableEntity:          "Request is well-formed but cannot be processed.",
	http.StatusLocked:                       "Resource is locked.",
	http.StatusFailedDependency:             "Dependency failure prevents request processing.",
	http.StatusTooEarly:                     "Request sent too early, try again later.",
	http.StatusUpgradeRequired:              "Upgrade required to access this resource.",
	http.StatusPreconditionRequired:         "Preconditions must be met before request processing.",
	http.StatusTooManyRequests:              "Too many requests, slow down.",
	http.StatusRequestHeaderFieldsTooLarge:  "Request headers are too large.",
	http.StatusUnavailableForLegalReasons:   "Resource unavailable due to legal restrictions.",

	// Server errors (5xx)
	http.StatusInternalServerError:           "An unexpected error occurred, please try again later.",
	http.StatusNotImplemented:                "Requested functionality is not implemented.",
	http.StatusBadGateway:                    "Invalid response from upstream server.",
	http.StatusServiceUnavailable:            "Service temporarily unavailable, try again later.",
	http.StatusGatewayTimeout:                "Upstream server did not respond in time.",
	http.StatusHTTPVersionNotSupported:       "HTTP version not supported.",
	http.StatusVariantAlsoNegotiates:         "Variant negotiation failed.",
	http.StatusInsufficientStorage:           "Server has insufficient storage to complete the request.",
	http.StatusLoopDetected:                  "Infinite loop detected in request processing.",
	http.StatusNotExtended:                   "Further extensions required to process request.",
	http.StatusNetworkAuthenticationRequired: "Authentication required to access the network.",
}
