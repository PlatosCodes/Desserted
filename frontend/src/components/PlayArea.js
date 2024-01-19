// src/components/PlayArea.js
import React, { useState, useRef, useEffect } from 'react';
import { Paper, Typography, Button, Chip, Stack, Select, MenuItem, FormControl, InputLabel } from '@mui/material';
import { sendMessage } from '../services/websocketService';

const PlayArea = ({ game_id, playerGameId, selectedCards, setSelectedCards, setPlayerHand, playerHand, currentPlayerTurn, playerNumber}) => {
    const [dessertName, setDessertName] = useState('');
    const [errorMessage, setErrorMessage] = useState('');
    const [specialCardType, setSpecialCardType] = useState('');
    const selectedCardsRef = useRef(selectedCards);

    useEffect(() => {
        selectedCardsRef.current = selectedCards;
    }, [selectedCards]);

    const resetSelections = () => {
        setDessertName('');
        setSpecialCardType('');
    };

    const handlePlaySpecialCard = () => {
        try {
            sendMessage({
                type: 'playSpecialCard',
                data: {
                    game_id: parseInt(game_id, 10), 
                    player_game_id: parseInt(playerGameId, 10),
                    card_type: specialCardType,
                    card_id: parseInt(selectedCards[0], 10)
                }
            });

        } catch (error) {
            console.error('Error playing special card:', error);
            setErrorMessage('Failed to play special card. Please try again.');
        }
        resetSelections();
        setErrorMessage('');
    };

    const handlePlayDessert = async () => {
        if (!dessertName || selectedCards.length === 0) {
            setErrorMessage("Please select a dessert and at least one card");
            return;
        }
    
        try {
            // Prepare the data for the WebSocket message
            const dessertData = {
                game_id: parseInt(game_id, 10),
                player_game_id: parseInt(playerGameId, 10),
                dessert_name: dessertName,
                card_ids: selectedCards.map(card_id => parseInt(card_id, 10))
            };
            // Send the message through WebSocket
            sendMessage({ type: 'playDessert', data: dessertData });
            setPlayerHand(prevHand => prevHand.filter(card => !selectedCardsRef.current.includes(card.card_id)));
            setSelectedCards([]);
        } catch (error) {
            console.error('Error playing dessert:', error);
            setErrorMessage('Failed to play dessert. Please try again.');
        }
        resetSelections();
        setErrorMessage('');
    };
    

    return (
        <Paper elevation={3} style={{ padding: '20px', minHeight: '200px' }}>
            <Typography variant="h6">Play Area</Typography>
            <Stack direction="row" spacing={1}>
                {selectedCards.map((cardId, index) => {
                const card = playerHand.find(card => card.card_id === cardId);
                const cardName = card?.name || 'Unknown Card';
                return <Chip key={index} label={cardName} />;
                })}
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
                    <MenuItem value="Marble Cake">Marble Cake</MenuItem>
                    <MenuItem value="Triple Chocolate Brownies">Triple Chocolate Brownies</MenuItem>
                    <MenuItem value="Gourmet Truffles">Gourmet Truffles</MenuItem>
                    <MenuItem value="Raspberry Chocolate Cheesecake">Raspberry Chocolate Cheesecake</MenuItem>
                    <MenuItem value="Gold Leaf Cupcakes">Gold Leaf Cupcakes</MenuItem>
                </Select>
            </FormControl>
            <Button variant="contained" color="primary" 
                onClick={handlePlayDessert} 
                disabled={playerNumber !== currentPlayerTurn} 
                style={{ marginTop: '10px' }}>
                    Play Dessert
            </Button>
            <p></p>
            <FormControl fullWidth>
                <InputLabel id="special-card-select-label">Special Card</InputLabel>
                <Select
                    labelId="special-card-select-label"
                    value={specialCardType}
                    label="Special Card"
                    onChange={(e) => setSpecialCardType(e.target.value)}
                >
                    <MenuItem value="RefreshPantry">Refresh Pantry</MenuItem>
                    <MenuItem value="StealCard">Steal Card</MenuItem>
                    <MenuItem value="InstantBake">Instant Bake</MenuItem>
                    <MenuItem value="Sabotage">Sabotage</MenuItem>
                </Select>
            </FormControl>
            <Button 
                variant="contained" 
                color="primary"
                onClick={handlePlaySpecialCard}
                disabled={!specialCardType || currentPlayerTurn !== parseInt(playerNumber,10)}
                style={{ marginTop: '10px' }}>
                Play Special Card
            </Button>
            {errorMessage && <Typography color="error">{errorMessage}</Typography>}
        </Paper>
    );
};

export default PlayArea;
