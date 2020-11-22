#!/bin/bash
set -eux

[ -d build ] || mkdir build

if [ ! -e build/build.ninja ]; then
	meson . build \
		-Dgallium-drivers=panfrost,kmsro,swrast
fi

ninja -C build

sudo rm -vfr \
	/usr/local/lib/aarch64-linux-gnu/dri \
	/usr/local/lib/aarch64-linux-gnu/libEGL* \
	/usr/local/lib/aarch64-linux-gnu/libGL* \
	/usr/local/lib/aarch64-linux-gnu/libXvMC* \
	/usr/local/lib/aarch64-linux-gnu/libgbm* \
	/usr/local/lib/aarch64-linux-gnu/libgl* \
	/usr/local/lib/aarch64-linux-gnu/libxatracker* \
	/usr/local/lib/aarch64-linux-gnu/pkgconfig \
	/usr/local/lib/aarch64-linux-gnu/vdpau

sudo ninja -C build install
sudo ldconfig /usr/local/lib
