# GetDiscoveryMemberpropertiesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**Address** | Pointer to **string** | The Grid member address IP address. | [optional] [readonly] 
**CliCredentials** | Pointer to [**[]DiscoveryMemberpropertiesCliCredentials**](DiscoveryMemberpropertiesCliCredentials.md) | Discovery CLI credentials. | [optional] 
**DefaultSeedRouters** | Pointer to [**[]DiscoveryMemberpropertiesDefaultSeedRouters**](DiscoveryMemberpropertiesDefaultSeedRouters.md) | Default seed routers. | [optional] [readonly] 
**DiscoveryMember** | Pointer to **string** | The name of the network discovery Grid member. | [optional] [readonly] 
**EnableService** | Pointer to **bool** | Determines if the discovery service is enabled. | [optional] 
**GatewaySeedRouters** | Pointer to [**[]DiscoveryMemberpropertiesGatewaySeedRouters**](DiscoveryMemberpropertiesGatewaySeedRouters.md) | Gateway seed routers. | [optional] [readonly] 
**IsSa** | Pointer to **bool** | Determines if the standalone mode for discovery network monitor is enabled or not. | [optional] 
**Role** | Pointer to **string** | Discovery member role. | [optional] 
**ScanInterfaces** | Pointer to [**[]DiscoveryMemberpropertiesScanInterfaces**](DiscoveryMemberpropertiesScanInterfaces.md) | Discovery networks to which the member is assigned. | [optional] 
**SdnConfigs** | Pointer to [**[]DiscoveryMemberpropertiesSdnConfigs**](DiscoveryMemberpropertiesSdnConfigs.md) | List of SDN/SDWAN controller configurations. | [optional] 
**SeedRouters** | Pointer to [**[]DiscoveryMemberpropertiesSeedRouters**](DiscoveryMemberpropertiesSeedRouters.md) | Seed routers. | [optional] 
**Snmpv1v2Credentials** | Pointer to [**[]DiscoveryMemberpropertiesSnmpv1v2Credentials**](DiscoveryMemberpropertiesSnmpv1v2Credentials.md) | Discovery SNMP v1 and v2 credentials. | [optional] 
**Snmpv3Credentials** | Pointer to [**[]DiscoveryMemberpropertiesSnmpv3Credentials**](DiscoveryMemberpropertiesSnmpv3Credentials.md) | Discovery SNMP v3 credentials. | [optional] 
**UseCliCredentials** | Pointer to **bool** | Use flag for: cli_credentials | [optional] 
**UseSnmpv1v2Credentials** | Pointer to **bool** | Use flag for: snmpv1v2_credentials | [optional] 
**UseSnmpv3Credentials** | Pointer to **bool** | Use flag for: snmpv3_credentials | [optional] 
**Uuid** | Pointer to **string** | Universally Unique ID assigned for this object | [optional] [readonly] 
**Result** | Pointer to [**DiscoveryMemberproperties**](DiscoveryMemberproperties.md) |  | [optional] 

## Methods

### NewGetDiscoveryMemberpropertiesResponse

`func NewGetDiscoveryMemberpropertiesResponse() *GetDiscoveryMemberpropertiesResponse`

NewGetDiscoveryMemberpropertiesResponse instantiates a new GetDiscoveryMemberpropertiesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDiscoveryMemberpropertiesResponseWithDefaults

`func NewGetDiscoveryMemberpropertiesResponseWithDefaults() *GetDiscoveryMemberpropertiesResponse`

NewGetDiscoveryMemberpropertiesResponseWithDefaults instantiates a new GetDiscoveryMemberpropertiesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetDiscoveryMemberpropertiesResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetDiscoveryMemberpropertiesResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetDiscoveryMemberpropertiesResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetDiscoveryMemberpropertiesResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAddress

`func (o *GetDiscoveryMemberpropertiesResponse) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *GetDiscoveryMemberpropertiesResponse) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *GetDiscoveryMemberpropertiesResponse) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *GetDiscoveryMemberpropertiesResponse) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetCliCredentials

`func (o *GetDiscoveryMemberpropertiesResponse) GetCliCredentials() []DiscoveryMemberpropertiesCliCredentials`

GetCliCredentials returns the CliCredentials field if non-nil, zero value otherwise.

### GetCliCredentialsOk

`func (o *GetDiscoveryMemberpropertiesResponse) GetCliCredentialsOk() (*[]DiscoveryMemberpropertiesCliCredentials, bool)`

GetCliCredentialsOk returns a tuple with the CliCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCliCredentials

`func (o *GetDiscoveryMemberpropertiesResponse) SetCliCredentials(v []DiscoveryMemberpropertiesCliCredentials)`

SetCliCredentials sets CliCredentials field to given value.

### HasCliCredentials

`func (o *GetDiscoveryMemberpropertiesResponse) HasCliCredentials() bool`

HasCliCredentials returns a boolean if a field has been set.

### GetDefaultSeedRouters

`func (o *GetDiscoveryMemberpropertiesResponse) GetDefaultSeedRouters() []DiscoveryMemberpropertiesDefaultSeedRouters`

GetDefaultSeedRouters returns the DefaultSeedRouters field if non-nil, zero value otherwise.

### GetDefaultSeedRoutersOk

`func (o *GetDiscoveryMemberpropertiesResponse) GetDefaultSeedRoutersOk() (*[]DiscoveryMemberpropertiesDefaultSeedRouters, bool)`

GetDefaultSeedRoutersOk returns a tuple with the DefaultSeedRouters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSeedRouters

`func (o *GetDiscoveryMemberpropertiesResponse) SetDefaultSeedRouters(v []DiscoveryMemberpropertiesDefaultSeedRouters)`

SetDefaultSeedRouters sets DefaultSeedRouters field to given value.

### HasDefaultSeedRouters

`func (o *GetDiscoveryMemberpropertiesResponse) HasDefaultSeedRouters() bool`

HasDefaultSeedRouters returns a boolean if a field has been set.

### GetDiscoveryMember

`func (o *GetDiscoveryMemberpropertiesResponse) GetDiscoveryMember() string`

GetDiscoveryMember returns the DiscoveryMember field if non-nil, zero value otherwise.

### GetDiscoveryMemberOk

`func (o *GetDiscoveryMemberpropertiesResponse) GetDiscoveryMemberOk() (*string, bool)`

GetDiscoveryMemberOk returns a tuple with the DiscoveryMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryMember

`func (o *GetDiscoveryMemberpropertiesResponse) SetDiscoveryMember(v string)`

SetDiscoveryMember sets DiscoveryMember field to given value.

### HasDiscoveryMember

`func (o *GetDiscoveryMemberpropertiesResponse) HasDiscoveryMember() bool`

HasDiscoveryMember returns a boolean if a field has been set.

### GetEnableService

`func (o *GetDiscoveryMemberpropertiesResponse) GetEnableService() bool`

GetEnableService returns the EnableService field if non-nil, zero value otherwise.

### GetEnableServiceOk

`func (o *GetDiscoveryMemberpropertiesResponse) GetEnableServiceOk() (*bool, bool)`

GetEnableServiceOk returns a tuple with the EnableService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableService

`func (o *GetDiscoveryMemberpropertiesResponse) SetEnableService(v bool)`

SetEnableService sets EnableService field to given value.

### HasEnableService

`func (o *GetDiscoveryMemberpropertiesResponse) HasEnableService() bool`

HasEnableService returns a boolean if a field has been set.

### GetGatewaySeedRouters

`func (o *GetDiscoveryMemberpropertiesResponse) GetGatewaySeedRouters() []DiscoveryMemberpropertiesGatewaySeedRouters`

GetGatewaySeedRouters returns the GatewaySeedRouters field if non-nil, zero value otherwise.

### GetGatewaySeedRoutersOk

`func (o *GetDiscoveryMemberpropertiesResponse) GetGatewaySeedRoutersOk() (*[]DiscoveryMemberpropertiesGatewaySeedRouters, bool)`

GetGatewaySeedRoutersOk returns a tuple with the GatewaySeedRouters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewaySeedRouters

`func (o *GetDiscoveryMemberpropertiesResponse) SetGatewaySeedRouters(v []DiscoveryMemberpropertiesGatewaySeedRouters)`

SetGatewaySeedRouters sets GatewaySeedRouters field to given value.

### HasGatewaySeedRouters

`func (o *GetDiscoveryMemberpropertiesResponse) HasGatewaySeedRouters() bool`

HasGatewaySeedRouters returns a boolean if a field has been set.

### GetIsSa

`func (o *GetDiscoveryMemberpropertiesResponse) GetIsSa() bool`

GetIsSa returns the IsSa field if non-nil, zero value otherwise.

### GetIsSaOk

`func (o *GetDiscoveryMemberpropertiesResponse) GetIsSaOk() (*bool, bool)`

GetIsSaOk returns a tuple with the IsSa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSa

`func (o *GetDiscoveryMemberpropertiesResponse) SetIsSa(v bool)`

SetIsSa sets IsSa field to given value.

### HasIsSa

`func (o *GetDiscoveryMemberpropertiesResponse) HasIsSa() bool`

HasIsSa returns a boolean if a field has been set.

### GetRole

`func (o *GetDiscoveryMemberpropertiesResponse) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *GetDiscoveryMemberpropertiesResponse) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *GetDiscoveryMemberpropertiesResponse) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *GetDiscoveryMemberpropertiesResponse) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetScanInterfaces

`func (o *GetDiscoveryMemberpropertiesResponse) GetScanInterfaces() []DiscoveryMemberpropertiesScanInterfaces`

GetScanInterfaces returns the ScanInterfaces field if non-nil, zero value otherwise.

### GetScanInterfacesOk

`func (o *GetDiscoveryMemberpropertiesResponse) GetScanInterfacesOk() (*[]DiscoveryMemberpropertiesScanInterfaces, bool)`

GetScanInterfacesOk returns a tuple with the ScanInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanInterfaces

`func (o *GetDiscoveryMemberpropertiesResponse) SetScanInterfaces(v []DiscoveryMemberpropertiesScanInterfaces)`

SetScanInterfaces sets ScanInterfaces field to given value.

### HasScanInterfaces

`func (o *GetDiscoveryMemberpropertiesResponse) HasScanInterfaces() bool`

HasScanInterfaces returns a boolean if a field has been set.

### GetSdnConfigs

`func (o *GetDiscoveryMemberpropertiesResponse) GetSdnConfigs() []DiscoveryMemberpropertiesSdnConfigs`

GetSdnConfigs returns the SdnConfigs field if non-nil, zero value otherwise.

### GetSdnConfigsOk

`func (o *GetDiscoveryMemberpropertiesResponse) GetSdnConfigsOk() (*[]DiscoveryMemberpropertiesSdnConfigs, bool)`

GetSdnConfigsOk returns a tuple with the SdnConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdnConfigs

`func (o *GetDiscoveryMemberpropertiesResponse) SetSdnConfigs(v []DiscoveryMemberpropertiesSdnConfigs)`

SetSdnConfigs sets SdnConfigs field to given value.

### HasSdnConfigs

`func (o *GetDiscoveryMemberpropertiesResponse) HasSdnConfigs() bool`

HasSdnConfigs returns a boolean if a field has been set.

### GetSeedRouters

`func (o *GetDiscoveryMemberpropertiesResponse) GetSeedRouters() []DiscoveryMemberpropertiesSeedRouters`

GetSeedRouters returns the SeedRouters field if non-nil, zero value otherwise.

### GetSeedRoutersOk

`func (o *GetDiscoveryMemberpropertiesResponse) GetSeedRoutersOk() (*[]DiscoveryMemberpropertiesSeedRouters, bool)`

GetSeedRoutersOk returns a tuple with the SeedRouters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeedRouters

`func (o *GetDiscoveryMemberpropertiesResponse) SetSeedRouters(v []DiscoveryMemberpropertiesSeedRouters)`

SetSeedRouters sets SeedRouters field to given value.

### HasSeedRouters

`func (o *GetDiscoveryMemberpropertiesResponse) HasSeedRouters() bool`

HasSeedRouters returns a boolean if a field has been set.

### GetSnmpv1v2Credentials

`func (o *GetDiscoveryMemberpropertiesResponse) GetSnmpv1v2Credentials() []DiscoveryMemberpropertiesSnmpv1v2Credentials`

GetSnmpv1v2Credentials returns the Snmpv1v2Credentials field if non-nil, zero value otherwise.

### GetSnmpv1v2CredentialsOk

`func (o *GetDiscoveryMemberpropertiesResponse) GetSnmpv1v2CredentialsOk() (*[]DiscoveryMemberpropertiesSnmpv1v2Credentials, bool)`

GetSnmpv1v2CredentialsOk returns a tuple with the Snmpv1v2Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmpv1v2Credentials

`func (o *GetDiscoveryMemberpropertiesResponse) SetSnmpv1v2Credentials(v []DiscoveryMemberpropertiesSnmpv1v2Credentials)`

SetSnmpv1v2Credentials sets Snmpv1v2Credentials field to given value.

### HasSnmpv1v2Credentials

`func (o *GetDiscoveryMemberpropertiesResponse) HasSnmpv1v2Credentials() bool`

HasSnmpv1v2Credentials returns a boolean if a field has been set.

### GetSnmpv3Credentials

`func (o *GetDiscoveryMemberpropertiesResponse) GetSnmpv3Credentials() []DiscoveryMemberpropertiesSnmpv3Credentials`

GetSnmpv3Credentials returns the Snmpv3Credentials field if non-nil, zero value otherwise.

### GetSnmpv3CredentialsOk

`func (o *GetDiscoveryMemberpropertiesResponse) GetSnmpv3CredentialsOk() (*[]DiscoveryMemberpropertiesSnmpv3Credentials, bool)`

GetSnmpv3CredentialsOk returns a tuple with the Snmpv3Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmpv3Credentials

`func (o *GetDiscoveryMemberpropertiesResponse) SetSnmpv3Credentials(v []DiscoveryMemberpropertiesSnmpv3Credentials)`

SetSnmpv3Credentials sets Snmpv3Credentials field to given value.

### HasSnmpv3Credentials

`func (o *GetDiscoveryMemberpropertiesResponse) HasSnmpv3Credentials() bool`

HasSnmpv3Credentials returns a boolean if a field has been set.

### GetUseCliCredentials

`func (o *GetDiscoveryMemberpropertiesResponse) GetUseCliCredentials() bool`

GetUseCliCredentials returns the UseCliCredentials field if non-nil, zero value otherwise.

### GetUseCliCredentialsOk

`func (o *GetDiscoveryMemberpropertiesResponse) GetUseCliCredentialsOk() (*bool, bool)`

GetUseCliCredentialsOk returns a tuple with the UseCliCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCliCredentials

`func (o *GetDiscoveryMemberpropertiesResponse) SetUseCliCredentials(v bool)`

SetUseCliCredentials sets UseCliCredentials field to given value.

### HasUseCliCredentials

`func (o *GetDiscoveryMemberpropertiesResponse) HasUseCliCredentials() bool`

HasUseCliCredentials returns a boolean if a field has been set.

### GetUseSnmpv1v2Credentials

`func (o *GetDiscoveryMemberpropertiesResponse) GetUseSnmpv1v2Credentials() bool`

GetUseSnmpv1v2Credentials returns the UseSnmpv1v2Credentials field if non-nil, zero value otherwise.

### GetUseSnmpv1v2CredentialsOk

`func (o *GetDiscoveryMemberpropertiesResponse) GetUseSnmpv1v2CredentialsOk() (*bool, bool)`

GetUseSnmpv1v2CredentialsOk returns a tuple with the UseSnmpv1v2Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSnmpv1v2Credentials

`func (o *GetDiscoveryMemberpropertiesResponse) SetUseSnmpv1v2Credentials(v bool)`

SetUseSnmpv1v2Credentials sets UseSnmpv1v2Credentials field to given value.

### HasUseSnmpv1v2Credentials

`func (o *GetDiscoveryMemberpropertiesResponse) HasUseSnmpv1v2Credentials() bool`

HasUseSnmpv1v2Credentials returns a boolean if a field has been set.

### GetUseSnmpv3Credentials

`func (o *GetDiscoveryMemberpropertiesResponse) GetUseSnmpv3Credentials() bool`

GetUseSnmpv3Credentials returns the UseSnmpv3Credentials field if non-nil, zero value otherwise.

### GetUseSnmpv3CredentialsOk

`func (o *GetDiscoveryMemberpropertiesResponse) GetUseSnmpv3CredentialsOk() (*bool, bool)`

GetUseSnmpv3CredentialsOk returns a tuple with the UseSnmpv3Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSnmpv3Credentials

`func (o *GetDiscoveryMemberpropertiesResponse) SetUseSnmpv3Credentials(v bool)`

SetUseSnmpv3Credentials sets UseSnmpv3Credentials field to given value.

### HasUseSnmpv3Credentials

`func (o *GetDiscoveryMemberpropertiesResponse) HasUseSnmpv3Credentials() bool`

HasUseSnmpv3Credentials returns a boolean if a field has been set.

### GetUuid

`func (o *GetDiscoveryMemberpropertiesResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *GetDiscoveryMemberpropertiesResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *GetDiscoveryMemberpropertiesResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *GetDiscoveryMemberpropertiesResponse) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetResult

`func (o *GetDiscoveryMemberpropertiesResponse) GetResult() DiscoveryMemberproperties`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetDiscoveryMemberpropertiesResponse) GetResultOk() (*DiscoveryMemberproperties, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetDiscoveryMemberpropertiesResponse) SetResult(v DiscoveryMemberproperties)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetDiscoveryMemberpropertiesResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


