package main // import "code.rightpaddle.cn/audio_record"
/*
#cgo LDFLAGS: -L${SRCDIR}/cpp/libs  -L${SRCDIR}/cpp -lmyagora -lstdc++ -pthread -lpthread -lrecorder -lrt -lm


#include "cpp/agora_api.h"
#include "stdlib.h"

extern void onError(RecordingEnginex *record, int error, int stat_code);
void on_error_cgo(RecordingEnginex *record, int error, int stat_code){
	onError(record, error, stat_code);
}

extern void onWarning(RecordingEnginex *record, unsigned int uid);
void on_warning_cgo(RecordingEnginex *record, unsigned int uid){
	onWarning(record, uid);
}

extern onJoinChannelSuccess(RecordingEnginex *record, const char* channelId , unsigned int uid);
void on_join_channel_success_cgo(RecordingEnginex *record, const char* channelId, unsigned int uid){
	onJoinChannelSuccess(record, channelId, uid);
}
*/
import "C"

import (
	"fmt"
	"unsafe"
)

type UserJoinInfos struct {
	StorageDir string
}

type AudioFrame struct {
}

type VideoFrame struct {
}

type AudioVolumeInfo struct {
}

type RecordingConfig struct {
}

type RecordingEngineProperties struct {
}

func RecordingConfigToCgo(rf RecordingConfig) *C.RecordingConfig {

	return C.CreateRecordingConfig()
}

type IAgoraRecord interface {
	CreateChannel(appId string, channelKey string, name string,uid uint, config RecordingConfig) int
	SetVideoMixLayout() int
	LeaveChannel() int
	StoppedOnError() int
	Release() int
	Stopped() int
	UpdateMixModeSetting(width int, height int, isVideoMix bool) int
	GetRecorderProperties() RecordingEngineProperties
	UpdateStorageDir(dir string)
	UpdateLayoutSetting(layoutMode int, maxVertPreLayoutUid int)
	StartService() int
	StopService() int
	SetVideoMixingLayout() int
	GetConfigInfo() RecordingConfig
	SetUserBackground(uid uint, imagePath string) int
	SetLogLevel(level int) int
	SetMediaKeepTime(timeMs uint)

	SetOnError(func(int, int))
	SetOnWarning(func(warn int))
	SetOnJoinChannelSuccess(func(string, uint))
	SetOnLeaveChannel(func(int))
	SetOnUserJoined(func(uint, UserJoinInfos))
	SetOnUserOffline(func(uid uint, reason int))
	SetAudioFrameReceived(func(uint, AudioFrame))
	SetVideoFrameReceived(func(uint, VideoFrame))
	SetOnActiveSpeaker(func(uint))
	SetOnAudioVolumeIndication(func(AudioVolumeInfo, uint))
	SetOnFirstRemoteVideoDecoded(func(uid uint, width int, height int, elapsed int))
	SetOnFirstRemoteAudioFrame(func(uid uint, elapsed int))
	SetOnReceivingStreamStatusChanged(func(receivingAudio bool, receivingVideo bool))
	SetOnConnectionLost(func())
	SetOnConnectionInterrupted(func())
}

type defaultAgoraRecord struct {
	RecordingEnginex *C.RecordingEnginex

	mOnError                        func(int, int)
	mOnWarning                      func(warn int)
	mOnJoinChannelSuccess           func(string, uint)
	mOnLeaveChannel                 func(int)
	mOnUserJoined                   func(uint, UserJoinInfos)
	mOnUserOffline                  func(uid uint, reason int)
	mAudioFrameReceived             func(uint, AudioFrame)
	mVideoFrameReceived             func(uint, VideoFrame)
	mOnActiveSpeaker                func(uint)
	mOnAudioVolumeIndication        func(AudioVolumeInfo, uint)
	mOnFirstRemoteVideoDecoded      func(uid uint, width int, height int, elapsed int)
	mOnFirstRemoteAudioFrame        func(uid uint, elapsed int)
	mOnReceivingStreamStatusChanged func(receivingAudio bool, receivingVideo bool)
	mOnConnectionLost               func()
	mOnConnectionInterrupted        func()
}

func NewAgoraRecord() IAgoraRecord {
	ar := &defaultAgoraRecord{}
	ar.RecordingEnginex = C.CreateRecordingEngine()
	C.SetOnError(ar.RecordingEnginex, (C.pOnError)(unsafe.Pointer(C.on_error_cgo)))

	funcsMap.Store(ar.RecordingEnginex, ar)

	return ar
}

func (ar *defaultAgoraRecord) CreateChannel(appId string, channelKey string, name string, uid uint,config RecordingConfig) int {
	ad := C.CString(appId)
	chk := C.CString(channelKey)
	nm := C.CString(name)
	ccf := RecordingConfigToCgo(config)
	rt := C.CreateChannel(ar.RecordingEnginex, ad, chk, nm, C.uint(uid), ccf)

	C.free(unsafe.Pointer(ad))
	C.free(unsafe.Pointer(chk))
	C.free(unsafe.Pointer(nm))
	return int(rt)
}

func (ar *defaultAgoraRecord) SetVideoMixLayout() int {
	return 0
}

func (ar *defaultAgoraRecord) LeaveChannel() int {
	return 0
}
func (ar *defaultAgoraRecord) StoppedOnError() int {
	return 0
}
func (ar *defaultAgoraRecord) Release() int {
	return 0
}
func (ar *defaultAgoraRecord) Stopped() int {
	return 0
}
func (ar *defaultAgoraRecord) UpdateMixModeSetting(width int, height int, isVideoMix bool) int {
	return 0
}
func (ar *defaultAgoraRecord) GetRecorderProperties() RecordingEngineProperties {
	return RecordingEngineProperties{}
}
func (ar *defaultAgoraRecord) UpdateStorageDir(dir string) {
	return
}
func (ar *defaultAgoraRecord) UpdateLayoutSetting(layoutMode int, maxVertPreLayoutUid int) {
	return
}
func (ar *defaultAgoraRecord) StartService() int {
	return 0
}
func (ar *defaultAgoraRecord) StopService() int {
	return 0
}
func (ar *defaultAgoraRecord) SetVideoMixingLayout() int {
	return 0
}
func (ar *defaultAgoraRecord) GetConfigInfo() RecordingConfig {
	return RecordingConfig{}
}
func (ar *defaultAgoraRecord) SetUserBackground(uid uint, imagePath string) int {
	return 0
}
func (ar *defaultAgoraRecord) SetLogLevel(level int) int {
	return 0
}
func (ar *defaultAgoraRecord) SetMediaKeepTime(timeMs uint) {
	return
}

func (ar *defaultAgoraRecord) SetOnError(onError func(int, int)) {
	ar.mOnError = onError
}
func (ar *defaultAgoraRecord) SetOnWarning(rnWarning func(warn int)) {}
func (ar *defaultAgoraRecord) SetOnJoinChannelSuccess(onJoinChannelSuccess func(string, uint)) {
	ar.mOnJoinChannelSuccess = onJoinChannelSuccess
}
func (ar *defaultAgoraRecord) SetOnLeaveChannel(onLeaveChannel func(int)) {
	ar.mOnLeaveChannel = onLeaveChannel
}
func (ar *defaultAgoraRecord) SetOnUserJoined(onUserJoined func(uint, UserJoinInfos)) {
	ar.mOnUserJoined = onUserJoined
}
func (ar *defaultAgoraRecord) SetOnUserOffline(onUserOffline func(uid uint, reason int)) {
	ar.mOnUserOffline = onUserOffline
}
func (ar *defaultAgoraRecord) SetAudioFrameReceived(audioFrameReceived func(uint, AudioFrame)) {
	ar.mAudioFrameReceived = audioFrameReceived
}
func (ar *defaultAgoraRecord) SetVideoFrameReceived(videoFrameReceived func(uint, VideoFrame)) {
	ar.mVideoFrameReceived = videoFrameReceived
}
func (ar *defaultAgoraRecord) SetOnActiveSpeaker(onActiveSpeaker func(uint)) {
	ar.mOnActiveSpeaker = onActiveSpeaker
}
func (ar *defaultAgoraRecord) SetOnAudioVolumeIndication(onAudioVolumeIndication func(AudioVolumeInfo, uint)) {
	ar.mOnAudioVolumeIndication = onAudioVolumeIndication
}
func (ar *defaultAgoraRecord) SetOnFirstRemoteVideoDecoded(onFirstRemoteVideoDecoded func(uid uint, width int, height int, elapsed int)) {
	ar.mOnFirstRemoteVideoDecoded = onFirstRemoteVideoDecoded
}
func (ar *defaultAgoraRecord) SetOnFirstRemoteAudioFrame(onFirstRemoteAudioFrame func(uid uint, elapsed int)) {
	ar.mOnFirstRemoteAudioFrame = onFirstRemoteAudioFrame
}
func (ar *defaultAgoraRecord) SetOnReceivingStreamStatusChanged(onReceivingStreamStatusChanged func(receivingAudio bool, receivingVideo bool)) {
	ar.mOnReceivingStreamStatusChanged = onReceivingStreamStatusChanged
}
func (ar *defaultAgoraRecord) SetOnConnectionLost(onConnectionLost func()) {
	ar.mOnConnectionLost = onConnectionLost
}
func (ar *defaultAgoraRecord) SetOnConnectionInterrupted(onConnectionInterrupted func()) {
	ar.mOnConnectionInterrupted = onConnectionInterrupted
}

func main() {
	a := 1
	fmt.Println(a)

	record := C.CreateRecordingEngine()
	//C.SetOnErrorEventHandler(record, (C.ponError)(unsafe.Pointer(C.on_error_cgo)))
	fmt.Println(record, C.SetOnError)

}
