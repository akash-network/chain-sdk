# Changelog

## [1.0.0-alpha.29](https://github.com/akash-network/chain-sdk/compare/ts/v1.0.0-alpha.28...ts/v1.0.0-alpha.29) (2026-03-31)


### Features

* adds possibility to specify cert serial number ([#274](https://github.com/akash-network/chain-sdk/issues/274)) ([5322905](https://github.com/akash-network/chain-sdk/commit/532290565a8a8112276372411480400431e7a409))


### Bug Fixes

* **ts:** adds retry to chain sdk queries on connection errors ([#277](https://github.com/akash-network/chain-sdk/issues/277)) ([49c5d6b](https://github.com/akash-network/chain-sdk/commit/49c5d6bddeab5de7194c53cf0b625edacede908b))
* **ts:** handles service.expose as an optional value ([#247](https://github.com/akash-network/chain-sdk/issues/247)) ([c56179e](https://github.com/akash-network/chain-sdk/commit/c56179ef4b0d624ccbe853fb93dd42de244e3728))

## [1.0.0-alpha.28](https://github.com/akash-network/chain-sdk/compare/ts/v1.0.0-alpha.27...ts/v1.0.0-alpha.28) (2026-03-30)


### Features

* add EventAggregatedPrice ([#269](https://github.com/akash-network/chain-sdk/issues/269)) ([5c5cdea](https://github.com/akash-network/chain-sdk/commit/5c5cdea57bb1be7b8c1796e2f10d16064c2b013c))
* include spread in the ledger records ([#262](https://github.com/akash-network/chain-sdk/issues/262)) ([48edd00](https://github.com/akash-network/chain-sdk/commit/48edd000a0186be1cf3791080817690538244c75))
* oracle api v2 ([#271](https://github.com/akash-network/chain-sdk/issues/271)) ([9bf1e08](https://github.com/akash-network/chain-sdk/commit/9bf1e0855620fd5b9a87a74268bbbbe4a3249131))
* **oracle:** add future time drift param ([#272](https://github.com/akash-network/chain-sdk/issues/272)) ([05f70f8](https://github.com/akash-network/chain-sdk/commit/05f70f8a40ee7d650f35c5aee4c0eefdf3908d79))


### Bug Fixes

* **oracle:** params tag IDs ([#273](https://github.com/akash-network/chain-sdk/issues/273)) ([602a567](https://github.com/akash-network/chain-sdk/commit/602a56793ee1726699f9c66be6c80f62416cf667))

## [1.0.0-alpha.27](https://github.com/akash-network/chain-sdk/compare/ts/v1.0.0-alpha.26...ts/v1.0.0-alpha.27) (2026-03-18)


### Features

* implement ledger failed record for bme ([#255](https://github.com/akash-network/chain-sdk/issues/255)) ([4dc78aa](https://github.com/akash-network/chain-sdk/commit/4dc78aa46dd0538a2d7ce099e274ca47115581c4))
* **ts:** adds yaml.raw helper ([#260](https://github.com/akash-network/chain-sdk/issues/260)) ([28ead7b](https://github.com/akash-network/chain-sdk/commit/28ead7b910320702d38443d73a4dd00236032ce6))


### Bug Fixes

* add ledger failed event ([#257](https://github.com/akash-network/chain-sdk/issues/257)) ([40c2596](https://github.com/akash-network/chain-sdk/commit/40c2596fe4e03bbbd30254b5b40063075f5b5a0b))

## [1.0.0-alpha.26](https://github.com/akash-network/chain-sdk/compare/ts/v1.0.0-alpha.25...ts/v1.0.0-alpha.26) (2026-03-13)


### Features

* **bme:** add FundVault rpc ([#250](https://github.com/akash-network/chain-sdk/issues/250)) ([bf19b1a](https://github.com/akash-network/chain-sdk/commit/bf19b1a7fe2a1e2f14987267efed547a1843476b))

## [1.0.0-alpha.25](https://github.com/akash-network/chain-sdk/compare/ts/v1.0.0-alpha.24...ts/v1.0.0-alpha.25) (2026-03-09)


### Features

* **bme:** add endblocker params ([#245](https://github.com/akash-network/chain-sdk/issues/245)) ([4f6a1d2](https://github.com/akash-network/chain-sdk/commit/4f6a1d2e2e1bf3f25af59bf89f24e18f24b999a7))
* **go/sdl:** add permissions params ([#200](https://github.com/akash-network/chain-sdk/issues/200)) ([38be010](https://github.com/akash-network/chain-sdk/commit/38be010ce056239b063ab9a66d70ad17b571ab24))
* **market:** add OfferPrices and BidMinDeposits to market proto ([#242](https://github.com/akash-network/chain-sdk/issues/242)) ([0433094](https://github.com/akash-network/chain-sdk/commit/0433094ddd9ca07e905a7d7c8b5b272fb267ea50))
* **provider:** add BidScreening RPC and validation proto ([#243](https://github.com/akash-network/chain-sdk/issues/243)) ([1fadf95](https://github.com/akash-network/chain-sdk/commit/1fadf955bae5d60e4a79701ad1e7591d64f0b0a8))
* **ts:** adds support for chain transactions ([#246](https://github.com/akash-network/chain-sdk/issues/246)) ([0d175db](https://github.com/akash-network/chain-sdk/commit/0d175db778634eeb65815972f3c6263ab590d619))


### Bug Fixes

* **ts:** sort storage keys during manifest generation ([#233](https://github.com/akash-network/chain-sdk/issues/233)) ([1d5e0e4](https://github.com/akash-network/chain-sdk/commit/1d5e0e4c8bd9d417f7829ab9d65467054d695341))

## [1.0.0-alpha.24](https://github.com/akash-network/chain-sdk/compare/ts/v1.0.0-alpha.23...ts/v1.0.0-alpha.24) (2026-02-12)


### Features

* **ts:** support service permissions in SDL manifest ([b58ab52](https://github.com/akash-network/chain-sdk/commit/b58ab52554cf08f911569b9e115c715ddd29d32a))

## [1.0.0-alpha.23](https://github.com/akash-network/chain-sdk/compare/ts/v1.0.0-alpha.22...ts/v1.0.0-alpha.23) (2026-02-11)


### Features

* allows uact denom ([#222](https://github.com/akash-network/chain-sdk/issues/222)) ([39e2329](https://github.com/akash-network/chain-sdk/commit/39e23295da6114afde9d5642907e3f4458ee9051))
* implements generate manifest function ([#225](https://github.com/akash-network/chain-sdk/issues/225)) ([6ad9d64](https://github.com/akash-network/chain-sdk/commit/6ad9d6488c8af74ae1a91783353469cbbb368964))

## [1.0.0-alpha.22](https://github.com/akash-network/chain-sdk/compare/ts/v1.0.0-alpha.21...ts/v1.0.0-alpha.22) (2026-02-09)


### Features

* **ts:** adds sdl permissions validation ([419e3b7](https://github.com/akash-network/chain-sdk/commit/419e3b7f73ec3c7b49926ec079d4cf357d61853c))

## [1.0.0-alpha.21](https://github.com/akash-network/chain-sdk/compare/ts/v1.0.0-alpha.20...ts/v1.0.0-alpha.21) (2026-02-09)


### Features

* **ts:** adds sdl permissions validation ([419e3b7](https://github.com/akash-network/chain-sdk/commit/419e3b7f73ec3c7b49926ec079d4cf357d61853c))

## [1.0.0-alpha.21](https://github.com/akash-network/chain-sdk/compare/ts/v1.0.0-alpha.20...ts/v1.0.0-alpha.21) (2026-01-25)

### Code Refactoring

* **ts:** reduces generated types size

## [1.0.0-alpha.20](https://github.com/akash-network/chain-sdk/compare/ts/v1.0.0-alpha.19...ts/v1.0.0-alpha.20) (2026-01-20)


### ⚠ BREAKING CHANGES

* **ts:** simplifies TxClient interface ([#210](https://github.com/akash-network/chain-sdk/issues/210))

### Features

* **ts:** expose validateSDL and its primitives ([#204](https://github.com/akash-network/chain-sdk/issues/204)) ([38b2e98](https://github.com/akash-network/chain-sdk/commit/38b2e981c6c3000def92e9c1918bcae53f2356d7))


### Code Refactoring

* **ts:** simplifies TxClient interface ([#210](https://github.com/akash-network/chain-sdk/issues/210)) ([42c02b3](https://github.com/akash-network/chain-sdk/commit/42c02b30639d6ae68079fac58afc3cc8ebdf9f08))

## [1.0.0-alpha.19](https://github.com/akash-network/chain-sdk/compare/ts/v1.0.0-alpha.18...ts/v1.0.0-alpha.19) (2026-01-13)


### ⚠ BREAKING CHANGES

* **ts:** adds SDL json-schema and refactors validation ([#184](https://github.com/akash-network/chain-sdk/issues/184))

### Features

* **ts:** adds retry options to grpc transports ([#174](https://github.com/akash-network/chain-sdk/issues/174)) ([02bd13b](https://github.com/akash-network/chain-sdk/commit/02bd13bfb0a4f24a2770c57dc96081cd235b6115))


### Bug Fixes

* ceil resource convert string ([#172](https://github.com/akash-network/chain-sdk/issues/172)) ([4a2663b](https://github.com/akash-network/chain-sdk/commit/4a2663b2307bc5e4fe8f2449a9b57c62e885b02f))


### Code Refactoring

* **ts:** adds SDL json-schema and refactors validation ([#184](https://github.com/akash-network/chain-sdk/issues/184)) ([4b67a98](https://github.com/akash-network/chain-sdk/commit/4b67a989dbd204f52f40937e5434a75f57b495e5))
