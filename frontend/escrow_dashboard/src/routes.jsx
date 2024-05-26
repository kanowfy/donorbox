import { Routes, Route } from "react-router-dom";
import Login from "./pages/auth/Login";
import Logout from "./pages/auth/Logout";
import Home from "./pages/Home";
import Layout from "./pages/Layout";
import SetupTransfer from "./pages/SetupTransfer";
import ManagePayout from "./pages/ManagePayout";
import ManageRefund from "./pages/ManageRefund";

const AppRoutes = () => {
  return (
    <Routes>
      <Route path="/" element={<Layout />}>
        <Route index element={<Home />} />
        <Route path="manage/payout" element={<ManagePayout />} />
        <Route path="manage/refund" element={<ManageRefund />} />
      </Route>
      <Route path="/login" element={<Login />} />
      <Route path="/logout" element={<Logout />} />
      <Route path="/transfer-setup" element={<SetupTransfer />} />
    </Routes>
  );
};

export default AppRoutes;
