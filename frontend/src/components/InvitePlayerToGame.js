// src/components/InvitePlayerToGame.js
import React, { useState } from 'react';
import { TextField, Button } from '@mui/material';
import apiService from '../services/apiService';
import { Snackbar, CircularProgress } from '@mui/material';


const InvitePlayerToGame = ({ gameId }) => {
    const [inviteeUsername, setInviteeUsername] = useState('');
    const [inviterPlayerId] = useState('currentPlayerId'); // Replace with actual player ID from state/context
    const [loading, setLoading] = useState(false);

    const handleInvitePlayer = async () => {
        setLoading(true);
        try {
            await apiService.invitePlayerToGame({ inviterPlayerId, inviteeUsername, gameId });
            setSnackbarMessage('Friend request sent successfully!');
            setSnackbarSeverity('success');
        } catch (error) {
            setSnackbarMessage('Failed to send friend request.');
            setSnackbarSeverity('error');
        } finally {
            setLoading(false);
            setSnackbarOpen(true);
        }
    };

    const handleCloseSnackbar = () => {
        setSnackbarOpen(false);
    };

    return (
        <div>
            <TextField 
                label="Username" 
                value={inviteeUsername} 
                onChange={(e) => setInviteeUsername(e.target.value)} 
            />
            {loading ? <CircularProgress /> : <Button onClick={handleInvitePlayer}>Invite to Game</Button>}
        </div>
    );
};

export default InvitePlayerToGame;
