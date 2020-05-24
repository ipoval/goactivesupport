#!/usr/bin/env bash

source ./bin/helpers.bash

if [ -f version.go ]; then
  VERSION_DEFAULT=$(getAppVersion)
else
  VERSION_DEFAULT="0.1.1"
fi

echo -n "Version [x.x.x, default:${VERSION_DEFAULT}]: "
read version
version=${version:-"${VERSION_DEFAULT}"}

echo "Tagging: ${version}"
git tag "v${version}"
git push origin tag "v${version}"

echo "Done"
