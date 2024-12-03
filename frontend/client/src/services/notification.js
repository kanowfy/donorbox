import axios from 'axios';
import { BASE_URL } from '../constants';

const getNotifications = async (token, userID) => {
    const response = await axios.get(`${BASE_URL}/notifications/${userID}`, {
        headers: { Authorization: `Bearer ${token}` },
    })
    return response.data;
}

const updateReadNotification = async (token, id) => {
    const response = await axios.post(`${BASE_URL}/notifications/${id}/read`, {
        headers: { Authorization: `Bearer ${token}` },
    })
    return response.data;
}

export default { getNotifications, updateReadNotification }
