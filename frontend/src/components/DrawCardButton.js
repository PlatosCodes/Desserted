// src/components/DrawCardButton.js

import React from 'react';
import { Button } from '@mui/material';
import apiService from '../services/apiService';

const DrawCardButton = ({ gameId, playerGameId }) => {
    const handleDrawCard = async () => {
        try {
            const card = await apiService.drawCard({ gameId, playerGameId });
            // Update game state with the new card
        } catch (error) {
            // Handle error
        }
    };

    return <Button onClick={handleDrawCard}>Draw Card</Button>;
};

export default DrawCardButton;
