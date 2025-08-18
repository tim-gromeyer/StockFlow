import React, { useState, useCallback } from 'react';
import './App.css';
import Balance from './components/Balance';
import Portfolio from './components/Portfolio';
import Trade from './components/Trade';
import LiveTicker from './components/LiveTicker';

function App() {
    const [refreshKey, setRefreshKey] = useState(0);

    const handleTrade = useCallback(() => {
        setRefreshKey(prevKey => prevKey + 1);
    }, []);

    return (
        <div className="App">
            <header className="App-header">
                <h1>StockFlow</h1>
            </header>
            <div className="App-body">
                <div className="left-panel">
                    <LiveTicker />
                    <Balance refreshKey={refreshKey} />
                    <Portfolio refreshKey={refreshKey} />
                </div>
                <div className="right-panel">
                    <Trade onTrade={handleTrade} />
                </div>
            </div>
        </div>
    );
}

export default App;