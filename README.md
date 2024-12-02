
# Point Identity Network Overview

The **Point Identity Network** is a permissioned blockchain specifically designed for managing digital identities and access rights. It empowers individuals to regain control over their personal data and online interactions by offering:

- **Scalability**: Handles a growing number of use cases and participants efficiently.
- **Interoperability**: Ensures seamless integration across different platforms and ecosystems.
- **Security**: Provides a trusted and tamper-proof [Verifiable Data Registry (VDR)](https://www.w3.org/TR/did-core/#dfn-verifiable-data-registry) for implementing Self-Sovereign Identity (SSI) solutions.

## Key Features

- **Built on Cosmos-SDK**: Leverages the [Cosmos-SDK framework](https://tendermint.com/sdk/) to ensure high performance, modularity, and a robust architecture.
- **W3C Compliance**: Fully compatible with [W3C Decentralized Identifier (DID) standards](https://www.w3.org/TR/did-core/), enabling trusted and standards-aligned digital identity solutions.

The Point Identity Network is a foundational technology for building decentralized, user-centric digital ecosystems, emphasizing privacy, transparency, and control.

## Features

- Register, Update and Deactivate DID Documents
- Store/Update Credential Schema
- Store/Update status of a Verifiable Credential
- Stake `$PIDX` tokens
- Submit Governance Proposals
- Transfer `$PIDX` tokens within and across different Tendermint-based blockchains using IBC
- Deploy CosmWasm Smart Contracts

## Prerequisite

Following are the prerequisites that needs to be installed:

- Golang (Installation Guide: https://go.dev/doc/install) (version: 1.21)
- make
- jq

## Get started

### Local Binary

1. Clone this repository and install the binary:
   ```sh
   git clone https://github.com/pointidentity/pidx-node.git
   cd pidx-node
   make install
   ```

> The binary `pidx-noded` is usually generated in `$HOME/go/bin` directory. Run `pidx-noded --help` to explore its functionalities

2. Run the following script to setup a single-node blockchain. Please note that the following script requires `jq` to be installed.
   ```sh
   bash ./scripts/localnet-single-node/setup.sh
   ```

3. Start `pidx-noded`:
   ```sh
   pidx-noded start
   ```

### Docker

To run a single node `pidx-node` docker container, follow the below steps:

1. Pull the image:
   ```sh
   docker pull ghcr.io/pointidentity/pidx-node:latest
   ```

2. Run the following:
   ```sh
   docker run --rm -d \
	-p 26657:26657 -p 1317:1317 -p 26656:26656 -p 9090:9090 \
	--name pidx-node-container \
	ghcr.io/pointidentity/pidx-node start
   ```

## Documentation

| Topic | Reference |
| ----- | ---- |
| Decentralised Identifiers | https://docs.pointidentity.com/self-sovereign-identity/decentralized-identifier-did |
| Credential Schema | https://docs.pointidentity.com/self-sovereign-identity-ssi/schema |
| Verifiable Credential Status | https://docs.pointidentity.com/self-sovereign-identity/verifiable-credential-vc/credential-revocation-registry |


Please contact [support@pointidentity.com](mailto:support@pointidentity.com) for any questions.
