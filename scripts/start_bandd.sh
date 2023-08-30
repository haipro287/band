# start bandchain
bandd start \
--p2p.laddr tcp://0.0.0.0:36656 \
--rpc.laddr tcp://0.0.0.0:36657 \
--grpc-web.address 0.0.0.0:10091 \
--grpc.address 0.0.0.0:10090 \
--rpc.pprof_laddr 0.0.0.0:7060 \
