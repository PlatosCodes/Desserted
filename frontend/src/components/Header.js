// src/components/Header.js
import { styled } from '@mui/material/styles';
import { AppBar, Toolbar, Typography, IconButton } from '@mui/material';
import CakeIcon from '@mui/icons-material/Cake';
import LibraryAddCheckIcon from '@mui/icons-material/LibraryAddCheck';
import StarIcon from '@mui/icons-material/Star';
import { useDispatch, useSelector } from 'react-redux';
import { logoutUser, selectAuthenticated } from '../features/user/userSlice';
import { useNavigate, useLocation, Link } from 'react-router-dom';
import { GradientButton } from '../commonStyledComponents';
import MailIcon from '@mui/icons-material/Mail';
import Box from '@mui/material/Box';



const PREFIX = 'Header';

const classes = {
    appBar: `${PREFIX}-appBar`,
    toolbar: `${PREFIX}-toolbar`,
    title: `${PREFIX}-title`,
    logo: `${PREFIX}-logo`,
    logoutButton: `${PREFIX}-logoutButton`,
    bottomNav: `${PREFIX}-bottomNav`
};

const StyledAppBar = styled(AppBar)((
    {
        theme
    }
) => ({
    [`&.${classes.appBar}`]: {
        backgroundColor: 'rgba(255, 250, 245, 0.85)',
        backdropFilter: 'blur(5px)',
    },

    [`& .${classes.toolbar}`]: {
        display: 'flex',
        justifyContent: 'space-between',
        padding: theme.spacing(0, 3),
    },

    [`& .${classes.title}`]: {
        display: 'flex',
        alignItems: 'center',
        fontSize: '1.5rem',
        letterSpacing: '1px',
        color: 'black'
    },

    [`& .${classes.logo}`]: {
        marginRight: theme.spacing(2),
        height: '2rem',
        width: '2rem',
    },

    [`& .${classes.logoutButton}`]: {
        backgroundColor: '#F72585',
        color: '#FFFFFF',
        borderRadius: '8px',
        padding: theme.spacing(1, 2),
        '&:hover': {
            backgroundColor: '#D43F8D',
        }
    },

    [`& .${classes.bottomNav}`]: {
        boxShadow: 'none',
        backgroundColor: 'rgba(255, 250, 245, 0.85)',

    }
}));

const NavigationAction = styled(Box)(({ theme, selected }) => ({
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
    justifyContent: 'center',
    color: selected ? theme.palette.secondary.main : theme.palette.action.active,
    paddingBottom: theme.spacing(1),
  }));

const Header = () => {

    const dispatch = useDispatch();
    const navigate = useNavigate();
    const location = useLocation();

    const handleCakeIconClick = () => {
        sessionStorage.removeItem("cakeListCurrentPage");
        navigate("/create_game");
      };
      
    const handleLogout = async () => {
        try {
            dispatch(logoutUser());
            localStorage.removeItem('userData');
            localStorage.removeItem('isAuthenticated');
            navigate('/login');
        } catch (error) {
            console.error("Error during logout: ", error);
        }
    };

    const isAuthenticated = useSelector(selectAuthenticated);

    const isSelected = (path) => location.pathname === path;

    return (
        <StyledAppBar position="static" className={classes.appBar}>
            <Toolbar className={classes.toolbar}>
            <Box sx={{ display: 'flex', alignItems: 'center', marginRight: '120px'}}>
                    <Typography variant="h6" className={classes.title}>
                        <Link to="/dashboard" style={{ textDecoration: 'none', color: 'inherit' }}>
                            <img src='images/cake.webp' alt="App Logo" className={classes.logo} />
                            Desserted
                        </Link>
                    </Typography>
                </Box>

                {isAuthenticated && (
                    <Box sx={{ display: 'flex', justifyContent: 'space-around', flexGrow: 1 }}>
                        <NavigationAction selected={isSelected('/dashboard')}>
                            <IconButton color="inherit" component={Link} to="/dashboard">
                                <LibraryAddCheckIcon />
                            </IconButton>
                            <Typography variant="caption">Dashboard</Typography>
                        </NavigationAction>
                        <NavigationAction selected={isSelected('/create_game')}>
                            <IconButton color="inherit" onClick={handleCakeIconClick} component={Link} to="/create_game">
                                <CakeIcon />
                            </IconButton>
                            <Typography variant="caption">Create Game</Typography>
                        </NavigationAction>
                        <NavigationAction selected={isSelected('/profile')}>
                            <IconButton color="inherit" component={Link} to="/profile">
                                <StarIcon />
                            </IconButton>
                            <Typography variant="caption">Profile</Typography>
                        </NavigationAction>
                        <NavigationAction selected={isSelected('/friend_requests')}>
                            <IconButton color="inherit" component={Link} to="/friend_requests">
                                <MailIcon />
                            </IconButton>
                            <Typography variant="caption">Friend Requests</Typography>
                        </NavigationAction>
                    </Box>
                )}

                {isAuthenticated && (
                    <Box sx={{ display: 'flex', justifyContent: 'flex-end', marginLeft: '120px'}}>
                        <GradientButton variant="contained" onClick={handleLogout}>
                            Logout
                        </GradientButton>
                    </Box>
                )}
            </Toolbar>
        </StyledAppBar>
    );
};

export default Header;