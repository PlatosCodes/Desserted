// src/views/UserProfile.js
import React, { useState } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { TextField, Button, Paper, Typography, Container } from '@mui/material';
import { updateUserProfile } from '../features/user/userSlice';
import { selectUser } from '../features/user/userSlice';
import { styled } from '@mui/material/styles';

// Styled components
const StyledContainer = styled(Container)(({ theme }) => ({
  padding: theme.spacing(3),
  display: 'flex',
  flexDirection: 'column',
  alignItems: 'center',
  marginTop: theme.spacing(8),
}));


// Need to implement this in backend first -- but after I decide what belongs here. For now, placeholder stuff


const UserProfile = () => {
    const user = useSelector(selectUser);
    const [formData, setFormData] = useState({ username: user?.username, email: user?.email, password: '' });
    const dispatch = useDispatch();

    const handleChange = (e) => {
        setFormData({ ...formData, [e.target.name]: e.target.value });
    };

    const handleSubmit = (e) => {
        e.preventDefault();
        dispatch(updateUserProfile(formData));
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
        </Paper>
    );
};

export default UserProfile;