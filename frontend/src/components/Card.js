import React from 'react';
import styled from 'styled-components';
import { motion } from 'framer-motion';

const CardInner = styled.div`
  position: relative;
  width: 100%;
  height: 100%;
  text-align: center;
  transition: transform 0.8s ease;
  transform-style: preserve-3d;
`;

const CardContainer = styled(motion.div)`
  width: 120px;
  min-width: 120px; // Prevents cards from becoming too narrow
  height: 168px;
  border-radius: 10px;
  overflow: hidden;
  cursor: pointer;
  perspective: 1000px;
  &:hover {
    box-shadow: 0 8px 16px rgba(0, 0, 0, 0.6);
    transition: box-shadow 0.3s ease;
  }
`;


const CardFace = styled.div`
  position: absolute;
  width: 100%;
  height: 100%;
  backface-visibility: hidden;
  border-radius: 10px;
  overflow: hidden;
`;

const CardFront = styled(CardFace)`
  background: url(${props => props.bgImage}) no-repeat center center;
  background-size: cover;
  display: flex;
  align-items: flex-end;
  justify-content: center;
`;

const CardBack = styled(CardFace)`
  background-color: #1a1a1d;
  color: white;
  transform: rotateY(180deg);
  display: flex;
  align-items: center;
  justify-content: center;
`;

const CardTitle = styled.div`
  color: white;
  text-shadow: 2px 2px 4px rgba(0, 0, 0, 0.8);
  font-size: 12px;
  margin-bottom: 10px;
  padding: 5px;
  background-color: rgba(0, 0, 0, 0.25);
  border-radius: 5px;
`;

const Card = ({ ingredient, onSelect, isSelected }) => {
  const cardVariants = {
    hover: { scale: 1.05 },
    tap: { scale: 0.95 }
  };

  const cardName = ingredient?.name || 'Unknown Card';

  return (
    <CardContainer
      onClick={() => onSelect(ingredient.card_id)}
      variants={cardVariants}
      initial={false}
      animate={isSelected ? 'tap' : 'hover'}
      whileHover="hover"
      whileTap="tap"
    >
      <CardInner>
        <CardFront bgImage={`/images/${cardName.toLowerCase().replace(/ /g, '_')}.webp`} alt={cardName}>
          <CardTitle>{cardName}</CardTitle>
        </CardFront>
        <CardBack>
          <p>Desserted</p>
        </CardBack>
      </CardInner>
    </CardContainer>
  );
};

export default Card;
