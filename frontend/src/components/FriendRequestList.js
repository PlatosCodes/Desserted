// src/components/FriendRequestList.js
import React, { useEffect, useState } from 'react';
import apiService from '../services/apiService';
import { List, ListItem, Button, ListItemText, Snackbar } from '@mui/material';

const FriendRequestList = () => {
    const [friendRequests, setFriendRequests] = useState([]);
    const [snackbarOpen, setSnackbarOpen] = useState(false);
    const [snackbarMessage, setSnackbarMessage] = useState('');
    const [userId, setUserId] = useState(null); // Fetch from user's state

    useEffect(() => {
        const fetchFriendRequests = async () => {
            try {
                const data = await apiService.listFriendRequests(userId);
                setFriendRequests(data);
            } catch (error) {
                console.error('Error fetching friend requests:', error);
            }
        };

        fetchFriendRequests();
    }, [userId]);

    const handleAccept = async (friendshipId) => {
        try {
            await apiService.acceptFriendRequest({ userId, friendshipId });
            setSnackbarMessage('Friend request accepted successfully!');
            setSnackbarOpen(true);
            // Remove the accepted request from the list
            setFriendRequests(prevRequests => prevRequests.filter(request => request.friendshipId !== friendshipId));
        } catch (error) {
            setSnackbarMessage('Failed to accept friend request.');
            setSnackbarOpen(true);
        }
    };

    if (!friendRequests) return <p>Loading friend requests...</p>;

    return (
        <>
            <List>
                {friendRequests.map(request => (
                    <ListItem key={request.friendshipId}>
                        <ListItemText primary={request.frienderUsername} />
                        <Button onClick={() => handleAccept(request.friendshipId)}>Accept</Button>
                    </ListItem>
                ))}
            </List>
            <Snackbar open={snackbarOpen} autoHideDuration={6000} onClose={() => setSnackbarOpen(false)} message={snackbarMessage} />
        </>
    );
};

export default FriendRequestList;