# GetMsserverAdsitesSiteResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**Domain** | Pointer to **string** | The reference to the Active Directory Domain to which the site belongs. | [optional] 
**Name** | Pointer to **string** | The name of the site properties object for the Active Directory Sites. | [optional] 
**Networks** | Pointer to **[]string** | The list of networks to which the device interfaces belong. | [optional] 
**Result** | Pointer to [**MsserverAdsitesSite**](MsserverAdsitesSite.md) |  | [optional] 

## Methods

### NewGetMsserverAdsitesSiteResponse

`func NewGetMsserverAdsitesSiteResponse() *GetMsserverAdsitesSiteResponse`

NewGetMsserverAdsitesSiteResponse instantiates a new GetMsserverAdsitesSiteResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMsserverAdsitesSiteResponseWithDefaults

`func NewGetMsserverAdsitesSiteResponseWithDefaults() *GetMsserverAdsitesSiteResponse`

NewGetMsserverAdsitesSiteResponseWithDefaults instantiates a new GetMsserverAdsitesSiteResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetMsserverAdsitesSiteResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetMsserverAdsitesSiteResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetMsserverAdsitesSiteResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetMsserverAdsitesSiteResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetDomain

`func (o *GetMsserverAdsitesSiteResponse) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *GetMsserverAdsitesSiteResponse) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *GetMsserverAdsitesSiteResponse) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *GetMsserverAdsitesSiteResponse) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetName

`func (o *GetMsserverAdsitesSiteResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetMsserverAdsitesSiteResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetMsserverAdsitesSiteResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetMsserverAdsitesSiteResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetworks

`func (o *GetMsserverAdsitesSiteResponse) GetNetworks() []string`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *GetMsserverAdsitesSiteResponse) GetNetworksOk() (*[]string, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *GetMsserverAdsitesSiteResponse) SetNetworks(v []string)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *GetMsserverAdsitesSiteResponse) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.

### GetResult

`func (o *GetMsserverAdsitesSiteResponse) GetResult() MsserverAdsitesSite`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetMsserverAdsitesSiteResponse) GetResultOk() (*MsserverAdsitesSite, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetMsserverAdsitesSiteResponse) SetResult(v MsserverAdsitesSite)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetMsserverAdsitesSiteResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


