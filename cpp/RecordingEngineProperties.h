#ifndef __C_RECRODING_ENGINE_PROPERTIES_H__
#define __C_RECRODING_ENGINE_PROPERTIES_H__



#include "common.h"
#ifdef __cplusplus
extern "C" {
#endif

typedef struct RecordingEngineProperties RecordingEngineProperties;

AGORA_API const char * REPGetStorageDir(RecordingEngineProperties*rep);

#ifdef __cplusplus
};
#endif
#endif  //__C_RECRODING_ENGINE_PROPERTIES_H__