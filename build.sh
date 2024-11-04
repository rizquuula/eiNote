export VERSION=0.1.0

buildUiCmd="docker build -f Dockerfile.ui -t einote:ui-$VERSION ."
echo "##### RUNNING: $buildUiCmd"
eval $buildUiCmd

buildCoreCmd="docker build -f Dockerfile.core -t einote:core-$VERSION ."
echo "##### RUNNING: $buildCoreCmd"
eval $buildCoreCmd
