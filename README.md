# google-mrf-interface

MRF and Media Measurement Schema API

In order to see the available endpoints see the https://github.com/essence-tech/google-mrf-interface/blob/master/google_mrf.proto file. A
Python and Go client are provided for communication with the service.

## Clients

### Python
Install from pip with
```
pip install git+git://github.com/essence-tech/google-mrf-interface.git@master
```

```python
import googlemrf


client = googlemrf.get_client('ess-lon-cmo-003:4041', 'google-mrf.pem')

# Single item.
p = client.ValidateProduct(googlemrf.Query(name='TV'))

# List all MRFs.
mrf = client.MRFs(googlemrf.Empty())
for mrf in mrfs:
    print mrf

```

### Go

```go
package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"

	client "github.com/essence-tech/google-mrf-interface"
)

func main() {
	certificate, _ := tls.LoadX509KeyPair(
		"google-mrf.crt",
		"google-mrf.key",
	)
	certPool := x509.NewCertPool()
	bs, _ := ioutil.ReadFile("essence-systems.crt")
	certPool.AppendCertsFromPEM(bs)
	tc := &tls.Config{
		ServerName:   "google-mrf",
		Certificates: []tls.Certificate{certificate},
		RootCAs:      certPool,
	}

	c := client.New("127.0.0.1:4041", tc)
	list, err := c.MRFs()
	if err != nil {
		panic(err)
	}

	for _, m := range list {
		fmt.Println(m.PrimaryRegion)
	}
}
```
