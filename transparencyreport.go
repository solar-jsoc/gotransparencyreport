package transparencyreport

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	baseURL = "https://transparencyreport.google.com/transparencyreport/api/v3/httpsreport/ct"
)

var (
	HTTPClient = http.DefaultClient
)

type Cert struct {
	Subject      string
	Issuer       string
	SerialNumber string
	NotBefore    time.Time
	NotAfter     time.Time
	DNSNames     []string
}

func (c *Cert) UnmarshalJSON(data []byte) error {
	var notBefore, notAfter int64
	cert := []interface{}{&c.SerialNumber, &c.Subject, &c.Issuer, &notBefore, &notAfter, nil, nil, &c.DNSNames}
	proxy := []interface{}{nil, &cert, nil}

	err := json.Unmarshal(data, &proxy)
	if err != nil {
		return err
	}

	c.NotBefore = time.Unix(notBefore/1000, 0)
	c.NotAfter = time.Unix(notAfter/1000, 0)
	if len(strings.SplitN(c.SerialNumber, ":", 2)[0]) == 1 {
		c.SerialNumber = "0" + c.SerialNumber
	}

	return nil
}

type certInfo struct {
	Hash string
}

func (c *certInfo) UnmarshalJSON(data []byte) error {
	proxy := []interface{}{nil, nil, nil, nil, nil, &c.Hash, nil, nil, nil}
	return json.Unmarshal(data, &proxy)
}

type certsResponse struct {
	Certs      []certInfo
	NextPageID string
}

func (r *certsResponse) UnmarshalJSON(data []byte) error {
	pageInfo := []interface{}{nil, &r.NextPageID, nil, nil, nil}
	proxy := []interface{}{nil, &r.Certs, nil, &pageInfo}
	return json.Unmarshal(data, &proxy)
}

func get(path string, values url.Values) ([]byte, error) {
	response, err := HTTPClient.Get(baseURL + path + "?" + values.Encode())
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP status: %d", response.StatusCode)
	}

	return ioutil.ReadAll(response.Body)
}

// Search certs for domain
func Search(domain string, includeExpired, includeSubdomains bool) ([]Cert, error) {
	if domain == "" {
		return nil, fmt.Errorf("Domain is empty")
	}
	body, err := get("/certsearch", url.Values{
		"domain":             {domain},
		"include_expired":    {fmt.Sprintf("%t", includeExpired)},
		"include_subdomains": {fmt.Sprintf("%t", includeSubdomains)},
	})
	if err != nil {
		return nil, fmt.Errorf("Get domains error: %w", err)
	}

	response := certsResponse{}
	err = unmarshalJSON(body, &response)
	if err != nil {
		return nil, err
	}

	certificates := []Cert{}
	for {
		for _, crt := range response.Certs {
			updCrt, err := getCert(crt.Hash)
			if err != nil {
				return nil, fmt.Errorf("Parse certificate by hash error: %w", err)
			}
			certificates = append(certificates, updCrt)
		}

		if response.NextPageID == "" {
			break
		}

		body, err = get("/certsearch/page", url.Values{"p": {response.NextPageID}})
		if err != nil {
			return nil, err
		}

		response = certsResponse{}
		err = unmarshalJSON(body, &response)
		if err != nil {
			return nil, err
		}
	}
	return certificates, nil
}

func getCert(hash string) (Cert, error) {
	cert := Cert{}
	body, err := get("/certbyhash", url.Values{"hash": {hash}})
	if err != nil {
		return cert, err
	}

	err = unmarshalJSON(body, &cert)
	return cert, err
}

func unmarshalJSON(data []byte, v interface{}) error {
	data = bytes.TrimSpace(bytes.TrimPrefix(data, []byte(")]}'")))
	if !bytes.HasPrefix(data, []byte("[[")) {
		return fmt.Errorf("Unknown prefix")
	}

	return json.Unmarshal(data[1:len(data)-1], v)
}
