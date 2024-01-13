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
        console.log("WebSocket Message Received:", event);

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
            case 'dessertPlayedUpdate':
                console.log("VIVA", data.data)
                
                let message = `Player ${data.data.player_number} played ${data.data.dessert_name} dessert for ${data.data.dessert_score} points!`;
            
                if (data.data.player_number === parseInt(player_number, 10)) {
                    message = data.data.success 
                        ? `You whipped up a ${data.data.dessert_name} dessert for ${data.data.dessert_score} points!` 
                        : `Failed to play ${data.data.dessert_name}. Check your ingredients!`;
                    
                    if (data.data.success) {
                        setPlayerHand(prevHand => prevHand.filter(card => !selectedCardsRef.current.includes(card.card_id)));
                        setSelectedCards([]);
                    }
                }
                setSnackbarMessage(message);
                setSnackbarOpen(true);
                break;
            case 'stealCardDetailedNotification':
                handleStealCardDetailedNotification(data);
                break;
            case 'stealCardGenericNotification':
                handleStealCardGenericNotification(data);
                break;
            case 'scoreUpdate':
                updateScores(data.score_data);
                break;
            case 'endTurnUpdate':
                setCurrentPlayerTurn(data.end_turn_data.current_player_number);
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
            {/* Desserts List */}
            <div className="desserts-list">
                <Typography variant="h5">Desserts List</Typography>
                <ul>
                    <li>Cake (10 Points) - Requires Flour, Sugar, Eggs</li>
                    <li>Pie (15 Points) - Requires Flour, Butter, Berries</li>
                    <li>Chocolate Chip Cookies (20 Points) - Requires Flour, Sugar, Dark Chocolate</li>
                    <li>Cheesecake (25 Points) - Requires Cream Cheese, Eggs, Vanilla</li>
                    <li>Marble Cake (30 Points) - Requires Flour, Sugar, Eggs, Butter, Vanilla, Chocolate</li>
                    <li>Triple Chocolate Brownies (35 Points) - Requires Chocolate, Butter, Sugar, Flour, Eggs</li>
                    <li>Gourmet Truffles (40 Points) - Requires Chocolate, Cream Cheese, Honey</li>
                    <li>Raspberry Chocolate Cheesecake (45 Points) - Requires Cream Cheese, Eggs, Sugar, Vanilla, Chocolate, Berries</li>
                    <li>Gold Leaf Cupcakes (50 Points) - Requires Flour, Sugar, Butter, Edible Gold Leaf</li>
                </ul>
            </div>

            {/* Special Cards */}
            <div className="special-cards">
                <Typography variant="h5">Special Cards</Typography>
                <ul>
                    <li>Wildcard Ingredient: Can substitute any one ingredient.</li>
                    <li>Steal Card: Take a card at random from another player’s hand.</li>
                    <li>Double Points: Doubles the points of the next dessert you play.</li>
                    <li>Refresh Hand: Discard your hand and draw the same number of cards.</li>
                    <li>Instant Bake: Play a dessert without using a turn.</li>
                    <li>Mystery Ingredient: Add random extra points (1-10) to a dessert.</li>
                    <li>Sabotage: Force another player to skip their turn.</li>
                    <li>Glass of Milk: Add 3 points to your dessert.</li>
                </ul>
            </div>
        </Container>
    );
};

export default GameboardView;