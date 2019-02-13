# -*- Autoconf -*-
# Process this file with autoconf to produce a configure script.

AC_PREREQ([2.69])
AC_INIT([go-opennsl], [0.1], [admin@example.net])
AM_INIT_AUTOMAKE([foreign])
LT_INIT([dlopen])

# Checks for programs.
AC_PROG_CC
AM_PROG_LIBTOOL
AC_PROG_RANLIB
AC_PROG_INSTALL

# Checks for libraries.
# FIXME: Replace `main' with a function in `-lc':
AC_CHECK_LIB([c], [main])

# Checks for header files.
AC_CHECK_HEADERS([arpa/inet.h fcntl.h float.h limits.h netinet/in.h stdint.h stdlib.h string.h sys/socket.h unistd.h])

# Checks for typedefs, structures, and compiler characteristics.
AC_TYPE_SIZE_T
AC_TYPE_UINT32_T

# Checks for library functions.

# export variables
AC_SUBST(libopennsl_stub_version_info, 1:0:0)
AC_CONFIG_FILES([
    Makefile
    libopennsl/Makefile
])
AC_OUTPUT
