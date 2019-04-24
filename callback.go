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
	agora, exist := funcsMap.Load(record)
	if !exist {
		//return C.int(ReConsumeLater)
		return
	}
	if agora.(*defaultAgoraRecord).mOnWarning !=nil{
		agora.(*defaultAgoraRecord).mOnWarning(warn)
	}
}

//export onJoinChannelSuccess
func onJoinChannelSuccess(record *C.RecordingEnginex, channel *C.char, uid uint) {
	agora, exist := funcsMap.Load(record)
	if !exist {
		//return C.int(ReConsumeLater)
		return
	}
	if agora.(*defaultAgoraRecord).mOnJoinChannelSuccess !=nil{
		agora.(*defaultAgoraRecord).mOnJoinChannelSuccess(C.GoString(channel), uid)
	}
}

//export onLeaveChannel
func onLeaveChannel(record *C.RecordingEnginex, code int) {

	agora, exist := funcsMap.Load(record)
	if !exist {
		//return C.int(ReConsumeLater)
		return
	}
	if agora.(*defaultAgoraRecord).mOnLeaveChannel !=nil{
		agora.(*defaultAgoraRecord).mOnLeaveChannel(code)
	}
}

//export onUserJoined
func onUserJoined(record *C.RecordingEnginex, uid uint, cinfos *C.UserJoinInfos) {
	fmt.Println(record, uid)
	agora, exist := funcsMap.Load(record)
	if !exist {
		//return C.int(ReConsumeLater)
		return
	}
	//todo
	if agora.(*defaultAgoraRecord).mOnUserJoined !=nil{
		infos := CUserJoinInfosToGo(cinfos)
		agora.(*defaultAgoraRecord).mOnUserJoined(uid, infos)
	}
}

//export onUserOffline
func onUserOffline(record *C.RecordingEnginex, uid uint, reason int) {
	agora, exist := funcsMap.Load(record)
	if !exist {
		//return C.int(ReConsumeLater)
		return
	}
	if agora.(*defaultAgoraRecord).mOnUserOffline !=nil{
		agora.(*defaultAgoraRecord).mOnUserOffline(uid, reason)
	}

}

//export audioFrameReceived
func audioFrameReceived(record *C.RecordingEnginex, uid uint, cframe *C.AudioFrame) {
	agora, exist := funcsMap.Load(record)
	if !exist {
		//return C.int(ReConsumeLater)
		return
	}
	//todo
	if agora.(*defaultAgoraRecord).mAudioFrameReceived !=nil{
		frame := CAudioFrameToGo(cframe)
		agora.(*defaultAgoraRecord).mAudioFrameReceived(uid, frame)
	}

}

//export videoFrameReceived
func videoFrameReceived(record *C.RecordingEnginex, uid uint, cframe *C.VideoFrame) {
	agora, exist := funcsMap.Load(record)
	if !exist {
		//return C.int(ReConsumeLater)
		return
	}
	//todo
	if agora.(*defaultAgoraRecord).mVideoFrameReceived != nil {
		frame := CVideoFrameToGo(cframe)
		agora.(*defaultAgoraRecord).mVideoFrameReceived(uid, frame)
	}
}

//export onActiveSpeaker
func onActiveSpeaker(record *C.RecordingEnginex, uid uint) {
	agora, exist := funcsMap.Load(record)
	if !exist {
		//return C.int(ReConsumeLater)
		return
	}
	if agora.(*defaultAgoraRecord).mOnActiveSpeaker !=nil{
		agora.(*defaultAgoraRecord).mOnActiveSpeaker(uid)
	}

}

//export onAudioVolumeIndication
func onAudioVolumeIndication(record *C.RecordingEnginex, cspeakers *C.AudioVolumeInfo, speakerNum uint) {
	agora, exist := funcsMap.Load(record)
	if !exist {
		//return C.int(ReConsumeLater)
		return
	}
	if agora.(*defaultAgoraRecord).mOnAudioVolumeIndication !=nil{
		//todo AudioVolumeInfo
		speakers := CAudioVolumeInfo(cspeakers)
		agora.(*defaultAgoraRecord).mOnAudioVolumeIndication(speakers, speakerNum)
	}

}

//export onFirstRemoteVideoDecoded
func onFirstRemoteVideoDecoded(record *C.RecordingEnginex, uid uint, width int, height int, elapsed int) {
	agora, exist := funcsMap.Load(record)
	if !exist {
		//return C.int(ReConsumeLater)
		return
	}
	if agora.(*defaultAgoraRecord).mOnFirstRemoteVideoDecoded !=nil{
		agora.(*defaultAgoraRecord).mOnFirstRemoteVideoDecoded(uid, width, height, elapsed)
	}

}

//export onFirstRemoteAudioFrame
func onFirstRemoteAudioFrame(record *C.RecordingEnginex, uid uint, elapsed int) {
	agora, exist := funcsMap.Load(record)
	if !exist {
		//return C.int(ReConsumeLater)
		return
	}
	if agora.(*defaultAgoraRecord).mOnFirstRemoteAudioFrame !=nil{
		agora.(*defaultAgoraRecord).mOnFirstRemoteAudioFrame(uid, elapsed)
	}

}

//export onReceivingStreamStatusChanged
func onReceivingStreamStatusChanged(record *C.RecordingEnginex, receivingAudio int, receivingVideo int) {

	agora, exist := funcsMap.Load(record)
	if !exist {
		//return C.int(ReConsumeLater)
		return
	}
	if agora.(*defaultAgoraRecord).mOnReceivingStreamStatusChanged !=nil{
		ra := true
		if receivingAudio==0{
			ra = false
		}
		rv := true
		if receivingVideo==0{
			rv = false
		}
		agora.(*defaultAgoraRecord).mOnReceivingStreamStatusChanged(ra, rv)
	}
}

//export onConnectionLost
func onConnectionLost(record *C.RecordingEnginex) {

	agora, exist := funcsMap.Load(record)
	if !exist {
		//return C.int(ReConsumeLater)
		return
	}
	if agora.(*defaultAgoraRecord).mOnConnectionLost !=nil{
		agora.(*defaultAgoraRecord).mOnConnectionLost ()
	}
}

//export onConnectionInterrupted
func onConnectionInterrupted(record *C.RecordingEnginex) {

	agora, exist := funcsMap.Load(record)
	if !exist {
		//return C.int(ReConsumeLater)
		return
	}
	if agora.(*defaultAgoraRecord).mOnConnectionInterrupted !=nil{
		agora.(*defaultAgoraRecord).mOnConnectionInterrupted()
	}
}
