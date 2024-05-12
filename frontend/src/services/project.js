import axios from 'axios';
import { BASE_URL } from '../constants';

const getAllProjects = async () => {
    const response = await axios.get(`${BASE_URL}/projects`);
    return response.data;
}

export default { getAllProjects }
