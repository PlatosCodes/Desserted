// src/views/Gameboard.js
import React, { useState, useEffect, useCallback } from 'react';
import { useSelector } from 'react-redux';
import { Container, Grid, Typography, Button, CircularProgress } from '@mui/material';
import Hand from '../components/Hand';
import PlayArea from '../components/PlayArea';
import Scoreboard from '../components/Scoreboard';
import { connectWebSocket, sendMessage, closeWebSocket } from '../services/websocketService';
import { selectUser } from '../features/user/userSlice';
import Cookie from 'js-cookie';
import apiService from '../services/apiService';
import { useParams } from 'react-router-dom';
import { useWebSocket } from '../hooks/useWebSocket'; 

const GameboardView = () => {
    const user = useSelector(selectUser);
    const { game_id, player_game_id, player_number } = useParams();
    const [game, setGame] = useState(null); 
    const [playerHand, setPlayerHand] = useState([]);
    const [selectedCards, setSelectedCards] = useState([]);
    const [playerScores, setPlayerScores] = useState([]);
    const [currentPlayerTurn, setCurrentPlayerTurn] = useState(null)
    const [isLoading, setIsLoading] = useState(false);

    useEffect(() => {
        const token = Cookie.get('access_token');
        const ws = connectWebSocket(token, handleMessage);
        
        // Reset the game state and fetch new data when game_id changes
        resetGameState();
        fetchPlayerHand();

        return () => {
            closeWebSocket();
            resetGameState(); 
        };
    }, [user.id, game_id]);

    const resetGameState = () => {
        setPlayerHand([]);
        setSelectedCards([]);
        setPlayerScores([]);
        setCurrentPlayerTurn(null);
    };

    useEffect(() => {
        const fetchGame = async () => {
            setIsLoading(true);
            try {
                const gameData = await apiService.getGameDetails(game_id);                
                setGame(gameData);
                setCurrentPlayerTurn(gameData.game.current_player_number)
            } catch (error) {
                console.error('Error fetching game data:', error);
                // Handle error appropriately
            } finally {
                setIsLoading(false);
            }
        };
        fetchGame();
    }, [game_id]);

    const fetchPlayerHand = async () => {
        setIsLoading(true);
        try {
            const handData = await apiService.getPlayerHand(player_game_id);
            setPlayerHand(handData.player_hand);
            console.log("PlayerHand: ", handData)
        } catch (error) {
            console.error('Error fetching player hand:', error);
        } finally {
            setIsLoading(false);
        }
    };

    useEffect(() => {
        const fetchScores = async () => {
            try {
                const playersData = await apiService.listGamePlayers( {game_id: parseInt(game_id, 10) });
                setPlayerScores(playersData.players.map(player => ({
                    id: player.player_id,
                    player_number: player.player_number,
                    // name: player.username, 
                    score: typeof player.player_score === 'object' ? 
                    (player.player_score.Valid ? player.player_score.Int32 : 0) :
                    (player.player_score !== undefined ? player.player_score : 0)
         })));
            } catch (error) {
                console.error('Error fetching game players:', error);
            }
        };
        fetchScores();
    }, []);
    
    const handleMessage = useCallback((event) => {
        const data = JSON.parse(event.data);
        console.log("WebSocket Message Received:", data);
    
        switch (data.type) {
            case 'drawCardResponse':
                setPlayerHand(prevHand => {
                    const newHand = [...prevHand, data.card];
                    // Sort the new hand by card_id
                    return newHand.sort((a, b) => parseInt(a.card_id) - parseInt(b.card_id));
                });
                break;
            case 'scoreUpdate':
                updateScores(data.players);
                break;
            case 'endTurn':
                setCurrentPlayerTurn(data.game.current_player_number);
                break;
            default:
                console.warn('Unhandled message type:', data.type);
        }
    }, [setPlayerHand, setPlayerScores, setCurrentPlayerTurn, currentPlayerTurn]);

    const handleCardSelect = (cardId) => {
        setSelectedCards(prevSelectedCards => {
            const newSelectedCards = new Set(prevSelectedCards);
            if (newSelectedCards.has(cardId)) {
                newSelectedCards.delete(cardId);
            } else {
                newSelectedCards.add(cardId);
            }
            return Array.from(newSelectedCards);
        });
    };

    const handleDrawCard = () => {
        if (game.game.current_player_number === parseInt(player_number,10)) {
        sendMessage({ type: 'drawCard', data: { game_id: parseInt(game_id, 10), 
                                                player_game_id: parseInt(playerHand[0].player_game_id, 10), 
                                                player_number: parseInt(player_number,10), 
                                                player_hand_id: parseInt(playerHand[0].player_hand_id, 10)} });
    } else {
            alert("It's not your turn!");
        }
    };

    const handleEndTurn = () => {
        if (game.game.current_player_number === parseInt(player_number,10)) {
            sendMessage({ type: 'endTurn', data: { game_id: parseInt(game_id, 10), player_game_id: parseInt(playerHand[0].player_game_id, 10) } });
        }
      };

      const updateScores = (players) => {
        const updatedScores = players.map(player => ({
            id: player.player_id,
            score: player.player_score || 0,
        }));
        setPlayerScores(updatedScores);
    };

    if (isLoading) return <CircularProgress />;
    if (!game) return <Typography variant="h6">Game not found or error loading game</Typography>;
    
    return (
        <Container>
            <Typography variant="h4" gutterBottom>Game Board</Typography>
            <Scoreboard players={playerScores} />
            <Button onClick={handleDrawCard} disabled={currentPlayerTurn !== parseInt(player_number,10)}>
                Draw Card
            </Button>            
            <Grid container spacing={3}>
                <Grid item xs={12}>
                    <PlayArea
                        playerGameId={player_game_id}
                        selectedCards={selectedCards}
                        setSelectedCards={setSelectedCards}
                        setPlayerHand={setPlayerHand}
                        playerHand={playerHand}
                        currentPlayerTurn={parseInt(currentPlayerTurn,10)}
                        playerNumber={parseInt(player_number,10)}
                    />
                </Grid>
                <Grid item xs={12}>
                    <Hand cards={playerHand} onCardSelect={handleCardSelect} selectedCards={selectedCards} />
                </Grid>
            </Grid>
            <Button onClick={handleEndTurn} disabled={currentPlayerTurn !== parseInt(player_number,10)}>
                End Turn
            </Button>
        </Container>
    );
};

export default GameboardView;