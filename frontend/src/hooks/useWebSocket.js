// src/hooks/useWebSocket.js
import { useEffect, useState, useCallback } from 'react';
import Cookie from 'js-cookie'

export const useWebSocket = (userId, onMessageHandler) => {
    const [socket, setSocket] = useState(null);

    const connect = useCallback(() => {
        const token = Cookie.get('access_token');
        const newSocket = new WebSocket(`ws://localhost:8082/ws?access_token=${token}`);

        newSocket.onopen = () => console.log('WebSocket Connected');
        newSocket.onerror = (error) => console.error('WebSocket Error:', error);
        newSocket.onmessage = onMessageHandler;
        newSocket.onclose = (event) => console.log('WebSocket Disconnected', event.code, event.reason);

        setSocket(newSocket);
    }, [userId, onMessageHandler]);

    const disconnect = useCallback(() => {
        if (socket) {
            socket.close();
        }
    }, [socket]);

    const sendMessage = useCallback((message) => {
        if (socket && socket.readyState === WebSocket.OPEN) {
            socket.send(JSON.stringify(message));
        } else {
            console.error("WebSocket is not connected.");
        }
    }, [socket]);

    // Ensure WebSocket is closed when the component unmounts
    useEffect(() => {
        return () => disconnect();
    }, [disconnect]);

    return { connect, disconnect, sendMessage };
};
