// src/views/Login.js
import React, { useState } from 'react';
import { useMutation } from 'react-query';
import { useDispatch } from 'react-redux';
import { loginUser } from '../features/user/userSlice';
import { TextField, Button, Typography, Grid, Paper, Alert } from '@mui/material';
import { Helmet } from 'react-helmet';
import { useNavigate, Link } from 'react-router-dom';
import apiService from '../services/apiService';
import Cookie from 'js-cookie';
import '../dessert-animation.css';
import { styled } from '@mui/material/styles';

// Styled components
const StyledGrid = styled(Grid)(({ theme }) => ({
  height: '100vh',
  backgroundImage: 'url(/lionel.jpg)',
  backgroundSize: 'cover',
  backgroundPosition: 'center',
}));

const StyledPaper = styled(Paper)(({ theme }) => ({
  padding: theme.spacing(4),
  display: 'flex',
  flexDirection: 'column',
  alignItems: 'center',
}));

const FormRow = styled('div')(({ theme }) => ({
  display: 'flex',
  justifyContent: 'center',
}));

const Footer = styled('footer')(({ theme }) => ({
  marginTop: theme.spacing(4),
}));

const Login = () => {
  const dispatch = useDispatch();
  const navigate = useNavigate();
  const [formData, setFormData] = useState({ username: '', password: '' });

  // React Query mutation for login
  const loginMutation = useMutation(apiService.loginUser, {
    onSuccess: (response) => {
      const { access_token, refresh_token, session_id, user } = response;
      Cookie.set('access_token', access_token);
      Cookie.set('refresh_token', refresh_token);
      Cookie.set('session_id', session_id);
      dispatch(loginUser({ session_id, access_token, user }));
      localStorage.setItem('userData', JSON.stringify(user));
      localStorage.setItem('isAuthenticated', 'true');
      navigate('/dashboard');
    },
    onError: (error) => {
      console.error('Login Error:', error);
    },
  });

  const handleChange = (e) => {
    setFormData({ ...formData, [e.target.name]: e.target.value });
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    if (!formData.username || !formData.password) {
      loginMutation.setError('Please enter your username and password.');
      return;
    }
    loginMutation.mutate(formData);
  };

  return (
    <StyledGrid container justifyContent="center" alignItems="center">
      <div className="dessert-animation">
        <img src="frontend/public/images/desserts.webp" width='360' align="top" alt="Desserts Animation" />
      </div>
      <Helmet><title>Login - Desserted</title></Helmet>
      <Grid item xs={12} sm={8} md={5}>
        <StyledPaper elevation={5}>
          <Typography variant="h2" align="center">Desserted</Typography>
          <main>
            <Typography variant="h4" align="center">Login</Typography>
            {loginMutation.isError && (
              <Alert severity="error">
                {loginMutation.error?.response?.data?.error || 'An unexpected error occurred.'}
              </Alert>
            )}           
            <form onSubmit={handleSubmit} style={{ width: '100%', marginTop: 1, textAlign: "center"}}>
              <FormRow>
                <TextField label="Username" name="username" value={formData.username} onChange={handleChange} sx={{ flex: 1, m: 0.5 }} />
                <TextField label="Password" name="password" type="password" value={formData.password} onChange={handleChange} sx={{ flex: 1, m: 0.5 }} />
              </FormRow>
              <Button type="submit" fullWidth variant="contained" color="primary" sx={{ mt: 3, mb: 2 }}>Login</Button>
              <Typography variant="body2" align="center" sx={{ mt: 1.5, mb: 1.5 }}>
                Don't have an account? <Link to="/register">Register</Link>
              </Typography>
            </form>
            <Typography variant="h6" align="center">The Dessert Card Game.</Typography>
          </main>
          <Footer>
            <Typography align="center">&copy; 2023 Desserted</Typography>
          </Footer>
        </StyledPaper>
      </Grid>
    </StyledGrid>
  );
};

export default Login;
