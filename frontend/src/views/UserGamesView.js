// src/views/UserGamesView.js
import React, { useState, useEffect } from 'react';
import { Container, Typography, CircularProgress, List, ListItem, ListItemText, Button, Alert } from '@mui/material';
import apiService from '../services/apiService';

const UserGamesView = () => {
  const [userGames, setUserGames] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    const fetchUserGames = async () => {
      try {
        const data = await apiService.listUserGames(); // Assuming this API call exists
        setUserGames(data);
      } catch (err) {
        setError(err.message || 'Error fetching games.');
      } finally {
        setLoading(false);
      }
    };

    fetchUserGames();
  }, []);

  if (loading) return <CircularProgress />;

  return (
    <Container>
      <Typography variant="h4">My Games</Typography>
      {error && <Alert severity="error">{error}</Alert>}
      <List>
        {userGames.map(game => (
          <ListItem key={game.id}>
            <ListItemText primary={game.title} secondary={`Status: ${game.status}`} />
            <Button variant="contained" color="primary">
              Join Game
            </Button>
          </ListItem>
        ))}
      </List>
    </Container>
  );
};

export default UserGamesView;
