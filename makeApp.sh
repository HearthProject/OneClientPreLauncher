#!/bin/bash

VERSION="$1"

echo "Generating osx app for $VERSION"

chmod +x "build/out/OneClient-$VERSION-darwin-amd64"

./appify "build/out/OneClient-$VERSION-darwin-amd64" "OneClient-$VERSION-osx"

mkdir "OneClient-$VERSION-osx.app/Contents/Resources"
cp icon.icns "OneClient-$VERSION-osx.app/Contents/Resources/icon.icns"
cp Info.plist "OneClient-$VERSION-osx.app/Contents/Info.plist"

mv "OneClient-$VERSION-osx.app" "build/out/OneClient-$VERSION-osx.app"

cd build/out/

zip -r "OneClient-$VERSION-osx.zip" "OneClient-$VERSION-osx.app"
cd ..
cd ..

#Create a dmg file here