# audio_record
> 声网录制sdk的golang封装
1. 依赖linux环境(不支持macos/win)
2. 编译
> cd audio_record/cpp&make

# 示例
```golang
package main // import "github.com/niwho/record_demo"

import "fmt"
import "github.com/niwho/audio_record"

type MyRecord struct {
    audio_record.IAgoraRecord
}

func (mr *MyRecord) OnError(err int, statCode int) {

}

func (mr *MyRecord) OnWarn(warn int) {

}

func main() {
    record := &MyRecord{
        IAgoraRecord: audio_record.NewAgoraRecord(),
    }
    record.SetOnError(record.OnError)
    record.SetOnWarning(record.OnWarn)

    fmt.Println("record:", record.IAgoraRecord)
}

```
