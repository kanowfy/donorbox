import axios from 'axios';
import { BASE_URL } from '../constants';

const login = async (email, password) => {
    const response = await axios.post(`${BASE_URL}/escrow/login`, { email, password });
    return response.data;
}

const getCurrent = async (token) => {
    const response = await axios.get(`${BASE_URL}/escrow/authenticated`, {
        headers: { Authorization: `Bearer ${token}` }
    });

    return response.data;
}

export default { login, getCurrent };
