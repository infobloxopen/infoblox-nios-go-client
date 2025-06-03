# GridDnsLoggingCategories

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

### NewGridDnsLoggingCategories

`func NewGridDnsLoggingCategories() *GridDnsLoggingCategories`

NewGridDnsLoggingCategories instantiates a new GridDnsLoggingCategories object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGridDnsLoggingCategoriesWithDefaults

`func NewGridDnsLoggingCategoriesWithDefaults() *GridDnsLoggingCategories`

NewGridDnsLoggingCategoriesWithDefaults instantiates a new GridDnsLoggingCategories object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogDtcGslb

`func (o *GridDnsLoggingCategories) GetLogDtcGslb() bool`

GetLogDtcGslb returns the LogDtcGslb field if non-nil, zero value otherwise.

### GetLogDtcGslbOk

`func (o *GridDnsLoggingCategories) GetLogDtcGslbOk() (*bool, bool)`

GetLogDtcGslbOk returns a tuple with the LogDtcGslb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogDtcGslb

`func (o *GridDnsLoggingCategories) SetLogDtcGslb(v bool)`

SetLogDtcGslb sets LogDtcGslb field to given value.

### HasLogDtcGslb

`func (o *GridDnsLoggingCategories) HasLogDtcGslb() bool`

HasLogDtcGslb returns a boolean if a field has been set.

### GetLogDtcHealth

`func (o *GridDnsLoggingCategories) GetLogDtcHealth() bool`

GetLogDtcHealth returns the LogDtcHealth field if non-nil, zero value otherwise.

### GetLogDtcHealthOk

`func (o *GridDnsLoggingCategories) GetLogDtcHealthOk() (*bool, bool)`

GetLogDtcHealthOk returns a tuple with the LogDtcHealth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogDtcHealth

`func (o *GridDnsLoggingCategories) SetLogDtcHealth(v bool)`

SetLogDtcHealth sets LogDtcHealth field to given value.

### HasLogDtcHealth

`func (o *GridDnsLoggingCategories) HasLogDtcHealth() bool`

HasLogDtcHealth returns a boolean if a field has been set.

### GetLogGeneral

`func (o *GridDnsLoggingCategories) GetLogGeneral() bool`

GetLogGeneral returns the LogGeneral field if non-nil, zero value otherwise.

### GetLogGeneralOk

`func (o *GridDnsLoggingCategories) GetLogGeneralOk() (*bool, bool)`

GetLogGeneralOk returns a tuple with the LogGeneral field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogGeneral

`func (o *GridDnsLoggingCategories) SetLogGeneral(v bool)`

SetLogGeneral sets LogGeneral field to given value.

### HasLogGeneral

`func (o *GridDnsLoggingCategories) HasLogGeneral() bool`

HasLogGeneral returns a boolean if a field has been set.

### GetLogClient

`func (o *GridDnsLoggingCategories) GetLogClient() bool`

GetLogClient returns the LogClient field if non-nil, zero value otherwise.

### GetLogClientOk

`func (o *GridDnsLoggingCategories) GetLogClientOk() (*bool, bool)`

GetLogClientOk returns a tuple with the LogClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogClient

`func (o *GridDnsLoggingCategories) SetLogClient(v bool)`

SetLogClient sets LogClient field to given value.

### HasLogClient

`func (o *GridDnsLoggingCategories) HasLogClient() bool`

HasLogClient returns a boolean if a field has been set.

### GetLogConfig

`func (o *GridDnsLoggingCategories) GetLogConfig() bool`

GetLogConfig returns the LogConfig field if non-nil, zero value otherwise.

### GetLogConfigOk

`func (o *GridDnsLoggingCategories) GetLogConfigOk() (*bool, bool)`

GetLogConfigOk returns a tuple with the LogConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogConfig

`func (o *GridDnsLoggingCategories) SetLogConfig(v bool)`

SetLogConfig sets LogConfig field to given value.

### HasLogConfig

`func (o *GridDnsLoggingCategories) HasLogConfig() bool`

HasLogConfig returns a boolean if a field has been set.

### GetLogDatabase

`func (o *GridDnsLoggingCategories) GetLogDatabase() bool`

GetLogDatabase returns the LogDatabase field if non-nil, zero value otherwise.

### GetLogDatabaseOk

`func (o *GridDnsLoggingCategories) GetLogDatabaseOk() (*bool, bool)`

GetLogDatabaseOk returns a tuple with the LogDatabase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogDatabase

`func (o *GridDnsLoggingCategories) SetLogDatabase(v bool)`

SetLogDatabase sets LogDatabase field to given value.

### HasLogDatabase

`func (o *GridDnsLoggingCategories) HasLogDatabase() bool`

HasLogDatabase returns a boolean if a field has been set.

### GetLogDnssec

`func (o *GridDnsLoggingCategories) GetLogDnssec() bool`

GetLogDnssec returns the LogDnssec field if non-nil, zero value otherwise.

### GetLogDnssecOk

`func (o *GridDnsLoggingCategories) GetLogDnssecOk() (*bool, bool)`

GetLogDnssecOk returns a tuple with the LogDnssec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogDnssec

`func (o *GridDnsLoggingCategories) SetLogDnssec(v bool)`

SetLogDnssec sets LogDnssec field to given value.

### HasLogDnssec

`func (o *GridDnsLoggingCategories) HasLogDnssec() bool`

HasLogDnssec returns a boolean if a field has been set.

### GetLogLameServers

`func (o *GridDnsLoggingCategories) GetLogLameServers() bool`

GetLogLameServers returns the LogLameServers field if non-nil, zero value otherwise.

### GetLogLameServersOk

`func (o *GridDnsLoggingCategories) GetLogLameServersOk() (*bool, bool)`

GetLogLameServersOk returns a tuple with the LogLameServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogLameServers

`func (o *GridDnsLoggingCategories) SetLogLameServers(v bool)`

SetLogLameServers sets LogLameServers field to given value.

### HasLogLameServers

`func (o *GridDnsLoggingCategories) HasLogLameServers() bool`

HasLogLameServers returns a boolean if a field has been set.

### GetLogNetwork

`func (o *GridDnsLoggingCategories) GetLogNetwork() bool`

GetLogNetwork returns the LogNetwork field if non-nil, zero value otherwise.

### GetLogNetworkOk

`func (o *GridDnsLoggingCategories) GetLogNetworkOk() (*bool, bool)`

GetLogNetworkOk returns a tuple with the LogNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogNetwork

`func (o *GridDnsLoggingCategories) SetLogNetwork(v bool)`

SetLogNetwork sets LogNetwork field to given value.

### HasLogNetwork

`func (o *GridDnsLoggingCategories) HasLogNetwork() bool`

HasLogNetwork returns a boolean if a field has been set.

### GetLogNotify

`func (o *GridDnsLoggingCategories) GetLogNotify() bool`

GetLogNotify returns the LogNotify field if non-nil, zero value otherwise.

### GetLogNotifyOk

`func (o *GridDnsLoggingCategories) GetLogNotifyOk() (*bool, bool)`

GetLogNotifyOk returns a tuple with the LogNotify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogNotify

`func (o *GridDnsLoggingCategories) SetLogNotify(v bool)`

SetLogNotify sets LogNotify field to given value.

### HasLogNotify

`func (o *GridDnsLoggingCategories) HasLogNotify() bool`

HasLogNotify returns a boolean if a field has been set.

### GetLogQueries

`func (o *GridDnsLoggingCategories) GetLogQueries() bool`

GetLogQueries returns the LogQueries field if non-nil, zero value otherwise.

### GetLogQueriesOk

`func (o *GridDnsLoggingCategories) GetLogQueriesOk() (*bool, bool)`

GetLogQueriesOk returns a tuple with the LogQueries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogQueries

`func (o *GridDnsLoggingCategories) SetLogQueries(v bool)`

SetLogQueries sets LogQueries field to given value.

### HasLogQueries

`func (o *GridDnsLoggingCategories) HasLogQueries() bool`

HasLogQueries returns a boolean if a field has been set.

### GetLogQueryRewrite

`func (o *GridDnsLoggingCategories) GetLogQueryRewrite() bool`

GetLogQueryRewrite returns the LogQueryRewrite field if non-nil, zero value otherwise.

### GetLogQueryRewriteOk

`func (o *GridDnsLoggingCategories) GetLogQueryRewriteOk() (*bool, bool)`

GetLogQueryRewriteOk returns a tuple with the LogQueryRewrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogQueryRewrite

`func (o *GridDnsLoggingCategories) SetLogQueryRewrite(v bool)`

SetLogQueryRewrite sets LogQueryRewrite field to given value.

### HasLogQueryRewrite

`func (o *GridDnsLoggingCategories) HasLogQueryRewrite() bool`

HasLogQueryRewrite returns a boolean if a field has been set.

### GetLogResponses

`func (o *GridDnsLoggingCategories) GetLogResponses() bool`

GetLogResponses returns the LogResponses field if non-nil, zero value otherwise.

### GetLogResponsesOk

`func (o *GridDnsLoggingCategories) GetLogResponsesOk() (*bool, bool)`

GetLogResponsesOk returns a tuple with the LogResponses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogResponses

`func (o *GridDnsLoggingCategories) SetLogResponses(v bool)`

SetLogResponses sets LogResponses field to given value.

### HasLogResponses

`func (o *GridDnsLoggingCategories) HasLogResponses() bool`

HasLogResponses returns a boolean if a field has been set.

### GetLogResolver

`func (o *GridDnsLoggingCategories) GetLogResolver() bool`

GetLogResolver returns the LogResolver field if non-nil, zero value otherwise.

### GetLogResolverOk

`func (o *GridDnsLoggingCategories) GetLogResolverOk() (*bool, bool)`

GetLogResolverOk returns a tuple with the LogResolver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogResolver

`func (o *GridDnsLoggingCategories) SetLogResolver(v bool)`

SetLogResolver sets LogResolver field to given value.

### HasLogResolver

`func (o *GridDnsLoggingCategories) HasLogResolver() bool`

HasLogResolver returns a boolean if a field has been set.

### GetLogSecurity

`func (o *GridDnsLoggingCategories) GetLogSecurity() bool`

GetLogSecurity returns the LogSecurity field if non-nil, zero value otherwise.

### GetLogSecurityOk

`func (o *GridDnsLoggingCategories) GetLogSecurityOk() (*bool, bool)`

GetLogSecurityOk returns a tuple with the LogSecurity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogSecurity

`func (o *GridDnsLoggingCategories) SetLogSecurity(v bool)`

SetLogSecurity sets LogSecurity field to given value.

### HasLogSecurity

`func (o *GridDnsLoggingCategories) HasLogSecurity() bool`

HasLogSecurity returns a boolean if a field has been set.

### GetLogUpdate

`func (o *GridDnsLoggingCategories) GetLogUpdate() bool`

GetLogUpdate returns the LogUpdate field if non-nil, zero value otherwise.

### GetLogUpdateOk

`func (o *GridDnsLoggingCategories) GetLogUpdateOk() (*bool, bool)`

GetLogUpdateOk returns a tuple with the LogUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogUpdate

`func (o *GridDnsLoggingCategories) SetLogUpdate(v bool)`

SetLogUpdate sets LogUpdate field to given value.

### HasLogUpdate

`func (o *GridDnsLoggingCategories) HasLogUpdate() bool`

HasLogUpdate returns a boolean if a field has been set.

### GetLogXferIn

`func (o *GridDnsLoggingCategories) GetLogXferIn() bool`

GetLogXferIn returns the LogXferIn field if non-nil, zero value otherwise.

### GetLogXferInOk

`func (o *GridDnsLoggingCategories) GetLogXferInOk() (*bool, bool)`

GetLogXferInOk returns a tuple with the LogXferIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogXferIn

`func (o *GridDnsLoggingCategories) SetLogXferIn(v bool)`

SetLogXferIn sets LogXferIn field to given value.

### HasLogXferIn

`func (o *GridDnsLoggingCategories) HasLogXferIn() bool`

HasLogXferIn returns a boolean if a field has been set.

### GetLogXferOut

`func (o *GridDnsLoggingCategories) GetLogXferOut() bool`

GetLogXferOut returns the LogXferOut field if non-nil, zero value otherwise.

### GetLogXferOutOk

`func (o *GridDnsLoggingCategories) GetLogXferOutOk() (*bool, bool)`

GetLogXferOutOk returns a tuple with the LogXferOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogXferOut

`func (o *GridDnsLoggingCategories) SetLogXferOut(v bool)`

SetLogXferOut sets LogXferOut field to given value.

### HasLogXferOut

`func (o *GridDnsLoggingCategories) HasLogXferOut() bool`

HasLogXferOut returns a boolean if a field has been set.

### GetLogUpdateSecurity

`func (o *GridDnsLoggingCategories) GetLogUpdateSecurity() bool`

GetLogUpdateSecurity returns the LogUpdateSecurity field if non-nil, zero value otherwise.

### GetLogUpdateSecurityOk

`func (o *GridDnsLoggingCategories) GetLogUpdateSecurityOk() (*bool, bool)`

GetLogUpdateSecurityOk returns a tuple with the LogUpdateSecurity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogUpdateSecurity

`func (o *GridDnsLoggingCategories) SetLogUpdateSecurity(v bool)`

SetLogUpdateSecurity sets LogUpdateSecurity field to given value.

### HasLogUpdateSecurity

`func (o *GridDnsLoggingCategories) HasLogUpdateSecurity() bool`

HasLogUpdateSecurity returns a boolean if a field has been set.

### GetLogRateLimit

`func (o *GridDnsLoggingCategories) GetLogRateLimit() bool`

GetLogRateLimit returns the LogRateLimit field if non-nil, zero value otherwise.

### GetLogRateLimitOk

`func (o *GridDnsLoggingCategories) GetLogRateLimitOk() (*bool, bool)`

GetLogRateLimitOk returns a tuple with the LogRateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogRateLimit

`func (o *GridDnsLoggingCategories) SetLogRateLimit(v bool)`

SetLogRateLimit sets LogRateLimit field to given value.

### HasLogRateLimit

`func (o *GridDnsLoggingCategories) HasLogRateLimit() bool`

HasLogRateLimit returns a boolean if a field has been set.

### GetLogRpz

`func (o *GridDnsLoggingCategories) GetLogRpz() bool`

GetLogRpz returns the LogRpz field if non-nil, zero value otherwise.

### GetLogRpzOk

`func (o *GridDnsLoggingCategories) GetLogRpzOk() (*bool, bool)`

GetLogRpzOk returns a tuple with the LogRpz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogRpz

`func (o *GridDnsLoggingCategories) SetLogRpz(v bool)`

SetLogRpz sets LogRpz field to given value.

### HasLogRpz

`func (o *GridDnsLoggingCategories) HasLogRpz() bool`

HasLogRpz returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


