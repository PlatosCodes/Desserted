// App.js
import React, { lazy, Suspense, useEffect } from 'react';
import { BrowserRouter as Router, Routes, Route, Navigate } from 'react-router-dom';
import { useSelector, useDispatch } from 'react-redux';
import { ThemeProvider } from '@mui/material';
import { QueryClient, QueryClientProvider } from 'react-query';
import { ErrorBoundary } from 'react-error-boundary';
import theme from './theme';
import { checkUserSession, logoutUser, selectAuthenticated } from './features/user/userSlice';
import Header from './components/Header';
import ErrorFallback from './components/ErrorFallback';
import Cookie from 'js-cookie'



function ProtectedRoute({ element }) {
    const dispatch = useDispatch();

    const isAuthenticated = useSelector(selectAuthenticated);

    useEffect(() => {
        const token = Cookie.get('token');
        const isSessionValid = checkUserSession();
        if (!isSessionValid) {
            dispatch(logoutUser());
            localStorage.removeItem('userData');
        }
    }, [dispatch]);

    return element;
}

function App() {
    const queryClient = new QueryClient();

    const Login = lazy(() => import('./views/Login'));
    const Register = lazy(() => import('./views/Register'));
    const GameBoard = lazy(() => import('./views/MainGameView'));
    const UserProfile = lazy(() => import('./views/UserProfile'));
    const Dashboard = lazy(() => import('./views/DashboardView'));


    return (
        <ThemeProvider theme={theme}>
            <QueryClientProvider client={queryClient}>
                <Router>
                    <ErrorBoundary FallbackComponent={ErrorFallback}>
                        <Header />
                        <Suspense fallback={<div>Loading...</div>}>
                            <Routes>
                                <Route path="/login" element={<Login />} />
                                <Route path="/register" element={<Register />} />
                                <Route path="/gameboard" element={<ProtectedRoute><GameBoard /></ProtectedRoute>} />
                                <Route path="/profile" element={<ProtectedRoute><UserProfile /></ProtectedRoute>} />
                                <Route path="/dashboard" element={<ProtectedRoute><Dashboard /></ProtectedRoute>} />
                                <Route path="/" element={
                                    ProtectedRoute.isAuthenticated ? <Navigate to="/dashboard" replace /> : <Navigate to="/login" replace />
                                    } />
                            </Routes>
                        </Suspense>
                    </ErrorBoundary>
                </Router>
            </QueryClientProvider>
        </ThemeProvider>
    );
}

export default App;
