// src/components/GameInviteList.js
import React, { useEffect, useState } from 'react';
import apiService from '../services/apiService';
import { List, ListItem, Button, ListItemText, Snackbar } from '@mui/material';

const GameInviteList = ({ gameId }) => {
    const [gameInvites, setGameInvites] = useState([]);
    const [snackbarOpen, setSnackbarOpen] = useState(false);
    const [snackbarMessage, setSnackbarMessage] = useState('');

    useEffect(() => {
        const fetchGameInvites = async () => {
            try {
                const response = await apiService.listGamePlayers({ gameId });
                setGameInvites(response.players);
            } catch (error) {
                console.error('Error fetching game invites:', error);
            }
        };

        fetchGameInvites();
    }, [gameId]);

    const handleAcceptInvite = async (inviteId) => {
        try {
            await apiService.acceptGameInvite({ inviteId, gameId });
            setSnackbarMessage('Game invite accepted successfully!');
            setSnackbarOpen(true);
            // Optionally, remove the accepted invite from the list
            setGameInvites(prevInvites => prevInvites.filter(invite => invite.id !== inviteId));
        } catch (error) {
            setSnackbarMessage('Failed to accept game invite.');
            setSnackbarOpen(true);
        }
    };


    return (
        <>
            <List>
                {gameInvites.map(invite => (
                    <ListItem key={invite.id}>
                        <ListItemText primary={`Invite from ${invite.playerName}`} />
                        <Button onClick={() => handleAcceptInvite(invite.id)}>Accept</Button>
                    </ListItem>
                ))}
            </List>
            <Snackbar open={snackbarOpen} autoHideDuration={6000} onClose={() => setSnackbarOpen(false)} message={snackbarMessage} />
        </>
    );
};

export default GameInviteList;