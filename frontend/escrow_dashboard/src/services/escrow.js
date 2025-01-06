import axios from "axios";
import { BASE_URL } from "../constants";

const getStats = async (token) => {
  const response = await axios.get(`${BASE_URL}/escrow/statistics`, {
    headers: { Authorization: `Bearer ${token}` },
  });
  return response.data;
};

const reviewProject = async (token, data) => {
  await axios.post(`${BASE_URL}/escrow/approve/project`, data, {
    headers: { Authorization: `Bearer ${token}` },
  });
};

const reviewVerification = async (token, data) => {
  await axios.post(`${BASE_URL}/escrow/approve/verification`, data, {
    headers: { Authorization: `Bearer ${token}` },
  });
};

const resolveMilestone = async(token, id, payload) => {
  await axios.post(`${BASE_URL}/escrow/resolve/${id}`, payload, {
    headers: { Authorization: `Bearer ${token}` },
  });
}

const reviewProof = async (token, data) => {
  await axios.post(`${BASE_URL}/escrow/approve/proof`, data, {
    headers: { Authorization: `Bearer ${token}` },
  });
};

const reviewReport = async (token, data) => {
  await axios.post(`${BASE_URL}/escrow/review/report`, data, {
    headers: { Authorization: `Bearer ${token}` },
  });
};

const generateReport = async (token, data) => {
  const response = await axios.post(`${BASE_URL}/reports/generate`, data, {
    headers: { Authorization: `Bearer ${token}` },
    responseType: "blob"
  });

  return response.data;
};

const resolveDispute = async (token, data) => {
  await axios.post(`${BASE_URL}/escrow/resolve/dispute`, data, {
    headers: { Authorization: `Bearer ${token}` },
  });
};

export default { getStats, reviewProject, reviewVerification, resolveMilestone, reviewProof, reviewReport, generateReport, resolveDispute };
