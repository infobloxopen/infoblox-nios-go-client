# GetDtcResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**AddCertificate** | Pointer to **map[string]interface{}** |  | [optional] 
**DtcGetObjectGridState** | Pointer to **map[string]interface{}** |  | [optional] 
**DtcObjectDisable** | Pointer to **map[string]interface{}** |  | [optional] 
**DtcObjectEnable** | Pointer to **map[string]interface{}** |  | [optional] 
**GenerateEaTopologyDb** | Pointer to **map[string]interface{}** |  | [optional] 
**ImportMaxminddb** | Pointer to **map[string]interface{}** |  | [optional] 
**Query** | Pointer to **map[string]interface{}** |  | [optional] 
**Result** | Pointer to [**Dtc**](Dtc.md) |  | [optional] 

## Methods

### NewGetDtcResponse

`func NewGetDtcResponse() *GetDtcResponse`

NewGetDtcResponse instantiates a new GetDtcResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDtcResponseWithDefaults

`func NewGetDtcResponseWithDefaults() *GetDtcResponse`

NewGetDtcResponseWithDefaults instantiates a new GetDtcResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetDtcResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetDtcResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetDtcResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetDtcResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAddCertificate

`func (o *GetDtcResponse) GetAddCertificate() map[string]interface{}`

GetAddCertificate returns the AddCertificate field if non-nil, zero value otherwise.

### GetAddCertificateOk

`func (o *GetDtcResponse) GetAddCertificateOk() (*map[string]interface{}, bool)`

GetAddCertificateOk returns a tuple with the AddCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddCertificate

`func (o *GetDtcResponse) SetAddCertificate(v map[string]interface{})`

SetAddCertificate sets AddCertificate field to given value.

### HasAddCertificate

`func (o *GetDtcResponse) HasAddCertificate() bool`

HasAddCertificate returns a boolean if a field has been set.

### GetDtcGetObjectGridState

`func (o *GetDtcResponse) GetDtcGetObjectGridState() map[string]interface{}`

GetDtcGetObjectGridState returns the DtcGetObjectGridState field if non-nil, zero value otherwise.

### GetDtcGetObjectGridStateOk

`func (o *GetDtcResponse) GetDtcGetObjectGridStateOk() (*map[string]interface{}, bool)`

GetDtcGetObjectGridStateOk returns a tuple with the DtcGetObjectGridState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtcGetObjectGridState

`func (o *GetDtcResponse) SetDtcGetObjectGridState(v map[string]interface{})`

SetDtcGetObjectGridState sets DtcGetObjectGridState field to given value.

### HasDtcGetObjectGridState

`func (o *GetDtcResponse) HasDtcGetObjectGridState() bool`

HasDtcGetObjectGridState returns a boolean if a field has been set.

### GetDtcObjectDisable

`func (o *GetDtcResponse) GetDtcObjectDisable() map[string]interface{}`

GetDtcObjectDisable returns the DtcObjectDisable field if non-nil, zero value otherwise.

### GetDtcObjectDisableOk

`func (o *GetDtcResponse) GetDtcObjectDisableOk() (*map[string]interface{}, bool)`

GetDtcObjectDisableOk returns a tuple with the DtcObjectDisable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtcObjectDisable

`func (o *GetDtcResponse) SetDtcObjectDisable(v map[string]interface{})`

SetDtcObjectDisable sets DtcObjectDisable field to given value.

### HasDtcObjectDisable

`func (o *GetDtcResponse) HasDtcObjectDisable() bool`

HasDtcObjectDisable returns a boolean if a field has been set.

### GetDtcObjectEnable

`func (o *GetDtcResponse) GetDtcObjectEnable() map[string]interface{}`

GetDtcObjectEnable returns the DtcObjectEnable field if non-nil, zero value otherwise.

### GetDtcObjectEnableOk

`func (o *GetDtcResponse) GetDtcObjectEnableOk() (*map[string]interface{}, bool)`

GetDtcObjectEnableOk returns a tuple with the DtcObjectEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtcObjectEnable

`func (o *GetDtcResponse) SetDtcObjectEnable(v map[string]interface{})`

SetDtcObjectEnable sets DtcObjectEnable field to given value.

### HasDtcObjectEnable

`func (o *GetDtcResponse) HasDtcObjectEnable() bool`

HasDtcObjectEnable returns a boolean if a field has been set.

### GetGenerateEaTopologyDb

`func (o *GetDtcResponse) GetGenerateEaTopologyDb() map[string]interface{}`

GetGenerateEaTopologyDb returns the GenerateEaTopologyDb field if non-nil, zero value otherwise.

### GetGenerateEaTopologyDbOk

`func (o *GetDtcResponse) GetGenerateEaTopologyDbOk() (*map[string]interface{}, bool)`

GetGenerateEaTopologyDbOk returns a tuple with the GenerateEaTopologyDb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateEaTopologyDb

`func (o *GetDtcResponse) SetGenerateEaTopologyDb(v map[string]interface{})`

SetGenerateEaTopologyDb sets GenerateEaTopologyDb field to given value.

### HasGenerateEaTopologyDb

`func (o *GetDtcResponse) HasGenerateEaTopologyDb() bool`

HasGenerateEaTopologyDb returns a boolean if a field has been set.

### GetImportMaxminddb

`func (o *GetDtcResponse) GetImportMaxminddb() map[string]interface{}`

GetImportMaxminddb returns the ImportMaxminddb field if non-nil, zero value otherwise.

### GetImportMaxminddbOk

`func (o *GetDtcResponse) GetImportMaxminddbOk() (*map[string]interface{}, bool)`

GetImportMaxminddbOk returns a tuple with the ImportMaxminddb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportMaxminddb

`func (o *GetDtcResponse) SetImportMaxminddb(v map[string]interface{})`

SetImportMaxminddb sets ImportMaxminddb field to given value.

### HasImportMaxminddb

`func (o *GetDtcResponse) HasImportMaxminddb() bool`

HasImportMaxminddb returns a boolean if a field has been set.

### GetQuery

`func (o *GetDtcResponse) GetQuery() map[string]interface{}`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *GetDtcResponse) GetQueryOk() (*map[string]interface{}, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *GetDtcResponse) SetQuery(v map[string]interface{})`

SetQuery sets Query field to given value.

### HasQuery

`func (o *GetDtcResponse) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetResult

`func (o *GetDtcResponse) GetResult() Dtc`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetDtcResponse) GetResultOk() (*Dtc, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetDtcResponse) SetResult(v Dtc)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetDtcResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


