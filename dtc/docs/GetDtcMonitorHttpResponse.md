# GetDtcMonitorHttpResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**Ciphers** | Pointer to **string** | An optional cipher list for a secure HTTP/S connection. | [optional] 
**ClientCert** | Pointer to **string** | An optional client certificate, supplied in a secure HTTP/S mode if present. | [optional] 
**Comment** | Pointer to **string** | Comment for this DTC monitor; maximum 256 characters. | [optional] 
**ContentCheck** | Pointer to **string** | The content check type. | [optional] 
**ContentCheckInput** | Pointer to **string** | A portion of response to use as input for content check. | [optional] 
**ContentCheckOp** | Pointer to **string** | A content check success criteria operator. | [optional] 
**ContentCheckRegex** | Pointer to **string** | A content check regular expression. | [optional] 
**ContentExtractGroup** | Pointer to **int64** | A content extraction sub-expression to extract. | [optional] 
**ContentExtractType** | Pointer to **string** | A content extraction expected type for the extracted data. | [optional] 
**ContentExtractValue** | Pointer to **string** | A content extraction value to compare with extracted result. | [optional] 
**EnableSni** | Pointer to **bool** | Determines whether the Server Name Indication (SNI) for HTTPS monitor is enabled. | [optional] 
**ExtAttrsPlus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrsMinus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**Interval** | Pointer to **int64** | The interval for TCP health check. | [optional] 
**Name** | Pointer to **string** | The display name for this DTC monitor. | [optional] 
**Port** | Pointer to **int64** | Port for TCP requests. | [optional] 
**Request** | Pointer to **string** | An HTTP request to send. | [optional] 
**Result** | Pointer to [**DtcMonitorHttp**](DtcMonitorHttp.md) |  | [optional] 
**ResultCode** | Pointer to **int64** | The expected return code. | [optional] 
**RetryDown** | Pointer to **int64** | The value of how many times the server should appear as down to be treated as dead after it was alive. | [optional] 
**RetryUp** | Pointer to **int64** | The value of how many times the server should appear as up to be treated as alive after it was dead. | [optional] 
**Secure** | Pointer to **bool** | The connection security status. | [optional] 
**Timeout** | Pointer to **int64** | The timeout for TCP health check in seconds. | [optional] 
**ValidateCert** | Pointer to **bool** | Determines whether the validation of the remote server&#39;s certificate is enabled. | [optional] 

## Methods

### NewGetDtcMonitorHttpResponse

`func NewGetDtcMonitorHttpResponse() *GetDtcMonitorHttpResponse`

NewGetDtcMonitorHttpResponse instantiates a new GetDtcMonitorHttpResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDtcMonitorHttpResponseWithDefaults

`func NewGetDtcMonitorHttpResponseWithDefaults() *GetDtcMonitorHttpResponse`

NewGetDtcMonitorHttpResponseWithDefaults instantiates a new GetDtcMonitorHttpResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetDtcMonitorHttpResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetDtcMonitorHttpResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetDtcMonitorHttpResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetDtcMonitorHttpResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetCiphers

`func (o *GetDtcMonitorHttpResponse) GetCiphers() string`

GetCiphers returns the Ciphers field if non-nil, zero value otherwise.

### GetCiphersOk

`func (o *GetDtcMonitorHttpResponse) GetCiphersOk() (*string, bool)`

GetCiphersOk returns a tuple with the Ciphers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiphers

`func (o *GetDtcMonitorHttpResponse) SetCiphers(v string)`

SetCiphers sets Ciphers field to given value.

### HasCiphers

`func (o *GetDtcMonitorHttpResponse) HasCiphers() bool`

HasCiphers returns a boolean if a field has been set.

### GetClientCert

`func (o *GetDtcMonitorHttpResponse) GetClientCert() string`

GetClientCert returns the ClientCert field if non-nil, zero value otherwise.

### GetClientCertOk

`func (o *GetDtcMonitorHttpResponse) GetClientCertOk() (*string, bool)`

GetClientCertOk returns a tuple with the ClientCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCert

`func (o *GetDtcMonitorHttpResponse) SetClientCert(v string)`

SetClientCert sets ClientCert field to given value.

### HasClientCert

`func (o *GetDtcMonitorHttpResponse) HasClientCert() bool`

HasClientCert returns a boolean if a field has been set.

### GetComment

`func (o *GetDtcMonitorHttpResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetDtcMonitorHttpResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetDtcMonitorHttpResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetDtcMonitorHttpResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetContentCheck

`func (o *GetDtcMonitorHttpResponse) GetContentCheck() string`

GetContentCheck returns the ContentCheck field if non-nil, zero value otherwise.

### GetContentCheckOk

`func (o *GetDtcMonitorHttpResponse) GetContentCheckOk() (*string, bool)`

GetContentCheckOk returns a tuple with the ContentCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentCheck

`func (o *GetDtcMonitorHttpResponse) SetContentCheck(v string)`

SetContentCheck sets ContentCheck field to given value.

### HasContentCheck

`func (o *GetDtcMonitorHttpResponse) HasContentCheck() bool`

HasContentCheck returns a boolean if a field has been set.

### GetContentCheckInput

`func (o *GetDtcMonitorHttpResponse) GetContentCheckInput() string`

GetContentCheckInput returns the ContentCheckInput field if non-nil, zero value otherwise.

### GetContentCheckInputOk

`func (o *GetDtcMonitorHttpResponse) GetContentCheckInputOk() (*string, bool)`

GetContentCheckInputOk returns a tuple with the ContentCheckInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentCheckInput

`func (o *GetDtcMonitorHttpResponse) SetContentCheckInput(v string)`

SetContentCheckInput sets ContentCheckInput field to given value.

### HasContentCheckInput

`func (o *GetDtcMonitorHttpResponse) HasContentCheckInput() bool`

HasContentCheckInput returns a boolean if a field has been set.

### GetContentCheckOp

`func (o *GetDtcMonitorHttpResponse) GetContentCheckOp() string`

GetContentCheckOp returns the ContentCheckOp field if non-nil, zero value otherwise.

### GetContentCheckOpOk

`func (o *GetDtcMonitorHttpResponse) GetContentCheckOpOk() (*string, bool)`

GetContentCheckOpOk returns a tuple with the ContentCheckOp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentCheckOp

`func (o *GetDtcMonitorHttpResponse) SetContentCheckOp(v string)`

SetContentCheckOp sets ContentCheckOp field to given value.

### HasContentCheckOp

`func (o *GetDtcMonitorHttpResponse) HasContentCheckOp() bool`

HasContentCheckOp returns a boolean if a field has been set.

### GetContentCheckRegex

`func (o *GetDtcMonitorHttpResponse) GetContentCheckRegex() string`

GetContentCheckRegex returns the ContentCheckRegex field if non-nil, zero value otherwise.

### GetContentCheckRegexOk

`func (o *GetDtcMonitorHttpResponse) GetContentCheckRegexOk() (*string, bool)`

GetContentCheckRegexOk returns a tuple with the ContentCheckRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentCheckRegex

`func (o *GetDtcMonitorHttpResponse) SetContentCheckRegex(v string)`

SetContentCheckRegex sets ContentCheckRegex field to given value.

### HasContentCheckRegex

`func (o *GetDtcMonitorHttpResponse) HasContentCheckRegex() bool`

HasContentCheckRegex returns a boolean if a field has been set.

### GetContentExtractGroup

`func (o *GetDtcMonitorHttpResponse) GetContentExtractGroup() int64`

GetContentExtractGroup returns the ContentExtractGroup field if non-nil, zero value otherwise.

### GetContentExtractGroupOk

`func (o *GetDtcMonitorHttpResponse) GetContentExtractGroupOk() (*int64, bool)`

GetContentExtractGroupOk returns a tuple with the ContentExtractGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentExtractGroup

`func (o *GetDtcMonitorHttpResponse) SetContentExtractGroup(v int64)`

SetContentExtractGroup sets ContentExtractGroup field to given value.

### HasContentExtractGroup

`func (o *GetDtcMonitorHttpResponse) HasContentExtractGroup() bool`

HasContentExtractGroup returns a boolean if a field has been set.

### GetContentExtractType

`func (o *GetDtcMonitorHttpResponse) GetContentExtractType() string`

GetContentExtractType returns the ContentExtractType field if non-nil, zero value otherwise.

### GetContentExtractTypeOk

`func (o *GetDtcMonitorHttpResponse) GetContentExtractTypeOk() (*string, bool)`

GetContentExtractTypeOk returns a tuple with the ContentExtractType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentExtractType

`func (o *GetDtcMonitorHttpResponse) SetContentExtractType(v string)`

SetContentExtractType sets ContentExtractType field to given value.

### HasContentExtractType

`func (o *GetDtcMonitorHttpResponse) HasContentExtractType() bool`

HasContentExtractType returns a boolean if a field has been set.

### GetContentExtractValue

`func (o *GetDtcMonitorHttpResponse) GetContentExtractValue() string`

GetContentExtractValue returns the ContentExtractValue field if non-nil, zero value otherwise.

### GetContentExtractValueOk

`func (o *GetDtcMonitorHttpResponse) GetContentExtractValueOk() (*string, bool)`

GetContentExtractValueOk returns a tuple with the ContentExtractValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentExtractValue

`func (o *GetDtcMonitorHttpResponse) SetContentExtractValue(v string)`

SetContentExtractValue sets ContentExtractValue field to given value.

### HasContentExtractValue

`func (o *GetDtcMonitorHttpResponse) HasContentExtractValue() bool`

HasContentExtractValue returns a boolean if a field has been set.

### GetEnableSni

`func (o *GetDtcMonitorHttpResponse) GetEnableSni() bool`

GetEnableSni returns the EnableSni field if non-nil, zero value otherwise.

### GetEnableSniOk

`func (o *GetDtcMonitorHttpResponse) GetEnableSniOk() (*bool, bool)`

GetEnableSniOk returns a tuple with the EnableSni field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSni

`func (o *GetDtcMonitorHttpResponse) SetEnableSni(v bool)`

SetEnableSni sets EnableSni field to given value.

### HasEnableSni

`func (o *GetDtcMonitorHttpResponse) HasEnableSni() bool`

HasEnableSni returns a boolean if a field has been set.

### GetExtAttrsPlus

`func (o *GetDtcMonitorHttpResponse) GetExtAttrsPlus() map[string]ExtAttrs`

GetExtAttrsPlus returns the ExtAttrsPlus field if non-nil, zero value otherwise.

### GetExtAttrsPlusOk

`func (o *GetDtcMonitorHttpResponse) GetExtAttrsPlusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsPlusOk returns a tuple with the ExtAttrsPlus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsPlus

`func (o *GetDtcMonitorHttpResponse) SetExtAttrsPlus(v map[string]ExtAttrs)`

SetExtAttrsPlus sets ExtAttrsPlus field to given value.

### HasExtAttrsPlus

`func (o *GetDtcMonitorHttpResponse) HasExtAttrsPlus() bool`

HasExtAttrsPlus returns a boolean if a field has been set.

### GetExtAttrsMinus

`func (o *GetDtcMonitorHttpResponse) GetExtAttrsMinus() map[string]ExtAttrs`

GetExtAttrsMinus returns the ExtAttrsMinus field if non-nil, zero value otherwise.

### GetExtAttrsMinusOk

`func (o *GetDtcMonitorHttpResponse) GetExtAttrsMinusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsMinusOk returns a tuple with the ExtAttrsMinus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsMinus

`func (o *GetDtcMonitorHttpResponse) SetExtAttrsMinus(v map[string]ExtAttrs)`

SetExtAttrsMinus sets ExtAttrsMinus field to given value.

### HasExtAttrsMinus

`func (o *GetDtcMonitorHttpResponse) HasExtAttrsMinus() bool`

HasExtAttrsMinus returns a boolean if a field has been set.

### GetExtAttrs

`func (o *GetDtcMonitorHttpResponse) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *GetDtcMonitorHttpResponse) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *GetDtcMonitorHttpResponse) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *GetDtcMonitorHttpResponse) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetInterval

`func (o *GetDtcMonitorHttpResponse) GetInterval() int64`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *GetDtcMonitorHttpResponse) GetIntervalOk() (*int64, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *GetDtcMonitorHttpResponse) SetInterval(v int64)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *GetDtcMonitorHttpResponse) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetName

`func (o *GetDtcMonitorHttpResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetDtcMonitorHttpResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetDtcMonitorHttpResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetDtcMonitorHttpResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPort

`func (o *GetDtcMonitorHttpResponse) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *GetDtcMonitorHttpResponse) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *GetDtcMonitorHttpResponse) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *GetDtcMonitorHttpResponse) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetRequest

`func (o *GetDtcMonitorHttpResponse) GetRequest() string`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *GetDtcMonitorHttpResponse) GetRequestOk() (*string, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *GetDtcMonitorHttpResponse) SetRequest(v string)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *GetDtcMonitorHttpResponse) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetResult

`func (o *GetDtcMonitorHttpResponse) GetResult() DtcMonitorHttp`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetDtcMonitorHttpResponse) GetResultOk() (*DtcMonitorHttp, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetDtcMonitorHttpResponse) SetResult(v DtcMonitorHttp)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetDtcMonitorHttpResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetResultCode

`func (o *GetDtcMonitorHttpResponse) GetResultCode() int64`

GetResultCode returns the ResultCode field if non-nil, zero value otherwise.

### GetResultCodeOk

`func (o *GetDtcMonitorHttpResponse) GetResultCodeOk() (*int64, bool)`

GetResultCodeOk returns a tuple with the ResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultCode

`func (o *GetDtcMonitorHttpResponse) SetResultCode(v int64)`

SetResultCode sets ResultCode field to given value.

### HasResultCode

`func (o *GetDtcMonitorHttpResponse) HasResultCode() bool`

HasResultCode returns a boolean if a field has been set.

### GetRetryDown

`func (o *GetDtcMonitorHttpResponse) GetRetryDown() int64`

GetRetryDown returns the RetryDown field if non-nil, zero value otherwise.

### GetRetryDownOk

`func (o *GetDtcMonitorHttpResponse) GetRetryDownOk() (*int64, bool)`

GetRetryDownOk returns a tuple with the RetryDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryDown

`func (o *GetDtcMonitorHttpResponse) SetRetryDown(v int64)`

SetRetryDown sets RetryDown field to given value.

### HasRetryDown

`func (o *GetDtcMonitorHttpResponse) HasRetryDown() bool`

HasRetryDown returns a boolean if a field has been set.

### GetRetryUp

`func (o *GetDtcMonitorHttpResponse) GetRetryUp() int64`

GetRetryUp returns the RetryUp field if non-nil, zero value otherwise.

### GetRetryUpOk

`func (o *GetDtcMonitorHttpResponse) GetRetryUpOk() (*int64, bool)`

GetRetryUpOk returns a tuple with the RetryUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryUp

`func (o *GetDtcMonitorHttpResponse) SetRetryUp(v int64)`

SetRetryUp sets RetryUp field to given value.

### HasRetryUp

`func (o *GetDtcMonitorHttpResponse) HasRetryUp() bool`

HasRetryUp returns a boolean if a field has been set.

### GetSecure

`func (o *GetDtcMonitorHttpResponse) GetSecure() bool`

GetSecure returns the Secure field if non-nil, zero value otherwise.

### GetSecureOk

`func (o *GetDtcMonitorHttpResponse) GetSecureOk() (*bool, bool)`

GetSecureOk returns a tuple with the Secure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecure

`func (o *GetDtcMonitorHttpResponse) SetSecure(v bool)`

SetSecure sets Secure field to given value.

### HasSecure

`func (o *GetDtcMonitorHttpResponse) HasSecure() bool`

HasSecure returns a boolean if a field has been set.

### GetTimeout

`func (o *GetDtcMonitorHttpResponse) GetTimeout() int64`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *GetDtcMonitorHttpResponse) GetTimeoutOk() (*int64, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *GetDtcMonitorHttpResponse) SetTimeout(v int64)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *GetDtcMonitorHttpResponse) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetValidateCert

`func (o *GetDtcMonitorHttpResponse) GetValidateCert() bool`

GetValidateCert returns the ValidateCert field if non-nil, zero value otherwise.

### GetValidateCertOk

`func (o *GetDtcMonitorHttpResponse) GetValidateCertOk() (*bool, bool)`

GetValidateCertOk returns a tuple with the ValidateCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidateCert

`func (o *GetDtcMonitorHttpResponse) SetValidateCert(v bool)`

SetValidateCert sets ValidateCert field to given value.

### HasValidateCert

`func (o *GetDtcMonitorHttpResponse) HasValidateCert() bool`

HasValidateCert returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


