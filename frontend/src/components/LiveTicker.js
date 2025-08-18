import React, { useState, useEffect } from 'react';

const LiveTicker = () => {
    const [prices, setPrices] = useState({});

    useEffect(() => {
        const ws = new WebSocket('ws://localhost:8080/ws/prices');

        ws.onopen = () => {
            console.log('WebSocket connected');
        };

        ws.onmessage = (event) => {
            const updates = JSON.parse(event.data);
            setPrices(prevPrices => {
                const newPrices = { ...prevPrices };
                updates.forEach(update => {
                    newPrices[update.symbol] = update;
                });
                return newPrices;
            });
        };

        ws.onclose = () => {
            console.log('WebSocket disconnected');
        };

        ws.onerror = (error) => {
            console.error('WebSocket error:', error);
        };

        return () => {
            ws.close();
        };
    }, []);

    return (
        <div>
            <h2>Live Stock Prices</h2>
            {Object.keys(prices).length === 0 ? (
                <p>Connecting to live updates...</p>
            ) : (
                <div>
                    {Object.entries(prices).map(([symbol, data]) => (
                        <p key={symbol}>
                            {symbol}: 
                            <span style={{
                                color: data.price > data.prevPrice ? 'green' : (data.price < data.prevPrice ? 'red' : 'black')
                            }}>
                                ${data.price.toFixed(2)}
                            </span>
                        </p>
                    ))}
                </div>
            )}
        </div>
    );
};

export default LiveTicker;
