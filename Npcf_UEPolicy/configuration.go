/*
 * Npcf_UEPolicyControl
 *
 * UE Policy Control Service API
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Npcf_UEPolicy

import (
	"net/http"
	"strings"
)

// contextKeys are used to identify the type of value in the context.
// Since these are string, it is possible to get a short description of the
// context key for logging and debugging using key.String().

type contextKey string

func (c contextKey) String() string {
	return "auth " + string(c)
}

var (
	// ContextOAuth2 takes an oauth2.TokenSource as authentication for the request.
	ContextOAuth2	= contextKey("token")

	// ContextBasicAuth takes BasicAuth as authentication for the request.
	ContextBasicAuth	= contextKey("basic")

	// ContextAccessToken takes a string oauth2 access token as authentication for the request.
	ContextAccessToken	= contextKey("accesstoken")

	// ContextAPIKey takes an APIKey as authentication for the request
	ContextAPIKey	= contextKey("apikey")
)

// BasicAuth provides basic http authentication to a request passed via context using ContextBasicAuth
type BasicAuth struct {
	UserName	string	`json:"userName,omitempty"`
	Password	string	`json:"password,omitempty"`
}

// APIKey provides API key based authentication to a request passed via context using ContextAPIKey
type APIKey struct {
	Key	string
	Prefix	string
}

type Configuration struct {
	Scheme		string	`json:"scheme,omitempty"`
	url		string
	basePath	string
	host		string
	defaultHeader	map[string]string
	userAgent	string
	httpClient	*http.Client
}

func NewConfiguration() *Configuration {
	cfg := &Configuration{
		basePath:	"https://example.com/",
		url:		"{apiRoot}",
		defaultHeader:	make(map[string]string),
		userAgent:	"OpenAPI-Generator/1.0.0/go",
	}
	return cfg
}

func (c *Configuration) SetBasePath(apiRoot string) {
	url := c.url

	// Replace apiRoot
	url = strings.Replace(url, "{"+"apiRoot"+"}", apiRoot, -1)

	c.basePath = url
}

func (c *Configuration) BasePath() string {
	return c.basePath
}

func (c *Configuration) Host() string {
	return c.host
}

func (c *Configuration) SetHost(host string) {
	c.host = host
}

func (c *Configuration) UserAgent() string {
	return c.userAgent
}

func (c *Configuration) SetUserAgent(userAgent string) {
	c.userAgent = userAgent
}

func (c *Configuration) DefaultHeader() map[string]string {
	return c.defaultHeader
}

func (c *Configuration) AddDefaultHeader(key string, value string) {
	c.defaultHeader[key] = value
}

func (c *Configuration) HTTPClient() *http.Client {
	return c.httpClient
}
