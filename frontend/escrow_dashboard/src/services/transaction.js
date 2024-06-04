import axios from 'axios';
import { BASE_URL } from '../constants';

const getAll = async (token) => {
    const response = await axios.get(`${BASE_URL}/transactions`, {
        headers: { Authorization: `Bearer ${token}` }
    });
    return response.data;
}

export default { getAll }
