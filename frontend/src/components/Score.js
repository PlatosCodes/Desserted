import styled from 'styled-components';
import { motion } from 'framer-motion';

const StyledScore = styled(motion.div)`
  font-size: 24px;
  color: #333;
  background-color: #f0f0f0;
  padding: 10px 20px;
  border-radius: 5px;
  box-shadow: 0px 4px 8px rgba(0,0,0,0.1);
  margin: 10px;
  display: inline-block;
`;

const Score = ({ playerScore }) => {
  return (
    <StyledScore
      animate={{ scale: 1.2, rotate: 360 }}
      transition={{ duration: 1 }}
    >
      {playerScore}
    </StyledScore>
  );
};
export default Score;