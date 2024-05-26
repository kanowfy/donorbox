import axios from 'axios';
import { BASE_URL } from '../constants';

const getProjectStats = async (projectID) => {
    const response = await axios.get(`${BASE_URL}/projects/${projectID}/backings/stats`);
    return response.data;
}

export default { getProjectStats }
