## Redis Go
Provides easy to use API to operate Redis db.

### How to use it?
`go get github.com/lhdhtrc/redis-go`

```go
package main

import (
	redis "github.com/lhdhtrc/redis-go/pkg"
	"go.uber.org/zap"
)

func main() {
	client, err := redis.New(&redis.Config{})
}
```

### Finally
- If you feel good, click on star.
- If you have a good suggestion, please ask the issue.