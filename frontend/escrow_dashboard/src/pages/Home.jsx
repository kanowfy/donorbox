import { useNavigate } from "react-router-dom";
import { useAuthContext } from "../context/AuthContext";
import { useEffect, useState } from "react";
import InfoCard from "../components/dashboard/InfoCard";
import escrow from "../services/escrow";
import utils from "../utils/utils";
import { Bar, Doughnut, Line } from "react-chartjs-2";
import { Chart as ChartJS } from "chart.js/auto";
import { CategoryIndexMap } from "../constants";

const Home = () => {
  const navigate = useNavigate();
  const { token } = useAuthContext();
  const [stats, setStats] = useState();
  const [categories, setCategories] = useState();

  useEffect(() => {
    if (!token) {
      navigate("/login");
    }

    const fetchStats = async () => {
      try {
        const response = await escrow.getStats(token);
        setStats(response.statistics);
        console.log(response.statistics);
        setCategories(response.statistics.categories_count);
      } catch (err) {
        console.error(err);
      }
    };

    fetchStats();
  }, [token, navigate]);

  const getCount = (i) => {
    const categ = categories?.filter(c => c.id === i)[0];
    if (categ) {
      return categ.count;
    } else {
      return 0;
    }
  }

  return (
    <div className="p-10 bg-slate-200 w-full space-y-7 min-h-screen">
      <div className="grid lg:grid-cols-4 gap-7 sm:grid-cols-2">
        <InfoCard
          title="Total donated fund"
          content={`₫${stats?.total_fund.toLocaleString()}`}
        />
        <InfoCard
          title="Number of donations"
          content={`${stats?.donation_count}`}
        />
        <InfoCard
          title="Number of pending projects"
          content={`${stats?.project_count.pending}`}
        />
        <InfoCard
          title="Number of pending user verifications"
          content={`${stats?.verification_count}`}
        />
      </div>
      <div className="grid grid-cols-5 space-x-3">
        <div className=" col-span-3 h-96 min-h-60 bg-white rounded-lg p-10 flex justify-center">
          {/* Project Categories distribution */}
          <Bar
            data={{
              labels: Object.keys(CategoryIndexMap),
              datasets: [
                {
                  label: "Number of Projects by category",
                  data: [getCount(1), getCount(2), getCount(3), getCount(4), getCount(5), getCount(6), getCount(7), getCount(8), getCount(9)],
                },
              ],
            }}
          />
        </div>
        <div className=" col-span-2 h-96 min-h-60 bg-white rounded-lg p-10 flex justify-center">
          <Doughnut
            data={{
              labels: ["Pending", "Ongoing", "Finished", "Rejected"],
              datasets: [
                {
                  label: "Number of Projects",
                  data: [
                    stats?.project_count.pending,
                    stats?.project_count.ongoing,
                    stats?.project_count.finished,
                    stats?.project_count.rejected,
                  ],
                },
              ],
            }}
          />
        </div>
      </div>
      <div>
        <div className="h-96 min-h-60 bg-white rounded-lg p-10 flex flex-col items-center space-y-2">
          {/* Monthly fund release */}
          <Line
            data={{
              labels: ["January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"],
              datasets: [
                {
                  label: "Fund donated per Month",
                  data: [20000000, 30000000, 25000000, 21000000, 26000000, 15000000, 16000000, 24000000, 25000000, 27000000, 29000000, 28000000],
                },
              ],
            }}
          />
        </div>
      </div>
    </div>
  );
};

export default Home;
