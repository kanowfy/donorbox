import axios from "axios";
import { BASE_URL } from "../constants";

const getEnded = async () => {
  const response = await axios.get(`${BASE_URL}/projects/ended`);
  return response.data;
};

const getPending = async (token) => {
  const response = await axios.get(`${BASE_URL}/projects/pending`, {
    headers: { Authorization: `Bearer ${token} `},
  });
  return response.data;
};

const getUnresolvedMilestones = async (token) => {
  const response = await axios.get(`${BASE_URL}/milestones/unresolved`, {
    headers: { Authorization: `Bearer ${token}` },
  });
  return response.data;
}

export default { getEnded, getPending, getUnresolvedMilestones };
