// src/views/ActivateAccount.js
import React, { useEffect, useState } from 'react';
import { Typography, Container, Paper } from '@mui/material';
import apiService from '../services/apiService';

function ActivateAccount() {
    const [status, setStatus] = useState('Loading...');
    
    useEffect(() => {
        const params = new URLSearchParams(window.location.search);
        const userId = params.get('user_id');
        const activationToken = params.get('activation_token');

        if (!userId || !activationToken) {
            setStatus('Invalid activation link.');
            return;
        }

        apiService.post('http://localhost:8082/v1/activate', {
            user_id: parseInt(userId, 10),
            activation_token: activationToken
        }).then(response => {
            setStatus('User successfully activated!');
        }).catch(error => {
            setStatus('Failed to activate user: ' + error.message);
        });
    }, []);
    
    return (
        <Container component={Paper} sx={{ padding: 3, display: 'flex', flexDirection: 'column', alignItems: 'center', marginTop: 8, maxWidth: 'sm' }}>
            <Typography variant="h4" gutterBottom>
                Activation Status
            </Typography>
            <Typography>{status}</Typography>
        </Container>
    );
}

export default ActivateAccount;
