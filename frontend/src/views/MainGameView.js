import React, { useEffect, useContext } from 'react';
import GameBoard from '../components/GameBoard';
import { useApi } from '../hooks/useApi';
import apiService from '../services/apiService';
import { Container, Typography, CircularProgress } from '@mui/material';
import { useGame } from '../context/GameContext';

const MainGameView = () => {
    const { gameState, updateGameState } = useGame();
    const { request } = useApi(apiService.getGameDetails);

    useEffect(() => {
        if (!gameState) {
            // Fetch initial game state
            const fetchGameData = async () => {
                const data = await request();
                updateGameState(data);
            };
            fetchGameData();
        }
    }, [gameState, request, updateGameState]);

    if (!gameState) {
        return <Container><CircularProgress /></Container>;
    }

    return (
        <Container>
            <Typography variant="h4" gutterBottom>Main Game</Typography>
            <GameBoard players={gameState.players} />
        </Container>
    );
};

export default MainGameView;
