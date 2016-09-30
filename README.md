# google-mrf-interface

MRF and Media Measurement Schema API

In order to see the available endpoints see the https://github.com/essence-tech/google-mrf-interface/blob/master/google_mrf.proto file. A
Python and Go client are provided for communication with the service.

## Clients

### Python
Install from pip with
```
pip install git+git://github.com/essence-tech/google-mrf-interace.git@master
```

```python
import googlemrf


client = googlemrf.get_client('ess-lon-cmo-003:4041')

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
    "github.com/essence-tech/google-mrf-inteface"
)

c = client.New("ess-lon-cmo-003:4041")
p, err := c.ValidateProduct("TV")
if err != nil {
    panic(err)
}
...
```
