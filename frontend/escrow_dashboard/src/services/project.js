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

const getFundedMilestones = async (token) => {
  const response = await axios.get(`${BASE_URL}/milestones/funded`, {
    headers: { Authorization: `Bearer ${token}` },
  });
  return response.data;
}

const getProjectReports = async (token) => {
  const response = await axios.get(`${BASE_URL}/projects/reports`, {
    headers: { Authorization: `Bearer ${token}` },
  });
  return response.data;
}

const getDisputed = async (token) => {
  const response = await axios.get(`${BASE_URL}/projects/disputed`, {
    headers: { Authorization: `Bearer ${token}` },
  });
  return response.data;
}

export default { getEnded, getPending, getFundedMilestones, getProjectReports, getDisputed };
