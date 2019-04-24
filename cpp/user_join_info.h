#ifndef __C_USER_JOIN_INFO_H__
#define __C_USER_JOIN_INFO_H__


#include "common.h"
#ifdef __cplusplus
extern "C" {
#endif

typedef struct UserJoinInfos UserJoinInfos;


//AGORA_API VideoMixingLayout* GetVideoMixingLayoutDefault();
AGORA_API int UJISetRecordingDir(UserJoinInfos*, const char * dir);
AGORA_API const char * UJIGetRecordingDir(UserJoinInfos*);






#ifdef __cplusplus
};
#endif
#endif  //__C_USER_JOIN_INFO_H__
