import { Outlet } from "react-router-dom";
import Header from "../components/Header";
import Footer from "../components/Footer";

const Layout = () => {
  return (
    <div>
      <div>
        <Header />
      </div>
      <div className="flex flex-col min-h-screen">
        <Outlet />
      </div>
      <div className="px-6 mt-auto">
        <Footer />
      </div>
    </div>
  );
};

export default Layout;
