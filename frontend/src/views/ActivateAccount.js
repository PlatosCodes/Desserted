// ActivateAccount.js
import React, { useEffect, useState } from 'react';
import { Typography, Container, Paper, makeStyles } from '@mui/materials/styles';
import axiosInstance from '../services/apiService';

const useStyles = makeStyles((theme) => ({
  root: {
    padding: theme.spacing(3),
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
    marginTop: theme.spacing(8),
  },
}));

function ActivateAccount() {
    const classes = useStyles();
    const [status, setStatus] = useState('Loading...');
    
    useEffect(() => {
        const params = new URLSearchParams(window.location.search);
        const userId = params.get('user_id');
        const activationToken = params.get('activation_token');

        if (!userId || !activationToken) {
            setStatus('Invalid activation link.');
            return;
        }

        axiosInstance.post('/activate', {
            user_id: parseInt(userId, 10),
            activation_token: activationToken
        }).then(response => {
            setStatus('User successfully activated!');
        }).catch(error => {
            setStatus('Failed to activate user.');
        });
    }, []);
    
    return (
        <Container component={Paper} className={classes.root} maxWidth="sm">
            <Typography variant="h4" gutterBottom>
                Activation Status
            </Typography>
            <Typography>{status}</Typography>
        </Container>
    );
}

export default ActivateAccount;
