// src/hooks/useUserFriends.js
import { useQuery } from 'react-query';
import apiService from '../services/apiService';

export const useUserFriends = (userId) => {
  return useQuery(
    ['userFriends', userId],
    async () => {
      const response = await apiService.listUserFriends({ user_id: userId, limit: 10, offset: 0 });
          
      return response.friendships || [];
    },
    { enabled: !!userId }
  );
};
