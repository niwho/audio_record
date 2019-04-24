package main

/*
#cgo LDFLAGS: -L${SRCDIR}/cpp/libs  -L${SRCDIR}/cpp -lrecorder -lmyagora -lstdc++

#include "cpp/agora_api.h"

*/
import "C"

import (
	"fmt"
	"sync"
)

var funcsMap sync.Map

//export onError
func onError(record *C.RecordingEnginex, err, statCode int) {
	agora, exist := funcsMap.Load(record)
	if !exist {
		//return C.int(ReConsumeLater)
		return
	}
	if agora.(*defaultAgoraRecord).mOnError !=nil{
		agora.(*defaultAgoraRecord).mOnError(err, statCode)
	}
	fmt.Println(record, err, statCode)
}

//export onWarning
func onWarning(record *C.RecordingEnginex, warn int) {
	fmt.Println(warn)
}

//export onJoinChannelSuccess
func onJoinChannelSuccess(record *C.RecordingEnginex, channel *C.char, uid uint) {

}

//export onLeaveChannel
func onLeaveChannel(record *C.RecordingEnginex, code int) {

}

//export onUserJoined
func onUserJoined(record *C.RecordingEnginex, uid uint, infos *C.UserJoinInfos) {
	fmt.Println(record, uid)
}

//export onUserOffline
func onUserOffline(record *C.RecordingEnginex, uid uint, reason int) {

}

//export audioFrameReceived
func audioFrameReceived(record *C.RecordingEnginex, uid uint, frame *C.AudioFrame) {

}

//export videoFrameReceived
func videoFrameReceived(record *C.RecordingEnginex, uid uint, frame *C.VideoFrame) {

}

//export onActiveSpeaker
func onActiveSpeaker(record *C.RecordingEnginex, uid uint) {

}

//export onAudioVolumeIndication
func onAudioVolumeIndication(record *C.RecordingEnginex, speakers *C.AudioVolumeInfo, speakerNum uint) {

}

//export onFirstRemoteVideoDecoded
func onFirstRemoteVideoDecoded(record *C.RecordingEnginex, uid uint, width int, height int, elapsed int) {

}

//export onFirstRemoteAudioFrame
func onFirstRemoteAudioFrame(record *C.RecordingEnginex, uid uint, elapsed int) {

}

//export onReceivingStreamStatusChanged
func onReceivingStreamStatusChanged(record *C.RecordingEnginex, receivingAudio int, receivingVideo int) {

}

//export onConnectionLost
func onConnectionLost(record *C.RecordingEnginex) {

}

//export onConnectionInterrupted
func onConnectionInterrupted(record *C.RecordingEnginex) {

}
