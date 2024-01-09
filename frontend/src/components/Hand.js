import React from 'react';
import styled from 'styled-components';
import Card from './Card';

const HandContainer = styled.div`
  display: flex;
  overflow-x: auto;
  padding: 10px;
  gap: 10px;
  transition: all 0.3s ease-in-out;

  @media (min-width: 768px) {
    justify-content: center;
    flex-wrap: wrap;
  }
`;

const Hand = ({ cards, onCardSelect, selectedCards }) => (
  <HandContainer>
    {cards?.map(card => (
      <Card
        key={card.player_hand_id}
        ingredient={card}
        onSelect={onCardSelect}
        isSelected={selectedCards.includes(card.card_id)}
      />
    ))}
  </HandContainer>
);

export default Hand;
