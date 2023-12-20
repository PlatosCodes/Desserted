// src/components/Hand.js
import React from 'react';
import styled from 'styled-components';
import Card from './Card';

const HandContainer = styled.div`
  display: flex;
  overflow-x: auto;
  padding: 10px;
  gap: 10px;
  transition: all 0.3s ease-in-out;

  // Enhance responsiveness
  @media (min-width: 768px) {
    justify-content: center;
    flex-wrap: wrap;
  }

  // Subtle hover effect for interactivity
  &:hover {
    transform: translateY(-10px);
  }
`;

const Hand = ({ cards }) => (
  <HandContainer aria-label="Player's hand">
    {cards.map(card => (
      <Card key={card.id} {...card} />
    ))}
  </HandContainer>
);

export default Hand;
