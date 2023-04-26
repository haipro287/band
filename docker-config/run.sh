#!/bin/bash

bandd init $1 --chain-id bandchain

cp /chain/docker-config/$1/priv_validator_key.json ~/.band/config/priv_validator_key.json
cp /chain/docker-config/$1/node_key.json ~/.band/config/node_key.json
cp /chain/docker-config/genesis.json ~/.band/config/genesis.json
cp -r /chain/docker-config/files ~/.band

sed -E -i \
  "s/timeout_commit = \".*\"/timeout_commit = \"3s\"/" \
  ~/.band/config/config.toml

if [ "$1" == "query-node" ];then
    cp /chain/docker-config/app.toml ~/.band/config/app.toml
fi

if [ "$1" == "emitter" ];then
    sleep 60
    bandd start --with-emitter test@kafka:9092 \
    --p2p.persistent_peers 11392b605378063b1c505c0ab123f04bd710d7d7@multi-validator1-node:26656,0851086afcd835d5a6fb0ffbf96fcdf74fec742e@multi-validator2-node:26656,63808bd64f2ec19acb2a494c8ce8467c595f6fba@multi-validator3-node:26656,7b58b086dd915a79836eb8bfa956aeb9488d13b0@multi-validator4-node:26656
elif [ "$1" == "query-node" ];then
    bandd start --rpc.laddr tcp://0.0.0.0:26657 \
    --p2p.persistent_peers 11392b605378063b1c505c0ab123f04bd710d7d7@multi-validator1-node:26656,0851086afcd835d5a6fb0ffbf96fcdf74fec742e@multi-validator2-node:26656,63808bd64f2ec19acb2a494c8ce8467c595f6fba@multi-validator3-node:26656,7b58b086dd915a79836eb8bfa956aeb9488d13b0@multi-validator4-node:26656    --with-request-search=sqlite3:rqsearch.db \
    --with-pricer=1
else
    bandd start --rpc.laddr tcp://0.0.0.0:26657 \
    --p2p.persistent_peers 11392b605378063b1c505c0ab123f04bd710d7d7@multi-validator1-node:26656,0851086afcd835d5a6fb0ffbf96fcdf74fec742e@multi-validator2-node:26656,63808bd64f2ec19acb2a494c8ce8467c595f6fba@multi-validator3-node:26656,7b58b086dd915a79836eb8bfa956aeb9488d13b0@multi-validator4-node:26656
fi
