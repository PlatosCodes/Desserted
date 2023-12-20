import React from 'react';
import { Button } from '@mui/material';
import apiService from '../services/apiService';

const StartGameButton = ({ gameId }) => {
    const handleStartGame = async () => {
        await apiService.startGame({ gameId });
        // Handle the game start logic, e.g., redirect to the game board
    };

    return <Button onClick={handleStartGame}>Start Game</Button>;
};

export default StartGameButton;
