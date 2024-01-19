// src/components/FriendRequestList.js
import React from 'react';
import { List, ListItem, Button, ListItemText } from '@mui/material';

const FriendRequestList = ({ friendRequests, onAccept }) => {
    return (
        <List>
            {friendRequests.map(request => (
                <ListItem key={request.friendshipId} sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
                    <ListItemText primary={request.frienderUsername} />
                    <Button variant="contained" color="primary" onClick={() => onAccept(request.friendshipId)}>
                        Accept
                    </Button>
                </ListItem>
            ))}
        </List>
    );
};

export default FriendRequestList;
