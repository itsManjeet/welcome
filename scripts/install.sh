#!/bin/sh

DESTDIR=${DESTDIR:-'/'}

mkdir -p ${DESTDIR}/usr/{bin,share/pixmaps}
mkdir -p ${DESTDIR}/etc/autostart/

cp assets/welcome.desktop ${DESTDIR}/etc/autostart/
go build -o ${DESTDIR}/usr/bin/welcome main.go

cp assets/*.png ${DESTDIR}/usr/share/pixmaps/