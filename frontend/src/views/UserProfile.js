// src/views/UserProfile.js
import React, { useState } from 'react';
import { useMutation } from 'react-query';
import { useDispatch, useSelector } from 'react-redux';
import { TextField, Button, Paper, Typography, Alert, Container } from '@mui/material';
import { updateUserProfile } from '../features/user/userSlice';
import { selectUser } from '../features/user/userSlice';
import apiService from '../services/apiService';
import { styled } from '@mui/material/styles';

// Styled components
const StyledContainer = styled(Container)(({ theme }) => ({
  padding: theme.spacing(3),
  display: 'flex',
  flexDirection: 'column',
  alignItems: 'center',
  marginTop: theme.spacing(8),
}));

const UserProfile = () => {
    const user = useSelector(selectUser);
    const [formData, setFormData] = useState({ username: user?.username, email: user?.email, password: '' });
    const dispatch = useDispatch();
    const updateProfileMutation = useMutation(apiService.updateUser);

    const handleChange = (e) => {
        setFormData({ ...formData, [e.target.name]: e.target.value });
    };

    const handleSubmit = (e) => {
        e.preventDefault();
        updateProfileMutation.mutate(formData);
        dispatch(updateUserProfile(formData));
        setFormData({ ...formData, password: '' });
    };

    return (
        <Paper>
            <Typography variant="h4">Update Profile</Typography>
            <form onSubmit={handleSubmit}>
                <TextField label="Username" name="username" value={formData.username} onChange={handleChange} />
                <TextField label="Email" name="email" value={formData.email} onChange={handleChange} />
                <TextField label="Password" name="password" type="password" value={formData.password} onChange={handleChange} />
                <Button type="submit">Update</Button>
            </form>
            {updateProfileMutation.isError && (
                <Alert severity="error">{updateProfileMutation.error.message}</Alert>
            )}
        </Paper>
    );
};

export default UserProfile;
        