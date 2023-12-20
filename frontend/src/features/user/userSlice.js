// src/features/user/userSlice.js
import { createSlice, createAsyncThunk } from '@reduxjs/toolkit';
import apiService from '../../services/apiService';
import Cookie from 'js-cookie';

export const checkUserSession = createAsyncThunk(
    'user/checkSession',
    async () => {
        try {
            const response = await apiService.verifySession();
            return response.data.isAuthenticated;
        } catch (error) {
            return false;
        }
    }
);

export const updateUserProfile = createAsyncThunk(
    'user/updateUserProfile',
    async (updateData, { rejectWithValue }) => {
        try {
            const response = await apiService.updateUser(updateData);
            return response.user;
        } catch (error) {
            return rejectWithValue(error.response.data);
        }
    }
);

export const userSlice = createSlice({
    name: 'user',
    initialState: {
        userData: localStorage.getItem('userData') || null,
        isAuthenticated: localStorage.getItem('isAuthenticated') === 'true',
        isChecking: false,
        isAdmin: false,
        error: null,
    },
    reducers: {
        loginUser: (state, action) => {
            const { session_id, token, user } = action.payload;
            Cookie.set('session_id', session_id);
            Cookie.set('token', token);
            localStorage.setItem('userData', JSON.stringify(user));
            state.userData = user;
            state.isAuthenticated = true;
        },
        
        logoutUser: (state) => {
            Cookie.remove('session_id');
            Cookie.remove('token');
            localStorage.removeItem('userData');
            state.userData = null;
            state.isAuthenticated = false;
        },
    },
    extraReducers: (builder) => {
        builder
            .addCase(checkUserSession.pending, (state) => {
                state.isChecking = true;
            })
            .addCase(checkUserSession.fulfilled, (state, action) => {
                state.isChecking = false;
                state.isAuthenticated = action.payload;
                if (!action.payload) {
                    state.userData = null;
                }
            })
            .addCase(checkUserSession.rejected, (state, action) => {
                state.isChecking = false;
                state.error = action.error.message;
                state.isAuthenticated = false;
                state.userData = null;
            })
            .addCase(updateUserProfile.fulfilled, (state, action) => {
                state.userData = action.payload;
                localStorage.setItem('userData', JSON.stringify(action.payload));
            });
    }
});

export const { loginUser, logoutUser } = userSlice.actions;
export const selectUser = (state) => {
    return typeof state.user.userData === 'string' ? JSON.parse(state.user.userData) : state.user.userData;
};
export const selectAuthenticated = (state) => state.user.isAuthenticated;

export default userSlice.reducer;
