import React, { useState, useEffect } from 'react';
import { useSelector } from 'react-redux';
import { selectUser } from '../features/user/userSlice';
import apiService from '../services/apiService';
import { Button, Container, Typography, Alert, Checkbox, List, ListItem, ListItemText, CircularProgress } from '@mui/material';
import { useNavigate } from 'react-router-dom';

const CreateGameView = () => {
    const user = useSelector(selectUser);
    const [feedback, setFeedback] = useState('');
    const [friends, setFriends] = useState([]);
    const [selectedFriends, setSelectedFriends] = useState([]);
    const [loading, setLoading] = useState(true);
    const navigate = useNavigate();

    useEffect(() => {
        const fetchFriends = async () => {
            if (!user || !user.id) {
                setFeedback("User not logged in or User ID is missing");
                setLoading(false);
                return;
            }
            try {
                const response = await apiService.listUserFriends({ user_id: user.id, limit: 10, offset: 0 });
                setFriends(response.friendships || []);
            } catch (err) {
                setFeedback(err.message || 'Error fetching friends.');
            } finally {
                setLoading(false);
            }
        };

        fetchFriends();
    }, [user]);

    const handleCheckboxChange = (friendUsername) => {
        setSelectedFriends(prev => {
            if (prev.includes(friendUsername)) {
                return prev.filter(username => username !== friendUsername);
            } else {
                return [...prev, friendUsername];
            }
        });
    };

    const handleCreateGame = async () => {
        if (selectedFriends.length === 0) {
            setFeedback('Please select at least one friend to invite.');
            return;
        }

        setLoading(true);
        try {
            const gameResponse = await apiService.createGame({ created_by: user.id });
            const gameId = gameResponse.game.game_id; // Adjust according to actual response structure
            await Promise.all(selectedFriends.map(username => 
                apiService.invitePlayerToGame({ inviterPlayerId: user.id, inviteeUsernames: [username], gameId: gameId })
            ));
            setFeedback('Game created and friends invited successfully!');
            navigate('/dashboard');
        } catch (error) {
            setFeedback('Failed to create the game or invite friends.');
        } finally {
            setLoading(false);
        }
    };

    if (loading) return <CircularProgress />;

    return (
        <Container>
            <Typography variant="h4">Create New Game</Typography>
            <List>
                {friends.map(friend => (
                    <ListItem key={friend.friendshipId}>
                        <Checkbox
                            checked={selectedFriends.includes(friend.friend_username)}
                            onChange={() => handleCheckboxChange(friend.friend_username)}
                        />
                        <ListItemText primary={friend.friend_username} />
                    </ListItem>
                ))}
            </List>
            <Button onClick={handleCreateGame} disabled={loading}>Create Game and Invite Friends</Button>
            {feedback && <Alert severity="info">{feedback}</Alert>}
        </Container>
    );
};

export default CreateGameView;
