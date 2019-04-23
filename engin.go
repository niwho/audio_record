package main // import "code.rightpaddle.cn/audio_record"
/*
#cgo LDFLAGS: -L${SRCDIR}/cpp/libs  -L${SRCDIR}/cpp -lmyagora -lstdc++ -pthread -lpthread -lrecorder -lrt -lm


#include "cpp/agora_api.h"
extern void OnUserJoinedCallBack(RecordingEnginex *record ,unsigned int uid);
void OnUserJoinedCallBack_cgo(RecordingEnginex *record ,unsigned int uid){
	OnUserJoinedCallBack(record, uid);
}

extern void OnError(RecordingEnginex *record, int error, int stat_code);
void on_error_cgo(RecordingEnginex *record, int error, int stat_code){
	OnError(record, error, stat_code);
}
*/
import "C"

import (
	"fmt"
)

func main() {
	a := 1
	fmt.Println(a)

	record := C.CreateRecordingEngine()
	//C.SetOnErrorEventHandler(record, (C.ponError)(unsafe.Pointer(C.on_error_cgo)))
	fmt.Println(record, C.SetOnErrorEventHandler)

}
