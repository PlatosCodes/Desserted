import React, { useState } from 'react';
import { TextField, Button, Typography, Grid, Paper } from '@mui/material';
import { Helmet } from 'react-helmet';
import { useNavigate, Link } from 'react-router-dom';
import { useDispatch } from 'react-redux';
import { loginUser } from '../features/user/userSlice';
import axiosInstance from '../services/apiService';
import './../dessert-animation.css';
import { styled } from '@mui/material/styles';
import apiService from '../services/apiService';
import Cookie from 'js-cookie'

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
  const [error, setError] = useState(null);

  const handleChange = (e) => {
    setFormData({ ...formData, [e.target.name]: e.target.value });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    if (!formData.username || !formData.password) {
        setError('Please enter your username and password.');
        return;
    }

    try {
      const response = await apiService.loginUser(formData);
      const { access_token, session_id, user } = response;
      
      if (access_token && session_id) {
        Cookie.set('token', access_token);
        Cookie.set('session_id', session_id);
        dispatch(loginUser(user));
        localStorage.setItem('userData', JSON.stringify(user));
        localStorage.setItem('isAuthenticated', 'true');
        navigate('/dashboard');
      } else {
          // Handle the error if session_id or access_token is missing
          setError('Login failed. Please try again.');
      }
    } catch (err) {
        console.error(err);
        setError(err?.response?.data?.error || 'An unexpected error occurred.');
    }
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
            {error && <Typography color="error">{error}</Typography>}
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
