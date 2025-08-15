import React, { useState, useEffect } from 'react';
import axios from 'axios';

const API_URL = 'http://localhost:8080';

const Balance = ({ refreshKey }) => {
    const [balance, setBalance] = useState(0);

    useEffect(() => {
        axios.get(`${API_URL}/balance/1`)
            .then(response => {
                setBalance(response.data.balance);
            })
            .catch(error => {
                console.error('Error fetching balance:', error);
            });
    }, [refreshKey]);

    return (
        <div>
            <h2>Your Balance</h2>
            <p>${balance.toFixed(2)}</p>
        </div>
    );
};

export default Balance;
