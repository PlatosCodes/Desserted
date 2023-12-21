// src/hooks/useGameInvites.js
import { useQuery } from 'react-query';
import apiService from '../services/apiService';

export const useGameInvites = (userId) => {
    return useQuery(['gameInvites', userId], () => apiService.listGameInvites({ user_id: userId }));
};
