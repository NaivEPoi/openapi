/*
 * NSSF NS Selection
 *
 * NSSF Network Slice Selection Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nnssf_NSSelection

import (
	"context"
	"strconv"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/antihax/optional"
	"golang.org/x/net/http2"
	"golang.org/x/oauth2/clientcredentials"
	"golang.org/x/oauth2"

	"github.com/free5gc/openapi"
	. "github.com/free5gc/openapi/models"
)

// Linger please
var (
	_ context.Context
)

type NetworkSliceInformationDocumentApiService service

/*
NetworkSliceInformationDocumentApiService Retrieve the Network Slice Selection Information
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param nfType NF type of the NF service consumer
 * @param nfId NF Instance ID of the NF service consumer
 * @param optional nil or *NSSelectionGetParamOpts - Optional Parameters:
 * @param "SliceInfoRequestForRegistration" (optional.Interface of SliceInfoForRegistration) -  Requested network slice information during Registration procedure
 * @param "SliceInfoRequestForPduSession" (optional.Interface of SliceInfoForPduSession) -  Requested network slice information during PDU session establishment procedure
 * @param "HomePlmnId" (optional.Interface of PlmnId) -  PLMN ID of the HPLMN
 * @param "Tai" (optional.Interface of Tai) -  TAI of the UE
 * @param "SupportedFeatures" (optional.String) -  Features required to be supported by the NFs in the target slice instance
@return AuthorizedNetworkSliceInfo
*/

type NSSelectionGetParamOpts struct {
	SliceInfoRequestForRegistration	optional.Interface
	SliceInfoRequestForPduSession	optional.Interface
	HomePlmnId			optional.Interface
	Tai				optional.Interface
	SupportedFeatures		optional.String
}

func (a *NetworkSliceInformationDocumentApiService) NSSelectionGet(ctx context.Context, nfType NfType, nfId string, localVarOptionals *NSSelectionGetParamOpts) (AuthorizedNetworkSliceInfo, *http.Response, error) {
	var (
		localVarHTTPMethod	= strings.ToUpper("Get")
		localVarPostBody	interface{}
		localVarFormFileName	string
		localVarFileName	string
		localVarFileBytes	[]byte
		localVarReturnValue	AuthorizedNetworkSliceInfo
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath() + "/network-slice-information"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("nf-type", openapi.ParameterToString(nfType, ""))
	localVarQueryParams.Add("nf-id", openapi.ParameterToString(nfId, ""))
	if localVarOptionals != nil && localVarOptionals.SliceInfoRequestForRegistration.IsSet() {
		localVarQueryParams.Add("slice-info-request-for-registration", openapi.ParameterToString(localVarOptionals.SliceInfoRequestForRegistration.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SliceInfoRequestForPduSession.IsSet() {
		localVarQueryParams.Add("slice-info-request-for-pdu-session", openapi.ParameterToString(localVarOptionals.SliceInfoRequestForPduSession.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.HomePlmnId.IsSet() {
		localVarQueryParams.Add("home-plmn-id", openapi.ParameterToString(localVarOptionals.HomePlmnId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Tai.IsSet() {
		localVarQueryParams.Add("tai", openapi.ParameterToString(localVarOptionals.Tai.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SupportedFeatures.IsSet() && localVarOptionals.SupportedFeatures.Value() != "" {
		localVarQueryParams.Add("supported-features", openapi.ParameterToString(localVarOptionals.SupportedFeatures.Value(), ""))
	}

	localVarHTTPContentTypes := []string{"application/json"}

	localVarHeaderParams["Content-Type"] = localVarHTTPContentTypes[0]	// use the first content type specified in 'consumes'

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := openapi.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	scopes := []string{"nnssf-nsselection",}
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
		additional_params.Add("targetNfType", "NSSF")
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
	case 200:
		err = openapi.Deserialize(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
		}
		return localVarReturnValue, localVarHTTPResponse, nil
	case 400:
		var v ProblemDetails
		err = openapi.Deserialize(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarReturnValue, localVarHTTPResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarReturnValue, localVarHTTPResponse, apiError
	case 401:
		var v ProblemDetails
		err = openapi.Deserialize(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarReturnValue, localVarHTTPResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarReturnValue, localVarHTTPResponse, apiError
	case 403:
		var v ProblemDetails
		err = openapi.Deserialize(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarReturnValue, localVarHTTPResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarReturnValue, localVarHTTPResponse, apiError
	case 404:
		var v ProblemDetails
		err = openapi.Deserialize(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarReturnValue, localVarHTTPResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarReturnValue, localVarHTTPResponse, apiError
	case 406:
		return localVarReturnValue, localVarHTTPResponse, nil
	case 414:
		var v ProblemDetails
		err = openapi.Deserialize(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarReturnValue, localVarHTTPResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarReturnValue, localVarHTTPResponse, apiError
	case 429:
		var v ProblemDetails
		err = openapi.Deserialize(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarReturnValue, localVarHTTPResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarReturnValue, localVarHTTPResponse, apiError
	case 500:
		var v ProblemDetails
		err = openapi.Deserialize(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarReturnValue, localVarHTTPResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarReturnValue, localVarHTTPResponse, apiError
	case 503:
		var v ProblemDetails
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
