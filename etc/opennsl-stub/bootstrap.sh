#! /bin/bash
# -*- coding: utf-8 -*-

which make     > /dev/null || { echo "make not installed."     ; exit 1; }
which automake > /dev/null || { echo "automake not installed." ; exit 1; }
which autoconf > /dev/null || { echo "autoconf not installed." ; exit 1; }

aclocal
automake -ac
autoconf
./configure $OPTS

