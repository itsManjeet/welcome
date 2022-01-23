#!/bin/sh

echo "package app

const UI = \`$(cat ui.glade)\`" > app/ui.go