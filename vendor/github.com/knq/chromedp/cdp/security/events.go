package security

// Code generated by chromedp-gen. DO NOT EDIT.

import (
	cdp "github.com/knq/chromedp/cdp"
)

// EventSecurityStateChanged the security state of the page changed.
type EventSecurityStateChanged struct {
	SecurityState         State                  `json:"securityState"`         // Security state.
	SchemeIsCryptographic bool                   `json:"schemeIsCryptographic"` // True if the page was loaded over cryptographic transport such as HTTPS.
	Explanations          []*StateExplanation    `json:"explanations"`          // List of explanations for the security state. If the overall security state is `insecure` or `warning`, at least one corresponding explanation should be included.
	InsecureContentStatus *InsecureContentStatus `json:"insecureContentStatus"` // Information about insecure content on the page.
	Summary               string                 `json:"summary,omitempty"`     // Overrides user-visible description of the state.
}

// EventCertificateError there is a certificate error. If overriding
// certificate errors is enabled, then it should be handled with the
// handleCertificateError command. Note: this event does not fire if the
// certificate error has been allowed internally.
type EventCertificateError struct {
	EventID    int64  `json:"eventId"`    // The ID of the event.
	ErrorType  string `json:"errorType"`  // The type of the error.
	RequestURL string `json:"requestURL"` // The url that was requested.
}

// EventTypes all event types in the domain.
var EventTypes = []cdp.MethodType{
	cdp.EventSecuritySecurityStateChanged,
	cdp.EventSecurityCertificateError,
}
