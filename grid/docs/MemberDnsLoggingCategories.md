# MemberDnsLoggingCategories

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LogDtcGslb** | Pointer to **bool** | Determines whether the DTC GSLB activity is captured or not. | [optional] 
**LogDtcHealth** | Pointer to **bool** | Determines whether the DTC health monitoring information is captured or not. | [optional] 
**LogGeneral** | Pointer to **bool** | Determines whether the BIND messages that are not specifically classified are captured or not. | [optional] 
**LogClient** | Pointer to **bool** | Determines whether the client requests are captured or not. | [optional] 
**LogConfig** | Pointer to **bool** | Determines whether the configuration file parsing is captured or not. | [optional] 
**LogDatabase** | Pointer to **bool** | Determines whether the BIND&#39;s internal database processes are captured or not. | [optional] 
**LogDnssec** | Pointer to **bool** | Determines whether the DNSSEC-signed responses are captured or not. | [optional] 
**LogLameServers** | Pointer to **bool** | Determines whether the bad delegation instances are captured or not. | [optional] 
**LogNetwork** | Pointer to **bool** | Determines whether the network operation messages are captured or not. | [optional] 
**LogNotify** | Pointer to **bool** | Determines whether the asynchronous zone change notification messages are captured or not. | [optional] 
**LogQueries** | Pointer to **bool** | Determines whether the query messages are captured or not. | [optional] 
**LogQueryRewrite** | Pointer to **bool** | Determines whether the query rewrite messages are captured or not. | [optional] 
**LogResponses** | Pointer to **bool** | Determines whether the response messages are captured or not. | [optional] 
**LogResolver** | Pointer to **bool** | Determines whether the DNS resolution instances, including recursive queries from resolvers are captured or not. | [optional] 
**LogSecurity** | Pointer to **bool** | Determines whether the approved and denied requests are captured or not. | [optional] 
**LogUpdate** | Pointer to **bool** | Determines whether the dynamic update instances are captured or not. | [optional] 
**LogXferIn** | Pointer to **bool** | Determines whether the zone transfer messages from the remote name servers to the appliance are captured or not. | [optional] 
**LogXferOut** | Pointer to **bool** | Determines whether the zone transfer messages from the Infoblox appliance to remote name servers are captured or not. | [optional] 
**LogUpdateSecurity** | Pointer to **bool** | Determines whether the security update messages are captured or not. | [optional] 
**LogRateLimit** | Pointer to **bool** | Determines whether the rate limit messages are captured or not. | [optional] 
**LogRpz** | Pointer to **bool** | Determines whether the Response Policy Zone messages are captured or not. | [optional] 

## Methods

### NewMemberDnsLoggingCategories

`func NewMemberDnsLoggingCategories() *MemberDnsLoggingCategories`

NewMemberDnsLoggingCategories instantiates a new MemberDnsLoggingCategories object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberDnsLoggingCategoriesWithDefaults

`func NewMemberDnsLoggingCategoriesWithDefaults() *MemberDnsLoggingCategories`

NewMemberDnsLoggingCategoriesWithDefaults instantiates a new MemberDnsLoggingCategories object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogDtcGslb

`func (o *MemberDnsLoggingCategories) GetLogDtcGslb() bool`

GetLogDtcGslb returns the LogDtcGslb field if non-nil, zero value otherwise.

### GetLogDtcGslbOk

`func (o *MemberDnsLoggingCategories) GetLogDtcGslbOk() (*bool, bool)`

GetLogDtcGslbOk returns a tuple with the LogDtcGslb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogDtcGslb

`func (o *MemberDnsLoggingCategories) SetLogDtcGslb(v bool)`

SetLogDtcGslb sets LogDtcGslb field to given value.

### HasLogDtcGslb

`func (o *MemberDnsLoggingCategories) HasLogDtcGslb() bool`

HasLogDtcGslb returns a boolean if a field has been set.

### GetLogDtcHealth

`func (o *MemberDnsLoggingCategories) GetLogDtcHealth() bool`

GetLogDtcHealth returns the LogDtcHealth field if non-nil, zero value otherwise.

### GetLogDtcHealthOk

`func (o *MemberDnsLoggingCategories) GetLogDtcHealthOk() (*bool, bool)`

GetLogDtcHealthOk returns a tuple with the LogDtcHealth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogDtcHealth

`func (o *MemberDnsLoggingCategories) SetLogDtcHealth(v bool)`

SetLogDtcHealth sets LogDtcHealth field to given value.

### HasLogDtcHealth

`func (o *MemberDnsLoggingCategories) HasLogDtcHealth() bool`

HasLogDtcHealth returns a boolean if a field has been set.

### GetLogGeneral

`func (o *MemberDnsLoggingCategories) GetLogGeneral() bool`

GetLogGeneral returns the LogGeneral field if non-nil, zero value otherwise.

### GetLogGeneralOk

`func (o *MemberDnsLoggingCategories) GetLogGeneralOk() (*bool, bool)`

GetLogGeneralOk returns a tuple with the LogGeneral field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogGeneral

`func (o *MemberDnsLoggingCategories) SetLogGeneral(v bool)`

SetLogGeneral sets LogGeneral field to given value.

### HasLogGeneral

`func (o *MemberDnsLoggingCategories) HasLogGeneral() bool`

HasLogGeneral returns a boolean if a field has been set.

### GetLogClient

`func (o *MemberDnsLoggingCategories) GetLogClient() bool`

GetLogClient returns the LogClient field if non-nil, zero value otherwise.

### GetLogClientOk

`func (o *MemberDnsLoggingCategories) GetLogClientOk() (*bool, bool)`

GetLogClientOk returns a tuple with the LogClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogClient

`func (o *MemberDnsLoggingCategories) SetLogClient(v bool)`

SetLogClient sets LogClient field to given value.

### HasLogClient

`func (o *MemberDnsLoggingCategories) HasLogClient() bool`

HasLogClient returns a boolean if a field has been set.

### GetLogConfig

`func (o *MemberDnsLoggingCategories) GetLogConfig() bool`

GetLogConfig returns the LogConfig field if non-nil, zero value otherwise.

### GetLogConfigOk

`func (o *MemberDnsLoggingCategories) GetLogConfigOk() (*bool, bool)`

GetLogConfigOk returns a tuple with the LogConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogConfig

`func (o *MemberDnsLoggingCategories) SetLogConfig(v bool)`

SetLogConfig sets LogConfig field to given value.

### HasLogConfig

`func (o *MemberDnsLoggingCategories) HasLogConfig() bool`

HasLogConfig returns a boolean if a field has been set.

### GetLogDatabase

`func (o *MemberDnsLoggingCategories) GetLogDatabase() bool`

GetLogDatabase returns the LogDatabase field if non-nil, zero value otherwise.

### GetLogDatabaseOk

`func (o *MemberDnsLoggingCategories) GetLogDatabaseOk() (*bool, bool)`

GetLogDatabaseOk returns a tuple with the LogDatabase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogDatabase

`func (o *MemberDnsLoggingCategories) SetLogDatabase(v bool)`

SetLogDatabase sets LogDatabase field to given value.

### HasLogDatabase

`func (o *MemberDnsLoggingCategories) HasLogDatabase() bool`

HasLogDatabase returns a boolean if a field has been set.

### GetLogDnssec

`func (o *MemberDnsLoggingCategories) GetLogDnssec() bool`

GetLogDnssec returns the LogDnssec field if non-nil, zero value otherwise.

### GetLogDnssecOk

`func (o *MemberDnsLoggingCategories) GetLogDnssecOk() (*bool, bool)`

GetLogDnssecOk returns a tuple with the LogDnssec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogDnssec

`func (o *MemberDnsLoggingCategories) SetLogDnssec(v bool)`

SetLogDnssec sets LogDnssec field to given value.

### HasLogDnssec

`func (o *MemberDnsLoggingCategories) HasLogDnssec() bool`

HasLogDnssec returns a boolean if a field has been set.

### GetLogLameServers

`func (o *MemberDnsLoggingCategories) GetLogLameServers() bool`

GetLogLameServers returns the LogLameServers field if non-nil, zero value otherwise.

### GetLogLameServersOk

`func (o *MemberDnsLoggingCategories) GetLogLameServersOk() (*bool, bool)`

GetLogLameServersOk returns a tuple with the LogLameServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogLameServers

`func (o *MemberDnsLoggingCategories) SetLogLameServers(v bool)`

SetLogLameServers sets LogLameServers field to given value.

### HasLogLameServers

`func (o *MemberDnsLoggingCategories) HasLogLameServers() bool`

HasLogLameServers returns a boolean if a field has been set.

### GetLogNetwork

`func (o *MemberDnsLoggingCategories) GetLogNetwork() bool`

GetLogNetwork returns the LogNetwork field if non-nil, zero value otherwise.

### GetLogNetworkOk

`func (o *MemberDnsLoggingCategories) GetLogNetworkOk() (*bool, bool)`

GetLogNetworkOk returns a tuple with the LogNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogNetwork

`func (o *MemberDnsLoggingCategories) SetLogNetwork(v bool)`

SetLogNetwork sets LogNetwork field to given value.

### HasLogNetwork

`func (o *MemberDnsLoggingCategories) HasLogNetwork() bool`

HasLogNetwork returns a boolean if a field has been set.

### GetLogNotify

`func (o *MemberDnsLoggingCategories) GetLogNotify() bool`

GetLogNotify returns the LogNotify field if non-nil, zero value otherwise.

### GetLogNotifyOk

`func (o *MemberDnsLoggingCategories) GetLogNotifyOk() (*bool, bool)`

GetLogNotifyOk returns a tuple with the LogNotify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogNotify

`func (o *MemberDnsLoggingCategories) SetLogNotify(v bool)`

SetLogNotify sets LogNotify field to given value.

### HasLogNotify

`func (o *MemberDnsLoggingCategories) HasLogNotify() bool`

HasLogNotify returns a boolean if a field has been set.

### GetLogQueries

`func (o *MemberDnsLoggingCategories) GetLogQueries() bool`

GetLogQueries returns the LogQueries field if non-nil, zero value otherwise.

### GetLogQueriesOk

`func (o *MemberDnsLoggingCategories) GetLogQueriesOk() (*bool, bool)`

GetLogQueriesOk returns a tuple with the LogQueries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogQueries

`func (o *MemberDnsLoggingCategories) SetLogQueries(v bool)`

SetLogQueries sets LogQueries field to given value.

### HasLogQueries

`func (o *MemberDnsLoggingCategories) HasLogQueries() bool`

HasLogQueries returns a boolean if a field has been set.

### GetLogQueryRewrite

`func (o *MemberDnsLoggingCategories) GetLogQueryRewrite() bool`

GetLogQueryRewrite returns the LogQueryRewrite field if non-nil, zero value otherwise.

### GetLogQueryRewriteOk

`func (o *MemberDnsLoggingCategories) GetLogQueryRewriteOk() (*bool, bool)`

GetLogQueryRewriteOk returns a tuple with the LogQueryRewrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogQueryRewrite

`func (o *MemberDnsLoggingCategories) SetLogQueryRewrite(v bool)`

SetLogQueryRewrite sets LogQueryRewrite field to given value.

### HasLogQueryRewrite

`func (o *MemberDnsLoggingCategories) HasLogQueryRewrite() bool`

HasLogQueryRewrite returns a boolean if a field has been set.

### GetLogResponses

`func (o *MemberDnsLoggingCategories) GetLogResponses() bool`

GetLogResponses returns the LogResponses field if non-nil, zero value otherwise.

### GetLogResponsesOk

`func (o *MemberDnsLoggingCategories) GetLogResponsesOk() (*bool, bool)`

GetLogResponsesOk returns a tuple with the LogResponses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogResponses

`func (o *MemberDnsLoggingCategories) SetLogResponses(v bool)`

SetLogResponses sets LogResponses field to given value.

### HasLogResponses

`func (o *MemberDnsLoggingCategories) HasLogResponses() bool`

HasLogResponses returns a boolean if a field has been set.

### GetLogResolver

`func (o *MemberDnsLoggingCategories) GetLogResolver() bool`

GetLogResolver returns the LogResolver field if non-nil, zero value otherwise.

### GetLogResolverOk

`func (o *MemberDnsLoggingCategories) GetLogResolverOk() (*bool, bool)`

GetLogResolverOk returns a tuple with the LogResolver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogResolver

`func (o *MemberDnsLoggingCategories) SetLogResolver(v bool)`

SetLogResolver sets LogResolver field to given value.

### HasLogResolver

`func (o *MemberDnsLoggingCategories) HasLogResolver() bool`

HasLogResolver returns a boolean if a field has been set.

### GetLogSecurity

`func (o *MemberDnsLoggingCategories) GetLogSecurity() bool`

GetLogSecurity returns the LogSecurity field if non-nil, zero value otherwise.

### GetLogSecurityOk

`func (o *MemberDnsLoggingCategories) GetLogSecurityOk() (*bool, bool)`

GetLogSecurityOk returns a tuple with the LogSecurity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogSecurity

`func (o *MemberDnsLoggingCategories) SetLogSecurity(v bool)`

SetLogSecurity sets LogSecurity field to given value.

### HasLogSecurity

`func (o *MemberDnsLoggingCategories) HasLogSecurity() bool`

HasLogSecurity returns a boolean if a field has been set.

### GetLogUpdate

`func (o *MemberDnsLoggingCategories) GetLogUpdate() bool`

GetLogUpdate returns the LogUpdate field if non-nil, zero value otherwise.

### GetLogUpdateOk

`func (o *MemberDnsLoggingCategories) GetLogUpdateOk() (*bool, bool)`

GetLogUpdateOk returns a tuple with the LogUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogUpdate

`func (o *MemberDnsLoggingCategories) SetLogUpdate(v bool)`

SetLogUpdate sets LogUpdate field to given value.

### HasLogUpdate

`func (o *MemberDnsLoggingCategories) HasLogUpdate() bool`

HasLogUpdate returns a boolean if a field has been set.

### GetLogXferIn

`func (o *MemberDnsLoggingCategories) GetLogXferIn() bool`

GetLogXferIn returns the LogXferIn field if non-nil, zero value otherwise.

### GetLogXferInOk

`func (o *MemberDnsLoggingCategories) GetLogXferInOk() (*bool, bool)`

GetLogXferInOk returns a tuple with the LogXferIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogXferIn

`func (o *MemberDnsLoggingCategories) SetLogXferIn(v bool)`

SetLogXferIn sets LogXferIn field to given value.

### HasLogXferIn

`func (o *MemberDnsLoggingCategories) HasLogXferIn() bool`

HasLogXferIn returns a boolean if a field has been set.

### GetLogXferOut

`func (o *MemberDnsLoggingCategories) GetLogXferOut() bool`

GetLogXferOut returns the LogXferOut field if non-nil, zero value otherwise.

### GetLogXferOutOk

`func (o *MemberDnsLoggingCategories) GetLogXferOutOk() (*bool, bool)`

GetLogXferOutOk returns a tuple with the LogXferOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogXferOut

`func (o *MemberDnsLoggingCategories) SetLogXferOut(v bool)`

SetLogXferOut sets LogXferOut field to given value.

### HasLogXferOut

`func (o *MemberDnsLoggingCategories) HasLogXferOut() bool`

HasLogXferOut returns a boolean if a field has been set.

### GetLogUpdateSecurity

`func (o *MemberDnsLoggingCategories) GetLogUpdateSecurity() bool`

GetLogUpdateSecurity returns the LogUpdateSecurity field if non-nil, zero value otherwise.

### GetLogUpdateSecurityOk

`func (o *MemberDnsLoggingCategories) GetLogUpdateSecurityOk() (*bool, bool)`

GetLogUpdateSecurityOk returns a tuple with the LogUpdateSecurity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogUpdateSecurity

`func (o *MemberDnsLoggingCategories) SetLogUpdateSecurity(v bool)`

SetLogUpdateSecurity sets LogUpdateSecurity field to given value.

### HasLogUpdateSecurity

`func (o *MemberDnsLoggingCategories) HasLogUpdateSecurity() bool`

HasLogUpdateSecurity returns a boolean if a field has been set.

### GetLogRateLimit

`func (o *MemberDnsLoggingCategories) GetLogRateLimit() bool`

GetLogRateLimit returns the LogRateLimit field if non-nil, zero value otherwise.

### GetLogRateLimitOk

`func (o *MemberDnsLoggingCategories) GetLogRateLimitOk() (*bool, bool)`

GetLogRateLimitOk returns a tuple with the LogRateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogRateLimit

`func (o *MemberDnsLoggingCategories) SetLogRateLimit(v bool)`

SetLogRateLimit sets LogRateLimit field to given value.

### HasLogRateLimit

`func (o *MemberDnsLoggingCategories) HasLogRateLimit() bool`

HasLogRateLimit returns a boolean if a field has been set.

### GetLogRpz

`func (o *MemberDnsLoggingCategories) GetLogRpz() bool`

GetLogRpz returns the LogRpz field if non-nil, zero value otherwise.

### GetLogRpzOk

`func (o *MemberDnsLoggingCategories) GetLogRpzOk() (*bool, bool)`

GetLogRpzOk returns a tuple with the LogRpz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogRpz

`func (o *MemberDnsLoggingCategories) SetLogRpz(v bool)`

SetLogRpz sets LogRpz field to given value.

### HasLogRpz

`func (o *MemberDnsLoggingCategories) HasLogRpz() bool`

HasLogRpz returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


