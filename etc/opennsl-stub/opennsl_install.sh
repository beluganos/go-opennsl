#! /bin/bash
# -*- coding: utf-8 -*-

# Copyright (C) 2018 Nippon Telegraph and Telephone Corporation.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#    http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
# implied.
# See the License for the specific language governing permissions and
# limitations under the License.

WB_NAME=$2

GIT_URL=https://github.com/Broadcom-Switch/OpenNSL.git
WORKDIR=$HOME/opennsl
#WORKDIR=/opt/work/opennsl

do_git_sparse() {
    git config core.sparseCheckout true

    echo /include      >> .git/info/sparse-checkout
    echo /bin/$WB_NAME >> .git/info/sparse-checkout
}

do_clone() {
    if [ -e $WORKDIR ]; then
        return
    fi

    mkdir -p $WORKDIR

    pushd $WORKDIR

    git init
    git remote add origin $GIT_URL

    if [ "$WB_NAME" != "all" ]; then
        do_git_sparse
    fi

    git fetch --depth 1 origin master
    git pull  --depth 1 origin master

    popd
}

do_usage() {
    echo "Usage: $0 <clone | pkg-config | install> <as5712 | as7712 | ...>"
    exit 1
}

do_make_pkgconfig() {
    cat <<EOF > libopennsl.pc
Name: libopennsl
Description: OpenNSL library
Version: 3.5.0.1
Libs: -L${WORKDIR}/bin/$WB_NAME -lopennsl
Cflags: -I${WORKDIR}/include -DINCLUDE_L3
EOF
}

do_install_pkgconfig() {
    sudo install -pm 644 libopennsl.pc /usr/lib/pkgconfig/libopennsl.pc
}

do_check() {
    if [ "x$WB_NAME" = "x" ]; then
        do_usage
    fi
}

do_show() {
    pkg-config --cflags libopennsl || { echo "error!!"; }
    pkg-config --libs   libopennsl || { echo "error!!"; }
}

do_check
case $1 in
    clone)
        do_clone
        ;;

    pkg-config)
        do_make_pkgconfig
        do_install_pkgconfig
        ;;

    install)
        do_clone
        do_make_pkgconfig
        do_install_pkgconfig
        ;;
    show)
        do_show
        ;;
    *)
        do_usage
        ;;
esac
