//#include <iostream>

//#include "common.h"
#include "agorasdk/Agorasdk.h"
#include "agora_api.h"
using namespace std;

#ifdef __cplusplus
extern "C" {
#endif

// RecordingConfig
AGORA_API RecordingConfig*  CreateRecordingConfig(){
    return NULL;
}

AGORA_API RecordingEnginex* CreateRecordingEngine(){
   return (RecordingEnginex*)new agora::AgoraSdk ();
}


AGORA_API void SetOnError(RecordingEnginex* handler, pOnError onError){

}

AGORA_API int CreateChannel(RecordingEnginex*record,const char * appid, const char* channelKey, const char* name, unsigned int uid, RecordingConfig* config){
    // 这些指正对象是否需要memcpy（free）
    return ((agora::AgoraSdk*)record)->createChannel(appid, channelKey, name, ( agora::linuxsdk::uid_t)uid, *((agora::recording::RecordingConfig*)config));
}
AGORA_API int SetVideoMixLayout(RecordingEnginex*record){
    return 0;
}
AGORA_API int LeaveChannel(RecordingEnginex*record){
return 0;
}
AGORA_API int StoppedOnError(RecordingEnginex*record){
return 0;
}
AGORA_API int Release(RecordingEnginex*record){
return 0;
}
AGORA_API int Stopped(RecordingEnginex*record){
return 0;
}
AGORA_API void UpdateMixModeSetting(RecordingEnginex *record, int width, int height, int isVideoMix){
return ;
}
AGORA_API RecordingEngineProperties* GetRecorderProperties(RecordingEnginex*record){
return NULL;
}
AGORA_API void UpdateStorageDir(const char* dir){
return ;
}
AGORA_API void UpdateLayoutSetting(RecordingEnginex*record, int layoutMode, int maxVertPreLayoutUid){
return ;
}
AGORA_API int StartService(RecordingEnginex*record){
return 0;
}
AGORA_API int StopService(RecordingEnginex*record){
return 0;
}
AGORA_API int SetVideoMixingLayout(RecordingEnginex*record){
return 0;
}
AGORA_API RecordingConfig* GetConfigInfo(RecordingEnginex*record){
return NULL;
}
AGORA_API int SetUserBackground(RecordingEnginex*record, unsigned int uid, const char* image_path){
return 0;
}
AGORA_API int SetLogLevel(RecordingEnginex*record, int level){
return 0;
}
AGORA_API void SetMediaKeepTime(RecordingEnginex*record, unsigned int time_ms){
return ;
}


#ifdef __cplusplus
};
#endif
