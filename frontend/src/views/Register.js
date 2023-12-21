// src/views/Register.ks
import React, { useState } from 'react';
import { useMutation } from 'react-query';
import { useNavigate, Link } from 'react-router-dom';
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
  const navigate = useNavigate();

  const registerMutation = useMutation(apiService.createUser, {
      onSuccess: () => {
          navigate('/login');
      },
      onError: (error) => {
          console.error("Error:", error);
      },
  });

  const validateForm = () => {
      if (!username || !email || !password) {
        registerMutation.setError("All fields are required!");
        return false;
      }
      const emailRegex = /^[\w-]+(\.[\w-]+)*@([\w-]+\.)+[a-zA-Z]{2,7}$/;
      if (!email.match(emailRegex)) {
        registerMutation.setError("Please enter a valid email!");
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
          {registerMutation.isError && <Alert severity="error">{registerMutation.error.message}</Alert>}
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
