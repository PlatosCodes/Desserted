// src/theme.js
import { createTheme } from '@mui/material/styles';
// import BottomNavigationAction from '@mui/material/BottomNavigationAction';


const theme = createTheme({
  palette: {
    primary: {
      main: '#fbc2eb',
    },
    secondary: {
      main: '#a6c1ee', // Choose a color that complements your primary color
    },  
    // background: {
    //   default: 'url(/images/background.webp)', // Assuming you have a default background for the whole app
    // },
  },
  typography: {
    fontFamily: 'Pacifico, cursive',
    allVariants: {
      color: '#6d4c41',
    },
  },
  components: {
    MuiAppBar: {
      styleOverrides: {
        root: {
          backgroundColor: 'rgba(255, 250, 245, 0.85)',
          backdropFilter: 'blur(5px)',
        },
      },
    },
    // MuiButton: {
    //   styleOverrides: {
    //     root: {
    //       boxShadow: '0px 4px 10px rgba(0, 0, 0, 0.8)',
    //       fontWeight: 'bold',
    //     },
    //   },
    // },
  },
  // MuiBottomNavigationAction: {
  //   styleOverrides: {
  //     root: {
  //       color: '#6d4c41', // default icon color
  //       '&.Mui-selected': {
  //         color: '#fbc2eb', // selected icon color
  //       },
  //       '& .MuiBottomNavigationAction-label': {
  //         fontSize: '0.7rem', // adjust as needed
  //         '&.Mui-selected': {
  //           fontSize: '0.7rem', // adjust as needed
  //         },
  //       },
  //     },
  //   },
  // },
});

export default theme;
