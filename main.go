package pkicsr

import (
	"crypto/x509"
	"encoding/base64"
	"strings"
)

// Data struct is the main struct
type Data struct {
	CSR                string     `json:"csr,omitempty"`
	Version            int        `json:"version,omitempty"`
	CommonName         string     `json:"common_name,omitempty"`
	EmailAddresses     []string   `json:"email_addresses,omitempty"`
	DNSNames           []string   `json:"dns_names,omitempty"`
	SignatureAlgorithm string     `json:"signature_algorithm,omitempty"`
	PublicKeyAlgorithm string     `json:"public_key_algorithm,omitempty"`
	Country            []string   `json:"country,omitempty"`
	Locality           []string   `json:"locality,omitempty"`
	Organization       []string   `json:"organization,omitempty"`
	OrganizationalUnit []string   `json:"organizational_unit,omitempty"`
	PostalCode         []string   `json:"postal_code,omitempty"`
	Province           []string   `json:"province,omitempty"`
	SerialNumber       string     `json:"serial_number,omitempty"`
	StreetAddress      []string   `json:"street_address,omitempty"`
	Subject            string     `json:"subject,omitempty"`
	Controls           []*control `json:"control,omitempty"`
	Error              bool       `json:"error,omitempty"`
	ErrorMessage       string     `json:"errormessage,omitempty"`
}

type control struct {
	Message  string `json:"message,omitempty"`
	Blocking bool   `json:"blocking,omitempty"`
}

// Get function, main function of this module.
func Get(csrPOST string) *Data {
	response := new(Data)
	response.CSR = csrPOST

	// Change CSR format
	csrPOST = strings.Replace(csrPOST, "-----BEGIN CERTIFICATE REQUEST-----", "", -1)
	csrPOST = strings.Replace(csrPOST, "-----END CERTIFICATE REQUEST-----", "", -1)
	csrPOST = strings.Replace(csrPOST, "\n", "", -1)
	csrPOST = strings.Replace(csrPOST, "\t", "", -1)
	csrPOST = strings.Replace(csrPOST, " ", "", -1)
	csrBytes := fromBase64(csrPOST)
	csr, err := x509.ParseCertificateRequest(csrBytes)
	if err != nil {
		response.Error = true
		response.ErrorMessage = err.Error()
		return response
	}
/*
	if caadata.DNSSEC == true {
		caacontrol := new(control)
		caacontrol.Message = "DNSSEC found on domain, no DNS errors may occur."
		caacontrol.Blocking = false
		caadata.Controls = append(caadata.Controls, caacontrol)
	}
*/

	response.Version = csr.Version
	response.CommonName = csr.Subject.CommonName
	response.EmailAddresses = csr.EmailAddresses
	response.DNSNames = csr.DNSNames
	response.SignatureAlgorithm = csr.SignatureAlgorithm.String()
	response.PublicKeyAlgorithm = csr.PublicKeyAlgorithm.String()
	response.Country = csr.Subject.Country
	response.Locality = csr.Subject.Locality
	response.Organization = csr.Subject.Organization
	response.OrganizationalUnit = csr.Subject.OrganizationalUnit
	response.PostalCode = csr.Subject.PostalCode
	response.Province = csr.Subject.Province
	response.SerialNumber = csr.Subject.SerialNumber
	response.StreetAddress = csr.Subject.StreetAddress
	response.Subject = csr.Subject.String()

	// response.Parsed = csr
	return response
}

func fromBase64(in string) []byte {
	out := make([]byte, base64.StdEncoding.DecodedLen(len(in)))
	n, err := base64.StdEncoding.Decode(out, []byte(in))
	if err != nil {
		panic("failed to base64 decode")
	}
	return out[:n]
}
