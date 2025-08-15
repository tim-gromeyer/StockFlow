import React, { useState, useEffect } from 'react';
import axios from 'axios';

const API_URL = 'http://localhost:8080';

const Portfolio = ({ refreshKey }) => {
    const [portfolio, setPortfolio] = useState([]);
    const [totalValue, setTotalValue] = useState(0);

    useEffect(() => {
        axios.get(`${API_URL}/portfolio/1`)
            .then(response => {
                setPortfolio(response.data.portfolio || []);
                setTotalValue(response.data.total_value || 0);
            })
            .catch(error => {
                console.error('Error fetching portfolio:', error);
            });
    }, [refreshKey]);

    return (
        <div>
            <h2>Your Portfolio</h2>
            <table>
                <thead>
                    <tr>
                        <th>Symbol</th>
                        <th>Quantity</th>
                    </tr>
                </thead>
                <tbody>
                    {portfolio.map(stock => (
                        <tr key={stock.StockSymbol}>
                            <td>{stock.StockSymbol}</td>
                            <td>{stock.Quantity}</td>
                        </tr>
                    ))}
                </tbody>
            </table>
            <h3>Total Value: ${totalValue.toFixed(2)}</h3>
        </div>
    );
};

export default Portfolio;
