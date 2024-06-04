import axios from 'axios';
import { BASE_URL } from '../constants';

const getStats = async (token) => {
    const response = await axios.get(`${BASE_URL}/escrow/statistics`, {
        headers: { Authorization: `Bearer ${token}` }
    });
    return response.data;
}

export default { getStats }
