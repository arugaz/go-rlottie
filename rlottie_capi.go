package rlottie

/*
#cgo !windows LDFLAGS: -lm
#cgo windows CFLAGS: -DRLOTTIE_BUILD=0
#cgo windows CXXFLAGS: -DRLOTTIE_BUILD=0
#cgo CXXFLAGS: -std=c++14 -fno-exceptions -fno-asynchronous-unwind-tables -fno-rtti -Wall -fvisibility=hidden -Wnon-virtual-dtor -Woverloaded-virtual -Wno-unused-parameter -Wno-unused-value
#include "rlottie_capi.h"
*/
import "C"
import "unsafe"

// Lottie Animation
type Lottie_Animation *C.Lottie_Animation

// Set Lottie Configuration Cache Size
func LottieConfigureModelCacheSize(size uint) {
	C.lottie_configure_model_cache_size(C.size_t(size))
}

// Create New Lottie From Data
func LottieAnimationFromData(str string, key string, resource_path string) Lottie_Animation {
	return C.lottie_animation_from_data(C.CString(str), C.CString(key), C.CString(resource_path))
}

// Get Lottie Size
func LottieAnimationGetSize(animation Lottie_Animation) (uint, uint) {
	var width, height C.size_t
	C.lottie_animation_get_size(animation, &width, &height)
	return uint(width), uint(height)
}

// Get Total Frame
func LottieAnimationGetTotalframe(animation Lottie_Animation) uint {
	return uint(C.lottie_animation_get_totalframe(animation))
}

// Get Framerate
func LottieAnimationGetFramerate(animation Lottie_Animation) float64 {
	return float64(C.lottie_animation_get_framerate(animation))
}

// Get Lottie Duration
func LottieAnimationGetDuration(animation Lottie_Animation) float64 {
	return float64(C.lottie_animation_get_duration(animation))
}

// Get Pos Frame
func LottieAnimationGetFrameAtPos(animation Lottie_Animation, pos float32) uint {
	return uint(C.lottie_animation_get_frame_at_pos(animation, C.float(pos)))
}

// Render Lottie Animation
func LottieAnimationRender(animation Lottie_Animation, frame_num uint, buffer []byte, width uint, height uint, bytes_per_line uint) {
	C.lottie_animation_render(animation, C.size_t(frame_num), (*C.uint32_t)(unsafe.Pointer(&buffer[0])), C.size_t(width), C.size_t(height), C.size_t(bytes_per_line))
}

// Delete Lottie Animation
func LottieAnimationDestroy(animation Lottie_Animation) {
	C.lottie_animation_destroy(animation)
}
