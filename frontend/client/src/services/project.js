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

const search = async (query) => {
    const response = await axios.post(`${BASE_URL}/projects/search`, {
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

const getCategoryByName = async (name) => {
    const response = await axios.get(`${BASE_URL}/categories/${name}`);
    return response.data;
}

const getByCategoryID = async (id) => {
    const response = await axios.get(`${BASE_URL}/projects?category=${id}`);
    return response.data;
}

const createProof = async (token, data) => {
    const response = await axios.post(`${BASE_URL}/milestones/proofs`, data, {
        headers: { Authorization: `Bearer ${token}` }
    })

    return response.data;
}

const createReport = async (id, data) => {
    const response = await axios.post(`${BASE_URL}/projects/${id}/reports`, data)
    return response.data;
}



export default { getAll, getForUser, search, getOne, create, getCategoryByName, getByCategoryID, createProof, createReport }
