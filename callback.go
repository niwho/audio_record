package main

/*
#cgo LDFLAGS: -L${SRCDIR}/cpp/libs  -L${SRCDIR}/cpp -lrecorder -lmyagora -lstdc++

#include "cpp/agora_api.h"

*/
import "C"

import (
	"fmt"
)

//export OnError
func OnError(record *C.RecordingEnginex, error1, statCode C.int) {
	fmt.Println(record, error1, statCode)
}

//export OnUserJoinedCallBack
func OnUserJoinedCallBack(record *C.RecordingEnginex, uid C.uint) {
	fmt.Println(record, uid)
}

//export OnUserJoinedCallBack
func OnUserJoinedCallBack123(record *C.RecordingEnginex, uid C.uint) {
	fmt.Println(record, uid)
}
