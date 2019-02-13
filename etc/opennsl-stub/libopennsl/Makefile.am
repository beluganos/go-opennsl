lib_LTLIBRARIES = libopennsl.la

libopennsl_la_SOURCES = \
	version.c \
	types.c \
	logger.c \
	driver.c \
	init.c \
	port.c \
	switch.c \
	link.c \
	tunnel.c \
	vlan.c \
	vxlan.c \
	l2.c \
	l2X.c \
	l3.c \
	mpls.c \
	field.c \
	multicast.c \
	stg.c \
	rx.c \
	tx.c \
	pkt.c \
	trunk.c \
	stat.c \
	cosq.c \
	bitop.c \
	mirror.c \
	libopennsl.c

libopennsl_la_CFLAGS = `pkg-config --cflags libopennsl`
libopennsl_la_LDFLAGS = -version-info $(libopennsl_stub_version_info)
