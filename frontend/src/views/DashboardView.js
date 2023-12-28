// src/views/DashboardView.js
import React, { useContext } from 'react';
import { useSelector } from 'react-redux';
import { useNavigate } from 'react-router-dom';
import { Container, Typography, Button, CircularProgress, Alert } from '@mui/material';
import { selectUser } from '../features/user/userSlice';
import { useActivePlayerGames } from '../hooks/useActivePlayerGames';
import apiService from '../services/apiService';
// import { useGame } from '../context/GameContext';


const DashboardView = () => {
    // const { gameState, updateGameState } = useContext(useGame);

    const user = useSelector(selectUser);
    const navigate = useNavigate();
    const { data: activeGames, isLoading, isError, error } = useActivePlayerGames(user.id);
    console.log(activeGames)
    const handleGameClick = async (game_id) => {
        try {
            const data = await apiService.getGameDetails(game_id);
            navigate(`/gameboard/${game_id}`);
        } catch (error) {
            // handle error
            console.log(error);
        }
    };
    
    if (isLoading) return <CircularProgress />;
    if (isError) return <Alert severity="error">{error.message}</Alert>;

    return (
        <Container aligncontent={'center'}>
            <Typography variant="h4" textAlign={'center'}>Welcome to Desserted, {user?.username}</Typography>
            <Typography variant="h5">Your Active Games</Typography>
            {activeGames && activeGames.length > 0 ? (
                activeGames.map((game, index) => (
                    <Button 
                        key={index}
                        variant="outlined"
                        onClick={() => handleGameClick(game.game_id)}
                    >
                        Game ID: {game.game_id}, Status: {game.player_status}, PlayerGameID: {game.player_game}
                    </Button>
                ))
            ) : (
                <Typography>No active games found.</Typography>
            )}
        </Container>
    );
};

export default DashboardView;
