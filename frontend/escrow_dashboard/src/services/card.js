import axios from 'axios';
import { BASE_URL } from '../constants';

const setupTransfer = async (token, card) => {
    const response = await axios.post(`${BASE_URL}/escrow/transfer`, card, {
        headers: { Authorization: `Bearer ${token}` }
    });
    return response.data;
}

export default { setupTransfer };
