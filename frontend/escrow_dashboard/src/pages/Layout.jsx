import { Outlet } from "react-router-dom";
import Sidenav from "../components/Sidenav";

const Layout = () => {
  return (
    <div className="w-full grid grid-cols-12">
      <div className="col-span-2">
        <Sidenav />
      </div>
      <div className="col-span-10">
        <Outlet />
      </div>
    </div>
  );
};

export default Layout;
