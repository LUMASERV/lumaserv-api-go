# lumaserv-api-go
Go Client for the LUMASERV API

## Usage
```go
import (
"github.com/lumaserv/lumaserv-api-go/core"
)

client := core.NewClient("YOUR_API_TOKEN")
res, _, err := client.GetServers()
```