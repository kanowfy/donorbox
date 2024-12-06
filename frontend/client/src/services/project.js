import axios from 'axios';
import { BASE_URL } from '../constants';

const getAll = async () => {
    const response = await axios.get(`${BASE_URL}/projects`);
    return response.data;
}

const getForUser = async (token) => {
    const response = await axios.get(`${BASE_URL}/projects/authenticated`, {
        headers: { Authorization: `Bearer ${token}` }
    });

    return response.data;
}

const getOne = async (id) => {
    const response = await axios.get(`${BASE_URL}/projects/${id}`)
    return response.data;
}

const search = async (query, page, pageSize) => {
    const response = await axios.post(`${BASE_URL}/projects/search?page=${page}&page_size=${pageSize}`, {
        "query": query
    }
    )

    return response.data;
}

const create = async (token, data) => {
    const response = await axios.post(`${BASE_URL}/projects`, data, {
        headers: { Authorization: `Bearer ${token}` }
    })

    return response.data;
}

const createUpdate = async (token, data) => {
    const response = await axios.post(`${BASE_URL}/projects/updates`, data, {
        headers: { Authorization: `Bearer ${token}` }
    })

    return response.data;
}

const getUpdates = async (projectID) => {
    const response = await axios.get(`${BASE_URL}/projects/${projectID}/updates`)
    return response.data;
}

const getCategoryByName = async (name) => {
    const response = await axios.get(`${BASE_URL}/categories/${name}`);
    return response.data;
}

const getByCategoryID = async (id) => {
    const response = await axios.get(`${BASE_URL}/projects?category=${id}`);
    return response.data;
}


export default { getAll, getForUser, search, getOne, create, createUpdate, getUpdates, getCategoryByName, getByCategoryID }
