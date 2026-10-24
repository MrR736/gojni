#ifndef __JNITOOLS_H__
#define __JNITOOLS_H__

#include <jni.h>

#ifdef __cplusplus
extern "C" {
#endif

#define MAKE_GOJNI_TOOLS(type) \
static inline type JNICALL gojni_null_##type(void) { return NULL; } /* \ */
// static inline int JNICALL gojni_is_null_##type(type v) { return v == NULL; }

MAKE_GOJNI_TOOLS(jarray)
MAKE_GOJNI_TOOLS(jweak)
MAKE_GOJNI_TOOLS(jstring)
MAKE_GOJNI_TOOLS(jclass)
MAKE_GOJNI_TOOLS(jmethodID)
MAKE_GOJNI_TOOLS(jfieldID)
MAKE_GOJNI_TOOLS(jobject)
MAKE_GOJNI_TOOLS(jobjectArray)
MAKE_GOJNI_TOOLS(jthrowable)

MAKE_GOJNI_TOOLS(jbooleanArray)
MAKE_GOJNI_TOOLS(jbyteArray)
MAKE_GOJNI_TOOLS(jcharArray)
MAKE_GOJNI_TOOLS(jshortArray)
MAKE_GOJNI_TOOLS(jintArray)
MAKE_GOJNI_TOOLS(jlongArray)
MAKE_GOJNI_TOOLS(jfloatArray)
MAKE_GOJNI_TOOLS(jdoubleArray)

#define MAKE_GOJNI_TOOLS_JVALUE(type,arr) \
static inline jvalue gojni_jvalue_##type(type v) { \
	jvalue x; \
	x.arr = v; \
	return x; \
}

MAKE_GOJNI_TOOLS_JVALUE(jint,i)
MAKE_GOJNI_TOOLS_JVALUE(jbyte,b)
MAKE_GOJNI_TOOLS_JVALUE(jboolean, z)
MAKE_GOJNI_TOOLS_JVALUE(jchar, c)
MAKE_GOJNI_TOOLS_JVALUE(jshort, s)
MAKE_GOJNI_TOOLS_JVALUE(jlong, j)
MAKE_GOJNI_TOOLS_JVALUE(jfloat, f)
MAKE_GOJNI_TOOLS_JVALUE(jdouble, d)
MAKE_GOJNI_TOOLS_JVALUE(jobject, l)


#undef MAKE_GOJNI_TOOLS_JVALUE
#undef MAKE_GOJNI_TOOLS

#ifdef __cplusplus
}
#endif

#endif