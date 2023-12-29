//src/components/Card.js
import React, {useState} from 'react';
import styled from 'styled-components';

const CardInner = styled.div`
  position: relative;
  width: 100%;
  height: 100%;
  text-align: center;
  transition: transform 0.8s;
  transform-style: preserve-3d;
`;

const CardFront = styled.div`
  display: flex;
  align-items: center;
  justify-content: center;
  background-color: white;
  border-radius: 10px;
  color: black;
  font-size: 16px;
  /* Add more styles as needed */
`;

const CardBack = styled.div`
  background-color: #1a1a1a;
  border-radius: 10px;
  color: white;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 20px;
  /* You can add a pattern or logo here */
`;


const CardContainer = styled.div`
  width: 100px;
  height: 140px;
  border: 1px solid #ccc;
  border-radius: 10px;
  overflow: hidden;
  background-color: #fff;
  box-shadow: 0 2px 5px rgba(0, 0, 0, 0.2);
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: space-between;
  padding: 10px;
  cursor: pointer;
  transform: ${({ isSelected }) => isSelected ? 'scale(1.1)' : 'none'};
  transition: transform 0.3s;
  &:hover ${CardInner} {
    transform: rotateY(180deg);
  }
  &:hover {
    transform: scale(1.05);
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.25);
  }
  &:hover {
    box-shadow: 0 6px 12px rgba(0, 0, 0, 0.3);
  }
`;

const CardImage = styled.img`
  width: 80%;
  height: auto;
`;

const CardText = styled.div`
  margin-top: 10px; // Adjust as needed
`;

const DessertIcons = styled.div`
  display: flex;
  justify-content: center;
  width: 100%;
`;

const DessertIcon = styled.img`
  width: 20px;
  height: 20px;
  margin: 0 2px;
`;

const Card = ({ ingredient, desserts, onSelect }) => {
  const [isSelected, setIsSelected] = useState(false);

  const handleCardClick = () => {
    setIsSelected(!isSelected);
    // onSelect(ingredient.id);
  };  
  
  return (
    <CardContainer onClick={() => onSelect(ingredient.card_id)} isSelected={isSelected}>
    <CardInner> 
        <CardFront>
          <CardImage src={`/images/${ingredient.card_name.toLowerCase()}.webp`} alt={ingredient.card_name} />
          {/* <CardImage src={`/images/egg.webp`} alt={ingredient.name} /> */}
          <CardText>{ingredient.card_name}</CardText>
          <DessertIcons>
            <DessertIcon key="CAKE" src={`/images/cake.webp`} alt="cake" />
          {/* {desserts.map(dessert => (
            <DessertIcon key={dessert.id} src={`/images/${dessert.icon}`} alt={dessert.name} />
          ))} */}
          </DessertIcons>
        </CardFront>
        <CardBack>
          <p>Desserted</p>
        </CardBack>
      </CardInner>
    </CardContainer>
  );
};

export default Card;