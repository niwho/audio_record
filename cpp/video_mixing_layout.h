#ifndef __C_VIDEO_MIXING_LAYOUT_H__
#define __C_VIDEO_MIXING_LAYOUT_H__

#include "common.h"
#ifdef __cplusplus
extern "C" {
#endif

typedef struct VideoMixingLayout VideoMixingLayout;
typedef struct Region Region;


AGORA_API VideoMixingLayout* GetVideoMixingLayoutDefault();
AGORA_API int VMLSetCanvasWidth(VideoMixingLayout*,int);
AGORA_API int VMLSetCanvasHeight(VideoMixingLayout*,int);
AGORA_API int VMLSetBackgroundColor(VideoMixingLayout*, const char* );
AGORA_API int VMLSetAppData(VideoMixingLayout*, const char* , unsigned int);
AGORA_API int VMLSetRegions(Region*);


AGORA_API Region* GetRegionDefault();
AGORA_API int RegionSetUid(Region*, unsigned int);
AGORA_API int RegionSetX(Region*, double);
AGORA_API int RegionSetY(Region*, double);
AGORA_API int RegionSetWidth(Region*, double);
AGORA_API int RegionSetHeight(Region*, double);
AGORA_API int RegionSetZOrder(Region*, int);
AGORA_API int RegionSetAlpha(Region*, double);
AGORA_API int RegionSetRenderMode(Region*, int);
AGORA_API int RegionISetItem(Region*, int, Region*);




#ifdef __cplusplus
};
#endif
#endif  //__C_VIDEO_MIXING_LAYOUT_H__
