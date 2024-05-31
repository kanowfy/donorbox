import axios from 'axios';
import { BASE_URL } from '../constants';

const getProjectStats = async (projectID) => {
    const response = await axios.get(`${BASE_URL}/projects/${projectID}/backings/stats`);
    return response.data;
}

const backProject = async (projectID, req) => {
    const response = await axios.post(`${BASE_URL}/projects/${projectID}/backings`, req);
    return response.data;
}

export default { getProjectStats, backProject }
