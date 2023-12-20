// src/components/PlayerHand.js
import React, { useEffect, useState } from 'react';
import apiService from '../services/apiService';
import Card from './Card';

const PlayerHand = ({ playerGameId }) => {
    const [hand, setHand] = useState([]);
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState('');

    useEffect(() => {
        const fetchHand = async () => {
            setLoading(true);
            try {
                const playerHand = await apiService.getPlayerHand(playerGameId);
                setHand(playerHand);
            } catch (error) {
                setError('Failed to load player hand.');
            } finally {
                setLoading(false);
            }
        };

        fetchHand();
    }, [playerGameId]);


    return (
        <div>
            {loading ? <CircularProgress /> : hand.map(card => (<Card key={card.cardId} {...card} />))}
            {error && <p>{error}</p>}
        </div>
    );
};

export default PlayerHand;
