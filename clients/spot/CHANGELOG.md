### Changelog

## 1.12.0 - 2026-08-19

### Changed (1)

- Updated `github.com/binance/binance-connector-go/common/v2` library to version `2.7.0`.

## 1.11.0 - 2026-07-15

### Changed (1)

- Updated `github.com/binance/binance-connector-go/common/v2` library to version `2.6.0`.

## 1.10.0 - 2026-07-14

### Changed (77)

- Updated `github.com/binance/binance-connector-go/common/v2` library to version `2.4.0`.

#### REST API

- Modified parameter `cancelRestrictions`:
  - enum removed: `NEW`, `PARTIALLY_FILLED`
  - affected methods:
    - `deleteOrder()` (`DELETE /api/v3/order`)
    - `orderCancelReplace()` (`POST /api/v3/order/cancelReplace`)
- Modified parameter `newOrderRespType`:
  - enum removed: `MARKET`, `LIMIT`
  - affected methods:
    - `newOrder()` (`POST /api/v3/order`)
    - `orderCancelReplace()` (`POST /api/v3/order/cancelReplace`)
    - `orderOco()` (`POST /api/v3/order/oco`)
    - `orderTest()` (`POST /api/v3/order/test`)
    - `orderListOco()` (`POST /api/v3/orderList/oco`)
    - `orderListOpo()` (`POST /api/v3/orderList/opo`)
    - `orderListOpoco()` (`POST /api/v3/orderList/opoco`)
    - `orderListOto()` (`POST /api/v3/orderList/oto`)
    - `orderListOtoco()` (`POST /api/v3/orderList/otoco`)
    - `sorOrder()` (`POST /api/v3/sor/order`)
    - `sorOrderTest()` (`POST /api/v3/sor/order/test`)
- Modified parameter `pegOffsetType`:
  - enum removed: `NON_REPRESENTABLE`
  - affected methods:
    - `newOrder()` (`POST /api/v3/order`)
    - `orderCancelReplace()` (`POST /api/v3/order/cancelReplace`)
    - `orderTest()` (`POST /api/v3/order/test`)
- Modified parameter `pegPriceType`:
  - enum removed: `NON_REPRESENTABLE`
  - affected methods:
    - `newOrder()` (`POST /api/v3/order`)
    - `orderCancelReplace()` (`POST /api/v3/order/cancelReplace`)
    - `orderTest()` (`POST /api/v3/order/test`)
- Modified parameter `permissions`:
  - items: enum added: `SPOT`, `MARGIN`, `LEVERAGED`, `TRD_GRP_002`, `TRD_GRP_003`, `TRD_GRP_004`, `TRD_GRP_005`, `TRD_GRP_006`, `TRD_GRP_007`, `TRD_GRP_008`, `TRD_GRP_009`, `TRD_GRP_010`, `TRD_GRP_011`, `TRD_GRP_012`, `TRD_GRP_013`, `TRD_GRP_014`, `TRD_GRP_015`, `TRD_GRP_016`, `TRD_GRP_017`, `TRD_GRP_018`, `TRD_GRP_019`, `TRD_GRP_020`, `TRD_GRP_021`, `TRD_GRP_022`, `TRD_GRP_023`, `TRD_GRP_024`, `TRD_GRP_025`
  - affected methods:
    - `exchangeInfo()` (`GET /api/v3/exchangeInfo`)
- Modified parameter `selfTradePreventionMode`:
  - enum removed: `NON_REPRESENTABLE`
  - affected methods:
    - `newOrder()` (`POST /api/v3/order`)
    - `orderCancelReplace()` (`POST /api/v3/order/cancelReplace`)
    - `orderOco()` (`POST /api/v3/order/oco`)
    - `orderTest()` (`POST /api/v3/order/test`)
    - `orderListOco()` (`POST /api/v3/orderList/oco`)
    - `orderListOpo()` (`POST /api/v3/orderList/opo`)
    - `orderListOpoco()` (`POST /api/v3/orderList/opoco`)
    - `orderListOto()` (`POST /api/v3/orderList/oto`)
    - `orderListOtoco()` (`POST /api/v3/orderList/otoco`)
    - `sorOrder()` (`POST /api/v3/sor/order`)
    - `sorOrderTest()` (`POST /api/v3/sor/order/test`)
- Modified parameter `symbolStatus`:
  - enum removed: `END_OF_DAY`, `NON_REPRESENTABLE`
  - affected methods:
    - `depth()` (`GET /api/v3/depth`)
    - `exchangeInfo()` (`GET /api/v3/exchangeInfo`)
    - `executionRules()` (`GET /api/v3/executionRules`)
    - `referencePriceCalculation()` (`GET /api/v3/referencePrice/calculation`)
    - `tickerBookTicker()` (`GET /api/v3/ticker/bookTicker`)
    - `tickerPrice()` (`GET /api/v3/ticker/price`)
    - `tickerTradingDay()` (`GET /api/v3/ticker/tradingDay`)
- Modified parameter `symbolStatus`:
  - enum removed: `END_OF_DAY`, `NON_REPRESENTABLE`
  - affected methods:
    - `ticker()` (`GET /api/v3/ticker`)
    - `ticker24hr()` (`GET /api/v3/ticker/24hr`)
- Modified parameter `timeInForce`:
  - enum removed: `NON_REPRESENTABLE`
  - affected methods:
    - `newOrder()` (`POST /api/v3/order`)
    - `orderCancelReplace()` (`POST /api/v3/order/cancelReplace`)
    - `orderTest()` (`POST /api/v3/order/test`)
    - `sorOrder()` (`POST /api/v3/sor/order`)
    - `sorOrderTest()` (`POST /api/v3/sor/order/test`)
- Modified parameter `type`:
  - enum removed: `NON_REPRESENTABLE`
  - affected methods:
    - `newOrder()` (`POST /api/v3/order`)
    - `orderCancelReplace()` (`POST /api/v3/order/cancelReplace`)
    - `orderTest()` (`POST /api/v3/order/test`)
- Modified parameter `type`:
  - enum removed: `STOP_LOSS`, `STOP_LOSS_LIMIT`, `TAKE_PROFIT`, `TAKE_PROFIT_LIMIT`, `LIMIT_MAKER`, `NON_REPRESENTABLE`
  - affected methods:
    - `sorOrder()` (`POST /api/v3/sor/order`)
    - `sorOrderTest()` (`POST /api/v3/sor/order/test`)
- Modified parameter `windowSize`:
  - enum added: `7d`
  - affected methods:
    - `ticker()` (`GET /api/v3/ticker`)
- Modified response for `allOrders()` (`GET /api/v3/allOrders`):
  - items: property `workingFloor` added
  - items: property `preventedMatchId` added
  - items: property `strategyId` added
  - items: property `expiryReason` added
  - items: property `pegOffsetType` added
  - items: property `trailingDelta` added
  - items: property `peggedPrice` added
  - items: property `usedSor` added
  - items: property `pegOffsetValue` added
  - items: property `trailingTime` added
  - items: property `pegPriceType` added
  - items: property `preventedQuantity` added
  - items: property `strategyType` added
  - items: item property `workingFloor` added
  - items: item property `preventedMatchId` added
  - items: item property `strategyId` added
  - items: item property `expiryReason` added
  - items: item property `pegOffsetType` added
  - items: item property `trailingDelta` added
  - items: item property `peggedPrice` added
  - items: item property `usedSor` added
  - items: item property `pegOffsetValue` added
  - items: item property `trailingTime` added
  - items: item property `pegPriceType` added
  - items: item property `preventedQuantity` added
  - items: item property `strategyType` added

- Modified response for `depth()` (`GET /api/v3/depth`):
  - `asks`.items: minItems `0` → `2`
  - `asks`.items: maxItems `null` → `2`
  - `bids`.items: minItems `0` → `2`
  - `bids`.items: maxItems `null` → `2`

- Modified response for `exchangeInfo()` (`GET /api/v3/exchangeInfo`):
  - property `sors` added
  - `exchangeFilters`.items: oneOf modified
  - `symbols`.items.`filters`.items: oneOf modified
  - `symbols`.items.`filters`.items: oneOf modified

- Modified response for `klines()` (`GET /api/v3/klines`):
  - items.items: oneOf added 2 schema(s)
  - items.items: oneOf removed 2 schema(s)

- Modified response for `myFilters()` (`GET /api/v3/myFilters`):
  - `assetFilters`.items: oneOf modified
  - `exchangeFilters`.items: oneOf modified
  - `symbolFilters`.items: oneOf modified

- Modified response for `getOpenOrders()` (`GET /api/v3/openOrders`):
  - items: property `pegPriceType` added
  - items: property `expiryReason` added
  - items: property `peggedPrice` added
  - items: property `pegOffsetValue` added
  - items: property `strategyId` added
  - items: property `preventedQuantity` added
  - items: property `strategyType` added
  - items: property `workingFloor` added
  - items: property `trailingDelta` added
  - items: property `trailingTime` added
  - items: property `preventedMatchId` added
  - items: property `usedSor` added
  - items: property `pegOffsetType` added
  - items: item property `pegPriceType` added
  - items: item property `expiryReason` added
  - items: item property `peggedPrice` added
  - items: item property `pegOffsetValue` added
  - items: item property `strategyId` added
  - items: item property `preventedQuantity` added
  - items: item property `strategyType` added
  - items: item property `workingFloor` added
  - items: item property `trailingDelta` added
  - items: item property `trailingTime` added
  - items: item property `preventedMatchId` added
  - items: item property `usedSor` added
  - items: item property `pegOffsetType` added

- Modified response for `deleteOrder()` (`DELETE /api/v3/order`):
  - property `strategyType` added
  - property `usedSor` added
  - property `pegPriceType` added
  - property `expiryReason` added
  - property `peggedPrice` added
  - property `icebergQty` added
  - property `pegOffsetValue` added
  - property `stopPrice` added
  - property `trailingDelta` added
  - property `workingFloor` added
  - property `strategyId` added
  - property `pegOffsetType` added
  - property `trailingTime` added
  - property `preventedQuantity` added
  - property `preventedMatchId` added

- Modified response for `getOrder()` (`GET /api/v3/order`):
  - property `expiryReason` added
  - property `pegPriceType` added
  - property `pegOffsetType` added
  - property `peggedPrice` added
  - property `usedSor` added
  - property `workingFloor` added
  - property `pegOffsetValue` added
  - property `preventedQuantity` added
  - property `strategyType` added
  - property `strategyId` added
  - property `trailingTime` added
  - property `preventedMatchId` added
  - property `trailingDelta` added

- Modified response for `orderAmendKeepPriority()` (`PUT /api/v3/order/amend/keepPriority`):
  - `amendedOrder`: property `preventedQuantity` added
  - `amendedOrder`: property `pegOffsetValue` added
  - `amendedOrder`: property `expiryReason` added
  - `amendedOrder`: property `pegPriceType` added
  - `amendedOrder`: property `strategyType` added
  - `amendedOrder`: property `peggedPrice` added
  - `amendedOrder`: property `preventedMatchId` added
  - `amendedOrder`: property `trailingTime` added
  - `amendedOrder`: property `trailingDelta` added
  - `amendedOrder`: property `pegOffsetType` added
  - `amendedOrder`: property `strategyId` added
  - `amendedOrder`: property `usedSor` added
  - `amendedOrder`: property `icebergQty` added
  - `amendedOrder`: property `workingFloor` added
  - `amendedOrder`: property `stopPrice` added

- Modified response for `orderCancelReplace()` (`POST /api/v3/order/cancelReplace`):
  - property `code` deleted
  - property `data` deleted
  - property `msg` deleted
  - `cancelResponse`: property `pegOffsetType` added
  - `cancelResponse`: property `pegPriceType` added
  - `cancelResponse`: property `preventedMatchId` added
  - `cancelResponse`: property `strategyType` added
  - `cancelResponse`: property `trailingDelta` added
  - `cancelResponse`: property `usedSor` added
  - `cancelResponse`: property `workingFloor` added
  - `cancelResponse`: property `expiryReason` added
  - `cancelResponse`: property `stopPrice` added
  - `cancelResponse`: property `trailingTime` added
  - `cancelResponse`: property `strategyId` added
  - `cancelResponse`: property `icebergQty` added
  - `cancelResponse`: property `pegOffsetValue` added
  - `cancelResponse`: property `peggedPrice` added
  - `cancelResponse`: property `preventedQuantity` added
  - `newOrderResponse`: property `pegOffsetValue` added
  - `newOrderResponse`: property `strategyType` added
  - `newOrderResponse`: property `peggedPrice` added
  - `newOrderResponse`: property `trailingDelta` added
  - `newOrderResponse`: property `expiryReason` added
  - `newOrderResponse`: property `preventedQuantity` added
  - `newOrderResponse`: property `usedSor` added
  - `newOrderResponse`: property `pegOffsetType` added
  - `newOrderResponse`: property `pegPriceType` added
  - `newOrderResponse`: property `preventedMatchId` added
  - `newOrderResponse`: property `stopPrice` added
  - `newOrderResponse`: property `strategyId` added
  - `newOrderResponse`: property `workingFloor` added
  - `newOrderResponse`: property `icebergQty` added
  - `newOrderResponse`: property `trailingTime` added
  - `newOrderResponse`.`fills`.items: type `string` → `object`
  - `newOrderResponse`.`fills`.items: property `tradeId` added
  - `newOrderResponse`.`fills`.items: property `commission` added
  - `newOrderResponse`.`fills`.items: property `commissionAsset` added
  - `newOrderResponse`.`fills`.items: property `price` added
  - `newOrderResponse`.`fills`.items: property `qty` added
  - `newOrderResponse`.`fills`.items: item property `tradeId` added
  - `newOrderResponse`.`fills`.items: item property `commission` added
  - `newOrderResponse`.`fills`.items: item property `commissionAsset` added
  - `newOrderResponse`.`fills`.items: item property `price` added
  - `newOrderResponse`.`fills`.items: item property `qty` added

- Modified response for `deleteOrderList()` (`DELETE /api/v3/orderList`):
  - `orderReports`.items: property `pegOffsetType` added
  - `orderReports`.items: property `peggedPrice` added
  - `orderReports`.items: property `trailingTime` added
  - `orderReports`.items: property `strategyId` added
  - `orderReports`.items: property `strategyType` added
  - `orderReports`.items: property `pegOffsetValue` added
  - `orderReports`.items: property `icebergQty` added
  - `orderReports`.items: property `preventedMatchId` added
  - `orderReports`.items: property `usedSor` added
  - `orderReports`.items: property `expiryReason` added
  - `orderReports`.items: property `workingFloor` added
  - `orderReports`.items: property `preventedQuantity` added
  - `orderReports`.items: property `trailingDelta` added
  - `orderReports`.items: property `pegPriceType` added
  - `orderReports`.items: property `selfTradePreventionMode` deleted
  - `orderReports`.items: item property `pegOffsetType` added
  - `orderReports`.items: item property `peggedPrice` added
  - `orderReports`.items: item property `trailingTime` added
  - `orderReports`.items: item property `strategyId` added
  - `orderReports`.items: item property `strategyType` added
  - `orderReports`.items: item property `pegOffsetValue` added
  - `orderReports`.items: item property `icebergQty` added
  - `orderReports`.items: item property `preventedMatchId` added
  - `orderReports`.items: item property `usedSor` added
  - `orderReports`.items: item property `expiryReason` added
  - `orderReports`.items: item property `workingFloor` added
  - `orderReports`.items: item property `preventedQuantity` added
  - `orderReports`.items: item property `trailingDelta` added
  - `orderReports`.items: item property `pegPriceType` added
  - `orderReports`.items: item property `selfTradePreventionMode` deleted

- Modified response for `orderListOpo()` (`POST /api/v3/orderList/opo`):
  - `orderReports`.items: property `icebergQty` added
  - `orderReports`.items: property `pegOffsetValue` added
  - `orderReports`.items: property `usedSor` added
  - `orderReports`.items: property `stopPrice` added
  - `orderReports`.items: property `strategyId` added
  - `orderReports`.items: property `pegPriceType` added
  - `orderReports`.items: property `pegOffsetType` added
  - `orderReports`.items: property `workingFloor` added
  - `orderReports`.items: property `trailingTime` added
  - `orderReports`.items: property `strategyType` added
  - `orderReports`.items: property `peggedPrice` added
  - `orderReports`.items: property `preventedMatchId` added
  - `orderReports`.items: property `preventedQuantity` added
  - `orderReports`.items: property `expiryReason` added
  - `orderReports`.items: property `trailingDelta` added
  - `orderReports`.items: item property `icebergQty` added
  - `orderReports`.items: item property `pegOffsetValue` added
  - `orderReports`.items: item property `usedSor` added
  - `orderReports`.items: item property `stopPrice` added
  - `orderReports`.items: item property `strategyId` added
  - `orderReports`.items: item property `pegPriceType` added
  - `orderReports`.items: item property `pegOffsetType` added
  - `orderReports`.items: item property `workingFloor` added
  - `orderReports`.items: item property `trailingTime` added
  - `orderReports`.items: item property `strategyType` added
  - `orderReports`.items: item property `peggedPrice` added
  - `orderReports`.items: item property `preventedMatchId` added
  - `orderReports`.items: item property `preventedQuantity` added
  - `orderReports`.items: item property `expiryReason` added
  - `orderReports`.items: item property `trailingDelta` added

- Modified response for `orderListOpoco()` (`POST /api/v3/orderList/opoco`):
  - `orderReports`.items: property `expiryReason` added
  - `orderReports`.items: property `trailingDelta` added
  - `orderReports`.items: property `usedSor` added
  - `orderReports`.items: property `peggedPrice` added
  - `orderReports`.items: property `pegPriceType` added
  - `orderReports`.items: property `preventedQuantity` added
  - `orderReports`.items: property `icebergQty` added
  - `orderReports`.items: property `pegOffsetValue` added
  - `orderReports`.items: property `trailingTime` added
  - `orderReports`.items: property `workingFloor` added
  - `orderReports`.items: property `strategyId` added
  - `orderReports`.items: property `preventedMatchId` added
  - `orderReports`.items: property `strategyType` added
  - `orderReports`.items: property `pegOffsetType` added
  - `orderReports`.items: item property `expiryReason` added
  - `orderReports`.items: item property `trailingDelta` added
  - `orderReports`.items: item property `usedSor` added
  - `orderReports`.items: item property `peggedPrice` added
  - `orderReports`.items: item property `pegPriceType` added
  - `orderReports`.items: item property `preventedQuantity` added
  - `orderReports`.items: item property `icebergQty` added
  - `orderReports`.items: item property `pegOffsetValue` added
  - `orderReports`.items: item property `trailingTime` added
  - `orderReports`.items: item property `workingFloor` added
  - `orderReports`.items: item property `strategyId` added
  - `orderReports`.items: item property `preventedMatchId` added
  - `orderReports`.items: item property `strategyType` added
  - `orderReports`.items: item property `pegOffsetType` added

- Modified response for `orderListOto()` (`POST /api/v3/orderList/oto`):
  - `orderReports`.items: property `pegOffsetValue` added
  - `orderReports`.items: property `strategyId` added
  - `orderReports`.items: property `trailingTime` added
  - `orderReports`.items: property `pegPriceType` added
  - `orderReports`.items: property `icebergQty` added
  - `orderReports`.items: property `expiryReason` added
  - `orderReports`.items: property `workingFloor` added
  - `orderReports`.items: property `preventedMatchId` added
  - `orderReports`.items: property `peggedPrice` added
  - `orderReports`.items: property `usedSor` added
  - `orderReports`.items: property `strategyType` added
  - `orderReports`.items: property `trailingDelta` added
  - `orderReports`.items: property `stopPrice` added
  - `orderReports`.items: property `pegOffsetType` added
  - `orderReports`.items: property `preventedQuantity` added
  - `orderReports`.items: item property `pegOffsetValue` added
  - `orderReports`.items: item property `strategyId` added
  - `orderReports`.items: item property `trailingTime` added
  - `orderReports`.items: item property `pegPriceType` added
  - `orderReports`.items: item property `icebergQty` added
  - `orderReports`.items: item property `expiryReason` added
  - `orderReports`.items: item property `workingFloor` added
  - `orderReports`.items: item property `preventedMatchId` added
  - `orderReports`.items: item property `peggedPrice` added
  - `orderReports`.items: item property `usedSor` added
  - `orderReports`.items: item property `strategyType` added
  - `orderReports`.items: item property `trailingDelta` added
  - `orderReports`.items: item property `stopPrice` added
  - `orderReports`.items: item property `pegOffsetType` added
  - `orderReports`.items: item property `preventedQuantity` added

- Modified response for `orderListOtoco()` (`POST /api/v3/orderList/otoco`):
  - `orderReports`.items: property `preventedQuantity` added
  - `orderReports`.items: property `expiryReason` added
  - `orderReports`.items: property `pegPriceType` added
  - `orderReports`.items: property `pegOffsetValue` added
  - `orderReports`.items: property `trailingDelta` added
  - `orderReports`.items: property `usedSor` added
  - `orderReports`.items: property `icebergQty` added
  - `orderReports`.items: property `pegOffsetType` added
  - `orderReports`.items: property `peggedPrice` added
  - `orderReports`.items: property `preventedMatchId` added
  - `orderReports`.items: property `strategyType` added
  - `orderReports`.items: property `trailingTime` added
  - `orderReports`.items: property `strategyId` added
  - `orderReports`.items: property `workingFloor` added
  - `orderReports`.items: item property `preventedQuantity` added
  - `orderReports`.items: item property `expiryReason` added
  - `orderReports`.items: item property `pegPriceType` added
  - `orderReports`.items: item property `pegOffsetValue` added
  - `orderReports`.items: item property `trailingDelta` added
  - `orderReports`.items: item property `usedSor` added
  - `orderReports`.items: item property `icebergQty` added
  - `orderReports`.items: item property `pegOffsetType` added
  - `orderReports`.items: item property `peggedPrice` added
  - `orderReports`.items: item property `preventedMatchId` added
  - `orderReports`.items: item property `strategyType` added
  - `orderReports`.items: item property `trailingTime` added
  - `orderReports`.items: item property `strategyId` added
  - `orderReports`.items: item property `workingFloor` added

- Modified response for `ticker()` (`GET /api/v3/ticker`):
  - oneOf modified

- Modified response for `ticker24hr()` (`GET /api/v3/ticker/24hr`):
  - oneOf modified

- Modified response for `tickerBookTicker()` (`GET /api/v3/ticker/bookTicker`):
  - oneOf modified

- Modified response for `tickerPrice()` (`GET /api/v3/ticker/price`):
  - oneOf modified

- Modified response for `tickerTradingDay()` (`GET /api/v3/ticker/tradingDay`):
  - oneOf modified

- Modified response for `uiKlines()` (`GET /api/v3/uiKlines`):
  - items.items: oneOf added 2 schema(s)
  - items.items: oneOf removed 2 schema(s)

- Marked `orderOco()` (`POST /api/v3/order/oco`) as deprecated.

#### WebSocket API

- Modified parameter `cancelRestrictions`:
  - enum removed: `NEW`, `PARTIALLY_FILLED`
  - affected methods:
    - `orderCancel()` (`order.cancel` method)
    - `orderCancelReplace()` (`order.cancelReplace` method)
- Modified parameter `newOrderRespType`:
  - enum removed: `MARKET`, `LIMIT`
  - affected methods:
    - `orderCancelReplace()` (`order.cancelReplace` method)
    - `orderPlace()` (`order.place` method)
    - `orderTest()` (`order.test` method)
    - `orderListPlace()` (`orderList.place` method)
    - `orderListPlaceOco()` (`orderList.place.oco` method)
    - `orderListPlaceOpo()` (`orderList.place.opo` method)
    - `orderListPlaceOpoco()` (`orderList.place.opoco` method)
    - `orderListPlaceOto()` (`orderList.place.oto` method)
    - `orderListPlaceOtoco()` (`orderList.place.otoco` method)
    - `sorOrderPlace()` (`sor.order.place` method)
    - `sorOrderTest()` (`sor.order.test` method)
- Modified parameter `pegOffsetType`:
  - enum removed: `NON_REPRESENTABLE`
  - affected methods:
    - `orderCancelReplace()` (`order.cancelReplace` method)
    - `orderPlace()` (`order.place` method)
    - `orderTest()` (`order.test` method)
- Modified parameter `pegPriceType`:
  - enum removed: `NON_REPRESENTABLE`
  - affected methods:
    - `orderCancelReplace()` (`order.cancelReplace` method)
    - `orderPlace()` (`order.place` method)
    - `orderTest()` (`order.test` method)
- Modified parameter `selfTradePreventionMode`:
  - enum removed: `NON_REPRESENTABLE`
  - affected methods:
    - `orderCancelReplace()` (`order.cancelReplace` method)
    - `orderPlace()` (`order.place` method)
    - `orderTest()` (`order.test` method)
    - `orderListPlace()` (`orderList.place` method)
    - `orderListPlaceOco()` (`orderList.place.oco` method)
    - `orderListPlaceOpo()` (`orderList.place.opo` method)
    - `orderListPlaceOpoco()` (`orderList.place.opoco` method)
    - `orderListPlaceOto()` (`orderList.place.oto` method)
    - `orderListPlaceOtoco()` (`orderList.place.otoco` method)
    - `sorOrderPlace()` (`sor.order.place` method)
    - `sorOrderTest()` (`sor.order.test` method)
- Modified parameter `symbolStatus`:
  - enum removed: `END_OF_DAY`, `NON_REPRESENTABLE`
  - affected methods:
    - `depth()` (`depth` method)
    - `exchangeInfo()` (`exchangeInfo` method)
    - `executionRules()` (`executionRules` method)
    - `ticker()` (`ticker` method)
    - `ticker24hr()` (`ticker.24hr` method)
    - `tickerBook()` (`ticker.book` method)
    - `tickerPrice()` (`ticker.price` method)
    - `tickerTradingDay()` (`ticker.tradingDay` method)
- Modified parameter `symbolStatus`:
  - enum removed: `END_OF_DAY`, `NON_REPRESENTABLE`
  - affected methods:
    - `referencePriceCalculation()` (`referencePrice.calculation` method)
- Modified parameter `timeInForce`:
  - enum removed: `NON_REPRESENTABLE`
  - affected methods:
    - `orderCancelReplace()` (`order.cancelReplace` method)
    - `orderPlace()` (`order.place` method)
    - `orderTest()` (`order.test` method)
    - `sorOrderPlace()` (`sor.order.place` method)
    - `sorOrderTest()` (`sor.order.test` method)
- Modified parameter `type`:
  - enum removed: `NON_REPRESENTABLE`
  - affected methods:
    - `orderCancelReplace()` (`order.cancelReplace` method)
    - `orderPlace()` (`order.place` method)
    - `orderTest()` (`order.test` method)
- Modified parameter `type`:
  - enum removed: `STOP_LOSS`, `STOP_LOSS_LIMIT`, `TAKE_PROFIT`, `TAKE_PROFIT_LIMIT`, `LIMIT_MAKER`, `NON_REPRESENTABLE`
  - affected methods:
    - `sorOrderPlace()` (`sor.order.place` method)
    - `sorOrderTest()` (`sor.order.test` method)
- Modified parameter `windowSize`:
  - enum added: `7d`
  - affected methods:
    - `ticker()` (`ticker` method)
- Modified response for `allOrders()` (`allOrders` method):
  - `result`.items: property `pegOffsetType` added
  - `result`.items: property `peggedPrice` added
  - `result`.items: property `workingFloor` added
  - `result`.items: property `strategyType` added
  - `result`.items: property `pegOffsetValue` added
  - `result`.items: property `strategyId` added
  - `result`.items: property `pegPriceType` added
  - `result`.items: property `usedSor` added
  - `result`.items: property `trailingDelta` added
  - `result`.items: property `trailingTime` added
  - `result`.items: property `expiryReason` added
  - `result`.items: item property `pegOffsetType` added
  - `result`.items: item property `peggedPrice` added
  - `result`.items: item property `workingFloor` added
  - `result`.items: item property `strategyType` added
  - `result`.items: item property `pegOffsetValue` added
  - `result`.items: item property `strategyId` added
  - `result`.items: item property `pegPriceType` added
  - `result`.items: item property `usedSor` added
  - `result`.items: item property `trailingDelta` added
  - `result`.items: item property `trailingTime` added
  - `result`.items: item property `expiryReason` added

- Modified response for `depth()` (`depth` method):
  - `result`.`asks`.items: minItems `0` → `2`
  - `result`.`asks`.items: maxItems `null` → `2`
  - `result`.`bids`.items: minItems `0` → `2`
  - `result`.`bids`.items: maxItems `null` → `2`

- Modified response for `exchangeInfo()` (`exchangeInfo` method):
  - property `symbols` added
  - property `timezone` added
  - property `exchangeFilters` added
  - property `serverTime` added
  - property `sors` added
  - property `status` deleted
  - property `id` deleted
  - property `result` deleted

- Modified response for `klines()` (`klines` method):
  - `result`.items: minItems `0` → `12`
  - `result`.items: maxItems `null` → `12`
  - `result`.items.items: oneOf added 2 schema(s)
  - `result`.items.items: oneOf removed 2 schema(s)

- Modified response for `myFilters()` (`myFilters` method):
  - property `assetFilters` added
  - property `exchangeFilters` added
  - property `symbolFilters` added
  - property `result` deleted
  - property `status` deleted
  - property `id` deleted

- Modified response for `openOrdersCancelAll()` (`openOrders.cancelAll` method):
  - `result`.items: property `workingFloor` added
  - `result`.items: property `expiryReason` added
  - `result`.items: property `peggedPrice` added
  - `result`.items: property `usedSor` added
  - `result`.items: property `preventedMatchId` added
  - `result`.items: property `preventedQuantity` added
  - `result`.items: property `pegOffsetValue` added
  - `result`.items: property `pegOffsetType` added
  - `result`.items: property `pegPriceType` added
  - `result`.items.`orderReports`.items: property `pegOffsetType` added
  - `result`.items.`orderReports`.items: property `preventedMatchId` added
  - `result`.items.`orderReports`.items: property `pegPriceType` added
  - `result`.items.`orderReports`.items: property `strategyId` added
  - `result`.items.`orderReports`.items: property `preventedQuantity` added
  - `result`.items.`orderReports`.items: property `peggedPrice` added
  - `result`.items.`orderReports`.items: property `usedSor` added
  - `result`.items.`orderReports`.items: property `icebergQty` added
  - `result`.items.`orderReports`.items: property `expiryReason` added
  - `result`.items.`orderReports`.items: property `strategyType` added
  - `result`.items.`orderReports`.items: property `trailingDelta` added
  - `result`.items.`orderReports`.items: property `trailingTime` added
  - `result`.items.`orderReports`.items: property `workingFloor` added
  - `result`.items.`orderReports`.items: property `pegOffsetValue` added
  - `result`.items.`orderReports`.items: item property `pegOffsetType` added
  - `result`.items.`orderReports`.items: item property `preventedMatchId` added
  - `result`.items.`orderReports`.items: item property `pegPriceType` added
  - `result`.items.`orderReports`.items: item property `strategyId` added
  - `result`.items.`orderReports`.items: item property `preventedQuantity` added
  - `result`.items.`orderReports`.items: item property `peggedPrice` added
  - `result`.items.`orderReports`.items: item property `usedSor` added
  - `result`.items.`orderReports`.items: item property `icebergQty` added
  - `result`.items.`orderReports`.items: item property `expiryReason` added
  - `result`.items.`orderReports`.items: item property `strategyType` added
  - `result`.items.`orderReports`.items: item property `trailingDelta` added
  - `result`.items.`orderReports`.items: item property `trailingTime` added
  - `result`.items.`orderReports`.items: item property `workingFloor` added
  - `result`.items.`orderReports`.items: item property `pegOffsetValue` added
  - `result`.items: item property `workingFloor` added
  - `result`.items: item property `expiryReason` added
  - `result`.items: item property `peggedPrice` added
  - `result`.items: item property `usedSor` added
  - `result`.items: item property `preventedMatchId` added
  - `result`.items: item property `preventedQuantity` added
  - `result`.items: item property `pegOffsetValue` added
  - `result`.items: item property `pegOffsetType` added
  - `result`.items: item property `pegPriceType` added
  - `result`.items.`orderReports`.items: property `pegOffsetType` added
  - `result`.items.`orderReports`.items: property `preventedMatchId` added
  - `result`.items.`orderReports`.items: property `pegPriceType` added
  - `result`.items.`orderReports`.items: property `strategyId` added
  - `result`.items.`orderReports`.items: property `preventedQuantity` added
  - `result`.items.`orderReports`.items: property `peggedPrice` added
  - `result`.items.`orderReports`.items: property `usedSor` added
  - `result`.items.`orderReports`.items: property `icebergQty` added
  - `result`.items.`orderReports`.items: property `expiryReason` added
  - `result`.items.`orderReports`.items: property `strategyType` added
  - `result`.items.`orderReports`.items: property `trailingDelta` added
  - `result`.items.`orderReports`.items: property `trailingTime` added
  - `result`.items.`orderReports`.items: property `workingFloor` added
  - `result`.items.`orderReports`.items: property `pegOffsetValue` added
  - `result`.items.`orderReports`.items: item property `pegOffsetType` added
  - `result`.items.`orderReports`.items: item property `preventedMatchId` added
  - `result`.items.`orderReports`.items: item property `pegPriceType` added
  - `result`.items.`orderReports`.items: item property `strategyId` added
  - `result`.items.`orderReports`.items: item property `preventedQuantity` added
  - `result`.items.`orderReports`.items: item property `peggedPrice` added
  - `result`.items.`orderReports`.items: item property `usedSor` added
  - `result`.items.`orderReports`.items: item property `icebergQty` added
  - `result`.items.`orderReports`.items: item property `expiryReason` added
  - `result`.items.`orderReports`.items: item property `strategyType` added
  - `result`.items.`orderReports`.items: item property `trailingDelta` added
  - `result`.items.`orderReports`.items: item property `trailingTime` added
  - `result`.items.`orderReports`.items: item property `workingFloor` added
  - `result`.items.`orderReports`.items: item property `pegOffsetValue` added

- Modified response for `openOrdersStatus()` (`openOrders.status` method):
  - `result`.items: property `trailingDelta` added
  - `result`.items: property `preventedMatchId` added
  - `result`.items: property `preventedQuantity` added
  - `result`.items: property `strategyId` added
  - `result`.items: property `usedSor` added
  - `result`.items: property `pegOffsetValue` added
  - `result`.items: property `strategyType` added
  - `result`.items: property `expiryReason` added
  - `result`.items: property `pegPriceType` added
  - `result`.items: property `trailingTime` added
  - `result`.items: property `pegOffsetType` added
  - `result`.items: property `peggedPrice` added
  - `result`.items: property `workingFloor` added
  - `result`.items: item property `trailingDelta` added
  - `result`.items: item property `preventedMatchId` added
  - `result`.items: item property `preventedQuantity` added
  - `result`.items: item property `strategyId` added
  - `result`.items: item property `usedSor` added
  - `result`.items: item property `pegOffsetValue` added
  - `result`.items: item property `strategyType` added
  - `result`.items: item property `expiryReason` added
  - `result`.items: item property `pegPriceType` added
  - `result`.items: item property `trailingTime` added
  - `result`.items: item property `pegOffsetType` added
  - `result`.items: item property `peggedPrice` added
  - `result`.items: item property `workingFloor` added

- Modified response for `orderAmendKeepPriority()` (`order.amend.keepPriority` method):
  - `result`.`amendedOrder`: property `usedSor` added
  - `result`.`amendedOrder`: property `workingFloor` added
  - `result`.`amendedOrder`: property `pegOffsetType` added
  - `result`.`amendedOrder`: property `trailingDelta` added
  - `result`.`amendedOrder`: property `stopPrice` added
  - `result`.`amendedOrder`: property `pegPriceType` added
  - `result`.`amendedOrder`: property `pegOffsetValue` added
  - `result`.`amendedOrder`: property `trailingTime` added
  - `result`.`amendedOrder`: property `strategyType` added
  - `result`.`amendedOrder`: property `icebergQty` added
  - `result`.`amendedOrder`: property `preventedQuantity` added
  - `result`.`amendedOrder`: property `expiryReason` added
  - `result`.`amendedOrder`: property `peggedPrice` added
  - `result`.`amendedOrder`: property `preventedMatchId` added
  - `result`.`amendedOrder`: property `strategyId` added

- Modified response for `orderCancel()` (`order.cancel` method):
  - `result`: property `peggedPrice` added
  - `result`: property `pegOffsetType` added
  - `result`: property `pegPriceType` added
  - `result`: property `expiryReason` added
  - `result`: property `preventedQuantity` added
  - `result`: property `usedSor` added
  - `result`: property `trailingTime` added
  - `result`: property `workingFloor` added
  - `result`: property `preventedMatchId` added
  - `result`: property `pegOffsetValue` added
  - `result`.`orderReports`.items: property `expiryReason` added
  - `result`.`orderReports`.items: property `strategyType` added
  - `result`.`orderReports`.items: property `usedSor` added
  - `result`.`orderReports`.items: property `pegPriceType` added
  - `result`.`orderReports`.items: property `trailingTime` added
  - `result`.`orderReports`.items: property `preventedQuantity` added
  - `result`.`orderReports`.items: property `pegOffsetType` added
  - `result`.`orderReports`.items: property `strategyId` added
  - `result`.`orderReports`.items: property `trailingDelta` added
  - `result`.`orderReports`.items: property `workingFloor` added
  - `result`.`orderReports`.items: property `pegOffsetValue` added
  - `result`.`orderReports`.items: property `preventedMatchId` added
  - `result`.`orderReports`.items: property `icebergQty` added
  - `result`.`orderReports`.items: property `peggedPrice` added
  - `result`.`orderReports`.items: item property `expiryReason` added
  - `result`.`orderReports`.items: item property `strategyType` added
  - `result`.`orderReports`.items: item property `usedSor` added
  - `result`.`orderReports`.items: item property `pegPriceType` added
  - `result`.`orderReports`.items: item property `trailingTime` added
  - `result`.`orderReports`.items: item property `preventedQuantity` added
  - `result`.`orderReports`.items: item property `pegOffsetType` added
  - `result`.`orderReports`.items: item property `strategyId` added
  - `result`.`orderReports`.items: item property `trailingDelta` added
  - `result`.`orderReports`.items: item property `workingFloor` added
  - `result`.`orderReports`.items: item property `pegOffsetValue` added
  - `result`.`orderReports`.items: item property `preventedMatchId` added
  - `result`.`orderReports`.items: item property `icebergQty` added
  - `result`.`orderReports`.items: item property `peggedPrice` added

- Modified response for `orderCancelReplace()` (`order.cancelReplace` method):
  - `result`.`cancelResponse`: property `expiryReason` added
  - `result`.`cancelResponse`: property `peggedPrice` added
  - `result`.`cancelResponse`: property `pegOffsetValue` added
  - `result`.`cancelResponse`: property `pegOffsetType` added
  - `result`.`cancelResponse`: property `strategyType` added
  - `result`.`cancelResponse`: property `icebergQty` added
  - `result`.`cancelResponse`: property `strategyId` added
  - `result`.`cancelResponse`: property `trailingTime` added
  - `result`.`cancelResponse`: property `preventedMatchId` added
  - `result`.`cancelResponse`: property `trailingDelta` added
  - `result`.`cancelResponse`: property `workingFloor` added
  - `result`.`cancelResponse`: property `pegPriceType` added
  - `result`.`cancelResponse`: property `stopPrice` added
  - `result`.`cancelResponse`: property `usedSor` added
  - `result`.`cancelResponse`: property `preventedQuantity` added
  - `result`.`newOrderResponse`: property `preventedQuantity` added
  - `result`.`newOrderResponse`: property `pegOffsetValue` added
  - `result`.`newOrderResponse`: property `strategyId` added
  - `result`.`newOrderResponse`: property `workingFloor` added
  - `result`.`newOrderResponse`: property `stopPrice` added
  - `result`.`newOrderResponse`: property `pegPriceType` added
  - `result`.`newOrderResponse`: property `peggedPrice` added
  - `result`.`newOrderResponse`: property `pegOffsetType` added
  - `result`.`newOrderResponse`: property `trailingTime` added
  - `result`.`newOrderResponse`: property `preventedMatchId` added
  - `result`.`newOrderResponse`: property `usedSor` added
  - `result`.`newOrderResponse`: property `expiryReason` added
  - `result`.`newOrderResponse`: property `trailingDelta` added
  - `result`.`newOrderResponse`: property `strategyType` added
  - `result`.`newOrderResponse`: property `icebergQty` added

- Modified response for `orderPlace()` (`order.place` method):
  - `result`: property `stopPrice` added
  - `result`: property `usedSor` added
  - `result`: property `icebergQty` added
  - `result`: property `pegOffsetType` added
  - `result`: property `pegOffsetValue` added
  - `result`: property `preventedMatchId` added
  - `result`: property `trailingTime` added
  - `result`: property `peggedPrice` added
  - `result`: property `expiryReason` added
  - `result`: property `pegPriceType` added
  - `result`: property `preventedQuantity` added
  - `result`: property `trailingDelta` added
  - `result`: property `workingFloor` added
  - `result`: property `strategyId` added
  - `result`: property `strategyType` added

- Modified response for `orderStatus()` (`order.status` method):
  - `result`: property `pegOffsetType` added
  - `result`: property `workingFloor` added
  - `result`: property `pegOffsetValue` added
  - `result`: property `usedSor` added
  - `result`: property `expiryReason` added
  - `result`: property `pegPriceType` added
  - `result`: property `peggedPrice` added

- Modified response for `orderListCancel()` (`orderList.cancel` method):
  - `result`.`orderReports`.items: property `preventedQuantity` added
  - `result`.`orderReports`.items: property `strategyId` added
  - `result`.`orderReports`.items: property `trailingTime` added
  - `result`.`orderReports`.items: property `strategyType` added
  - `result`.`orderReports`.items: property `trailingDelta` added
  - `result`.`orderReports`.items: property `workingFloor` added
  - `result`.`orderReports`.items: property `icebergQty` added
  - `result`.`orderReports`.items: property `pegPriceType` added
  - `result`.`orderReports`.items: property `expiryReason` added
  - `result`.`orderReports`.items: property `pegOffsetValue` added
  - `result`.`orderReports`.items: property `peggedPrice` added
  - `result`.`orderReports`.items: property `preventedMatchId` added
  - `result`.`orderReports`.items: property `pegOffsetType` added
  - `result`.`orderReports`.items: property `usedSor` added
  - `result`.`orderReports`.items: item property `preventedQuantity` added
  - `result`.`orderReports`.items: item property `strategyId` added
  - `result`.`orderReports`.items: item property `trailingTime` added
  - `result`.`orderReports`.items: item property `strategyType` added
  - `result`.`orderReports`.items: item property `trailingDelta` added
  - `result`.`orderReports`.items: item property `workingFloor` added
  - `result`.`orderReports`.items: item property `icebergQty` added
  - `result`.`orderReports`.items: item property `pegPriceType` added
  - `result`.`orderReports`.items: item property `expiryReason` added
  - `result`.`orderReports`.items: item property `pegOffsetValue` added
  - `result`.`orderReports`.items: item property `peggedPrice` added
  - `result`.`orderReports`.items: item property `preventedMatchId` added
  - `result`.`orderReports`.items: item property `pegOffsetType` added
  - `result`.`orderReports`.items: item property `usedSor` added

- Modified response for `orderListPlace()` (`orderList.place` method):
  - `result`.`orderReports`.items: property `trailingTime` added
  - `result`.`orderReports`.items: property `workingFloor` added
  - `result`.`orderReports`.items: property `strategyType` added
  - `result`.`orderReports`.items: property `trailingDelta` added
  - `result`.`orderReports`.items: property `icebergQty` added
  - `result`.`orderReports`.items: property `pegOffsetValue` added
  - `result`.`orderReports`.items: property `expiryReason` added
  - `result`.`orderReports`.items: property `usedSor` added
  - `result`.`orderReports`.items: property `pegOffsetType` added
  - `result`.`orderReports`.items: property `preventedMatchId` added
  - `result`.`orderReports`.items: property `preventedQuantity` added
  - `result`.`orderReports`.items: property `strategyId` added
  - `result`.`orderReports`.items: property `pegPriceType` added
  - `result`.`orderReports`.items: property `peggedPrice` added
  - `result`.`orderReports`.items: item property `trailingTime` added
  - `result`.`orderReports`.items: item property `workingFloor` added
  - `result`.`orderReports`.items: item property `strategyType` added
  - `result`.`orderReports`.items: item property `trailingDelta` added
  - `result`.`orderReports`.items: item property `icebergQty` added
  - `result`.`orderReports`.items: item property `pegOffsetValue` added
  - `result`.`orderReports`.items: item property `expiryReason` added
  - `result`.`orderReports`.items: item property `usedSor` added
  - `result`.`orderReports`.items: item property `pegOffsetType` added
  - `result`.`orderReports`.items: item property `preventedMatchId` added
  - `result`.`orderReports`.items: item property `preventedQuantity` added
  - `result`.`orderReports`.items: item property `strategyId` added
  - `result`.`orderReports`.items: item property `pegPriceType` added
  - `result`.`orderReports`.items: item property `peggedPrice` added

- Modified response for `orderListPlaceOco()` (`orderList.place.oco` method):
  - `result`.`orderReports`.items: property `pegOffsetType` added
  - `result`.`orderReports`.items: property `preventedQuantity` added
  - `result`.`orderReports`.items: property `expiryReason` added
  - `result`.`orderReports`.items: property `strategyType` added
  - `result`.`orderReports`.items: property `preventedMatchId` added
  - `result`.`orderReports`.items: property `pegOffsetValue` added
  - `result`.`orderReports`.items: property `trailingTime` added
  - `result`.`orderReports`.items: property `strategyId` added
  - `result`.`orderReports`.items: property `pegPriceType` added
  - `result`.`orderReports`.items: property `icebergQty` added
  - `result`.`orderReports`.items: property `usedSor` added
  - `result`.`orderReports`.items: property `trailingDelta` added
  - `result`.`orderReports`.items: property `workingFloor` added
  - `result`.`orderReports`.items: property `peggedPrice` added
  - `result`.`orderReports`.items: item property `pegOffsetType` added
  - `result`.`orderReports`.items: item property `preventedQuantity` added
  - `result`.`orderReports`.items: item property `expiryReason` added
  - `result`.`orderReports`.items: item property `strategyType` added
  - `result`.`orderReports`.items: item property `preventedMatchId` added
  - `result`.`orderReports`.items: item property `pegOffsetValue` added
  - `result`.`orderReports`.items: item property `trailingTime` added
  - `result`.`orderReports`.items: item property `strategyId` added
  - `result`.`orderReports`.items: item property `pegPriceType` added
  - `result`.`orderReports`.items: item property `icebergQty` added
  - `result`.`orderReports`.items: item property `usedSor` added
  - `result`.`orderReports`.items: item property `trailingDelta` added
  - `result`.`orderReports`.items: item property `workingFloor` added
  - `result`.`orderReports`.items: item property `peggedPrice` added

- Modified response for `orderListPlaceOpo()` (`orderList.place.opo` method):
  - `result`.`orderReports`.items: property `preventedMatchId` added
  - `result`.`orderReports`.items: property `trailingTime` added
  - `result`.`orderReports`.items: property `pegOffsetValue` added
  - `result`.`orderReports`.items: property `expiryReason` added
  - `result`.`orderReports`.items: property `peggedPrice` added
  - `result`.`orderReports`.items: property `preventedQuantity` added
  - `result`.`orderReports`.items: property `stopPrice` added
  - `result`.`orderReports`.items: property `trailingDelta` added
  - `result`.`orderReports`.items: property `icebergQty` added
  - `result`.`orderReports`.items: property `strategyType` added
  - `result`.`orderReports`.items: property `pegOffsetType` added
  - `result`.`orderReports`.items: property `strategyId` added
  - `result`.`orderReports`.items: property `usedSor` added
  - `result`.`orderReports`.items: property `workingFloor` added
  - `result`.`orderReports`.items: property `pegPriceType` added
  - `result`.`orderReports`.items: item property `preventedMatchId` added
  - `result`.`orderReports`.items: item property `trailingTime` added
  - `result`.`orderReports`.items: item property `pegOffsetValue` added
  - `result`.`orderReports`.items: item property `expiryReason` added
  - `result`.`orderReports`.items: item property `peggedPrice` added
  - `result`.`orderReports`.items: item property `preventedQuantity` added
  - `result`.`orderReports`.items: item property `stopPrice` added
  - `result`.`orderReports`.items: item property `trailingDelta` added
  - `result`.`orderReports`.items: item property `icebergQty` added
  - `result`.`orderReports`.items: item property `strategyType` added
  - `result`.`orderReports`.items: item property `pegOffsetType` added
  - `result`.`orderReports`.items: item property `strategyId` added
  - `result`.`orderReports`.items: item property `usedSor` added
  - `result`.`orderReports`.items: item property `workingFloor` added
  - `result`.`orderReports`.items: item property `pegPriceType` added

- Modified response for `orderListPlaceOpoco()` (`orderList.place.opoco` method):
  - `result`.`orderReports`.items: property `peggedPrice` added
  - `result`.`orderReports`.items: property `preventedMatchId` added
  - `result`.`orderReports`.items: property `icebergQty` added
  - `result`.`orderReports`.items: property `trailingDelta` added
  - `result`.`orderReports`.items: property `pegPriceType` added
  - `result`.`orderReports`.items: property `pegOffsetType` added
  - `result`.`orderReports`.items: property `workingFloor` added
  - `result`.`orderReports`.items: property `usedSor` added
  - `result`.`orderReports`.items: property `strategyType` added
  - `result`.`orderReports`.items: property `trailingTime` added
  - `result`.`orderReports`.items: property `expiryReason` added
  - `result`.`orderReports`.items: property `strategyId` added
  - `result`.`orderReports`.items: property `preventedQuantity` added
  - `result`.`orderReports`.items: property `pegOffsetValue` added
  - `result`.`orderReports`.items: item property `peggedPrice` added
  - `result`.`orderReports`.items: item property `preventedMatchId` added
  - `result`.`orderReports`.items: item property `icebergQty` added
  - `result`.`orderReports`.items: item property `trailingDelta` added
  - `result`.`orderReports`.items: item property `pegPriceType` added
  - `result`.`orderReports`.items: item property `pegOffsetType` added
  - `result`.`orderReports`.items: item property `workingFloor` added
  - `result`.`orderReports`.items: item property `usedSor` added
  - `result`.`orderReports`.items: item property `strategyType` added
  - `result`.`orderReports`.items: item property `trailingTime` added
  - `result`.`orderReports`.items: item property `expiryReason` added
  - `result`.`orderReports`.items: item property `strategyId` added
  - `result`.`orderReports`.items: item property `preventedQuantity` added
  - `result`.`orderReports`.items: item property `pegOffsetValue` added

- Modified response for `orderListPlaceOto()` (`orderList.place.oto` method):
  - `result`.`orderReports`.items: property `peggedPrice` added
  - `result`.`orderReports`.items: property `strategyId` added
  - `result`.`orderReports`.items: property `strategyType` added
  - `result`.`orderReports`.items: property `expiryReason` added
  - `result`.`orderReports`.items: property `preventedMatchId` added
  - `result`.`orderReports`.items: property `workingFloor` added
  - `result`.`orderReports`.items: property `icebergQty` added
  - `result`.`orderReports`.items: property `trailingDelta` added
  - `result`.`orderReports`.items: property `pegOffsetValue` added
  - `result`.`orderReports`.items: property `stopPrice` added
  - `result`.`orderReports`.items: property `preventedQuantity` added
  - `result`.`orderReports`.items: property `trailingTime` added
  - `result`.`orderReports`.items: property `pegOffsetType` added
  - `result`.`orderReports`.items: property `usedSor` added
  - `result`.`orderReports`.items: property `pegPriceType` added
  - `result`.`orderReports`.items: item property `peggedPrice` added
  - `result`.`orderReports`.items: item property `strategyId` added
  - `result`.`orderReports`.items: item property `strategyType` added
  - `result`.`orderReports`.items: item property `expiryReason` added
  - `result`.`orderReports`.items: item property `preventedMatchId` added
  - `result`.`orderReports`.items: item property `workingFloor` added
  - `result`.`orderReports`.items: item property `icebergQty` added
  - `result`.`orderReports`.items: item property `trailingDelta` added
  - `result`.`orderReports`.items: item property `pegOffsetValue` added
  - `result`.`orderReports`.items: item property `stopPrice` added
  - `result`.`orderReports`.items: item property `preventedQuantity` added
  - `result`.`orderReports`.items: item property `trailingTime` added
  - `result`.`orderReports`.items: item property `pegOffsetType` added
  - `result`.`orderReports`.items: item property `usedSor` added
  - `result`.`orderReports`.items: item property `pegPriceType` added

- Modified response for `orderListPlaceOtoco()` (`orderList.place.otoco` method):
  - `result`.`orderReports`.items: property `pegPriceType` added
  - `result`.`orderReports`.items: property `preventedMatchId` added
  - `result`.`orderReports`.items: property `workingFloor` added
  - `result`.`orderReports`.items: property `icebergQty` added
  - `result`.`orderReports`.items: property `trailingTime` added
  - `result`.`orderReports`.items: property `pegOffsetValue` added
  - `result`.`orderReports`.items: property `pegOffsetType` added
  - `result`.`orderReports`.items: property `trailingDelta` added
  - `result`.`orderReports`.items: property `expiryReason` added
  - `result`.`orderReports`.items: property `preventedQuantity` added
  - `result`.`orderReports`.items: property `strategyType` added
  - `result`.`orderReports`.items: property `usedSor` added
  - `result`.`orderReports`.items: property `peggedPrice` added
  - `result`.`orderReports`.items: property `strategyId` added
  - `result`.`orderReports`.items: item property `pegPriceType` added
  - `result`.`orderReports`.items: item property `preventedMatchId` added
  - `result`.`orderReports`.items: item property `workingFloor` added
  - `result`.`orderReports`.items: item property `icebergQty` added
  - `result`.`orderReports`.items: item property `trailingTime` added
  - `result`.`orderReports`.items: item property `pegOffsetValue` added
  - `result`.`orderReports`.items: item property `pegOffsetType` added
  - `result`.`orderReports`.items: item property `trailingDelta` added
  - `result`.`orderReports`.items: item property `expiryReason` added
  - `result`.`orderReports`.items: item property `preventedQuantity` added
  - `result`.`orderReports`.items: item property `strategyType` added
  - `result`.`orderReports`.items: item property `usedSor` added
  - `result`.`orderReports`.items: item property `peggedPrice` added
  - `result`.`orderReports`.items: item property `strategyId` added

- Modified response for `referencePrice()` (`referencePrice` method):
  - property `rateLimits` added

- Modified response for `referencePriceCalculation()` (`referencePrice.calculation` method):
  - property `rateLimits` added

- Modified response for `sorOrderPlace()` (`sor.order.place` method):
  - `result`.items: property `pegOffsetType` added
  - `result`.items: property `preventedMatchId` added
  - `result`.items: property `trailingDelta` added
  - `result`.items: property `icebergQty` added
  - `result`.items: property `peggedPrice` added
  - `result`.items: property `stopPrice` added
  - `result`.items: property `strategyType` added
  - `result`.items: property `trailingTime` added
  - `result`.items: property `pegOffsetValue` added
  - `result`.items: property `expiryReason` added
  - `result`.items: property `preventedQuantity` added
  - `result`.items: property `strategyId` added
  - `result`.items: property `pegPriceType` added
  - `result`.items: item property `pegOffsetType` added
  - `result`.items: item property `preventedMatchId` added
  - `result`.items: item property `trailingDelta` added
  - `result`.items: item property `icebergQty` added
  - `result`.items: item property `peggedPrice` added
  - `result`.items: item property `stopPrice` added
  - `result`.items: item property `strategyType` added
  - `result`.items: item property `trailingTime` added
  - `result`.items: item property `pegOffsetValue` added
  - `result`.items: item property `expiryReason` added
  - `result`.items: item property `preventedQuantity` added
  - `result`.items: item property `strategyId` added
  - `result`.items: item property `pegPriceType` added

- Modified response for `ticker()` (`ticker` method):
  - oneOf modified

- Modified response for `ticker24hr()` (`ticker.24hr` method):
  - oneOf modified

- Modified response for `tickerBook()` (`ticker.book` method):
  - oneOf modified

- Modified response for `tickerPrice()` (`ticker.price` method):
  - oneOf modified

- Modified response for `uiKlines()` (`uiKlines` method):
  - `result`.items: minItems `0` → `12`
  - `result`.items: maxItems `null` → `12`
  - `result`.items.items: oneOf added 2 schema(s)
  - `result`.items.items: oneOf removed 2 schema(s)

- Marked `orderListPlace()` (`orderList.place` method) as deprecated.

#### WebSocket Streams

- Modified parameter `updateSpeed`:
  - enum added: `100ms`
  - affected methods:
    - `partialBookDepth()` (`<symbol>@depth<levels>@<updateSpeed>` stream)
    - `diffBookDepth()` (`<symbol>@depth@<updateSpeed>` stream)
- Modified response for `partialBookDepth()` (`<symbol>@depth<levels>@<updateSpeed>` stream):
  - `asks`.items: minItems `0` → `2`
  - `asks`.items: maxItems `null` → `2`
  - `bids`.items: minItems `0` → `2`
  - `bids`.items: maxItems `null` → `2`

- Modified response for `diffBookDepth()` (`<symbol>@depth@<updateSpeed>` stream):
  - `a`.items: minItems `0` → `2`
  - `a`.items: maxItems `null` → `2`
  - `b`.items: minItems `0` → `2`
  - `b`.items: maxItems `null` → `2`

## 1.9.0 - 2026-06-16

### Changed (1)

- Updated dependencies

## 1.8.0 - 2026-05-20

### Added (3)

#### REST API

- `historicalBlockTrades()` (`GET /api/v3/historicalBlockTrades`)

#### WebSocket API

- `blockTradesHistorical()` (`blockTrades.historical` method)

#### WebSocket Streams

- `blockTrade()` (`<symbol>@blockTrade` stream)

### Changed (3)

#### REST API

- Modified parameter `selfTradePreventionMode`:
  - enum added: `TRANSFER`
  - affected methods:
    - `newOrder()` (`POST /api/v3/order`)
    - `orderCancelReplace()` (`POST /api/v3/order/cancelReplace`)
    - `orderOco()` (`POST /api/v3/order/oco`)
    - `orderTest()` (`POST /api/v3/order/test`)
    - `orderListOco()` (`POST /api/v3/orderList/oco`)
    - `orderListOpo()` (`POST /api/v3/orderList/opo`)
    - `orderListOpoco()` (`POST /api/v3/orderList/opoco`)
    - `orderListOto()` (`POST /api/v3/orderList/oto`)
    - `orderListOtoco()` (`POST /api/v3/orderList/otoco`)
    - `sorOrder()` (`POST /api/v3/sor/order`)
    - `sorOrderTest()` (`POST /api/v3/sor/order/test`)
- Marked `orderOco()` (`POST /api/v3/order/oco`) as deprecated.

#### WebSocket API

- Modified parameter `selfTradePreventionMode`:
  - enum added: `TRANSFER`
  - affected methods:
    - `orderCancelReplace()` (`order.cancelReplace` method)
    - `orderPlace()` (`order.place` method)
    - `orderTest()` (`order.test` method)
    - `orderListPlace()` (`orderList.place` method)
    - `orderListPlaceOco()` (`orderList.place.oco` method)
    - `orderListPlaceOpo()` (`orderList.place.opo` method)
    - `orderListPlaceOpoco()` (`orderList.place.opoco` method)
    - `orderListPlaceOto()` (`orderList.place.oto` method)
    - `orderListPlaceOtoco()` (`orderList.place.otoco` method)
    - `sorOrderPlace()` (`sor.order.place` method)
    - `sorOrderTest()` (`sor.order.test` method)

## 1.7.0 - 2026-04-20

### Changed (1)

#### WebSocket API

- Modified response for `referencePrice()` (`referencePrice` method):
  - `result`: property `code` added
  - `result`: property `msg` added

## 1.6.0 - 2026-03-26

### Changed (1)

- Updated `github.com/binance/binance-connector-go/common/v2` library to version `2.2.0`.

## 1.5.0 - 2026-03-16

### Added (7)

#### REST API

- `executionRules()` (`GET /api/v3/executionRules`)
- `referencePrice()` (`GET /api/v3/referencePrice`)
- `referencePriceCalculation()` (`GET /api/v3/referencePrice/calculation`)

#### WebSocket API

- `executionRules()` (`executionRules` method)
- `referencePrice()` (`referencePrice` method)
- `referencePriceCalculation()` (`referencePrice.calculation` method)

#### WebSocket Streams

- `referencePrice()` (`<symbol>@referencePrice` stream)

### Changed (1)

- Updated `github.com/binance/binance-connector-go/common/v2` library to version `2.1.0`.

## 1.4.0 - 2026-02-13

### Changed (1)

- Updated `github.com/binance/binance-connector-go/common/v2` library to version `2.0.0`.

## 1.3.0 - 2026-01-29

### Changed (1)

#### WebSocket API

- Added parameter `recvWindow`
  - affected methods:
    - `userDataStreamSubscribeSignature()` (`userDataStream.subscribe.signature` method)

## 1.2.0 - 2026-01-23

### Changed (1)

- Updated `github.com/binance/binance-connector-go/common` library to version `1.2.0`.

## 1.1.0 - 2026-01-13

### Changed (1)

- Updated `github.com/binance/binance-connector-go/common` library to version `1.1.0`.

## 1.0.0 - 2025-12-17

- Initial release 
