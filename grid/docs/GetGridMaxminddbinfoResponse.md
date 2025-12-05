# GetGridMaxminddbinfoResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**BinaryMajorVersion** | Pointer to **int64** | The major version of DB binary format. | [optional] [readonly] 
**BinaryMinorVersion** | Pointer to **int64** | The minor version of DB binary format. | [optional] [readonly] 
**BuildTime** | Pointer to **int64** | The time at which the DB was built. | [optional] [readonly] 
**DatabaseType** | Pointer to **string** | The structure of data records (\&quot;GeoLite2-Country\&quot;, GeoLite2-City\&quot;, etc.). | [optional] [readonly] 
**DeploymentTime** | Pointer to **int64** | The time at which the current Topology DB was deployed. | [optional] [readonly] 
**Member** | Pointer to **string** | The member for testing the connection. | [optional] [readonly] 
**TopologyType** | Pointer to **string** | The topology type. | [optional] [readonly] 
**Uuid** | Pointer to **string** | Universally Unique ID assigned for this object | [optional] [readonly] 
**Result** | Pointer to [**GridMaxminddbinfo**](GridMaxminddbinfo.md) |  | [optional] 

## Methods

### NewGetGridMaxminddbinfoResponse

`func NewGetGridMaxminddbinfoResponse() *GetGridMaxminddbinfoResponse`

NewGetGridMaxminddbinfoResponse instantiates a new GetGridMaxminddbinfoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetGridMaxminddbinfoResponseWithDefaults

`func NewGetGridMaxminddbinfoResponseWithDefaults() *GetGridMaxminddbinfoResponse`

NewGetGridMaxminddbinfoResponseWithDefaults instantiates a new GetGridMaxminddbinfoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetGridMaxminddbinfoResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetGridMaxminddbinfoResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetGridMaxminddbinfoResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetGridMaxminddbinfoResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetBinaryMajorVersion

`func (o *GetGridMaxminddbinfoResponse) GetBinaryMajorVersion() int64`

GetBinaryMajorVersion returns the BinaryMajorVersion field if non-nil, zero value otherwise.

### GetBinaryMajorVersionOk

`func (o *GetGridMaxminddbinfoResponse) GetBinaryMajorVersionOk() (*int64, bool)`

GetBinaryMajorVersionOk returns a tuple with the BinaryMajorVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryMajorVersion

`func (o *GetGridMaxminddbinfoResponse) SetBinaryMajorVersion(v int64)`

SetBinaryMajorVersion sets BinaryMajorVersion field to given value.

### HasBinaryMajorVersion

`func (o *GetGridMaxminddbinfoResponse) HasBinaryMajorVersion() bool`

HasBinaryMajorVersion returns a boolean if a field has been set.

### GetBinaryMinorVersion

`func (o *GetGridMaxminddbinfoResponse) GetBinaryMinorVersion() int64`

GetBinaryMinorVersion returns the BinaryMinorVersion field if non-nil, zero value otherwise.

### GetBinaryMinorVersionOk

`func (o *GetGridMaxminddbinfoResponse) GetBinaryMinorVersionOk() (*int64, bool)`

GetBinaryMinorVersionOk returns a tuple with the BinaryMinorVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryMinorVersion

`func (o *GetGridMaxminddbinfoResponse) SetBinaryMinorVersion(v int64)`

SetBinaryMinorVersion sets BinaryMinorVersion field to given value.

### HasBinaryMinorVersion

`func (o *GetGridMaxminddbinfoResponse) HasBinaryMinorVersion() bool`

HasBinaryMinorVersion returns a boolean if a field has been set.

### GetBuildTime

`func (o *GetGridMaxminddbinfoResponse) GetBuildTime() int64`

GetBuildTime returns the BuildTime field if non-nil, zero value otherwise.

### GetBuildTimeOk

`func (o *GetGridMaxminddbinfoResponse) GetBuildTimeOk() (*int64, bool)`

GetBuildTimeOk returns a tuple with the BuildTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildTime

`func (o *GetGridMaxminddbinfoResponse) SetBuildTime(v int64)`

SetBuildTime sets BuildTime field to given value.

### HasBuildTime

`func (o *GetGridMaxminddbinfoResponse) HasBuildTime() bool`

HasBuildTime returns a boolean if a field has been set.

### GetDatabaseType

`func (o *GetGridMaxminddbinfoResponse) GetDatabaseType() string`

GetDatabaseType returns the DatabaseType field if non-nil, zero value otherwise.

### GetDatabaseTypeOk

`func (o *GetGridMaxminddbinfoResponse) GetDatabaseTypeOk() (*string, bool)`

GetDatabaseTypeOk returns a tuple with the DatabaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseType

`func (o *GetGridMaxminddbinfoResponse) SetDatabaseType(v string)`

SetDatabaseType sets DatabaseType field to given value.

### HasDatabaseType

`func (o *GetGridMaxminddbinfoResponse) HasDatabaseType() bool`

HasDatabaseType returns a boolean if a field has been set.

### GetDeploymentTime

`func (o *GetGridMaxminddbinfoResponse) GetDeploymentTime() int64`

GetDeploymentTime returns the DeploymentTime field if non-nil, zero value otherwise.

### GetDeploymentTimeOk

`func (o *GetGridMaxminddbinfoResponse) GetDeploymentTimeOk() (*int64, bool)`

GetDeploymentTimeOk returns a tuple with the DeploymentTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentTime

`func (o *GetGridMaxminddbinfoResponse) SetDeploymentTime(v int64)`

SetDeploymentTime sets DeploymentTime field to given value.

### HasDeploymentTime

`func (o *GetGridMaxminddbinfoResponse) HasDeploymentTime() bool`

HasDeploymentTime returns a boolean if a field has been set.

### GetMember

`func (o *GetGridMaxminddbinfoResponse) GetMember() string`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *GetGridMaxminddbinfoResponse) GetMemberOk() (*string, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *GetGridMaxminddbinfoResponse) SetMember(v string)`

SetMember sets Member field to given value.

### HasMember

`func (o *GetGridMaxminddbinfoResponse) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetTopologyType

`func (o *GetGridMaxminddbinfoResponse) GetTopologyType() string`

GetTopologyType returns the TopologyType field if non-nil, zero value otherwise.

### GetTopologyTypeOk

`func (o *GetGridMaxminddbinfoResponse) GetTopologyTypeOk() (*string, bool)`

GetTopologyTypeOk returns a tuple with the TopologyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopologyType

`func (o *GetGridMaxminddbinfoResponse) SetTopologyType(v string)`

SetTopologyType sets TopologyType field to given value.

### HasTopologyType

`func (o *GetGridMaxminddbinfoResponse) HasTopologyType() bool`

HasTopologyType returns a boolean if a field has been set.

### GetUuid

`func (o *GetGridMaxminddbinfoResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *GetGridMaxminddbinfoResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *GetGridMaxminddbinfoResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *GetGridMaxminddbinfoResponse) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetResult

`func (o *GetGridMaxminddbinfoResponse) GetResult() GridMaxminddbinfo`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetGridMaxminddbinfoResponse) GetResultOk() (*GridMaxminddbinfo, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetGridMaxminddbinfoResponse) SetResult(v GridMaxminddbinfo)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetGridMaxminddbinfoResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


