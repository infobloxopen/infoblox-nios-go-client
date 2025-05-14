package option

import (
	"github.com/unasra/nios-go-client/internal"
	"net/http"
)

// ClientOption is a function that applies configuration options to the API Client.
type ClientOption func(configuration *internal.Configuration)

// WithNIOSHostUrl returns a ClientOption that sets the URL for Infoblox NIOS Portal
// Can also be configured using the `NIOS_HOST_URL` environment variable.
// Required
func WithNIOSHostUrl(NIOSHostURL string) ClientOption {
	return func(configuration *internal.Configuration) {
		if NIOSHostURL != "" {
			configuration.NIOSHostURL = NIOSHostURL
		}
	}
}

// WithNIOSAuth returns a ClientOption that sets the NIOSAuth for accessing the NIOS WAPI.
// Can also be configured by using the `NIOS_AUTH` environment variable.
//
// Required.
func WithNIOSAuth(NIOSAuth string) ClientOption {
	return func(configuration *internal.Configuration) {
		if NIOSAuth != "" {
			configuration.NIOSAuth = NIOSAuth
		}
	}
}

// WithHTTPClient returns a ClientOption that sets the HTTPClient to use for the SDK.
// Optional. The default HTTPClient will be used if not provided.
func WithHTTPClient(httpClient *http.Client) ClientOption {
	return func(configuration *internal.Configuration) {
		if httpClient != nil {
			configuration.HTTPClient = httpClient
		}
	}
}

// WithDefaultTags returns a ClientOption that sets the tags the client can set by default for objects that has tags support.
// Optional.
func WithDefaultTags(defaultTags map[string]string) ClientOption {
	return func(configuration *internal.Configuration) {
		configuration.DefaultTags = defaultTags
	}
}

// WithClientName returns a ClientOption that sets the name of the client using the SDK.
// This can be used to identify the client in the audit logs.
// Optional. If not provided, the client name will be set to "nios-go-client".
func WithClientName(clientName string) ClientOption {
	return func(configuration *internal.Configuration) {
		if clientName != "" {
			configuration.ClientName = clientName
		}
	}
}

// WithDebug returns a ClientOption that sets the debug mode.
// Enabling the debug flag will write the request and response to the log.
func WithDebug(debug bool) ClientOption {
	return func(configuration *internal.Configuration) {
		configuration.Debug = debug
	}
}
