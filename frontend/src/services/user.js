import axios from 'axios';
import { BASE_URL } from '../constants';

const login = async (email, password) => {
    const response = await axios.post(`${BASE_URL}/users/login`, { email, password });
    return response.data;
}

const register = async (data) => {
    const response = await axios.post(`${BASE_URL}/users/register`, {
        email: data.email,
        password: data.password,
        first_name: data.first_name,
        last_name: data.last_name
    });
    return response.data;
}

const getCurrent = async (token) => {
    const response = await axios.get(`${BASE_URL}/users/authenticated`, {
        headers: { Authorization: `Bearer ${token}` }
    });

    return response.data;
}

const getOne = async (id) => {
    const response = await axios.get(`${BASE_URL}/users/${id}`);
    return response.data;
}

const getToken = async () => {
    const response = await axios.get(`${BASE_URL}/users/auth/google/token`, {
        withCredentials: true
    })
    return response.data;
}

const verify = async (token) => {
    const response = await axios.post(`${BASE_URL}/users/verify?token=${token}`)
    return response.data;
}

export default { login, register, getCurrent, getOne, getToken, verify };
