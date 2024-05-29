import axios from 'axios';
import { BASE_URL } from '../constants';

const payout = async (token, projectID) => {
    const response = await axios.post(`${BASE_URL}/escrow/${projectID}/payout`, null, {
        headers: { Authorization: `Bearer ${token}` }
    });
    return response.data;
}

const refund = async (token, projectID) => {
    const response = await axios.post(`${BASE_URL}/escrow/${projectID}/refund`, null, {
        headers: { Authorization: `Bearer ${token}` }
    });
    return response.data;
}

export default { payout, refund }
