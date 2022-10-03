/*
 * Npcf_PolicyAuthorization Service API
 *
 * This is the Policy Authorization Service
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Npcf_PolicyAuthorization

import (
	"context"
	"strconv"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/free5gc/openapi"
	"golang.org/x/net/http2"
	"golang.org/x/oauth2/clientcredentials"
	"golang.org/x/oauth2"
	"github.com/free5gc/openapi/models"
)

// Linger please
var (
	_ context.Context
)

type EventsSubscriptionDocumentApiService service

/*
 EventsSubscriptionDocumentApiService deletes the Events Subscription subresource
  * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  * @param appSessionId string identifying the Individual Application Session Context resource
*/

func (a *EventsSubscriptionDocumentApiService) DeleteEventsSubsc(ctx context.Context, appSessionId string) (*http.Response, error) {
	var (
		localVarHTTPMethod	= strings.ToUpper("Delete")
		localVarPostBody	interface{}
		localVarFormFileName	string
		localVarFileName	string
		localVarFileBytes	[]byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath() + "/app-sessions/{appSessionId}/events-subscription"
	localVarPath = strings.Replace(localVarPath, "{"+"appSessionId"+"}", fmt.Sprintf("%v", appSessionId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarHTTPContentTypes := []string{"application/json"}

	localVarHeaderParams["Content-Type"] = localVarHTTPContentTypes[0]	// use the first content type specified in 'consumes'

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := openapi.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	scopes := []string{"npcf-policyauthorization",}
	additional_params, ok := ctx.Value(openapi.ContextOAuthAdditionalParams).(url.Values)
	if !ok {
		return nil, fmt.Errorf("OAuth parameters are invalid")
	}
	oauth, err := strconv.ParseBool(additional_params["OAuth"][0])
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	if oauth {
		tokenUrl := fmt.Sprintf("%v/oauth2/token", additional_params["NrfUri"][0])
		additional_params.Del("NrfUri")
		additional_params.Del("OAuth")
		additional_params.Add("targetNfType", "PCF")
		conf := &clientcredentials.Config{Scopes: scopes, TokenURL: tokenUrl, AuthStyle: oauth2.AuthStyleInParams, EndpointParams: additional_params}
		http_client := &http.Client{Transport: &http2.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}}
		ctx = context.WithValue(ctx, oauth2.HTTPClient, http_client)
		token, err := conf.Token(ctx)
		if err != nil {
			return nil, fmt.Errorf(err.Error())
		}
		ctx = context.WithValue(ctx, openapi.ContextAccessToken, token.AccessToken)
	}

	r, err := openapi.PrepareRequest(ctx, a.client.cfg, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := openapi.CallAPI(a.client.cfg, r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	apiError := openapi.GenericOpenAPIError{
		RawBody:	localVarBody,
		ErrorStatus:	localVarHTTPResponse.Status,
	}

	switch localVarHTTPResponse.StatusCode {
	case 204:
		return localVarHTTPResponse, nil
	case 400:
		var v models.ProblemDetails
		err = openapi.Deserialize(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarHTTPResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarHTTPResponse, apiError
	case 401:
		var v models.ProblemDetails
		err = openapi.Deserialize(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarHTTPResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarHTTPResponse, apiError
	case 403:
		var v models.ProblemDetails
		err = openapi.Deserialize(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarHTTPResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarHTTPResponse, apiError
	case 404:
		var v models.ProblemDetails
		err = openapi.Deserialize(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarHTTPResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarHTTPResponse, apiError
	case 429:
		var v models.ProblemDetails
		err = openapi.Deserialize(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarHTTPResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarHTTPResponse, apiError
	case 500:
		var v models.ProblemDetails
		err = openapi.Deserialize(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarHTTPResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarHTTPResponse, apiError
	case 503:
		var v models.ProblemDetails
		err = openapi.Deserialize(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarHTTPResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarHTTPResponse, apiError
	default:
		return localVarHTTPResponse, nil
	}
}

/*
 EventsSubscriptionDocumentApiService creates or modifies an Events Subscription subresource
  * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  * @param appSessionId string identifying the Events Subscription resource
  * @param eventsSubscReqData Creation or modification of an Events Subscription resource.
 @return models.UpdateEventsSubscResponse
*/

func (a *EventsSubscriptionDocumentApiService) UpdateEventsSubsc(ctx context.Context, appSessionId string, eventsSubscReqData models.EventsSubscReqData) (models.UpdateEventsSubscResponse, *http.Response, error) {
	var (
		localVarHTTPMethod	= strings.ToUpper("Put")
		localVarPostBody	interface{}
		localVarFormFileName	string
		localVarFileName	string
		localVarFileBytes	[]byte
		localVarReturnValue	models.UpdateEventsSubscResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath() + "/app-sessions/{appSessionId}/events-subscription"
	localVarPath = strings.Replace(localVarPath, "{"+"appSessionId"+"}", fmt.Sprintf("%v", appSessionId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarHTTPContentTypes := []string{"application/json"}

	localVarHeaderParams["Content-Type"] = localVarHTTPContentTypes[0]	// use the first content type specified in 'consumes'

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := openapi.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	// body params
	localVarPostBody = &eventsSubscReqData
	scopes := []string{"npcf-policyauthorization",}
	additional_params, ok := ctx.Value(openapi.ContextOAuthAdditionalParams).(url.Values)
	if !ok {
		return localVarReturnValue, nil, fmt.Errorf("OAuth parameters are invalid")
	}
	oauth, err := strconv.ParseBool(additional_params["OAuth"][0])
	if err != nil {
		return localVarReturnValue, nil, fmt.Errorf(err.Error())
	}
	if oauth {
		tokenUrl := fmt.Sprintf("%v/oauth2/token", additional_params["NrfUri"][0])
		additional_params.Del("NrfUri")
		additional_params.Del("OAuth")
		additional_params.Add("targetNfType", "PCF")
		conf := &clientcredentials.Config{Scopes: scopes, TokenURL: tokenUrl, AuthStyle: oauth2.AuthStyleInParams, EndpointParams: additional_params}
		http_client := &http.Client{Transport: &http2.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}}
		ctx = context.WithValue(ctx, oauth2.HTTPClient, http_client)
		token, err := conf.Token(ctx)
		if err != nil {
			return localVarReturnValue, nil, fmt.Errorf(err.Error())
		}
		ctx = context.WithValue(ctx, openapi.ContextAccessToken, token.AccessToken)
	}

	r, err := openapi.PrepareRequest(ctx, a.client.cfg, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := openapi.CallAPI(a.client.cfg, r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	apiError := openapi.GenericOpenAPIError{
		RawBody:	localVarBody,
		ErrorStatus:	localVarHTTPResponse.Status,
	}

	switch localVarHTTPResponse.StatusCode {
	case 201:
		err = openapi.Deserialize(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
		}
		return localVarReturnValue, localVarHTTPResponse, nil
	case 200:
		err = openapi.Deserialize(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
		}
		return localVarReturnValue, localVarHTTPResponse, nil
	case 204:
		return localVarReturnValue, localVarHTTPResponse, nil
	case 400:
		var v models.ProblemDetails
		err = openapi.Deserialize(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarReturnValue, localVarHTTPResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarReturnValue, localVarHTTPResponse, apiError
	case 401:
		var v models.ProblemDetails
		err = openapi.Deserialize(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarReturnValue, localVarHTTPResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarReturnValue, localVarHTTPResponse, apiError
	case 403:
		var v models.ProblemDetails
		err = openapi.Deserialize(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarReturnValue, localVarHTTPResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarReturnValue, localVarHTTPResponse, apiError
	case 404:
		var v models.ProblemDetails
		err = openapi.Deserialize(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarReturnValue, localVarHTTPResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarReturnValue, localVarHTTPResponse, apiError
	case 411:
		var v models.ProblemDetails
		err = openapi.Deserialize(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarReturnValue, localVarHTTPResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarReturnValue, localVarHTTPResponse, apiError
	case 413:
		var v models.ProblemDetails
		err = openapi.Deserialize(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarReturnValue, localVarHTTPResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarReturnValue, localVarHTTPResponse, apiError
	case 415:
		var v models.ProblemDetails
		err = openapi.Deserialize(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarReturnValue, localVarHTTPResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarReturnValue, localVarHTTPResponse, apiError
	case 429:
		var v models.ProblemDetails
		err = openapi.Deserialize(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarReturnValue, localVarHTTPResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarReturnValue, localVarHTTPResponse, apiError
	case 500:
		var v models.ProblemDetails
		err = openapi.Deserialize(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarReturnValue, localVarHTTPResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarReturnValue, localVarHTTPResponse, apiError
	case 503:
		var v models.ProblemDetails
		err = openapi.Deserialize(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarReturnValue, localVarHTTPResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarReturnValue, localVarHTTPResponse, apiError
	default:
		return localVarReturnValue, localVarHTTPResponse, nil
	}
}
