#!/bin/bash

rm -rf ~/.yoda

# config chain id
yoda config chain-id bandchain

yoda config node tcp://localhost:36657

# add validator to yoda config
yoda config validator $(bandd keys show validator -a --bech val --keyring-backend test)

# setup execution endpoint
# yoda config executor "rest:$EXECUTOR_URL/live/master?timeout=10s"
yoda config executor "rest:http://localhost:5000?timeout=10s"

# setup broadcast-timeout to yoda config
yoda config broadcast-timeout "5m"

# setup rpc-poll-interval to yoda config
yoda config rpc-poll-interval "1s"

# setup max-try to yoda config
yoda config max-try 5

echo "y" | bandd tx oracle activate --from validator --keyring-backend test --chain-id bandchain --node tcp://localhost:36657
# wait for activation transaction success
sleep 2

for i in $(eval echo {1..1})
do
  # add reporter key
  yoda keys add reporter$i
done

# send band tokens to reporters
echo "y" | bandd tx bank send  validator $(yoda keys list -a) 1000000uband --keyring-backend test --chain-id bandchain --node tcp://localhost:36657

# wait for sending band tokens transaction success
sleep 2

# add reporter to bandchain
echo "y" | bandd tx oracle add-reporters $(yoda keys list -a) --from validator --keyring-backend test --chain-id bandchain --node tcp://localhost:36657

# wait for addding reporter transaction success
sleep 2

# run yoda
yoda run
