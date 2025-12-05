# GetDtcServerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**AutoCreateHostRecord** | Pointer to **bool** | Enabling this option will auto-create a single read-only A/AAAA/CNAME record corresponding to the configured hostname and update it if the hostname changes. | [optional] 
**Comment** | Pointer to **string** | Comment for the DTC Server; maximum 256 characters. | [optional] 
**Disable** | Pointer to **bool** | Determines whether the DTC Server is disabled or not. When this is set to False, the fixed address is enabled. | [optional] 
**ExtAttrsPlus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrsMinus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**Health** | Pointer to [**DtcServerHealth**](DtcServerHealth.md) |  | [optional] 
**Host** | Pointer to **string** | The address or FQDN of the server. | [optional] 
**Monitors** | Pointer to [**[]DtcServerMonitors**](DtcServerMonitors.md) | List of IP/FQDN and monitor pairs to be used for additional monitoring. | [optional] 
**Name** | Pointer to **string** | The DTC Server display name. | [optional] 
**SniHostname** | Pointer to **string** | The hostname for Server Name Indication (SNI) in FQDN format. | [optional] 
**UseSniHostname** | Pointer to **bool** | Use flag for: sni_hostname | [optional] 
**Uuid** | Pointer to **string** | Universally Unique ID assigned for this object | [optional] [readonly] 
**Result** | Pointer to [**DtcServer**](DtcServer.md) |  | [optional] 

## Methods

### NewGetDtcServerResponse

`func NewGetDtcServerResponse() *GetDtcServerResponse`

NewGetDtcServerResponse instantiates a new GetDtcServerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDtcServerResponseWithDefaults

`func NewGetDtcServerResponseWithDefaults() *GetDtcServerResponse`

NewGetDtcServerResponseWithDefaults instantiates a new GetDtcServerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetDtcServerResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetDtcServerResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetDtcServerResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetDtcServerResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAutoCreateHostRecord

`func (o *GetDtcServerResponse) GetAutoCreateHostRecord() bool`

GetAutoCreateHostRecord returns the AutoCreateHostRecord field if non-nil, zero value otherwise.

### GetAutoCreateHostRecordOk

`func (o *GetDtcServerResponse) GetAutoCreateHostRecordOk() (*bool, bool)`

GetAutoCreateHostRecordOk returns a tuple with the AutoCreateHostRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCreateHostRecord

`func (o *GetDtcServerResponse) SetAutoCreateHostRecord(v bool)`

SetAutoCreateHostRecord sets AutoCreateHostRecord field to given value.

### HasAutoCreateHostRecord

`func (o *GetDtcServerResponse) HasAutoCreateHostRecord() bool`

HasAutoCreateHostRecord returns a boolean if a field has been set.

### GetComment

`func (o *GetDtcServerResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetDtcServerResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetDtcServerResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetDtcServerResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDisable

`func (o *GetDtcServerResponse) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *GetDtcServerResponse) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *GetDtcServerResponse) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *GetDtcServerResponse) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetExtAttrsPlus

`func (o *GetDtcServerResponse) GetExtAttrsPlus() map[string]ExtAttrs`

GetExtAttrsPlus returns the ExtAttrsPlus field if non-nil, zero value otherwise.

### GetExtAttrsPlusOk

`func (o *GetDtcServerResponse) GetExtAttrsPlusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsPlusOk returns a tuple with the ExtAttrsPlus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsPlus

`func (o *GetDtcServerResponse) SetExtAttrsPlus(v map[string]ExtAttrs)`

SetExtAttrsPlus sets ExtAttrsPlus field to given value.

### HasExtAttrsPlus

`func (o *GetDtcServerResponse) HasExtAttrsPlus() bool`

HasExtAttrsPlus returns a boolean if a field has been set.

### GetExtAttrsMinus

`func (o *GetDtcServerResponse) GetExtAttrsMinus() map[string]ExtAttrs`

GetExtAttrsMinus returns the ExtAttrsMinus field if non-nil, zero value otherwise.

### GetExtAttrsMinusOk

`func (o *GetDtcServerResponse) GetExtAttrsMinusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsMinusOk returns a tuple with the ExtAttrsMinus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsMinus

`func (o *GetDtcServerResponse) SetExtAttrsMinus(v map[string]ExtAttrs)`

SetExtAttrsMinus sets ExtAttrsMinus field to given value.

### HasExtAttrsMinus

`func (o *GetDtcServerResponse) HasExtAttrsMinus() bool`

HasExtAttrsMinus returns a boolean if a field has been set.

### GetExtAttrs

`func (o *GetDtcServerResponse) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *GetDtcServerResponse) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *GetDtcServerResponse) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *GetDtcServerResponse) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetHealth

`func (o *GetDtcServerResponse) GetHealth() DtcServerHealth`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *GetDtcServerResponse) GetHealthOk() (*DtcServerHealth, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *GetDtcServerResponse) SetHealth(v DtcServerHealth)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *GetDtcServerResponse) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetHost

`func (o *GetDtcServerResponse) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *GetDtcServerResponse) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *GetDtcServerResponse) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *GetDtcServerResponse) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetMonitors

`func (o *GetDtcServerResponse) GetMonitors() []DtcServerMonitors`

GetMonitors returns the Monitors field if non-nil, zero value otherwise.

### GetMonitorsOk

`func (o *GetDtcServerResponse) GetMonitorsOk() (*[]DtcServerMonitors, bool)`

GetMonitorsOk returns a tuple with the Monitors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitors

`func (o *GetDtcServerResponse) SetMonitors(v []DtcServerMonitors)`

SetMonitors sets Monitors field to given value.

### HasMonitors

`func (o *GetDtcServerResponse) HasMonitors() bool`

HasMonitors returns a boolean if a field has been set.

### GetName

`func (o *GetDtcServerResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetDtcServerResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetDtcServerResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetDtcServerResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSniHostname

`func (o *GetDtcServerResponse) GetSniHostname() string`

GetSniHostname returns the SniHostname field if non-nil, zero value otherwise.

### GetSniHostnameOk

`func (o *GetDtcServerResponse) GetSniHostnameOk() (*string, bool)`

GetSniHostnameOk returns a tuple with the SniHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSniHostname

`func (o *GetDtcServerResponse) SetSniHostname(v string)`

SetSniHostname sets SniHostname field to given value.

### HasSniHostname

`func (o *GetDtcServerResponse) HasSniHostname() bool`

HasSniHostname returns a boolean if a field has been set.

### GetUseSniHostname

`func (o *GetDtcServerResponse) GetUseSniHostname() bool`

GetUseSniHostname returns the UseSniHostname field if non-nil, zero value otherwise.

### GetUseSniHostnameOk

`func (o *GetDtcServerResponse) GetUseSniHostnameOk() (*bool, bool)`

GetUseSniHostnameOk returns a tuple with the UseSniHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSniHostname

`func (o *GetDtcServerResponse) SetUseSniHostname(v bool)`

SetUseSniHostname sets UseSniHostname field to given value.

### HasUseSniHostname

`func (o *GetDtcServerResponse) HasUseSniHostname() bool`

HasUseSniHostname returns a boolean if a field has been set.

### GetUuid

`func (o *GetDtcServerResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *GetDtcServerResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *GetDtcServerResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *GetDtcServerResponse) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetResult

`func (o *GetDtcServerResponse) GetResult() DtcServer`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetDtcServerResponse) GetResultOk() (*DtcServer, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetDtcServerResponse) SetResult(v DtcServer)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetDtcServerResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


