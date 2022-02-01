#!/bin/sh

DESTDIR=${DESTDIR:-'/'}

mkdir -p ${DESTDIR}/usr/{bin,share/pixmaps}
mkdir -p ${DESTDIR}/etc/xdg/autostart/

cp assets/welcome.desktop ${DESTDIR}/etc/autostart/
./scripts/genui.sh

sed -i "s|assets/looks.svg|/usr/share/welcome/looks.svg|g" app/ui.go
go build -o ${DESTDIR}/usr/bin/welcome main.go

cp assets/*.png ${DESTDIR}/usr/share/pixmaps/
install -v -D -m 0755 assets/looks.svg -t ${DESTDIR}/usr/share/welcome/
