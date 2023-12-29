// src/hooks/useActivePlayerGames.js
import { useQuery } from 'react-query';
import apiService from '../services/apiService';

export const useActivePlayerGames = (userId) => {
  return useQuery(
    ['activePlayerGames', userId],
    async () => {
      const response = await apiService.listActivePlayerGames({ player_id: userId });
      console.log("RESPONSE: ", response)
      return response; 
    },
    { enabled: !!userId }
  );
};
