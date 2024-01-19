// src/views/Login.js
import React, { useState } from 'react';
import { useMutation } from 'react-query';
import { useDispatch } from 'react-redux';
import { loginUser } from '../features/user/userSlice';
import { TextField, Button, Typography, Grid, Paper, Alert, InputAdornment } from '@mui/material';
import { Helmet } from 'react-helmet';
import { useNavigate, Link } from 'react-router-dom';
import apiService from '../services/apiService';
import Cookie from 'js-cookie';
import '../dessert-animation.css';
import { styled } from '@mui/material/styles';
import CakeIcon from '@mui/icons-material/Cake';
import VpnKeyIcon from '@mui/icons-material/VpnKey';

// Styled components
const StyledGrid = styled(Grid)(({ theme }) => ({
  height: 'calc(100vh - 64px)',
  backgroundImage: 'url(/images/background.webp)',
  backgroundSize: 'cover',
  backgroundPosition: 'center',
  display: 'flex', // Ensure flex display for centering
  justifyContent: 'center', // Center horizontally
  alignItems: 'center', // Center vertically
}));

const StyledPaper = styled(Paper)(({ theme }) => ({
  padding: theme.spacing(4),
  margin: 'auto',
  display: 'flex',
  flexDirection: 'column',
  alignItems: 'center',
  backgroundColor: 'rgba(255, 250, 245, 0.85)', // Semi-transparent off-white
  backdropFilter: 'blur(5px)', // Apply a blur effect to the background
  borderRadius: theme.shape.borderRadius, // Optional: rounded corners
  boxShadow: theme.shadows[10], // Elevate the paper with a stronger shadow
  border: '1px solid rgba(255, 255, 255, 0.7)', // Optional: subtle border
  maxWidth: 400, // Set a max width for smaller screens
}));

const StyledTextField = styled(TextField)(({ theme }) => ({
  '& label.Mui-focused': {
    color: 'chocolate', // Color of the label text when focused
  },
  '& .MuiOutlinedInput-root': {
    '& fieldset': {
      borderColor: 'tan', // Border color
    },
    '&:hover fieldset': {
      borderColor: 'goldenrod', // Border color on hover
    },
    '&.Mui-focused fieldset': {
      borderColor: 'chocolate', // Border color when focused
      boxShadow: `0 0 0 2px ${theme.palette.primary.main}`, // Glow effect when focused
    },
    backgroundColor: 'rgba(255, 248, 240, 0.9)', // Light background for input
  },
  '& .MuiInputAdornment-root .MuiSvgIcon-root': {
    color: 'tan', // Icon color
  },
}));

// Define a gradient button with a shadow
const GradientButton = styled(Button)(({ theme }) => ({
  backgroundImage: 'linear-gradient(to right, #fbc2eb 0%, #a6c1ee 100%)', // Gradient from pink to blue
  color: theme.palette.getContrastText('#a6c1ee'), // Text color for good contrast
  boxShadow: 'inset 0 2px 4px 0 rgba(0,0,0,0.2)', // Inner shadow for an embossed look
  '&:hover': {
    boxShadow: '0 4px 20px 0 rgba(0,0,0,0.25)', // Shadow on hover for depth
    backgroundImage: 'linear-gradient(to right, #fbc2eb 0%, #a1c4fd 100%)', // Slightly different gradient on hover
  },
}));

const FormRow = styled('div')(({ theme }) => ({
  display: 'flex',
  justifyContent: 'center',
}));

const Footer = styled('footer')(({ theme }) => ({
  marginTop: theme.spacing(4),
}));

// const LaceFooter = styled(Footer)(({ theme }) => ({
//   padding: theme.spacing(2),
//   background: 'linear-gradient(rgba(255, 255, 255, 0.9), rgba(255, 255, 255, 0.7))',
//   borderTop: '1px solid rgba(255, 255, 255, 0.8)',
//   position: 'relative',
//   '&::after': {
//     content: '""',
//     position: 'absolute',
//     top: 0,
//     left: 0,
//     right: 0,
//     height: 10,
//     backgroundImage: 'url("image/lace.png")', // Replace with your lace pattern
//     backgroundRepeat: 'repeat-x',
//   },
// }));

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
      <Helmet><title>Login - Desserted</title></Helmet>
      <Grid item xs={12} sm={8} md={5}>
        <StyledPaper elevation={5}>
          <Typography variant="h2" align="center" sx={{ fontFamily: 'Pacifico, cursive', color: '#6d4c41', fontSize: '4rem' }}>
            Desserted
          </Typography>
          <main>
            <Typography variant="h5" align="center" sx = {{ color: '#6d4c41' }}>Login</Typography>
            {loginMutation.isError && (
              <Alert severity="error">
                {loginMutation.error?.response?.data?.error || 'An unexpected error occurred.'}
              </Alert>
            )}
            <form onSubmit={handleSubmit} style={{ width: '100%', marginTop: 1, textAlign: "center"}}>
              <FormRow>
                <StyledTextField label="Username" name="username" value={formData.username} onChange={handleChange} sx={{ flex: 1, m: 0.5 }} InputProps={{
    startAdornment: (
      <InputAdornment position="start">
        <CakeIcon />
      </InputAdornment>
    ),
  }} />
                <StyledTextField label="Password" name="password" type="password" value={formData.password} onChange={handleChange} sx={{ flex: 1, m: 0.5 }} InputProps={{
    startAdornment: (
      <InputAdornment position="start">
        <VpnKeyIcon />
      </InputAdornment>
    ),
  }} />
              </FormRow>
              <GradientButton type="submit" fullWidth variant="contained" sx={{ mt: 3, mb: 2 }}>Login</GradientButton>
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
