// /src/services/apiService.js 
import axios from 'axios';
import axiosRetry from 'axios-retry';
import Cookie from 'js-cookie';

const API_URL = 'http://localhost:8082';

const axiosInstance = axios.create({
    baseURL: API_URL,
    withCredentials: true,
    timeout: 10000,
});

axiosRetry(axiosInstance, { retries: 3, retryDelay: axiosRetry.exponentialDelay });

axiosInstance.interceptors.request.use(config => {
    const token = Cookie.get('access_token');
    console.log("Token:", token); // Add this line to log the token

    if (token && !config.url.endsWith('/check_session')) {
        config.headers.Authorization = `Bearer ${token}`;
    }
    return config;
});

axiosInstance.interceptors.response.use(
    response => response,
    async error => {
        const originalRequest = error.config;
        if (error.response?.status === 401 && !originalRequest._retry) {
            originalRequest._retry = true;
            try {
                const { data } = await axiosInstance.post('/renew_access');
                Cookie.set('access_token', data.access_token);
                axiosInstance.defaults.headers.Authorization = `Bearer ${data.access_token}`;
                return axiosInstance(originalRequest);
            } catch (refreshError) {
                console.error("Unable to refresh token", refreshError);
                return Promise.reject(refreshError);
            }
        }
        return Promise.reject(error);
    }
);
const handleRequestError = (error, action) => {
    console.error(`Error during ${action}:`, error?.response || error);
    throw error;
};

// Define your API methods
const apiService = {
    createUser: async (userData) => {
        try {
            const response = await axiosInstance.post('/v1/create_user', userData);
            return response.data;
        } catch (error) {
            handleRequestError(error, 'creating user');
        }
    },

    loginUser: async (loginData) => {
        try {
            const response = await axiosInstance.post('/v1/login_user', loginData);
            const data = response.data;
            return data;
        } catch (error) {
            handleRequestError(error, 'user login');
        }
    },

    logoutUser: async () => {
        try {
            const response = await axiosInstance.post('/v1/logout');
            const data = response.data;
            return data;
        } catch (error) {
            handleRequestError(error, 'user logout');
        }
    },

    verifySession: async () => {
        try {
            const sessionId = Cookie.get("session_id");
            console.log(sessionId)
            if (!sessionId) {
                return { isAuthenticated: false };
            }
            const response = await axiosInstance.post('/v1/check_session', { session_id: sessionId });
            return response.data;
        } catch (error) {
            handleRequestError(error, 'checking user session');
        }
    },

    updateUser: async (updateData) => {
        try {
            const response = await axiosInstance.patch('/v1/update_user', updateData);
            return response.data;
        } catch (error) {
            handleRequestError(error, 'updating user');
        }
    },
    

    createGame: async (gameData) => {
        try {
            const response = await axiosInstance.post('/v1/create_game', gameData);
            return response.data;
        } catch (error) {
            handleRequestError(error, 'creating game');
        }
    },

    createFriendship: async ({ frienderId, friendeeUsername }) => {
        try {
            const response = await axiosInstance.post('/v1/create_friendship', { frienderId, friendeeUsername });
            return response.data;
        } catch (error) {
            handleRequestError(error, 'creating friendship');
        }
    },

    listFriendRequests: async (userId) => {
        try {
            const response = await axiosInstance.get(`/v1/list_friend_requests/${userId}`);
            return response.data.friendRequests;
        } catch (error) {
            handleRequestError(error, 'listing friend requests');
        }
    },
    
    acceptFriendRequest: async ({ userId, friendshipId }) => {
        try {
            await axiosInstance.post('/v1/accept_friend_request', { userId, friendshipId });
        } catch (error) {
            handleRequestError(error, 'accepting friend request');
        }
    },

    listUserFriends: async ({ user_id, limit = 10, offset = 0 }) => {
        try {
            const response = await axiosInstance.get(`/v1/list_user_friends/${user_id}?limit=${limit}&offset=${offset}`);
            return response.data;
        } catch (error) {
            handleRequestError(error, 'listing user friends');
        }
    },

    invitePlayerToGame: async ({ inviterPlayerId, inviteeUsername, gameId }) => {
        try {
            const response = await axiosInstance.post('/v1/invite_player_to_game', { inviterPlayerId, inviteeUsername, gameId });
            return response.data;
        } catch (error) {
            handleRequestError(error, 'inviting player to game');
        }
    },

    acceptGameInvite: async (inviteData) => {
        try {
            const response = await axiosInstance.post('/v1/accept_game_invite', inviteData);
            return response.data;
        } catch (error) {
            handleRequestError(error, 'accepting game invite');
        }
    },

    listGamePlayers: async () => {
        try {
            const response = await axiosInstance.get('/v1/list_game_players');
            return response.data;
        } catch (error) {
            handleRequestError(error, 'listing game players');
        }
    },

    startGame: async (startData) => {
        try {
            const response = await axiosInstance.post('/v1/start_game', startData);
            return response.data;
        } catch (error) {
            handleRequestError(error, 'starting game');
        }
    },

    getPlayerHand: async (playerGameId) => {
        try {
            const response = await axiosInstance.get(`/v1/get_player_hand?playerGameId=${playerGameId}`);
            return response.data.playerHand;
        } catch (error) {
            handleRequestError(error, 'getting player hand');
        }
    },

    playDessert: async (dessertData) => {
        try {
            const response = await axiosInstance.post('/v1/play_dessert', dessertData);
            return response.data;
        } catch (error) {
            handleRequestError(error, 'playing dessert');
        }
    },

    drawCard: async ({ gameId, playerGameId }) => {
        try {
            const response = await axiosInstance.post('/v1/draw_card', { gameId, playerGameId });
            return response.data;
        } catch (error) {
            handleRequestError(error, 'drawing a card');
        }
    },

    getPlayerGameData: async (playerGameId) => {
        try {
            const response = await axiosInstance.get(`/v1/get_player_game/${playerGameId}`);
            return response.data;
        } catch (error) {
            handleRequestError(error, 'getting player game data');
        }
    },

};

export default apiService;