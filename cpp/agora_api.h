#ifndef __C_AGORA_API_H__
#define __C_AGORA_API_H__

#include "common.h"
//#include "record.h"
#include "user_join_info.h"
#include "recording_config.h"
#include "RecordingEngineProperties.h"
#include "miscellaneous.h"
#ifdef __cplusplus
extern "C" {
#endif

typedef struct RecordingEnginex RecordingEnginex;
//typedef struct UserJoinInfos UserJoinInfos; //user_join_info.h
//typedef struct AudioFrame AudioFrame;//miscellaneous.h
//typedef struct VideoFrame VideoFrame;//miscellaneous.h
//typedef struct AudioVolumeInfo AudioVolumeInfo;//miscellaneous.h
//typedef struct RecordingConfig RecordingConfig; //recording_config.h
//typedef struct RecordingEngineProperties RecordingEngineProperties; //RecordingEngineProperties.h

// UserJoinInfos

// AudioFrame


// VideoFrame

// AudioVolumeInfo

// RecordingConfig


// RecordingEngineProperties


typedef void (*pOnError)(RecordingEnginex*record, int error, int stat_code);
typedef void (*pOnWarning)(RecordingEnginex*record, int warn);
typedef void (*pOnJoinChannelSuccess)(RecordingEnginex*record, const char* channelId, unsigned int uid);
typedef void (*pOnLeaveChannel)(RecordingEnginex*record, int code);
typedef void (*pOnUserJoined)(RecordingEnginex*record, unsigned int uid, UserJoinInfos* infos);
typedef void (*pOnUserOffline)(RecordingEnginex*record, unsigned int uid, int reason);
typedef void (*pAudioFrameReceived)(RecordingEnginex*record, unsigned int uid, AudioFrame * frame);
typedef void (*pVideoFrameReceived)(RecordingEnginex*record, unsigned int uid, VideoFrame * frame);
typedef void (*pOnActiveSpeaker)(RecordingEnginex*record, unsigned int uid);
typedef void (*pOnAudioVolumeIndication)(RecordingEnginex*record, AudioVolumeInfo* speakers, unsigned int speakerNum);
typedef void (*pOnFirstRemoteVideoDecoded)(RecordingEnginex*record, unsigned int uid, int width, int height, int elapsed);
typedef void (*pOnFirstRemoteAudioFrame)(RecordingEnginex*record, unsigned int uid, int elapsed);
typedef void (*pOnReceivingStreamStatusChanged)(RecordingEnginex*record, int receivingAudio, int receivingVideo);
typedef void (*pOnConnectionLost)(RecordingEnginex*record);
typedef void (*pOnConnectionInterrupted)(RecordingEnginex*record);

AGORA_API RecordingEnginex* CreateRecordingEngine();
AGORA_API int CreateChannel(RecordingEnginex*record,const char * appid, const char* channelKey, const char* name, unsigned int uid, RecordingConfig* config);
AGORA_API int SetVideoMixLayout(RecordingEnginex*record);
AGORA_API int LeaveChannel(RecordingEnginex*record);
AGORA_API int StoppedOnError(RecordingEnginex*record);
AGORA_API int Release(RecordingEnginex*record);
AGORA_API int Stopped(RecordingEnginex*record);
AGORA_API void UpdateMixModeSetting(RecordingEnginex *record, int width, int height, int isVideoMix);
AGORA_API RecordingEngineProperties* GetRecorderProperties(RecordingEnginex*record);
AGORA_API void UpdateStorageDir(RecordingEnginex *record, const char* dir);
AGORA_API void UpdateLayoutSetting(RecordingEnginex*record, int layoutMode, int maxVertPreLayoutUid);
AGORA_API int StartService(RecordingEnginex*record);
AGORA_API int StopService(RecordingEnginex*record);
AGORA_API int SetVideoMixingLayout(RecordingEnginex*record);
AGORA_API RecordingConfig* GetConfigInfo(RecordingEnginex*record);
AGORA_API int SetUserBackground(RecordingEnginex*record, unsigned int uid, const char* image_path);
AGORA_API int SetLogLevel(RecordingEnginex*record, int level);
AGORA_API void SetMediaKeepTime(RecordingEnginex*record, unsigned int time_ms);



AGORA_API void SetOnError(RecordingEnginex* handler, pOnError onError);
AGORA_API void SetOnWarning(RecordingEnginex* handler, pOnWarning onWarning);
AGORA_API void SetOnJoinChannelSuccess(RecordingEnginex* handler, pOnJoinChannelSuccess onJoinChannelSuccess);
AGORA_API void SetOnLeaveChannel (RecordingEnginex* handler, pOnLeaveChannel onLeaveChannel);
AGORA_API void SetOnUserJoined (RecordingEnginex* handler, pOnUserJoined onUserJoined);
AGORA_API void SetOnUserOffline (RecordingEnginex* handler, pOnUserOffline onUserOffline);
AGORA_API void SetAudioFrameReceived(RecordingEnginex* handler, pAudioFrameReceived audioFrameReceived);
AGORA_API void SetVideoFrameReceived(RecordingEnginex* handler, pVideoFrameReceived videoFrameReceived);
AGORA_API void SetOnActiveSpeaker(RecordingEnginex* handler, pOnActiveSpeaker onActiveSpeaker);
AGORA_API void SetOnAudioVolumeIndication(RecordingEnginex* handler, pOnAudioVolumeIndication onAudioVolumeIndication);
AGORA_API void SetOnFirstRemoteVideoDecoded(RecordingEnginex* handler, pOnFirstRemoteVideoDecoded onFirstRemoteVideoDecoded);
AGORA_API void SetOnFirstRemoteAudioFrame(RecordingEnginex* handler, pOnFirstRemoteAudioFrame onFirstRemoteAudioFrame);
AGORA_API void SetOnReceivingStreamStatusChanged(RecordingEnginex* handler, pOnReceivingStreamStatusChanged onReceivingStreamStatusChanged);
AGORA_API void SetOnConnectionLost(RecordingEnginex* handler, pOnConnectionLost onConnectionLost);
AGORA_API void SetOnConnectionInterrupted(RecordingEnginex* handler, pOnConnectionInterrupted onConnectionInterrupted);

#ifdef __cplusplus
};
#endif
#endif  //__C_AGORA_API_H__
