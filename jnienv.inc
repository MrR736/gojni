JNIEXPORT jint JNICALL gojni_GetVersion(JNIEnv *env);

JNIEXPORT jclass JNICALL gojni_DefineClass(JNIEnv *env, const char *name, jobject loader, const jbyte *buf,jsize len);
JNIEXPORT jclass JNICALL gojni_FindClass(JNIEnv *env, const char *name);

JNIEXPORT jmethodID JNICALL gojni_FromReflectedMethod(JNIEnv *env, jobject method);
JNIEXPORT jfieldID JNICALL gojni_FromReflectedField(JNIEnv *env, jobject field);

JNIEXPORT jobject JNICALL gojni_ToReflectedMethod(JNIEnv *env, jclass cls, jmethodID methodID, jboolean isStatic);

JNIEXPORT jclass JNICALL gojni_GetSuperclass(JNIEnv *env, jclass sub);
JNIEXPORT jboolean JNICALL gojni_IsAssignableFrom(JNIEnv *env, jclass sub, jclass sup);

JNIEXPORT jobject JNICALL gojni_ToReflectedField(JNIEnv *env, jclass cls, jfieldID fieldID, jboolean isStatic);

JNIEXPORT jint JNICALL gojni_Throw(JNIEnv *env, jthrowable obj);
JNIEXPORT jint JNICALL gojni_ThrowNew(JNIEnv *env, jclass clazz, const char *msg);
JNIEXPORT jthrowable JNICALL gojni_ExceptionOccurred(JNIEnv *env);
JNIEXPORT void JNICALL gojni_ExceptionDescribe(JNIEnv *env);
JNIEXPORT void JNICALL gojni_ExceptionClear(JNIEnv *env);
JNIEXPORT void JNICALL gojni_FatalError(JNIEnv *env, const char *msg);

JNIEXPORT jint JNICALL gojni_PushLocalFrame(JNIEnv *env, jint capacity);
JNIEXPORT jobject JNICALL gojni_PopLocalFrame(JNIEnv *env, jobject result);

JNIEXPORT jobject JNICALL gojni_NewGlobalRef(JNIEnv *env, jobject lobj);
JNIEXPORT void JNICALL gojni_DeleteGlobalRef(JNIEnv *env, jobject gref);
JNIEXPORT void JNICALL gojni_DeleteLocalRef(JNIEnv *env, jobject obj);
JNIEXPORT jboolean JNICALL gojni_IsSameObject(JNIEnv *env, jobject obj1, jobject obj2);
JNIEXPORT jobject JNICALL gojni_NewLocalRef(JNIEnv *env, jobject ref);
JNIEXPORT jint JNICALL gojni_EnsureLocalCapacity(JNIEnv *env, jint capacity);

JNIEXPORT jobject JNICALL gojni_AllocObject(JNIEnv *env, jclass clazz);

JNIEXPORT jobject JNICALL gojni_NewObjectA(JNIEnv *env, jclass clazz, jmethodID methodID, const jvalue *args);

JNIEXPORT jclass JNICALL gojni_GetObjectClass(JNIEnv *env, jobject obj);
JNIEXPORT jboolean JNICALL gojni_IsInstanceOf(JNIEnv *env, jobject obj, jclass clazz);

JNIEXPORT jmethodID JNICALL gojni_GetMethodID(JNIEnv *env, jclass clazz, const char *name, const char *sig);

JNIEXPORT jobject JNICALL gojni_CallObjectMethodA(JNIEnv *env, jobject obj, jmethodID methodID, const jvalue * args);

JNIEXPORT jboolean JNICALL gojni_CallBooleanMethodA(JNIEnv *env, jobject obj, jmethodID methodID, const jvalue * args);

JNIEXPORT jbyte JNICALL gojni_CallByteMethodA(JNIEnv *env, jobject obj, jmethodID methodID, const jvalue *args);

JNIEXPORT jchar JNICALL gojni_CallCharMethodA(JNIEnv *env, jobject obj, jmethodID methodID, const jvalue *args);

JNIEXPORT jshort JNICALL gojni_CallShortMethodA(JNIEnv *env, jobject obj, jmethodID methodID, const jvalue *args);

JNIEXPORT jint JNICALL gojni_CallIntMethodA(JNIEnv *env, jobject obj, jmethodID methodID, const jvalue *args);

JNIEXPORT jlong JNICALL gojni_CallLongMethodA(JNIEnv *env, jobject obj, jmethodID methodID, const jvalue *args);

JNIEXPORT jfloat JNICALL gojni_CallFloatMethodA(JNIEnv *env, jobject obj, jmethodID methodID, const jvalue *args);

JNIEXPORT jdouble JNICALL gojni_CallDoubleMethodA(JNIEnv *env, jobject obj, jmethodID methodID, const jvalue *args);

JNIEXPORT void JNICALL gojni_CallVoidMethodA(JNIEnv *env, jobject obj, jmethodID methodID, const jvalue * args);

JNIEXPORT jobject JNICALL gojni_CallNonvirtualObjectMethodA(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID,const jvalue * args);

JNIEXPORT jboolean JNICALL gojni_CallNonvirtualBooleanMethodA(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID,const jvalue * args);

JNIEXPORT jbyte JNICALL gojni_CallNonvirtualByteMethodA(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID,const jvalue *args);

JNIEXPORT jchar JNICALL gojni_CallNonvirtualCharMethodA(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID,const jvalue *args);

JNIEXPORT jshort JNICALL gojni_CallNonvirtualShortMethodA(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID,const jvalue *args);

JNIEXPORT jint JNICALL gojni_CallNonvirtualIntMethodA(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID,const jvalue *args);

JNIEXPORT jlong JNICALL gojni_CallNonvirtualLongMethodA(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID,const jvalue *args);

JNIEXPORT jfloat JNICALL gojni_CallNonvirtualFloatMethodA(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID,const jvalue *args);

JNIEXPORT jdouble JNICALL gojni_CallNonvirtualDoubleMethodA(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID,const jvalue *args);

JNIEXPORT void JNICALL gojni_CallNonvirtualVoidMethodA(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID,const jvalue * args);

JNIEXPORT jfieldID JNICALL gojni_GetFieldID(JNIEnv *env, jclass clazz, const char *name, const char *sig);

JNIEXPORT jobject JNICALL gojni_GetObjectField(JNIEnv *env, jobject obj, jfieldID fieldID);
JNIEXPORT jboolean JNICALL gojni_GetBooleanField(JNIEnv *env, jobject obj, jfieldID fieldID);
JNIEXPORT jbyte JNICALL gojni_GetByteField(JNIEnv *env, jobject obj, jfieldID fieldID);
JNIEXPORT jchar JNICALL gojni_GetCharField(JNIEnv *env, jobject obj, jfieldID fieldID);
JNIEXPORT jshort JNICALL gojni_GetShortField(JNIEnv *env, jobject obj, jfieldID fieldID);
JNIEXPORT jint JNICALL gojni_GetIntField(JNIEnv *env, jobject obj, jfieldID fieldID);
JNIEXPORT jlong JNICALL gojni_GetLongField(JNIEnv *env, jobject obj, jfieldID fieldID);
JNIEXPORT jfloat JNICALL gojni_GetFloatField(JNIEnv *env, jobject obj, jfieldID fieldID);
JNIEXPORT jdouble JNICALL gojni_GetDoubleField(JNIEnv *env, jobject obj, jfieldID fieldID);

JNIEXPORT void JNICALL gojni_SetObjectField(JNIEnv *env, jobject obj, jfieldID fieldID, jobject val);
JNIEXPORT void JNICALL gojni_SetBooleanField(JNIEnv *env, jobject obj, jfieldID fieldID, jboolean val);
JNIEXPORT void JNICALL gojni_SetByteField(JNIEnv *env, jobject obj, jfieldID fieldID, jbyte val);
JNIEXPORT void JNICALL gojni_SetCharField(JNIEnv *env, jobject obj, jfieldID fieldID, jchar val);
JNIEXPORT void JNICALL gojni_SetShortField(JNIEnv *env, jobject obj, jfieldID fieldID, jshort val);
JNIEXPORT void JNICALL gojni_SetIntField(JNIEnv *env, jobject obj, jfieldID fieldID, jint val);
JNIEXPORT void JNICALL gojni_SetLongField(JNIEnv *env, jobject obj, jfieldID fieldID, jlong val);
JNIEXPORT void JNICALL gojni_SetFloatField(JNIEnv *env, jobject obj, jfieldID fieldID, jfloat val);
JNIEXPORT void JNICALL gojni_SetDoubleField(JNIEnv *env, jobject obj, jfieldID fieldID, jdouble val);

JNIEXPORT jmethodID JNICALL gojni_GetStaticMethodID(JNIEnv *env, jclass clazz, const char *name, const char *sig);

JNIEXPORT jobject JNICALL gojni_CallStaticObjectMethodA(JNIEnv *env, jclass clazz, jmethodID methodID, const jvalue *args);

JNIEXPORT jboolean JNICALL gojni_CallStaticBooleanMethodA(JNIEnv *env, jclass clazz, jmethodID methodID, const jvalue *args);

JNIEXPORT jbyte JNICALL gojni_CallStaticByteMethodA(JNIEnv *env, jclass clazz, jmethodID methodID, const jvalue *args);

JNIEXPORT jchar JNICALL gojni_CallStaticCharMethodA(JNIEnv *env, jclass clazz, jmethodID methodID, const jvalue *args);

JNIEXPORT jshort JNICALL gojni_CallStaticShortMethodA(JNIEnv *env, jclass clazz, jmethodID methodID, const jvalue *args);

JNIEXPORT jint JNICALL gojni_CallStaticIntMethodA(JNIEnv *env, jclass clazz, jmethodID methodID, const jvalue *args);

JNIEXPORT jlong JNICALL gojni_CallStaticLongMethodA(JNIEnv *env, jclass clazz, jmethodID methodID, const jvalue *args);

JNIEXPORT jfloat JNICALL gojni_CallStaticFloatMethodA(JNIEnv *env, jclass clazz, jmethodID methodID, const jvalue *args);

JNIEXPORT jdouble JNICALL gojni_CallStaticDoubleMethodA(JNIEnv *env, jclass clazz, jmethodID methodID, const jvalue *args);

JNIEXPORT void JNICALL gojni_CallStaticVoidMethodA(JNIEnv *env, jclass cls, jmethodID methodID, const jvalue * args);

JNIEXPORT jfieldID JNICALL gojni_GetStaticFieldID(JNIEnv *env, jclass clazz, const char *name, const char *sig);
JNIEXPORT jobject JNICALL gojni_GetStaticObjectField(JNIEnv *env, jclass clazz, jfieldID fieldID);
JNIEXPORT jboolean JNICALL gojni_GetStaticBooleanField(JNIEnv *env, jclass clazz, jfieldID fieldID);
JNIEXPORT jbyte JNICALL gojni_GetStaticByteField(JNIEnv *env, jclass clazz, jfieldID fieldID);
JNIEXPORT jchar JNICALL gojni_GetStaticCharField(JNIEnv *env, jclass clazz, jfieldID fieldID);
JNIEXPORT jshort JNICALL gojni_GetStaticShortField(JNIEnv *env, jclass clazz, jfieldID fieldID);
JNIEXPORT jint JNICALL gojni_GetStaticIntField(JNIEnv *env, jclass clazz, jfieldID fieldID);
JNIEXPORT jlong JNICALL gojni_GetStaticLongField(JNIEnv *env, jclass clazz, jfieldID fieldID);
JNIEXPORT jfloat JNICALL gojni_GetStaticFloatField(JNIEnv *env, jclass clazz, jfieldID fieldID);
JNIEXPORT jdouble JNICALL gojni_GetStaticDoubleField(JNIEnv *env, jclass clazz, jfieldID fieldID);

JNIEXPORT void JNICALL gojni_SetStaticObjectField(JNIEnv *env, jclass clazz, jfieldID fieldID, jobject value);
JNIEXPORT void JNICALL gojni_SetStaticBooleanField(JNIEnv *env, jclass clazz, jfieldID fieldID, jboolean value);
JNIEXPORT void JNICALL gojni_SetStaticByteField(JNIEnv *env, jclass clazz, jfieldID fieldID, jbyte value);
JNIEXPORT void JNICALL gojni_SetStaticCharField(JNIEnv *env, jclass clazz, jfieldID fieldID, jchar value);
JNIEXPORT void JNICALL gojni_SetStaticShortField(JNIEnv *env, jclass clazz, jfieldID fieldID, jshort value);
JNIEXPORT void JNICALL gojni_SetStaticIntField(JNIEnv *env, jclass clazz, jfieldID fieldID, jint value);
JNIEXPORT void JNICALL gojni_SetStaticLongField(JNIEnv *env, jclass clazz, jfieldID fieldID, jlong value);
JNIEXPORT void JNICALL gojni_SetStaticFloatField(JNIEnv *env, jclass clazz, jfieldID fieldID, jfloat value);
JNIEXPORT void JNICALL gojni_SetStaticDoubleField(JNIEnv *env, jclass clazz, jfieldID fieldID, jdouble value);

JNIEXPORT jstring JNICALL gojni_NewString(JNIEnv *env, const jchar *unicode, jsize len);
JNIEXPORT jsize JNICALL gojni_GetStringLength(JNIEnv *env, jstring str);
JNIEXPORT const jchar *JNICALL gojni_GetStringChars(JNIEnv *env, jstring str, jboolean *isCopy);
JNIEXPORT void JNICALL gojni_ReleaseStringChars(JNIEnv *env, jstring str, const jchar *chars);

JNIEXPORT jstring JNICALL gojni_NewStringUTF(JNIEnv *env, const char *utf);
JNIEXPORT jsize JNICALL gojni_GetStringUTFLength(JNIEnv *env, jstring str);
JNIEXPORT const char* JNICALL gojni_GetStringUTFChars(JNIEnv *env, jstring str, jboolean *isCopy);
JNIEXPORT void JNICALL gojni_ReleaseStringUTFChars(JNIEnv *env, jstring str, const char* chars);


JNIEXPORT jsize JNICALL gojni_GetArrayLength(JNIEnv *env, jarray array);

JNIEXPORT jobjectArray JNICALL gojni_NewObjectArray(JNIEnv *env, jsize len, jclass clazz, jobject init);
JNIEXPORT jobject JNICALL gojni_GetObjectArrayElement(JNIEnv *env, jobjectArray array, jsize index);
JNIEXPORT void JNICALL gojni_SetObjectArrayElement(JNIEnv *env, jobjectArray array, jsize index, jobject val);

JNIEXPORT jbooleanArray JNICALL gojni_NewBooleanArray(JNIEnv *env, jsize len);
JNIEXPORT jbyteArray JNICALL gojni_NewByteArray(JNIEnv *env, jsize len);
JNIEXPORT jcharArray JNICALL gojni_NewCharArray(JNIEnv *env, jsize len);
JNIEXPORT jshortArray JNICALL gojni_NewShortArray(JNIEnv *env, jsize len);
JNIEXPORT jintArray JNICALL gojni_NewIntArray(JNIEnv *env, jsize len);
JNIEXPORT jlongArray JNICALL gojni_NewLongArray(JNIEnv *env, jsize len);
JNIEXPORT jfloatArray JNICALL gojni_NewFloatArray(JNIEnv *env, jsize len);
JNIEXPORT jdoubleArray JNICALL gojni_NewDoubleArray(JNIEnv *env, jsize len);

JNIEXPORT jboolean * JNICALL gojni_GetBooleanArrayElements(JNIEnv *env, jbooleanArray array, jboolean *isCopy);
JNIEXPORT jbyte * JNICALL gojni_GetByteArrayElements(JNIEnv *env, jbyteArray array, jboolean *isCopy);
JNIEXPORT jchar * JNICALL gojni_GetCharArrayElements(JNIEnv *env, jcharArray array, jboolean *isCopy);
JNIEXPORT jshort * JNICALL gojni_GetShortArrayElements(JNIEnv *env, jshortArray array, jboolean *isCopy);
JNIEXPORT jint * JNICALL gojni_GetIntArrayElements(JNIEnv *env, jintArray array, jboolean *isCopy);
JNIEXPORT jlong * JNICALL gojni_GetLongArrayElements(JNIEnv *env, jlongArray array, jboolean *isCopy);
JNIEXPORT jfloat * JNICALL gojni_GetFloatArrayElements(JNIEnv *env, jfloatArray array, jboolean *isCopy);
JNIEXPORT jdouble * JNICALL gojni_GetDoubleArrayElements(JNIEnv *env, jdoubleArray array, jboolean *isCopy);

JNIEXPORT void JNICALL gojni_ReleaseBooleanArrayElements(JNIEnv *env, jbooleanArray array, jboolean *elems, jint mode);
JNIEXPORT void JNICALL gojni_ReleaseByteArrayElements(JNIEnv *env, jbyteArray array, jbyte *elems, jint mode);
JNIEXPORT void JNICALL gojni_ReleaseCharArrayElements(JNIEnv *env, jcharArray array, jchar *elems, jint mode);
JNIEXPORT void JNICALL gojni_ReleaseShortArrayElements(JNIEnv *env, jshortArray array, jshort *elems, jint mode);
JNIEXPORT void JNICALL gojni_ReleaseIntArrayElements(JNIEnv *env, jintArray array, jint *elems, jint mode);
JNIEXPORT void JNICALL gojni_ReleaseLongArrayElements(JNIEnv *env, jlongArray array, jlong *elems, jint mode);
JNIEXPORT void JNICALL gojni_ReleaseFloatArrayElements(JNIEnv *env, jfloatArray array, jfloat *elems, jint mode);
JNIEXPORT void JNICALL gojni_ReleaseDoubleArrayElements(JNIEnv *env, jdoubleArray array, jdouble *elems, jint mode);

JNIEXPORT void JNICALL gojni_GetBooleanArrayRegion(JNIEnv *env, jbooleanArray array, jsize start, jsize l, jboolean *buf);
JNIEXPORT void JNICALL gojni_GetByteArrayRegion(JNIEnv *env, jbyteArray array, jsize start, jsize len, jbyte *buf);
JNIEXPORT void JNICALL gojni_GetCharArrayRegion(JNIEnv *env, jcharArray array, jsize start, jsize len, jchar *buf);
JNIEXPORT void JNICALL gojni_GetShortArrayRegion(JNIEnv *env, jshortArray array, jsize start, jsize len, jshort *buf);
JNIEXPORT void JNICALL gojni_GetIntArrayRegion(JNIEnv *env, jintArray array, jsize start, jsize len, jint *buf);
JNIEXPORT void JNICALL gojni_GetLongArrayRegion(JNIEnv *env, jlongArray array, jsize start, jsize len, jlong *buf);
JNIEXPORT void JNICALL gojni_GetFloatArrayRegion(JNIEnv *env, jfloatArray array, jsize start, jsize len, jfloat *buf);
JNIEXPORT void JNICALL gojni_GetDoubleArrayRegion(JNIEnv *env, jdoubleArray array, jsize start, jsize len, jdouble *buf);

JNIEXPORT void JNICALL gojni_SetBooleanArrayRegion(JNIEnv *env, jbooleanArray array, jsize start, jsize l, const jboolean *buf);
JNIEXPORT void JNICALL gojni_SetByteArrayRegion(JNIEnv *env, jbyteArray array, jsize start, jsize len, const jbyte *buf);
JNIEXPORT void JNICALL gojni_SetCharArrayRegion(JNIEnv *env, jcharArray array, jsize start, jsize len, const jchar *buf);
JNIEXPORT void JNICALL gojni_SetShortArrayRegion(JNIEnv *env, jshortArray array, jsize start, jsize len, const jshort *buf);
JNIEXPORT void JNICALL gojni_SetIntArrayRegion(JNIEnv *env, jintArray array, jsize start, jsize len, const jint *buf);
JNIEXPORT void JNICALL gojni_SetLongArrayRegion(JNIEnv *env, jlongArray array, jsize start, jsize len, const jlong *buf);
JNIEXPORT void JNICALL gojni_SetFloatArrayRegion(JNIEnv *env, jfloatArray array, jsize start, jsize len, const jfloat *buf);
JNIEXPORT void JNICALL gojni_SetDoubleArrayRegion(JNIEnv *env, jdoubleArray array, jsize start, jsize len, const jdouble *buf);

JNIEXPORT jint JNICALL gojni_RegisterNatives(JNIEnv *env, jclass clazz, const JNINativeMethod *methods,jint nMethods);
JNIEXPORT jint JNICALL gojni_UnregisterNatives(JNIEnv *env, jclass clazz);

JNIEXPORT jint JNICALL gojni_MonitorEnter(JNIEnv *env, jobject obj);
JNIEXPORT jint JNICALL gojni_MonitorExit(JNIEnv *env, jobject obj);

JNIEXPORT jint JNICALL gojni_GetJavaVM(JNIEnv *env, JavaVM **vm);

JNIEXPORT void JNICALL gojni_GetStringRegion(JNIEnv *env, jstring str, jsize start, jsize len, jchar *buf);
JNIEXPORT void JNICALL gojni_GetStringUTFRegion(JNIEnv *env, jstring str, jsize start, jsize len, char *buf);

JNIEXPORT void * JNICALL gojni_GetPrimitiveArrayCritical(JNIEnv *env, jarray array, jboolean *isCopy);
JNIEXPORT void JNICALL gojni_ReleasePrimitiveArrayCritical(JNIEnv *env, jarray array, void *carray, jint mode);

JNIEXPORT const jchar * JNICALL gojni_GetStringCritical(JNIEnv *env, jstring string, jboolean *isCopy);
JNIEXPORT void JNICALL gojni_ReleaseStringCritical(JNIEnv *env, jstring string, const jchar *cstring);

JNIEXPORT jweak JNICALL gojni_NewWeakGlobalRef(JNIEnv *env, jobject obj);
JNIEXPORT void JNICALL gojni_DeleteWeakGlobalRef(JNIEnv *env, jweak ref);

JNIEXPORT jboolean JNICALL gojni_ExceptionCheck(JNIEnv *env);

JNIEXPORT jobject JNICALL gojni_NewDirectByteBuffer(JNIEnv* env, void* address, jlong capacity);
JNIEXPORT void* JNICALL gojni_GetDirectBufferAddress(JNIEnv* env, jobject buf);
JNIEXPORT jlong JNICALL gojni_GetDirectBufferCapacity(JNIEnv* env, jobject buf);
JNIEXPORT jobjectRefType JNICALL gojni_GetObjectRefType(JNIEnv* env, jobject obj);
JNIEXPORT jobject JNICALL gojni_GetModule(JNIEnv* env, jclass clazz);
JNIEXPORT jboolean JNICALL gojni_IsVirtualThread(JNIEnv* env, jobject obj);