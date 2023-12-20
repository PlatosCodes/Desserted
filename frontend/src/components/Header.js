import React, {useEffect, useState} from 'react';
import { styled } from '@mui/material/styles';
import { AppBar, Toolbar, Typography, Button, BottomNavigation, BottomNavigationAction } from '@mui/material';
import CakeIcon from '@mui/icons-material/Cake';
import LibraryAddCheckIcon from '@mui/icons-material/LibraryAddCheck';
import StarIcon from '@mui/icons-material/Star';
import { useDispatch, useSelector } from 'react-redux';
import { logoutUser, selectAuthenticated } from '../features/user/userSlice';
import { useNavigate, useLocation, Link } from 'react-router-dom';

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
        backgroundColor: '#1E213A',
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
        backgroundColor: 'white',
        boxShadow: 'none',
    }
}));

const Header = () => {

    const dispatch = useDispatch();
    const navigate = useNavigate();
    const location = useLocation();
    const [value, setValue] = useState(location.pathname);

    useEffect(() => {
        setValue(location.pathname);
    }, [location.pathname]);


    const handleCakeIconClick = () => {
        sessionStorage.removeItem("cakeListCurrentPage");
        navigate("/dashboard");
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

    return (
        <StyledAppBar position="static" className={classes.appBar}>
            <Toolbar className={classes.toolbar}>
            <Typography variant="h6" className={classes.title}>
                <Link to="/dashboard" style={{ textDecoration: 'none', display: 'flex', alignItems: 'center', color: 'inherit' }}>
                    <img src='./images/cake.webp' alt="App Logo" className={classes.logo} />
                    Desserted
                </Link>
            </Typography>
                {isAuthenticated && (
                    <Button variant="contained" className={classes.logoutButton} onClick={handleLogout}>
                        Logout
                    </Button>
                )}
            </Toolbar>
            {isAuthenticated && (
                <BottomNavigation 
                value={value} 
                onChange={(event, newValue) => setValue(newValue)} 
                className={classes.bottomNav}
                >
                <BottomNavigationAction label="Dashboard" value="/dashboard" icon={<LibraryAddCheckIcon />} component={Link} to="/dashboard" />
                <BottomNavigationAction label="Cakes" value="/dashboard" icon={<CakeIcon />} component={Link} onClick={handleCakeIconClick} to="/dashboard" />
                <BottomNavigationAction label="Profile" value="/dashboard" icon={<StarIcon />} component={Link} to="/user_profile" />
            </BottomNavigation>
                )}
        </StyledAppBar>
    );
}

export default Header;
