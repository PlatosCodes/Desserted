// src/views/DashboardView.js
import React from 'react';
import { useSelector } from 'react-redux';
import { useMutation, useQueryClient } from 'react-query';
import { useNavigate } from 'react-router-dom';
import { Container, Typography, Button, CircularProgress, Alert, Grid } from '@mui/material';
import { selectUser } from '../features/user/userSlice';
import { useActivePlayerGames } from '../hooks/useActivePlayerGames';
import apiService from '../services/apiService';
import GameInvitesView from './GameInvitesView'

const DashboardView = () => {
    const user = useSelector(selectUser);
    const navigate = useNavigate();
    const queryClient = useQueryClient();
    const { data: activeGames, isLoading, isError, error } = useActivePlayerGames(user.id);

    const startGameMutation = useMutation(apiService.startGame, {
        onSuccess: () => {
            // Invalidate and refetch active games data
            queryClient.invalidateQueries(['activePlayerGames', user.id]);
        },
        onError: (error) => {
            console.error("Error starting game:", error);
            // TODO: Handle error appropriately
        },
    });

    const handleGameClick = (game_id, player_game_id, player_number) => {
        navigate(`/gameboard/${game_id}/${player_game_id}/${player_number}`);
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
                            onClick={() => handleGameClick(game.game_id, game.player_game_id, game.player_number)}
                        >
                            Game ID: {game.game_id}, Status: {game.status}, Player Number: {game.player_number}, Creator: {game.created_by}
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
            <Grid container spacing={3}>
                <Grid item xs={12}>
                    <GameInvitesView />
                </Grid>
            </Grid>
        </Container>
    );
};

export default DashboardView;
