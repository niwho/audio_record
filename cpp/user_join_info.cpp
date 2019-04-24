#include <string>
#include <stdlib.h>
#include <string.h>
#include "user_join_info.h"
#include "include/IAgoraRecordingEngine.h"

AGORA_API int UJISetRecordingDir(UserJoinInfos* userJoinInfos, const char * dir){
    ((agora::linuxsdk::UserJoinInfos*)userJoinInfos)->storageDir = dir;
    return 0;
}

AGORA_API const char * UJIGetRecordingDir(UserJoinInfos* userJoinInfos){
    return  ((agora::linuxsdk::UserJoinInfos*)userJoinInfos)->storageDir;
}