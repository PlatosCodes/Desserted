// DashboardView.js
import React from 'react';
import { useSelector } from 'react-redux';
import { Container, Typography } from '@mui/material';
import { selectUser } from '../features/user/userSlice';
import UserProfile from './UserProfile';

const DashboardView = () => {
    const user = useSelector(selectUser);
    console.log("User:", user)

    return (
        <Container aligncontent={'center'}>
            <p></p>
            <Typography variant="h4" textAlign={'center'}>Welcome to Desserted, {user?.username}</Typography>
            <UserProfile userId={user?.id} />
            {/* I will add more components related to the user dashboard */}
        </Container>
    );
};

export default DashboardView;
