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


const Login = lazy(() => import('./views/Login'));
const Register = lazy(() => import('./views/Register'));
const GameBoard = lazy(() => import('./views/MainGameView'));
const UserProfile = lazy(() => import('./views/UserProfile'));
const Dashboard = lazy(() => import('./views/DashboardView'));
const GameInvitesView = lazy(() => import('./views/GameInvitesView'));
const FriendsView = lazy(() => import('./views/FriendsView'));
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
                                    <Route path="/login" element={<Login />} />
                                    <Route path="/register" element={<Register />} />
                                    <Route path="/gameboard" element={<ProtectedRoute element={<GameBoard />} />} />
                                    <Route path="/profile" element={<ProtectedRoute element={<UserProfile />} />} />
                                    <Route path="/dashboard" element={<ProtectedRoute element={<Dashboard />} />} />
                                    <Route path="/game-invites" element={<ProtectedRoute element={<GameInvitesView />} />} />
                                    <Route path="/friends" element={<ProtectedRoute element={<FriendsView />} />} />
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
