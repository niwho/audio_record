#include "miscellaneous.h"
#include "include/IAgoraLinuxSdkCommon.h"
#include "include/IAgoraRecordingEngine.h"

unsigned int AVIGetVolume(AudioVolumeInfo* avi){

    return (unsigned int)(((agora::linuxsdk::AudioVolumeInfo*)avi) -> uid);
}
unsigned int AVIGetUid(AudioVolumeInfo* avi){

    return ((agora::linuxsdk::AudioVolumeInfo*)avi) -> volume;
}

int AudioGetType(AudioFrame*frame){

    return int(((agora::linuxsdk::AudioFrame*)frame) -> type);
}

AudioPcmFrame* AudioGetPcmFrame(AudioFrame*frame){

    return (AudioPcmFrame*)(&((((agora::linuxsdk::AudioFrame*)frame))->frame));
}

AudioAacFrame* AudioGetAacFrame(AudioFrame*frame){

    return (AudioAacFrame*)(&((((agora::linuxsdk::AudioFrame*)frame))->frame));
}

unsigned long long PcmFrameGetMs(AudioPcmFrame*frame){

    return ((agora::linuxsdk::AudioPcmFrame*)frame) -> frame_ms_;
}

unsigned int PcmFrameGetChannelNum(AudioPcmFrame*frame){
    return ((agora::linuxsdk::AudioPcmFrame*)frame) -> channels_;

}

unsigned int PcmFrameGetSampleBits(AudioPcmFrame*frame){

    return ((agora::linuxsdk::AudioPcmFrame*)frame) -> sample_bits_;
}

unsigned int PcmFrameGetSampleRates(AudioPcmFrame*frame){

    return ((agora::linuxsdk::AudioPcmFrame*)frame) -> sample_rates_;
}

unsigned int PcmFrameGetSamples(AudioPcmFrame*frame){

    return ((agora::linuxsdk::AudioPcmFrame*)frame) -> samples_;
}

const unsigned char* PcmFrameGetBufData(AudioPcmFrame*frame){

    return ((agora::linuxsdk::AudioPcmFrame*)frame) -> pcmBuf_;
}

unsigned int PcmFrameGetBufSize(AudioPcmFrame*frame){

    return ((agora::linuxsdk::AudioPcmFrame*)frame) -> pcmBufSize_;
}

const unsigned char* AacFrameBufData(AudioAacFrame*frame){

    return ((agora::linuxsdk::AudioAacFrame*)frame) -> aacBuf_;
}

unsigned int AacFrameGetBufSize(AudioAacFrame*frame){

    return ((agora::linuxsdk::AudioAacFrame*)frame) -> aacBufSize_;
}

unsigned long long AacFrameGetMs(AudioAacFrame*frame){

    return ((agora::linuxsdk::AudioAacFrame*)frame) -> frame_ms_;
}

unsigned int AacFrameGetChannelNum(AudioAacFrame*frame){

    return ((agora::linuxsdk::AudioAacFrame*)frame) -> channels_;
}

unsigned int AacFrameGetBitRates(AudioAacFrame*frame){
    return ((agora::linuxsdk::AudioAacFrame*)frame) -> bitrate_;
}
