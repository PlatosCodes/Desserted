// src/components/ErrorFallback.js
import React from 'react';
import { Button, Container, Typography } from '@mui/material';

function ErrorFallback({ error, resetErrorBoundary }) {
    return (
        <Container>
            <Typography variant="h5" color="error">Something went wrong:</Typography>
            <Typography variant="subtitle1">{error.message}</Typography>
            <Button variant="contained" color="primary" onClick={resetErrorBoundary}>
                Try Again
            </Button>
        </Container>
    );
}

export default ErrorFallback;
