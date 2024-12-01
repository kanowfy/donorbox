import { useNavigate } from "react-router-dom";
import { useAuthContext } from "../context/AuthContext";
import { useEffect, useState } from "react";
import InfoCard from "../components/dashboard/InfoCard";
import DashboardPieChart from "../components/dashboard/DasboardPieChart";
import escrow from "../services/escrow";
import utils from "../utils/utils";
import DashboardLineChart from "../components/dashboard/DashboardLineChart";

const Home = () => {
  const navigate = useNavigate();
  const { token } = useAuthContext();
  const [stats, setStats] = useState();
  const [transactions, setTransactions] = useState();

  useEffect(() => {
    if (!token) {
      navigate("/login");
    }

    const fetchStats = async () => {
      try {
        const response = await escrow.getStats(token);
        setStats(response.stats);
        setTransactions(response.transactions);
      } catch (err) {
        console.error(err);
      }
    };

    fetchStats();
  }, [token, navigate]);

  return (
    <div className="p-10 bg-slate-200 w-full space-y-7 min-h-screen">
      {/*
      <div className="grid lg:grid-cols-4 gap-7 sm:grid-cols-2">
        <InfoCard
          title="Total donated fund"
          content={stats?.balance && `₫${utils.formatNumber(stats?.balance)}`}
        />
        <InfoCard
          title="Number of donations"
          content={`${stats?.ongoing ?? 0}`}
        />
        <InfoCard
          title="Number of pending projects"
          content={`${stats?.completed_payout ?? 0}`}
        />
        <InfoCard
          title="Number of finished projects"
          content={`${stats?.completed_refund ?? 0}`}
        />
      </div>
      <div className="grid grid-cols-2">
        <DashboardPieChart stats={stats} />
        <DashboardLineChart transactions={transactions} />
      </div>*/}
    </div>
  );
};

export default Home;
