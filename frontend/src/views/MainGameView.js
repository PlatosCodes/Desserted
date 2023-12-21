// src/views/MainGameView.js
import React, { useEffect, useContext } from 'react';
import { Container, Typography, CircularProgress, Alert } from '@mui/material';
import GameBoard from './GameBoard';
import { useGame } from '../context/GameContext';
import apiService from '../services/apiService';

const MainGameView = () => {
    const { gameState, setGameState } = useContext(useGame);

    useEffect(() => {
        const fetchGameData = async () => {
            try {
                const data = await apiService.getGameDetails();
                setGameState(data);
            } catch (error) {
                // Error handling
                console.error('Failed to fetch game data:', error);
                setGameState({ ...gameState, error: 'Failed to load game data.' });
            }
        };

        if (!gameState) {
            fetchGameData();
        }
    }, [gameState, setGameState]);

    if (!gameState) {
        return <Container><CircularProgress /></Container>;
    }

    if (gameState.error) {
        return <Container><Alert severity="error">{gameState.error}</Alert></Container>;
    }

    return (
        <Container>
            <Typography variant="h4" gutterBottom>Main Game</Typography>
            <GameBoard />
        </Container>
    );
};

export default MainGameView;
