// src/views/GameBoard.js
import React, { useState, useEffect } from 'react';
import { Container, Grid, Paper, Typography, Button } from '@mui/material';
import Hand from '../components/Hand';
import Scoreboard from '../components/Scoreboard';
import PlayArea from '../components/PlayArea';
import apiService from '../services/apiService';
import { useGame } from '../context/GameContext';

const GameBoard = () => {
    const { gameState } = useGame();
    const [playerHand, setPlayerHand] = useState([]);

    useEffect(() => {
        // Fetch player hand and update state
        const fetchPlayerHand = async () => {
            try {
                const handData = await apiService.getPlayerHand(/* playerId */);
                setPlayerHand(handData);
            } catch (error) {
                console.error('Error fetching player hand:', error);
            }
        };
        fetchPlayerHand();
    }, []);

    return (
        <Container>
            <Typography variant="h4" gutterBottom>Game Board</Typography>
            <Grid container spacing={3}>
                <Grid item xs={12} md={6}>
                    <Scoreboard players={gameState.players} />
                </Grid>
                <Grid item xs={12} md={6}>
                <PlayArea 
                    playerHand={playerHand}
                />
                </Grid>
                <Grid item xs={12}>
                    <Paper elevation={3}>
                        <Hand cards={playerHand} />
                    </Paper>
                </Grid>
            </Grid>
        </Container>
    );
};

export default GameBoard;
