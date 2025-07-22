package main

import (
	"fmt"
	"strings"
)

//"os"

func ConvertStringSliceToHCL(input []string) string {
	var quotedStrings []string
	for _, s := range input {
		quotedStrings = append(quotedStrings, fmt.Sprintf("%q", s))
	}
	return fmt.Sprintf("[%s]", strings.Join(quotedStrings, ", "))
}

// ConvertMapToHCL serializes a map[string]any into HCL format.
func ConvertMapToHCL(data map[string]any) string {
	return convertMapToHCLWithIndent(data, 0)
}

func convertMapToHCLWithIndent(data map[string]any, indentLevel int) string {
	var keyValues []string
	keyIndent := strings.Repeat("  ", indentLevel+1)
	baseIndent := strings.Repeat("  ", indentLevel)

	// Iterate through all key-value pairs in the map
	for key, value := range data {
		// Clean up the key name (remove underscore prefix if present)
		cleanKey := key
		if strings.HasPrefix(key, "_") {
			cleanKey = key[1:]
		}

		// Format the value based on its type
		var formattedValue string

		switch v := value.(type) {
		case []map[string]any:
			// Handle slice of maps
			formattedValue = convertSliceOfMapsToHCLWithIndent(v, indentLevel+1)
		case map[string]any:
			// Handle nested map
			formattedValue = convertMapToHCLWithIndent(v, indentLevel+1)
		case []string:
			// Handle string slice
			formattedValue = ConvertStringSliceToHCL(v)
		case string:
			formattedValue = fmt.Sprintf("%q", v)
		case int, int64, float64:
			formattedValue = fmt.Sprintf("%v", v)
		case bool:
			formattedValue = fmt.Sprintf("%t", v)
		default:
			formattedValue = fmt.Sprintf("%q", fmt.Sprintf("%v", v))
		}

		keyValues = append(keyValues, fmt.Sprintf("%s%-10s = %s", keyIndent, cleanKey, formattedValue))
	}

	return fmt.Sprintf("{\n%s\n%s}", strings.Join(keyValues, "\n"), baseIndent)
}

func ConvertSliceOfMapsToHCL(data []map[string]any) string {
	return convertSliceOfMapsToHCLWithIndent(data, 0)
}

func convertSliceOfMapsToHCLWithIndent(data []map[string]any, indentLevel int) string {
	var blocks []string
	baseIndent := strings.Repeat("  ", indentLevel)
	itemIndent := strings.Repeat("  ", indentLevel+1)
	keyIndent := strings.Repeat("  ", indentLevel+2)

	for _, item := range data {
		var keyValues []string

		// Iterate through all key-value pairs in the map
		for key, value := range item {
			// Clean up the key name (remove underscore prefix if present)
			cleanKey := key
			if strings.HasPrefix(key, "_") {
				cleanKey = key[1:]
			}

			var formattedValue string

			switch v := value.(type) {
			case []map[string]any:
				// Handle nested slice of maps recursively with proper indentation
				nestedHCL := convertSliceOfMapsToHCLWithIndent(v, indentLevel+2)
				formattedValue = nestedHCL
			case map[string]any:
				// Handle nested map with proper indentation
				nestedHCL := convertMapToHCLWithIndent(v, indentLevel+2)
				formattedValue = nestedHCL
			case []string:
				// Handle string slice
				formattedValue = ConvertStringSliceToHCL(v)
			case string:
				formattedValue = fmt.Sprintf("%q", v)
			case int, int64, float64:
				formattedValue = fmt.Sprintf("%v", v)
			case bool:
				formattedValue = fmt.Sprintf("%t", v)
			default:
				formattedValue = fmt.Sprintf("%q", fmt.Sprintf("%v", v))
			}

			keyValues = append(keyValues, fmt.Sprintf("%s%-10s = %s", keyIndent, cleanKey, formattedValue))
		}

		block := fmt.Sprintf("%s{\n%s\n%s}", itemIndent, strings.Join(keyValues, "\n"), itemIndent)
		blocks = append(blocks, block)
	}

	if indentLevel == 0 {
		return fmt.Sprintf("[\n%s\n]", strings.Join(blocks, ",\n"))
	} else {
		return fmt.Sprintf("[\n%s\n%s]", strings.Join(blocks, ",\n"), baseIndent)
	}
}
func main() {
	// Test with complex nested data including map[string]any and []string
	updateForwarding := []map[string]any{
		{
			"_struct":    "addressac",
			"address":    "10.0.0.0",
			"permission": "ALLOW",
			"tags":       []string{"production", "critical", "database"},
			"metadata": map[string]any{
				"_struct":     "metadata",
				"environment": "prod",
				"region":      "us-west-2",
				"cost":        150.50,
				"active":      true,
			},
		},
		{
			"_struct": "addressac",
			"address": "10.0.0.0",
			"permission": []map[string]any{
				{
					"_struct": "addressac",
					"address": "10.0.0.0",
					"permission": []map[string]any{
						{
							"_struct":    "addressac",
							"address":    "10.0.0.0",
							"permission": "ALLOW",
						},
					},
				},
				{
					"_struct":    "addressac",
					"address":    "10.0.0.0",
					"permission": "ALLOW",
					"config": map[string]any{
						"_struct": "config",
						"timeout": 30,
						"retry":   true,
						"hosts":   []string{"host1", "host2"},
					},
				},
			},
		},
	}

	result := ConvertSliceOfMapsToHCL(updateForwarding)
	fmt.Println(result)
	// NIOS_HOST_URL := "https://172.28.83.91"
	// //NIOS_AUTH := "admin:infoblox"
	// apiClient := client.NewAPIClient(
	// 	option.WithNIOSHostUrl(NIOS_HOST_URL),
	// 	option.WithNIOSUsername("admin"),
	// 	option.WithNIOSPassword("Infoblox@123"),
	// 	option.WithDebug(true),
	// )

	// dnsClient := apiClient.DNSAPI

	// // //Create Host record
	// // reqBody := dns.RecordHost{
	// // 	Name: dns.PtrString("testshrikanthost.example.com"),
	// // 	View: dns.PtrString("default"),
	// // 	Ipv4addrs: []string{
	// // 		"10.0.0.0",
	// // 	},
	// // }

	// // resp, httpRes, err := dnsClient.RecordHostAPI.Create(context.Background()).RecordHost(reqBody).Execute()
	// // if err != nil {
	// // 	fmt.Printf("Error getting host records: %v\nFull HTTP response: %v\n", err, httpRes)
	// // }

	// _, _, _ = dnsClient.RecordAAPI.List(context.Background()).MaxResults(10).Execute()

}
