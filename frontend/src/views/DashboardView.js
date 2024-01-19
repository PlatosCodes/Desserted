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
import { styled } from '@mui/material/styles';


const StyledContainer = styled(Container)(({ theme }) => ({
    backgroundImage: 'url(/images/background3.webp)',
    backgroundSize: 'cover',
    backgroundPosition: 'center',
    backgroundColor: 'rgba(255, 255, 255, 0.5)',
    padding: theme.spacing(4),
    minWidth: '100vw',
    minHeight: 'calc(100vh - 120px)',
    backgroundRepeat: 'no-repeat',
    overflow: 'hidden',
}));

const VibrantButton = styled(Button)(({ theme }) => ({
    backgroundColor: '#ffff80', // A vibrant color
    color: '#000', // White text color for contrast
    margin: theme.spacing(1), // Add some margin
    '&:hover': {
        backgroundColor: '#ffff00', // Darker shade for hover state
    },
    boxShadow: '0px 4px 10px rgba(0, 0, 0, 0.8)', // Shadow for depth
    fontWeight: 'bold', // Optional: makes the text bold
}));


const VibrantTypography = styled(Typography)(({ theme}) => ({
    color: '#00ccff', // choose a color that contrasts well with your background
    textShadow: '10px 0px 8px rgba(0, 0, 0, 0.8)', // optional: adds a shadow for better legibility
    fontWeight: 'bold', // makes the font bolder
    fontSize: '4rem', // adjust the font size as needed
}));

const SmallVibrantTypography = styled(Typography)(({ theme }) => ({
    color: '#6d4c41',
    textShadow: '10px 4px 8px rgba(0, 0, 0, 0.5)',
    fontWeight: 'bold',
    fontSize: '2rem',
}));


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
        <StyledContainer aligncontent={'center'}>
            <Container>
            <VibrantTypography variant="h4" textAlign={'center'}>Welcome to Desserted, {user?.username}</VibrantTypography>
            <SmallVibrantTypography variant="h5">Your Active Games</SmallVibrantTypography>
            {activeGames && activeGames.length > 0 ? (
                activeGames.map((game, index) => (
                    <div key={index}>
                    <VibrantButton 
                        variant="contained" // Use 'contained' for a solid background
                        onClick={() => handleGameClick(game.game_id, game.player_game_id, game.player_number)}
                    >
                            Game ID: {game.game_id}, Status: {game.status}, Player Number: {game.player_number}, Creator: {game.created_by}
                        </VibrantButton>
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
                <VibrantTypography>No active games found.</VibrantTypography>
            )}
            <Grid container spacing={3}>
                <Grid item xs={12}>
                    <GameInvitesView />
                </Grid>
            </Grid>
            </Container>
        </StyledContainer>
    );
};

export default DashboardView;
