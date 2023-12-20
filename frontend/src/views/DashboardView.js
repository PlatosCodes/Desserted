// src/views/DashboardView.js
import React from 'react';
import { useSelector } from 'react-redux';
import { Container, Typography } from '@mui/material';
import { selectUser } from '../features/user/userSlice';
import UserProfile from './UserProfile';

const DashboardView = () => {
    const user = useSelector(selectUser);

    return (
        <Container aligncontent={'center'}>
            <p></p>
            <Typography variant="h4" textAlign={'center'}>Welcome to Desserted, {user?.username}</Typography>


        </Container>
    );
};

export default DashboardView;
