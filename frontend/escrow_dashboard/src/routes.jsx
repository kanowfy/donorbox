import { Routes, Route } from "react-router-dom";
import Login from "./pages/auth/Login";
import Logout from "./pages/auth/Logout";
import Home from "./pages/Home";
import Layout from "./pages/Layout";
import TransactionAudits from "./pages/TransactionAudits";
import ManageApplications from "./pages/ManageApplications";
import ManageMilestones from "./pages/ManageMilestones";

const AppRoutes = () => {
  return (
    <Routes>
      <Route path="/" element={<Layout />}>
        <Route index element={<Home />} />
        <Route path="manage/applications" element={<ManageApplications />} />
        <Route path="manage/milestones" element={<ManageMilestones />} />
        <Route path="transactions" element={<TransactionAudits />} />
      </Route>
      <Route path="/login" element={<Login />} />
      <Route path="/logout" element={<Logout />} />
    </Routes>
  );
};

export default AppRoutes;
