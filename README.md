# go-opennsl

OpenNSL library for Go lang.

## Pre-condition
Acqurire header file and liberally from [OpenNSL](https://github.com/Broadcom-Switch/OpenNSL), and execute pkg-config.

### register pkg-config

```
$ cat /usr/lib/pkgconfig/libopennsl.pc 

Name: libopennsl
Description: OpenNSL library
Version: 3.5.0.1
Libs: -L/opt/work/opennsl/bin/as7712 -lopennsl
Cflags: -I/opt/work/opennsl/include -DINCLUDE_L3
```

### execute pkg-config

```
$ pkg-config --cflags libopennsl
-DINCLUDE_L3 -I/opt/work/opennsl/include

$ pkg-config --libs libopennsl
-L/opt/work/opennsl/bin/as7712 -lopennsl
```

## Build and Test

```
$ go get github.com/beluganos/go-opennsl
```
