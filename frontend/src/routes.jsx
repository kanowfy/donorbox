import { Routes, Route } from "react-router-dom";
import Layout from "./pages/Layout";
import Home from "./pages/Home";
import About from "./pages/About";
import Login from "./pages/Login";
import Profile from "./pages/Profile";
import Logout from "./pages/Logout";
import Search from "./pages/Search";
import Project from "./pages/Project";
import Register from "./pages/Register";
import SocialLogin from "./pages/SocialLogin";
import RegisterSuccess from "./pages/RegisterSuccess";
import Verify from "./pages/Verify";

const AppRoutes = () => {
  return (
    <Routes>
      <Route path="/" element={<Layout />}>
        <Route index element={<Home />} />
        <Route path="/search" element={<Search />} />
        <Route path="/about" element={<About />} />
        <Route path="/profile" element={<Profile />} />
        <Route path="/fundraiser/:id" element={<Project />} />
      </Route>
      <Route path="/register" element={<Register />} />
      <Route path="/register/success" element={<RegisterSuccess />} />
      <Route path="/login" element={<Login />} />
      <Route path="/login/google" element={<SocialLogin />} />
      <Route path="/logout" element={<Logout />} />
      <Route path="/verify" element={<Verify />} />
    </Routes>
  );
};

export default AppRoutes;
