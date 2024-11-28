import axios from "axios";
import { BASE_URL } from "../constants";

const getStats = async (token) => {
  const response = await axios.get(`${BASE_URL}/escrow/statistics`, {
    headers: { Authorization: `Bearer ${token}` },
  });
  return response.data;
};

const reviewProject = async (token, data) => {
  console.log(data);
  await axios.post(`${BASE_URL}/escrow/approve`, data, {
    headers: { Authorization: `Bearer ${token}` },
  });
};

export default { getStats, reviewProject };
