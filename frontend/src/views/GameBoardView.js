// src/views/GameBoard.js
import React, { useState, useEffect, useCallback } from 'react';
import { useSelector } from 'react-redux';
import { Container, Grid, Typography, Button } from '@mui/material';
import Hand from '../components/Hand';
import PlayArea from '../components/PlayArea';
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

    useEffect(() => {
        const token = Cookie.get('access_token');
        const ws = connectWebSocket(token, handleMessage);
        fetchPlayerHand();

        return () => closeWebSocket();
    }, [user.id]);

    const handleMessage = useCallback((event) => {
        const data = JSON.parse(event.data);

        if (data.type === 'drawCardResponse') {
            // Update the player's hand with the new card information
            setPlayerHand([...playerHand, data.card]); // Assuming 'data.card' is the new card drawn
        }

        // Handle other types of messages here
    }, [playerHand]);

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
        console.log("Game ID: ", parseInt(game_id, 10), "PlayerHand: ", parseInt(playerHand[0].player_game_id, 10))
        sendMessage({ type: 'drawCard', data: { game_id: parseInt(game_id, 10), player_game_id: parseInt(playerHand[0].player_game_id, 10) } });
        fetchPlayerHand();
    };

    return (
        <Container>
            <Typography variant="h4" gutterBottom>Game Board</Typography>
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




// import React, { useState, useEffect } from 'react';
// import { useSelector } from 'react-redux';
// import { Container, Grid, Paper, Typography } from '@mui/material';
// import Hand from '../components/Hand';
// import apiService from '../services/apiService';
// import { selectUser } from '../features/user/userSlice';

// const GameBoard = () => {
//     const user = useSelector(selectUser);
//     const [playerHand, setPlayerHand] = useState([]);

//     useEffect(() => {
//         const fetchPlayerHand = async () => {
//             try {
//                 const handData = await apiService.getPlayerHand(user.id);
//                 setPlayerHand(handData.player_hand);
//             } catch (error) {
//                 console.error('Error fetching player hand:', error);
//             }
//         };
//         fetchPlayerHand();
//     }, [user.id]);

//     return (
//         <Container>
//             <Typography variant="h4" gutterBottom>Game Board</Typography>
//             <Grid container spacing={3}>
//                 <Grid item xs={12} md={6}>
//                     {/* Scoreboard or other components */}
//                 </Grid>
//                 <Grid item xs={12} md={6}>
//                     <Hand cards={playerHand} />
//                 </Grid>
//             </Grid>
//         </Container>
//     );
// };

// export default GameBoard;
