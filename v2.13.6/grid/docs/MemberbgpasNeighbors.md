# MemberbgpasNeighbors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Interface** | Pointer to **string** | The interface that sends BGP advertisement information. | [optional] 
**NeighborIp** | Pointer to **string** | The IP address of the BGP neighbor. | [optional] 
**RemoteAs** | Pointer to **int64** | The remote AS number of the BGP neighbor. | [optional] 
**AuthenticationMode** | Pointer to **string** | The BGP authentication mode. | [optional] 
**BgpNeighborPass** | Pointer to **string** | The password for a BGP neighbor. This is required only if authentication_mode is set to \&quot;MD5\&quot;. When the password is entered, the value is preserved even if authentication_mode is changed to \&quot;NONE\&quot;. This is a write-only attribute. | [optional] 
**Comment** | Pointer to **string** | User comments for this BGP neighbor. | [optional] 
**Multihop** | Pointer to **bool** | Determines if the multi-hop support is enabled or not. | [optional] 
**MultihopTtl** | Pointer to **int64** | The Time To Live (TTL) value for multi-hop. Valid values are between 1 and 255. | [optional] 
**BfdTemplate** | Pointer to **string** | The BFD template name. | [optional] 
**EnableBfd** | Pointer to **bool** | Determines if BFD is enabled or not. | [optional] 
**EnableBfdDnscheck** | Pointer to **bool** | Determines if BFD internal DNS monitor is enabled or not. | [optional] 

## Methods

### NewMemberbgpasNeighbors

`func NewMemberbgpasNeighbors() *MemberbgpasNeighbors`

NewMemberbgpasNeighbors instantiates a new MemberbgpasNeighbors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberbgpasNeighborsWithDefaults

`func NewMemberbgpasNeighborsWithDefaults() *MemberbgpasNeighbors`

NewMemberbgpasNeighborsWithDefaults instantiates a new MemberbgpasNeighbors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterface

`func (o *MemberbgpasNeighbors) GetInterface() string`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *MemberbgpasNeighbors) GetInterfaceOk() (*string, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *MemberbgpasNeighbors) SetInterface(v string)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *MemberbgpasNeighbors) HasInterface() bool`

HasInterface returns a boolean if a field has been set.

### GetNeighborIp

`func (o *MemberbgpasNeighbors) GetNeighborIp() string`

GetNeighborIp returns the NeighborIp field if non-nil, zero value otherwise.

### GetNeighborIpOk

`func (o *MemberbgpasNeighbors) GetNeighborIpOk() (*string, bool)`

GetNeighborIpOk returns a tuple with the NeighborIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeighborIp

`func (o *MemberbgpasNeighbors) SetNeighborIp(v string)`

SetNeighborIp sets NeighborIp field to given value.

### HasNeighborIp

`func (o *MemberbgpasNeighbors) HasNeighborIp() bool`

HasNeighborIp returns a boolean if a field has been set.

### GetRemoteAs

`func (o *MemberbgpasNeighbors) GetRemoteAs() int64`

GetRemoteAs returns the RemoteAs field if non-nil, zero value otherwise.

### GetRemoteAsOk

`func (o *MemberbgpasNeighbors) GetRemoteAsOk() (*int64, bool)`

GetRemoteAsOk returns a tuple with the RemoteAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAs

`func (o *MemberbgpasNeighbors) SetRemoteAs(v int64)`

SetRemoteAs sets RemoteAs field to given value.

### HasRemoteAs

`func (o *MemberbgpasNeighbors) HasRemoteAs() bool`

HasRemoteAs returns a boolean if a field has been set.

### GetAuthenticationMode

`func (o *MemberbgpasNeighbors) GetAuthenticationMode() string`

GetAuthenticationMode returns the AuthenticationMode field if non-nil, zero value otherwise.

### GetAuthenticationModeOk

`func (o *MemberbgpasNeighbors) GetAuthenticationModeOk() (*string, bool)`

GetAuthenticationModeOk returns a tuple with the AuthenticationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationMode

`func (o *MemberbgpasNeighbors) SetAuthenticationMode(v string)`

SetAuthenticationMode sets AuthenticationMode field to given value.

### HasAuthenticationMode

`func (o *MemberbgpasNeighbors) HasAuthenticationMode() bool`

HasAuthenticationMode returns a boolean if a field has been set.

### GetBgpNeighborPass

`func (o *MemberbgpasNeighbors) GetBgpNeighborPass() string`

GetBgpNeighborPass returns the BgpNeighborPass field if non-nil, zero value otherwise.

### GetBgpNeighborPassOk

`func (o *MemberbgpasNeighbors) GetBgpNeighborPassOk() (*string, bool)`

GetBgpNeighborPassOk returns a tuple with the BgpNeighborPass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpNeighborPass

`func (o *MemberbgpasNeighbors) SetBgpNeighborPass(v string)`

SetBgpNeighborPass sets BgpNeighborPass field to given value.

### HasBgpNeighborPass

`func (o *MemberbgpasNeighbors) HasBgpNeighborPass() bool`

HasBgpNeighborPass returns a boolean if a field has been set.

### GetComment

`func (o *MemberbgpasNeighbors) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *MemberbgpasNeighbors) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *MemberbgpasNeighbors) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *MemberbgpasNeighbors) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetMultihop

`func (o *MemberbgpasNeighbors) GetMultihop() bool`

GetMultihop returns the Multihop field if non-nil, zero value otherwise.

### GetMultihopOk

`func (o *MemberbgpasNeighbors) GetMultihopOk() (*bool, bool)`

GetMultihopOk returns a tuple with the Multihop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultihop

`func (o *MemberbgpasNeighbors) SetMultihop(v bool)`

SetMultihop sets Multihop field to given value.

### HasMultihop

`func (o *MemberbgpasNeighbors) HasMultihop() bool`

HasMultihop returns a boolean if a field has been set.

### GetMultihopTtl

`func (o *MemberbgpasNeighbors) GetMultihopTtl() int64`

GetMultihopTtl returns the MultihopTtl field if non-nil, zero value otherwise.

### GetMultihopTtlOk

`func (o *MemberbgpasNeighbors) GetMultihopTtlOk() (*int64, bool)`

GetMultihopTtlOk returns a tuple with the MultihopTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultihopTtl

`func (o *MemberbgpasNeighbors) SetMultihopTtl(v int64)`

SetMultihopTtl sets MultihopTtl field to given value.

### HasMultihopTtl

`func (o *MemberbgpasNeighbors) HasMultihopTtl() bool`

HasMultihopTtl returns a boolean if a field has been set.

### GetBfdTemplate

`func (o *MemberbgpasNeighbors) GetBfdTemplate() string`

GetBfdTemplate returns the BfdTemplate field if non-nil, zero value otherwise.

### GetBfdTemplateOk

`func (o *MemberbgpasNeighbors) GetBfdTemplateOk() (*string, bool)`

GetBfdTemplateOk returns a tuple with the BfdTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfdTemplate

`func (o *MemberbgpasNeighbors) SetBfdTemplate(v string)`

SetBfdTemplate sets BfdTemplate field to given value.

### HasBfdTemplate

`func (o *MemberbgpasNeighbors) HasBfdTemplate() bool`

HasBfdTemplate returns a boolean if a field has been set.

### GetEnableBfd

`func (o *MemberbgpasNeighbors) GetEnableBfd() bool`

GetEnableBfd returns the EnableBfd field if non-nil, zero value otherwise.

### GetEnableBfdOk

`func (o *MemberbgpasNeighbors) GetEnableBfdOk() (*bool, bool)`

GetEnableBfdOk returns a tuple with the EnableBfd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableBfd

`func (o *MemberbgpasNeighbors) SetEnableBfd(v bool)`

SetEnableBfd sets EnableBfd field to given value.

### HasEnableBfd

`func (o *MemberbgpasNeighbors) HasEnableBfd() bool`

HasEnableBfd returns a boolean if a field has been set.

### GetEnableBfdDnscheck

`func (o *MemberbgpasNeighbors) GetEnableBfdDnscheck() bool`

GetEnableBfdDnscheck returns the EnableBfdDnscheck field if non-nil, zero value otherwise.

### GetEnableBfdDnscheckOk

`func (o *MemberbgpasNeighbors) GetEnableBfdDnscheckOk() (*bool, bool)`

GetEnableBfdDnscheckOk returns a tuple with the EnableBfdDnscheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableBfdDnscheck

`func (o *MemberbgpasNeighbors) SetEnableBfdDnscheck(v bool)`

SetEnableBfdDnscheck sets EnableBfdDnscheck field to given value.

### HasEnableBfdDnscheck

`func (o *MemberbgpasNeighbors) HasEnableBfdDnscheck() bool`

HasEnableBfdDnscheck returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


