// GameBoard.js
import React from 'react';
import { Grid, Paper, Typography } from '@mui/material';

const GameBoard = ({ players }) => {
    return (
        <Grid container spacing={2}>
            {/* Iterate over players or game elements */}
            {players.map((player, index) => (
                <Grid item key={index} xs={12} sm={6} md={4}>
                    <Paper>
                        <Typography variant="h6">{player.name}</Typography>
                        {/* Display player-specific data */}
                    </Paper>
                </Grid>
            ))}
        </Grid>
    );
};

export default GameBoard;
