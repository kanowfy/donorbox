import { Routes, Route } from "react-router-dom";
import Login from "./pages/auth/Login";
import Logout from "./pages/auth/Logout";
import Home from "./pages/Home";
import Layout from "./pages/Layout";
import ManageProjectApplications from "./pages/ManageProjectApplications";
import ManageMilestones from "./pages/ManageMilestones";
import ManageUserVerifications from "./pages/ManageUserVerifications";
import ManageDocuments from "./pages/ManageDocuments";
import AuditTrails from "./pages/AuditTrails";

const AppRoutes = () => {
  return (
    <Routes>
      <Route path="/" element={<Layout />}>
        <Route index element={<Home />} />
        <Route path="manage/verifications" element={<ManageUserVerifications />} />
        <Route path="manage/projects" element={<ManageProjectApplications />} />
        <Route path="manage/milestones" element={<ManageMilestones />} />
        <Route path="manage/documents" element={<ManageDocuments />} />
        <Route path="audits" element={<AuditTrails />} />
      </Route>
      <Route path="/login" element={<Login />} />
      <Route path="/logout" element={<Logout />} />
    </Routes>
  );
};

export default AppRoutes;
