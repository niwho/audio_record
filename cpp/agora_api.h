#ifndef __C_AGORA_API_H__
#define __C_AGORA_API_H__

#include "common.h"
//#include "record.h"
#ifdef __cplusplus
extern "C" {
#endif

typedef struct RecordingEnginex RecordingEnginex;

typedef void (*ponError)(RecordingEnginex*record, int error, int stat_code);

AGORA_API RecordingEnginex* CreateRecordingEngine();
AGORA_API void SetOnErrorEventHandler(RecordingEnginex* handler, ponError onError);

#ifdef __cplusplus
};
#endif
#endif  //__C_AGORA_API_H__
