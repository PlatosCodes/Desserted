// src/views/MainGameView.js
import React, { useEffect, useContext } from 'react';
import { Container, Typography, CircularProgress, Alert } from '@mui/material';
import Gameboard from './Gameboard';
import { useGame } from '../context/GameContext';
import apiService from '../services/apiService';
import { connectWebSocket, closeWebSocket, sendMessage } from '../services/websocketService';
import Cookie from 'js-cookie';

const MainGameView = () => {
    const { gameState, updateGameState } = useContext(useGame);

    useEffect(() => {
        const fetchGameData = async () => {
            try {
                const data = await apiService.getGameDetails();
                updateGameState(data);
            } catch (error) {
                console.error('Failed to fetch game data:', error);
                updateGameState({ error: 'Failed to load game data.' });
            }
        };

        if (!gameState) {
            fetchGameData();
        }

        // Connect to WebSocket
        const token = Cookie.get('access_token');
        const ws = connectWebSocket(token, (event) => {
            const message = JSON.parse(event.data);
            updateGameState(message);
        });

        // Clean up on unmount
        return () => {
            closeWebSocket();
        };
    }, [gameState, updateGameState]);

    if (!gameState) {
        return <Container><CircularProgress /></Container>;
    }

    if (gameState.error) {
        return <Container><Alert severity="error">{gameState.error}</Alert></Container>;
    }

    return (
        <Container>
            <Typography variant="h4" gutterBottom>Main Game</Typography>
            <Gameboard gameState={gameState} />
        </Container>
    );
};

export default MainGameView;
