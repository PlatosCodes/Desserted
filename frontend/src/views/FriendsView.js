// src/views/FriendsView.js
import React from 'react';
import { useQuery } from 'react-query';
import { useSelector } from 'react-redux';
import { Container, Typography, CircularProgress, List, ListItem, ListItemText, Alert } from '@mui/material';
import { selectUser } from '../features/user/userSlice';
import { useUserFriends } from '../hooks/useUserFriends';

// Function to capitalize the first letter of a string
const capitalizeFirstLetter = (str) => {
  if (typeof str !== 'string') return str;
  return str.charAt(0).toUpperCase() + str.slice(1);
};

const FriendsView = () => {
  const user = useSelector(selectUser);
  const { data: friends, isLoading, isError, error } = useUserFriends(user.id);

  if (isLoading) return <CircularProgress />;
  if (isError) return <Alert severity="error">{error.message}</Alert>;
  return (
    <Container>
      <Typography variant="h4">Friends</Typography>
      {error && <Alert severity="error">{error}</Alert>}
      {friends?.length === 0 ? (
        <Typography>No friends found.</Typography>
      ) : (
        <List>
          {friends?.map((friend, index) => (
            <ListItem key={friend.friendshipId || index}>
              <ListItemText primary={`Friend ${index + 1}: ${capitalizeFirstLetter(friend.friend_username)}`} />
            </ListItem>
          ))}
        </List>
      )}
    </Container>
  );
};

export default FriendsView;