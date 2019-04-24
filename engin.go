package main // import "code.rightpaddle.cn/audio_record"
/*
#cgo LDFLAGS: -L${SRCDIR}/cpp/libs  -L${SRCDIR}/cpp -lmyagora -lstdc++ -pthread -lpthread -lrecorder -lrt -lm


#include "cpp/agora_api.h"
#include "stdlib.h"

extern void onError(RecordingEnginex *record, int error, int stat_code);
void onErrorCgo(RecordingEnginex *record, int error, int stat_code){
	onError(record, error, stat_code);
}

extern void onWarning(RecordingEnginex *record, unsigned int uid);
void onWarningCgo(RecordingEnginex *record, unsigned int uid){
	onWarning(record, uid);
}

extern void onJoinChannelSuccess(RecordingEnginex *record, const char* channelId , unsigned int uid);
void onJoinChannelSuccessCgo(RecordingEnginex *record, const char* channelId, unsigned int uid){
	onJoinChannelSuccess(record, channelId, uid);
}

extern void onLeaveChannel(RecordingEnginex*record, int code);
void onLeaveChannelCgo(RecordingEnginex*record, int code){
onLeaveChannel(record, code);
}
extern void onUserJoined(RecordingEnginex*record, unsigned int uid, UserJoinInfos* infos);
void onUserJoinedCgo(RecordingEnginex*record, unsigned int uid, UserJoinInfos* infos){
onUserJoined(record, uid, infos);
}
extern void onUserOffline(RecordingEnginex*record, unsigned int uid, int reason);
void onUserOfflineCgo(RecordingEnginex*record, unsigned int uid, int reason){
onUserOffline(record, uid, reason);
}
extern void audioFrameReceived(RecordingEnginex*record, unsigned int uid, AudioFrame * frame);
void audioFrameReceivedCgo(RecordingEnginex*record, unsigned int uid, AudioFrame * frame){
audioFrameReceived(record, uid, frame);
}
extern void videoFrameReceived(RecordingEnginex*record, unsigned int uid, VideoFrame * frame);
void videoFrameReceivedCgo(RecordingEnginex*record, unsigned int uid, VideoFrame * frame){
videoFrameReceived(record, uid, frame);
}
extern void onActiveSpeaker(RecordingEnginex*record, unsigned int uid);
void onActiveSpeakerCgo(RecordingEnginex*record, unsigned int uid){
onActiveSpeaker(record, uid);
}
extern void onAudioVolumeIndication(RecordingEnginex*record, AudioVolumeInfo* speakers, unsigned int speakerNum);
void onAudioVolumeIndicationCgo(RecordingEnginex*record, AudioVolumeInfo* speakers, unsigned int speakerNum){
onAudioVolumeIndication(record, speakers, speakerNum);
}
extern void onFirstRemoteVideoDecoded(RecordingEnginex*record, unsigned int uid, int width, int height, int elapsed);
void onFirstRemoteVideoDecodedCgo(RecordingEnginex*record, unsigned int uid, int width, int height, int elapsed){
onFirstRemoteVideoDecoded(record, uid, width, height, elapsed);
}
extern void onFirstRemoteAudioFrame(RecordingEnginex*record, unsigned int uid, int elapsed);
void onFirstRemoteAudioFrameCgo(RecordingEnginex*record, unsigned int uid, int elapsed){
onFirstRemoteAudioFrame(record, uid, elapsed);
}
extern void onReceivingStreamStatusChanged(RecordingEnginex*record, int receivingAudio, int receivingVideo);
void onReceivingStreamStatusChangedCgo(RecordingEnginex*record, int receivingAudio, int receivingVideo){
onReceivingStreamStatusChanged(record, receivingAudio, receivingVideo);
}
extern void onConnectionLost(RecordingEnginex*record);
void onConnectionLostCgo(RecordingEnginex*record){
onConnectionLost(record);
}
extern void onConnectionInterrupted(RecordingEnginex*record);
void onConnectionInterruptedCgo(RecordingEnginex*record){
onConnectionInterrupted(record);
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
	ChannelProfile int
	IsAudioOnly bool
}

type RecordingEngineProperties struct {
	StorageDir string
}

func RecordingConfigToCgo(config RecordingConfig) *C.RecordingConfig {

	cconfig := C.CreateRecordingConfig()

	C.RCSetChannelProfileType(cconfig, 1)

	return cconfig
}

func CRecordingConfigToGo(cconfig *C.RecordingConfig) RecordingConfig {

	channelProfile := int(C.RCGetChannelProfileType(cconfig))
	audioOnly  := C.int(C.RCGetIsAudioOnly(cconfig))
	isAudioOnly := true
	if audioOnly==0{
		isAudioOnly = false
	}
	config := RecordingConfig{
		ChannelProfile: channelProfile,
		IsAudioOnly: isAudioOnly,
	}

	return config
}

func CUserJoinInfosToGo(cuserJoinInfos *C.UserJoinInfos) UserJoinInfos{
	return UserJoinInfos{}
}

func CAudioFrameToGo(ccf *C.AudioFrame) AudioFrame{

	return AudioFrame{}
}

func CVideoFrameToGo(ccf *C.VideoFrame) VideoFrame{

	return VideoFrame{}
}

func CAudioVolumeInfo(cavi *C.AudioVolumeInfo) AudioVolumeInfo{

	return AudioVolumeInfo{}
}


func RecorderPropertiesToGo(rp *C.RecordingEngineProperties) RecordingEngineProperties{

	dir := C.GoString(C.REPGetStorageDir(rp))
	goRp := RecordingEngineProperties{
		StorageDir :dir,
	}
	return goRp
}

type IAgoraRecord interface {
	CreateChannel(appId string, channelKey string, name string,uid uint, config RecordingConfig) int
	SetVideoMixLayout() int
	LeaveChannel() int
	StoppedOnError() int
	Release() int
	Stopped() int
	UpdateMixModeSetting(width int, height int, isVideoMix bool)
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

	C.SetOnError(ar.RecordingEnginex, (C.pOnError)(unsafe.Pointer(C.onErrorCgo)))
	C.SetOnWarning(ar.RecordingEnginex,(C.pOnWarning)(unsafe.Pointer(C.onWarningCgo)))
	C.SetOnJoinChannelSuccess(ar.RecordingEnginex,(C.pOnJoinChannelSuccess)(unsafe.Pointer(C.onJoinChannelSuccessCgo)))
	C.SetOnLeaveChannel(ar.RecordingEnginex,(C.pOnLeaveChannel)(unsafe.Pointer(C.onLeaveChannelCgo)))
	C.SetOnUserJoined(ar.RecordingEnginex,(C.pOnUserJoined)(unsafe.Pointer(C.onUserJoinedCgo)))
	C.SetOnUserOffline(ar.RecordingEnginex,(C.pOnUserOffline)(unsafe.Pointer(C.onUserOfflineCgo)))
	C.SetAudioFrameReceived(ar.RecordingEnginex,(C.pAudioFrameReceived)(unsafe.Pointer(C.audioFrameReceivedCgo)))
	C.SetVideoFrameReceived(ar.RecordingEnginex,(C.pVideoFrameReceived)(unsafe.Pointer(C.videoFrameReceivedCgo)))
	C.SetOnActiveSpeaker(ar.RecordingEnginex,(C.pOnActiveSpeaker)(unsafe.Pointer(C.onActiveSpeakerCgo)))
	C.SetOnAudioVolumeIndication(ar.RecordingEnginex,(C.pOnAudioVolumeIndication)(unsafe.Pointer(C.onAudioVolumeIndicationCgo)))
	C.SetOnFirstRemoteVideoDecoded(ar.RecordingEnginex,(C.pOnFirstRemoteVideoDecoded)(unsafe.Pointer(C.onFirstRemoteVideoDecodedCgo)))
	C.SetOnFirstRemoteAudioFrame(ar.RecordingEnginex,(C.pOnFirstRemoteAudioFrame)(unsafe.Pointer(C.onFirstRemoteAudioFrameCgo)))
	C.SetOnReceivingStreamStatusChanged(ar.RecordingEnginex,(C.pOnReceivingStreamStatusChanged)(unsafe.Pointer(C.onReceivingStreamStatusChangedCgo)))
	C.SetOnConnectionLost(ar.RecordingEnginex,(C.pOnConnectionLost)(unsafe.Pointer(C.onConnectionLostCgo)))
	C.SetOnConnectionInterrupted(ar.RecordingEnginex,(C.pOnConnectionInterrupted)(unsafe.Pointer(C.onConnectionInterruptedCgo)))

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
	C.FreeRecordingConfig(ccf) // 应当在对象释放时删除？？
	return int(rt)
}

func (ar *defaultAgoraRecord) SetVideoMixLayout() int {

	rt := C.SetVideoMixLayout(ar.RecordingEnginex)
	return int(rt)
}

func (ar *defaultAgoraRecord) LeaveChannel() int {
	rt := C.LeaveChannel(ar.RecordingEnginex)
	return int(rt)
}
func (ar *defaultAgoraRecord) StoppedOnError() int {
	rt := C.StoppedOnError(ar.RecordingEnginex)
	return int(rt)
}
func (ar *defaultAgoraRecord) Release() int {
	rt := C.Release(ar.RecordingEnginex)
	return int(rt)
}
func (ar *defaultAgoraRecord) Stopped() int {
	rt := C.Stopped(ar.RecordingEnginex)
	return int(rt)
}
func (ar *defaultAgoraRecord) UpdateMixModeSetting(width int, height int, isVideoMix bool) {
	videoMix := 0
	if isVideoMix{
		videoMix = 1
	}
	C.UpdateMixModeSetting(ar.RecordingEnginex, C.int(width), C.int(height), C.int(videoMix))
	return
}
func (ar *defaultAgoraRecord) GetRecorderProperties() RecordingEngineProperties {
	rp := C.GetRecorderProperties(ar.RecordingEnginex)
	goRp := RecorderPropertiesToGo(rp)
	return goRp
}
func (ar *defaultAgoraRecord) UpdateStorageDir(dir string) {
	d := C.CString(dir)
	C.UpdateStorageDir(ar.RecordingEnginex, d)
	C.free(unsafe.Pointer(d))
	return
}
func (ar *defaultAgoraRecord) UpdateLayoutSetting(layoutMode int, maxVertPreLayoutUid int) {
	C.UpdateLayoutSetting(ar.RecordingEnginex, C.int(layoutMode), C.int(maxVertPreLayoutUid))
	return
}
func (ar *defaultAgoraRecord) StartService() int {
	rt := C.StartService(ar.RecordingEnginex)
	return int(rt)
}
func (ar *defaultAgoraRecord) StopService() int {
	rt := C.StopService(ar.RecordingEnginex)
	return int(rt)
}

func (ar *defaultAgoraRecord) SetVideoMixingLayout() int {
	rt := C.SetVideoMixingLayout(ar.RecordingEnginex)
	return int(rt)
}
func (ar *defaultAgoraRecord) GetConfigInfo() RecordingConfig {
	//todo
	cconfig := C.GetConfigInfo(ar.RecordingEnginex)
	return CRecordingConfigToGo(cconfig)
}
func (ar *defaultAgoraRecord) SetUserBackground(uid uint, imagePath string) int {
	img := C.CString(imagePath)
	rt := C.SetVideoMixLayout(ar.RecordingEnginex, uid, img)
	C.free(unsafe.Pointer(img))
	return int(rt)
}
func (ar *defaultAgoraRecord) SetLogLevel(level int) int {
	rt := C.SetLogLevel(ar.RecordingEnginex, C.int(level))
	return int(rt)
}
func (ar *defaultAgoraRecord) SetMediaKeepTime(timeMs uint) {
	C.SetMediaKeepTime(ar.RecordingEnginex, C.uint(timeMs))
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
