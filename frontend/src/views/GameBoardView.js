// src/views/Gameboard.js
import React, { useState, useEffect, useCallback, useRef } from 'react';
import { useSelector } from 'react-redux';
import { Container, Grid, Typography, Button, CircularProgress, Snackbar } from '@mui/material';
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
    const [snackbarOpen, setSnackbarOpen] = useState(false);
    const [snackbarMessage, setSnackbarMessage] = useState('');
    const selectedCardsRef = useRef(selectedCards);

    useEffect(() => {
        selectedCardsRef.current = selectedCards;
    }, [selectedCards]);


    useEffect(() => {
        const token = Cookie.get('access_token');
        const ws = connectWebSocket(token, game_id, player_game_id, handleMessage);
        
        // Reset the game state and fetch new data when game_id changes
        resetGameState();
        fetchPlayerHand();

        return () => {
            closeWebSocket();
            resetGameState(); 
        };
    }, [user.id, game_id, player_game_id]);

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

    const handleStealCardDetailedNotification = useCallback((data) => {
        if (data.playerGameID === parseInt(player_game_id, 10)) {
            // Current player stole a card
            setPlayerHand(prevHand => prevHand.filter(card => !selectedCardsRef.current.includes(card.card_id)));
            setPlayerHand(prevHand => {
                const newHand = [...prevHand, data.card];
                // Sort the new hand by card_id
                return newHand.sort((a, b) => parseInt(a.card_id) - parseInt(b.card_id));
            });          
            setSelectedCards([]);
            setSnackbarMessage(`You stole ${data.card.name} from Player ${data.targetPlayerID}`);
        } else if (data.targetPlayerID === parseInt(player_game_id, 10)) {
            // Current player had a card stolen
            setPlayerHand(prevHand => prevHand.filter(card => card.card_id.toString() !== data.card.card_id.toString()));
            setSnackbarMessage(`Player ${data.playerGameID} stole your ${data.card.name}`);
        }
        setSnackbarOpen(true);
    }, [player_game_id, setSnackbarMessage, setSnackbarOpen, setPlayerHand]);

    
    const handleStealCardGenericNotification = useCallback((data) => {
        // Handle generic notification for other players
        setSnackbarMessage(data.notificationText);
        setSnackbarOpen(true);
    }, [setSnackbarMessage, setSnackbarOpen]);

    
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
            case 'refreshPantry':
                if (Array.isArray(data.hand)) {
                    setPlayerHand(data.hand);
                    setSnackbarMessage("Pantry has been refreshed!");
                    setSnackbarOpen(true);
                } else {
                    console.error('Invalid hand data received', data.hand);
                }
                break;
            case 'dessertResponse':
                // Assuming 'data' contains information about the played dessert
                if(data.success) {
                    setPlayerHand(prevHand => prevHand.filter(card => !selectedCardsRef.current.includes(card.card_id)));
                    setSelectedCards([]);
                    setSnackbarMessage(data.message || "Successfully play dessert");
                } else {
                    // Handle failure (e.g., display an error message)
                    setSnackbarMessage(data.message || "Failed to play dessert");
                    setSnackbarOpen(true);
                }
                break;
            case 'stealCardDetailedNotification':
                handleStealCardDetailedNotification(data);
                break;
            case 'stealCardGenericNotification':
                handleStealCardGenericNotification(data);
                break;
            case 'scoreUpdate':
                updateScores(data.players);
                break;
            case 'endTurn':
                console.log("I DID RECEIVE!")
                setCurrentPlayerTurn(data.game.current_player_number);
                break;
            case 'error':
                setSnackbarMessage(data.message);
                setSnackbarOpen(true);
                break;
            default:
                console.warn('Unhandled message type:', data.type);
        }
    }, [setPlayerHand, setPlayerScores, setCurrentPlayerTurn, setSnackbarMessage, setSnackbarOpen, 
        handleStealCardDetailedNotification, handleStealCardGenericNotification]);

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

    // Snackbar close handler
    const handleSnackbarClose = () => {
        setSnackbarOpen(false);
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
                        game_id={game_id}
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
            <Snackbar
                open={snackbarOpen}
                autoHideDuration={6000}
                onClose={handleSnackbarClose}
                message={snackbarMessage}
            />
        </Container>
    );
};

export default GameboardView;