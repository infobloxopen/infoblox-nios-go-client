# GridMaxminddbinfo

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

## Methods

### NewGridMaxminddbinfo

`func NewGridMaxminddbinfo() *GridMaxminddbinfo`

NewGridMaxminddbinfo instantiates a new GridMaxminddbinfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGridMaxminddbinfoWithDefaults

`func NewGridMaxminddbinfoWithDefaults() *GridMaxminddbinfo`

NewGridMaxminddbinfoWithDefaults instantiates a new GridMaxminddbinfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GridMaxminddbinfo) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GridMaxminddbinfo) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GridMaxminddbinfo) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GridMaxminddbinfo) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetBinaryMajorVersion

`func (o *GridMaxminddbinfo) GetBinaryMajorVersion() int64`

GetBinaryMajorVersion returns the BinaryMajorVersion field if non-nil, zero value otherwise.

### GetBinaryMajorVersionOk

`func (o *GridMaxminddbinfo) GetBinaryMajorVersionOk() (*int64, bool)`

GetBinaryMajorVersionOk returns a tuple with the BinaryMajorVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryMajorVersion

`func (o *GridMaxminddbinfo) SetBinaryMajorVersion(v int64)`

SetBinaryMajorVersion sets BinaryMajorVersion field to given value.

### HasBinaryMajorVersion

`func (o *GridMaxminddbinfo) HasBinaryMajorVersion() bool`

HasBinaryMajorVersion returns a boolean if a field has been set.

### GetBinaryMinorVersion

`func (o *GridMaxminddbinfo) GetBinaryMinorVersion() int64`

GetBinaryMinorVersion returns the BinaryMinorVersion field if non-nil, zero value otherwise.

### GetBinaryMinorVersionOk

`func (o *GridMaxminddbinfo) GetBinaryMinorVersionOk() (*int64, bool)`

GetBinaryMinorVersionOk returns a tuple with the BinaryMinorVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryMinorVersion

`func (o *GridMaxminddbinfo) SetBinaryMinorVersion(v int64)`

SetBinaryMinorVersion sets BinaryMinorVersion field to given value.

### HasBinaryMinorVersion

`func (o *GridMaxminddbinfo) HasBinaryMinorVersion() bool`

HasBinaryMinorVersion returns a boolean if a field has been set.

### GetBuildTime

`func (o *GridMaxminddbinfo) GetBuildTime() int64`

GetBuildTime returns the BuildTime field if non-nil, zero value otherwise.

### GetBuildTimeOk

`func (o *GridMaxminddbinfo) GetBuildTimeOk() (*int64, bool)`

GetBuildTimeOk returns a tuple with the BuildTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildTime

`func (o *GridMaxminddbinfo) SetBuildTime(v int64)`

SetBuildTime sets BuildTime field to given value.

### HasBuildTime

`func (o *GridMaxminddbinfo) HasBuildTime() bool`

HasBuildTime returns a boolean if a field has been set.

### GetDatabaseType

`func (o *GridMaxminddbinfo) GetDatabaseType() string`

GetDatabaseType returns the DatabaseType field if non-nil, zero value otherwise.

### GetDatabaseTypeOk

`func (o *GridMaxminddbinfo) GetDatabaseTypeOk() (*string, bool)`

GetDatabaseTypeOk returns a tuple with the DatabaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseType

`func (o *GridMaxminddbinfo) SetDatabaseType(v string)`

SetDatabaseType sets DatabaseType field to given value.

### HasDatabaseType

`func (o *GridMaxminddbinfo) HasDatabaseType() bool`

HasDatabaseType returns a boolean if a field has been set.

### GetDeploymentTime

`func (o *GridMaxminddbinfo) GetDeploymentTime() int64`

GetDeploymentTime returns the DeploymentTime field if non-nil, zero value otherwise.

### GetDeploymentTimeOk

`func (o *GridMaxminddbinfo) GetDeploymentTimeOk() (*int64, bool)`

GetDeploymentTimeOk returns a tuple with the DeploymentTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentTime

`func (o *GridMaxminddbinfo) SetDeploymentTime(v int64)`

SetDeploymentTime sets DeploymentTime field to given value.

### HasDeploymentTime

`func (o *GridMaxminddbinfo) HasDeploymentTime() bool`

HasDeploymentTime returns a boolean if a field has been set.

### GetMember

`func (o *GridMaxminddbinfo) GetMember() string`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *GridMaxminddbinfo) GetMemberOk() (*string, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *GridMaxminddbinfo) SetMember(v string)`

SetMember sets Member field to given value.

### HasMember

`func (o *GridMaxminddbinfo) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetTopologyType

`func (o *GridMaxminddbinfo) GetTopologyType() string`

GetTopologyType returns the TopologyType field if non-nil, zero value otherwise.

### GetTopologyTypeOk

`func (o *GridMaxminddbinfo) GetTopologyTypeOk() (*string, bool)`

GetTopologyTypeOk returns a tuple with the TopologyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopologyType

`func (o *GridMaxminddbinfo) SetTopologyType(v string)`

SetTopologyType sets TopologyType field to given value.

### HasTopologyType

`func (o *GridMaxminddbinfo) HasTopologyType() bool`

HasTopologyType returns a boolean if a field has been set.

### GetUuid

`func (o *GridMaxminddbinfo) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *GridMaxminddbinfo) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *GridMaxminddbinfo) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *GridMaxminddbinfo) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


