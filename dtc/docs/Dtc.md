# Dtc

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

## Methods

### NewDtc

`func NewDtc() *Dtc`

NewDtc instantiates a new Dtc object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtcWithDefaults

`func NewDtcWithDefaults() *Dtc`

NewDtcWithDefaults instantiates a new Dtc object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *Dtc) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *Dtc) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *Dtc) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *Dtc) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAddCertificate

`func (o *Dtc) GetAddCertificate() map[string]interface{}`

GetAddCertificate returns the AddCertificate field if non-nil, zero value otherwise.

### GetAddCertificateOk

`func (o *Dtc) GetAddCertificateOk() (*map[string]interface{}, bool)`

GetAddCertificateOk returns a tuple with the AddCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddCertificate

`func (o *Dtc) SetAddCertificate(v map[string]interface{})`

SetAddCertificate sets AddCertificate field to given value.

### HasAddCertificate

`func (o *Dtc) HasAddCertificate() bool`

HasAddCertificate returns a boolean if a field has been set.

### GetDtcGetObjectGridState

`func (o *Dtc) GetDtcGetObjectGridState() map[string]interface{}`

GetDtcGetObjectGridState returns the DtcGetObjectGridState field if non-nil, zero value otherwise.

### GetDtcGetObjectGridStateOk

`func (o *Dtc) GetDtcGetObjectGridStateOk() (*map[string]interface{}, bool)`

GetDtcGetObjectGridStateOk returns a tuple with the DtcGetObjectGridState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtcGetObjectGridState

`func (o *Dtc) SetDtcGetObjectGridState(v map[string]interface{})`

SetDtcGetObjectGridState sets DtcGetObjectGridState field to given value.

### HasDtcGetObjectGridState

`func (o *Dtc) HasDtcGetObjectGridState() bool`

HasDtcGetObjectGridState returns a boolean if a field has been set.

### GetDtcObjectDisable

`func (o *Dtc) GetDtcObjectDisable() map[string]interface{}`

GetDtcObjectDisable returns the DtcObjectDisable field if non-nil, zero value otherwise.

### GetDtcObjectDisableOk

`func (o *Dtc) GetDtcObjectDisableOk() (*map[string]interface{}, bool)`

GetDtcObjectDisableOk returns a tuple with the DtcObjectDisable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtcObjectDisable

`func (o *Dtc) SetDtcObjectDisable(v map[string]interface{})`

SetDtcObjectDisable sets DtcObjectDisable field to given value.

### HasDtcObjectDisable

`func (o *Dtc) HasDtcObjectDisable() bool`

HasDtcObjectDisable returns a boolean if a field has been set.

### GetDtcObjectEnable

`func (o *Dtc) GetDtcObjectEnable() map[string]interface{}`

GetDtcObjectEnable returns the DtcObjectEnable field if non-nil, zero value otherwise.

### GetDtcObjectEnableOk

`func (o *Dtc) GetDtcObjectEnableOk() (*map[string]interface{}, bool)`

GetDtcObjectEnableOk returns a tuple with the DtcObjectEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtcObjectEnable

`func (o *Dtc) SetDtcObjectEnable(v map[string]interface{})`

SetDtcObjectEnable sets DtcObjectEnable field to given value.

### HasDtcObjectEnable

`func (o *Dtc) HasDtcObjectEnable() bool`

HasDtcObjectEnable returns a boolean if a field has been set.

### GetGenerateEaTopologyDb

`func (o *Dtc) GetGenerateEaTopologyDb() map[string]interface{}`

GetGenerateEaTopologyDb returns the GenerateEaTopologyDb field if non-nil, zero value otherwise.

### GetGenerateEaTopologyDbOk

`func (o *Dtc) GetGenerateEaTopologyDbOk() (*map[string]interface{}, bool)`

GetGenerateEaTopologyDbOk returns a tuple with the GenerateEaTopologyDb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateEaTopologyDb

`func (o *Dtc) SetGenerateEaTopologyDb(v map[string]interface{})`

SetGenerateEaTopologyDb sets GenerateEaTopologyDb field to given value.

### HasGenerateEaTopologyDb

`func (o *Dtc) HasGenerateEaTopologyDb() bool`

HasGenerateEaTopologyDb returns a boolean if a field has been set.

### GetImportMaxminddb

`func (o *Dtc) GetImportMaxminddb() map[string]interface{}`

GetImportMaxminddb returns the ImportMaxminddb field if non-nil, zero value otherwise.

### GetImportMaxminddbOk

`func (o *Dtc) GetImportMaxminddbOk() (*map[string]interface{}, bool)`

GetImportMaxminddbOk returns a tuple with the ImportMaxminddb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportMaxminddb

`func (o *Dtc) SetImportMaxminddb(v map[string]interface{})`

SetImportMaxminddb sets ImportMaxminddb field to given value.

### HasImportMaxminddb

`func (o *Dtc) HasImportMaxminddb() bool`

HasImportMaxminddb returns a boolean if a field has been set.

### GetQuery

`func (o *Dtc) GetQuery() map[string]interface{}`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *Dtc) GetQueryOk() (*map[string]interface{}, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *Dtc) SetQuery(v map[string]interface{})`

SetQuery sets Query field to given value.

### HasQuery

`func (o *Dtc) HasQuery() bool`

HasQuery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


