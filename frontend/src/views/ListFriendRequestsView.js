// src/views/ListFriendRequestsView.js
import React from 'react';
import { useQuery, useMutation, useQueryClient } from 'react-query';
import { Container, Typography, CircularProgress, Alert, Snackbar } from '@mui/material';
import { useSelector } from 'react-redux';
import { selectUser } from '../features/user/userSlice';
import apiService from '../services/apiService';
import FriendRequestList from '../components/FriendRequestList';

const ListFriendRequestsView = () => {
    const user = useSelector(selectUser);
    const queryClient = useQueryClient();

    // Fetching friend requests
    const { data: friendRequests, isLoading, isError, error, isSuccess } = useQuery(
        ['friendRequests', user.id],
        () => apiService.listFriendRequests(user.id),
        { enabled: !!user.id }
    );

    // Mutation for accepting friend requests
    const acceptFriendRequestMutation = useMutation(apiService.acceptFriendRequest, {
        onSuccess: (_, { friendshipId }) => {
            // Update the cache for friendRequests
            queryClient.setQueryData(['friendRequests', user.id], (oldData) => {
                return oldData.filter(request => request.friendshipId !== friendshipId);
            });
        },
    });

    // Function to handle the acceptance of a friend request
    const handleAccept = (friendshipId) => {
        acceptFriendRequestMutation.mutate({ userId: user.id, friendshipId });
    };

    if (isLoading) return <CircularProgress />;

    if (isError) return <Alert severity="error">{error.message}</Alert>;

    return (
        <Container>
            <Typography variant="h4" style={{ marginBottom: 16 }}>Friend Requests</Typography>
            {isSuccess && friendRequests.length === 0 && (
                <Typography variant="subtitle1">No new friend requests.</Typography>
            )}
            {isSuccess && friendRequests.length > 0 && (
                <FriendRequestList friendRequests={friendRequests} onAccept={handleAccept} />
            )}
            {acceptFriendRequestMutation.isError && (
                <Snackbar
                    open={acceptFriendRequestMutation.isError}
                    autoHideDuration={6000}
                    message="Error accepting friend request"
                />
            )}
        </Container>
    );
};

export default ListFriendRequestsView;
