// //src/components/GameBoard.js
import React from 'react';
import styled from 'styled-components';
import Hand from './Hand';
import { motion } from 'framer-motion';
import { Typography } from '@mui/material';

const BoardContainer = styled(motion.div)`
  padding: 20px;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  background-color: #eaeaea;
  border-radius: 15px;
  box-shadow: 0px 10px 20px rgba(0,0,0,0.2);
  transition: all 0.5s ease-in-out;
`;

const PlayerInfo = styled.div`
  margin-bottom: 20px;
`;

const GameBoard = ({ players }) => (
  <BoardContainer
    initial={{ scale: 0.8, opacity: 0 }}
    animate={{ scale: 1, opacity: 1 }}
    transition={{ duration: 0.8 }}
  >
    {players.map(player => (
      <div key={player.id}>
        <PlayerInfo>
          <Typography variant="h6">{player.name}'s Hand</Typography>
          {/* Display other player information here */}
        </PlayerInfo>
        <Hand cards={player.hand} />
      </div>
    ))}
  </BoardContainer>
);

export default GameBoard;
