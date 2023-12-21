// src/hooks/useFriendRequests.js
import { useQuery } from 'react-query';
import apiService from '../services/apiService';

export const useFriendRequests = (userId) => {
    return useQuery(
        ['friendRequests', userId],
        () => apiService.listFriendRequests(userId),
        { enabled: !!userId }
    );
};
