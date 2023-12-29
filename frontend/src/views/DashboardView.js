// src/views/DashboardView.js
import React, { useContext } from 'react';
import { useSelector } from 'react-redux';
import { useNavigate } from 'react-router-dom';
import { Container, Typography, Button, CircularProgress, Alert } from '@mui/material';
import { selectUser } from '../features/user/userSlice';
import { useActivePlayerGames } from '../hooks/useActivePlayerGames';
import apiService from '../services/apiService';
import { useQueryClient } from 'react-query';

// import { useGame } from '../context/GameContext';


const DashboardView = () => {
    // const { gameState, updateGameState } = useContext(useGame);
    const user = useSelector(selectUser);
    const navigate = useNavigate();
    const queryClient = useQueryClient();
    const { data: activeGames, isLoading, isError, error } = useActivePlayerGames(user.id);
    const handleGameClick = async (game_id, player_game_id) => {
        try {
            const data = await apiService.getGameDetails(game_id);
            navigate(`/gameboard/${game_id}/${player_game_id}`);
        } catch (error) {
            // handle error
            console.log(error);
        }
    };

    const handleStartGame = async (game_id) => {
        try {
            await apiService.startGame(game_id);
            // const updatedActiveGames = await apiService.listActivePlayerGames({ player_id: user.id });
            queryClient.invalidateQueries(['activePlayerGames', user.id]);

        } catch (error) {
            console.error("Error starting game:", error);
        }
    };
    
    if (isLoading) return <CircularProgress />;
    if (isError) return <Alert severity="error">{error.message}</Alert>;

    return (
        <Container aligncontent={'center'}>
            <Typography variant="h4" textAlign={'center'}>Welcome to Desserted, {user?.username}</Typography>
            <Typography variant="h5">Your Active Games</Typography>
            {/* Loop through games and categorize based on status */}
            {activeGames && activeGames.length > 0 ? (
                activeGames.map((game, index) => (
                    <div key={index}>
                        <Button 
                            variant="outlined"
                            onClick={() => handleGameClick(game.game_id, game.player_game)}
                        >
                            Game ID: {game.game_id}, Status: {game.status}, Player Game ID: {game.player_game}, Creator: { game.created_by }
                        </Button>
                        {game.status === 'waiting' && game.created_by === user.id && (
                            <Button onClick={() => 
                                handleStartGame(game.game_id)}>Start Game</Button>
                        )}
                        {game.status === 'waiting' && game.created_by !== user.id && (
                            <Typography>Waiting for creator to start.</Typography>
                        )}
                    </div>
                ))
            ) : (
                <Typography>No active games found.</Typography>
            )}
        </Container>
    );
};

export default DashboardView;
