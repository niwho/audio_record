#include "RecordingEngineProperties.h"
#include "include/IAgoraRecordingEngine.h"

AGORA_API const char * REPGetStorageDir(RecordingEngineProperties*rep){
    return ((agora::recording::RecordingEngineProperties*)rep)->storageDir;
}
