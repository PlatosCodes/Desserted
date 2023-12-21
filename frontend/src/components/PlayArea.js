// src/components/PlayArea.js
import React, { useState } from 'react';
import { Paper, Typography, Button, Chip, Stack, Select, MenuItem, FormControl, InputLabel } from '@mui/material';
import apiService from '../services/apiService';

const PlayArea = ({ playerGameId, playerHand, refreshPlayerData }) => {
    const [selectedCards, setSelectedCards] = useState([]);
    const [dessertName, setDessertName] = useState('');
    const [errorMessage, setErrorMessage] = useState('');

    const handleCardSelect = (cardId) => {
        if (selectedCards.includes(cardId)) {
            setSelectedCards(selectedCards.filter(id => id !== cardId));
        } else {
            setSelectedCards([...selectedCards, cardId]);
        }
    };

    const handlePlayDessert = async () => {
        if (!dessertName || selectedCards.length === 0) {
            setErrorMessage("Please select a dessert and at least one card");
            return;
        }

        try {
            const response = await apiService.playDessert({ playerGameId, dessertName, cardIds: selectedCards });
            console.log(response);
            setSelectedCards([]);
            setDessertName('');
            refreshPlayerData();
        } catch (error) {
            console.error('Error playing dessert:', error);
            setErrorMessage('Failed to play dessert. Please try again.');
        }
    };

    return (
        <Paper elevation={3} style={{ padding: '20px', minHeight: '200px' }}>
            <Typography variant="h6">Play Area</Typography>
            <Stack direction="row" spacing={1}>
                {selectedCards.map((cardId, index) => (
                    <Chip key={index} label={playerHand.find(card => card.id === cardId).name} />
                ))}
            </Stack>
            <FormControl fullWidth>
                <InputLabel id="dessert-select-label">Dessert</InputLabel>
                <Select
                    labelId="dessert-select-label"
                    value={dessertName}
                    label="Dessert"
                    onChange={(e) => setDessertName(e.target.value)}
                >
                    {/* List of desserts */}
                    <MenuItem value="Cake">Cake</MenuItem>
                    {/* Add other dessert options here */}
                </Select>
            </FormControl>
            <Button variant="contained" color="primary" onClick={handlePlayDessert} style={{ marginTop: '10px' }}>
                Play Dessert
            </Button>
            {errorMessage && <Typography color="error">{errorMessage}</Typography>}
        </Paper>
    );
};

export default PlayArea;
