// src/views/GameInvitesView.js
import React from 'react';
import { Container, Typography, CircularProgress, Button, List, ListItem, ListItemText, Alert } from '@mui/material';
import { useSelector } from 'react-redux';
import { selectUser } from '../features/user/userSlice';
import { useGameInvites } from '../hooks/useGameInvites';
import { useMutation, useQueryClient } from 'react-query';
import apiService from '../services/apiService';

const GameInvitesView = () => {
  const user = useSelector(selectUser);
  const queryClient = useQueryClient();
  const { data: invitesData, isLoading, isError, error, refetch } = useGameInvites(user.id);
  const gameInvites = invitesData?.game_invite ?? []; 

  const acceptInviteMutation = useMutation(apiService.acceptGameInvite, {
    onMutate: async (variables) => {
      await queryClient.cancelQueries(['gameInvites', user.id]);

      const previousData = queryClient.getQueryData(['gameInvites', user.id]);
      
      queryClient.setQueryData(['gameInvites', user.id], oldData => ({
        ...oldData,
        game_invite: oldData.game_invite.filter(invite => invite.game_id !== variables.game_id),
      }));

      return { previousData };
    },
    onError: (err, variables, context) => {
      queryClient.setQueryData(['gameInvites', user.id], context.previousData);
    },
    onSettled: () => {
      refetch();
    },
  });

  const handleAcceptInvite = (inviteId, gameId) => {
    acceptInviteMutation.mutate({ invitee_player_id: inviteId, game_id: gameId });
  };

  if (isLoading) return <CircularProgress />;
  if (isError) return <Alert severity="error">{error.message}</Alert>;

  return (
    <Container>
      <Typography variant="h4">Game Invites</Typography>
      {error && <Alert severity="error">{error}</Alert>}
      <List>
        {gameInvites.map(invite => (
          <ListItem key={invite.game_invitation_id}>
            <ListItemText primary={`Game invite from player ID: ${invite.invitee_player_id} for game ID: ${invite.game_id}`} />
            <Button variant="contained" color="primary" onClick={() => handleAcceptInvite(invite.invitee_player_id, invite.game_id)}>
              Accept
            </Button>
          </ListItem>
        ))}
      </List>
    </Container>
  );
};

export default GameInvitesView;
