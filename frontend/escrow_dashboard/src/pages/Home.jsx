import { useNavigate } from "react-router-dom";
import { useAuthContext } from "../context/AuthContext";
import { useEffect } from "react";
import InfoCard from "../components/dashboard/InfoCard";
import DashboardPieChart from "../components/dashboard/DasboardPieChart";
import DashboardBarChart from "../components/dashboard/DashboardBarChart";

const Home = () => {
  const navigate = useNavigate();
  const { token } = useAuthContext();

  useEffect(() => {
    if (!token) {
      navigate("/login");
    }
  }, [token, navigate]);

  return (
    <div className="p-10 bg-slate-200 w-full space-y-7 h-screen">
      <div className="grid lg:grid-cols-4 gap-7 sm:grid-cols-2">
        <InfoCard title="Balance" content="₫50,172,000" />
        <InfoCard title="Total active Fundraisers" content="100" />
        <InfoCard title="Number of Successful Payouts" content="10" />
        <InfoCard title="Number of Refunded Projects" content="2" />
      </div>
      <div className="grid grid-cols-10 gap-7">
        <div className="col-span-4">
          <DashboardPieChart />
        </div>
        <div className="col-span-6">
          <DashboardBarChart />
        </div>
      </div>
    </div>
  );
};

export default Home;
