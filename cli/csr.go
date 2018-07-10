package main

import (
	"encoding/json"
	"fmt"

	"github.com/binaryfigments/csr"
)

func main() {
	test := `-----BEGIN CERTIFICATE REQUEST-----
	MIICuzCCAaMCAQAwdjELMAkGA1UEBhMCTkwxHDAaBgNVBAMME3d3dy5kb2RneWRv
	bnV0cy5jb20xEDAOBgNVBAcMB1V0cmVjaHQxGDAWBgNVBAoMD0RvZGd5IERvbnV0
	cyBCVjEPMA0GA1UECAwGVXJlY2h0MQwwCgYDVQQLDANJQ1QwggEiMA0GCSqGSIb3
	DQEBAQUAA4IBDwAwggEKAoIBAQDuQSaCpyxyTMfl5h0MVA4AjhPMVw+3+ia0LSYQ
	UQ25UPY+4A+dodWAz+oScdEM/dThfsFDmXcoeILmGdZZcSusGPpEw7D3E+svqvaG
	0m1vcjQjHoyYkadE6AWqbSAAhYzZ565oCERonzHpkt8TzXBV5dZVVF7BPcY+WTx/
	yCDi+mXQnMd9ngJ0hdmwL9rctTFygnD4eWDGvrbuljLHidp2qyKfKt2+DVKtywxd
	p3C/c5YQFPtoyoWnahZIManEd3DRljvyQdsxO/9jcuzVmWjpBthpdPlK3ev+RUVd
	9ZliMtjteLhWMzjskytfwQ1sackl+RDj9Q/GJpJ7UFXUgUajAgMBAAGgADANBgkq
	hkiG9w0BAQsFAAOCAQEAvXuZSRiPmmFnRI6/nzIt5bFe8Igp7Jmt0gi4C60sa630
	y/oGR0W3ILhFlgjRJkDRQlTyht4mpk7SuXTqiq/SXxFxe8nvZm1+n7x0K3uhcw9W
	/D+eis2RMao8DgJVey/uXbJTzKxlHoTAwqE/jfKOzrNih6nDOxCv5kM8F08IpAD2
	kloyXC4Q6FBUCmY6nKOULk5x6FpPIZoIytN0yZrTvfRithbAPvKjF2w6sTyVob6Y
	zLIuAtMYWI3ZdFqCSZ0tXLOfwjol8+cx11dCZFGcZKBT5V2twsOv9wOGPPxlg+fB
	BhGlPMFNvcuTydH+58O5MI2MG+woyWV2EQd84kZ3yw==
	-----END CERTIFICATE REQUEST-----`

	data := pkicsr.Get(test)
	json, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s\n", json)
}
