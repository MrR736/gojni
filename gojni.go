package gojni

/*
 #cgo CFLAGS: -I/usr/lib/jvm/default-java/include
 #cgo CFLAGS: -I/usr/lib/jvm/default-java/include/linux
 #include <stdlib.h>
 #include <stdarg.h>
 #include <jni.h>
 #include "javavm.h"
 #include "jnienv.h"
 #include "jnitools.h"
*/
import "C"

import (
	"unicode/utf16"
	"unsafe"
)

type (
	jboolean = C.jboolean
	jchar    = C.jchar
	jshort   = C.jshort
	jfloat   = C.jfloat
	jdouble  = C.jdouble
	jbyte    = C.jbyte
	jint     = C.jint
	jlong    = C.jlong
	jsize    = C.jsize

	jobject         = C.jobject
	jclass          = C.jclass
	jthrowable      = C.jthrowable
	jstring         = C.jstring
	jarray          = C.jarray
	jbooleanArray   = C.jbooleanArray
	jbyteArray      = C.jbyteArray
	jcharArray      = C.jcharArray
	jshortArray     = C.jshortArray
	jintArray       = C.jintArray
	jlongArray      = C.jlongArray
	jfloatArray     = C.jfloatArray
	jdoubleArray    = C.jdoubleArray
	jobjectArray    = C.jobjectArray
	jweak           = C.jweak
	jvalue          = C.jvalue
	jfieldID        = C.jfieldID
	jmethodID       = C.jmethodID
	jobjectRefType  = C.jobjectRefType

	JNINativeMethod = C.JNINativeMethod

	JNIEnv          = C.JNIEnv
	JavaVM          = C.JavaVM
	JavaVMInitArgs  = C.JavaVMInitArgs
	JavaVMAttachArgs = C.JavaVMAttachArgs

	u16string []jchar
)

// Exported JNI primitive types.
type (
	JBoolean = jboolean
	JChar    = jchar
	JShort   = jshort
	JFloat   = jfloat
	JDouble  = jdouble
	JByte    = jbyte
	JInt     = jint
	JLong    = jlong
	JSize    = jsize
)

// Exported JNI reference types.
type (
	JObject         = jobject
	JClass          = jclass
	JThrowable      = jthrowable
	JString         = jstring
	JArray          = jarray
	JBooleanArray   = jbooleanArray
	JByteArray      = jbyteArray
	JCharArray      = jcharArray
	JShortArray     = jshortArray
	JIntArray       = jintArray
	JLongArray      = jlongArray
	JFloatArray     = jfloatArray
	JDoubleArray    = jdoubleArray
	JObjectArray    = jobjectArray
	JWeak           = jweak
	JValue          = jvalue
	JFieldID        = jfieldID
	JMethodID       = jmethodID
	JObjectRefType  = jobjectRefType
)

const (
	JNIFalse jboolean = 0
	JNITrue  jboolean = 1

	JNI_OK        jint = 0
	JNI_ERR       jint = -1
	JNI_EDETACHED jint = -2
	JNI_EVERSION  jint = -3
	JNI_ENOMEM    jint = -4
	JNI_EEXIST    jint = -5
	JNI_EINVAL    jint = -6

	JNI_COMMIT jint = 1
	JNI_ABORT  jint = 2

	JNI_VERSION_1_1 jint = 0x00010001
	JNI_VERSION_1_2 jint = 0x00010002
	JNI_VERSION_1_4 jint = 0x00010004
	JNI_VERSION_1_6 jint = 0x00010006
	JNI_VERSION_1_8 jint = 0x00010008
	JNI_VERSION_9   jint = 0x00090000
	JNI_VERSION_10  jint = 0x000a0000
	JNI_VERSION_19  jint = 0x00130000
	JNI_VERSION_20  jint = 0x00140000
	JNI_VERSION_21  jint = 0x00150000
)

type jnull struct {
	Class        jclass
	String       jstring
	Throwable    jthrowable
	Object       jobject
	Weak         jweak
	MethodID     jmethodID
	FieldID      jfieldID

	ObjectArray  jobjectArray
	BooleanArray jbooleanArray
	ByteArray    jbyteArray
	IntArray     jintArray
	CharArray    jcharArray
	ShortArray   jshortArray
	LongArray    jlongArray
	FloatArray   jfloatArray
	DoubleArray  jdoubleArray
}

func newJNull() jnull {
	return jnull{
		Class:        C.gojni_null_jclass(),
		String:       C.gojni_null_jstring(),
		Throwable:    C.gojni_null_jthrowable(),
		Object:       C.gojni_null_jobject(),
		Weak:          C.gojni_null_jweak(),
		MethodID:     C.gojni_null_jmethodID(),
		FieldID:      C.gojni_null_jfieldID(),

		ObjectArray:  C.gojni_null_jobjectArray(),
		BooleanArray: C.gojni_null_jbooleanArray(),
		ByteArray:    C.gojni_null_jbyteArray(),
		IntArray:     C.gojni_null_jintArray(),
		CharArray:    C.gojni_null_jcharArray(),
		ShortArray:   C.gojni_null_jshortArray(),
		LongArray:    C.gojni_null_jlongArray(),
		FloatArray:   C.gojni_null_jfloatArray(),
		DoubleArray:  C.gojni_null_jdoubleArray(),
	}
}

var jnil = newJNull()


func U16String(s string) u16string {
	runes := []rune(s)
	units := utf16.Encode(runes)

	out := make(u16string, len(units))
	for i, unit := range units {
		out[i] = jchar(unit)
	}
	return out
}

func (s u16string) String() string {
	units := make([]uint16, len(s))
	for i, unit := range s {
		units[i] = uint16(unit)
	}
	return string(utf16.Decode(units))
}

func goJNIValues(args ...any) []jvalue {
	values := make([]jvalue, len(args))

	for i, arg := range args {
		switch v := arg.(type) {
			case int32:
				values[i] = jvalue(C.gojni_jvalue_jint(C.jint(v)))

			case jint:
				values[i] = jvalue(C.gojni_jvalue_jint(C.jint(v)))

			case int8:
				values[i] = jvalue(C.gojni_jvalue_jbyte(jbyte(v)))

			case jbyte:
				values[i] = jvalue(C.gojni_jvalue_jbyte(jbyte(v)))

			case uint8:
				values[i] = jvalue(C.gojni_jvalue_jboolean(jboolean(v)))

			case jboolean:
				values[i] = jvalue(C.gojni_jvalue_jboolean(jboolean(v)))

			case uint16:
				values[i] = jvalue(C.gojni_jvalue_jchar(jchar(v)))

			case jchar:
				values[i] = jvalue(C.gojni_jvalue_jchar(jchar(v)))

			case int16:
				values[i] = jvalue(C.gojni_jvalue_jshort(jshort(v)))

			case jshort:
				values[i] = jvalue(C.gojni_jvalue_jshort(jshort(v)))

			case int64:
				values[i] = jvalue(C.gojni_jvalue_jlong(jlong(v)))

			case jlong:
				values[i] = jvalue(C.gojni_jvalue_jlong(jlong(v)))

			case float32:
				values[i] = jvalue(C.gojni_jvalue_jfloat(jfloat(v)))

			case jfloat:
				values[i] = jvalue(C.gojni_jvalue_jfloat(jfloat(v)))

			case float64:
				values[i] = jvalue(C.gojni_jvalue_jdouble(jdouble(v)))

			case jdouble:
				values[i] = jvalue(C.gojni_jvalue_jdouble(jdouble(v)))

			case jobject:
				values[i] = jvalue(C.gojni_jvalue_jobject(jobject(v)))

			default:
				return nil
		}
	}

	return values
}

type GoJNIEnv struct {
	ptr *JNIEnv
}

func WrapGoJNIEnv(env *JNIEnv) *GoJNIEnv {
	if env == nil {
		return nil
	}
	return &GoJNIEnv{ptr: env}
}

type goJavaVM struct {
	ptr *JavaVM
}

func WrapgoJavaVM(vm *JavaVM) *goJavaVM {
	if vm == nil {
		return nil
	}
	return &goJavaVM{ptr: vm}
}

func (env *GoJNIEnv) GetVersion() jint {
	if env == nil || env.ptr == nil {
		return JNI_ERR
	}
	return jint(C.gojni_GetVersion(env.ptr))
}

func (env *GoJNIEnv) DefineClass(name string,loader jobject,buf []jbyte) jclass {
	if env == nil || env.ptr == nil {
		return jnil.Class
	}

	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	var cbuf *jbyte
	if len(buf) > 0 {
		cbuf = (*jbyte)(unsafe.Pointer(&buf[0]))
	}

	return jclass(C.gojni_DefineClass(
		env.ptr,
		cname,
		loader,
		cbuf,
		jsize(len(buf)),
	))
}

func (env *GoJNIEnv) FindClass(name string) jclass {
	if env == nil || env.ptr == nil {
		return jnil.Class
	}

	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	return jclass(C.gojni_FindClass(
		env.ptr,
		cname,
	))
}

func (env *GoJNIEnv) FromReflectedMethod(method jobject) jmethodID {
	if env == nil || env.ptr == nil {
		return jnil.MethodID
	}

	return jmethodID(C.gojni_FromReflectedMethod(
		env.ptr,
		method,
	))
}

func (env *GoJNIEnv) FromReflectedField(field jobject) jfieldID {
	if env == nil || env.ptr == nil {
		return jnil.FieldID
	}

	return jfieldID(C.gojni_FromReflectedField(
		env.ptr,
		field,
	))
}

func (env *GoJNIEnv) GetSuperclass(sub jclass) jclass {
	if env == nil || env.ptr == nil {
		return jnil.Class
	}

	return jclass(C.gojni_GetSuperclass(
		env.ptr,
		sub,
	))
}

func (env *GoJNIEnv) IsAssignableFrom(sub jclass,sup jclass) jboolean {
	if env == nil || env.ptr == nil {
		return JNIFalse
	}

	return jboolean(C.gojni_IsAssignableFrom(
		env.ptr,
		sub,
		sup,
	))
}

func (env *GoJNIEnv) ToReflectedField(cls jclass,fieldID jfieldID,isStatic jboolean) jobject {
	if env == nil || env.ptr == nil {
		return jnil.Object
	}

	return jobject(C.gojni_ToReflectedField(
		env.ptr,
		cls,
		fieldID,
		isStatic,
	))
}

func (env *GoJNIEnv) Throw(obj jthrowable) jint {
	if env == nil || env.ptr == nil {
		return JNI_EINVAL
	}

	return jint(C.gojni_Throw(
		env.ptr,
		obj,
	))
}

func (env *GoJNIEnv) ThrowNew(cls jclass,name string) jint {
	if env == nil || env.ptr == nil {
		return JNI_EINVAL
	}

	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	return jint(C.gojni_ThrowNew(
		env.ptr,
		cls,
		cname,
	))
}

func (env *GoJNIEnv) ExceptionOccurred() jthrowable {
	if env == nil || env.ptr == nil {
		return jnil.Throwable
	}
	return jthrowable(C.gojni_ExceptionOccurred(
		env.ptr,
	))
}

func (env *GoJNIEnv) ExceptionDescribe() {
	if env == nil || env.ptr == nil {
		return
	}
	C.gojni_ExceptionDescribe(
		env.ptr,
	)
}

func (env *GoJNIEnv) ExceptionClear() {
	if env == nil || env.ptr == nil {
		return
	}
	C.gojni_ExceptionClear(
		env.ptr,
	)
}

func (env *GoJNIEnv) FatalError(msg string) {
	if env == nil || env.ptr == nil {
		return
	}

	cmsg := C.CString(msg)
	defer C.free(unsafe.Pointer(cmsg))

	C.gojni_FatalError(
		env.ptr,
		cmsg,
	)
}

func (env *GoJNIEnv) PushLocalFrame(capacity jint) jint {
	if env == nil || env.ptr == nil {
		return JNI_EINVAL
	}

	return jint(C.gojni_PushLocalFrame(
		env.ptr,
		capacity,
	))
}

func (env *GoJNIEnv) PopLocalFrame(result jobject) jobject {
	if env == nil || env.ptr == nil {
		return jnil.Object
	}

	return jobject(C.gojni_PopLocalFrame(
		env.ptr,
		result,
	))
}

func (env *GoJNIEnv) NewGlobalRef(lobj jobject) jobject {
	if env == nil || env.ptr == nil {
		return jnil.Object
	}

	return jobject(C.gojni_NewGlobalRef(
		env.ptr,
		lobj,
	))
}

func (env *GoJNIEnv) DeleteGlobalRef(obj jobject) {
	if env == nil || env.ptr == nil {
		return
	}

	C.gojni_DeleteGlobalRef(
		env.ptr,
		obj,
	)
}

func (env *GoJNIEnv) DeleteLocalRef(obj jobject) {
	if env == nil || env.ptr == nil {
		return
	}

	C.gojni_DeleteLocalRef(
		env.ptr,
		obj,
	)
}

func (env *GoJNIEnv) IsSameObject(obj1 jobject,obj2 jobject) jboolean {
	if env == nil || env.ptr == nil {
		return JNIFalse
	}

	return jboolean(C.gojni_IsSameObject(
		env.ptr,
		obj1,
		obj2,
	))
}

func (env *GoJNIEnv) NewLocalRef(lobj jobject) jobject {
	if env == nil || env.ptr == nil {
		return jnil.Object
	}

	return jobject(C.gojni_NewLocalRef(
		env.ptr,
		lobj,
	))
}

func (env *GoJNIEnv) EnsureLocalCapacity(capacity jint) jint {
	if env == nil || env.ptr == nil {
		return JNI_EINVAL
	}

	return jint(C.gojni_EnsureLocalCapacity(
		env.ptr,
		capacity,
	))
}

func (env *GoJNIEnv) AllocObject(cls jclass) jobject {
	if env == nil || env.ptr == nil {
		return jnil.Object
	}

	return jobject(C.gojni_AllocObject(
		env.ptr,
		cls,
	))
}

func (env *GoJNIEnv) NewObjectA(clazz jclass,methodID jmethodID,args []jvalue,) jobject {
	if env == nil || env.ptr == nil {
		return jnil.Object
	}

	var cargs *jvalue
	if len(args) > 0 {
		cargs = (*jvalue)(unsafe.Pointer(&args[0]))
	}

	return jobject(C.gojni_NewObjectA(
		env.ptr,
		clazz,
		methodID,
		cargs,
	))
}

func (env *GoJNIEnv) NewObject(clazz jclass,methodID jmethodID,args ...any,) jobject {
	if env == nil || env.ptr == nil {
		return jnil.Object
	}
	values := goJNIValues(args...)
	if len(args) > 0 && values == nil {
		return jnil.Object
	}
	return env.NewObjectA(clazz, methodID, values)
}


func (env *GoJNIEnv) GetObjectClass(obj jobject,) jclass {
	if env == nil || env.ptr == nil {
		return jnil.Class
	}

	return jclass(C.gojni_GetObjectClass(
		env.ptr,
		obj,
	))
}

func (env *GoJNIEnv) IsInstanceOf(obj jobject,clazz jclass) jboolean {
	if env == nil || env.ptr == nil {
		return JNIFalse
	}

	return jboolean(C.gojni_IsInstanceOf(
		env.ptr,
		obj,
		clazz,
	))
}

func (env *GoJNIEnv) GetMethodID(cls jclass,name string,sig string) jmethodID {
	if env == nil || env.ptr == nil {
		return jnil.MethodID
	}

	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	csig := C.CString(sig)
	defer C.free(unsafe.Pointer(csig))

	return jmethodID(C.gojni_GetMethodID(
		env.ptr,
		cls,
		cname,
		csig,
	))
}

func (env *GoJNIEnv) CallObjectMethodA(obj jobject,methodID jmethodID,args []jvalue,) jobject {
	if env == nil || env.ptr == nil {
		return jnil.Object
	}

	var cargs *jvalue
	if len(args) > 0 {
		cargs = (*jvalue)(unsafe.Pointer(&args[0]))
	}

	return jobject(C.gojni_CallObjectMethodA(
		env.ptr,
		obj,
		methodID,
		cargs,
	))
}

func (env *GoJNIEnv) CallObjectMethod(obj jobject,methodID jmethodID,args ...any,) jobject {
	if env == nil || env.ptr == nil {
		return jnil.Object
	}
	values := goJNIValues(args...)
	if len(args) > 0 && values == nil {
		return jnil.Object
	}
	return env.CallObjectMethodA(obj, methodID, values)
}


func (env *GoJNIEnv) CallBooleanMethodA(obj jobject,methodID jmethodID,args []jvalue,) jboolean {
	if env == nil || env.ptr == nil {
		return JNIFalse
	}

	var cargs *jvalue
	if len(args) > 0 {
		cargs = (*jvalue)(unsafe.Pointer(&args[0]))
	}

	return jboolean(C.gojni_CallBooleanMethodA(
		env.ptr,
		obj,
		methodID,
		cargs,
	))
}

func (env *GoJNIEnv) CallBooleanMethod(obj jobject,methodID jmethodID,args ...any,) jboolean {
	if env == nil || env.ptr == nil {
		return JNIFalse
	}
	values := goJNIValues(args...)
	if len(args) > 0 && values == nil {
		return JNIFalse
	}
	return env.CallBooleanMethodA(obj, methodID, values)
}


func (env *GoJNIEnv) CallByteMethodA(obj jobject,methodID jmethodID,args []jvalue,) jbyte {
	if env == nil || env.ptr == nil {
		return 0
	}

	var cargs *jvalue
	if len(args) > 0 {
		cargs = (*jvalue)(unsafe.Pointer(&args[0]))
	}

	return jbyte(C.gojni_CallByteMethodA(
		env.ptr,
		obj,
		methodID,
		cargs,
	))
}

func (env *GoJNIEnv) CallByteMethod(obj jobject,methodID jmethodID,args ...any,) jbyte {
	if env == nil || env.ptr == nil {
		return 0
	}
	values := goJNIValues(args...)
	if len(args) > 0 && values == nil {
		return 0
	}
	return env.CallByteMethodA(obj, methodID, values)
}

func (env *GoJNIEnv) CallCharMethodA(obj jobject,methodID jmethodID,args []jvalue,) jchar {
	if env == nil || env.ptr == nil {
		return jchar(0)
	}

	var cargs *jvalue
	if len(args) > 0 {
		cargs = (*jvalue)(unsafe.Pointer(&args[0]))
	}

	return jchar(C.gojni_CallCharMethodA(
		env.ptr,
		obj,
		methodID,
		cargs,
	))
}

func (env *GoJNIEnv) CallCharMethod(obj jobject,methodID jmethodID,args ...any,) jchar {
	if env == nil || env.ptr == nil {
		return 0
	}
	values := goJNIValues(args...)
	if len(args) > 0 && values == nil {
		return 0
	}
	return env.CallCharMethodA(obj, methodID, values)
}

func (env *GoJNIEnv) CallShortMethodA(obj jobject,methodID jmethodID,args []jvalue,) jshort {
	if env == nil || env.ptr == nil {
		return jshort(JNI_EINVAL)
	}

	var cargs *jvalue
	if len(args) > 0 {
		cargs = (*jvalue)(unsafe.Pointer(&args[0]))
	}

	return jshort(C.gojni_CallShortMethodA(
		env.ptr,
		obj,
		methodID,
		cargs,
	))
}

func (env *GoJNIEnv) CallShortMethod(obj jobject,methodID jmethodID,args ...any,) jshort {
	if env == nil || env.ptr == nil {
		return jshort(JNI_EINVAL)
	}
	values := goJNIValues(args...)
	if len(args) > 0 && values == nil {
		return jshort(JNI_ERR)
	}
	return env.CallShortMethodA(obj, methodID, values)
}


func (env *GoJNIEnv) CallIntMethodA(obj jobject,methodID jmethodID,args []jvalue,) jint {
	if env == nil || env.ptr == nil {
		return JNI_EINVAL
	}

	var cargs *jvalue
	if len(args) > 0 {
		cargs = (*jvalue)(unsafe.Pointer(&args[0]))
	}

	return jint(C.gojni_CallIntMethodA(
		env.ptr,
		obj,
		methodID,
		cargs,
	))
}

func (env *GoJNIEnv) CallIntMethod(obj jobject,methodID jmethodID,args ...any,) jint {
	if env == nil || env.ptr == nil {
		return JNI_EINVAL
	}
	values := goJNIValues(args...)
	if len(args) > 0 && values == nil {
		return JNI_ERR
	}
	return env.CallIntMethodA(obj, methodID, values)
}

func (env *GoJNIEnv) CallLongMethodA(obj jobject,methodID jmethodID,args []jvalue,) jlong {
	if env == nil || env.ptr == nil {
		return jlong(JNI_EINVAL)
	}

	var cargs *jvalue
	if len(args) > 0 {
		cargs = (*jvalue)(unsafe.Pointer(&args[0]))
	}

	return jlong(C.gojni_CallLongMethodA(
		env.ptr,
		obj,
		methodID,
		cargs,
	))
}

func (env *GoJNIEnv) CallLongMethod(obj jobject,methodID jmethodID,args ...any,) jlong {
	if env == nil || env.ptr == nil {
		return jlong(JNI_EINVAL)
	}
	values := goJNIValues(args...)
	if len(args) > 0 && values == nil {
		return jlong(JNI_ERR)
	}
	return env.CallLongMethodA(obj, methodID, values)
}

func (env *GoJNIEnv) CallFloatMethodA(obj jobject,methodID jmethodID,args []jvalue,) jfloat {
	if env == nil || env.ptr == nil {
		return jfloat(JNI_EINVAL)
	}

	var cargs *jvalue
	if len(args) > 0 {
		cargs = (*jvalue)(unsafe.Pointer(&args[0]))
	}

	return jfloat(C.gojni_CallFloatMethodA(
		env.ptr,
		obj,
		methodID,
		cargs,
	))
}

func (env *GoJNIEnv) CallFloatMethod(obj jobject,methodID jmethodID,args ...any,) jfloat {
	if env == nil || env.ptr == nil {
		return jfloat(JNI_EINVAL)
	}
	values := goJNIValues(args...)
	if len(args) > 0 && values == nil {
		return jfloat(JNI_ERR)
	}
	return env.CallFloatMethodA(obj, methodID, values)
}

func (env *GoJNIEnv) CallDoubleMethodA(obj jobject,methodID jmethodID,args []jvalue,) jdouble {
	if env == nil || env.ptr == nil {
		return jdouble(JNI_EINVAL)
	}

	var cargs *jvalue
	if len(args) > 0 {
		cargs = (*jvalue)(unsafe.Pointer(&args[0]))
	}

	return jdouble(C.gojni_CallDoubleMethodA(
		env.ptr,
		obj,
		methodID,
		cargs,
	))
}

func (env *GoJNIEnv) CallDoubleMethod(obj jobject,methodID jmethodID,args ...any,) jdouble {
	if env == nil || env.ptr == nil {
		return jdouble(JNI_EINVAL)
	}
	values := goJNIValues(args...)
	if len(args) > 0 && values == nil {
		return jdouble(JNI_ERR)
	}
	return env.CallDoubleMethodA(obj, methodID, values)
}

func (env *GoJNIEnv) CallVoidMethodA(obj jobject,methodID jmethodID,args []jvalue,) {
	if env == nil || env.ptr == nil {
		return
	}

	var cargs *jvalue
	if len(args) > 0 {
		cargs = (*jvalue)(unsafe.Pointer(&args[0]))
	}

	C.gojni_CallVoidMethodA(
		env.ptr,
		obj,
		methodID,
		cargs,
	)
}

func (env *GoJNIEnv) CallVoidMethod(obj jobject,methodID jmethodID,args ...any,) {
	if env == nil || env.ptr == nil {
		return
	}
	values := goJNIValues(args...)
	if len(args) > 0 && values == nil {
		return
	}
	env.CallVoidMethodA(obj, methodID, values)
}


func (env *GoJNIEnv) CallNonvirtualObjectMethodA(obj jobject,clazz jclass,methodID jmethodID,args []jvalue,) jobject {
	if env == nil || env.ptr == nil {
		return jnil.Object
	}

	var cargs *jvalue
	if len(args) > 0 {
		cargs = (*jvalue)(unsafe.Pointer(&args[0]))
	}

	return jobject(C.gojni_CallNonvirtualObjectMethodA(
		env.ptr,
		obj,
		clazz,
		methodID,
		cargs,
	))
}

func (env *GoJNIEnv) CallNonvirtualObjectMethod(obj jobject,clazz jclass,methodID jmethodID,args ...any,) jobject {
	if env == nil || env.ptr == nil {
		return jnil.Object
	}
	values := goJNIValues(args...)
	if len(args) > 0 && values == nil {
		return jnil.Object
	}
	return env.CallNonvirtualObjectMethodA(obj, clazz, methodID, values)
}


func (env *GoJNIEnv) CallNonvirtualBooleanMethodA(obj jobject,clazz jclass,methodID jmethodID,args []jvalue,) jboolean {
	if env == nil || env.ptr == nil {
		return JNIFalse
	}

	var cargs *jvalue
	if len(args) > 0 {
		cargs = (*jvalue)(unsafe.Pointer(&args[0]))
	}

	return jboolean(C.gojni_CallNonvirtualBooleanMethodA(
		env.ptr,
		obj,
		clazz,
		methodID,
		cargs,
	))
}

func (env *GoJNIEnv) CallNonvirtualBooleanMethod(obj jobject,clazz jclass,methodID jmethodID,args ...any,) jboolean {
	if env == nil || env.ptr == nil {
		return JNIFalse
	}
	values := goJNIValues(args...)
	if len(args) > 0 && values == nil {
		return JNIFalse
	}
	return env.CallNonvirtualBooleanMethodA(obj, clazz, methodID, values)
}


func (env *GoJNIEnv) CallNonvirtualByteMethodA(obj jobject,clazz jclass,methodID jmethodID,args []jvalue,) jbyte {
	if env == nil || env.ptr == nil {
		return 0
	}

	var cargs *jvalue
	if len(args) > 0 {
		cargs = (*jvalue)(unsafe.Pointer(&args[0]))
	}

	return jbyte(C.gojni_CallNonvirtualByteMethodA(
		env.ptr,
		obj,
		clazz,
		methodID,
		cargs,
	))
}

func (env *GoJNIEnv) CallNonvirtualByteMethod(obj jobject,clazz jclass,methodID jmethodID,args ...any,) jbyte {
	if env == nil || env.ptr == nil {
		return 0
	}
	values := goJNIValues(args...)
	if len(args) > 0 && values == nil {
		return 0
	}
	return env.CallNonvirtualByteMethodA(obj, clazz, methodID, values)
}

func (env *GoJNIEnv) CallNonvirtualCharMethodA(obj jobject,clazz jclass,methodID jmethodID,args []jvalue,) jchar {
	if env == nil || env.ptr == nil {
		return jchar(0)
	}

	var cargs *jvalue
	if len(args) > 0 {
		cargs = (*jvalue)(unsafe.Pointer(&args[0]))
	}

	return jchar(C.gojni_CallNonvirtualCharMethodA(
		env.ptr,
		obj,
		clazz,
		methodID,
		cargs,
	))
}

func (env *GoJNIEnv) CallNonvirtualCharMethod(obj jobject,clazz jclass,methodID jmethodID,args ...any,) jchar {
	if env == nil || env.ptr == nil {
		return 0
	}
	values := goJNIValues(args...)
	if len(args) > 0 && values == nil {
		return 0
	}
	return env.CallNonvirtualCharMethodA(obj, clazz, methodID, values)
}

func (env *GoJNIEnv) CallNonvirtualShortMethodA(obj jobject,clazz jclass,methodID jmethodID,args []jvalue,) jshort {
	if env == nil || env.ptr == nil {
		return jshort(JNI_EINVAL)
	}

	var cargs *jvalue
	if len(args) > 0 {
		cargs = (*jvalue)(unsafe.Pointer(&args[0]))
	}

	return jshort(C.gojni_CallNonvirtualShortMethodA(
		env.ptr,
		obj,
		clazz,
		methodID,
		cargs,
	))
}

func (env *GoJNIEnv) CallNonvirtualShortMethod(obj jobject,clazz jclass,methodID jmethodID,args ...any,) jshort {
	if env == nil || env.ptr == nil {
		return jshort(JNI_EINVAL)
	}
	values := goJNIValues(args...)
	if len(args) > 0 && values == nil {
		return jshort(JNI_ERR)
	}
	return env.CallNonvirtualShortMethodA(obj, clazz, methodID, values)
}


func (env *GoJNIEnv) CallNonvirtualIntMethodA(obj jobject,clazz jclass,methodID jmethodID,args []jvalue,) jint {
	if env == nil || env.ptr == nil {
		return JNI_EINVAL
	}

	var cargs *jvalue
	if len(args) > 0 {
		cargs = (*jvalue)(unsafe.Pointer(&args[0]))
	}

	return jint(C.gojni_CallNonvirtualIntMethodA(
		env.ptr,
		obj,
		clazz,
		methodID,
		cargs,
	))
}

func (env *GoJNIEnv) CallNonvirtualIntMethod(obj jobject,clazz jclass,methodID jmethodID,args ...any,) jint {
	if env == nil || env.ptr == nil {
		return JNI_EINVAL
	}
	values := goJNIValues(args...)
	if len(args) > 0 && values == nil {
		return JNI_ERR
	}
	return env.CallNonvirtualIntMethodA(obj, clazz, methodID, values)
}

func (env *GoJNIEnv) CallNonvirtualLongMethodA(obj jobject,clazz jclass,methodID jmethodID,args []jvalue,) jlong {
	if env == nil || env.ptr == nil {
		return jlong(JNI_EINVAL)
	}

	var cargs *jvalue
	if len(args) > 0 {
		cargs = (*jvalue)(unsafe.Pointer(&args[0]))
	}

	return jlong(C.gojni_CallNonvirtualLongMethodA(
		env.ptr,
		obj,
		clazz,
		methodID,
		cargs,
	))
}

func (env *GoJNIEnv) CallNonvirtualLongMethod(obj jobject,clazz jclass,methodID jmethodID,args ...any,) jlong {
	if env == nil || env.ptr == nil {
		return jlong(JNI_EINVAL)
	}
	values := goJNIValues(args...)
	if len(args) > 0 && values == nil {
		return jlong(JNI_ERR)
	}
	return env.CallNonvirtualLongMethodA(obj, clazz, methodID, values)
}

func (env *GoJNIEnv) CallNonvirtualFloatMethodA(obj jobject,clazz jclass,methodID jmethodID,args []jvalue,) jfloat {
	if env == nil || env.ptr == nil {
		return jfloat(JNI_EINVAL)
	}

	var cargs *jvalue
	if len(args) > 0 {
		cargs = (*jvalue)(unsafe.Pointer(&args[0]))
	}

	return jfloat(C.gojni_CallNonvirtualFloatMethodA(
		env.ptr,
		obj,
		clazz,
		methodID,
		cargs,
	))
}

func (env *GoJNIEnv) CallNonvirtualFloatMethod(obj jobject,clazz jclass,methodID jmethodID,args ...any,) jfloat {
	if env == nil || env.ptr == nil {
		return jfloat(JNI_EINVAL)
	}
	values := goJNIValues(args...)
	if len(args) > 0 && values == nil {
		return jfloat(JNI_ERR)
	}
	return env.CallNonvirtualFloatMethodA(obj, clazz, methodID, values)
}

func (env *GoJNIEnv) CallNonvirtualDoubleMethodA(obj jobject,clazz jclass,methodID jmethodID,args []jvalue,) jdouble {
	if env == nil || env.ptr == nil {
		return jdouble(JNI_EINVAL)
	}

	var cargs *jvalue
	if len(args) > 0 {
		cargs = (*jvalue)(unsafe.Pointer(&args[0]))
	}

	return jdouble(C.gojni_CallNonvirtualDoubleMethodA(
		env.ptr,
		obj,
		clazz,
		methodID,
		cargs,
	))
}

func (env *GoJNIEnv) CallNonvirtualDoubleMethod(obj jobject,clazz jclass,methodID jmethodID,args ...any,) jdouble {
	if env == nil || env.ptr == nil {
		return jdouble(JNI_EINVAL)
	}
	values := goJNIValues(args...)
	if len(args) > 0 && values == nil {
		return jdouble(JNI_ERR)
	}
	return env.CallNonvirtualDoubleMethodA(obj, clazz, methodID, values)
}

func (env *GoJNIEnv) CallNonvirtualVoidMethodA(obj jobject,clazz jclass,methodID jmethodID,args []jvalue,) {
	if env == nil || env.ptr == nil {
		return
	}

	var cargs *jvalue
	if len(args) > 0 {
		cargs = (*jvalue)(unsafe.Pointer(&args[0]))
	}

	C.gojni_CallNonvirtualVoidMethodA(
		env.ptr,
		obj,
		clazz,
		methodID,
		cargs,
	)
}

func (env *GoJNIEnv) CallNonvirtualVoidMethod(obj jobject,clazz jclass,methodID jmethodID,args ...any,) {
	if env == nil || env.ptr == nil {
		return
	}
	values := goJNIValues(args...)
	if len(args) > 0 && values == nil {
		return
	}
	env.CallNonvirtualVoidMethodA(obj, clazz, methodID, values)
}

func (env *GoJNIEnv) GetFieldID(cls jclass,name string,sig string) jfieldID {
	if env == nil || env.ptr == nil {
		return jnil.FieldID
	}

	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	csig := C.CString(sig)
	defer C.free(unsafe.Pointer(csig))

	return jfieldID(C.gojni_GetFieldID(
		env.ptr,
		cls,
		cname,
		csig,
	))
}

func (env *GoJNIEnv) GetObjectField(obj jobject,fieldID jfieldID) jobject {
	if env == nil || env.ptr == nil {
		return jnil.Object
	}

	return jobject(C.gojni_GetObjectField(
		env.ptr,
		obj,
		fieldID,
	))
}

func (env *GoJNIEnv) GetBooleanField(obj jobject,fieldID jfieldID) jboolean {
	if env == nil || env.ptr == nil {
		return JNIFalse
	}

	return jboolean(C.gojni_GetBooleanField(
		env.ptr,
		obj,
		fieldID,
	))
}

func (env *GoJNIEnv) GetByteField(obj jobject,fieldID jfieldID) jbyte {
	if env == nil || env.ptr == nil {
		return 0
	}

	return jbyte(C.gojni_GetByteField(
		env.ptr,
		obj,
		fieldID,
	))
}

func (env *GoJNIEnv) GetCharField(obj jobject,fieldID jfieldID) jchar {
	if env == nil || env.ptr == nil {
		return jchar(0)
	}

	return jchar(C.gojni_GetCharField(
		env.ptr,
		obj,
		fieldID,
	))
}

func (env *GoJNIEnv) GetShortField(obj jobject,fieldID jfieldID) jshort {
	if env == nil || env.ptr == nil {
		return jshort(JNI_EINVAL)
	}

	return jshort(C.gojni_GetShortField(
		env.ptr,
		obj,
		fieldID,
	))
}

func (env *GoJNIEnv) GetIntField(obj jobject,fieldID jfieldID) jint {
	if env == nil || env.ptr == nil {
		return JNI_EINVAL
	}

	return jint(C.gojni_GetIntField(
		env.ptr,
		obj,
		fieldID,
	))
}

func (env *GoJNIEnv) GetLongField(obj jobject,fieldID jfieldID) jlong {
	if env == nil || env.ptr == nil {
		return jlong(JNI_EINVAL)
	}

	return jlong(C.gojni_GetLongField(
		env.ptr,
		obj,
		fieldID,
	))
}

func (env *GoJNIEnv) GetFloatField(obj jobject,fieldID jfieldID) jfloat {
	if env == nil || env.ptr == nil {
		return jfloat(JNI_EINVAL)
	}

	return jfloat(C.gojni_GetFloatField(
		env.ptr,
		obj,
		fieldID,
	))
}

func (env *GoJNIEnv) GetDoubleField(obj jobject,fieldID jfieldID) jdouble {
	if env == nil || env.ptr == nil {
		return jdouble(JNI_EINVAL)
	}

	return jdouble(C.gojni_GetDoubleField(
		env.ptr,
		obj,
		fieldID,
	))
}

func (env *GoJNIEnv) SetObjectField(obj jobject,fieldID jfieldID,val jobject) {
	if env == nil || env.ptr == nil {
		return
	}

	C.gojni_SetObjectField(
		env.ptr,
		obj,
		fieldID,
		val,
	)
}

func (env *GoJNIEnv) SetBooleanField(obj jobject,fieldID jfieldID,val jboolean) {
	if env == nil || env.ptr == nil {
		return
	}

	C.gojni_SetBooleanField(
		env.ptr,
		obj,
		fieldID,
		val,
	)
}

func (env *GoJNIEnv) SetCharField(obj jobject,fieldID jfieldID,val jchar) {
	if env == nil || env.ptr == nil {
		return
	}

	C.gojni_SetCharField(
		env.ptr,
		obj,
		fieldID,
		val,
	)
}

func (env *GoJNIEnv) SetShortField(obj jobject,fieldID jfieldID,val jshort) {
	if env == nil || env.ptr == nil {
		return
	}

	C.gojni_SetShortField(
		env.ptr,
		obj,
		fieldID,
		val,
	)
}

func (env *GoJNIEnv) SetIntField(obj jobject,fieldID jfieldID,val jint) {
	if env == nil || env.ptr == nil {
		return
	}

	C.gojni_SetIntField(
		env.ptr,
		obj,
		fieldID,
		val,
	)
}

func (env *GoJNIEnv) SetLongField(obj jobject,fieldID jfieldID,val jlong) {
	if env == nil || env.ptr == nil {
		return
	}

	C.gojni_SetLongField(
		env.ptr,
		obj,
		fieldID,
		val,
	)
}

func (env *GoJNIEnv) SetFloatField(obj jobject,fieldID jfieldID,val jfloat) {
	if env == nil || env.ptr == nil {
		return
	}

	C.gojni_SetFloatField(
		env.ptr,
		obj,
		fieldID,
		val,
	)
}

func (env *GoJNIEnv) SetDoubleField(obj jobject,fieldID jfieldID,val jdouble) {
	if env == nil || env.ptr == nil {
		return
	}

	C.gojni_SetDoubleField(
		env.ptr,
		obj,
		fieldID,
		val,
	)
}

func (env *GoJNIEnv) GetStaticMethodID(clazz jclass,name string,sig string) jmethodID {
	if env == nil || env.ptr == nil {
		return jnil.MethodID
	}

	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	csig := C.CString(sig)
	defer C.free(unsafe.Pointer(csig))

	return jmethodID(C.gojni_GetStaticMethodID(
		env.ptr,
		clazz,
		cname,
		csig,
	))
}

func (env *GoJNIEnv) CallStaticObjectMethodA(clazz jclass,methodID jmethodID,args []jvalue,) jobject {
	if env == nil || env.ptr == nil {
		return jnil.Object
	}

	var cargs *jvalue
	if len(args) > 0 {
		cargs = (*jvalue)(unsafe.Pointer(&args[0]))
	}

	return jobject(C.gojni_CallStaticObjectMethodA(
		env.ptr,
		clazz,
		methodID,
		cargs,
	))
}

func (env *GoJNIEnv) CallStaticObjectMethod(clazz jclass,methodID jmethodID,args ...any,) jobject {
	if env == nil || env.ptr == nil {
		return jnil.Object
	}
	values := goJNIValues(args...)
	if len(args) > 0 && values == nil {
		return jnil.Object
	}
	return env.CallStaticObjectMethodA(clazz, methodID, values)
}

func (env *GoJNIEnv) CallStaticBooleanMethodA(clazz jclass,methodID jmethodID,args []jvalue,) jboolean {
	if env == nil || env.ptr == nil {
		return JNIFalse
	}

	var cargs *jvalue
	if len(args) > 0 {
		cargs = (*jvalue)(unsafe.Pointer(&args[0]))
	}

	return jboolean(C.gojni_CallStaticBooleanMethodA(
		env.ptr,
		clazz,
		methodID,
		cargs,
	))
}

func (env *GoJNIEnv) CallStaticBooleanMethod(clazz jclass,methodID jmethodID,args ...any,) jboolean {
	if env == nil || env.ptr == nil {
		return JNIFalse
	}
	values := goJNIValues(args...)
	if len(args) > 0 && values == nil {
		return JNIFalse
	}
	return env.CallStaticBooleanMethodA(clazz, methodID, values)
}

func (env *GoJNIEnv) CallStaticByteMethodA(clazz jclass,methodID jmethodID,args []jvalue,) jbyte {
	if env == nil || env.ptr == nil {
		return 0
	}

	var cargs *jvalue
	if len(args) > 0 {
		cargs = (*jvalue)(unsafe.Pointer(&args[0]))
	}

	return jbyte(C.gojni_CallStaticByteMethodA(
		env.ptr,
		clazz,
		methodID,
		cargs,
	))
}

func (env *GoJNIEnv) CallStaticByteMethod(clazz jclass,methodID jmethodID,args ...any,) jbyte {
	if env == nil || env.ptr == nil {
		return 0
	}
	values := goJNIValues(args...)
	if len(args) > 0 && values == nil {
		return 0
	}
	return env.CallStaticByteMethodA(clazz, methodID, values)
}

func (env *GoJNIEnv) CallStaticCharMethodA(clazz jclass,methodID jmethodID,args []jvalue,) jchar {
	if env == nil || env.ptr == nil {
		return jchar(0)
	}

	var cargs *jvalue
	if len(args) > 0 {
		cargs = (*jvalue)(unsafe.Pointer(&args[0]))
	}

	return jchar(C.gojni_CallStaticCharMethodA(
		env.ptr,
		clazz,
		methodID,
		cargs,
	))
}

func (env *GoJNIEnv) CallStaticCharMethod(clazz jclass,methodID jmethodID,args ...any,) jchar {
	if env == nil || env.ptr == nil {
		return jchar(0)
	}
	values := goJNIValues(args...)
	if len(args) > 0 && values == nil {
		return jchar(0)
	}
	return env.CallStaticCharMethodA(clazz, methodID, values)
}

func (env *GoJNIEnv) CallStaticShortMethodA(clazz jclass,methodID jmethodID,args []jvalue,) jshort {
	if env == nil || env.ptr == nil {
		return jshort(JNI_EINVAL)
	}

	var cargs *jvalue
	if len(args) > 0 {
		cargs = (*jvalue)(unsafe.Pointer(&args[0]))
	}

	return jshort(C.gojni_CallStaticShortMethodA(
		env.ptr,
		clazz,
		methodID,
		cargs,
	))
}

func (env *GoJNIEnv) CallStaticShortMethod(clazz jclass,methodID jmethodID,args ...any,) jshort {
	if env == nil || env.ptr == nil {
		return jshort(JNI_EINVAL)
	}
	values := goJNIValues(args...)
	if len(args) > 0 && values == nil {
		return jshort(JNI_ERR)
	}
	return env.CallStaticShortMethodA(clazz, methodID, values)
}

func (env *GoJNIEnv) CallStaticIntMethodA(clazz jclass,methodID jmethodID,args []jvalue,) jint {
	if env == nil || env.ptr == nil {
		return JNI_EINVAL
	}

	var cargs *jvalue
	if len(args) > 0 {
		cargs = (*jvalue)(unsafe.Pointer(&args[0]))
	}

	return jint(C.gojni_CallStaticIntMethodA(
		env.ptr,
		clazz,
		methodID,
		cargs,
	))
}

func (env *GoJNIEnv) CallStaticIntMethod(clazz jclass,methodID jmethodID,args ...any,) jint {
	if env == nil || env.ptr == nil {
		return JNI_EINVAL
	}
	values := goJNIValues(args...)
	if len(args) > 0 && values == nil {
		return JNI_ERR
	}
	return env.CallStaticIntMethodA(clazz, methodID, values)
}

func (env *GoJNIEnv) CallStaticLongMethodA(clazz jclass,methodID jmethodID,args []jvalue,) jlong {
	if env == nil || env.ptr == nil {
		return jlong(JNI_EINVAL)
	}

	var cargs *jvalue
	if len(args) > 0 {
		cargs = (*jvalue)(unsafe.Pointer(&args[0]))
	}

	return jlong(C.gojni_CallStaticLongMethodA(
		env.ptr,
		clazz,
		methodID,
		cargs,
	))
}

func (env *GoJNIEnv) CallStaticLongMethod(clazz jclass,methodID jmethodID,args ...any,) jlong {
	if env == nil || env.ptr == nil {
		return jlong(JNI_EINVAL)
	}
	values := goJNIValues(args...)
	if len(args) > 0 && values == nil {
		return jlong(JNI_ERR)
	}
	return env.CallStaticLongMethodA(clazz, methodID, values)
}

func (env *GoJNIEnv) CallStaticFloatMethodA(clazz jclass,methodID jmethodID,args []jvalue,) jfloat {
	if env == nil || env.ptr == nil {
		return jfloat(JNI_EINVAL)
	}

	var cargs *jvalue
	if len(args) > 0 {
		cargs = (*jvalue)(unsafe.Pointer(&args[0]))
	}

	return jfloat(C.gojni_CallStaticFloatMethodA(
		env.ptr,
		clazz,
		methodID,
		cargs,
	))
}

func (env *GoJNIEnv) CallStaticFloatMethod(clazz jclass,methodID jmethodID,args ...any,) jfloat {
	if env == nil || env.ptr == nil {
		return jfloat(JNI_EINVAL)
	}
	values := goJNIValues(args...)
	if len(args) > 0 && values == nil {
		return jfloat(JNI_ERR)
	}
	return env.CallStaticFloatMethodA(clazz, methodID, values)
}

func (env *GoJNIEnv) CallStaticDoubleMethodA(clazz jclass,methodID jmethodID,args []jvalue,) jdouble {
	if env == nil || env.ptr == nil {
		return jdouble(JNI_EINVAL)
	}

	var cargs *jvalue
	if len(args) > 0 {
		cargs = (*jvalue)(unsafe.Pointer(&args[0]))
	}

	return jdouble(C.gojni_CallStaticDoubleMethodA(
		env.ptr,
		clazz,
		methodID,
		cargs,
	))
}

func (env *GoJNIEnv) CallStaticDoubleMethod(clazz jclass,methodID jmethodID,args ...any,) jdouble {
	if env == nil || env.ptr == nil {
		return jdouble(JNI_EINVAL)
	}
	values := goJNIValues(args...)
	if len(args) > 0 && values == nil {
		return jdouble(JNI_ERR)
	}
	return env.CallStaticDoubleMethodA(clazz, methodID, values)
}

func (env *GoJNIEnv) CallStaticVoidMethodA(clazz jclass,methodID jmethodID,args []jvalue,) {
	if env == nil || env.ptr == nil {
		return
	}

	var cargs *jvalue
	if len(args) > 0 {
		cargs = (*jvalue)(unsafe.Pointer(&args[0]))
	}

	C.gojni_CallStaticVoidMethodA(
		env.ptr,
		clazz,
		methodID,
		cargs,
	)
}

func (env *GoJNIEnv) CallStaticVoidMethod(clazz jclass,methodID jmethodID,args ...any,) {
	if env == nil || env.ptr == nil {
		return
	}
	values := goJNIValues(args...)
	if len(args) > 0 && values == nil {
		return
	}
	env.CallStaticVoidMethodA(clazz, methodID, values)
}

func (env *GoJNIEnv) GetStaticFieldID(cls jclass,name string,sig string) jfieldID {
	if env == nil || env.ptr == nil {
		return jnil.FieldID
	}

	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	csig := C.CString(sig)
	defer C.free(unsafe.Pointer(csig))

	return jfieldID(C.gojni_GetStaticFieldID(
		env.ptr,
		cls,
		cname,
		csig,
	))
}

func (env *GoJNIEnv) GetStaticObjectField(clazz jclass,fieldID jfieldID) jobject {
	if env == nil || env.ptr == nil {
		return jnil.Object
	}

	return jobject(C.gojni_GetStaticObjectField(
		env.ptr,
		clazz,
		fieldID,
	))
}

func (env *GoJNIEnv) GetStaticBooleanField(clazz jclass,fieldID jfieldID) jboolean {
	if env == nil || env.ptr == nil {
		return JNIFalse
	}

	return jboolean(C.gojni_GetStaticBooleanField(
		env.ptr,
		clazz,
		fieldID,
	))
}

func (env *GoJNIEnv) GetStaticByteField(clazz jclass,fieldID jfieldID) jbyte {
	if env == nil || env.ptr == nil {
		return 0
	}

	return jbyte(C.gojni_GetStaticByteField(
		env.ptr,
		clazz,
		fieldID,
	))
}

func (env *GoJNIEnv) GetStaticCharField(clazz jclass,fieldID jfieldID) jchar {
	if env == nil || env.ptr == nil {
		return jchar(0)
	}

	return jchar(C.gojni_GetStaticCharField(
		env.ptr,
		clazz,
		fieldID,
	))
}

func (env *GoJNIEnv) GetStaticShortField(clazz jclass,fieldID jfieldID) jshort {
	if env == nil || env.ptr == nil {
		return jshort(JNI_EINVAL)
	}

	return jshort(C.gojni_GetStaticShortField(
		env.ptr,
		clazz,
		fieldID,
	))
}

func (env *GoJNIEnv) GetStaticIntField(clazz jclass,fieldID jfieldID) jint {
	if env == nil || env.ptr == nil {
		return JNI_EINVAL
	}

	return jint(C.gojni_GetStaticIntField(
		env.ptr,
		clazz,
		fieldID,
	))
}

func (env *GoJNIEnv) GetStaticLongField(clazz jclass,fieldID jfieldID) jlong {
	if env == nil || env.ptr == nil {
		return jlong(JNI_EINVAL)
	}

	return jlong(C.gojni_GetStaticLongField(
		env.ptr,
		clazz,
		fieldID,
	))
}

func (env *GoJNIEnv) GetStaticFloatField(clazz jclass,fieldID jfieldID) jfloat {
	if env == nil || env.ptr == nil {
		return jfloat(JNI_EINVAL)
	}

	return jfloat(C.gojni_GetStaticFloatField(
		env.ptr,
		clazz,
		fieldID,
	))
}

func (env *GoJNIEnv) GetStaticDoubleField(clazz jclass,fieldID jfieldID) jdouble {
	if env == nil || env.ptr == nil {
		return jdouble(JNI_EINVAL)
	}

	return jdouble(C.gojni_GetStaticDoubleField(
		env.ptr,
		clazz,
		fieldID,
	))
}

func (env *GoJNIEnv) SetStaticObjectField(clazz jclass,fieldID jfieldID,value jobject) {
	if env == nil || env.ptr == nil {
		return
	}

	C.gojni_SetStaticObjectField(
		env.ptr,
		clazz,
		fieldID,
		value,
	)
}

func (env *GoJNIEnv) SetStaticBooleanField(clazz jclass,fieldID jfieldID,value jboolean) {
	if env == nil || env.ptr == nil {
		return
	}

	C.gojni_SetStaticBooleanField(
		env.ptr,
		clazz,
		fieldID,
		value,
	)
}

func (env *GoJNIEnv) SetStaticByteField(clazz jclass,fieldID jfieldID,value jbyte) {
	if env == nil || env.ptr == nil {
		return
	}

	C.gojni_SetStaticByteField(
		env.ptr,
		clazz,
		fieldID,
		value,
	)
}

func (env *GoJNIEnv) SetStaticCharField(clazz jclass,fieldID jfieldID,value jchar) {
	if env == nil || env.ptr == nil {
		return
	}

	C.gojni_SetStaticCharField(
		env.ptr,
		clazz,
		fieldID,
		value,
	)
}

func (env *GoJNIEnv) SetStaticShortField(clazz jclass,fieldID jfieldID,value jshort) {
	if env == nil || env.ptr == nil {
		return
	}

	C.gojni_SetStaticShortField(
		env.ptr,
		clazz,
		fieldID,
		value,
	)
}

func (env *GoJNIEnv) SetStaticIntField(clazz jclass,fieldID jfieldID,value jint) {
	if env == nil || env.ptr == nil {
		return
	}

	C.gojni_SetStaticIntField(
		env.ptr,
		clazz,
		fieldID,
		value,
	)
}

func (env *GoJNIEnv) SetStaticLongField(clazz jclass,fieldID jfieldID,value jlong) {
	if env == nil || env.ptr == nil {
		return
	}

	C.gojni_SetStaticLongField(
		env.ptr,
		clazz,
		fieldID,
		value,
	)
}

func (env *GoJNIEnv) SetStaticFloatField(clazz jclass,fieldID jfieldID,value jfloat) {
	if env == nil || env.ptr == nil {
		return
	}

	C.gojni_SetStaticFloatField(
		env.ptr,
		clazz,
		fieldID,
		value,
	)
}

func (env *GoJNIEnv) SetStaticDoubleField(clazz jclass,fieldID jfieldID,value jdouble) {
	if env == nil || env.ptr == nil {
		return
	}

	C.gojni_SetStaticDoubleField(
		env.ptr,
		clazz,
		fieldID,
		value,
	)
}

func (env *GoJNIEnv) NewString(s u16string) jstring {
	if env == nil || env.ptr == nil {
		return jnil.String
	}

	var chars *jchar
	if len(s) > 0 {
		chars = (*jchar)(unsafe.Pointer(&s[0]))
	}

	return jstring(C.gojni_NewString(
		env.ptr,
		chars,
		jsize(len(s)),
	))
}

func (env *GoJNIEnv) GetStringLength(str jstring) jsize {
	if env == nil || env.ptr == nil {
		return 0
	}

	return jsize(C.gojni_GetStringLength(
		env.ptr,
		str,
	))
}

func (env *GoJNIEnv) GetStringChars(str jstring,isCopy *jboolean,) u16string {
	if env == nil || env.ptr == nil {
		return nil
	}

	chars := C.gojni_GetStringChars(
		env.ptr,
		str,
		isCopy,
	)
	if chars == nil {
		return nil
	}
	defer C.gojni_ReleaseStringChars(env.ptr, str, chars)

	length := C.gojni_GetStringLength(env.ptr, str)
	if length < 0 {
		return nil
	}
	if length == 0 {
		return u16string{}
	}

	src := unsafe.Slice((*C.jchar)(unsafe.Pointer(chars)), int(length))
	out := make(u16string, int(length))
	copy(out, src)

	return out
}

func (env *GoJNIEnv) NewStringUTF(s string) jstring {
	if env == nil || env.ptr == nil {
		return jnil.String
	}

	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))

	return jstring(C.gojni_NewStringUTF(
		env.ptr,
		cs,
	))
}

func (env *GoJNIEnv) GetStringUTFLength(str jstring) jsize {
	if env == nil || env.ptr == nil {
		return 0
	}

	return jsize(C.gojni_GetStringUTFLength(
		env.ptr,
		str,
	))
}

func (env *GoJNIEnv) GetStringUTFChars(str jstring,isCopy *jboolean,) string {
	if env == nil || env.ptr == nil {
		return ""
	}

	cstr := C.gojni_GetStringUTFChars(env.ptr, str, isCopy)
	if cstr == nil {
		return ""
	}
	defer C.gojni_ReleaseStringUTFChars(env.ptr, str, cstr)

	return C.GoString(cstr)
}

func (env *GoJNIEnv) GetArrayLength(array jarray) jsize {
	if env == nil || env.ptr == nil {
		return 0
	}

	return jsize(C.gojni_GetArrayLength(
		env.ptr,
		array,
	))
}

func (env *GoJNIEnv) NewObjectArray(leng jsize,clazz jclass,init jobject) jobjectArray {
	if env == nil || env.ptr == nil {
		return jnil.ObjectArray
	}
	return jobjectArray(C.gojni_NewObjectArray(
		env.ptr,
		leng,
		clazz,
		init,
	))
}

func (env *GoJNIEnv) GetObjectArrayElement(array jobjectArray,index jsize) jobject {
	if env == nil || env.ptr == nil {
		return jnil.Object
	}
	return jobject(C.gojni_GetObjectArrayElement(
		env.ptr,
		array,
		index,
	))
}

func (env *GoJNIEnv) SetObjectArrayElement(array jobjectArray,index jsize,val jobject) {
	if env == nil || env.ptr == nil {
		return
	}
	C.gojni_SetObjectArrayElement(
		env.ptr,
		array,
		index,
		val,
	)
}

func (env *GoJNIEnv) NewBooleanArray(leng jsize,) jbooleanArray {
	if env == nil || env.ptr == nil {
		return jnil.BooleanArray
	}

	return jbooleanArray(C.gojni_NewBooleanArray(
		env.ptr,
		leng,
	))
}

func (env *GoJNIEnv) NewByteArray(leng jsize,) jbyteArray {
	if env == nil || env.ptr == nil {
		return jnil.ByteArray
	}

	return jbyteArray(C.gojni_NewByteArray(
		env.ptr,
		leng,
	))
}
func (env *GoJNIEnv) NewCharArray(leng jsize,) jcharArray {
	if env == nil || env.ptr == nil {
		return jnil.CharArray
	}

	return jcharArray(C.gojni_NewCharArray(
		env.ptr,
		leng,
	))
}

func (env *GoJNIEnv) NewShortArray(leng jsize,) jshortArray {
	if env == nil || env.ptr == nil {
		return jnil.ShortArray
	}

	return jshortArray(C.gojni_NewShortArray(
		env.ptr,
		leng,
	))
}

func (env *GoJNIEnv) NewIntArray(leng jsize,) jintArray {
	if env == nil || env.ptr == nil {
		return jnil.IntArray
	}

	return jintArray(C.gojni_NewIntArray(
		env.ptr,
		leng,
	))
}

func (env *GoJNIEnv) NewLongArray(leng jsize,) jlongArray {
	if env == nil || env.ptr == nil {
		return jnil.LongArray
	}

	return jlongArray(C.gojni_NewLongArray(
		env.ptr,
		leng,
	))
}

func (env *GoJNIEnv) NewFloatArray(leng jsize,) jfloatArray {
	if env == nil || env.ptr == nil {
		return jnil.FloatArray
	}

	return jfloatArray(C.gojni_NewFloatArray(
		env.ptr,
		leng,
	))
}

func (env *GoJNIEnv) NewDoubleArray(leng jsize,) jdoubleArray {
	if env == nil || env.ptr == nil {
		return jnil.DoubleArray
	}

	return jdoubleArray(C.gojni_NewDoubleArray(
		env.ptr,
		leng,
	))
}

func (env *GoJNIEnv) GetBooleanArrayElements(
	array jbooleanArray,
	isCopy *jboolean,
) []jboolean {
	if env == nil || env.ptr == nil {
		return nil
	}

	length := env.GetArrayLength(jarray(array))
	if length < 0 {
		return nil
	}

	var cIsCopy jboolean
	var pIsCopy *jboolean
	if isCopy != nil {
		pIsCopy = &cIsCopy
	}

	elems := C.gojni_GetBooleanArrayElements(env.ptr, array, pIsCopy)
	if elems == nil {
		return nil
	}

	if isCopy != nil {
		*isCopy = jboolean(cIsCopy)
	}

	// Always release the JNI elements before returning.
	defer C.gojni_ReleaseBooleanArrayElements(
		env.ptr,
		array,
		elems,
		JNI_ABORT,
	)

	// Copy JNI memory into Go-owned memory.
	src := unsafe.Slice((*jboolean)(unsafe.Pointer(elems)), int(length))
	result := append([]jboolean(nil), src...)

	return result
}

func (env *GoJNIEnv) ReleaseBooleanArrayElements(
	array jbooleanArray,
	elems []jboolean,
	mode jint,
) {
	if env == nil || env.ptr == nil {
		return
	}

	if mode == JNI_ABORT {
		return
	}

	if mode != 0 && mode != JNI_COMMIT {
		return
	}

	if len(elems) == 0 {
		return
	}

	length := env.GetArrayLength(jarray(array))
	if length < 0 || len(elems) > int(length) {
		return
	}

	C.gojni_SetBooleanArrayRegion(
		env.ptr,
		array,
		0,
		jsize(len(elems)),
		(*jboolean)(unsafe.Pointer(&elems[0])),
	)
}

func (env *GoJNIEnv) GetByteArrayElements(
	array jbyteArray,
	isCopy *jboolean,
) []jbyte {
	if env == nil || env.ptr == nil {
		return nil
	}

	length := env.GetArrayLength(jarray(array))
	if length < 0 {
		return nil
	}

	var cIsCopy jboolean
	var pIsCopy *jboolean
	if isCopy != nil {
		pIsCopy = &cIsCopy
	}

	elems := C.gojni_GetByteArrayElements(env.ptr, array, pIsCopy)
	if elems == nil {
		return nil
	}

	if isCopy != nil {
		*isCopy = jboolean(cIsCopy)
	}

	defer C.gojni_ReleaseByteArrayElements(
		env.ptr,
		array,
		elems,
		JNI_ABORT,
	)

	if length == 0 {
		return []jbyte{}
	}

	src := unsafe.Slice((*jbyte)(unsafe.Pointer(elems)), int(length))
	result := make([]jbyte, int(length))
	copy(result, src)

	return result
}

func (env *GoJNIEnv) ReleaseByteArrayElements(
	array jbyteArray,
	elems []jbyte,
	mode jint,
) {
	if env == nil || env.ptr == nil {
		return
	}

	if mode == JNI_ABORT {
		return
	}

	if mode != 0 && mode != JNI_COMMIT {
		return
	}

	if len(elems) == 0 {
		return
	}

	length := env.GetArrayLength(jarray(array))
	if length < 0 || len(elems) > int(length) {
		return
	}

	C.gojni_SetByteArrayRegion(
		env.ptr,
		array,
		0,
		jsize(len(elems)),
		(*jbyte)(unsafe.Pointer(&elems[0])),
	)
}

func (env *GoJNIEnv) GetCharArrayElements(
	array jcharArray,
	isCopy *jboolean,
) []jchar {
	if env == nil || env.ptr == nil {
		return nil
	}

	length := env.GetArrayLength(jarray(array))
	if length < 0 {
		return nil
	}

	var cIsCopy jboolean
	var pIsCopy *jboolean
	if isCopy != nil {
		pIsCopy = &cIsCopy
	}

	elems := C.gojni_GetCharArrayElements(env.ptr, array, pIsCopy)
	if elems == nil {
		return nil
	}

	if isCopy != nil {
		*isCopy = jboolean(cIsCopy)
	}

	defer C.gojni_ReleaseCharArrayElements(
		env.ptr,
		array,
		elems,
		JNI_ABORT,
	)

	if length == 0 {
		return []jchar{}
	}

	src := unsafe.Slice((*jchar)(unsafe.Pointer(elems)), int(length))
	result := make([]jchar, int(length))
	copy(result, src)

	return result
}

func (env *GoJNIEnv) ReleaseCharArrayElements(
	array jcharArray,
	elems []jchar,
	mode jint,
) {
	if env == nil || env.ptr == nil {
		return
	}

	if mode == JNI_ABORT {
		return
	}

	if mode != 0 && mode != JNI_COMMIT {
		return
	}

	if len(elems) == 0 {
		return
	}

	length := env.GetArrayLength(jarray(array))
	if length < 0 || len(elems) > int(length) {
		return
	}

	C.gojni_SetCharArrayRegion(
		env.ptr,
		array,
		0,
		jsize(len(elems)),
		(*jchar)(unsafe.Pointer(&elems[0])),
	)
}

func (env *GoJNIEnv) GetShortArrayElements(
	array jshortArray,
	isCopy *jboolean,
) []jshort {
	if env == nil || env.ptr == nil {
		return nil
	}

	length := env.GetArrayLength(jarray(array))
	if length < 0 {
		return nil
	}

	var cIsCopy jboolean
	var pIsCopy *jboolean
	if isCopy != nil {
		pIsCopy = &cIsCopy
	}

	elems := C.gojni_GetShortArrayElements(env.ptr, array, pIsCopy)
	if elems == nil {
		return nil
	}

	if isCopy != nil {
		*isCopy = jboolean(cIsCopy)
	}

	defer C.gojni_ReleaseShortArrayElements(
		env.ptr,
		array,
		elems,
		JNI_ABORT,
	)

	if length == 0 {
		return []jshort{}
	}

	src := unsafe.Slice((*jshort)(unsafe.Pointer(elems)), int(length))
	result := make([]jshort, int(length))
	copy(result, src)

	return result
}

func (env *GoJNIEnv) ReleaseShortArrayElements(
	array jshortArray,
	elems []jshort,
	mode jint,
) {
	if env == nil || env.ptr == nil {
		return
	}

	if mode == JNI_ABORT {
		return
	}

	if mode != 0 && mode != JNI_COMMIT {
		return
	}

	if len(elems) == 0 {
		return
	}

	length := env.GetArrayLength(jarray(array))
	if length < 0 || len(elems) > int(length) {
		return
	}

	C.gojni_SetShortArrayRegion(
		env.ptr,
		array,
		0,
		jsize(len(elems)),
		(*jshort)(unsafe.Pointer(&elems[0])),
	)
}

func (env *GoJNIEnv) GetIntArrayElements(
	array jintArray,
	isCopy *jboolean,
) []jint {
	if env == nil || env.ptr == nil {
		return nil
	}

	length := env.GetArrayLength(jarray(array))
	if length < 0 {
		return nil
	}

	var cIsCopy jboolean
	var pIsCopy *jboolean
	if isCopy != nil {
		pIsCopy = &cIsCopy
	}

	elems := C.gojni_GetIntArrayElements(env.ptr, array, pIsCopy)
	if elems == nil {
		return nil
	}

	if isCopy != nil {
		*isCopy = jboolean(cIsCopy)
	}

	defer C.gojni_ReleaseIntArrayElements(
		env.ptr,
		array,
		elems,
		JNI_ABORT,
	)

	if length == 0 {
		return []jint{}
	}

	src := unsafe.Slice((*jint)(unsafe.Pointer(elems)), int(length))
	result := make([]jint, int(length))
	copy(result, src)

	return result
}

func (env *GoJNIEnv) ReleaseIntArrayElements(
	array jintArray,
	elems []jint,
	mode jint,
) {
	if env == nil || env.ptr == nil {
		return
	}

	if mode == JNI_ABORT {
		return
	}

	if mode != 0 && mode != JNI_COMMIT {
		return
	}

	if len(elems) == 0 {
		return
	}

	length := env.GetArrayLength(jarray(array))
	if length < 0 || len(elems) > int(length) {
		return
	}

	C.gojni_SetIntArrayRegion(
		env.ptr,
		array,
		0,
		jsize(len(elems)),
		(*C.jint)(unsafe.Pointer(&elems[0])),
	)
}

func (env *GoJNIEnv) GetLongArrayElements(
	array jlongArray,
	isCopy *jboolean,
) []jlong {
	if env == nil || env.ptr == nil {
		return nil
	}

	length := env.GetArrayLength(jarray(array))
	if length < 0 {
		return nil
	}

	var cIsCopy jboolean
	var pIsCopy *jboolean
	if isCopy != nil {
		pIsCopy = &cIsCopy
	}

	elems := C.gojni_GetLongArrayElements(env.ptr, array, pIsCopy)
	if elems == nil {
		return nil
	}

	if isCopy != nil {
		*isCopy = jboolean(cIsCopy)
	}

	defer C.gojni_ReleaseLongArrayElements(
		env.ptr,
		array,
		elems,
		JNI_ABORT,
	)

	if length == 0 {
		return []jlong{}
	}

	src := unsafe.Slice((*jlong)(unsafe.Pointer(elems)), int(length))
	result := make([]jlong, int(length))
	copy(result, src)

	return result
}

func (env *GoJNIEnv) ReleaseLongArrayElements(
	array jlongArray,
	elems []jlong,
	mode jint,
) {
	if env == nil || env.ptr == nil {
		return
	}

	if mode == JNI_ABORT {
		return
	}

	if mode != 0 && mode != JNI_COMMIT {
		return
	}

	if len(elems) == 0 {
		return
	}

	length := env.GetArrayLength(jarray(array))
	if length < 0 || len(elems) > int(length) {
		return
	}

	C.gojni_SetLongArrayRegion(
		env.ptr,
		array,
		0,
		jsize(len(elems)),
		(*jlong)(unsafe.Pointer(&elems[0])),
	)
}

func (env *GoJNIEnv) GetFloatArrayElements(
	array jfloatArray,
	isCopy *jboolean,
) []jfloat {
	if env == nil || env.ptr == nil {
		return nil
	}

	length := env.GetArrayLength(jarray(array))
	if length < 0 {
		return nil
	}

	var cIsCopy jboolean
	var pIsCopy *jboolean
	if isCopy != nil {
		pIsCopy = &cIsCopy
	}

	elems := C.gojni_GetFloatArrayElements(env.ptr, array, pIsCopy)
	if elems == nil {
		return nil
	}

	if isCopy != nil {
		*isCopy = jboolean(cIsCopy)
	}

	defer C.gojni_ReleaseFloatArrayElements(
		env.ptr,
		array,
		elems,
		JNI_ABORT,
	)

	if length == 0 {
		return []jfloat{}
	}

	src := unsafe.Slice((*jfloat)(unsafe.Pointer(elems)), int(length))
	result := make([]jfloat, int(length))
	copy(result, src)

	return result
}

func (env *GoJNIEnv) ReleaseFloatArrayElements(
	array jfloatArray,
	elems []jfloat,
	mode jint,
) {
	if env == nil || env.ptr == nil {
		return
	}

	if mode == JNI_ABORT {
		return
	}

	if mode != 0 && mode != JNI_COMMIT {
		return
	}

	if len(elems) == 0 {
		return
	}

	length := env.GetArrayLength(jarray(array))
	if length < 0 || len(elems) > int(length) {
		return
	}

	C.gojni_SetFloatArrayRegion(
		env.ptr,
		array,
		0,
		jsize(len(elems)),
		(*jfloat)(unsafe.Pointer(&elems[0])),
	)
}

func (env *GoJNIEnv) GetDoubleArrayElements(
	array jdoubleArray,
	isCopy *jboolean,
) []jdouble {
	if env == nil || env.ptr == nil {
		return nil
	}

	length := env.GetArrayLength(jarray(array))
	if length < 0 {
		return nil
	}

	var cIsCopy jboolean
	var pIsCopy *jboolean
	if isCopy != nil {
		pIsCopy = &cIsCopy
	}

	elems := C.gojni_GetDoubleArrayElements(env.ptr, array, pIsCopy)
	if elems == nil {
		return nil
	}

	if isCopy != nil {
		*isCopy = jboolean(cIsCopy)
	}

	defer C.gojni_ReleaseDoubleArrayElements(
		env.ptr,
		array,
		elems,
		JNI_ABORT,
	)

	if length == 0 {
		return []jdouble{}
	}

	src := unsafe.Slice((*jdouble)(unsafe.Pointer(elems)), int(length))
	result := make([]jdouble, int(length))
	copy(result, src)

	return result
}

func (env *GoJNIEnv) ReleaseDoubleArrayElements(
	array jdoubleArray,
	elems []jdouble,
	mode jint,
) {
	if env == nil || env.ptr == nil {
		return
	}

	if mode == JNI_ABORT {
		return
	}

	if mode != 0 && mode != JNI_COMMIT {
		return
	}

	if len(elems) == 0 {
		return
	}

	length := env.GetArrayLength(jarray(array))
	if length < 0 || len(elems) > int(length) {
		return
	}

	C.gojni_SetDoubleArrayRegion(
		env.ptr,
		array,
		0,
		jsize(len(elems)),
		(*jdouble)(unsafe.Pointer(&elems[0])),
	)
}

func (env *GoJNIEnv) GetBooleanArrayRegion(
	array jbooleanArray,
	start jsize, length jsize,
) []jboolean {
	if env == nil || env.ptr == nil {
		return nil
	}
	if start < 0 || length < 0 {
		return nil
	}

	result := make([]jboolean, int(length))
	if length == 0 {
		return result
	}

	C.gojni_GetBooleanArrayRegion(
		env.ptr, array, start, length,
		(*C.jboolean)(unsafe.Pointer(&result[0])),
	)
	return result
}

func (env *GoJNIEnv) GetByteArrayRegion(
	array jbyteArray,
	start jsize, length jsize,
) []jbyte {
	if env == nil || env.ptr == nil {
		return nil
	}
	if start < 0 || length < 0 {
		return nil
	}

	result := make([]jbyte, int(length))
	if length == 0 {
		return result
	}

	C.gojni_GetByteArrayRegion(
		env.ptr, array, start, length,
		(*C.jbyte)(unsafe.Pointer(&result[0])),
	)
	return result
}

func (env *GoJNIEnv) GetCharArrayRegion(
	array jcharArray,
	start jsize, length jsize,
) []jchar {
	if env == nil || env.ptr == nil {
		return nil
	}
	if start < 0 || length < 0 {
		return nil
	}

	result := make([]jchar, int(length))
	if length == 0 {
		return result
	}

	C.gojni_GetCharArrayRegion(
		env.ptr, array, start, length,
		(*C.jchar)(unsafe.Pointer(&result[0])),
	)
	return result
}

func (env *GoJNIEnv) GetShortArrayRegion(
	array jshortArray,
	start jsize, length jsize,
) []jshort {
	if env == nil || env.ptr == nil {
		return nil
	}
	if start < 0 || length < 0 {
		return nil
	}

	result := make([]jshort, int(length))
	if length == 0 {
		return result
	}

	C.gojni_GetShortArrayRegion(
		env.ptr, array, start, length,
		(*C.jshort)(unsafe.Pointer(&result[0])),
	)
	return result
}

func (env *GoJNIEnv) GetIntArrayRegion(
	array jintArray,
	start jsize, length jsize,
) []jint {
	if env == nil || env.ptr == nil {
		return nil
	}
	if start < 0 || length < 0 {
		return nil
	}

	result := make([]jint, int(length))
	if length == 0 {
		return result
	}

	C.gojni_GetIntArrayRegion(
		env.ptr, array, start, length,
		(*C.jint)(unsafe.Pointer(&result[0])),
	)
	return result
}

func (env *GoJNIEnv) GetLongArrayRegion(
	array jlongArray,
	start jsize, length jsize,
) []jlong {
	if env == nil || env.ptr == nil {
		return nil
	}
	if start < 0 || length < 0 {
		return nil
	}

	result := make([]jlong, int(length))
	if length == 0 {
		return result
	}

	C.gojni_GetLongArrayRegion(
		env.ptr, array, start, length,
		(*C.jlong)(unsafe.Pointer(&result[0])),
	)
	return result
}

func (env *GoJNIEnv) GetFloatArrayRegion(
	array jfloatArray,
	start jsize, length jsize,
) []jfloat {
	if env == nil || env.ptr == nil {
		return nil
	}
	if start < 0 || length < 0 {
		return nil
	}

	result := make([]jfloat, int(length))
	if length == 0 {
		return result
	}

	C.gojni_GetFloatArrayRegion(
		env.ptr, array, start, length,
		(*C.jfloat)(unsafe.Pointer(&result[0])),
	)
	return result
}

func (env *GoJNIEnv) GetDoubleArrayRegion(
	array jdoubleArray,
	start jsize, length jsize,
) []jdouble {
	if env == nil || env.ptr == nil {
		return nil
	}
	if start < 0 || length < 0 {
		return nil
	}

	result := make([]jdouble, int(length))
	if length == 0 {
		return result
	}

	C.gojni_GetDoubleArrayRegion(
		env.ptr, array, start, length,
		(*C.jdouble)(unsafe.Pointer(&result[0])),
	)
	return result
}

func (env *GoJNIEnv) SetBooleanArrayRegion(
	array jbooleanArray,
	start jsize,
	values []jboolean,
) jint {
	if env == nil || env.ptr == nil {
		return JNI_EINVAL
	}
	if start < 0 || len(values) > int(^uint32(0)>>1)-int(start) {
		return JNI_EINVAL
	}
	if len(values) == 0 {
		return JNI_OK
	}

	C.gojni_SetBooleanArrayRegion(
		env.ptr, array, start, jsize(len(values)),
		(*C.jboolean)(unsafe.Pointer(&values[0])),
	)
	return JNI_OK
}

func (env *GoJNIEnv) SetByteArrayRegion(
	array jbyteArray,
	start jsize,
	values []jbyte,
) jint {
	if env == nil || env.ptr == nil {
		return JNI_EINVAL
	}
	if start < 0 || len(values) > int(^uint32(0)>>1)-int(start) {
		return JNI_EINVAL
	}
	if len(values) == 0 {
		return JNI_OK
	}

	C.gojni_SetByteArrayRegion(
		env.ptr, array, start, jsize(len(values)),
		(*C.jbyte)(unsafe.Pointer(&values[0])),
	)
	return JNI_OK
}

func (env *GoJNIEnv) SetCharArrayRegion(
	array jcharArray,
	start jsize,
	values []jchar,
) jint {
	if env == nil || env.ptr == nil {
		return JNI_EINVAL
	}
	if start < 0 || len(values) > int(^uint32(0)>>1)-int(start) {
		return JNI_EINVAL
	}
	if len(values) == 0 {
		return JNI_OK
	}

	C.gojni_SetCharArrayRegion(
		env.ptr, array, start, jsize(len(values)),
		(*C.jchar)(unsafe.Pointer(&values[0])),
	)
	return JNI_OK
}

func (env *GoJNIEnv) SetShortArrayRegion(
	array jshortArray,
	start jsize,
	values []jshort,
) jint {
	if env == nil || env.ptr == nil {
		return JNI_EINVAL
	}
	if start < 0 || len(values) > int(^uint32(0)>>1)-int(start) {
		return JNI_EINVAL
	}
	if len(values) == 0 {
		return JNI_OK
	}

	C.gojni_SetShortArrayRegion(
		env.ptr, array, start, jsize(len(values)),
		(*C.jshort)(unsafe.Pointer(&values[0])),
	)
	return JNI_OK
}

func (env *GoJNIEnv) SetIntArrayRegion(
	array jintArray,
	start jsize,
	values []jint,
) jint {
	if env == nil || env.ptr == nil {
		return JNI_EINVAL
	}
	if start < 0 || len(values) > int(^uint32(0)>>1)-int(start) {
		return JNI_EINVAL
	}
	if len(values) == 0 {
		return JNI_OK
	}

	C.gojni_SetIntArrayRegion(
		env.ptr, array, start, jsize(len(values)),
		(*C.jint)(unsafe.Pointer(&values[0])),
	)
	return JNI_OK
}

func (env *GoJNIEnv) SetLongArrayRegion(
	array jlongArray,
	start jsize,
	values []jlong,
) jint {
	if env == nil || env.ptr == nil {
		return JNI_EINVAL
	}
	if start < 0 || len(values) > int(^uint32(0)>>1)-int(start) {
		return JNI_EINVAL
	}
	if len(values) == 0 {
		return JNI_OK
	}

	C.gojni_SetLongArrayRegion(
		env.ptr, array, start, jsize(len(values)),
		(*C.jlong)(unsafe.Pointer(&values[0])),
	)
	return JNI_OK
}

func (env *GoJNIEnv) SetFloatArrayRegion(
	array jfloatArray,
	start jsize,
	values []jfloat,
) jint {
	if env == nil || env.ptr == nil {
		return JNI_EINVAL
	}
	if start < 0 || len(values) > int(^uint32(0)>>1)-int(start) {
		return JNI_EINVAL
	}
	if len(values) == 0 {
		return JNI_OK
	}

	C.gojni_SetFloatArrayRegion(
		env.ptr, array, start, jsize(len(values)),
		(*C.jfloat)(unsafe.Pointer(&values[0])),
	)
	return JNI_OK
}

func (env *GoJNIEnv) SetDoubleArrayRegion(
	array jdoubleArray,
	start jsize,
	values []jdouble,
) jint {
	if env == nil || env.ptr == nil {
		return JNI_EINVAL
	}
	if start < 0 || len(values) > int(^uint32(0)>>1)-int(start) {
		return JNI_EINVAL
	}
	if len(values) == 0 {
		return JNI_OK
	}

	C.gojni_SetDoubleArrayRegion(
		env.ptr, array, start, jsize(len(values)),
		(*C.jdouble)(unsafe.Pointer(&values[0])),
	)
	return JNI_OK
}

func (env *GoJNIEnv) RegisterNatives(
	clazz jclass,
	methods []JNINativeMethod,
) jint {
	if env == nil || env.ptr == nil {
		return JNI_EINVAL
	}

	if len(methods) == 0 {
		return C.gojni_RegisterNatives(env.ptr, clazz, nil, 0)
	}

	return C.gojni_RegisterNatives(
		env.ptr,
		clazz,
		(*C.JNINativeMethod)(unsafe.Pointer(&methods[0])),
				       C.jint(len(methods)),
	)
}

func (env *GoJNIEnv) UnregisterNatives(clazz jclass) jint {
	if env == nil || env.ptr == nil {
		return JNI_EINVAL
	}

	return C.gojni_UnregisterNatives(env.ptr, clazz)
}

func (env *GoJNIEnv) MonitorEnter(obj jobject) jint {
	if env == nil || env.ptr == nil {
		return JNI_EINVAL
	}

	return C.gojni_MonitorEnter(env.ptr, obj)
}

func (env *GoJNIEnv) MonitorExit(obj jobject) jint {
	if env == nil || env.ptr == nil {
		return JNI_EINVAL
	}

	return C.gojni_MonitorExit(env.ptr, obj)
}

func (env *GoJNIEnv) GetJavaVM() (*goJavaVM, jint) {
	if env == nil || env.ptr == nil {
		return nil, JNI_EINVAL
	}

	var vm *JavaVM
	rc := C.gojni_GetJavaVM(env.ptr, &vm)
	if rc != JNI_OK {
		return nil, jint(rc)
	}

	return WrapgoJavaVM((*JavaVM)(unsafe.Pointer(vm))), jint(rc)
}

func (env *GoJNIEnv) GetStringRegion(
	str jstring,
	start jsize, length jsize,
) u16string {
	if env == nil || env.ptr == nil || start < 0 || length < 0 {
		return nil
	}

	result := make(u16string, int(length))
	if length == 0 {
		return result
	}

	C.gojni_GetStringRegion(
		env.ptr,
		str,
		start,
		length,
		(*C.jchar)(unsafe.Pointer(&result[0])),
	)
	return result
}

func modifiedUTF8LengthFromUTF16(
	env *GoJNIEnv,
	str jstring,
	start, length jsize,
) int {
	if env == nil || env.ptr == nil || start < 0 || length < 0 {
		return -1
	}
	if length == 0 {
		return 0
	}

	units := make([]jchar, int(length))
	C.gojni_GetStringRegion(
		env.ptr,
		str,
		start,
		length,
		(*C.jchar)(unsafe.Pointer(&units[0])),
	)

	n := 0
	for _, c := range units {
		switch {
			case c == 0:
				// Modified UTF-8 encodes U+0000 as C0 80.
				n += 2
			case c <= 0x7F:
				n++
			case c <= 0x7FF:
				n += 2
			default:
				// Each UTF-16 code unit, including each surrogate,
				// takes three bytes in Modified UTF-8.
				n += 3
		}
	}

	return n
}

func (env *GoJNIEnv) GetStringUTFRegion(
	str jstring,
	start jsize,
	length jsize,
) string {
	if env == nil || env.ptr == nil || start < 0 || length < 0 {
		return ""
	}
	if length == 0 {
		return ""
	}

	// Maximum Modified UTF-8 size is 3 bytes per UTF-16 code unit, plus one
	// byte because the JVM always writes a terminating NUL after the data.
	buf := make([]byte, int(length)*3+1)

	C.gojni_GetStringUTFRegion(
		env.ptr,
		str,
		start,
		length,
		(*C.char)(unsafe.Pointer(&buf[0])),
	)

	// GetStringUTFRegion does not report the number of bytes written.
	// Determine the actual encoded length before returning a Go string.
	n := modifiedUTF8LengthFromUTF16(env, str, start, length)
	if n < 0 || n > len(buf) {
		return ""
	}

	return string(buf[:n])
}


func (vm *goJavaVM) GetEnv(version jint) (jint, *GoJNIEnv) {
	if vm == nil || vm.ptr == nil {
		return JNI_ERR, nil
	}

	var env unsafe.Pointer

	result := C.gojni_GetEnv(
		vm.ptr,
		&env,
		C.jint(version),
	)

	return jint(result), WrapGoJNIEnv((*JNIEnv)(env))
}

func (env *GoJNIEnv) GetPrimitiveArrayCritical(
	array jarray,
	isCopy *jboolean,
) unsafe.Pointer {
	if env == nil || env.ptr == nil {
		return nil
	}

	var cIsCopy C.jboolean
	var pIsCopy *C.jboolean
	if isCopy != nil {
		pIsCopy = &cIsCopy
	}

	ptr := C.gojni_GetPrimitiveArrayCritical(env.ptr, array, pIsCopy)
	if ptr == nil {
		return nil
	}

	if isCopy != nil {
		*isCopy = jboolean(cIsCopy)
	}

	return ptr
}

func (env *GoJNIEnv) ReleasePrimitiveArrayCritical(
	array jarray,
	carray unsafe.Pointer,
	mode jint,
) {
	if env == nil || env.ptr == nil || carray == nil {
		return
	}

	C.gojni_ReleasePrimitiveArrayCritical(
		env.ptr,
		array,
		carray,
		mode,
	)
}

func (env *GoJNIEnv) GetStringCritical(
	str jstring,
	isCopy *jboolean,
) *jchar {
	if env == nil || env.ptr == nil {
		return nil
	}

	var cIsCopy C.jboolean
	var pIsCopy *C.jboolean
	if isCopy != nil {
		pIsCopy = &cIsCopy
	}

	ptr := C.gojni_GetStringCritical(env.ptr, str, pIsCopy)
	if ptr == nil {
		return nil
	}

	if isCopy != nil {
		*isCopy = jboolean(cIsCopy)
	}

	return (*jchar)(unsafe.Pointer(ptr))
}

func (env *GoJNIEnv) ReleaseStringCritical(
	str jstring,
	cstring *jchar,
) {
	if env == nil || env.ptr == nil || cstring == nil {
		return
	}

	C.gojni_ReleaseStringCritical(
		env.ptr,
		str,
		(*C.jchar)(unsafe.Pointer(cstring)),
	)
}

func (vm *goJavaVM) AttachCurrentThread(args unsafe.Pointer) (jint, *GoJNIEnv) {
	if vm == nil || vm.ptr == nil {
		return JNI_ERR, nil
	}

	var env unsafe.Pointer

	result := C.gojni_AttachCurrentThread(
		vm.ptr,
		&env,
		args,
	)

	return jint(result), WrapGoJNIEnv((*JNIEnv)(env))
}

func (env *GoJNIEnv) NewWeakGlobalRef(obj jobject) jweak {
	if env == nil || env.ptr == nil {
		return jnil.Weak
	}
	return C.gojni_NewWeakGlobalRef(env.ptr, obj)
}

func (env *GoJNIEnv) DeleteWeakGlobalRef(ref jweak) {
	if env == nil || env.ptr == nil {
		return
	}
	C.gojni_DeleteWeakGlobalRef(env.ptr, ref)
}

func (env *GoJNIEnv) ExceptionCheck() jboolean {
	if env == nil || env.ptr == nil {
		return jboolean(0)
	}
	return jboolean(C.gojni_ExceptionCheck(env.ptr))
}

func (env *GoJNIEnv) NewDirectByteBuffer(
	address unsafe.Pointer,
	capacity jlong,
) jobject {
	if env == nil || env.ptr == nil {
		return jnil.Object
	}
	return C.gojni_NewDirectByteBuffer(
		env.ptr,
		address,
		capacity,
	)
}

func (env *GoJNIEnv) GetDirectBufferAddress(buf jobject) unsafe.Pointer {
	if env == nil || env.ptr == nil {
		return nil
	}
	return C.gojni_GetDirectBufferAddress(env.ptr, buf)
}

func (env *GoJNIEnv) GetDirectBufferCapacity(buf jobject) jlong {
	if env == nil || env.ptr == nil {
		return jlong(-1)
	}
	return jlong(C.gojni_GetDirectBufferCapacity(env.ptr, buf))
}

func (env *GoJNIEnv) GetObjectRefType(obj jobject) jobjectRefType {
	if env == nil || env.ptr == nil {
		return jobjectRefType(0)
	}
	return jobjectRefType(C.gojni_GetObjectRefType(env.ptr, obj))
}

func (env *GoJNIEnv) GetModule(clazz jclass) jobject {
	if env == nil || env.ptr == nil {
		return jnil.Object
	}
	return C.gojni_GetModule(env.ptr, clazz)
}

func (env *GoJNIEnv) IsVirtualThread(obj jobject) jboolean {
	if env == nil || env.ptr == nil {
		return jboolean(0)
	}
	return jboolean(C.gojni_IsVirtualThread(env.ptr, obj))
}

func (vm *goJavaVM) AttachCurrentThreadAsDaemon(args unsafe.Pointer) (jint, *GoJNIEnv) {
	if vm == nil || vm.ptr == nil {
		return JNI_ERR, nil
	}

	var env unsafe.Pointer

	result := C.gojni_AttachCurrentThreadAsDaemon(
		vm.ptr,
		&env,
		args,
	)

	return jint(result), WrapGoJNIEnv((*JNIEnv)(env))
}

func (vm *goJavaVM) DestroyJavaVM() jint {
	if vm == nil || vm.ptr == nil {
		return JNI_ERR
	}
	result := C.gojni_DestroyJavaVM(vm.ptr)
	return jint(result)
}

func (vm *goJavaVM) DetachCurrentThread() jint {
	if vm == nil || vm.ptr == nil {
		return JNI_ERR
	}
	result := C.gojni_DetachCurrentThread(vm.ptr)
	return jint(result)
}