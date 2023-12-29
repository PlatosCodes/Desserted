// src/views/GameBoard.js
import React, { useState, useEffect, useCallback } from 'react';
import { useSelector } from 'react-redux';
import { Container, Grid, Typography, Button } from '@mui/material';
import Hand from '../components/Hand';
import PlayArea from '../components/PlayArea';
import Scoreboard from '../components/Scoreboard'
import { connectWebSocket, sendMessage, closeWebSocket } from '../services/websocketService';
import { selectUser } from '../features/user/userSlice';
import Cookie from 'js-cookie';
import apiService from '../services/apiService';
import { useParams } from 'react-router-dom';


const GameBoardView = () => {
    const user = useSelector(selectUser);
    const { game_id } = useParams();
    const [playerHand, setPlayerHand] = useState([]);
    const [selectedCards, setSelectedCards] = useState([]);
    const [playerScores, setPlayerScores] = useState([]);

    useEffect(() => {
        const token = Cookie.get('access_token');
        const ws = connectWebSocket(token, handleMessage);
        fetchPlayerHand();

        return () => closeWebSocket();
    }, [user.id]);

    useEffect(() => {
        const fetchScores = async () => {
            try {
                const playersData = await apiService.listGamePlayers( {game_id: parseInt(game_id, 10) });
                console.log("We trying here", playersData)
                setPlayerScores(playersData.players.map(player => ({
                    id: player.player_id,
                    // name: player.username, // Adjust according to your data structure
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
        console.log("WebSocket Message Received:", data); // Debugging log
    
        if (data.type === 'drawCardResponse') {
            setPlayerHand([...playerHand, data.card]);
        }
    
        if (data.type === 'scoreUpdate' && Array.isArray(data.players)) {
            console.log("Score Update Received:", data.players); // Debugging log
            setPlayerScores(data.players.map(player => ({
                id: player.player_id,
                score: typeof player.player_score === 'object' ? 
                       (player.player_score.Valid ? player.player_score.Int32 : 0) :
                       (player.player_score !== undefined ? player.player_score : 0)
            })));
            fetchPlayerHand();
        }
    }, [playerHand, playerScores]);
    

    const fetchPlayerHand = async () => {
        try {
            const handData = await apiService.getPlayerHand(user.id);
            setPlayerHand(handData.player_hand);
        } catch (error) {
            console.error('Error fetching player hand:', error);
        }
    };

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
        sendMessage({ type: 'drawCard', data: { game_id: parseInt(game_id, 10), player_game_id: parseInt(playerHand[0].player_game_id, 10) } });
        fetchPlayerHand();
    };

    return (
        <Container>
            <Typography variant="h4" gutterBottom>Game Board</Typography>
            <Scoreboard players={playerScores} />
            <Button onClick={handleDrawCard}>Draw Card</Button>
            <Grid container spacing={3}>
                <Grid item xs={12}>
                    <PlayArea
                        playerGameId={user.id}
                        selectedCards={selectedCards}
                        setSelectedCards={setSelectedCards}
                        fetchPlayerHand={fetchPlayerHand}
                        playerHand={playerHand}
                    />
                </Grid>
                <Grid item xs={12}>
                    <Hand cards={playerHand} onCardSelect={handleCardSelect} selectedCards={selectedCards} />
                </Grid>
            </Grid>
        </Container>
    );
};

export default GameBoardView;