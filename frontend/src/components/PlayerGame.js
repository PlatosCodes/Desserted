// PlayerGame.js
import React, { useState, useEffect } from 'react';
import { useSelector } from 'react-redux';
import apiService from './services/apiService';
import Score from './components/Score';

const PlayerGame = () => {
    const [playerGameData, setPlayerGameData] = useState(null);
    const userData = useSelector(state => state.user.userData); // Assuming this is where you store user data

    useEffect(() => {
        if (userData && userData.id) {
            const fetchPlayerGameData = async () => {
                const data = await apiService.getPlayerGameData(userData.id);
                setPlayerGameData(data);
            };
            fetchPlayerGameData();
        }
    }, [userData]);

    if (!playerGameData) {
        return <div>Loading player game data...</div>;
    }

    return (
        <div>
            <Score playerScore={playerGameData.playerScore} />
        </div>
    );
};

export default PlayerGame;
