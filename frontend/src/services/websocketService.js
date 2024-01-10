// src/services/websocketService.js
let socket;

export const connectWebSocket = (token, game_id, player_game_id, onMessageHandler) => {
  socket = new WebSocket(`ws://localhost:8082/ws?access_token=${token}&game_id=${game_id}&player_game_id=${player_game_id}`);

  socket.onopen = () => {
    console.log('WebSocket Connected');
  };

  socket.onmessage = onMessageHandler;

  socket.onerror = (error) => {
    console.error('WebSocket Error:', error);
  };

  socket.onclose = (event) => {
    console.log('WebSocket Disconnected', event.code, event.reason);
  };

  return socket;
};

export const sendMessage = (message) => {
  if (socket && socket.readyState === WebSocket.OPEN) {
      socket.send(JSON.stringify(message));
  } else {
      console.error("WebSocket is not connected.");
  }
};

export const closeWebSocket = () => {
  socket && socket.close();
};
