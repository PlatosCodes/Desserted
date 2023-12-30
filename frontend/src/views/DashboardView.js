// src/views/DashboardView.js
import React from 'react';
import { useSelector } from 'react-redux';
import { useMutation, useQueryClient } from 'react-query';
import { useNavigate } from 'react-router-dom';
import { Container, Typography, Button, CircularProgress, Alert } from '@mui/material';
import { selectUser } from '../features/user/userSlice';
import { useActivePlayerGames } from '../hooks/useActivePlayerGames';
import apiService from '../services/apiService';

const DashboardView = () => {
    const user = useSelector(selectUser);
    const navigate = useNavigate();
    const queryClient = useQueryClient();
    const { data: activeGames, isLoading, isError, error } = useActivePlayerGames(user.id);

    // React Query mutation for starting a game
    const startGameMutation = useMutation(apiService.startGame, {
        onSuccess: () => {
            // Invalidate and refetch active games data
            queryClient.invalidateQueries(['activePlayerGames', user.id]);
        },
        onError: (error) => {
            console.error("Error starting game:", error);
            // Handle error appropriately
        },
    });

    const handleGameClick = (game_id, player_game_id) => {
        navigate(`/gameboard/${game_id}/${player_game_id}`);
    };

    const handleStartGame = (game_id) => {
        startGameMutation.mutate(game_id);
    };

    if (isLoading) return <CircularProgress />;
    if (isError) return <Alert severity="error">{error.message}</Alert>;

    return (
        <Container aligncontent={'center'}>
            <Typography variant="h4" textAlign={'center'}>Welcome to Desserted, {user?.username}</Typography>
            <Typography variant="h5">Your Active Games</Typography>
            {activeGames && activeGames.length > 0 ? (
                activeGames.map((game, index) => (
                    <div key={index}>
                        <Button 
                            variant="outlined"
                            onClick={() => handleGameClick(game.game_id, game.player_game_id)}
                        >
                            Game ID: {game.game_id}, Status: {game.status}, Player Game ID: {game.player_game_id}, Creator: {game.created_by}
                        </Button>
                        {game.status === 'waiting' && game.created_by === user.id && (
                            <Button 
                                onClick={() => handleStartGame(game.game_id)}
                                disabled={startGameMutation.isLoading}
                            >
                                Start Game
                            </Button>
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
