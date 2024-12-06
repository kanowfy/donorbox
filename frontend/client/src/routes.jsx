import { Routes, Route } from "react-router-dom";
import Layout from "./pages/Layout";
import Home from "./pages/Home";
import Login from "./pages/auth/Login";
import Logout from "./pages/auth/Logout";
import Search from "./pages/project/Search";
import ViewProject from "./pages/project/ViewProject";
import Register from "./pages/auth/Register";
import SocialLogin from "./pages/auth/SocialLogin";
import RegisterSuccess from "./pages/auth/RegisterSuccess";
import Verify from "./pages/auth/Verify";
import CreateProject from "./pages/project/manage/CreateProject";
import Donate from "./pages/project/donation/Donate";
import AccountSettings from "./pages/AccountSettings";
import ProjectList from "./pages/project/manage/ProjectList";
import ManageDashboard from "./pages/project/manage/ManageDashboard";
import ManageDonations from "./pages/project/manage/ManageDonations";
import ManageUpdates from "./pages/project/manage/ManageUpdates";
import ManageLayout from "./pages/project/manage/ManageLayout";
import NotFound from "./pages/NotFound";
import ForgotPassword from "./pages/auth/ForgotPassword";
import ResetPassword from "./pages/auth/ResetPassword";
import ReportProject from "./pages/project/ReportProject";
import Payment from "./pages/project/donation/Payment";
import CheckoutForm from "./pages/project/donation/CheckoutForm";
import VerifyAccount from "./pages/VerifyAccount";
import Category from "./pages/project/Category";

const AppRoutes = () => {
  return (
    <Routes>
      <Route path="/" element={<Layout />}>
        <Route index element={<Home />} />
        <Route path="/search" element={<Search />} />
        <Route path="/category/:name" element={<Category />} />
        <Route path="/account/settings" element={<AccountSettings />} />
        <Route path="/account/verify" element={<VerifyAccount />} />
        <Route path="/fundraiser/:id" element={<ViewProject />} />
        <Route path="/start-fundraiser" element={<CreateProject />} />
        <Route path="/fundraisers" element={<ProjectList />} />
        <Route path="/fundraiser/:id/report" element={<ReportProject />} />
        <Route path="*" element={<NotFound />} />

        <Route path="/manage/:id" element={<ManageLayout />}>
          <Route index element={<ManageDashboard />} />
          <Route path="donations" element={<ManageDonations />} />
          <Route path="updates" element={<ManageUpdates />} />
        </Route>
        <Route path="/fundraiser/:id/donate" element={<Donate />} />
        <Route path="/fundraiser/:id/payment" element={<Payment />}>
          <Route path="checkout" element={<CheckoutForm />} />
        </Route>
      </Route>
      <Route path="/register" element={<Register />} />
      <Route path="/register/success" element={<RegisterSuccess />} />
      <Route path="/login" element={<Login />} />
      <Route path="/login/google" element={<SocialLogin />} />
      <Route path="/logout" element={<Logout />} />
      <Route path="/password/forgot" element={<ForgotPassword />} />
      <Route path="/password/reset" element={<ResetPassword />} />
      <Route path="/verify" element={<Verify />} />
    </Routes>
  );
};

export default AppRoutes;
