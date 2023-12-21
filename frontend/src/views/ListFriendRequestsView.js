// src/views/ListFriendRequestsView.js
import React, { useEffect, useState } from 'react';
import apiService from '../services/apiService';
import FriendRequestList from '../components/FriendRequestList';
import { Snackbar, Alert } from '@mui/material';
import { useSelector } from 'react-redux'
import { selectUser } from '../features/user/userSlice';

const ListFriendRequestsView = () => {
    const [friendRequests, setFriendRequests] = useState([]);
    const [snackbarOpen, setSnackbarOpen] = useState(false);
    const [snackbarMessage, setSnackbarMessage] = useState('');
    const [snackbarSeverity, setSnackbarSeverity] = useState('info');

    const user = useSelector(selectUser);

    useEffect(() => {
        const fetchFriendRequests = async () => {
            try {
                const response = await apiService.listFriendRequests(user.id);
                if (response && Array.isArray(response)) { 
                    setFriendRequests(response);
                }
            } catch (error) {
                console.error('Error fetching friend requests:', error);
                setSnackbarMessage('Failed to fetch friend requests.');
                setSnackbarSeverity('error');
                setSnackbarOpen(true);
            }
        };
    
        fetchFriendRequests();
    }, [user.id]);
    

    const handleAccept = async (friendshipId) => {
        try {
            await apiService.acceptFriendRequest({ userId: user.id, friendshipId });
            setSnackbarMessage('Friend request accepted successfully!');
            setSnackbarSeverity('success');
            setFriendRequests(prevRequests => prevRequests.filter(request => request.friendshipId !== friendshipId));
        } catch (error) {
            setSnackbarMessage('Failed to accept friend request.');
            setSnackbarSeverity('error');
        }
        setSnackbarOpen(true);
    };

    return (
        <>
            <FriendRequestList friendRequests={friendRequests} onAccept={handleAccept} />
            <Snackbar open={snackbarOpen} autoHideDuration={6000} onClose={() => setSnackbarOpen(false)}>
                <Alert onClose={() => setSnackbarOpen(false)} severity={snackbarSeverity} sx={{ width: '100%' }}>
                    {snackbarMessage}
                </Alert>
            </Snackbar>
        </>
    );
};

export default ListFriendRequestsView;
