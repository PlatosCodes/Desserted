// src/components/SendFriendRequest.js
import React, { useState } from 'react';
import { TextField, Button } from '@mui/material';
import apiService from '../services/apiService';
import { Snackbar } from '@mui/material';
import MuiAlert from '@mui/material/Alert';

const SendFriendRequest = () => {
    const [friendeeUsername, setFriendeeUsername] = useState('');
    const [userId] = useState('currentUser'); // Replace with actual user ID
    const [snackbarOpen, setSnackbarOpen] = useState(false);
    const [snackbarMessage, setSnackbarMessage] = useState('');
    const [snackbarSeverity, setSnackbarSeverity] = useState('success');

    const handleSendRequest = async () => {
        try {
            await apiService.createFriendship({ frienderId: userId, friendeeUsername });
            setSnackbarMessage('Friend request sent successfully!');
            setSnackbarSeverity('success');
        } catch (error) {
            setSnackbarMessage('Failed to send friend request.');
            setSnackbarSeverity('error');
        } finally {
            setSnackbarOpen(true);
        }
    };

    const handleCloseSnackbar = () => {
        setSnackbarOpen(false);
    };

    return (
        <>
            <div>
                <TextField 
                    label="Username" 
                    value={friendeeUsername} 
                    onChange={(e) => setFriendeeUsername(e.target.value)} 
                />
                <Button onClick={handleSendRequest}>Send Friend Request</Button>
            </div>
            <Snackbar open={snackbarOpen} autoHideDuration={6000} onClose={handleCloseSnackbar}>
                <MuiAlert onClose={handleCloseSnackbar} severity={snackbarSeverity} elevation={6} variant="filled">
                    {snackbarMessage}
                </MuiAlert>
            </Snackbar>
        </>
    );
};

export default SendFriendRequest;
