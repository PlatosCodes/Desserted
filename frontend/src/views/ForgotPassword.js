// src/views/ForgotPassword.js
import React, { useState } from 'react';
import { TextField, Button, Typography, Container, Paper, makeStyles } from '@mui/material';
import axiosInstance from '../services/apiService';

const useStyles = makeStyles((theme) => ({
  root: {
    padding: theme.spacing(3),
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
    marginTop: theme.spacing(8),
  },
  form: {
    marginTop: theme.spacing(2),
    width: '100%',
  },
  submit: {
    marginTop: theme.spacing(3),
  },
}));

const ForgotPassword = () => {
  const classes = useStyles();
  const [email, setEmail] = useState('');
  const [feedback, setFeedback] = useState('');

  const handleForgotPassword = async () => {
    try {
      await axiosInstance.post('v1//users/forgot_password', { email });
      setFeedback('Password reset link has been sent to your email.');
    } catch (err) {
      setFeedback('Error sending reset link. Please try again.');
    }
  };

  return (
    <Container component={Paper} className={classes.root} maxWidth="sm">
      <Typography variant="h6">Forgot Password</Typography>
      <form className={classes.form} noValidate>
        <TextField 
          label="Email"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
          fullWidth
          variant="outlined"
          margin="normal"
        />
        <Button 
          onClick={handleForgotPassword}
          fullWidth
          variant="contained"
          color="primary"
          className={classes.submit}
        >
          Send Reset Link
        </Button>
        {feedback && <Typography color="secondary">{feedback}</Typography>}
      </form>
    </Container>
  );
};

export default ForgotPassword;
