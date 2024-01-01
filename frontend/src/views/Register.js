// src/views/Register.js
import React, { useState } from 'react';
import { useMutation } from 'react-query';
import { useNavigate } from 'react-router-dom';
import { TextField, Button, Grid, Container, Typography, Paper, Alert } from '@mui/material';
import { styled } from '@mui/material/styles';
import apiService from '../services/apiService';

// Styled components
const StyledContainer = styled(Container)(({ theme }) => ({
  padding: theme.spacing(3),
  display: 'flex',
  flexDirection: 'column',
  alignItems: 'center',
}));

const Title = styled(Typography)(({ theme }) => ({
  marginBottom: theme.spacing(2),
}));

const Register = () => {
  const [username, setUsername] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');
  const navigate = useNavigate();

  const registerMutation = useMutation(apiService.createUser, {
    onSuccess: () => {
      navigate('/login');
    },
    onError: (error) => {
      // Extracting error response from Axios
      const serverResponse = error.response?.data;

      // Handling specific error types
      if (serverResponse?.field_violations) {
        // For field-specific errors
        const fieldErrors = serverResponse.field_violations.map(violation => `${violation.field}: ${violation.description}`).join(', ');
        setError(fieldErrors);
      } else if (serverResponse?.error) {
        // For general error messages from server
        setError(serverResponse.error);
      } else if (error.response?.status === 409) {
        // Handling unique constraint violations
        if (error.response.data.message.includes("users_username_key")) {
          setError('Username already exists.');
        } else if (error.response.data.message.includes("users_email_key")) {
          setError('Email already registered.');
        } else {
          setError('An unexpected error occurred.');
        }
      } else if (error.response?.status === 400) {
        setError('Invalid password. Please make sure password is between 6 and 100 characters.');
      } else {
        setError('An unexpected error occurred.');
      }
    },
  });

  const validateForm = () => {
    if (!username || !email || !password) {
      setError("All fields are required!");
      return false;
    }
    const emailRegex = /^[\w-]+(\.[\w-]+)*@([\w-]+\.)+[a-zA-Z]{2,7}$/;
    if (!email.match(emailRegex)) {
      setError("Please enter a valid email!");
      return false;
    }
    return true;
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    if (validateForm()) {
      registerMutation.mutate({ username, email, password });
    }
  };

  return (
    <StyledContainer component={Paper} maxWidth="xs">
      <Title variant="h4">Register</Title>
      {error && <Alert severity="error">{error}</Alert>}
      <Grid container spacing={3}>
        <Grid item xs={12}>
          <TextField fullWidth label="Username" variant="outlined" value={username} onChange={(e) => setUsername(e.target.value)} />
        </Grid>
        <Grid item xs={12}>
          <TextField fullWidth label="Password" variant="outlined" type="password" value={password} onChange={(e) => setPassword(e.target.value)} />
        </Grid>
        <Grid item xs={12}>
          <TextField fullWidth label="Email" variant="outlined" type="email" value={email} onChange={(e) => setEmail(e.target.value)} />
        </Grid>
        <Grid item xs={12}>
          <Button variant="contained" color="primary" onClick={handleSubmit}>Register</Button>
        </Grid>
      </Grid>
    </StyledContainer>
  );
}

export default Register;
