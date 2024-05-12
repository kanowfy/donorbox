import axios from 'axios';
import { BASE_URL } from '../constants';

const login = async (email, password) => {
    const response = await axios.post(`${BASE_URL}/users/login`, { email, password });
    return response.data;
}

const register = async (data) => {
    const response = await axios.post(`${BASE_URL}/users/register`, data);
    return response.data;
}

const getCurrentUser = async (token) => {
    const response = await axios.get(`${BASE_URL}/users`, {
        headers: { Authorization: `Bearer ${token}` }
    });

    return response.data;
}

export default { login, register, getCurrentUser };
