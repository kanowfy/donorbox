import axios from 'axios';
import { BASE_URL } from '../constants';

const ask = async (content) => {
    const response = await axios.post(`${BASE_URL}/rag/query`, {
        content: content,
    })
    return response.data;
}

export default { ask };
