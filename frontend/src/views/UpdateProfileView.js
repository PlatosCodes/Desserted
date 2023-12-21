import React, { useState } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { updateUserProfile, selectUser } from '../features/user/userSlice';
import { Container, TextField, Button, Typography, Alert } from '@mui/material';

const UpdateProfileView = () => {
    const user = useSelector(selectUser);
    const dispatch = useDispatch();
    const [formData, setFormData] = useState({ username: user.username, email: user.email });
    const [feedback, setFeedback] = useState('');

    const handleChange = (e) => {
        setFormData({ ...formData, [e.target.name]: e.target.value });
    };

    const handleSubmit = (e) => {
        e.preventDefault();
        dispatch(updateUserProfile(formData))
            .then(() => setFeedback('Profile updated successfully!'))
            .catch(() => setFeedback('Failed to update profile.'));
    };

    return (
        <Container>
            <Typography variant="h4">Update Profile</Typography>
            <form onSubmit={handleSubmit}>
                <TextField label="Username" name="username" value={formData.username} onChange={handleChange} />
                <TextField label="Email" name="email" value={formData.email} onChange={handleChange} />
                <Button type="submit">Update</Button>
            </form>
            {feedback && <Alert severity="info">{feedback}</Alert>}
        </Container>
    );
};

export default UpdateProfileView;
