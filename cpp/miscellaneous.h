#ifndef __C_MISCELLANEOUS_H__
#define __C_MISCELLANEOUS_H__

#include "common.h"

#ifdef __cplusplus
extern "C" {
#endif

typedef struct AudioFrame AudioFrame;//miscellaneous.h
typedef struct VideoFrame VideoFrame;//miscellaneous.h
typedef struct AudioVolumeInfo AudioVolumeInfo;//miscellaneous

typedef struct VideoYuvFrame VideoYuvFrame;
typedef struct VideoH264Frame VideoH264Frame;
typedef struct VideoJpgFrame VideoJpgFrame;


typedef struct AudioPcmFrame AudioPcmFrame;
typedef struct AudioAacFrame AudioAacFrame;










unsigned int AVIGetVolume(AudioVolumeInfo* avi);
unsigned int AVIGetUid(AudioVolumeInfo* avi);

int AudioGetType(AudioFrame*frame);
AudioPcmFrame* AudioGetPcmFrame(AudioFrame*frame);
AudioAacFrame* AudioGetAacFrame(AudioFrame*frame);

unsigned long long PcmFrameGetMs(AudioPcmFrame*frame);
unsigned int PcmFrameGetChannelNum(AudioPcmFrame*frame);
unsigned int PcmFrameGetSampleBits(AudioPcmFrame*frame);
unsigned int PcmFrameGetSampleRates(AudioPcmFrame*frame);
unsigned int PcmFrameGetSamples(AudioPcmFrame*frame);
const unsigned char* PcmFrameGetBufData(AudioPcmFrame*frame);
unsigned int PcmFrameGetBufSize(AudioPcmFrame*frame);

const unsigned char* AacFrameBufData(AudioAacFrame*frame);
unsigned int AacFrameGetBufSize(AudioAacFrame*frame);
unsigned long long AacFrameGetMs(AudioAacFrame*frame);
unsigned int AacFrameGetChannelNum(AudioAacFrame*frame);
unsigned int AacFrameGetBitRates(AudioAacFrame*frame);



#ifdef __cplusplus
};
#endif
#endif  //__C_MISCELLANEOUS_H__
