import { createSlice, createAsyncThunk } from '@reduxjs/toolkit';
import apiService from '../../services/apiService';
import Cookie from 'js-cookie'

export const checkUserSession = createAsyncThunk(
    'user/checkSession',
    async () => {
      try {
        const response = await apiService.verifySession()
        return response.data.isAuthenticated;
      } catch (error) {
        return false; 
      }
    }
  );
  

export const userSlice = createSlice({
    name: 'user',
    initialState: {
        userData: JSON.parse(localStorage.getItem('userData')) || null,
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
            state.isAuthenticated = action.payload; // Set based on actual response
            if (!action.payload) {
            state.userData = null; // Clear user data if not authenticated
            }
        })
        .addCase(checkUserSession.rejected, (state, action) => {
            state.isChecking = false;
            state.error = action.error.message;
            state.isAuthenticated = false;
            state.userData = null; // Clear user data on error
        });
}

});

export const { loginUser, logoutUser } = userSlice.actions;
export const selectUser = (state) => state.user.userData;
export const selectAuthenticated = (state) => state.user.isAuthenticated;
export default userSlice.reducer;
