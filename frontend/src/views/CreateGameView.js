import React, { useState } from 'react';
import { useSelector } from 'react-redux';
import { selectUser } from '../features/user/userSlice';
import apiService from '../services/apiService';
import { Button, TextField, Container, Typography, Alert } from '@mui/material';
import { useNavigate } from 'react-router-dom';

const CreateGameView = () => {
    const user = useSelector(selectUser);
    const [feedback, setFeedback] = useState('');
    const navigate = useNavigate()

    const handleCreateGame = async () => {
        try {
            await apiService.createGame({ created_by: user.id });
            setFeedback('Game created successfully!');
            navigate('/dashboard')
        } catch (error) {
            setFeedback('Failed to create the game.');
        }
    };

    return (
        <Container>
            <Typography variant="h4">Create New Game</Typography>
            <TextField label="Game Creator" value={user.id}  />
            <Button onClick={handleCreateGame}>Create Game</Button>
            {feedback && <Alert severity="info">{feedback}</Alert>}
        </Container>
    );
};

export default CreateGameView;
