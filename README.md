```
# mq-go2sky-hook

用于[work](https://github.com/qit-team/work) 包消息入队的hook


## Example
```go
package main

import (
    "log"
    "context"
    "time"
    
    ""github.com/qit-team/work"  // version>=0.3.12

    mqSkyHook "github.com/qit-team/mq-go2sky-hook"
	"github.com/SkyAPM/go2sky"
)


func main() {
	tracer, err := go2sky.NewTracer("127.0.0.1:11800")
	if err != nil {
		log.Fatal(err)
	}
	var ctx = context.Background()
	hook := mqSkyHook.NewHook(tracer)
	job := work.New()
	job.AddHook(hook)
	
	// do an equeue test for job
	...
}
```

```

```
