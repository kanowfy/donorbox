import { Outlet } from "react-router-dom";
import Header from "../components/Header";

const Layout = () => {
  return (
    <div>
      <Header />
      <Outlet />
      <div>footer</div>
    </div>
  );
};

export default Layout;
