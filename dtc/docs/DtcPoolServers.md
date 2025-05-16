# DtcPoolServers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Server** | Pointer to **string** | The server to link with. | [optional] 
**Ratio** | Pointer to **int64** | The weight of server. | [optional] 

## Methods

### NewDtcPoolServers

`func NewDtcPoolServers() *DtcPoolServers`

NewDtcPoolServers instantiates a new DtcPoolServers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtcPoolServersWithDefaults

`func NewDtcPoolServersWithDefaults() *DtcPoolServers`

NewDtcPoolServersWithDefaults instantiates a new DtcPoolServers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServer

`func (o *DtcPoolServers) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *DtcPoolServers) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *DtcPoolServers) SetServer(v string)`

SetServer sets Server field to given value.

### HasServer

`func (o *DtcPoolServers) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetRatio

`func (o *DtcPoolServers) GetRatio() int64`

GetRatio returns the Ratio field if non-nil, zero value otherwise.

### GetRatioOk

`func (o *DtcPoolServers) GetRatioOk() (*int64, bool)`

GetRatioOk returns a tuple with the Ratio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatio

`func (o *DtcPoolServers) SetRatio(v int64)`

SetRatio sets Ratio field to given value.

### HasRatio

`func (o *DtcPoolServers) HasRatio() bool`

HasRatio returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


