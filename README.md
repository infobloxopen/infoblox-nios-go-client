# Infoblox NIOS Go Client

A Go client library for interacting with Infoblox NIOS (Network Identity Operating System) via its WAPI (Web API). This library enables Go applications to automate DNS, DHCP, IPAM, and other network management tasks using Infoblox appliances.

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

You can configure the client using environment variables for convenience:

| Variable Name         | Description                                 | Example                                |
|---------------------- |---------------------------------------------|----------------------------------------|
| NIOS_WAPI_URL         | Infoblox WAPI endpoint URL                  | `https://nios.example.com/wapi/v2.10/` |
| NIOS_USERNAME         | Username for authentication                 | `admin`                                |
| NIOS_PASSWORD         | Password for authentication                 | `password`                             | |

### Using Configuration for the API Client

Instead ofusing environment variables, you can also configure the API client directly in your code. Here's an example of how to set up the configuration in Infoblox NIOS API client :

```go
import "os"

conn := infoblox.NewConnector(
    os.Getenv("NIOS_WAPI_URL"),
    os.Getenv("NIOS_USERNAME"),
    os.Getenv("NIOS_PASSWORD"),
)
```

## Usage

```go
package main
import (
    "context"
    "testing"
	
    "github.com/stretchr/testify/require"
    
    "github.com/infobloxopen/infoblox-nios-go-client/dns"
)

var readableAttributes = "aws_rte53_record_info,cloud_info,comment,creation_time,creator,ddns_principal,ddns_protected,disable,discovered_data,dns_name,extattrs,forbid_reclamation,ipv4addr,last_queried,ms_ad_user_data,name,reclaimable,shared_record_group,ttl,use_ttl,view,zone"

func CreateARecord(t *testing.T) {
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

## Supported Operations

- Create, read, update, delete (CRUD) for NIOS objects
- Search and filter objects
- Customizable request parameters

Refer to [Infoblox WAPI documentation](https://docs.infoblox.com/display/API/WAPI+Reference) for supported object types and fields.

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
