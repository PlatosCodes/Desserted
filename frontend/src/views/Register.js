// src/views/Register.ks
import React, { useState } from 'react';
import apiService from '../services/apiService';
import { useNavigate, Link } from 'react-router-dom';
import { TextField, Button, Grid, Container, Typography, Paper } from '@mui/material';
import { styled } from '@mui/material/styles';

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
    const [successMessage, setSuccessMessage] = useState('');
    const [error, setError] = useState(null);
    const navigate = useNavigate();

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
      
    const handleSubmit = async (e) => {
      e.preventDefault();
      const isValid = validateForm();
      if (!isValid) return;
      
      try {
          await apiService.createUser({ username, email, password });
          setSuccessMessage("Registration successful! Please check your email for a verification link.");
          setUsername('');
          setEmail('');
          setPassword('');
          navigate('/login');
      } catch (err) {
          console.log("Error:", err)
          setError(err.response && err.response.data && err.response.data.error ? err.response.data.error : 'An unexpected error occurred.');
      }
  };
    
    return (
        <StyledContainer component={Paper} maxWidth="xs">
          <Title variant="h4">Register</Title>
          {successMessage && <Typography color="primary">{successMessage}</Typography>}
          {error && <Typography color="error">{error}</Typography>}
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
