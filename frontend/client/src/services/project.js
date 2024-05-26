import axios from 'axios';
import { BASE_URL } from '../constants';

const getAll = async () => {
    const response = await axios.get(`${BASE_URL}/projects`);
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

export default { getAll, search, getOne }
