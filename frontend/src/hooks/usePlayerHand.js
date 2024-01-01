// src/hooks/usePlayerHand.js
import { useQuery } from 'react-query';
import apiService from '../services/apiService';

export const usePlayerHand = (player_game_id) => {
  return useQuery(
    ['playerHand', player_game_id],
    () => apiService.getPlayerHand(player_game_id),
    { enabled: !!player_game_id }
);
};
