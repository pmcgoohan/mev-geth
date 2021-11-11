# geth merger

This is a fork of [mev-geth](https://github.com/flashbots/mev-geth) with additional logging capabilities.

* Based on [mev-geth release v1.10.9-mev0.4](https://github.com/flashbots/mev-geth/tree/release/v1.10.9-mev0.4)
* Metabase dashboard: https://data.flashbots.net/dashboard/71-geth-logger-dashboard

A good starting point for the logger changes is [`newMultiworker(..)` in `miner/multi_worker.go`](https://github.com/flashbots/mev-geth-logger/blob/blocklog/miner/multi_worker.go#L118).
## Getting started

You can build, run and test mev-geth logger locally:

* Sending bundles and megabundles locally can be tested using [mev-geth-demo](https://github.com/flashbots/mev-geth-demo).
* You can use docker-compose to start a Postgres database and an graphical database interface (Adminer) (listening on [localhost:8093](http://localhost:8093/?pgsql=db&username=user&db=db&ns=public&select=mevgeth_log_summary)).

```bash
# Start the database container
docker-compose up

# Build geth
make geth

# Clear the local database
make clear-db

# Start built code in privnet mode (starts a fresh private net)
./run-privnet.sh

# Run the mev-geth-demo tests
make test-local
```

At this point, blocks will be produced and database entries written.

Database helpers:

```bash
# Connect to the database with psql:
docker exec -it mev-geth-private_db_1 /usr/bin/psql -h localhost -U user db

| Version | Spec                                                                                        |
| ------- | ------------------------------------------------------------------------------------------- |
| v0.3    | [MEV-Geth Spec v0.3](https://docs.flashbots.net/flashbots-auction/miners/mev-geth-spec/v03) |
| v0.2    | [MEV-Geth Spec v0.2](https://docs.flashbots.net/flashbots-auction/miners/mev-geth-spec/v02) |
| v0.1    | [MEV-Geth Spec v0.1](https://docs.flashbots.net/flashbots-auction/miners/mev-geth-spec/v01) |


`mev-geth-demo` setup setup:

```bash
mkdir _test
cd _test
git clone git@github.com:flashbots/mev-geth-demo.git
cd mev-geth-demo

# Install dependencies
yarn
```
