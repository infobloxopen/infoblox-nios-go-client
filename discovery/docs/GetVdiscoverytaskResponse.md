# GetVdiscoverytaskResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**AllowUnsecuredConnection** | Pointer to **bool** | Allow unsecured connection over HTTPS and bypass validation of the remote SSL certificate. | [optional] 
**AutoConsolidateCloudEa** | Pointer to **bool** | Whether to insert or update cloud EAs with discovery data. | [optional] 
**AutoConsolidateManagedTenant** | Pointer to **bool** | Whether to replace managed tenant with discovery tenant data. | [optional] 
**AutoConsolidateManagedVm** | Pointer to **bool** | Whether to replace managed virtual machine with discovery vm data. | [optional] 
**AutoCreateDnsHostnameTemplate** | Pointer to **string** | Template string used to generate host name. | [optional] 
**AutoCreateDnsRecord** | Pointer to **bool** | Control whether to create or update DNS record using discovered data. | [optional] 
**AutoCreateDnsRecordType** | Pointer to **string** | Indicates the type of record to create if the auto create DNS record is enabled. | [optional] 
**Comment** | Pointer to **string** | Comment on the task. | [optional] 
**CredentialsType** | Pointer to **string** | Credentials type used for connecting to the cloud management platform. | [optional] 
**DnsViewPrivateIp** | Pointer to **string** | The DNS view name for private IPs. | [optional] 
**DnsViewPublicIp** | Pointer to **string** | The DNS view name for public IPs. | [optional] 
**DomainName** | Pointer to **string** | The name of the domain to use with keystone v3. | [optional] 
**DriverType** | Pointer to **string** | Type of discovery driver. | [optional] 
**EnableFilter** | Pointer to **bool** | Enable filter for cloud discovery task | [optional] 
**Enabled** | Pointer to **bool** | Whether to enabled the cloud discovery or not. | [optional] 
**FqdnOrIp** | Pointer to **string** | FQDN or IP of the cloud management platform. | [optional] 
**IdentityVersion** | Pointer to **string** | Identity service version. | [optional] 
**LastRun** | Pointer to **int64** | Timestamp of last run. | [optional] [readonly] 
**Member** | Pointer to **string** | Member on which cloud discovery will be run. | [optional] 
**MergeData** | Pointer to **bool** | Whether to replace the old data with new or not. | [optional] 
**Name** | Pointer to **string** | Name of this cloud discovery task. Uniquely identify a task. | [optional] 
**NetworkFilter** | Pointer to **string** | Options to filter the networks in cdiscovery task. | [optional] 
**NetworkList** | Pointer to **[]string** | List of networks to filter in cdiscovery task. | [optional] 
**Password** | Pointer to **string** | Password used for connecting to the cloud management platform. | [optional] 
**Port** | Pointer to **int64** | Connection port used for connecting to the cloud management platform. | [optional] 
**PrivateNetworkView** | Pointer to **string** | Network view for private IPs. | [optional] 
**PrivateNetworkViewMappingPolicy** | Pointer to **string** | Mapping policy for the network view for private IPs in discovery data. | [optional] 
**Protocol** | Pointer to **string** | Connection protocol used for connecting to the cloud management platform. | [optional] 
**PublicNetworkView** | Pointer to **string** | Network view for public IPs. | [optional] 
**PublicNetworkViewMappingPolicy** | Pointer to **string** | Mapping policy for the network view for public IPs in discovery data. | [optional] 
**ScheduledRun** | Pointer to [**VdiscoverytaskScheduledRun**](VdiscoverytaskScheduledRun.md) |  | [optional] 
**ServiceAccountFile** | Pointer to **string** | The service_account_file for GCP. | [optional] 
**State** | Pointer to **string** | Current state of this task. | [optional] [readonly] 
**StateMsg** | Pointer to **string** | State message of the complete discovery process. | [optional] [readonly] 
**UpdateDnsViewPrivateIp** | Pointer to **bool** | If set to true, the appliance uses a specific DNS view for private IPs. | [optional] 
**UpdateDnsViewPublicIp** | Pointer to **bool** | If set to true, the appliance uses a specific DNS view for public IPs. | [optional] 
**UpdateMetadata** | Pointer to **bool** | Whether to update metadata as a result of this network discovery. | [optional] 
**UseIdentity** | Pointer to **bool** | If set true, all keystone connection will use \&quot;/identity\&quot; endpoint and port value will be ignored. | [optional] 
**Username** | Pointer to **string** | Username used for connecting to the cloud management platform. | [optional] 
**VdiscoveryControl** | Pointer to **map[string]interface{}** |  | [optional] 
**Result** | Pointer to [**Vdiscoverytask**](Vdiscoverytask.md) |  | [optional] 

## Methods

### NewGetVdiscoverytaskResponse

`func NewGetVdiscoverytaskResponse() *GetVdiscoverytaskResponse`

NewGetVdiscoverytaskResponse instantiates a new GetVdiscoverytaskResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetVdiscoverytaskResponseWithDefaults

`func NewGetVdiscoverytaskResponseWithDefaults() *GetVdiscoverytaskResponse`

NewGetVdiscoverytaskResponseWithDefaults instantiates a new GetVdiscoverytaskResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetVdiscoverytaskResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetVdiscoverytaskResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetVdiscoverytaskResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetVdiscoverytaskResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAllowUnsecuredConnection

`func (o *GetVdiscoverytaskResponse) GetAllowUnsecuredConnection() bool`

GetAllowUnsecuredConnection returns the AllowUnsecuredConnection field if non-nil, zero value otherwise.

### GetAllowUnsecuredConnectionOk

`func (o *GetVdiscoverytaskResponse) GetAllowUnsecuredConnectionOk() (*bool, bool)`

GetAllowUnsecuredConnectionOk returns a tuple with the AllowUnsecuredConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUnsecuredConnection

`func (o *GetVdiscoverytaskResponse) SetAllowUnsecuredConnection(v bool)`

SetAllowUnsecuredConnection sets AllowUnsecuredConnection field to given value.

### HasAllowUnsecuredConnection

`func (o *GetVdiscoverytaskResponse) HasAllowUnsecuredConnection() bool`

HasAllowUnsecuredConnection returns a boolean if a field has been set.

### GetAutoConsolidateCloudEa

`func (o *GetVdiscoverytaskResponse) GetAutoConsolidateCloudEa() bool`

GetAutoConsolidateCloudEa returns the AutoConsolidateCloudEa field if non-nil, zero value otherwise.

### GetAutoConsolidateCloudEaOk

`func (o *GetVdiscoverytaskResponse) GetAutoConsolidateCloudEaOk() (*bool, bool)`

GetAutoConsolidateCloudEaOk returns a tuple with the AutoConsolidateCloudEa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoConsolidateCloudEa

`func (o *GetVdiscoverytaskResponse) SetAutoConsolidateCloudEa(v bool)`

SetAutoConsolidateCloudEa sets AutoConsolidateCloudEa field to given value.

### HasAutoConsolidateCloudEa

`func (o *GetVdiscoverytaskResponse) HasAutoConsolidateCloudEa() bool`

HasAutoConsolidateCloudEa returns a boolean if a field has been set.

### GetAutoConsolidateManagedTenant

`func (o *GetVdiscoverytaskResponse) GetAutoConsolidateManagedTenant() bool`

GetAutoConsolidateManagedTenant returns the AutoConsolidateManagedTenant field if non-nil, zero value otherwise.

### GetAutoConsolidateManagedTenantOk

`func (o *GetVdiscoverytaskResponse) GetAutoConsolidateManagedTenantOk() (*bool, bool)`

GetAutoConsolidateManagedTenantOk returns a tuple with the AutoConsolidateManagedTenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoConsolidateManagedTenant

`func (o *GetVdiscoverytaskResponse) SetAutoConsolidateManagedTenant(v bool)`

SetAutoConsolidateManagedTenant sets AutoConsolidateManagedTenant field to given value.

### HasAutoConsolidateManagedTenant

`func (o *GetVdiscoverytaskResponse) HasAutoConsolidateManagedTenant() bool`

HasAutoConsolidateManagedTenant returns a boolean if a field has been set.

### GetAutoConsolidateManagedVm

`func (o *GetVdiscoverytaskResponse) GetAutoConsolidateManagedVm() bool`

GetAutoConsolidateManagedVm returns the AutoConsolidateManagedVm field if non-nil, zero value otherwise.

### GetAutoConsolidateManagedVmOk

`func (o *GetVdiscoverytaskResponse) GetAutoConsolidateManagedVmOk() (*bool, bool)`

GetAutoConsolidateManagedVmOk returns a tuple with the AutoConsolidateManagedVm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoConsolidateManagedVm

`func (o *GetVdiscoverytaskResponse) SetAutoConsolidateManagedVm(v bool)`

SetAutoConsolidateManagedVm sets AutoConsolidateManagedVm field to given value.

### HasAutoConsolidateManagedVm

`func (o *GetVdiscoverytaskResponse) HasAutoConsolidateManagedVm() bool`

HasAutoConsolidateManagedVm returns a boolean if a field has been set.

### GetAutoCreateDnsHostnameTemplate

`func (o *GetVdiscoverytaskResponse) GetAutoCreateDnsHostnameTemplate() string`

GetAutoCreateDnsHostnameTemplate returns the AutoCreateDnsHostnameTemplate field if non-nil, zero value otherwise.

### GetAutoCreateDnsHostnameTemplateOk

`func (o *GetVdiscoverytaskResponse) GetAutoCreateDnsHostnameTemplateOk() (*string, bool)`

GetAutoCreateDnsHostnameTemplateOk returns a tuple with the AutoCreateDnsHostnameTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCreateDnsHostnameTemplate

`func (o *GetVdiscoverytaskResponse) SetAutoCreateDnsHostnameTemplate(v string)`

SetAutoCreateDnsHostnameTemplate sets AutoCreateDnsHostnameTemplate field to given value.

### HasAutoCreateDnsHostnameTemplate

`func (o *GetVdiscoverytaskResponse) HasAutoCreateDnsHostnameTemplate() bool`

HasAutoCreateDnsHostnameTemplate returns a boolean if a field has been set.

### GetAutoCreateDnsRecord

`func (o *GetVdiscoverytaskResponse) GetAutoCreateDnsRecord() bool`

GetAutoCreateDnsRecord returns the AutoCreateDnsRecord field if non-nil, zero value otherwise.

### GetAutoCreateDnsRecordOk

`func (o *GetVdiscoverytaskResponse) GetAutoCreateDnsRecordOk() (*bool, bool)`

GetAutoCreateDnsRecordOk returns a tuple with the AutoCreateDnsRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCreateDnsRecord

`func (o *GetVdiscoverytaskResponse) SetAutoCreateDnsRecord(v bool)`

SetAutoCreateDnsRecord sets AutoCreateDnsRecord field to given value.

### HasAutoCreateDnsRecord

`func (o *GetVdiscoverytaskResponse) HasAutoCreateDnsRecord() bool`

HasAutoCreateDnsRecord returns a boolean if a field has been set.

### GetAutoCreateDnsRecordType

`func (o *GetVdiscoverytaskResponse) GetAutoCreateDnsRecordType() string`

GetAutoCreateDnsRecordType returns the AutoCreateDnsRecordType field if non-nil, zero value otherwise.

### GetAutoCreateDnsRecordTypeOk

`func (o *GetVdiscoverytaskResponse) GetAutoCreateDnsRecordTypeOk() (*string, bool)`

GetAutoCreateDnsRecordTypeOk returns a tuple with the AutoCreateDnsRecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCreateDnsRecordType

`func (o *GetVdiscoverytaskResponse) SetAutoCreateDnsRecordType(v string)`

SetAutoCreateDnsRecordType sets AutoCreateDnsRecordType field to given value.

### HasAutoCreateDnsRecordType

`func (o *GetVdiscoverytaskResponse) HasAutoCreateDnsRecordType() bool`

HasAutoCreateDnsRecordType returns a boolean if a field has been set.

### GetComment

`func (o *GetVdiscoverytaskResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetVdiscoverytaskResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetVdiscoverytaskResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetVdiscoverytaskResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCredentialsType

`func (o *GetVdiscoverytaskResponse) GetCredentialsType() string`

GetCredentialsType returns the CredentialsType field if non-nil, zero value otherwise.

### GetCredentialsTypeOk

`func (o *GetVdiscoverytaskResponse) GetCredentialsTypeOk() (*string, bool)`

GetCredentialsTypeOk returns a tuple with the CredentialsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialsType

`func (o *GetVdiscoverytaskResponse) SetCredentialsType(v string)`

SetCredentialsType sets CredentialsType field to given value.

### HasCredentialsType

`func (o *GetVdiscoverytaskResponse) HasCredentialsType() bool`

HasCredentialsType returns a boolean if a field has been set.

### GetDnsViewPrivateIp

`func (o *GetVdiscoverytaskResponse) GetDnsViewPrivateIp() string`

GetDnsViewPrivateIp returns the DnsViewPrivateIp field if non-nil, zero value otherwise.

### GetDnsViewPrivateIpOk

`func (o *GetVdiscoverytaskResponse) GetDnsViewPrivateIpOk() (*string, bool)`

GetDnsViewPrivateIpOk returns a tuple with the DnsViewPrivateIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsViewPrivateIp

`func (o *GetVdiscoverytaskResponse) SetDnsViewPrivateIp(v string)`

SetDnsViewPrivateIp sets DnsViewPrivateIp field to given value.

### HasDnsViewPrivateIp

`func (o *GetVdiscoverytaskResponse) HasDnsViewPrivateIp() bool`

HasDnsViewPrivateIp returns a boolean if a field has been set.

### GetDnsViewPublicIp

`func (o *GetVdiscoverytaskResponse) GetDnsViewPublicIp() string`

GetDnsViewPublicIp returns the DnsViewPublicIp field if non-nil, zero value otherwise.

### GetDnsViewPublicIpOk

`func (o *GetVdiscoverytaskResponse) GetDnsViewPublicIpOk() (*string, bool)`

GetDnsViewPublicIpOk returns a tuple with the DnsViewPublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsViewPublicIp

`func (o *GetVdiscoverytaskResponse) SetDnsViewPublicIp(v string)`

SetDnsViewPublicIp sets DnsViewPublicIp field to given value.

### HasDnsViewPublicIp

`func (o *GetVdiscoverytaskResponse) HasDnsViewPublicIp() bool`

HasDnsViewPublicIp returns a boolean if a field has been set.

### GetDomainName

`func (o *GetVdiscoverytaskResponse) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *GetVdiscoverytaskResponse) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *GetVdiscoverytaskResponse) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *GetVdiscoverytaskResponse) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetDriverType

`func (o *GetVdiscoverytaskResponse) GetDriverType() string`

GetDriverType returns the DriverType field if non-nil, zero value otherwise.

### GetDriverTypeOk

`func (o *GetVdiscoverytaskResponse) GetDriverTypeOk() (*string, bool)`

GetDriverTypeOk returns a tuple with the DriverType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverType

`func (o *GetVdiscoverytaskResponse) SetDriverType(v string)`

SetDriverType sets DriverType field to given value.

### HasDriverType

`func (o *GetVdiscoverytaskResponse) HasDriverType() bool`

HasDriverType returns a boolean if a field has been set.

### GetEnableFilter

`func (o *GetVdiscoverytaskResponse) GetEnableFilter() bool`

GetEnableFilter returns the EnableFilter field if non-nil, zero value otherwise.

### GetEnableFilterOk

`func (o *GetVdiscoverytaskResponse) GetEnableFilterOk() (*bool, bool)`

GetEnableFilterOk returns a tuple with the EnableFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableFilter

`func (o *GetVdiscoverytaskResponse) SetEnableFilter(v bool)`

SetEnableFilter sets EnableFilter field to given value.

### HasEnableFilter

`func (o *GetVdiscoverytaskResponse) HasEnableFilter() bool`

HasEnableFilter returns a boolean if a field has been set.

### GetEnabled

`func (o *GetVdiscoverytaskResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetVdiscoverytaskResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetVdiscoverytaskResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetVdiscoverytaskResponse) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetFqdnOrIp

`func (o *GetVdiscoverytaskResponse) GetFqdnOrIp() string`

GetFqdnOrIp returns the FqdnOrIp field if non-nil, zero value otherwise.

### GetFqdnOrIpOk

`func (o *GetVdiscoverytaskResponse) GetFqdnOrIpOk() (*string, bool)`

GetFqdnOrIpOk returns a tuple with the FqdnOrIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdnOrIp

`func (o *GetVdiscoverytaskResponse) SetFqdnOrIp(v string)`

SetFqdnOrIp sets FqdnOrIp field to given value.

### HasFqdnOrIp

`func (o *GetVdiscoverytaskResponse) HasFqdnOrIp() bool`

HasFqdnOrIp returns a boolean if a field has been set.

### GetIdentityVersion

`func (o *GetVdiscoverytaskResponse) GetIdentityVersion() string`

GetIdentityVersion returns the IdentityVersion field if non-nil, zero value otherwise.

### GetIdentityVersionOk

`func (o *GetVdiscoverytaskResponse) GetIdentityVersionOk() (*string, bool)`

GetIdentityVersionOk returns a tuple with the IdentityVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityVersion

`func (o *GetVdiscoverytaskResponse) SetIdentityVersion(v string)`

SetIdentityVersion sets IdentityVersion field to given value.

### HasIdentityVersion

`func (o *GetVdiscoverytaskResponse) HasIdentityVersion() bool`

HasIdentityVersion returns a boolean if a field has been set.

### GetLastRun

`func (o *GetVdiscoverytaskResponse) GetLastRun() int64`

GetLastRun returns the LastRun field if non-nil, zero value otherwise.

### GetLastRunOk

`func (o *GetVdiscoverytaskResponse) GetLastRunOk() (*int64, bool)`

GetLastRunOk returns a tuple with the LastRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRun

`func (o *GetVdiscoverytaskResponse) SetLastRun(v int64)`

SetLastRun sets LastRun field to given value.

### HasLastRun

`func (o *GetVdiscoverytaskResponse) HasLastRun() bool`

HasLastRun returns a boolean if a field has been set.

### GetMember

`func (o *GetVdiscoverytaskResponse) GetMember() string`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *GetVdiscoverytaskResponse) GetMemberOk() (*string, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *GetVdiscoverytaskResponse) SetMember(v string)`

SetMember sets Member field to given value.

### HasMember

`func (o *GetVdiscoverytaskResponse) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetMergeData

`func (o *GetVdiscoverytaskResponse) GetMergeData() bool`

GetMergeData returns the MergeData field if non-nil, zero value otherwise.

### GetMergeDataOk

`func (o *GetVdiscoverytaskResponse) GetMergeDataOk() (*bool, bool)`

GetMergeDataOk returns a tuple with the MergeData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeData

`func (o *GetVdiscoverytaskResponse) SetMergeData(v bool)`

SetMergeData sets MergeData field to given value.

### HasMergeData

`func (o *GetVdiscoverytaskResponse) HasMergeData() bool`

HasMergeData returns a boolean if a field has been set.

### GetName

`func (o *GetVdiscoverytaskResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetVdiscoverytaskResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetVdiscoverytaskResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetVdiscoverytaskResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetworkFilter

`func (o *GetVdiscoverytaskResponse) GetNetworkFilter() string`

GetNetworkFilter returns the NetworkFilter field if non-nil, zero value otherwise.

### GetNetworkFilterOk

`func (o *GetVdiscoverytaskResponse) GetNetworkFilterOk() (*string, bool)`

GetNetworkFilterOk returns a tuple with the NetworkFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkFilter

`func (o *GetVdiscoverytaskResponse) SetNetworkFilter(v string)`

SetNetworkFilter sets NetworkFilter field to given value.

### HasNetworkFilter

`func (o *GetVdiscoverytaskResponse) HasNetworkFilter() bool`

HasNetworkFilter returns a boolean if a field has been set.

### GetNetworkList

`func (o *GetVdiscoverytaskResponse) GetNetworkList() []string`

GetNetworkList returns the NetworkList field if non-nil, zero value otherwise.

### GetNetworkListOk

`func (o *GetVdiscoverytaskResponse) GetNetworkListOk() (*[]string, bool)`

GetNetworkListOk returns a tuple with the NetworkList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkList

`func (o *GetVdiscoverytaskResponse) SetNetworkList(v []string)`

SetNetworkList sets NetworkList field to given value.

### HasNetworkList

`func (o *GetVdiscoverytaskResponse) HasNetworkList() bool`

HasNetworkList returns a boolean if a field has been set.

### GetPassword

`func (o *GetVdiscoverytaskResponse) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *GetVdiscoverytaskResponse) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *GetVdiscoverytaskResponse) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *GetVdiscoverytaskResponse) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPort

`func (o *GetVdiscoverytaskResponse) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *GetVdiscoverytaskResponse) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *GetVdiscoverytaskResponse) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *GetVdiscoverytaskResponse) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetPrivateNetworkView

`func (o *GetVdiscoverytaskResponse) GetPrivateNetworkView() string`

GetPrivateNetworkView returns the PrivateNetworkView field if non-nil, zero value otherwise.

### GetPrivateNetworkViewOk

`func (o *GetVdiscoverytaskResponse) GetPrivateNetworkViewOk() (*string, bool)`

GetPrivateNetworkViewOk returns a tuple with the PrivateNetworkView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateNetworkView

`func (o *GetVdiscoverytaskResponse) SetPrivateNetworkView(v string)`

SetPrivateNetworkView sets PrivateNetworkView field to given value.

### HasPrivateNetworkView

`func (o *GetVdiscoverytaskResponse) HasPrivateNetworkView() bool`

HasPrivateNetworkView returns a boolean if a field has been set.

### GetPrivateNetworkViewMappingPolicy

`func (o *GetVdiscoverytaskResponse) GetPrivateNetworkViewMappingPolicy() string`

GetPrivateNetworkViewMappingPolicy returns the PrivateNetworkViewMappingPolicy field if non-nil, zero value otherwise.

### GetPrivateNetworkViewMappingPolicyOk

`func (o *GetVdiscoverytaskResponse) GetPrivateNetworkViewMappingPolicyOk() (*string, bool)`

GetPrivateNetworkViewMappingPolicyOk returns a tuple with the PrivateNetworkViewMappingPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateNetworkViewMappingPolicy

`func (o *GetVdiscoverytaskResponse) SetPrivateNetworkViewMappingPolicy(v string)`

SetPrivateNetworkViewMappingPolicy sets PrivateNetworkViewMappingPolicy field to given value.

### HasPrivateNetworkViewMappingPolicy

`func (o *GetVdiscoverytaskResponse) HasPrivateNetworkViewMappingPolicy() bool`

HasPrivateNetworkViewMappingPolicy returns a boolean if a field has been set.

### GetProtocol

`func (o *GetVdiscoverytaskResponse) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *GetVdiscoverytaskResponse) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *GetVdiscoverytaskResponse) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *GetVdiscoverytaskResponse) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetPublicNetworkView

`func (o *GetVdiscoverytaskResponse) GetPublicNetworkView() string`

GetPublicNetworkView returns the PublicNetworkView field if non-nil, zero value otherwise.

### GetPublicNetworkViewOk

`func (o *GetVdiscoverytaskResponse) GetPublicNetworkViewOk() (*string, bool)`

GetPublicNetworkViewOk returns a tuple with the PublicNetworkView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicNetworkView

`func (o *GetVdiscoverytaskResponse) SetPublicNetworkView(v string)`

SetPublicNetworkView sets PublicNetworkView field to given value.

### HasPublicNetworkView

`func (o *GetVdiscoverytaskResponse) HasPublicNetworkView() bool`

HasPublicNetworkView returns a boolean if a field has been set.

### GetPublicNetworkViewMappingPolicy

`func (o *GetVdiscoverytaskResponse) GetPublicNetworkViewMappingPolicy() string`

GetPublicNetworkViewMappingPolicy returns the PublicNetworkViewMappingPolicy field if non-nil, zero value otherwise.

### GetPublicNetworkViewMappingPolicyOk

`func (o *GetVdiscoverytaskResponse) GetPublicNetworkViewMappingPolicyOk() (*string, bool)`

GetPublicNetworkViewMappingPolicyOk returns a tuple with the PublicNetworkViewMappingPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicNetworkViewMappingPolicy

`func (o *GetVdiscoverytaskResponse) SetPublicNetworkViewMappingPolicy(v string)`

SetPublicNetworkViewMappingPolicy sets PublicNetworkViewMappingPolicy field to given value.

### HasPublicNetworkViewMappingPolicy

`func (o *GetVdiscoverytaskResponse) HasPublicNetworkViewMappingPolicy() bool`

HasPublicNetworkViewMappingPolicy returns a boolean if a field has been set.

### GetScheduledRun

`func (o *GetVdiscoverytaskResponse) GetScheduledRun() VdiscoverytaskScheduledRun`

GetScheduledRun returns the ScheduledRun field if non-nil, zero value otherwise.

### GetScheduledRunOk

`func (o *GetVdiscoverytaskResponse) GetScheduledRunOk() (*VdiscoverytaskScheduledRun, bool)`

GetScheduledRunOk returns a tuple with the ScheduledRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledRun

`func (o *GetVdiscoverytaskResponse) SetScheduledRun(v VdiscoverytaskScheduledRun)`

SetScheduledRun sets ScheduledRun field to given value.

### HasScheduledRun

`func (o *GetVdiscoverytaskResponse) HasScheduledRun() bool`

HasScheduledRun returns a boolean if a field has been set.

### GetServiceAccountFile

`func (o *GetVdiscoverytaskResponse) GetServiceAccountFile() string`

GetServiceAccountFile returns the ServiceAccountFile field if non-nil, zero value otherwise.

### GetServiceAccountFileOk

`func (o *GetVdiscoverytaskResponse) GetServiceAccountFileOk() (*string, bool)`

GetServiceAccountFileOk returns a tuple with the ServiceAccountFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountFile

`func (o *GetVdiscoverytaskResponse) SetServiceAccountFile(v string)`

SetServiceAccountFile sets ServiceAccountFile field to given value.

### HasServiceAccountFile

`func (o *GetVdiscoverytaskResponse) HasServiceAccountFile() bool`

HasServiceAccountFile returns a boolean if a field has been set.

### GetState

`func (o *GetVdiscoverytaskResponse) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GetVdiscoverytaskResponse) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GetVdiscoverytaskResponse) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *GetVdiscoverytaskResponse) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStateMsg

`func (o *GetVdiscoverytaskResponse) GetStateMsg() string`

GetStateMsg returns the StateMsg field if non-nil, zero value otherwise.

### GetStateMsgOk

`func (o *GetVdiscoverytaskResponse) GetStateMsgOk() (*string, bool)`

GetStateMsgOk returns a tuple with the StateMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateMsg

`func (o *GetVdiscoverytaskResponse) SetStateMsg(v string)`

SetStateMsg sets StateMsg field to given value.

### HasStateMsg

`func (o *GetVdiscoverytaskResponse) HasStateMsg() bool`

HasStateMsg returns a boolean if a field has been set.

### GetUpdateDnsViewPrivateIp

`func (o *GetVdiscoverytaskResponse) GetUpdateDnsViewPrivateIp() bool`

GetUpdateDnsViewPrivateIp returns the UpdateDnsViewPrivateIp field if non-nil, zero value otherwise.

### GetUpdateDnsViewPrivateIpOk

`func (o *GetVdiscoverytaskResponse) GetUpdateDnsViewPrivateIpOk() (*bool, bool)`

GetUpdateDnsViewPrivateIpOk returns a tuple with the UpdateDnsViewPrivateIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateDnsViewPrivateIp

`func (o *GetVdiscoverytaskResponse) SetUpdateDnsViewPrivateIp(v bool)`

SetUpdateDnsViewPrivateIp sets UpdateDnsViewPrivateIp field to given value.

### HasUpdateDnsViewPrivateIp

`func (o *GetVdiscoverytaskResponse) HasUpdateDnsViewPrivateIp() bool`

HasUpdateDnsViewPrivateIp returns a boolean if a field has been set.

### GetUpdateDnsViewPublicIp

`func (o *GetVdiscoverytaskResponse) GetUpdateDnsViewPublicIp() bool`

GetUpdateDnsViewPublicIp returns the UpdateDnsViewPublicIp field if non-nil, zero value otherwise.

### GetUpdateDnsViewPublicIpOk

`func (o *GetVdiscoverytaskResponse) GetUpdateDnsViewPublicIpOk() (*bool, bool)`

GetUpdateDnsViewPublicIpOk returns a tuple with the UpdateDnsViewPublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateDnsViewPublicIp

`func (o *GetVdiscoverytaskResponse) SetUpdateDnsViewPublicIp(v bool)`

SetUpdateDnsViewPublicIp sets UpdateDnsViewPublicIp field to given value.

### HasUpdateDnsViewPublicIp

`func (o *GetVdiscoverytaskResponse) HasUpdateDnsViewPublicIp() bool`

HasUpdateDnsViewPublicIp returns a boolean if a field has been set.

### GetUpdateMetadata

`func (o *GetVdiscoverytaskResponse) GetUpdateMetadata() bool`

GetUpdateMetadata returns the UpdateMetadata field if non-nil, zero value otherwise.

### GetUpdateMetadataOk

`func (o *GetVdiscoverytaskResponse) GetUpdateMetadataOk() (*bool, bool)`

GetUpdateMetadataOk returns a tuple with the UpdateMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateMetadata

`func (o *GetVdiscoverytaskResponse) SetUpdateMetadata(v bool)`

SetUpdateMetadata sets UpdateMetadata field to given value.

### HasUpdateMetadata

`func (o *GetVdiscoverytaskResponse) HasUpdateMetadata() bool`

HasUpdateMetadata returns a boolean if a field has been set.

### GetUseIdentity

`func (o *GetVdiscoverytaskResponse) GetUseIdentity() bool`

GetUseIdentity returns the UseIdentity field if non-nil, zero value otherwise.

### GetUseIdentityOk

`func (o *GetVdiscoverytaskResponse) GetUseIdentityOk() (*bool, bool)`

GetUseIdentityOk returns a tuple with the UseIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseIdentity

`func (o *GetVdiscoverytaskResponse) SetUseIdentity(v bool)`

SetUseIdentity sets UseIdentity field to given value.

### HasUseIdentity

`func (o *GetVdiscoverytaskResponse) HasUseIdentity() bool`

HasUseIdentity returns a boolean if a field has been set.

### GetUsername

`func (o *GetVdiscoverytaskResponse) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *GetVdiscoverytaskResponse) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *GetVdiscoverytaskResponse) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *GetVdiscoverytaskResponse) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetVdiscoveryControl

`func (o *GetVdiscoverytaskResponse) GetVdiscoveryControl() map[string]interface{}`

GetVdiscoveryControl returns the VdiscoveryControl field if non-nil, zero value otherwise.

### GetVdiscoveryControlOk

`func (o *GetVdiscoverytaskResponse) GetVdiscoveryControlOk() (*map[string]interface{}, bool)`

GetVdiscoveryControlOk returns a tuple with the VdiscoveryControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdiscoveryControl

`func (o *GetVdiscoverytaskResponse) SetVdiscoveryControl(v map[string]interface{})`

SetVdiscoveryControl sets VdiscoveryControl field to given value.

### HasVdiscoveryControl

`func (o *GetVdiscoverytaskResponse) HasVdiscoveryControl() bool`

HasVdiscoveryControl returns a boolean if a field has been set.

### GetResult

`func (o *GetVdiscoverytaskResponse) GetResult() Vdiscoverytask`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetVdiscoverytaskResponse) GetResultOk() (*Vdiscoverytask, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetVdiscoverytaskResponse) SetResult(v Vdiscoverytask)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetVdiscoverytaskResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


