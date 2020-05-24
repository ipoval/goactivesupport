#!/usr/bin/env bash

getAppVersion() {
  cat version.go | egrep -i 'version.*\d+\.\d+\.\d+' | cut -d"=" -f2 | tr -d '" '
}
