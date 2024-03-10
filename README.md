# algourl

Convert transactions on standard input to ARC-0026 compatible format with ASCII QR code output.

This utility allows to quickly online or offline a participation key, especially when your account is rekeyed and/or protected by hardware security device like Ledger.

## Quick Start

### Installation

```bash
go install github.com/algonode/algourl@latest
```

The `algourl` binary will be available (in most cases) @ `~/go/bin/algourl`

### Online an account with Mobile wallet 

⚠️ Make sure the participation key is generated and imported (see full examples)
</br>ℹ You do not need new key in case your online status got suspended for not voting.

```bash
goal account changeonlinestatus -a 7THLM2QWR2VOGCLIQGGASSTHV6GBZAB2OMYRCDEQ2K5PXP42YKIAUKLOHM -t - | ~/go/bin/algourl 
```

### Offline an account with Mobile wallet 
```bash
goal account changeonlinestatus -a 7THLM2QWR2VOGCLIQGGASSTHV6GBZAB2OMYRCDEQ2K5PXP42YKIAUKLOHM --online=0 -t - | ~/go/bin/algourl 
```

## Installation (full)

## Usage examples (full)

### New participation key

Initial state:
* No active participation key for the account on the node
* Account has some non zero balance (even 0.05 Algo is OK)
* Account private keys **are not** stored on the node (GOOD!)
* SSH access to the node :) 


#### Let's generate new participation keys bundle for an account



```bash
# goal account addpartkey --address=7THLM2QWR2VOGCLIQGGASSTHV6GBZAB2OMYRCDEQ2K5PXP42YKIAUKLOHM --roundFirstValid=36876865 --roundLastValid=36976865
Please stand by while generating keys. This might take a few minutes...
Participation key generation successful. Participation ID: TB3UFX2DUXSA7W7CDMWRAR4QJRALTXLGA54NPI2SIGD63XPJXX6A
```


## Example output

```bash

#> goal account changeonlinestatus -a 7THLM2QWR2VOGCLIQGGASSTHV6GBZAB2OMYRCDEQ2K5PXP42YKIAUKLOHM --online=0 -t - | algourl 

Paste below URL into your browser or scan QR code to online/offline the account
algorand://7THLM2QWR2VOGCLIQGGASSTHV6GBZAB2OMYRCDEQ2K5PXP42YKIAUKLOHM?type=keyreg&votefst=0&votelst=0

█████████████████████████████████████████████
██ ▄▄▄▄▄ █▀▄▄█ ▀▀█ ▄▀ ▀▀▀ ▀▄ ▀▀▀ ▀▀█ ▄▄▄▄▄ ██
██ █   █ ██▄▄ ▄▀▄▄   ▀▀█▄█ █▀▄██▄█ █ █   █ ██
██ █▄▄▄█ █▄███▄▄▀█▀█ ▀▀█▀▀▀ █▀█▄█▀██ █▄▄▄█ ██
██▄▄▄▄▄▄▄█▄█▄█▄█▄█ ▀ ▀ ▀ ▀ ▀▄█▄▀▄▀▄█▄▄▄▄▄▄▄██
██ █ █▄▀▄██▀▀██▄ ▄ ▀█▀ █▀ ▀▀█   ▄▄██▀▀▄▀█ ▀██
██ ▄ ▄▀▄▄██▀ ▀▄▀█▀▄▀ █▀▀ ▀█▀ ▀  █▀  ▄█▄▄▄█ ██
██▀▀▄▀▄█▄▀▀▀▀▀▀ █▄ ▄  █▀▀█▀███   █▀  ▄▄▀█▄▄██
██▀█▀▄▀█▄▄▄ ▀▄▀▄▀▀▄▀ ▀█▀▄██▀▄█ ▀ ██▀  ▄▄▄▄▄██
███  █ ▀▄▄▀  █▄▀█ ▀ ▀█ ▄██▀▀█▀ █▀▀   ▄█▀▄▄▄██
██▀ █▄▄▄▄▀█▀█▄▀ █▀▄▀ ██▀▀▀ ▀█▀ ▀▀███ ▄▀█▄▄▄██
██  ▄ ▀▀▄ ██▀▄   ▀█▀     ▀ ▄█ ▀▄   ▄▄▄█▄▄▀▀██
██ █▄█▀▀▄▀▀█▀▀▀   █  ▀  ▄█▀▀██▄ ▀▀ ▀ ▄▀██▄▄██
██████▄▀▄ ▄ ▄ ██ ▀▀▀█▀    ▀▀█▀▀ ▄█ ▄▄ █▄▀  ██
████▀▄  ▄▀▀█▄███▄█▀ ▄█  ▄█▄ ▄█▀ ██ ▄▄▀▄█▄▄▄██
███▀▄ █ ▄ ▀ ▄▀  ▀ █▀█▀ ▀██▀▄ █▀▀█▄▀  █▄██▄ ██
██▄▀▀▄▀▄▄    █▄█▀█▀ ▄▄ ▀▀▀ ▀▀█ ▀▀█ █▄▄█▄██ ██
██▄███▄▄▄▄ ▀▄███ ▀▄▄█▄   ▄▀  ▀ ▀█▀ ▄▄▄ ▀ ▀███
██ ▄▄▄▄▄ ████▄▀█▄▄ ▄█▀▀▀▄█▀▀▀██▀▄█ █▄█ █▄ ▄██
██ █   █ █▄▄ █ ███▀▀█  ▀▄  ▄█  ▄▄█▄▄▄  ▄▄ ▀██
██ █▄▄▄█ █▄▀ █ ▄ ▀▄  ██ ▄█▄▀▀█▀ ▀██▀█▄ ██ ▄██
██▄▄▄▄▄▄▄█▄▄█▄▄▄████▄██████▄▄██▄▄▄▄▄█▄▄██▄▄██

```


## TXN type support 

`algourl` utility supports the following transactions

- only participation key registration/deregistration transactions 
- unsigned raw MSGPack encoded files only

## Ecosystem support

> Generated URLs/QRs are ARC-0026 backward compatible but new standard that allows for non-pay transactions is  yet to be published

|Wallet|TX Support|Remarks|
|---|---|---|
|[Defly mobile](https://defly.app/)| KEYREG, PAY| |


# FAQ

### Q: Is it safe to use? Is it leaking keys ?

All information encoded into URL is public in nature (public keys) and safe for transport over untrusted medium (as all unsigned transactions are). Data in the URL/QR code contains no more information than written to the public ledger for everyone to see after any key registration transaction.
