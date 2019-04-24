#include <string>
#include <stdlib.h>
#include <string.h>
#include "recording_config.h"
#include "include/IAgoraRecordingEngine.h"

//using namespace agora::recording;
using namespace std;

AGORA_API RecordingConfig*  CreateRecordingConfig(){
    return NULL;
}

AGORA_API void FreeRecordingConfig(RecordingConfig* cfg){
    if (cfg!=NULL) {
        delete (agora::recording::RecordingConfig*)cfg;
    }
}

AGORA_API ::RecordingConfig* GetRecordingConfigDefault(){
    return (::RecordingConfig*)new agora::recording::RecordingConfig();
}

AGORA_API int RCSetChannelProfileType(RecordingConfig*config,int channelProfileType){
    ((agora::recording::RecordingConfig*)config)->channelProfile = (agora::linuxsdk::CHANNEL_PROFILE_TYPE)channelProfileType;
    return 0;
}

int RCSetIsAudioOnly(RecordingConfig*config,int isAudioOnly){
     ((agora::recording::RecordingConfig*)config)->isAudioOnly=bool(isAudioOnly);
     return 0;
}

int RCSetIsMixinEnabled(RecordingConfig*config,int isMixingEnabled){
     ((agora::recording::RecordingConfig*)config)->isMixingEnabled=bool(isMixingEnabled);
     return 0;
}

int RCSetDecryptionMode(RecordingConfig* config,const char*decryptionMode){
    //char* m = (char*)malloc(strlen(decryptionMode)+1);
    //strcpy(m, decryptionMode);
    ((agora::recording::RecordingConfig*)config)->decryptionMode = (char*)decryptionMode;

    return 0;
}


int RCSetSecret(RecordingConfig*config,const char*secret){
    ((agora::recording::RecordingConfig*)config)->secret = (char*)secret;
    return  0;
}

int RCSetIdleLimitSec(RecordingConfig*config,int idleLimitSec){
     ((agora::recording::RecordingConfig*)config)->idleLimitSec = idleLimitSec;
     return 0;
}

int RCSetAppliteDir(RecordingConfig*config,const char*appliteDir){
     //const string s = string(appliteDir);
     //((agora::recording::RecordingConfig*)config)->appliteDir =const_cast<char*>(s.c_str());

     char* m = (char*)malloc(strlen(appliteDir)+1);
     strcpy(m, appliteDir);
     ((agora::recording::RecordingConfig*)config)->appliteDir = (char*)appliteDir;
     return 0;
}

int RCSetLowUdpPort(RecordingConfig*config,int lowUdpPort){
     ((agora::recording::RecordingConfig*)config)->lowUdpPort = lowUdpPort;
     return 0;
}

int RCSetHighUdpPort(RecordingConfig*config,int highUdpPort){
     ((agora::recording::RecordingConfig*)config)->highUdpPort= highUdpPort;
     return 0;
}

AGORA_API int RCGetChannelProfileType(RecordingConfig*config){
    return ((agora::recording::RecordingConfig*)config)->channelProfile ;
}

AGORA_API int RCGetIsAudioOnly(RecordingConfig*config){
    return int(((agora::recording::RecordingConfig*)config)->isAudioOnly);
}
AGORA_API int RCGetIsMixinEnabled(RecordingConfig*config){
    return int(((agora::recording::RecordingConfig*)config)->isMixingEnabled);

}
AGORA_API const char* RCGetDecryptionMode(RecordingConfig*config){
    return ((agora::recording::RecordingConfig*)config)->decryptionMode;
}

AGORA_API const char* RCGetSecret(RecordingConfig*config){
    return  ((agora::recording::RecordingConfig*)config)->secret ;
}
AGORA_API int RCGetIdleLimitSec(RecordingConfig*config){
    return ((agora::recording::RecordingConfig*)config)->idleLimitSec;
}
AGORA_API const char* RCGetAppliteDir(RecordingConfig*config){
    return ((agora::recording::RecordingConfig*)config)->appliteDir;
}
AGORA_API int RCGetLowUdpPort(RecordingConfig*config){
     return ((agora::recording::RecordingConfig*)config)->lowUdpPort ;
}
AGORA_API int RCGetHighUdpPort(RecordingConfig*config){
    return  ((agora::recording::RecordingConfig*)config)->highUdpPort;
}