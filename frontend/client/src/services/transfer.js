import axios from 'axios';
import { BASE_URL } from '../constants';

const getCard = async (token, id) => {
    const response = await axios.get(`${BASE_URL}/cards/${id}/project`, {
        headers: { Authorization: `Bearer ${token}` }
    });

    return response.data;
}

const setupCard = async (token, projectID, card) => {
    const response = await axios.post(`${BASE_URL}/projects/${projectID}/transfer`, card, {
        headers: { Authorization: `Bearer ${token}` }
    });
    return response.data;
}

export default { getCard, setupCard }
