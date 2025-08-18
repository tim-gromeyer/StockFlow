# WebSocket API Documentation

This document provides information about the WebSocket API for the StockFlow application.

## Overview

The WebSocket API provides real-time stock price updates to connected clients.

## Connecting to the WebSocket

To connect to the WebSocket, use the following endpoint:

```
ws://localhost:8080/ws/prices
```

## Messages

### Server-to-Client Messages

The server broadcasts stock price updates to all connected clients. The message format is a JSON object with the following structure:

```json
{
  "symbol": "AAPL",
  "price": 150.00,
  "prevPrice": 149.50
}
```

*   `symbol`: The stock symbol.
*   `price`: The current price of the stock.
*   `prevPrice`: The previous price of the stock.

### Client-to-Server Messages

The client does not send any messages to the server. The communication is one-way (server-to-client).
