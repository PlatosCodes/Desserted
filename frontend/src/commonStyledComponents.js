// src/commonStyledComponents.js
import { styled } from '@mui/material/styles';
import Button from '@mui/material/Button';

export const GradientButton = styled(Button)(({ theme }) => ({
  backgroundImage: 'linear-gradient(to right, #fbc2eb 0%, #a6c1ee 100%)',
  color: theme.palette.getContrastText('#a6c1ee'),
  boxShadow: 'inset 0 2px 4px 0 rgba(0,0,0,0.2)',
  '&:hover': {
    boxShadow: '0 4px 20px 0 rgba(0,0,0,0.25)',
    backgroundImage: 'linear-gradient(to right, #fbc2eb 0%, #a1c4fd 100%)',
  },
}));