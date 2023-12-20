// MainGameView.js
import React, { useState, useEffect } from 'react';
import GameBoard from '../components/GameBoard';
import Score from '../components/Score';
import apiService from '../services/apiService';
import { Container, Typography, CircularProgress } from '@mui/material';

const MainGameView = () => {
    const [gameData, setGameData] = useState(null); // Renamed for clarity
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null); // State for handling errors

    useEffect(() => {
        // Fetch initial game state
        const fetchGameData = async () => {
            setLoading(true);
            try {
                const data = await apiService.getGameDetails(); // Ensure this method is implemented
                setGameData(data);
            } catch (error) {
                console.error('Error fetching game data:', error);
                setError('Failed to load game data.');
            }
            setLoading(false);
        };
        fetchGameData();
    }, []);

    if (loading) {
        return <Container><CircularProgress /></Container>;
    }

    if (error) {
        return (
            <Container>
                <Typography color="error">{error}</Typography>
            </Container>
        );
    }

    if (!gameData) {
        return (
            <Container>
                <Typography variant="h6">Game data not found.</Typography>
            </Container>
        );
    }

    return (
        <Container>
            <Typography variant="h4" gutterBottom>Main Game</Typography>
            <Score playerScore={gameData.playerScore} />
            <GameBoard players={gameData.players} />
        </Container>
    );
};

export default MainGameView;
