// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	dcl "github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"

	dataproc "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/dataproc"
	eventarc "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/eventarc"
)

func CreateDataprocClient(config *Config, userAgent, billingProject string) *dataproc.Client {
	dclClientOptions := dcl.WithHTTPClient(config.client)
	dclUserAgentOptions := dcl.WithUserAgent(userAgent)
	dclLoggerOptions := dcl.WithLogger(dclLogger{})
	var dclConfig *dcl.Config
	if config.UserProjectOverride && billingProject != "" {
		dclBillingProjectHeader := dcl.WithHeader("X-Goog-User-Project", billingProject)
		dclConfig = dcl.NewConfig(
			dclClientOptions,
			dclUserAgentOptions,
			dclLoggerOptions,
			dcl.WithBasePath(config.DataprocBasePath),
			dclBillingProjectHeader,
		)
	} else {
		dclConfig = dcl.NewConfig(
			dclClientOptions,
			dclUserAgentOptions,
			dclLoggerOptions,
			dcl.WithBasePath(config.DataprocBasePath),
		)
	}

	return dataproc.NewClient(dclConfig)
}

func CreateEventarcClient(config *Config, userAgent, billingProject string) *eventarc.Client {
	dclClientOptions := dcl.WithHTTPClient(config.client)
	dclUserAgentOptions := dcl.WithUserAgent(userAgent)
	dclLoggerOptions := dcl.WithLogger(dclLogger{})
	var dclConfig *dcl.Config
	if config.UserProjectOverride && billingProject != "" {
		dclBillingProjectHeader := dcl.WithHeader("X-Goog-User-Project", billingProject)
		dclConfig = dcl.NewConfig(
			dclClientOptions,
			dclUserAgentOptions,
			dclLoggerOptions,
			dcl.WithBasePath(config.EventarcBasePath),
			dclBillingProjectHeader,
		)
	} else {
		dclConfig = dcl.NewConfig(
			dclClientOptions,
			dclUserAgentOptions,
			dclLoggerOptions,
			dcl.WithBasePath(config.EventarcBasePath),
		)
	}

	return eventarc.NewClient(dclConfig)
}
