// src/components/Search.js
import React, { useState } from 'react';

const Search = () => {
    const [from, setFrom] = useState('');
    const [to, setTo] = useState('');
    const [results, setResults] = useState([]);

    const handleSearch = () => {
        const mockTransportOptions = [
            { mode: 'Bus', price: 25, from: 'New York', to: 'Washington' },
            { mode: 'Train', price: 40, from: 'New York', to: 'Washington' },
            { mode: 'Flight', price: 120, from: 'New York', to: 'Washington' },
            { mode: 'Bus', price: 30, from: 'Los Angeles', to: 'San Francisco' },
            { mode: 'Train', price: 50, from: 'Los Angeles', to: 'San Francisco' },
            { mode: 'Flight', price: 150, from: 'Los Angeles', to: 'San Francisco' }
        ];

        const filteredResults = mockTransportOptions.filter(
            option => option.from.toLowerCase() === from.toLowerCase() && option.to.toLowerCase() === to.toLowerCase()
        );
        setResults(filteredResults);
    };

    return (
        <div>
            <h2>Search Transport</h2>
            <input type="text" placeholder="From" value={from} onChange={(e) => setFrom(e.target.value)} required />
            <input type="text" placeholder="To" value={to} onChange={(e) => setTo(e.target.value)} required />
            <button onClick={handleSearch}>Search</button>
            <div>
                {results.length > 0 ? (
                    <ul>
                        {results.map((option, index) => (
                            <li key={index}>{option.mode}: ${option.price}</li>
                        ))}
                    </ul>
                ) : (
                    <p>No transport options found.</p>
                )}
            </div>
        </div>
    );
};

export default Search;