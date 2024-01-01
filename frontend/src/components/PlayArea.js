// src/components/PlayArea.js
import React, { useState } from 'react';
import { Paper, Typography, Button, Chip, Stack, Select, MenuItem, FormControl, InputLabel } from '@mui/material';
import apiService from '../services/apiService';
import { sendMessage } from '../services/websocketService';

const PlayArea = ({ playerGameId, selectedCards, setSelectedCards, setPlayerHand, playerHand, currentPlayerTurn, playerNumber}) => {
    const [dessertName, setDessertName] = useState('');
    const [errorMessage, setErrorMessage] = useState('');

    const handlePlayDessert = async () => {
        if (!dessertName || selectedCards.length === 0) {
            setErrorMessage("Please select a dessert and at least one card");
            return;
        }
    
        try {
            // Prepare the data for the WebSocket message
            const dessertData = {
                player_game_id: parseInt(playerGameId, 10),
                dessert_name: dessertName,
                card_ids: selectedCards.map(card_id => parseInt(card_id, 10))
            };
            // Send the message through WebSocket
            sendMessage({ type: 'playDessert', data: dessertData });
    
            // Remove played cards from hand
            setPlayerHand(prevHand => prevHand.filter(card => !selectedCards.includes(card.card_id)));
    
            // Resetting the selected cards
            setSelectedCards([]);
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
                    <Chip key={index} label={playerHand.find(card => card.card_id === cardId).card_name} />
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
                    <MenuItem value="Pie">Pie</MenuItem>
                    <MenuItem value="Chocolate Chip Cookies">Chocolate Chip Cookies</MenuItem>
                    <MenuItem value="Cheesecake">Cheesecake</MenuItem>
                    <MenuItem value="SaffronPannaCotta">Saffron Panna Cotta</MenuItem>


                    {/* Add other dessert options here */}

                </Select>
            </FormControl>
            <Button variant="contained" color="primary" 
                onClick={handlePlayDessert} 
                disabled={playerNumber !== currentPlayerTurn} 
                style={{ marginTop: '10px' }}>
                    Play Dessert
            </Button>
            {errorMessage && <Typography color="error">{errorMessage}</Typography>}
        </Paper>
    );
};

export default PlayArea;
