import axios from 'axios';
import { BASE_URL } from '../constants';

const getEnded = async () => {
    const response = await axios.get(`${BASE_URL}/projects/ended`);
    return response.data;
}

export default { getEnded }
