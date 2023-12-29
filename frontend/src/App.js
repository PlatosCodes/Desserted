// Import necessary modules and components
import React, { lazy, Suspense, useEffect } from 'react';
import { BrowserRouter as Router, Routes, Route, Navigate } from 'react-router-dom';
import { useSelector } from 'react-redux';
import { ThemeProvider } from '@mui/material';
import { QueryClient, QueryClientProvider } from 'react-query';
import { ErrorBoundary } from 'react-error-boundary';
import theme from './theme';
import { selectAuthenticated } from './features/user/userSlice';
import Header from './components/Header';
import ErrorFallback from './components/ErrorFallback';
import { GameProvider } from './context/GameContext';
import ListFriendRequestsView from './views/ListFriendRequestsView';


const Login = lazy(() => import('./views/Login'));
const Register = lazy(() => import('./views/Register'));
// const GameBoard = lazy(() => import('./views/MainGameView'));
const GameBoard = lazy(() => import('./views/GameBoardView'));
const UserProfile = lazy(() => import('./views/UserProfile'));
const UpdateProfile = lazy(() => import('./views/UpdateProfileView'));
const CreateGame = lazy(() => import('./views/CreateGameView'));
const Dashboard = lazy(() => import('./views/DashboardView'));
const GameInvitesView = lazy(() => import('./views/GameInvitesView'));
const FriendsView = lazy(() => import('./views/FriendsView'));
const FriendRequestsView = lazy(() => import('./views/ListFriendRequestsView'));
const UserGamesView = lazy(() => import('./views/UserGamesView'));

function App() {
    const queryClient = new QueryClient();
    const isAuthenticated = useSelector(selectAuthenticated);

    return (
        <GameProvider>
            <ThemeProvider theme={theme}>
                <QueryClientProvider client={queryClient}>
                    <Router>
                        <ErrorBoundary FallbackComponent={ErrorFallback}>
                            <Header />
                            <Suspense fallback={<div>Loading...</div>}>
                                <Routes>
                                    <Route path="/register" element={<Register />} />
                                    <Route path="/login" element={<Login />} />
                                    <Route path="/profile" element={<ProtectedRoute element={<UserProfile />} />} />
                                    <Route path="/update_profile" element={<ProtectedRoute element={<UpdateProfile />} />} />
                                    <Route path="/create_game" element={<ProtectedRoute element={<CreateGame />} />} />
                                    <Route path="/game-invites" element={<ProtectedRoute element={<GameInvitesView />} />} />
                                    <Route path="/dashboard" element={<ProtectedRoute element={<Dashboard />} />} />
                                    <Route path="/gameboard/:game_id" element={<GameBoard />} />
                                    <Route path="/friends" element={<ProtectedRoute element={<FriendsView />} />} />
                                    <Route path="/friend_requests" element={<ProtectedRoute element={<ListFriendRequestsView />} />} />
                                    <Route path="/my-games" element={<ProtectedRoute element={<UserGamesView />} />} />
                                    <Route path="/" element={
                                        isAuthenticated ? <Navigate to="/dashboard" replace /> : <Navigate to="/login" replace />
                                    } />
                                </Routes>
                            </Suspense>
                        </ErrorBoundary>
                    </Router>
                </QueryClientProvider>
            </ThemeProvider>
        </GameProvider>
    );
}

// ProtectedRoute component to guard private routes
function ProtectedRoute({ element }) {
    const isAuthenticated = useSelector(selectAuthenticated);

    if (!isAuthenticated) {
        return <Navigate to="/login" />;
    }

    return element;
}

export default App;
