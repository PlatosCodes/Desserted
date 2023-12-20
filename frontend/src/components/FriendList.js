//src/components/FriendList.js
import React, { useEffect, useState } from 'react';
import { useApi } from '../hooks/useApi';
import apiService from '../services/apiService';
import { List, ListItem, ListItemText } from '@mui/material';

const FriendList = () => {
    const { data: friends, request } = useApi(apiService.listUserFriends);
    const [userId, setUserId] = useState(null); // Assume this is fetched from the user's state

    useEffect(() => {
        if (userId) {
            request({ userId });
        }
    }, [userId, request]);

    if (!friends) return <p>Loading friends...</p>;

    return (
        <List>
            {friends.map(friend => (
                <ListItem key={friend.id}>
                    <ListItemText primary={friend.username} />
                </ListItem>
            ))}
        </List>
    );
};

export default FriendList;
