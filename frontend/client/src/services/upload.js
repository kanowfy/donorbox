import axios from 'axios';
import { BASE_URL } from '../constants';

const uploadImage = async (data) => {
    const response = await axios.post(`${BASE_URL}/upload/image`, data, {
        headers: {
            'Content-Type': 'multipart/form-data'
        }
    });

    return response.data;
}

export default { uploadImage }
