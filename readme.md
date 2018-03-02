This package contains root CA pools for use with a client `tls.Config`. They are
statically compiled into go code, and should work anywhere, regardless of what your system's
trusted CAs are.

The intent of this is for testing certificates to ensure compatibility with a variety of platforms.

You should **not** use this in any real-world workloads.