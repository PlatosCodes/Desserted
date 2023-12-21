// src/views/CreateGameView.js
import React, { useState } from 'react';
import { useSelector } from 'react-redux';
import { selectUser } from '../features/user/userSlice';
import { useUserFriends } from '../hooks/useUserFriends.js';
import { useMutation, useQueryClient } from 'react-query';
import apiService from '../services/apiService';
import { Button, Container, Typography, Alert, Snackbar, Checkbox, List, ListItem, ListItemText, CircularProgress } from '@mui/material';
import { useNavigate } from 'react-router-dom';

const CreateGameView = () => {
    const user = useSelector(selectUser);
    const { data: friends, isLoading, isError, error } = useUserFriends(user.id);    
    const [selectedFriends, setSelectedFriends] = useState([]);
    const [feedback, setFeedback] = useState('');
    const navigate = useNavigate();
    const queryClient = useQueryClient();

    const createGameMutation = useMutation(apiService.createGame, {
        onSuccess: async (gameResponse) => {
            const gameId = gameResponse.game.game_id;
            await Promise.all(selectedFriends.map(username => 
                apiService.invitePlayerToGame({ inviterPlayerId: user.id, inviteeUsernames: [username], gameId })
            ));
            queryClient.invalidateQueries(['userFriends', user.id]);
            navigate('/dashboard');
        },
        onError: (error) => {
            setFeedback('Failed to create the game. Please try again.');
        }
    });

    const handleCheckboxChange = (friendUsername) => {
        setSelectedFriends(prev => {
            if (prev.includes(friendUsername)) {
                return prev.filter(username => username !== friendUsername);
            } else {
                return [...prev, friendUsername];
            }
        });
    };

    const handleCreateGame = () => {
        if (selectedFriends.length === 0) {
            setFeedback('Please select at least one friend to invite.');
            return;
        }
        createGameMutation.mutate({ created_by: user.id });
    };

    if (isLoading) return <CircularProgress />;
    if (isError) return <Alert severity="error">{error.message}</Alert>;

    return (
        <Container>
            <Typography variant="h4">Create New Game</Typography>
            {error && <Alert severity="error">{error}</Alert>}
            {friends?.length === 0 ? (
                <Typography>No friends found.</Typography>
            ) : (
                <List>
                    {friends?.map((friend, index) => (
                        <ListItem key={friend.friendshipId}>
                            <Checkbox
                                checked={selectedFriends.includes(friend.friend_username)}
                                onChange={() => handleCheckboxChange(friend.friend_username)}
                            />
                            <ListItemText primary={friend.friend_username} />
                        </ListItem>
                    ))}
                </List>
            )}
            <Button onClick={handleCreateGame} disabled={isLoading}>Create Game and Invite Friends</Button>
            {feedback && (
                <Snackbar open={Boolean(feedback)} autoHideDuration={6000} onClose={() => setFeedback('')}>
                    <Alert onClose={() => setFeedback('')} severity={createGameMutation.isError ? 'error' : 'info'}>
                        {feedback}
                    </Alert>
                </Snackbar>
            )}
        </Container>
    );
};

export default CreateGameView;