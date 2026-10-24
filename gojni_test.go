package gojni

import (
	"reflect"
	"testing"
)

func TestU16StringRoundTrip(t *testing.T) {
	tests := []string{
		"",
		"hello",
		"こんにちは",
		"Hello, 世界 🌍",
		"😀",
		"\x00",
		"e\u0301",
	}

	for _, input := range tests {
		t.Run(input, func(t *testing.T) {
			got := U16String(input).String()
			if got != input {
				t.Errorf("U16String(%q).String() = %q; want %q", input, got, input)
			}
		})
	}
}

func TestU16StringLength(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"", 0},
		{"A", 1},
		{"こんにちは", 5},
		{"😀", 2}, // A supplementary character uses two UTF-16 code units.
	}

	for _, tt := range tests {
		got := len(U16String(tt.input))
		if got != tt.want {
			t.Errorf("len(U16String(%q)) = %d; want %d", tt.input, got, tt.want)
		}
	}
}

func TestJNIConstants(t *testing.T) {
	if JNI_OK != 0 {
		t.Errorf("JNI_OK = %d; want 0", JNI_OK)
	}
	if JNI_ERR != -1 {
		t.Errorf("JNI_ERR = %d; want -1", JNI_ERR)
	}
	if JNIFalse != 0 {
		t.Errorf("JNIFalse = %d; want 0", JNIFalse)
	}
	if JNITrue != 1 {
		t.Errorf("JNITrue = %d; want 1", JNITrue)
	}
	if JNI_VERSION_1_8 != 0x00010008 {
		t.Errorf("JNI_VERSION_1_8 = %#x; want %#x", JNI_VERSION_1_8, 0x00010008)
	}
}

func TestWrapNilPointers(t *testing.T) {
	if WrapGoJNIEnv(nil) != nil {
		t.Error("WrapGoJNIEnv(nil) should return nil")
	}
	if WrapgoJavaVM(nil) != nil {
		t.Error("WrapgoJavaVM(nil) should return nil")
	}
}

func TestNilJNIEnvMethods(t *testing.T) {
	var env *GoJNIEnv

	if got := env.GetVersion(); got != JNI_ERR {
		t.Errorf("GetVersion() = %d; want %d", got, JNI_ERR)
	}
	if got := env.FindClass("java/lang/Object"); got != jnil.Class {
		t.Errorf("FindClass() on nil environment = %v; want nil", got)
	}
	if got := env.GetMethodID(jnil.Class, "toString", "()Ljava/lang/String;"); got != nil {
		t.Errorf("GetMethodID() on nil environment = %v; want nil", got)
	}
	if got := env.NewString(U16String("test")); got != jnil.String {
		t.Errorf("NewString() on nil environment = %v; want nil", got)
	}
	if got := env.GetStringUTFChars(jnil.String, nil); got != "" {
		t.Errorf("GetStringUTFChars() on nil environment = %q; want empty string", got)
	}
	if got := env.GetIntArrayElements(jnil.IntArray, nil, JNI_ABORT); got != nil {
		t.Errorf("GetIntArrayElements() on nil environment = %v; want nil", got)
	}
	if got := env.GetStringRegion(jnil.String, 0, 1); got != nil {
		t.Errorf("GetStringRegion() on nil environment = %v; want nil", got)
	}
}

func TestNilJavaVMMethods(t *testing.T) {
	var vm *goJavaVM

	rc, env := vm.GetEnv(JNI_VERSION_1_8)
	if rc != JNI_ERR || env != nil {
		t.Errorf("GetEnv() = (%d, %v); want (%d, nil)", rc, env, JNI_ERR)
	}

	rc, env = vm.AttachCurrentThread(nil)
	if rc != JNI_ERR || env != nil {
		t.Errorf("AttachCurrentThread() = (%d, %v); want (%d, nil)", rc, env, JNI_ERR)
	}

	if got := vm.DestroyJavaVM(); got != JNI_ERR {
		t.Errorf("DestroyJavaVM() = %d; want %d", got, JNI_ERR)
	}
	if got := vm.DetachCurrentThread(); got != JNI_ERR {
		t.Errorf("DetachCurrentThread() = %d; want %d", got, JNI_ERR)
	}
}

func TestGoJNIValuesUnsupportedType(t *testing.T) {
	got := goJNIValues("unsupported")
	if got != nil {
		t.Errorf("goJNIValues() with unsupported argument = %v; want nil", got)
	}
}

func TestGoJNIValuesEmpty(t *testing.T) {
	got := goJNIValues()
	if !reflect.DeepEqual(got, []jvalue{}) {
		t.Errorf("goJNIValues() = %v; want empty slice", got)
	}
}
