# gojni

**A Go wrapper for the Java Native Interface (JNI) using cgo.**

`gojni` enables Go code to interact with a running Java Virtual Machine (JVM). It provides Go wrappers for JNI operations, including class and object access, Java method calls, fields, strings, arrays, exceptions, references, direct buffers, reflection, thread attachment, and native method registration.

## Features

- JNI environment and Java VM access
- Class lookup, inheritance checks, and type inspection
- Object creation and allocation
- Instance and static method calls
- Instance and static field access
- Java exception inspection, throwing, describing, and clearing
- Local, global, and weak references
- Java string conversion and UTF-16 access
- Primitive and object arrays
- Direct byte buffers
- Monitor enter and exit
- Reflection between Java members and JNI identifiers
- Java thread attachment and detachment
- Class definition and module access
- Native method registration and unregistration
- Go shared libraries callable from Java

## Requirements

- Go 1.24.4 or a compatible version supported by the project
- A JDK with JNI headers
- A C compiler supported by cgo
- Linux, or another platform with the appropriate JNI headers and build configuration

Check your toolchain:

```bash
go version
javac -version
gcc --version
```

Ensure cgo is enabled:

```bash
go env CGO_ENABLED
```

The output should be `1`.

## Installation

Use the module in your Go project:

```bash
go get <module-import-path>
```

Replace `<module-import-path>` with the repository's published Go module path.

For local development, an example project's `go.mod` can reference the local checkout:

```go
module example

go 1.24.4

require gojni v0.0.0

replace gojni => /path/to/gojni
```

Replace `/path/to/gojni` with the absolute path to your local repository.

## JNI header configuration

On Linux, the JNI headers are commonly located in the JDK's `include` directory and its platform-specific subdirectory.

For example:

```go
/*
#cgo CFLAGS: -I/usr/lib/jvm/default-java/include
#cgo CFLAGS: -I/usr/lib/jvm/default-java/include/linux
#include <jni.h>
*/
import "C"
```

Adjust these paths for your JDK installation. You can find the Java installation path with:

```bash
readlink -f "$(command -v javac)"
```

## Basic usage

The exact wrapper and type names depend on the exported API in your checkout. The following illustrates the general pattern:

```go
package main

/*
#include <jni.h>
*/
import "C"

import "gojni"

//export Java_com_example_MyClass_myfunc
func Java_com_example_MyClass_myfunc(
    env *C.JNIEnv,
    clazz C.jclass,
    value C.jint,
) C.jint {
    _ = clazz

    genv := gojni.WrapGoJNIEnv(
        (*gojni.JNIEnv)(unsafe.Pointer(env)),
    )
    if genv == nil {
        return value
    }

    return value
}

func main() {}
```

Add `unsafe` to the imports when using the pointer conversion shown above. Verify the `WrapGoJNIEnv` signature and environment representation against your version of `gojni`.

The `//export` directive exposes the Go function to cgo's generated C code. The exported symbol must match the Java native method's JNI naming convention.

For a static native Java method:

```java
package com.example;

public class MyClass {
    static {
        System.loadLibrary("example");
    }

    public static native int myfunc(int value);

    public static void main(String[] args) {
        System.out.println(myfunc(42));
    }
}
```

## Build a shared library

From the Go project containing `main.go`:

```bash
go build -buildmode=c-shared -o libexample.so .
```

This produces the shared library and a generated C header.

Verify that the JNI entry point is exported:

```bash
nm -D --defined-only libexample.so \
  | grep Java_com_example_MyClass_myfunc
```

Compile the Java class:

```bash
javac -d . java/com/example/MyClass.java
```

Run it:

```bash
java -Djava.library.path=. com.example.MyClass
```

The library name passed to `System.loadLibrary` omits the platform-specific prefix and extension. For example, `System.loadLibrary("example")` loads `libexample.so` on Linux when it is on the library search path.

## Arrays

`gojni` supports primitive Java arrays and object arrays.

The wrapper includes operations for:

- Obtaining array lengths
- Reading primitive array elements
- Reading and writing array regions
- Getting and setting object array elements

### Important: element access and copying

In the current implementation, methods such as `GetIntArrayElements` return a Go slice copied from the native JNI buffer. The native buffer is released before the method returns.

Consequently, modifying the returned slice does **not** automatically update the Java array.

For example:

```go
arr := genv.GetIntArrayElements(javaArray, nil, 0)
if len(arr) >= 2 {
    arr[1] = 90

    // Copy the updated data back to Java.
    genv.SetIntArrayRegion(javaArray, 0, arr)
}
```

The exact parameter types depend on the wrapper's method signatures.

### Release modes

The standard JNI release modes are:

| Mode | JNI meaning |
|---|---|
| `0` | Copy changes back and release the native buffer |
| `JNI_COMMIT` (`1`) | Copy changes back without releasing the native buffer |
| `JNI_ABORT` (`2`) | Discard changes to a copied native buffer and release it |

If a Go wrapper returns a copied Go slice and releases the native buffer internally, a custom `Release*ArrayElements` method cannot literally reproduce the native JNI lifecycle. Document and test such methods as copy-back APIs unless they retain the original native buffer until release.

## Exceptions

JNI exceptions can remain pending after an operation fails. Depending on the operation, check or clear exceptions before making further JNI calls that require no pending exception.

The wrapper includes exception operations such as:

- `ExceptionCheck`
- `ExceptionOccurred`
- `ExceptionClear`
- `ExceptionDescribe`
- `Throw`
- `ThrowNew`

Use these methods to propagate Java exceptions or handle errors appropriately. Do not silently clear exceptions unless that is intentional.

## References

JNI references have different lifetimes:

- **Local references** are generally scoped to a native call or local frame.
- **Global references** remain valid until explicitly deleted.
- **Weak global references** do not keep their Java objects alive.

Delete references when they are no longer needed. Use local frames when working with many temporary references.

## Threads and JVM access

A `JNIEnv` is thread-specific and must not be reused across unrelated threads.

For native-created threads, obtain the `JavaVM` and use the appropriate attach and detach operations before accessing JNI. Ensure that attached threads are detached when finished, and do not share a `JNIEnv` across threads.

## Testing

Run the Go package tests:

```bash
go test ./...
```

Run any Java integration tests using the repository's provided test scripts or documented commands.

The project's JNI test suite covers:

- Classes and type checks
- Exceptions and references
- Object creation and method calls
- Instance and static fields
- Strings
- Primitive and object arrays
- Direct buffers
- Monitors and reflection
- JVM access and thread attachment
- Class definition, modules, and native registration

Integration tests should run with a compatible JDK and the shared library built for the current platform.

## Platform considerations

JNI headers and shared-library formats vary by operating system.

- **Linux:** JNI headers usually include `include/linux`; shared libraries use `.so`.
- **Windows:** use the appropriate JDK headers and Windows C toolchain; shared libraries use `.dll`.
- **macOS:** use the appropriate JDK headers and platform build configuration; dynamic libraries typically use `.dylib`.

The cgo compiler configuration, JNI helper headers, library extension, and runtime search path must match the target platform.

## Limitations and safety

- JNI handles are valid only within their defined lifetime and thread context.
- JNI calls can leave Java exceptions pending.
- Native memory must be released according to the relevant JNI API's lifecycle.
- Go slices returned as copies do not automatically reflect changes in Java arrays.
- JNI operations are not all safe to call while an exception is pending.
- Go and Java garbage collectors manage different memory domains; native references must be handled explicitly.
- Cross-platform support depends on the build configuration and available JNI helpers.

## Development

Format Go source:

```bash
gofmt -w .
```

Run tests:

```bash
go test ./...
```

Build a shared library from a `package main` integration project:

```bash
go build -buildmode=c-shared -o libexample.so .
```

## Contributing

Contributions are welcome. When submitting changes:

1. Keep JNI behavior consistent with the JNI specification.
2. Handle null references, invalid ranges, and pending exceptions carefully.
3. Add tests for new wrapper methods.
4. Verify Java integration tests with a supported JDK.
5. Document any Go-specific behavior that differs from native JNI semantics.

## License

Licensed under the **MIT License**.

## Author

**MrR736**
<[MrR736@users.github.com](mailto:MrR736@users.github.com)>