# Infoblox NIOS Go Client

A Go client library for interacting with Infoblox NIOS via its WAPI (Web API). This library enables Go applications to automate DNS, DHCP, IPAM, and other network management tasks using Infoblox appliances.
The library is generated using the [OpenAPI Generator](https://openapi-generator.tech) project.

## Requirements

- Go (latest stable version recommended i.e. 1.24.4; minimum 1.18)
- Infoblox NIOS version 9.0.6 or higher

## Features

- Interact with Infoblox NIOS WAPI endpoints
- Manage DNS records, DHCP leases, IPAM objects, and more
- Authentication support (Basic Auth)
- Customizable HTTP client
- Error handling and response parsing

## Installation

```bash
go get github.com/infobloxopen/infoblox-nios-go-client
```

## Authentication

The client uses HTTP Basic Authentication by default. Ensure your credentials are correct and have appropriate permissions on the Infoblox appliance.


### Environment Variables

You can configure the authentication for the client using environment variables for convenience:

| Variable Name | Description                | Example                                |
|---------------|----------------------------|----------------------------------------|
| NIOS_HOST_URL | NIOS Grid URL              | `https://gridmaster.example.com` |
| NIOS_USERNAME | Username for authentication | `admin`                                |
| NIOS_PASSWORD | Password for authentication | `password`                             | |

### Using Configuration for the API Client

Instead of using environment variables, you can also configure the API client directly in your code. Here's an example of how to set up the configuration in Infoblox NIOS API client :

```go
import (
    "github.com/infobloxopen/infoblox-nios-go-client/client"
    "github.com/infobloxopen/infoblox-nios-go-client/client/option"
)

apiClient := client.NewAPIClient(
    option.WithNIOSHostUrl(NIOS_HOST_URL),
    option.WithNIOSUsername("username"),
    option.WithNIOSPassword("password"),
    option.WithDebug(true),
)
```

> Note: The Password is a secret and should be handled securely. Hardcoding the Password in your code is not recommended.

## Usage

You can either use an aggregated client to interact with multiple BloxOne APIs or create a client for a specific API.

#### Aggregated Client
You can use an aggregated client to interact with multiple APIs. The aggregated client is available in the `client` package.

Import the package in your code:

```go
import niosclient "github.com/infobloxopen/infoblox-nios-go-client/client"
```

To create a new API client, you can use the `NewAPIClient` function as shown below
```go
client := niosclient.NewAPIClient()
// Now you can access the API clients using the client object, e.g.:
dnsClient := client.DNSAPI
```

#### Specific API Client
Alternatively, you can create a client for a specific API using the API package. For example, to create a client for the DNS API:

```go
//import "github.com/infobloxopen/infoblox-nios-go-client/dns"
client := dns.NewAPIClient()
```

# Configuration

The `NewAPIClient` function accepts a variadic list of `option.ClientOption` functions that can be used to configure the client.
It requires the `option` package to be imported. You can import the package using:
```go
import "github.com/infobloxopen/infoblox-nios-go-client/option"
```

### Client Name
The client name is used to identify the client in the logs. By default, the client name is set to `nios-go-client`. You can change this using the `option.WithClientName` option. For example:
```go
client := niosclient.NewAPIClient(option.WithClientName("my-client"))
```


## Example Usage


```go
package main
import (
    "context"
    "testing"
	
    "github.com/stretchr/testify/require"
    
    "github.com/infobloxopen/infoblox-nios-go-client/dns"
)

var readableAttributes = "aws_rte53_record_info,cloud_info,comment,creation_time,creator,ddns_principal,ddns_protected,disable,discovered_data,dns_name,extattrs,forbid_reclamation,ipv4addr,last_queried,ms_ad_user_data,name,reclaimable,shared_record_group,ttl,use_ttl,view,zone"

func TestCreateARecord(t *testing.T) {
    apiClient := dns.NewAPIClient()
    RecordA := dns.RecordA{
        Comment: dns.PtrString("Example comment"),
        UseTtl:  dns.PtrBool(true),
        Ttl:     dns.PtrInt64(0),
        Name:    dns.PtrString("example_record.example.com"),
        Ipv4addr: &dns.RecordAIpv4addr{
            String: dns.PtrString("127.0.0.1"),
        },
    }
    resp, _ , err := apiClient.RecordAAPI.Create(context.Background()).
        RecordA(RecordA).
        ReturnFieldsPlus(readableAttributes).
        Execute()

    if err != nil {
        t.Errorf("Error: %v", err)
    }
    require.Nil(t, err)
    require.NotNil(t, resp)
}
```

## Debugging

To enable debug logging for troubleshooting and development, use the `option.WithDebug(true)` option when creating the API client:

```go
apiClient := client.NewAPIClient(
    option.WithDebug(true),
)
```

This will print detailed request and response information to the console, helping you diagnose issues with API calls.

## Supported Operations

- Create, read, update, delete (CRUD) for NIOS objects
- Search and filter objects
- Customizable request parameters

Refer to [Infoblox WAPI documentation](https://docs.infoblox.com/display/API/WAPI+Reference) for supported object types and fields.

## Supported Object Groups

This client library supports a comprehensive set of Infoblox NIOS objects organized into the following functional groups:

### DNS Management
Manage DNS zones, records, and related configurations:

`record:a`, `record:aaaa`, `record:alias`, `record:cname`, `record:mx`, `record:txt`, `record:ptr`, `record:srv`, `record:ns`, `record:host`, `record:dname`, `record:naptr`, `record:tlsa`, `record:caa`, `record:dhcid`, `record:dnskey`, `record:ds`, `record:nsec`, `record:nsec3`, `record:nsec3param`, `record:rrsig`, `record:unknown`, `record:host_ipv4addr`, `record:host_ipv6addr`, `zone_auth`, `zone_delegated`, `zone_forward`, `zone_stub`, `zone_rp`, `zone_auth_discrepancy`, `allrecords`, `recordnamepolicy`, `dns64group`, `allnsgroup`, `ddns:principalcluster`, `ddns:principalcluster:group`, `orderedresponsepolicyzones`, `sharedrecord:a`, `sharedrecord:aaaa`, `sharedrecord:cname`, `sharedrecord:mx`, `sharedrecord:srv`, `sharedrecord:txt`, `sharedrecordgroup`, `view`, `nsgroup`, `nsgroup:delegation`, `nsgroup:forwardingmember`, `nsgroup:forwardstubserver`, `nsgroup:stubmember`

### DHCP Management
Configure and manage DHCP services:

`dhcp:statistics`, `dhcpfailover`, `dhcpoptiondefinition`, `dhcpoptionspace`, `filterfingerprint`, `filtermac`, `filternac`, `filteroption`, `filterrelayagent`, `fingerprint`, `fixedaddress`, `fixedaddresstemplate`, `ipv6dhcpoptiondefinition`, `ipv6dhcpoptionspace`, `ipv6filteroption`, `ipv6fixedaddress`, `ipv6fixedaddresstemplate`, `ipv6range`, `ipv6rangetemplate`, `ipv6sharednetwork`, `lease`, `macfilteraddress`, `range`, `rangetemplate`, `roaminghost`, `sharednetwork`, `orderedranges`

### IP Address Management (IPAM)
Manage IP addresses, networks, and related resources:

`ipv4address`, `ipv6address`, `network`, `networkcontainer`, `networktemplate`, `networkview`, `ipv6network`, `ipv6networkcontainer`, `ipv6networktemplate`, `ipam:statistics`, `vlan`, `vlanview`, `vlanrange`, `bulkhost`, `bulkhostnametemplate`, `discoverytask`, `network_discovery`, `superhost`, `superhostchild`, `hostnamerewritepolicy`

### Grid Management
Manage grid infrastructure and member configuration:

`grid`, `grid:cloudapi`, `grid:cloudapi:cloudstatistics`, `grid:cloudapi:tenant`, `grid:cloudapi:vm`, `grid:cloudapi:vmaddress`, `grid:member:cloudapi`, `grid:dhcpproperties`, `grid:servicerestart:group`, `grid:servicerestart:status`, `grid:servicerestart:request`, `mastergrid`, `grid:dashboard`, `grid:dns`, `grid:filedistribution`, `grid:threatprotection`, `grid:threatinsight`, `grid:license_pool`, `grid:x509certificate`, `captiveportal`, `grid:license_pool_container`, `grid:maxminddbinfo`, `grid:servicerestart:group:order`, `grid:servicerestart:request:changedobject`, `restartservicestatus`, `natgroup`, `member`, `member:dhcpproperties`, `member:dns`, `member:filedistribution`, `member:license`, `member:threatprotection`, `member:parentalcontrol`, `member:threatinsight`, `member:cloudsync`, `memberdfp`, `gmcgroup`, `gmcschedule`, `distributionschedule`, `upgradegroup`, `upgradeschedule`, `upgradestatus`, `license:gridwide`, `extensibleattributedef`

### Security & Authentication
Manage user accounts, permissions, and authentication:

`admingroup`, `adminrole`, `adminuser`, `permission`, `userprofile`, `authpolicy`, `approvalworkflow`, `cacertificate`, `certificate:authservice`, `radius:authservice`, `tacacsplus:authservice`, `ldap_auth_service`, `saml:authservice`, `localuser:authservice`, `snmpuser`, `ftpuser`, `networkuser`, `ad_auth_service`, `hsm:allgroups`, `hsm:entrustnshieldgroup`, `hsm:thaleslunagroup`

### DNS Traffic Control (DTC)
Manage DNS traffic control and load balancing:

`dtc`, `dtc:server`, `dtc:pool`, `dtc:topology`, `dtc:lbdn`, `dtc:monitor`, `dtc:monitor:http`, `dtc:monitor:icmp`, `dtc:monitor:pdp`, `dtc:monitor:sip`, `dtc:monitor:snmp`, `dtc:monitor:tcp`, `dtc:record:a`, `dtc:record:aaaa`, `dtc:record:cname`, `dtc:record:naptr`, `dtc:record:srv`, `dtc:certificate`, `dtc:allrecords`, `dtc:object`, `dtc:topology:rule`, `dtc:topology:label`

### Response Policy Zones (RPZ)
Manage DNS security through response policy zones:

`allrpzrecords`, `record:rpz:a`, `record:rpz:aaaa`, `record:rpz:cname`, `record:rpz:mx`, `record:rpz:naptr`, `record:rpz:ptr`, `record:rpz:srv`, `record:rpz:txt`, `record:rpz:a:ipaddress`, `record:rpz:aaaa:ipaddress`, `record:rpz:cname:ipaddress`, `record:rpz:cname:ipaddressdn`, `record:rpz:cname:clientipaddress`, `record:rpz:cname:clientipaddressdn`

### Network Discovery
Manage network discovery and device information:

`discovery`, `discovery:credentialgroup`, `discovery:device`, `discovery:devicecomponent`, `discovery:deviceinterface`, `discovery:deviceneighbor`, `discovery:devicesupportbundle`, `discovery:diagnostictask`, `discovery:gridproperties`, `discovery:memberproperties`, `discovery:sdnnetwork`, `discovery:status`, `discovery:vrf`, `vdiscoverytask`

### Cloud Integration
Manage cloud DNS and user configurations:

`awsrte53taskgroup`, `awsuser`, `azurednstaskgroup`, `azureuser`, `gcpdnstaskgroup`, `gcpuser`, `multiregions`

### Federated Realms
Manage federated realm configurations:

`federatedrealms`, `fedipamop`

### Microsoft Server Integration
Integrate with Microsoft server environments:

`msserver`, `msserver:adsites:domain`, `msserver:adsites:site`, `msserver:dhcp`, `msserver:dns`, `mssuperscope`

### Access Control Lists (ACL)
Manage network access control:

`namedacl`

### Notifications
Configure notification systems:

`notification:rest:endpoint`, `notification:rest:template`, `notification:rule`

### Parental Control
Manage parental control and content filtering:

`parentalcontrol:avp`, `parentalcontrol:blockingpolicy`, `parentalcontrol:subscriber`, `parentalcontrol:subscriberrecord`, `parentalcontrol:subscribersite`

### Regional Internet Registry (RIR)
Manage RIR configurations:

`rir`, `rir:organization`

### Smart Folders
Organize and manage folder structures:

`smartfolder:children`, `smartfolder:global`, `smartfolder:personal`

### Threat Insight
Manage threat intelligence and insights:

`threatinsight:allowlist`, `threatinsight:cloudclient`, `threatinsight:insight_allowlist`, `threatinsight:moduleset`

### Threat Protection
Configure threat protection rules and policies:

`threatprotection:grid:rule`, `threatprotection:profile`, `threatprotection:profile:rule`, `threatprotection:rule`, `threatprotection:rulecategory`, `threatprotection:ruleset`, `threatprotection:ruletemplate`, `threatprotection:statistics`

### Miscellaneous
Additional utility and administrative objects:

`allendpoints`, `bfdtemplate`, `csvimporttask`, `datacollectioncluster`, `db_objects`, `dbsnapshot`, `deleted_objects`, `dxl:endpoint`, `fileop`, `kerberoskey`, `capacityreport`, `outbound:cloudclient`, `pxgrid:endpoint`, `ruleset`, `scavengingtask`, `scheduledtask`, `search`, `syslog:endpoint`, `taxii`, `tftpfiledir`

 ## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/fooBar`)
3. Commit your changes (`git commit -am 'Add some fooBar'`)
4. Push to the branch (`git push origin feature/fooBar`)
5. Create a new Pull Request

## License

This project is owned by Infoblox and available only to limited partners and customers.

## Support

For issues, please open a GitHub issue or contact your Infoblox representative.
