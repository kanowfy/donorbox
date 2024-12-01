import axios from 'axios';
import { BASE_URL } from '../constants';

const getPendingVerificationUsers = async (token) => {
    const response = await axios.get(`${BASE_URL}/users/pendingVerification`, {
        headers: { Authorization: `Bearer ${token}` }
    });

    return response.data;
}

export default { getPendingVerificationUsers };
