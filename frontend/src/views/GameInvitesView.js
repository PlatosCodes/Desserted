// src/views/GameInvitesView.js
import React, { useCallback } from 'react';
import { Container, Typography, CircularProgress, Button, List, ListItem, ListItemText, Alert } from '@mui/material';
import { useSelector } from 'react-redux';
import { selectUser } from '../features/user/userSlice';
import { useGameInvites } from '../hooks/useGameInvites';
import { useMutation, useQueryClient } from 'react-query';
import apiService from '../services/apiService';
import { styled } from '@mui/material/styles';


const SmallVibrantTypography = styled(Typography)(({ theme }) => ({
  color: '##000',
  textShadow: '10px 0px 8px rgba(0, 0, 0, 0.8)',
  fontWeight: 'bold',
  fontSize: '2rem',
}));


const GameInvitesView = () => {
  const user = useSelector(selectUser);
  const queryClient = useQueryClient();
  const { data: invitesData, isLoading, isError, error } = useGameInvites(user.id);
  const gameInvites = invitesData?.game_invite ?? [];

  const acceptInviteMutation = useMutation(apiService.acceptGameInvite, {
    onMutate: async (variables) => {
      await queryClient.cancelQueries(['gameInvites', user.id]);
      const previousData = queryClient.getQueryData(['gameInvites', user.id]);
      // Optimistic update
      queryClient.setQueryData(['gameInvites', user.id], oldData => ({
        ...oldData,
        game_invite: oldData.game_invite.filter(invite => invite.game_id !== variables.game_id),
      }));
      return { previousData };
    },
    onError: (err, variables, context) => {
      // Reset to previous data
      queryClient.setQueryData(['gameInvites', user.id], context.previousData);
      // Handle specific error (game already started)
      if (err?.response?.data?.message.includes('game has already started')) {
        alert("The game has already started.");
      } else {
        alert("An error occurred.");
      }
    },
    onSettled: () => {
      queryClient.invalidateQueries(['gameInvites', user.id]);
    },
  });

  const handleAcceptInvite = useCallback((inviteId, gameId) => {
    acceptInviteMutation.mutate({ invitee_player_id: inviteId, game_id: gameId });
  }, [acceptInviteMutation]);

  if (isLoading) return <CircularProgress />;
  if (isError) return <Alert severity="error">{error.message}</Alert>;

  return (
    <Container>
      <SmallVibrantTypography variant="h4">Game Invites</SmallVibrantTypography>
      <List>
        {gameInvites.map(invite => (
          <ListItem key={invite.game_invitation_id}>
            <ListItemText primary={`Game invite from player ID: ${invite.invitee_player_id} for game ID: ${invite.game_id}`} />
            <Button
              variant="contained"
              color="primary"
              onClick={() => handleAcceptInvite(invite.invitee_player_id, invite.game_id)}
              disabled={acceptInviteMutation.isLoading}
            >
              Accept
            </Button>
          </ListItem>
        ))}
      </List>
    </Container>
  );
};

export default GameInvitesView;
