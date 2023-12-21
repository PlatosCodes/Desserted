// src/components/Scoreboard.js
import React from 'react';
import { Paper, Typography, List, ListItem, ListItemText } from '@mui/material';

const Scoreboard = ({ players }) => {
    return (
        <Paper elevation={3} style={{ padding: '20px' }}>
            <Typography variant="h6">Scoreboard</Typography>
            <List>
                {players.map(player => (
                    <ListItem key={player.id}>
                        <ListItemText primary={player.name} secondary={`Score: ${player.score}`} />
                    </ListItem>
                ))}
            </List>
        </Paper>
    );
};

export default Scoreboard;
