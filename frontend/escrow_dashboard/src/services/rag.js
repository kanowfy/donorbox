import axios from "axios";
import { BASE_URL } from "../constants";

const addDocument = async (token, data) => {
  await axios.post(`${BASE_URL}/rag/documents`, data, {
    headers: { Authorization: `Bearer ${token}` },
  });
};

export default { addDocument }
