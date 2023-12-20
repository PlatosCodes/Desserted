// src/views/GameInvitesView.js
import React, { useState, useEffect } from 'react';
import { Container, Typography, CircularProgress, Button, List, ListItem, ListItemText, Alert } from '@mui/material';
import apiService from '../services/apiService';

const GameInvitesView = () => {
  const [gameInvites, setGameInvites] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    const fetchGameInvites = async () => {
      try {
        const data = await apiService.listGameInvites(); // Assuming this API call exists
        setGameInvites(data);
      } catch (err) {
        setError(err.message || 'Error fetching game invites.');
      } finally {
        setLoading(false);
      }
    };

    fetchGameInvites();
  }, []);

  const handleAcceptInvite = async (inviteId) => {
    try {
      await apiService.acceptGameInvite({ inviteId });
      setGameInvites(prev => prev.filter(invite => invite.id !== inviteId));
    } catch (err) {
      setError(err.message || 'Error accepting invite.');
    }
  };

  if (loading) return <CircularProgress />;

  return (
    <Container>
      <Typography variant="h4">Game Invites</Typography>
      {error && <Alert severity="error">{error}</Alert>}
      <List>
        {gameInvites.map(invite => (
          <ListItem key={invite.id}>
            <ListItemText primary={`Game invite from ${invite.senderName}`} />
            <Button variant="contained" color="primary" onClick={() => handleAcceptInvite(invite.id)}>
              Accept
            </Button>
          </ListItem>
        ))}
      </List>
    </Container>
  );
};

export default GameInvitesView;
