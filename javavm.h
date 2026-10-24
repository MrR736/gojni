#ifndef __JAVAVM_H__
#define __JAVAVM_H__

#include <jni.h>

#ifdef __cplusplus
extern "C" {
#endif

static inline jint JNICALL gojni_DestroyJavaVM(JavaVM *vm) {
	return (*vm)->DestroyJavaVM(vm);
}

static inline jint JNICALL gojni_AttachCurrentThread(JavaVM *vm, void **penv, void *args) {
	return (*vm)->AttachCurrentThread(vm, penv, args);
}

static inline jint JNICALL gojni_DetachCurrentThread(JavaVM *vm) {
	return (*vm)->DetachCurrentThread(vm);
}

static inline jint JNICALL gojni_GetEnv(JavaVM *vm, void **penv, jint version) {
	return (*vm)->GetEnv(vm, penv, version);
}

static inline jint JNICALL gojni_AttachCurrentThreadAsDaemon(JavaVM *vm, void **penv, void *args) {
	return (*vm)->AttachCurrentThreadAsDaemon(vm, penv, args);
}

#ifdef __cplusplus
}
#endif

#endif