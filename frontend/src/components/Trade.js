import React, { useState } from 'react';
import axios from 'axios';

const API_URL = 'http://localhost:8080';

const Trade = ({ onTrade }) => {
    const [symbol, setSymbol] = useState('AAPL');
    const [quantity, setQuantity] = useState(0);

    const handleTrade = (action) => {
        const tradeData = {
            user_id: 1,
            stock_symbol: symbol,
            quantity: parseInt(quantity, 10),
        };

        axios.post(`${API_URL}/${action}`, tradeData)
            .then(response => {
                alert('Trade successful!');
                onTrade(); // Callback to refresh data in parent
            })
            .catch(error => {
                console.error(`Error executing ${action}:`, error);
                alert(`Trade failed: ${error.response?.data?.error || 'Server error'}`);
            });
    };

    return (
        <div>
            <h2>Trade Stocks</h2>
            <div>
                <label>Stock Symbol:</label>
                <select value={symbol} onChange={e => setSymbol(e.target.value)}>
                    <option value="AAPL">AAPL</option>
                    <option value="GOOGL">GOOGL</option>
                    <option value="TSLA">TSLA</option>
                </select>
            </div>
            <div>
                <label>Quantity:</label>
                <input type="number" value={quantity} onChange={e => setQuantity(e.target.value)} />
            </div>
            <button onClick={() => handleTrade('buy')}>Buy</button>
            <button onClick={() => handleTrade('sell')}>Sell</button>
        </div>
    );
};

export default Trade;
