//#include <iostream>

//#include "common.h"
#include "agorasdk/Agorasdk.h"
#include "agora_api.h"
//using namespace std;

//#ifdef __cplusplus
//extern "C" {
//#endif

AGORA_API RecordingEnginex* CreateRecordingEngine(int a){
   return (RecordingEnginex*)new agora::AgoraSdk ();
}

AGORA_API void SetOnErrorEventHandler(RecordingEnginex* handler, void* onError){

}

//#ifdef __cplusplus
//};
//#endif
