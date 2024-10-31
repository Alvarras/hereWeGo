package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"
)

type (
	DNSResponse struct {
		Status   int        `json:"Status"`
		Tc       bool       `json:"TC"`
		Rd       bool       `json:"RD"`
		Ra       bool       `json:"RA"`
		Ad       bool       `json:"AD"`
		Cd       bool       `json:"CD"`
		Question []Question `json:"Question"`
		Answer   []Answer   `json:"Answer"`
	}
	Question struct {
		Name string `json:"name"`
		Type int    `json:"type"`
	}
	Answer struct {
		Name string `json:"name"`
		Type int    `json:"type"`
		TTL  int    `json:"TTL"`
		Data string `json:"data"`
	}
)

func getIPAddress(domain string) (string, error) {
	url := fmt.Sprintf("https://cloudflare-dns.com/dns-query?name=%s&type=A", domain)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("Error creating request: %v", err)
	}

	req.Header.Set("Accept", "application/dns-json")

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("Error sending request: %v", err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("Error reading response Body: %v", err)
	}

	var dnsResponse DNSResponse
	if err := json.Unmarshal(body, &dnsResponse); err != nil {
		return "", fmt.Errorf("Error decoding response body: %v", err)
	}

	if len(dnsResponse.Answer) == 0 {
		return dnsResponse.Answer[0].Data, nil
	}
	return string(body), nil

}

func TestAdress(t *testing.T) {
	type testCase struct {
		address   string
		expectErr bool
	}

	tests := []testCase{
		{"boot.dev", false},
		{"example.com", false},
		{"cloudflare.com", false},
	}

	// Additional test cases for submission
	if withSubmit {
		tests = append(tests, []testCase{
			{"iana.org", false},
		}...)
	}

	passCount := 0
	failCount := 0

	for _, test := range tests {
		output, err := getIPAddress(test.address)
		splitIP := strings.Split(output, ".")
		if err != nil && !test.expectErr {
			failCount++
			t.Errorf(`---------------------------------
URL:			%v
ExpectedErr:	%v
GotErr:			%v
Fail`, test.address, test.expectErr, err != nil)
		} else if len(splitIP) != 4 {
			failCount++
			t.Errorf(`---------------------------------
URL:			%v
Expected IP:	%v
Got:			%v
Fail`, test.address, true, output)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
URL:			%v
IP Address:		%v
Pass
`, test.address, output)
		}
	}

	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}

// withSubmit is set at compile time depending
// on which button is used to run the tests
var withSubmit = true
