// src/context/GameContext.js
import React, { createContext, useContext, useState } from 'react';

const GameContext = createContext();

export const useGame = () => useContext(GameContext);

export const GameProvider = ({ children }) => {
    const [gameState, setGameState] = useState(null);

    const updateGameState = (newState) => {
        setGameState(prevState => ({ ...prevState, ...newState }));
    };

    return (
        <GameContext.Provider value={{ gameState, updateGameState }}>
            {children}
        </GameContext.Provider>
    );
};

