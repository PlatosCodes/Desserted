// src/views/GameInvitesView.js
import React, { useState, useEffect } from 'react';
import { Container, Typography, CircularProgress, Button, List, ListItem, ListItemText, Alert } from '@mui/material';
import apiService from '../services/apiService';
import { useSelector } from 'react-redux';
import { selectUser } from '../features/user/userSlice';

const GameInvitesView = () => {
  const [gameInvites, setGameInvites] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  const user = useSelector(selectUser);

  useEffect(() => {
    const fetchGameInvites = async () => {
      try {
        console.log('Fetching game invites for user:', user);
        const data = await apiService.listGameInvites({ user_id: user.id });
        console.log('Fetched game invites:', data);
        if (data && data.game_invite && Array.isArray(data.game_invite)) {
          setGameInvites(data.game_invite);
        }
      } catch (err) {
        console.error('Error fetching game invites:', err);
        setError(err.message || 'Error fetching game invites.');
      } finally {
        setLoading(false);
      }
    };
  
    fetchGameInvites();
  }, [user.id]);
  

  const handleAcceptInvite = async (inviteId, gameId) => {
    try {
      await apiService.acceptGameInvite({ invitee_player_id: inviteId, game_id: gameId });
      setGameInvites(prev => prev.filter(invite => invite.game_id !== gameId));
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
          <ListItem key={invite.game_id}>
            <ListItemText primary={`Game invite from player ID: ${invite.invitee_player_id} for game ID: ${invite.game_id}`} />
            <Button variant="contained" color="primary" onClick={() => handleAcceptInvite(invite.invitee_player_id, invite.game_id)}>
              Accept
            </Button>
          </ListItem>
        ))}
      </List>
    </Container>
  );
};

export default GameInvitesView;
