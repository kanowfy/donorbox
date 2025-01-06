import axios from "axios";
import { BASE_URL } from "../constants";

const getAuditHistory = async (token) => {
  const response = await axios.get(`${BASE_URL}/audits`, {
    headers: { Authorization: `Bearer ${token}` },
  });
  return response.data;
};

export default { getAuditHistory };