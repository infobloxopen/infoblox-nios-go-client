# DiscoveryMemberproperties

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

## Methods

### NewDiscoveryMemberproperties

`func NewDiscoveryMemberproperties() *DiscoveryMemberproperties`

NewDiscoveryMemberproperties instantiates a new DiscoveryMemberproperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscoveryMemberpropertiesWithDefaults

`func NewDiscoveryMemberpropertiesWithDefaults() *DiscoveryMemberproperties`

NewDiscoveryMemberpropertiesWithDefaults instantiates a new DiscoveryMemberproperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *DiscoveryMemberproperties) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *DiscoveryMemberproperties) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *DiscoveryMemberproperties) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *DiscoveryMemberproperties) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAddress

`func (o *DiscoveryMemberproperties) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *DiscoveryMemberproperties) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *DiscoveryMemberproperties) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *DiscoveryMemberproperties) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetCliCredentials

`func (o *DiscoveryMemberproperties) GetCliCredentials() []DiscoveryMemberpropertiesCliCredentials`

GetCliCredentials returns the CliCredentials field if non-nil, zero value otherwise.

### GetCliCredentialsOk

`func (o *DiscoveryMemberproperties) GetCliCredentialsOk() (*[]DiscoveryMemberpropertiesCliCredentials, bool)`

GetCliCredentialsOk returns a tuple with the CliCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCliCredentials

`func (o *DiscoveryMemberproperties) SetCliCredentials(v []DiscoveryMemberpropertiesCliCredentials)`

SetCliCredentials sets CliCredentials field to given value.

### HasCliCredentials

`func (o *DiscoveryMemberproperties) HasCliCredentials() bool`

HasCliCredentials returns a boolean if a field has been set.

### GetDefaultSeedRouters

`func (o *DiscoveryMemberproperties) GetDefaultSeedRouters() []DiscoveryMemberpropertiesDefaultSeedRouters`

GetDefaultSeedRouters returns the DefaultSeedRouters field if non-nil, zero value otherwise.

### GetDefaultSeedRoutersOk

`func (o *DiscoveryMemberproperties) GetDefaultSeedRoutersOk() (*[]DiscoveryMemberpropertiesDefaultSeedRouters, bool)`

GetDefaultSeedRoutersOk returns a tuple with the DefaultSeedRouters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSeedRouters

`func (o *DiscoveryMemberproperties) SetDefaultSeedRouters(v []DiscoveryMemberpropertiesDefaultSeedRouters)`

SetDefaultSeedRouters sets DefaultSeedRouters field to given value.

### HasDefaultSeedRouters

`func (o *DiscoveryMemberproperties) HasDefaultSeedRouters() bool`

HasDefaultSeedRouters returns a boolean if a field has been set.

### GetDiscoveryMember

`func (o *DiscoveryMemberproperties) GetDiscoveryMember() string`

GetDiscoveryMember returns the DiscoveryMember field if non-nil, zero value otherwise.

### GetDiscoveryMemberOk

`func (o *DiscoveryMemberproperties) GetDiscoveryMemberOk() (*string, bool)`

GetDiscoveryMemberOk returns a tuple with the DiscoveryMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryMember

`func (o *DiscoveryMemberproperties) SetDiscoveryMember(v string)`

SetDiscoveryMember sets DiscoveryMember field to given value.

### HasDiscoveryMember

`func (o *DiscoveryMemberproperties) HasDiscoveryMember() bool`

HasDiscoveryMember returns a boolean if a field has been set.

### GetEnableService

`func (o *DiscoveryMemberproperties) GetEnableService() bool`

GetEnableService returns the EnableService field if non-nil, zero value otherwise.

### GetEnableServiceOk

`func (o *DiscoveryMemberproperties) GetEnableServiceOk() (*bool, bool)`

GetEnableServiceOk returns a tuple with the EnableService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableService

`func (o *DiscoveryMemberproperties) SetEnableService(v bool)`

SetEnableService sets EnableService field to given value.

### HasEnableService

`func (o *DiscoveryMemberproperties) HasEnableService() bool`

HasEnableService returns a boolean if a field has been set.

### GetGatewaySeedRouters

`func (o *DiscoveryMemberproperties) GetGatewaySeedRouters() []DiscoveryMemberpropertiesGatewaySeedRouters`

GetGatewaySeedRouters returns the GatewaySeedRouters field if non-nil, zero value otherwise.

### GetGatewaySeedRoutersOk

`func (o *DiscoveryMemberproperties) GetGatewaySeedRoutersOk() (*[]DiscoveryMemberpropertiesGatewaySeedRouters, bool)`

GetGatewaySeedRoutersOk returns a tuple with the GatewaySeedRouters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewaySeedRouters

`func (o *DiscoveryMemberproperties) SetGatewaySeedRouters(v []DiscoveryMemberpropertiesGatewaySeedRouters)`

SetGatewaySeedRouters sets GatewaySeedRouters field to given value.

### HasGatewaySeedRouters

`func (o *DiscoveryMemberproperties) HasGatewaySeedRouters() bool`

HasGatewaySeedRouters returns a boolean if a field has been set.

### GetIsSa

`func (o *DiscoveryMemberproperties) GetIsSa() bool`

GetIsSa returns the IsSa field if non-nil, zero value otherwise.

### GetIsSaOk

`func (o *DiscoveryMemberproperties) GetIsSaOk() (*bool, bool)`

GetIsSaOk returns a tuple with the IsSa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSa

`func (o *DiscoveryMemberproperties) SetIsSa(v bool)`

SetIsSa sets IsSa field to given value.

### HasIsSa

`func (o *DiscoveryMemberproperties) HasIsSa() bool`

HasIsSa returns a boolean if a field has been set.

### GetRole

`func (o *DiscoveryMemberproperties) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *DiscoveryMemberproperties) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *DiscoveryMemberproperties) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *DiscoveryMemberproperties) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetScanInterfaces

`func (o *DiscoveryMemberproperties) GetScanInterfaces() []DiscoveryMemberpropertiesScanInterfaces`

GetScanInterfaces returns the ScanInterfaces field if non-nil, zero value otherwise.

### GetScanInterfacesOk

`func (o *DiscoveryMemberproperties) GetScanInterfacesOk() (*[]DiscoveryMemberpropertiesScanInterfaces, bool)`

GetScanInterfacesOk returns a tuple with the ScanInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanInterfaces

`func (o *DiscoveryMemberproperties) SetScanInterfaces(v []DiscoveryMemberpropertiesScanInterfaces)`

SetScanInterfaces sets ScanInterfaces field to given value.

### HasScanInterfaces

`func (o *DiscoveryMemberproperties) HasScanInterfaces() bool`

HasScanInterfaces returns a boolean if a field has been set.

### GetSdnConfigs

`func (o *DiscoveryMemberproperties) GetSdnConfigs() []DiscoveryMemberpropertiesSdnConfigs`

GetSdnConfigs returns the SdnConfigs field if non-nil, zero value otherwise.

### GetSdnConfigsOk

`func (o *DiscoveryMemberproperties) GetSdnConfigsOk() (*[]DiscoveryMemberpropertiesSdnConfigs, bool)`

GetSdnConfigsOk returns a tuple with the SdnConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdnConfigs

`func (o *DiscoveryMemberproperties) SetSdnConfigs(v []DiscoveryMemberpropertiesSdnConfigs)`

SetSdnConfigs sets SdnConfigs field to given value.

### HasSdnConfigs

`func (o *DiscoveryMemberproperties) HasSdnConfigs() bool`

HasSdnConfigs returns a boolean if a field has been set.

### GetSeedRouters

`func (o *DiscoveryMemberproperties) GetSeedRouters() []DiscoveryMemberpropertiesSeedRouters`

GetSeedRouters returns the SeedRouters field if non-nil, zero value otherwise.

### GetSeedRoutersOk

`func (o *DiscoveryMemberproperties) GetSeedRoutersOk() (*[]DiscoveryMemberpropertiesSeedRouters, bool)`

GetSeedRoutersOk returns a tuple with the SeedRouters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeedRouters

`func (o *DiscoveryMemberproperties) SetSeedRouters(v []DiscoveryMemberpropertiesSeedRouters)`

SetSeedRouters sets SeedRouters field to given value.

### HasSeedRouters

`func (o *DiscoveryMemberproperties) HasSeedRouters() bool`

HasSeedRouters returns a boolean if a field has been set.

### GetSnmpv1v2Credentials

`func (o *DiscoveryMemberproperties) GetSnmpv1v2Credentials() []DiscoveryMemberpropertiesSnmpv1v2Credentials`

GetSnmpv1v2Credentials returns the Snmpv1v2Credentials field if non-nil, zero value otherwise.

### GetSnmpv1v2CredentialsOk

`func (o *DiscoveryMemberproperties) GetSnmpv1v2CredentialsOk() (*[]DiscoveryMemberpropertiesSnmpv1v2Credentials, bool)`

GetSnmpv1v2CredentialsOk returns a tuple with the Snmpv1v2Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmpv1v2Credentials

`func (o *DiscoveryMemberproperties) SetSnmpv1v2Credentials(v []DiscoveryMemberpropertiesSnmpv1v2Credentials)`

SetSnmpv1v2Credentials sets Snmpv1v2Credentials field to given value.

### HasSnmpv1v2Credentials

`func (o *DiscoveryMemberproperties) HasSnmpv1v2Credentials() bool`

HasSnmpv1v2Credentials returns a boolean if a field has been set.

### GetSnmpv3Credentials

`func (o *DiscoveryMemberproperties) GetSnmpv3Credentials() []DiscoveryMemberpropertiesSnmpv3Credentials`

GetSnmpv3Credentials returns the Snmpv3Credentials field if non-nil, zero value otherwise.

### GetSnmpv3CredentialsOk

`func (o *DiscoveryMemberproperties) GetSnmpv3CredentialsOk() (*[]DiscoveryMemberpropertiesSnmpv3Credentials, bool)`

GetSnmpv3CredentialsOk returns a tuple with the Snmpv3Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmpv3Credentials

`func (o *DiscoveryMemberproperties) SetSnmpv3Credentials(v []DiscoveryMemberpropertiesSnmpv3Credentials)`

SetSnmpv3Credentials sets Snmpv3Credentials field to given value.

### HasSnmpv3Credentials

`func (o *DiscoveryMemberproperties) HasSnmpv3Credentials() bool`

HasSnmpv3Credentials returns a boolean if a field has been set.

### GetUseCliCredentials

`func (o *DiscoveryMemberproperties) GetUseCliCredentials() bool`

GetUseCliCredentials returns the UseCliCredentials field if non-nil, zero value otherwise.

### GetUseCliCredentialsOk

`func (o *DiscoveryMemberproperties) GetUseCliCredentialsOk() (*bool, bool)`

GetUseCliCredentialsOk returns a tuple with the UseCliCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCliCredentials

`func (o *DiscoveryMemberproperties) SetUseCliCredentials(v bool)`

SetUseCliCredentials sets UseCliCredentials field to given value.

### HasUseCliCredentials

`func (o *DiscoveryMemberproperties) HasUseCliCredentials() bool`

HasUseCliCredentials returns a boolean if a field has been set.

### GetUseSnmpv1v2Credentials

`func (o *DiscoveryMemberproperties) GetUseSnmpv1v2Credentials() bool`

GetUseSnmpv1v2Credentials returns the UseSnmpv1v2Credentials field if non-nil, zero value otherwise.

### GetUseSnmpv1v2CredentialsOk

`func (o *DiscoveryMemberproperties) GetUseSnmpv1v2CredentialsOk() (*bool, bool)`

GetUseSnmpv1v2CredentialsOk returns a tuple with the UseSnmpv1v2Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSnmpv1v2Credentials

`func (o *DiscoveryMemberproperties) SetUseSnmpv1v2Credentials(v bool)`

SetUseSnmpv1v2Credentials sets UseSnmpv1v2Credentials field to given value.

### HasUseSnmpv1v2Credentials

`func (o *DiscoveryMemberproperties) HasUseSnmpv1v2Credentials() bool`

HasUseSnmpv1v2Credentials returns a boolean if a field has been set.

### GetUseSnmpv3Credentials

`func (o *DiscoveryMemberproperties) GetUseSnmpv3Credentials() bool`

GetUseSnmpv3Credentials returns the UseSnmpv3Credentials field if non-nil, zero value otherwise.

### GetUseSnmpv3CredentialsOk

`func (o *DiscoveryMemberproperties) GetUseSnmpv3CredentialsOk() (*bool, bool)`

GetUseSnmpv3CredentialsOk returns a tuple with the UseSnmpv3Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSnmpv3Credentials

`func (o *DiscoveryMemberproperties) SetUseSnmpv3Credentials(v bool)`

SetUseSnmpv3Credentials sets UseSnmpv3Credentials field to given value.

### HasUseSnmpv3Credentials

`func (o *DiscoveryMemberproperties) HasUseSnmpv3Credentials() bool`

HasUseSnmpv3Credentials returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


