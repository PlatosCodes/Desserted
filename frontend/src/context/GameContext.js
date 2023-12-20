// src/context/GameContext.js
import React, { createContext, useContext, useState, useEffect } from 'react';
import apiService from '../services/apiService';

const GameContext = createContext();

export const useGame = () => useContext(GameContext);

export const GameProvider = ({ children }) => {
    const [gameState, setGameState] = useState(null);

    useEffect(() => {
        const fetchGameDetails = async () => {
            try {
                const gameData = await apiService.getGameDetails(); // Assuming this API call exists
                setGameState(gameData);
            } catch (error) {
                console.error('Error fetching game details:', error);
            }
        };

        fetchGameDetails();
    }, []);

    return (
        <GameContext.Provider value={{ gameState, setGameState }}>
            {children}
        </GameContext.Provider>
    );
};
