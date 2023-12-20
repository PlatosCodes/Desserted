// src/views/FriendsView.js
import React, { useState, useEffect } from 'react';
import { useSelector } from 'react-redux';
import { Container, Typography, CircularProgress, List, ListItem, ListItemText, Alert } from '@mui/material';
import { selectUser } from '../features/user/userSlice';
import apiService from '../services/apiService';

const FriendsView = () => {
  const user = useSelector(selectUser);
  const [friends, setFriends] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    const fetchFriends = async () => {
      if (!user || !user.id) {
        setError("User not logged in or User ID is missing");
        setLoading(false);
        return;
      }
      try {
        const response = await apiService.listUserFriends({ user_id: user.id, limit: 10, offset: 0 });
        if (response && Array.isArray(response.friendships)) {
          setFriends(response.friendships);
        } else {
          setFriends([]); // Set to empty array if response is not an array
        }
      } catch (err) {
        setError(err.message || 'Error fetching friends.');
      } finally {
        setLoading(false);
      }
    };

    fetchFriends();
  }, [user]);

  if (loading) return <CircularProgress />;

  return (
    <Container>
      <Typography variant="h4">Friends</Typography>
      {error && <Alert severity="error">{error}</Alert>}
      {friends.length === 0 ? (
        <Typography>No friends found.</Typography>
      ) : (
        <List>
          {friends.map((friend, index) => (
            <ListItem key={friend.friendshipId || index}>
              <ListItemText primary={`Friend ID: ${friend.friendee_id}`} />
            </ListItem>
          ))}
        </List>
      )}
    </Container>
  );
};

export default FriendsView;
