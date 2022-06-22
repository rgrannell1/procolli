#! /usr/bin/env zsh

./bs/build.zsh

gecko --green '🥦 copying binary to /usr/bin/procolli'
sudo cp procolli /usr/bin/procolli
gecko --green '🥦 procolli installed!'
