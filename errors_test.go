package api

import (
	"errors"
	"net/http"
	"testing"
)

var allKnownCodes100 = []int{
	http.StatusContinue,
	http.StatusSwitchingProtocols,
	http.StatusProcessing,
	http.StatusEarlyHints,
}

var allKnownCodes200 = []int{
	// http.StatusOK, // should return nil, not an error
	http.StatusCreated,
	http.StatusAccepted,
	http.StatusNonAuthoritativeInfo,
	http.StatusNoContent,
	http.StatusResetContent,
	http.StatusPartialContent,
	http.StatusMultiStatus,
	http.StatusAlreadyReported,
	http.StatusIMUsed,
}

var allKnownCodes300 = []int{
	http.StatusMultipleChoices,
	http.StatusMovedPermanently,
	http.StatusFound,
	http.StatusSeeOther,
	http.StatusNotModified,
	http.StatusUseProxy,
	http.StatusTemporaryRedirect,
	http.StatusPermanentRedirect,
}

var allKnownCodes400 = []int{
	http.StatusBadRequest,
	http.StatusUnauthorized,
	http.StatusPaymentRequired,
	http.StatusForbidden,
	http.StatusNotFound,
	http.StatusMethodNotAllowed,
	http.StatusNotAcceptable,
	http.StatusProxyAuthRequired,
	http.StatusRequestTimeout,
	http.StatusConflict,
	http.StatusGone,
	http.StatusLengthRequired,
	http.StatusPreconditionFailed,
	http.StatusRequestEntityTooLarge,
	http.StatusRequestURITooLong,
	http.StatusUnsupportedMediaType,
	http.StatusRequestedRangeNotSatisfiable,
	http.StatusExpectationFailed,
	http.StatusTeapot,
	http.StatusMisdirectedRequest,
	http.StatusUnprocessableEntity,
	http.StatusLocked,
	http.StatusFailedDependency,
	http.StatusTooEarly,
	http.StatusUpgradeRequired,
	http.StatusPreconditionRequired,
	http.StatusTooManyRequests,
	http.StatusRequestHeaderFieldsTooLarge,
	http.StatusUnavailableForLegalReasons,
}

var allKnownCodes500 = []int{
	http.StatusInternalServerError,
	http.StatusNotImplemented,
	http.StatusBadGateway,
	http.StatusServiceUnavailable,
	http.StatusGatewayTimeout,
	http.StatusHTTPVersionNotSupported,
	http.StatusVariantAlsoNegotiates,
	http.StatusInsufficientStorage,
	http.StatusLoopDetected,
	http.StatusNotExtended,
	http.StatusNetworkAuthenticationRequired,
}
func TestErrors(t *testing.T) {

	t.Run("StatusOK", func(t *testing.T) {
		res := &Response{ Body: "", Status: http.StatusOK }
		err := newQueryError(res)
		if err != nil {
			t.Errorf("error generated, nil expected: %v", err)
		}
	})

	t.Run("100-199 status errors", func(t *testing.T) {
		for _, code := range allKnownCodes100 {
			res := &Response{ Body: "", Status: code }
			err := newQueryError(res)
			if !errors.Is(err, ErrInformation) {
				t.Errorf("for code %d, expected ErrInformation, got %v", code, err)
			}
		}
	})

	t.Run("200-299 status errors", func(t *testing.T) {
		for _, code := range allKnownCodes200 {
			res := &Response{ Body: "", Status: code }
			err := newQueryError(res)
			if !errors.Is(err, ErrSuccess) {
				t.Errorf("for code %d, expected ErrSuccess, got %v", code, err)
			}
		}
	})

	t.Run("300-399 status errors", func(t *testing.T) {
		for _, code := range allKnownCodes300 {
			res := &Response{ Body: "", Status: code }
			err := newQueryError(res)
			if !errors.Is(err, ErrRedirection) {
				t.Errorf("for code %d, expected ErrRedirection, got %v", code, err)
			}
		}
	})

	t.Run("400-499 status errors", func(t *testing.T) {
		for _, code := range allKnownCodes400 {
			res := &Response{ Body: "", Status: code }
			err := newQueryError(res)
			if !errors.Is(err, ErrClient) {
				t.Errorf("for code %d, expected ErrClient, got %v", code, err)
			}
		}
	})

	t.Run("500-599 status errors", func(t *testing.T) {
		for _, code := range allKnownCodes500 {
			res := &Response{ Body: "", Status: code }
			err := newQueryError(res)
			if !errors.Is(err, ErrServer) {
				t.Errorf("for code %d, expected ErrServer, got %v", code, err)
			}
		}
	})

	t.Run("error return values", func(t *testing.T) {
		for _, code := range allKnownCodes100 {
			res := &Response{ Body: "", Status: code }
			err := newQueryError(res)
			if err.Status() != code {
				t.Errorf("status code %d was not correctly returned by error (got %d)", code, err.Status())
			}
		}
	})

	t.Run("unsupported code handling", func(t *testing.T) {
		for _, code := range []int{99,601} {
			res := &Response{ Body: "", Status: code }
			err := newQueryError(res)
			if !errors.Is(err, ErrUnsupportedRange) {
				t.Errorf("for code %d, expected ErrUnsupportedRange, got %v", code, err)
			}
		}
	})

}
