import React, { useState, useEffect } from 'react';
import { Typography, Container, Paper } from '@mui/material';
import { styled } from '@mui/material/styles';
import axiosInstance from '../services/apiService';

// Styled components
const StyledContainer = styled(Container)(({ theme }) => ({
  padding: theme.spacing(3),
  display: 'flex',
  flexDirection: 'column',
  alignItems: 'center',
  marginTop: theme.spacing(8),
}));

const UserProfile = ({ userId }) => {
    const [user, setUser] = useState(null);

    useEffect(() => {
        const fetchUser = async () => {
            try {
                const response = await axiosInstance.get(`/users/${userId}`);
                setUser(response.data);
            } catch (err) {
                console.error(err);
                // Finish handling error gracefully here...
            }
        }

        fetchUser();
    }, [userId]);

    if (!user) return <Typography>Loading...</Typography>;

    return (
        <StyledContainer component={Paper} maxWidth="sm">
            <Typography variant="h4">{user.username}'s Profile</Typography>
            <Typography variant="h6">Email: {user.email}</Typography>
        </StyledContainer>
    );
}

export default UserProfile;
