// src/hooks/usePlayerScores.js
import { useQuery } from 'react-query';
import apiService from '../services/apiService';

export const usePlayerScores = (gameId) => {
  return useQuery(
    ['playerScores', gameId],
    () => apiService.listGamePlayers({ game_id: gameId }),
    { 
      enabled: !!gameId,
      select: (data) => data.players.map(player => ({
        id: player.player_id,
        score: typeof player.player_score === 'object' ? 
               (player.player_score.Valid ? player.player_score.Int32 : 0) :
               (player.player_score !== undefined ? player.player_score : 0)
      }))
    }
  );
};
