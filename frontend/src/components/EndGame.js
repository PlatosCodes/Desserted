// src/components/EndGame.js
import React from 'react';
import { Box, Typography, Button, useTheme } from '@mui/material';
import { motion } from 'framer-motion';
import Confetti from 'react-confetti';
// 
const EndGame = ({ winner, winningScore, winningMessage }) => {
    const theme = useTheme();
    // const [width, height] = "auto";

    const containerVariants = {
        hidden: { opacity: 0, y: 200 },
        visible: { opacity: 1, y: 0, transition: { delay: 0.5, duration: 0.8, type: 'spring' } }
    };

    const buttonVariants = {
        hover: { scale: 1.05, transition: { duration: 0.3 } },
        tap: { scale: 0.95 }
    };

    return (
        <Box
            component={motion.div}
            variants={containerVariants}
            initial="hidden"
            animate="visible"
            sx={{
                position: 'relative',
                display: 'flex',
                flexDirection: 'column',
                alignItems: 'center',
                justifyContent: 'center',
                minHeight: '100vh',
                backgroundColor: theme.palette.background.default,
                textAlign: 'center',
                padding: theme.spacing(3)
            }}
        >
            <Confetti  />
            <Typography variant="h2" gutterBottom>
                Game Over
            </Typography>
            <Typography variant="h4" gutterBottom>
                {`Congratulations, Player ${winner}!`}
            </Typography>
            <Typography variant="h5">
                {`Final Score: ${winningScore}`}
            </Typography>
            <Button
                component={motion.button}
                variants={buttonVariants}
                whileHover="hover"
                whileTap="tap"
                color="primary"
                variant="contained"
                // onClick={onRestart}
                sx={{ mt: 3 }}
            >
                Start a New Game - BUTTON COMING SOON
            </Button>
        </Box>
    );
};

export default EndGame;
